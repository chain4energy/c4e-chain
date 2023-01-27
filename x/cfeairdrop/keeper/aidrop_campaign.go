package keeper

import (
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func (k Keeper) CreateAidropCampaign(ctx sdk.Context, owner string, name string, description string, feegrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	k.Logger(ctx).Debug("create aidrop campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)
	if name == "" {
		k.Logger(ctx).Debug("create airdrop campaign empty campaign name")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "create airdrop campaign - empty campaign name error")
	}
	if description == "" {
		k.Logger(ctx).Error("create airdrop campaign empty campaign description")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "create airdrop campaign - empty campaign description error")
	}
	if startTime == nil {
		k.Logger(ctx).Error("create airdrop campaign start time is nil")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - start time is nil error")
	}
	if startTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Error("create airdrop campaign start time in the past", "startTime", startTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - start time in the past error (%s < %s)", startTime, ctx.BlockTime())
	}
	if endTime == nil {
		k.Logger(ctx).Error("create airdrop campaign end time is nil")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - end time is nil error")
	}
	if startTime.After(*endTime) {
		k.Logger(ctx).Error("create airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
	}
	if initialClaimFreeAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		initialClaimFreeAmount = &zeroInt
	}
	if initialClaimFreeAmount.IsNegative() {
		k.Logger(ctx).Error("create airdrop campaign initial claim free amount cannot be negative", "initialClaimFreeAmount", initialClaimFreeAmount)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}
	if feegrantAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		feegrantAmount = &zeroInt
	}
	if feegrantAmount.IsNegative() {
		k.Logger(ctx).Error("create airdrop campaign initial feegrant amount cannot be negative", "initialClaimFreeAmount", feegrantAmount)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create airdrop campaign - feegrant amount (%s) cannot be negative", feegrantAmount.String())
	}
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("create vesting account owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - owner parsing error: %s", owner).Error())
	}

	campaign := types.Campaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		FeegrantAmount:         *feegrantAmount,
		InitialClaimFreeAmount: *initialClaimFreeAmount,
		Enabled:                false,
		StartTime:              *startTime,
		EndTime:                *endTime,
		LockupPeriod:           *lockupPeriod,
		VestingPeriod:          *vestingPeriod,
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
		return sdkerrors.Wrapf(c4eerrors.ErrParsing, "edit airdrop campaign -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.Enabled == true {
		k.Logger(ctx).Error("edit airdrop campaign campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrParsing, "edit airdrop campaign -  campaign with id %d doesn't exist", campaignId)
	}
	if campaign.EndTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Error("edit airdrop campaign campaign doesn't exist", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrParsing, "edit airdrop campaign -  campaign with id %d doesn't exist", campaignId)
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
				return sdkerrors.Wrapf(c4eerrors.ErrParam, "edit airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
			}
			campaign.EndTime = *endTime
		} else {
			if startTime.After(campaign.EndTime) {
				k.Logger(ctx).Error("edit airdrop campaign start time is after end time", "startTime", startTime, "endTime", endTime)
				return sdkerrors.Wrapf(c4eerrors.ErrParam, "cedit airdrop campaign - start time is after end time error (%s > %s)", startTime, endTime)
			}
		}
		campaign.StartTime = *startTime
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
		k.Logger(ctx).Error("close airdrop campaign campaign campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "close airdrop campaign - campaign with id %d not found error", campaignId)
	}
	if campaign.EndTime.After(ctx.BlockTime()) {
		k.Logger(ctx).Debug("close airdrop campaign campaign is not over yet", "startTime", campaign.StartTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "close airdrop campaign - campaign with id %d campaign is not over yet (endtime - %s < %s)", campaignId, campaign.EndTime, ctx.BlockTime())
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("close airdrop campaign you are not the owner of this campaign")
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "close airdrop campaign - you are not the owner error")
	}
	if campaign.Enabled == false {
		k.Logger(ctx).Error("close airdrop campaign campaign is already closed")
		return sdkerrors.Wrap(types.ErrCampaignDisabled, fmt.Sprintf("close airdrop campaign - campaign with id %d is already closed or have not started yet error", campaignId))
	}
	campaign.Enabled = false
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) StartAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("start airdrop campaign", "owner", owner, "campaignId", campaignId)
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Debug("start airdrop campaign campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "start airdrop campaign campaign with id %d not found", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Debug("start airdrop campaign you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "start airdrop campaign you are not the owner of this campaign")
	}
	if campaign.Enabled == true {
		k.Logger(ctx).Debug("start airdrop campaign campaign has already started", "campaignId", campaignId)
		return sdkerrors.Wrap(c4eerrors.ErrAlreadyExists, fmt.Sprintf("start airdrop campaign campaign with id %d has already started", campaignId))
	}
	if campaign.StartTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Debug("start airdrop campaign campaign start time in the past", "startTime", campaign.StartTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "start airdrop campaign - campaign with id %d start time in the past error (%s < %s)", campaignId, campaign.StartTime, ctx.BlockTime())
	}
	campaign.Enabled = true
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) RemoveAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("start airdrop campaign", "owner", owner, "campaignId", campaignId)
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Error("start airdrop campaign: campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "start airdrop campaign campaign with id %d not found", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("start airdrop campaign you are not the owner")
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "start airdrop campaign you are not the owner")
	}
	if campaign.Enabled == true {
		k.Logger(ctx).Error("start airdrop campaign campaign has already started")
		return sdkerrors.Wrap(c4eerrors.ErrAlreadyExists, "start airdrop campaign campaign has already started")
	}

	k.RemoveCampaign(ctx, campaignId)
	k.RemoveAllMissionForCampaign(ctx, campaignId)
	return nil
}
