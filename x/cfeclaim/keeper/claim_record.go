package keeper

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"strconv"
	"time"
)

func (k Keeper) AddClaimRecords(ctx sdk.Context, owner string, campaignId uint64, claimRecords []*types.ClaimRecord) error {
	k.Logger(ctx).Debug("add user entries", "owner", owner, "campaignId", campaignId, "claimRecordsLength", len(claimRecords))
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
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)

	if campaign.CampaignType == types.VestingPoolCampaign {
		balances := k.bankKeeper.GetAllBalances(ctx, ownerAddress)
		if feegrantFeesSum.IsAnyGT(balances) {
			return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "owner balance is too small (%s < %s)", balances, feegrantFeesSum)
		}
		if err = k.vestingKeeper.AddVestingPoolReservation(ctx, owner, campaign.VestingPoolName, campaignId, amountSum.AmountOf(vestingDenom)); err != nil {
			return err
		}
	} else {
		feesAndClaimRecordsAmountSum := amountSum.Add(feegrantFeesSum...)
		balances := k.bankKeeper.GetAllBalances(ctx, ownerAddress)
		if feesAndClaimRecordsAmountSum.IsAnyGT(balances) {
			return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "owner balance is too small (%s < %s)", balances, feesAndClaimRecordsAmountSum)
		}
		if err = k.addClaimRecordsToDefaultCampaign(ctx, ownerAddress, amountSum); err != nil {
			return err
		}
	}

	campaign.CampaignCurrentAmount = campaign.CampaignCurrentAmount.Add(amountSum...)
	campaign.CampaignTotalAmount = campaign.CampaignTotalAmount.Add(amountSum...)
	k.SetCampaign(ctx, *campaign)

	err = k.setupAndSendFeegrant(ctx, ownerAddress, campaign, feegrantFeesSum, claimRecords, feegrantDenom)
	if err != nil {
		return err
	}

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

func (k Keeper) addClaimRecordsToDefaultCampaign(ctx sdk.Context, ownerAddress sdk.AccAddress, amountSum sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddress, types.ModuleName, amountSum)
}

func (k Keeper) DeleteClaimRecord(ctx sdk.Context, owner string, campaignId uint64, userAddress string) error {
	k.Logger(ctx).Debug("delete claim record", "owner", owner, "campaignId", campaignId, "userAddress", userAddress)
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return err
	}

	userEntry, claimRecordAmount, err := k.validateDeleteClaimRecord(ctx, owner, campaign, userAddress)
	if err != nil {
		return err
	}

	if err = k.sendCampaignCurrentAmountToOwner(ctx, &campaign, claimRecordAmount); err != nil {
		return err
	}

	err = k.deleteClaimRecordSendFeegrant(ctx, &campaign, userAddress)
	if err != nil {
		k.Logger(ctx).Debug("delete claim record send feegrant err", "err", err.Error())
	}

	for i, claimRecord := range userEntry.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			userEntry.ClaimRecords = append(userEntry.ClaimRecords[:i], userEntry.ClaimRecords[i+1:]...)
		}
	}

	k.SetUserEntry(ctx, userEntry)
	campaign.CampaignTotalAmount = campaign.CampaignTotalAmount.Sub(claimRecordAmount...)
	k.SetCampaign(ctx, campaign)

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
	if err = types.ValidateAddClaimRecords(claimRecords); err != nil {
		return nil, err
	}
	if err = ValidateCampaignNotEnded(ctx, campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (k Keeper) ValidateCampaignWhenAddedFromVestingPool(ctx sdk.Context, owner string, vestingPoolName string,
	lockupPeriod *time.Duration, vestingPeriod *time.Duration, free sdk.Dec) error {
	_, vestingPool, found := k.vestingKeeper.GetAccountVestingPool(ctx, owner, vestingPoolName)

	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pool %s not found for address %s", vestingPoolName, owner)
	}

	vestingType, err := k.vestingKeeper.GetVestingType(ctx, vestingPool.VestingType)
	if err != nil {
		return err
	}
	if lockupPeriod == nil {
		return errors.Wrap(c4eerrors.ErrParam, "lockup period cannot be nil for vesting pool campaign")
	}
	if vestingPeriod == nil {
		return errors.Wrap(c4eerrors.ErrParam, "lockup period cannot be nil for vesting pool campaign")
	}
	if err := vestingType.ValidateVestingPeriods(*lockupPeriod, *vestingPeriod); err != nil {
		return err
	}
	return vestingType.ValidateVestingFree(free)
}

