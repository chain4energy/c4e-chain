package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"golang.org/x/exp/slices"
	"strconv"
	"time"
)

func (k Keeper) AddClaimRecords(ctx sdk.Context, owner string, campaignId uint64, claimRecords []*types.ClaimRecord) error {
	ctx.Logger().Debug("add user entries", "owner", owner, "campaignId", campaignId, "claimRecordsLength", len(claimRecords))
	feegrantDenom := k.stakingKeeper.BondDenom(ctx)
	vestingDenom := k.vestingKeeper.Denom(ctx)

	campaign, err := k.ValidateAddClaimRecords(ctx, owner, campaignId, claimRecords)
	if err != nil {
		return err
	}

	usersEntries, amountSum, err := k.validateClaimRecords(ctx, campaign, claimRecords)
	if err != nil {
		return err
	}

	feegrantFeesSum := calculateFeegrantFeesSum(campaign.FeegrantAmount, int64(len(claimRecords)), feegrantDenom)
	feesAndClaimRecordsAmountSum := amountSum.Add(feegrantFeesSum...)
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)

	if campaign.CampaignType == types.VestingPoolCampaign {
		if err = k.vestingKeeper.AddVestingPoolReservation(ctx, owner, campaign.VestingPoolName, campaignId, amountSum.AmountOf(vestingDenom)); err != nil {
			return err
		}
	} else {
		if err = k.addClaimRecordsToDefaultAndDynamicCampaign(ctx, ownerAddress, campaign, amountSum, feesAndClaimRecordsAmountSum); err != nil {
			return err
		}
	}

	err = k.setupAndSendFeegrant(ctx, ownerAddress, campaign, feegrantFeesSum, claimRecords, feegrantDenom)
	if err != nil {
		return err
	}

	k.IncrementCampaignTotalAmount(ctx, types.CampaignTotalAmount{
		CampaignId: campaignId,
		Amount:     amountSum,
	})
	k.IncrementCampaignAmountLeft(ctx, types.CampaignAmountLeft{
		CampaignId: campaignId,
		Amount:     amountSum,
	})
	for _, userEntry := range usersEntries {
		k.SetUserEntry(ctx, *userEntry)
	}

	event := &types.AddClaimRecords{
		Owner:                   owner,
		CampaignId:              strconv.FormatUint(campaignId, 10),
		ClaimRecordsTotalAmount: amountSum.String(),
		ClaimRecordsNumber:      strconv.FormatInt(int64(len(claimRecords)), 10),
	}

	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Debug("add claim records emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) addClaimRecordsToVestingPoolCampaign(ctx sdk.Context, owner string, campaign *types.Campaign, amountSum sdk.Coins) error {
	accountVestingPools, found := k.vestingKeeper.GetAccountVestingPools(ctx, owner)
	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pools not found for address %s", owner)
	}

	for _, vestPool := range accountVestingPools.VestingPools {
		if vestPool.Name == campaign.VestingPoolName {
			vestPool.Sent = vestPool.Sent.Add(amountSum.AmountOf(k.vestingKeeper.Denom(ctx)))
			break
		}
	}

	k.vestingKeeper.SetAccountVestingPools(ctx, accountVestingPools)
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, cfevestingtypes.ModuleName, types.ModuleName, amountSum)
}

func (k Keeper) addClaimRecordsToDefaultAndDynamicCampaign(ctx sdk.Context, ownerAddress sdk.AccAddress, campaign *types.Campaign, amountSum sdk.Coins, feesAndClaimRecordsAmountSum sdk.Coins) error {
	allBalances := k.bankKeeper.GetAllBalances(ctx, ownerAddress)

	if slices.Contains(types.GetWhitelistedVestingAccounts(), campaign.Owner) { // TODO: probably to delete
		if err := k.ValidateCampaignWhenAddedFromVestingAccount(ctx, ownerAddress, campaign); err != nil {
			return err
		}
		if err := k.AddClaimRecordsFromWhitelistedVestingAccount(ctx, ownerAddress, amountSum); err != nil {
			return err
		}
		allBalances = k.bankKeeper.GetAllBalances(ctx, ownerAddress)
	}

	if !allBalances.IsAllGTE(feesAndClaimRecordsAmountSum) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "owner balance is too small (%s < %s)", allBalances, feesAndClaimRecordsAmountSum)
	}

	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddress, types.ModuleName, amountSum)
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

	if err = k.closeActionSwitch(ctx, closeAction, &campaign, claimRecordAmount); err != nil {
		return err
	}

	_ = k.deleteClaimRecordSendFeegrant(ctx, closeAction, &campaign, userAddress)

	for i, claimRecord := range userEntry.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			userEntry.ClaimRecords = append(userEntry.ClaimRecords[:i], userEntry.ClaimRecords[i+1:]...)
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

