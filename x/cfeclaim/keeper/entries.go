package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"strconv"
)

func (k Keeper) AddUsersEntries(ctx sdk.Context, owner string, campaignId uint64, claimRecords []*types.ClaimRecord) error {
	ctx.Logger().Debug("add user entries", "owner", owner, "campaignId", campaignId, "claimRecordsLength", len(claimRecords))
	feegrantDenom := k.stakingKeeper.BondDenom(ctx)

	campaign, err := k.ValidateAddUsersEntries(ctx, owner, campaignId, claimRecords)
	if err != nil {
		return err
	}

	usersEntries, claimRecordsAmountSum, err := k.validateClaimRecords(ctx, campaign, claimRecords)
	if err != nil {
		return err
	}

	feegrantFeesSum := calculateFeegrantFeesSum(campaign.FeegrantAmount, int64(len(claimRecords)), feegrantDenom)
	feesAndClaimRecordsAmountSum := claimRecordsAmountSum.Add(feegrantFeesSum...)

	if campaign.CampaignType != types.CampaignSale {
		ownerAddress, _ := sdk.AccAddressFromBech32(owner)
		allBalances := k.bankKeeper.GetAllBalances(ctx, ownerAddress)
		ctx.Logger().Debug("add user entries", "feegrantFeesSum", feegrantFeesSum, "allBalances", allBalances)

		if campaign.CampaignType == types.CampaignTeamdrop {
			if err = k.ValidateCampaignWhenAddedFromVestingAccount(ctx, ownerAddress, campaign); err != nil {
				return err
			}
			if err = k.AddClaimRecordsFromWhitelistedVestingAccount(ctx, ownerAddress, feesAndClaimRecordsAmountSum); err != nil {
				return err
			}
		}

		if !allBalances.IsAllGTE(feesAndClaimRecordsAmountSum) {
			return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "owner balance is too small (%s < %s)", allBalances, feesAndClaimRecordsAmountSum)
		}

		err = k.setupAndSendFeegrant(ctx, ownerAddress, campaign, feegrantFeesSum, claimRecords, feegrantDenom)
		if err != nil {
			return err
		}

		if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddress, types.ModuleName, claimRecordsAmountSum); err != nil {
			return err
		}
	} else if campaign.CampaignType == types.CampaignSale {
		accountVestingPools, err := k.ValidateCampaignWhenAddedFromVestingPool(ctx, campaign)
		if err != nil {
			return err
		}

		err = k.setupAndSendFeegrantForVestingPool(ctx, campaign, feegrantFeesSum, claimRecords, feegrantDenom)
		if err != nil {
			return err
		}

		if err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, cfevestingtypes.ModuleName, types.ModuleName, claimRecordsAmountSum); err != nil {
			return err
		}

		for _, vestPool := range accountVestingPools.VestingPools {
			if vestPool.Name == campaign.VestingPoolName {
				vestPool.Sent = vestPool.Sent.Add(feesAndClaimRecordsAmountSum.AmountOf(feegrantDenom))
				break
			}
		}

		k.vestingKeeper.SetAccountVestingPools(ctx, *accountVestingPools)
	}

	k.IncrementCampaignTotalAmount(ctx, types.CampaignTotalAmount{
		CampaignId: campaignId,
		Amount:     claimRecordsAmountSum,
	})
	k.IncrementCampaignAmountLeft(ctx, types.CampaignAmountLeft{
		CampaignId: campaignId,
		Amount:     claimRecordsAmountSum,
	})
	for _, userEntry := range usersEntries {
		k.SetUserEntry(ctx, *userEntry)
	}

	event := &types.AddClaimRecords{
		Owner:                   owner,
		CampaignId:              strconv.FormatUint(campaignId, 10),
		ClaimRecordsTotalAmount: claimRecordsAmountSum.String(),
		ClaimRecordsNumber:      strconv.FormatInt(int64(len(claimRecords)), 10),
	}

	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Debug("add claim records emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func calculateFeegrantFeesSum(feegrantAmount sdk.Int, claimRecordsNumber int64, feegrantDenom string) (feesSum sdk.Coins) {
	if feegrantAmount.GT(sdk.ZeroInt()) {
		feesSum = feesSum.Add(sdk.NewCoin(feegrantDenom, feegrantAmount.MulRaw(claimRecordsNumber)))
	}
	return
}

