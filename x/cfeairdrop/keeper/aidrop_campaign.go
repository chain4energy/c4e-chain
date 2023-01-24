package keeper

import (
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func (k Keeper) CreateAidropCampaign(ctx sdk.Context, owner string, name string, description string, feegrantAmount sdk.Int, initialClaimFreeAmount sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) error {
	k.Logger(ctx).Debug("create aidrop campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)
	if name == "" {
		k.Logger(ctx).Error("create airdrop campaign campaign: empty name ")
		return sdkerrors.Wrap(errortypes.ErrParam, "add mission to airdrop campaign empty name")
	}
	if description == "" {
		k.Logger(ctx).Error("create airdrop campaign campaign: empty description ")
		return sdkerrors.Wrap(errortypes.ErrParam, "add mission to airdrop campaign empty description")
	}
	if startTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Error("create airdrop campaign start time in the past", "startTime", startTime)
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - start time in the past error  (%s < %s)", startTime, ctx.BlockTime())
	}
	if startTime.After(endTime) {
		k.Logger(ctx).Error("create airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
	}
	if initialClaimFreeAmount.IsNegative() {
		k.Logger(ctx).Error("create airdrop campaign initial claim free amount cannot be negative", "initialClaimFreeAmount", initialClaimFreeAmount)
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}
	if feegrantAmount.IsNegative() {
		k.Logger(ctx).Error("create airdrop campaign initial feegrant amount cannot be negative", "initialClaimFreeAmount", feegrantAmount)
		return sdkerrors.Wrapf(errortypes.ErrParam, "create airdrop campaign - feegrant amount (%s) cannot be negative", feegrantAmount.String())
	}
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("create vesting account owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(errortypes.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - owner parsing error: %s", owner).Error())
	}

	campaign := types.Campaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		FeegrantAmount:         feegrantAmount,
		InitialClaimFreeAmount: initialClaimFreeAmount,
		Enabled:                false,
		StartTime:              &startTime,
		EndTime:                &endTime,
		LockupPeriod:           lockupPeriod,
		VestingPeriod:          vestingPeriod,
	}

	campaignId := k.AppendNewCampaign(ctx, campaign)
	missionInitial := types.NewInitialMission(campaignId)
	k.AppendNewMission(ctx, campaignId, *missionInitial)
	return nil
}

func (k Keeper) EditAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	k.Logger(ctx).Debug("edit airdrop campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)
	campaign, found := k.GetCampaign(
		ctx,
		campaignId,
	)
	if !found {
		k.Logger(ctx).Error("edit airdrop campaign campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrParsing, "edit airdrop campaign -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.Enabled == true {
		k.Logger(ctx).Error("edit airdrop campaign campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrParsing, "edit airdrop campaign -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("edit airdrop campaign you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "edit airdrop campaign - you are not the owner of campaign with id %d", campaignId)
	}
	if name != "" {
		campaign.Name = name
	}
	if description != "" {
		campaign.Description = description
	}
	if startTime != nil {
		if endTime != nil {
			if startTime.After(*endTime) {
				k.Logger(ctx).Error("edit airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
				return sdkerrors.Wrapf(errortypes.ErrParam, "edit airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
			}
			campaign.EndTime = endTime
		} else {
			if startTime.After(*campaign.EndTime) {
				k.Logger(ctx).Error("edit airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
				return sdkerrors.Wrapf(errortypes.ErrParam, "cedit airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
			}
		}
		campaign.StartTime = startTime
	}
	if vestingPeriod != nil {
		campaign.VestingPeriod = *vestingPeriod
	}
	if lockupPeriod != nil {
		campaign.LockupPeriod = *lockupPeriod
	}
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) CloseAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64, airdropCloseAction types.AirdropCloseAction) error {
	k.Logger(ctx).Debug("close airdrop campaign", "owner", owner, "campaignId", campaignId, "airdropCloseAction", airdropCloseAction)
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Error("close airdrop campaign campaign: campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrNotExists, "close airdrop campaign campaign with id %d not found", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("close airdrop campaign you are not the owner")
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "close airdrop campaign you are not the owner")
	}
	if campaign.Enabled == false {
		k.Logger(ctx).Error("close airdrop campaign campaign is already closed")
		return sdkerrors.Wrap(types.ErrCampaignDisabled, "close airdrop campaign campaign is already closed")
	}
	campaign.Enabled = false
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) StartAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("start airdrop campaign", "owner", owner, "campaignId", campaignId)
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Error("start airdrop campaign: campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrNotExists, "start airdrop campaign campaign with id %d not found", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("start airdrop campaign you are not the owner")
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "start airdrop campaign you are not the owner")
	}
	if campaign.Enabled == true {
		k.Logger(ctx).Error("start airdrop campaign campaign has already started")
		return sdkerrors.Wrap(errortypes.ErrAlreadyExists, "start airdrop campaign campaign has already started")
	}
	if campaign.StartTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Error("start airdrop campaign start time in the past", "startTime", campaign.StartTime)
		return sdkerrors.Wrapf(errortypes.ErrParam, "start airdrop campaign - start time in the past error  (%s < %s)", campaign.StartTime, ctx.BlockTime())
	}
	campaign.Enabled = true
	k.SetCampaign(ctx, campaign)
	return nil
}
