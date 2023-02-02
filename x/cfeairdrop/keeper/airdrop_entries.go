package keeper

import (
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/exp/slices"
	"strconv"
)

func (k Keeper) AddUsersEntries(ctx sdk.Context, owner string, campaignId uint64, airdropEntries []*types.ClaimRecord) error {
	ownerAddress, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Debug("add campaign entries owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "add campaign entries  - owner parsing error: %s", owner).Error())
	}
	campaign, found := k.GetCampaign(
		ctx,
		campaignId,
	)
	if !found {
		k.Logger(ctx).Debug("add campaign entries campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "add campaign entries -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Debug("add campaign entries you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrorInvalidSigner, "add campaign entries - you are not the owner of campaign with id %d", campaignId)
	}
	if campaign.Enabled == false {
		k.Logger(ctx).Debug("add campaign entries campaign is disabled", "campaignId", campaignId)
		return sdkerrors.Wrapf(types.ErrCampaignDisabled, "add campaign entries - campaign %d is disabled", campaignId)
	}
	if campaign.EndTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Debug("add campaign entries campaign is disabled", "campaignId", campaignId)
		return sdkerrors.Wrapf(types.ErrCampaignDisabled, fmt.Sprintf("add campaign entries - campaign %d is disabled (end time %s < %s)", campaignId, campaign.EndTime, ctx.BlockTime()))
	}
	var usersEntries []*types.UserEntry
	entriesAmountSum := sdk.NewCoins()
	allCampaignMissions, _ := k.AllMissionForCampaign(ctx, campaignId)
	for i, claimRecord := range airdropEntries {
		if claimRecord.Address == "" {
			k.Logger(ctx).Error("add campaign entries airdrop entry empty address", "claimRecordIndex", i)
			return sdkerrors.Wrapf(c4eerrors.ErrParam, "add campaign entries - airdrop entry empty address on index %d", i)
		}
		if len(claimRecord.Amount) == 0 {
			k.Logger(ctx).Error("add campaign entries airdrop entry must has at least one coin")
			return sdkerrors.Wrapf(c4eerrors.ErrParam, "add campaign entries - airdrop entry at index %d airdrop entry must has at least one coin", i)
		}

		allMissionsAmountSum := sdk.NewCoins()
		for _, mission := range allCampaignMissions {
			for _, amount := range claimRecord.Amount {
				allMissionsAmountSum = allMissionsAmountSum.Add(sdk.NewCoin(amount.Denom, mission.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()))
			}
		}
		initialClaimClaimable := claimRecord.Amount.Sub(allMissionsAmountSum)
		for _, coin := range claimRecord.Amount {
			if initialClaimClaimable.AmountOf(coin.Denom).LT(campaign.InitialClaimFreeAmount) {
				k.Logger(ctx).Error("add campaign entries airdrop entry amount < campaign initial claim free amount", "amount", coin.Amount, "initialClaimFreeAmount", campaign.InitialClaimFreeAmount, "claimRecordIndex", i)
				return sdkerrors.Wrapf(c4eerrors.ErrParam, "add campaign entries - airdrop entry at index %d initial claim amount %s < campaign initial claim free amount (%s)", i, initialClaimClaimable.AmountOf(coin.Denom), campaign.InitialClaimFreeAmount.String())
			}
			if coin.Amount.Equal(sdk.ZeroInt()) {
				k.Logger(ctx).Error("add campaign entries airdrop entry amount is 0", "amount", coin.Amount, "initialClaimFreeAmount", campaign.InitialClaimFreeAmount, "claimRecordIndex", i)
				return sdkerrors.Wrapf(c4eerrors.ErrParam, "add campaign entries - airdrop entry at index %d amount is 0", i)
			}
			entriesAmountSum = entriesAmountSum.Add(coin)
		}

		userEntry, err := k.addUserEntry(ctx, campaignId, claimRecord.Address, claimRecord.Amount)
		if err != nil {
			return err
		}
		usersEntries = append(usersEntries, userEntry)
	}

	feesAndEntriesSum := entriesAmountSum
	feesSum := sdk.NewCoins()

	bondedDenom := "uc4e"
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		feesSum = sdk.NewCoins(sdk.NewCoin(bondedDenom, campaign.FeegrantAmount.MulRaw(int64(len(airdropEntries)))))
		feesAndEntriesSum = feesAndEntriesSum.Add(feesSum...)
	}

	allBalances := k.bankKeeper.GetAllBalances(ctx, ownerAddress)
	if !allBalances.IsAllGTE(feesAndEntriesSum) {
		k.Logger(ctx).Error("add campaign entries airdrop entry owner balance is too small")
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, fmt.Sprintf("add campaign entries - owner balance is too small (%s < %s)", allBalances, feesAndEntriesSum))
	}

	if slices.Contains(types.GetWhitelistedVestingAccounts(), owner) {
		if err = k.AddClaimRecordsFromWhitelistedVestingAccount(ctx, owner, feesAndEntriesSum); err != nil {
			return err
		}
	}

	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		acc := k.NewModuleAccountSet(ctx, campaignId)
		if err = k.bankKeeper.SendCoins(ctx, ownerAddress, acc.GetAddress(), feesSum); err != nil {
			return err
		}
		if err = k.grantAllFeeAllowance(ctx, acc.GetAddress(), airdropEntries, sdk.NewCoins(sdk.NewCoin(bondedDenom, campaign.FeegrantAmount))); err != nil {
			return err
		}
	}

	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddress, types.ModuleName, entriesAmountSum); err != nil {
		return err
	}

	k.IncrementCampaignTotalAmount(ctx, types.CampaignTotalAmount{
		CampaignId: campaignId,
		Amount:     entriesAmountSum,
	})
	k.IncrementCampaignAmountLeft(ctx, types.CampaignAmountLeft{
		CampaignId: campaignId,
		Amount:     entriesAmountSum,
	})
	for _, userEntry := range usersEntries {
		k.SetUserEntry(ctx, *userEntry)
	}
	return nil
}