func (k Keeper) DeleteClaimRecord(ctx sdk.Context, owner string, campaignId uint64, userAddress string, closeAction types.CloseAction) error {
	k.Logger(ctx).Debug("delete claim record", "owner", owner, "campaignId", campaignId, "userAddress", userAddress)
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return err
	}
	userEntry, claimRecordAmount, err := k.validateDeleteClaimRecord(ctx, owner, campaign, userAddress)
	if err != nil {
		return err
	}

	if err = k.delteClaimRecordCloseActionSwitch(ctx, closeAction, &campaign, claimRecordAmount); err != nil {
		return err
	}

	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		granteeAddr, err := sdk.AccAddressFromBech32(userEntry.Address)
		if err != nil {
			return err
		}
		_, accountAddr := FeegrantAccountAddress(campaignId)
		_, err = k.feeGrantKeeper.GetAllowance(ctx, accountAddr, granteeAddr)
		if err != nil {
			k.revokeFeeAllowance(ctx, accountAddr, granteeAddr)
		}
	}

	k.SetUserEntry(ctx, userEntry)
	k.DecrementCampaignTotalAmount(ctx, campaignId, claimRecordAmount)

	event := &types.DeleteClaimRecord{
		Owner:             owner,
		CampaignId:        strconv.FormatUint(campaignId, 10),
		UserAddress:       userAddress,
		ClaimRecordAmount: claimRecordAmount.String(),
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Debug("delete claim record emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) delteClaimRecordCloseActionSwitch(ctx sdk.Context, CloseAction types.CloseAction, campaign *types.Campaign, campaignAmountLeft sdk.Coins) error {
	switch CloseAction {
	case types.CloseSendToCommunityPool:
		return k.delteClaimRecordCloseSendToCommunityPool(ctx, campaignAmountLeft)
	case types.CampaignCloseBurn:
		return k.delteClaimRecordCloseBurn(ctx, campaignAmountLeft)
	case types.CampaignCloseSendToOwner:
		return k.delteClaimRecordCloseSendToOwner(ctx, campaign, campaignAmountLeft)
	default:
		return errors.Wrap(sdkerrors.ErrInvalidType, "wrong campaign close action type")
	}
}

func (k Keeper) delteClaimRecordCloseSendToCommunityPool(ctx sdk.Context, campaignAmountLeft sdk.Coins) error {
	return k.distributionKeeper.FundCommunityPool(ctx, campaignAmountLeft, authtypes.NewModuleAddress(types.ModuleName))
}

func (k Keeper) delteClaimRecordCloseBurn(ctx sdk.Context, coinsToBurn sdk.Coins) error {
	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, coinsToBurn)
}

func (k Keeper) delteClaimRecordCloseSendToOwner(ctx sdk.Context, campaign *types.Campaign, campaignAmountLeft sdk.Coins) error {
	if campaign.CampaignType != types.CampaignSale {
		ownerAddress, _ := sdk.AccAddressFromBech32(campaign.Owner)
		return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, campaignAmountLeft)
	} else {
		vestingDenom := k.vestingKeeper.Denom(ctx)
		accountVestingPools, _ := k.vestingKeeper.GetAccountVestingPools(ctx, campaign.Owner)
		var vestingPool *cfevestingtypes.VestingPool
		for _, vestPool := range accountVestingPools.VestingPools {
			if vestPool.Name == campaign.VestingPoolName {
				vestingPool = vestPool
				break
			}
		}

		if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, cfevestingtypes.ModuleName, campaignAmountLeft); err != nil {
			return err
		}

		vestingPool.Sent = vestingPool.Sent.Add(campaignAmountLeft.AmountOf(vestingDenom))
		k.vestingKeeper.SetAccountVestingPools(ctx, accountVestingPools)
	}
	return nil
}

