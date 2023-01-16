package keeper

import (
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) missionIntialStep(ctx sdk.Context, log string, campaignId uint64, missionId uint64, address string, isHook bool) (*types.Campaign, *types.Mission, *types.UserAirdropEntries, error) {
	campaignConfig, campaignFound := k.GetCampaign(ctx, campaignId)
	if !campaignFound {
		k.Logger(ctx).Error(log+" - camapign not found", "campaignId", campaignId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "camapign not found: campaignId %d", campaignId)
	}
	if err := campaignConfig.IsEnabled(ctx.BlockTime()); err != nil {
		if isHook {
			k.Logger(ctx).Debug(log+" - camapign disabled", "campaignId", campaignId, "err", err)
			return nil, nil, nil, nil
		}
		k.Logger(ctx).Error(log+" - camapign disabled", "campaignId", campaignId, "err", err)
		return nil, nil, nil, sdkerrors.Wrapf(err, "campaign disabled - campaignId %d", campaignId)
	}
	k.Logger(ctx).Debug(log, "campaignId", campaignId, "missionId", missionId, "blockTime", ctx.BlockTime(), "campaigh start", campaignConfig.StartTime, "campaigh end", campaignConfig.EndTime)

	mission, missionFound := k.GetMission(ctx, campaignId, missionId)
	if !missionFound {
		k.Logger(ctx).Error(log+" - mission not found", "campaignId", campaignId, "missionId", missionId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "mission not found - campaignId %d, missionId %d", campaignId, missionId)
	}
	k.Logger(ctx).Debug(log, "mission", mission)

	userAirdropEntries, found := k.GetUserAirdropEntries(ctx, address)
	if !found {
		if isHook {
			k.Logger(ctx).Debug(log+" - claim record not found", "address", address)
			return nil, nil, nil, nil
		}
		k.Logger(ctx).Error(log+" - claim record not found", "address", address)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "claim record not found for address %s", address)
	}

	if !userAirdropEntries.HasCampaign(campaignId) {
		if isHook {
			k.Logger(ctx).Error(log+" - campaign record not found", "address", address, "campaignId", campaignId)
			return nil, nil, nil, nil
		}
		k.Logger(ctx).Error(log+" - campaign record not found", "address", address, "campaignId", campaignId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "campaign record with id: %d not found for address %s", campaignId, address)
	}
	return &campaignConfig, &mission, &userAirdropEntries, nil
}

func (k Keeper) ClaimInitialMission(ctx sdk.Context, campaignId uint64, missionId uint64, address string) error {
	campaignConfig, mission, userAirdropEntries, err := k.missionIntialStep(ctx, "claim initial mission", campaignId, missionId, address, false)
	if err != nil {
		return err
	}
	userAirdropEntries, err = k.completeMission(ctx, true, mission, userAirdropEntries)
	if err != nil {
		return err
	}
	userAirdropEntries, err = k.claimMission(ctx, true, campaignConfig, mission, userAirdropEntries)
	if err != nil {
		return err
	}
	k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	return nil
}

func (k Keeper) ClaimMission(ctx sdk.Context, campaignId uint64, missionId uint64, address string) error {
	campaignConfig, mission, userAirdropEntries, err := k.missionIntialStep(ctx, "claim mission", campaignId, missionId, address, false)
	if err != nil {
		return err
	}
	userAirdropEntries, err = k.claimMission(ctx, false, campaignConfig, mission, userAirdropEntries)
	if err != nil {
		return err
	}

	k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	return nil
}

