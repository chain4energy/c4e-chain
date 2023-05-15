package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"strconv"
)

func (k Keeper) InitialClaim(ctx sdk.Context, claimer string, campaignId uint64, additionalAddress string) error {
	k.Logger(ctx).Debug("initial claim", "claimer", claimer, "campaignId", campaignId, "additionalAddress", additionalAddress)

	addressToClaim := claimer
	if additionalAddress != "" {
		if err := k.validateAdditionalAddressToClaim(ctx, additionalAddress); err != nil {
			return err
		}
		addressToClaim = additionalAddress
	}

	campaign, mission, userEntry, err := k.missionFirstStep(ctx, campaignId, types.InitialMissionId, addressToClaim)
	if err != nil {
		return err
	}
	userEntry.ClaimAddress = addressToClaim

	userEntry, err = k.completeMission(mission, userEntry)
	if err != nil {
		return err
	}

	claimableAmount := k.calculateInitialClaimClaimableAmount(ctx, campaignId, userEntry)

	updatedFree, err := k.calculateInitialClaimFree(claimableAmount, campaign)
	if err != nil {
		return err
	}

	userEntry, err = k.claimMission(ctx, campaign, mission, userEntry, claimableAmount, updatedFree)
	if err != nil {
		return err
	}

	if campaign.FeegrantAmount.GT(math.ZeroInt()) {
		granteeAddr, err := sdk.AccAddressFromBech32(userEntry.Address)
		if err != nil {
			return err
		}
		_, accountAddr := CreateFeegrantAccountAddress(campaignId)
		if err = k.revokeFeeAllowance(ctx, accountAddr, granteeAddr); err != nil {
			return err
		}
	}

	k.SetUserEntry(ctx, *userEntry)

	event := &types.InitialClaim{
		Claimer:        claimer,
		CampaignId:     strconv.FormatUint(campaignId, 10),
		AddressToClaim: addressToClaim,
		Amount:         claimableAmount.String(),
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Debug("initial claim emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) Claim(ctx sdk.Context, campaignId uint64, missionId uint64, claimer string) error {
	k.Logger(ctx).Debug("claim", "claimer", claimer, "campaignId", campaignId, "missionId", missionId)

	campaign, mission, userEntry, err := k.missionFirstStep(ctx, campaignId, missionId, claimer)
	if err != nil {
		return err
	}

	if !userEntry.IsInitialMissionClaimed(campaignId) {
		return errors.Wrapf(types.ErrMissionNotCompleted, "initial mission not completed: address %s, campaignId: %d", userEntry.Address, campaignId)
	}

	if mission.MissionType == types.MissionClaim {
		userEntry, err = k.completeMission(mission, userEntry)
		if err != nil {
			return err
		}
	}

	claimableAmount, err := userEntry.ClaimableFromMission(mission)
	if err != nil {
		return err
	}
	userEntry, err = k.claimMission(ctx, campaign, mission, userEntry, claimableAmount, nil)
	if err != nil {
		return err
	}

	k.SetUserEntry(ctx, *userEntry)

	event := &types.Claim{
		Claimer:    claimer,
		CampaignId: strconv.FormatUint(campaignId, 10),
		MissionId:  strconv.FormatUint(missionId, 10),
		Amount:     claimableAmount.String(),
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Debug("claim emit event error", "event", event, "error", err.Error())
	}
	return nil
}

func (k Keeper) CompleteMissionFromHook(ctx sdk.Context, campaignId uint64, missionId uint64, address string) error {
	_, mission, userEntry, err := k.missionFirstStep(ctx, campaignId, missionId, address)
	if err != nil {
		k.Logger(ctx).Debug("complete mission from hook", "err", err.Error())
		return err
	}
	if !userEntry.IsInitialMissionClaimed(campaignId) {
		k.Logger(ctx).Debug("complete mission - initial mission not completed", "claimerAddress", address, "campaignId", campaignId, "missionId", missionId)
		return errors.Wrapf(types.ErrMissionNotCompleted, "initial mission not completed: address %s, campaignId: %d, missionId: %d", address, campaignId, 0)
	}
	userEntry, err = k.completeMission(mission, userEntry)
	if err != nil {
		return err
	}

	k.SetUserEntry(ctx, *userEntry)

	event := &types.CompleteMissionFromHook{
		CampaignId:  strconv.FormatUint(campaignId, 10),
		MissionId:   strconv.FormatUint(missionId, 10),
		UserAddress: address,
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Debug("complete mission from hook event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) completeMission(mission *types.Mission, userEntry *types.UserEntry) (*types.UserEntry, error) {
	campaignId := mission.CampaignId
	missionId := mission.Id
	address := userEntry.Address

	if userEntry.IsMissionCompleted(campaignId, missionId) {
		return nil, errors.Wrapf(types.ErrMissionCompleted, "address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if err := userEntry.CompleteMission(campaignId, missionId); err != nil {
		return nil, errors.Wrapf(types.ErrMissionCompletion, err.Error())
	}

	return userEntry, nil
}

func (k Keeper) claimMission(ctx sdk.Context, campaign *types.Campaign, mission *types.Mission, userEntry *types.UserEntry,
	claimableAmount sdk.Coins, updatedFee *sdk.Dec) (*types.UserEntry, error) {
	campaignId := mission.CampaignId
	missionId := mission.Id
	address := userEntry.ClaimAddress

	if !userEntry.IsMissionCompleted(campaignId, missionId) {
		return nil, errors.Wrapf(types.ErrMissionNotCompleted, "address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if userEntry.IsMissionClaimed(campaignId, missionId) {
		return nil, errors.Wrapf(types.ErrMissionClaimed, "mission already claimed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if err := userEntry.ClaimMission(campaignId, missionId); err != nil {
		return nil, errors.Wrapf(types.ErrMissionClaiming, err.Error())
	}

	free := campaign.Free
	if updatedFee != nil {
		free = *updatedFee
	}

	if campaign.CampaignType == types.VestingPoolCampaign {
		if err := k.vestingKeeper.SendReservedToNewVestingAccount(ctx, campaign.Owner, userEntry.ClaimAddress, campaign.VestingPoolName,
			claimableAmount.AmountOf(k.vestingKeeper.Denom(ctx)), campaign.Id, free, campaign.LockupPeriod, campaign.VestingPeriod); err != nil {
			return nil, err
		}
	} else {
		start := ctx.BlockTime().Add(campaign.LockupPeriod)
		end := start.Add(campaign.VestingPeriod)
		if _, err := k.vestingKeeper.SendToPeriodicContinuousVestingAccountFromModule(ctx, types.ModuleName, userEntry.ClaimAddress,
			claimableAmount, free, start.Unix(), end.Unix()); err != nil {
			return nil, errors.Wrapf(c4eerrors.ErrSendCoins, "send to claiming address %s error: "+err.Error(), userEntry.ClaimAddress)
		}
	}

	k.DecrementCampaignAmountLeft(ctx, campaignId, claimableAmount)
	return userEntry, nil
}

func (k Keeper) validateAdditionalAddressToClaim(ctx sdk.Context, additionalAddress string) error {
	addititonalAccAddress, err := sdk.AccAddressFromBech32(additionalAddress)
	if err != nil {
		return errors.Wrap(c4eerrors.ErrParsing, errors.Wrapf(err, "add mission to claim campaign - additionalAddress parsing error: %s", additionalAddress).Error())
	}

	if k.bankKeeper.BlockedAddr(addititonalAccAddress) {
		return errors.Wrapf(c4eerrors.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", additionalAddress)
	}

	account := k.accountKeeper.GetAccount(ctx, addititonalAccAddress)
	_, baseAccountOk := account.(*authtypes.BaseAccount)
	_, periodicContinuousVestingAccountOk := account.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	if baseAccountOk && periodicContinuousVestingAccountOk {
		return errors.Wrapf(c4eerrors.ErrInvalidAccountType, "account already exists and is not of PeriodicContinuousVestingAccount nor BaseAccount type, got: %T", account)
	}

	return nil
}

func (k Keeper) calculateInitialClaimClaimableAmount(ctx sdk.Context, campaignId uint64, userEntry *types.UserEntry) sdk.Coins {
	allCampaignMissions, _ := k.AllMissionForCampaign(ctx, campaignId)
	claimRecord := userEntry.GetClaimRecord(campaignId)
	allMissionsAmountSum := sdk.NewCoins()
	for _, mission := range allCampaignMissions {
		for _, amount := range claimRecord.Amount {
			allMissionsAmountSum = allMissionsAmountSum.Add(sdk.NewCoin(amount.Denom, mission.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()))
		}
	}
	return claimRecord.Amount.Sub(allMissionsAmountSum...)
}

func (k Keeper) calculateInitialClaimFree(claimableAmount sdk.Coins, campaign *types.Campaign) (*sdk.Dec, error) {
	minFreeAmount := campaign.Free
	for _, claimableAmountCoin := range claimableAmount {
		if claimableAmountCoin.Sub(sdk.NewCoin(claimableAmountCoin.Denom, campaign.InitialClaimFreeAmount)).IsNegative() {
			return nil, errors.Wrapf(c4eerrors.ErrSendCoins, "send to claim account  wrong send coins amount. %s < 1 token (1000000 %s)", claimableAmountCoin.String(), claimableAmountCoin.Denom)
		}
		free := sdk.NewDecFromInt(campaign.InitialClaimFreeAmount).Quo(sdk.NewDecFromInt(claimableAmountCoin.Amount))
		if minFreeAmount.LT(free) {
			minFreeAmount = free
		}
	}

	return &minFreeAmount, nil
}