func (k Keeper) addUserEntry(ctx sdk.Context, campaignId uint64, address string, allCoins sdk.Coins) (*types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(ctx, address)
	if !found {
		userEntry = types.UserEntry{Address: address}
	}
	if userEntry.HasCampaign(campaignId) {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrAlreadyExists, "campaignId %d already exists for address: %s", campaignId, address)
	}
	userEntry.ClaimRecords = append(userEntry.ClaimRecords, &types.ClaimRecord{CampaignId: campaignId, Amount: allCoins})
	return &userEntry, nil
}

func (k Keeper) DeleteClaimRecord(ctx sdk.Context, owner string, campaignId uint64, userAddress string) error {
	logger := ctx.Logger().With("delete claim record", "owner", owner, "campaignId", campaignId, "userAddress", userAddress)

	userEntry, claimRecordAmount, validationResult := k.ValidateRemoveClaimRecord(logger, ctx, owner, campaignId, userAddress)
	if validationResult != nil {
		return validationResult
	}

	k.SetUserEntry(ctx, userEntry)
	k.DecrementCampaignTotalAmount(ctx, campaignId, claimRecordAmount)

	return nil
}

func (k Keeper) grantAllFeeAllowance(ctx sdk.Context, moduleAddress sdk.AccAddress, airdropEntries []*types.ClaimRecord, grantAmount sdk.Coins) error {
	basicAllowance, err := codectypes.NewAnyWithValue(&feegranttypes.BasicAllowance{
		SpendLimit: grantAmount,
	})
	if err != nil {
		return err // TODO
	}
	allowedMsgAllowance := feegranttypes.AllowedMsgAllowance{
		Allowance:       basicAllowance,
		AllowedMessages: []string{"/chain4energy.c4echain.cfeairdrop.MsgInitialClaim"},
	}
	for _, claimRecord := range airdropEntries {
		granteeAddress, err := sdk.AccAddressFromBech32(claimRecord.Address)
		if err != nil {
			return err // TODO
		}
		feeAllowance, _ := k.feeGrantKeeper.GetAllowance(ctx, moduleAddress, granteeAddress)
		if feeAllowance != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
		}

		if err = k.feeGrantKeeper.GrantAllowance(ctx, moduleAddress, granteeAddress, &allowedMsgAllowance); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) revokeFeeAllowance(ctx sdk.Context, granter sdk.Address, grantee sdk.AccAddress) error {
	keeper, _ := (k.feeGrantKeeper).(feegrantkeeper.Keeper)
	feegrantMsgServer := feegrantkeeper.NewMsgServerImpl(keeper)
	msg := feegranttypes.MsgRevokeAllowance{
		Granter: granter.String(),
		Grantee: grantee.String(),
	}
	_, err := feegrantMsgServer.RevokeAllowance(sdk.WrapSDKContext(ctx), &msg)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) NewModuleAccountSet(ctx sdk.Context, campaignId uint64) *authtypes.ModuleAccount {
	moduleAddressName, accountAddr := FeegrantAccountAddress(campaignId)
	macc := &authtypes.ModuleAccount{
		BaseAccount: &authtypes.BaseAccount{
			Address: accountAddr.String(),
		},
		Name: moduleAddressName,
	}
	k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccount(ctx, macc))
	return macc
}