func (k Keeper) claimMission(ctx sdk.Context, initialClaim bool, campaignConfig *types.Campaign, mission *types.Mission, userAirdropEntries *types.UserAirdropEntries) (*types.UserAirdropEntries, error) {
	campaignId := mission.CampaignId
	missionId := mission.Id
	address := userAirdropEntries.Address
	if !userAirdropEntries.IsMissionCompleted(campaignId, missionId) {
		k.Logger(ctx).Error("claim mission - mission not completed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionNotCompleted, "mission not completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if userAirdropEntries.IsMissionClaimed(campaignId, missionId) {
		k.Logger(ctx).Error("claim mission - mission already claimed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionClaimed, "mission already claimed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if err := userAirdropEntries.ClaimMission(campaignId, missionId); err != nil {
		k.Logger(ctx).Error("claim mission - cannot claime", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionClaiming, err.Error())
	}

	claimableAmount := userAirdropEntries.ClaimableFromMission(mission)
	// TODO initial mission claim should have not waight but get whats left from ther missions

	claimable := sdk.NewCoins(sdk.NewCoin(k.Denom(ctx), claimableAmount))

	// calculate claimable after decay factor
	// decayInfo := k.GetParams(ctx).DecayInformation
	// claimable = decayInfo.ApplyDecayFactor(claimable, ctx.BlockTime())

	// check final claimable non-zero
	// if claimable.Empty() {
	// 	return types.ErrNoClaimable
	// }

	// decrease airdrop supply
	// airdropSupply.Amount = airdropSupply.Amount.Sub(claimable.AmountOf(airdropSupply.Denom))
	// if airdropSupply.Amount.IsNegative() {
	// 	return errors.Critical("airdrop supply is lower than total claimable")
	// }

	// send claimable to the user
	sendTo := address
	if len(userAirdropEntries.ClaimAddress) > 0 {
		sendTo = userAirdropEntries.ClaimAddress
	}
	claimer, err := sdk.AccAddressFromBech32(sendTo)
	if err != nil {
		return nil, sdkerrors.Wrapf(errortypes.ErrParsing, "wrong claiming address %s: "+err.Error(), sendTo)
	}
	start := ctx.BlockTime().Add(campaignConfig.LockupPeriod)
	end := start.Add(campaignConfig.VestingPeriod)
	if err := k.SendToAirdropAccount(ctx, claimer, claimable, start.Unix(), end.Unix(), initialClaim); err != nil {
		return nil, sdkerrors.Wrapf(errortypes.ErrSendCoins, "send to claiming address %s error: "+err.Error(), sendTo)
	}
	return userAirdropEntries, nil

}

func (k Keeper) CompleteMission(ctx sdk.Context, campaignId uint64, missionId uint64, address string, isHook bool) error {
	_, mission, userAirdropEntries, err := k.missionIntialStep(ctx, "complete mission", campaignId, missionId, address, isHook)
	if err != nil {
		return err
	}
	userAirdropEntries, err = k.completeMission(ctx, false, mission, userAirdropEntries)
	if err != nil {
		return err
	}
	k.SetUserAirdropEntries(ctx, *userAirdropEntries)
	return nil
}

// CompleteMission triggers the completion of the mission and distribute the claimable portion of airdrop to the user
// the method fails if the mission has already been completed
func (k Keeper) completeMission(ctx sdk.Context, isInitialClaim bool, mission *types.Mission, userAirdropEntries *types.UserAirdropEntries) (*types.UserAirdropEntries, error) {
	campaignId := mission.CampaignId
	missionId := mission.Id
	address := userAirdropEntries.Address
	// check if the mission is already complted for the claim record
	if userAirdropEntries.IsMissionCompleted(campaignId, missionId) {
		k.Logger(ctx).Error("complete mission - mission already completed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionCompleted, "mission already completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if !isInitialClaim {
		initialClaim, found := k.GetInitialClaim(ctx, campaignId)
		if !found {
			k.Logger(ctx).Error("complete mission - initial claim not found", "campaignId", campaignId)
			return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "initial claim not found - campaignId %d", campaignId)
		}
		if !userAirdropEntries.IsMissionClaimed(initialClaim.CampaignId, initialClaim.MissionId) {
			k.Logger(ctx).Error("complete mission - initial mission not completed", "address", address, "campaignId", campaignId, "missionId", missionId)
			return nil, sdkerrors.Wrapf(types.ErrMissionNotCompleted, "initial mission not completed: address %s, campaignId: %d, missionId: %d", address, initialClaim.CampaignId, initialClaim.MissionId)
		}
	}

	if err := userAirdropEntries.CompleteMission(campaignId, missionId); err != nil {
		k.Logger(ctx).Error("complete mission - cannot complete", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionCompletion, err.Error())
	}
	// k.SetUserAirdropEntries(ctx, userAirdropEntries)

	// err = ctx.EventManager().EmitTypedEvent(&types.EventMissionCompleted{
	// 	MissionID: missionID,
	// 	Claimer:   address,
	// })

	return userAirdropEntries, nil
}

func (k Keeper) AddMissionToAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType types.MissionType,
	weight sdk.Dec) error {
	k.Logger(ctx).Debug("add mission to airdrop campaign", "owner", owner, "campaignId", campaignId, "name", name,
		"description", description, "missionType", missionType, "weight", weight)
	if weight.GT(sdk.NewDec(1)) {
		k.Logger(ctx).Error("add mission to airdrop campaign weight is >= 1", "weight", weight)
		return sdkerrors.Wrapf(errortypes.ErrParam, "add mission to airdrop campaign weight is >= 1 (%s > 1)", weight.String())
	}
	if name == "" {
		k.Logger(ctx).Error("add mission to airdrop campaign: empty name ")
		return sdkerrors.Wrap(errortypes.ErrParam, "add mission to airdrop campaign empty name")
	}
	if description == "" {
		k.Logger(ctx).Error("add mission to airdrop campaign: empty description ")
		return sdkerrors.Wrap(errortypes.ErrParam, "add mission to airdrop campaign empty description")
	}
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("add mission to airdrop campaign owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(errortypes.ErrParsing, sdkerrors.Wrapf(err, "add mission to airdrop campaign - owner parsing error: %s", owner).Error())
	}
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Error("add mission to airdrop campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add mission to airdrop campaign - campaign with id %d not found", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("add mission to airdrop you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "add mission to airdrop campaign - you are not the owner of campaign with id %d", campaignId)
	}
	mission := types.Mission{
		CampaignId:  campaignId,
		Name:        name,
		Description: description,
		MissionType: missionType,
		Weight:      weight,
	}
	k.AppendNewMission(ctx, campaignId, mission)
	return nil
}