func (k Keeper) validateClaimRecords(ctx sdk.Context, campaign *types.Campaign, claimRecords []*types.ClaimRecord) (usersEntries []*types.UserEntry, amountSum sdk.Coins, err error) {
	for i, claimRecord := range claimRecords {
		amountSum = amountSum.Add(claimRecord.Amount...)

		userEntry, err := k.addClaimRecordToUserEntry(ctx, campaign.Id, claimRecord.Address, claimRecord.Amount)
		if err != nil {
			return nil, nil, types.WrapClaimRecordIndex(err, i)
		}
		usersEntries = append(usersEntries, userEntry)
	}
	return
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

func (k Keeper) NewModuleAccountSet(ctx sdk.Context, campaignId uint64) sdk.AccAddress {
	moduleAddressName, accountAddr := CreateFeegrantAccountAddress(campaignId)
	account := k.accountKeeper.GetAccount(ctx, accountAddr)
	if account != nil {
		return accountAddr
	}
	macc := &authtypes.ModuleAccount{
		BaseAccount: &authtypes.BaseAccount{
			Address: accountAddr.String(),
		},
		Name: moduleAddressName,
	}
	k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccount(ctx, macc))
	return accountAddr
}

func (k Keeper) validateDeleteClaimRecord(ctx sdk.Context, owner string, campaign types.Campaign, userAddress string) (types.UserEntry, sdk.Coins, error) {
	if campaign.Enabled == true {
		if !campaign.RemovableClaimRecords {
			return types.UserEntry{}, nil, errors.Wrap(sdkerrors.ErrInvalidType, "campaign must have RemovableClaimRecords flag set to true to be able to delete its entries")
		}
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

	amount := userEntry.GetClaimRecord(campaign.Id).Amount

	for _, claimedMissionId := range claimRecord.ClaimedMissions {
		mission, found := k.GetMission(ctx, campaign.Id, claimedMissionId)
		if !found {
			return types.UserEntry{}, nil, errors.Wrapf(sdkerrors.ErrNotFound, "mission with id %d not found", claimedMissionId)
		}
		if mission.MissionType == types.MissionInitialClaim {
			initiialClaimClaimableAmount := k.calculateInitialClaimClaimableAmount(ctx, campaign.Id, &userEntry)
			amount = amount.Sub(initiialClaimClaimableAmount...)
			continue
		}
		claimableFromMission, err := userEntry.ClaimableFromMission(&mission)
		if err != nil {
			return types.UserEntry{}, nil, err
		}
		amount = amount.Sub(claimableFromMission...)
	}

	return userEntry, amount, nil
}

func (k Keeper) ValidateUserEntry(ctx sdk.Context, userAddress string) (types.UserEntry, error) {
	userEntry, found := k.GetUserEntry(
		ctx,
		userAddress,
	)
	if !found {
		return types.UserEntry{}, errors.Wrapf(c4eerrors.ErrNotExists, "userEntry %s doesn't exist", userAddress)
	}

	return userEntry, nil
}

func ValidateClaimRecordExists(userEntry types.UserEntry, campaignId uint64) (claimRecordAmount *types.ClaimRecord, err error) {
	claimRecord := userEntry.GetClaimRecord(campaignId)
	if claimRecord == nil {
		return nil, errors.Wrapf(c4eerrors.ErrParsing, "campaign id %d claim entry doesn't exist", campaignId)
	}

	return claimRecord, nil
}
