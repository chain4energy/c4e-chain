package keeper

import (
	errortypes "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func (k Keeper) CreateAidropCampaign(ctx sdk.Context, owner string, name string, description string, denom string, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) error {
	k.Logger(ctx).Debug("create aidrop campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)
	if name == "" {
		k.Logger(ctx).Error("create airdrop campaign campaign: empty name ")
		return sdkerrors.Wrap(errortypes.ErrParam, "add mission to airdrop campaign empty name")
	}
	if denom == "" {
		k.Logger(ctx).Error("create airdrop campaign campaign: empty denom ")
		return sdkerrors.Wrap(errortypes.ErrParam, "add mission to airdrop campaign empty denom")
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
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("create vesting account owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(errortypes.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - owner parsing error: %s", owner).Error())
	}

	campaign := types.NewAirdropCampaign(owner, name, description, denom, startTime, endTime, lockupPeriod, vestingPeriod)
	k.AppendNewCampaign(ctx, *campaign)
	return nil
}

func (k Keeper) CloseAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64, burn bool, communityPoolSend bool) error {
	k.Logger(ctx).Debug("close airdrop campaign", "owner", owner, "campaignId", campaignId, "burn", burn,
		"communityPoolSend", communityPoolSend)
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Error("close airdrop campaign campaign: campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(errortypes.ErrNotExists, "close airdrop campaign campaign with id %d not found", campaignId)
	}
	if campaign.Owner != owner {
		k.Logger(ctx).Error("close airdrop campaign you are not the owner")
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "close airdrop campaign you are not the owner")
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
	campaign.Enabled = true
	k.SetCampaign(ctx, campaign)
	return nil
}