func (k Keeper) ValidateAddUsersEntries(ctx sdk.Context, owner string, campaignId uint64, claimRecords []*types.ClaimRecord) (*types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return nil, err
	}

	if err = ValidateOwner(campaign, owner); err != nil {
		return nil, err
	}
	if err = types.ValidateClaimRecords(claimRecords); err != nil {
		return nil, err
	}

	if campaign.CampaignType != types.CampaignSale {
		if err = ValidateCampaignIsNotDisabled(campaign); err != nil {
			return nil, err
		}
	}

	if err = ValidateCampaignNotEnded(ctx, campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (k Keeper) ValidateCampaignWhenAddedFromVestingAccount(ctx sdk.Context, ownerAcc sdk.AccAddress, campaign *types.Campaign) error {
	ownerAccount := k.accountKeeper.GetAccount(ctx, ownerAcc)
	if ownerAccount == nil {
		return errors.Wrapf(c4eerrors.ErrNotExists, "account %s doesn't exist", ownerAcc)
	}
	vestingAcc := ownerAccount.(*vestingtypes.ContinuousVestingAccount)
	vestingAccTimeDiff := vestingAcc.EndTime - vestingAcc.StartTime
	campaignLockupAndVestingSum := int64(campaign.VestingPeriod.Seconds() + campaign.LockupPeriod.Seconds())
	if campaignLockupAndVestingSum < vestingAccTimeDiff {
		return errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the duration of vesting and lockup must be equal to or greater than the remaining vesting time of the vesting account (%d < %d)", campaignLockupAndVestingSum, vestingAccTimeDiff))
	}
	return nil
}

func (k Keeper) ValidateCampaignWhenAddedFromVestingPool(ctx sdk.Context, campaign *types.Campaign) (*cfevestingtypes.AccountVestingPools, error) {
	accountVestingPools, _ := k.vestingKeeper.GetAccountVestingPools(ctx, campaign.Owner)
	var vestingPool *cfevestingtypes.VestingPool
	for _, vestPool := range accountVestingPools.VestingPools {
		if vestPool.Name == campaign.VestingPoolName {
			vestingPool = vestPool
			break
		}
	}

	vestingType, err := k.vestingKeeper.GetVestingType(ctx, vestingPool.VestingType)
	if err != nil {
		return nil, err
	}

	if vestingType.LockupPeriod > campaign.LockupPeriod {
		return nil, errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the duration of campaign lockup period must be equal to or greater than the vesting type lockup period (%d > %d)", vestingType.LockupPeriod, campaign.LockupPeriod))
	}

	if vestingType.VestingPeriod > campaign.VestingPeriod {
		return nil, errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the duration of campaign vesting period must be equal to or greater than the vesting type vesting period (%d > %d)", vestingType.VestingPeriod, campaign.VestingPeriod))
	}
	return &accountVestingPools, nil
}

func (k Keeper) validateClaimRecords(ctx sdk.Context, campaign *types.Campaign, claimRecords []*types.ClaimRecord) (usersEntries []*types.UserEntry, entriesAmountSum sdk.Coins, err error) {
	allCampaignMissions, _ := k.AllMissionForCampaign(ctx, campaign.Id)

	for i, claimRecord := range claimRecords {
		if err = k.validateInitialClaimFreeAmount(campaign.InitialClaimFreeAmount, allCampaignMissions, claimRecord); err != nil {
			return nil, nil, types.WrapClaimRecordIndex(err, i)
		}

		entriesAmountSum = entriesAmountSum.Add(claimRecord.Amount...)

		userEntry, err := k.addUserEntry(ctx, campaign.Id, claimRecord.Address, claimRecord.Amount)
		if err != nil {
			return nil, nil, types.WrapClaimRecordIndex(err, i)
		}
		usersEntries = append(usersEntries, userEntry)
	}
	return
}

func (k Keeper) validateInitialClaimFreeAmount(initialClaimFreeAmount sdk.Int, missions []types.Mission, claimRecord *types.ClaimRecord) error {
	allMissionsAmountSum := sdk.NewCoins()
	for _, mission := range missions {
		for _, amount := range claimRecord.Amount {
			allMissionsAmountSum = allMissionsAmountSum.Add(sdk.NewCoin(amount.Denom, mission.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()))
		}
	}
	initialClaimAmount := claimRecord.Amount.Sub(allMissionsAmountSum...)

	for _, coin := range claimRecord.Amount {
		if initialClaimAmount.AmountOf(coin.Denom).LT(initialClaimFreeAmount) {
			return errors.Wrapf(c4eerrors.ErrParam, "claim amount %s < campaign initial claim free amount (%s)", initialClaimAmount.AmountOf(coin.Denom), initialClaimFreeAmount.String())
		}
	}

	return nil
}

func (k Keeper) addUserEntry(ctx sdk.Context, campaignId uint64, address string, allCoins sdk.Coins) (*types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(ctx, address)
	if !found {
		userEntry = types.UserEntry{Address: address}
	}
	if userEntry.HasCampaign(campaignId) {
		return nil, errors.Wrapf(c4eerrors.ErrAlreadyExists, "campaignId %d already exists for address: %s", campaignId, address)
	}
	userEntry.ClaimRecords = append(userEntry.ClaimRecords, &types.ClaimRecord{CampaignId: campaignId, Amount: allCoins})
	return &userEntry, nil
}

func (k Keeper) setupAndSendFeegrant(ctx sdk.Context, ownerAcc sdk.AccAddress, campaign *types.Campaign, feegrantFeesSum sdk.Coins, claimRecords []*types.ClaimRecord, feegrantDenom string) error {
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		acc := k.NewModuleAccountSet(ctx, campaign.Id)
		if err := k.bankKeeper.SendCoins(ctx, ownerAcc, acc.GetAddress(), feegrantFeesSum); err != nil {
			return err
		}
		if err := k.grantFeeAllowanceToAllClaimRecords(ctx, acc.GetAddress(), claimRecords, sdk.NewCoins(sdk.NewCoin(feegrantDenom, campaign.FeegrantAmount))); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) setupAndSendFeegrantForVestingPool(ctx sdk.Context, campaign *types.Campaign, feegrantFeesSum sdk.Coins, claimRecords []*types.ClaimRecord, feegrantDenom string) error {
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		acc := k.NewModuleAccountSet(ctx, campaign.Id)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, cfevestingtypes.ModuleName, acc.GetAddress(), feegrantFeesSum); err != nil {
			return err
		}
		if err := k.grantFeeAllowanceToAllClaimRecords(ctx, acc.GetAddress(), claimRecords, sdk.NewCoins(sdk.NewCoin(feegrantDenom, campaign.FeegrantAmount))); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) grantFeeAllowanceToAllClaimRecords(ctx sdk.Context, moduleAddress sdk.AccAddress, claimEntries []*types.ClaimRecord, grantAmount sdk.Coins) error {
	basicAllowance, err := codectypes.NewAnyWithValue(&feegranttypes.BasicAllowance{
		SpendLimit: grantAmount,
	})
	if err != nil {
		return err
	}

	allowedMsgAllowance := feegranttypes.AllowedMsgAllowance{
		Allowance:       basicAllowance,
		AllowedMessages: []string{"/chain4energy.c4echain.cfeclaim.MsgInitialClaim"},
	}

	for _, claimRecord := range claimEntries {
		granteeAddress, _ := sdk.AccAddressFromBech32(claimRecord.Address)
		existingFeeAllowance, _ := k.feeGrantKeeper.GetAllowance(ctx, moduleAddress, granteeAddress)
		if existingFeeAllowance == nil {
			if err = k.feeGrantKeeper.GrantAllowance(ctx, moduleAddress, granteeAddress, &allowedMsgAllowance); err != nil {
				return err
			}
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

func (k Keeper) AddClaimRecordsFromWhitelistedVestingAccount(ctx sdk.Context, ownerAddress sdk.AccAddress, amount sdk.Coins) error {
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, ownerAddress)
	if spendableCoins.IsAllGTE(amount) {
		return nil
	}

	ownerAccount := k.accountKeeper.GetAccount(ctx, ownerAddress)
	if ownerAccount == nil {
		return errors.Wrapf(c4eerrors.ErrNotExists, "account %s doesn't exist", ownerAddress)
	}
	vestingAcc := ownerAccount.(*vestingtypes.ContinuousVestingAccount)

	if vestingAcc.DelegatedVesting.IsAllPositive() {
		balance := k.bankKeeper.GetAllBalances(ctx, ownerAddress)
		balanceWithoutOriginalVesting := balance.Add(vestingAcc.DelegatedVesting...).Sub(vestingAcc.OriginalVesting...)
		amount = amount.Add(balanceWithoutOriginalVesting...)
	}

	amount = amount.Sub(spendableCoins...)

	_, err := k.vestingKeeper.UnlockUnbondedContinuousVestingAccountCoins(ctx, ownerAddress, amount)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) validateDeleteClaimRecord(ctx sdk.Context, owner string, campaign types.Campaign, userAddress string) (types.UserEntry, sdk.Coins, error) {
	if err := ValidateCampaignTypeIsTeamdrop(campaign); err != nil {
		return types.UserEntry{}, nil, err
	}

	if err := ValidateOwner(campaign, owner); err != nil {
		return types.UserEntry{}, nil, err
	}

	userEntry, err := k.ValidateUserEntry(ctx, userAddress)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	claimRecordAmount, err := ValidateClaimRecordExists(userEntry, campaign.Id)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	return userEntry, claimRecordAmount, nil
}

func (k Keeper) ValidateUserEntry(ctx sdk.Context, userAddress string) (types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(
		ctx,
		userAddress,
	)
	if !found {
		return types.UserEntry{}, errors.Wrapf(c4eerrors.ErrParsing, "userEntry doesn't exist")
	}

	return userEntry, nil
}

func ValidateClaimRecordExists(userEntry types.UserEntry, campaignId uint64) (claimRecordAmount sdk.Coins, err error) {
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
		return nil, errors.Wrapf(c4eerrors.ErrParsing, "campaign id %d claim entry doesn't exist", campaignId)
	}

	return claimRecordAmount, nil
}

func ValidateCampaignTypeIsTeamdrop(campaign types.Campaign) error {
	if campaign.CampaignType != types.CampaignTeamdrop {
		return errors.Wrap(sdkerrors.ErrInvalidType, "ampaign must be of TEAMDROP type to be able to delete its entries")
	}
	return nil
}