func FeegrantAccountAddress(campaignId uint64) (string, sdk.AccAddress) {
	moduleAddressName := "fee-grant-" + strconv.FormatUint(campaignId, 10)
	return moduleAddressName, authtypes.NewModuleAddress(moduleAddressName)
}

func (k Keeper) AddClaimRecordsFromWhitelistedVestingAccount(ctx sdk.Context, from string, amount sdk.Coins) error {
	fromAddress, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		k.Logger(ctx).Error("delete user airdrop entry owner parsing error", "from", from, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "delete user airdrop entry - owner parsing error: %s", from).Error())
	}
	ak := k.accountKeeper
	bk := k.bankKeeper
	fromAcc := ak.GetAccount(ctx, fromAddress)
	if fromAcc == nil {
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "account %s doesn't exist", fromAddress)
	}

	spendableCoins := bk.SpendableCoins(ctx, fromAddress)
	if spendableCoins.IsAllGTE(amount) {
		return nil
	}

	vestingAcc := fromAcc.(*vestingtypes.ContinuousVestingAccount)

	balance := bk.GetAllBalances(ctx, fromAddress)
	if balance.IsAllGT(vestingAcc.OriginalVesting) {
		balanceWithoutOriginalVesting := balance.Sub(vestingAcc.OriginalVesting)
		amount = amount.Sub(balanceWithoutOriginalVesting)
	}

	lockedCoins := vestingAcc.LockedCoins(ctx.BlockTime()).Add(vestingAcc.DelegatedVesting...)
	spendableFromVesting := vestingAcc.OriginalVesting.Sub(lockedCoins)
	amountDiffFree := amount.Sub(spendableFromVesting)

	for _, coin := range amount {
		lockedPercentage := vestingAcc.OriginalVesting.AmountOf(coin.Denom).ToDec().Quo(lockedCoins.AmountOf(coin.Denom).ToDec())
		originalVestingDiff := amountDiffFree.AmountOf(coin.Denom).ToDec().Mul(lockedPercentage).TruncateInt()
		vestingAcc.OriginalVesting = vestingAcc.OriginalVesting.Sub(sdk.NewCoins(sdk.NewCoin(coin.Denom, originalVestingDiff)))
	}

	ak.SetAccount(ctx, vestingAcc)
	spendableCoins = bk.SpendableCoins(ctx, fromAddress)
	return nil
}

func (k Keeper) ValidateRemoveClaimRecord(logger log.Logger, ctx sdk.Context, owner string, campaignId uint64, userAddress string) (types.UserEntry, sdk.Coins, error) {
	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	if err = ValidateCampaignTypeIsTeamdrop(logger, campaign); err != nil {
		return types.UserEntry{}, nil, err
	}

	if err = ValidateOwner(logger, campaign, owner); err != nil {
		return types.UserEntry{}, nil, err
	}

	userEntry, err := k.ValidateUserEntry(logger, ctx, userAddress)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	claimRecordAmount, err := ValidateClaimRecordExists(logger, userEntry, campaignId)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	return userEntry, claimRecordAmount, nil
}

func (k Keeper) ValidateUserEntry(log log.Logger, ctx sdk.Context, userAddress string) (types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(
		ctx,
		userAddress,
	)
	if !found {
		log.Debug("delete user airdrop entry userEntry doesn't exist")
		return types.UserEntry{}, sdkerrors.Wrapf(c4eerrors.ErrParsing, "delete user airdrop entry - userEntry doesn't exist")
	}

	return userEntry, nil
}

func ValidateClaimRecordExists(log log.Logger, userEntry types.UserEntry, campaignId uint64) (claimRecordAmount sdk.Coins, err error) {
	claimRecordFound := false
	for i, claimRecord := range userEntry.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			claimRecordFound = true
			claimRecordAmount = claimRecord.Amount
			userEntry.ClaimRecords = append(userEntry.ClaimRecords[:i], userEntry.ClaimRecords[i+1:]...)
			break
		}
	}
	if !claimRecordFound {
		log.Debug("delete user airdrop entry airdrop entry doesn't exist", "campaignId", campaignId)
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParsing, "delete user airdrop entry -  campaign id %d airdrop entry doesn't exist", campaignId)
	}

	return claimRecordAmount, nil
}

func ValidateCampaignTypeIsTeamdrop(log log.Logger, campaign types.Campaign) error {
	if campaign.CampaignType != types.CampaignTeamdrop {
		log.Debug("campaign must be of TEAMDROP type to be able to delete its entries")
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "ampaign must be of TEAMDROP type to be able to delete its entries")
	}
	return nil
}