func (k Keeper) ValidateAddClaimRecords(ctx sdk.Context, owner string, campaignId uint64, claimRecords []*types.ClaimRecord) (*types.Campaign, error) {
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

	if campaign.CampaignType != types.VestingPoolCampaign {
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

func (k Keeper) ValidateCampaignWhenAddedFromVestingPool(ctx sdk.Context, owner string, vestingPoolName string,
	campaignLockupPeriod *time.Duration, campaignVestingPeriod *time.Duration) error {
	vestingPool, found := k.vestingKeeper.GetAccountVestingPool(ctx, owner, vestingPoolName)

	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pool %s not found for address %s", vestingPoolName, owner)
	}

	vestingType, err := k.vestingKeeper.GetVestingType(ctx, vestingPool.VestingType)
	if err != nil {
		return err
	}

	return validateVestingTimesForCampaign(vestingType, campaignLockupPeriod, campaignVestingPeriod)
}

func validateVestingTimesForCampaign(vestingType cfevestingtypes.VestingType, campaignLockupPeriod *time.Duration, campaignVestingPeriod *time.Duration) error {
	if vestingType.LockupPeriod > *campaignLockupPeriod {
		return errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the duration of campaign lockup period must be equal to or greater than the vesting type lockup period (%s > %s)", vestingType.LockupPeriod.String(), campaignLockupPeriod.String()))
	}

	if vestingType.VestingPeriod > *campaignVestingPeriod {
		return errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the duration of campaign vesting period must be equal to or greater than the vesting type vesting period (%s > %s)", vestingType.VestingPeriod.String(), campaignVestingPeriod.String()))
	}
	return nil
}

func (k Keeper) validateClaimRecords(ctx sdk.Context, campaign *types.Campaign, claimRecords []*types.ClaimRecord) (usersEntries []*types.UserEntry, entriesAmountSum sdk.Coins, err error) {
	allCampaignMissions, _ := k.AllMissionForCampaign(ctx, campaign.Id)

	for i, claimRecord := range claimRecords {
		if err = k.validateInitialClaimFreeAmount(campaign.InitialClaimFreeAmount, allCampaignMissions, claimRecord); err != nil {
			return nil, nil, types.WrapClaimRecordIndex(err, i)
		}

		entriesAmountSum = entriesAmountSum.Add(claimRecord.Amount...)

		userEntry, err := k.addClaimRecordToUserEntry(ctx, campaign.Id, claimRecord.Address, claimRecord.Amount)
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

func (k Keeper) addClaimRecordToUserEntry(ctx sdk.Context, campaignId uint64, address string, allCoins sdk.Coins) (*types.UserEntry, error) {
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
	moduleAddressName, accountAddr := CreateFeegrantAccountAddress(campaignId)
	macc := &authtypes.ModuleAccount{
		BaseAccount: &authtypes.BaseAccount{
			Address: accountAddr.String(),
		},
		Name: moduleAddressName,
	}
	k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccount(ctx, macc))
	return macc
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
	if err := ValidateCampaignTypeIsDynamic(campaign); err != nil {
		return types.UserEntry{}, nil, err
	}

	if err := ValidateOwner(campaign, owner); err != nil {
		return types.UserEntry{}, nil, err
	}

	userEntry, err := k.ValidateUserEntry(ctx, userAddress)
	if err != nil {
		return types.UserEntry{}, nil, err
	}

	claimRecord, err := ValidateClaimRecordExists(userEntry, campaign.Id)
	if err != nil {
		return types.UserEntry{}, nil, err
	}
	var amount sdk.Coins

	for _, claimedMissionId := range claimRecord.ClaimedMissions {
		mission, found := k.GetMission(ctx, campaign.Id, claimedMissionId)
		if !found {
			return types.UserEntry{}, nil, errors.Wrapf(sdkerrors.ErrNotFound, "mission with id %d not found", claimedMissionId)
		}
		for _, coin := range claimRecord.Amount {
			weightedAmount := sdk.NewDecFromInt(coin.Amount).Mul(mission.Weight).TruncateInt()
			amount = amount.Add(sdk.NewCoin(coin.Denom, weightedAmount))
		}
	}

	return userEntry, amount, nil
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

func ValidateClaimRecordExists(userEntry types.UserEntry, campaignId uint64) (claimRecordAmount *types.ClaimRecord, err error) {
	for _, claimRecord := range userEntry.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			return claimRecord, nil
		}
	}

	return nil, errors.Wrapf(c4eerrors.ErrParsing, "campaign id %d claim entry doesn't exist", campaignId)
}

func ValidateCampaignTypeIsDynamic(campaign types.Campaign) error {
	if campaign.CampaignType != types.DynamicCampaign {
		return errors.Wrap(sdkerrors.ErrInvalidType, "ampaign must be of DYNAMIC type to be able to delete its entries")
	}
	return nil
}
