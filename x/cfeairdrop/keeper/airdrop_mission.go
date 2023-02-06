package keeper

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func (k Keeper) AddMissionToCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType types.MissionType,
	weight sdk.Dec, claimStartDate *time.Time) error {
	k.Logger(ctx).Debug("add mission to airdrop campaign", "owner", owner, "campaignId", campaignId, "name", name,
		"description", description, "missionType", missionType, "weight", weight)

	if weight.IsNil() {
		k.Logger(ctx).Error("add mission to airdrop mission weight is nil")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "add mission to airdrop campaign weight is nil error")
	}

	if weight.GT(sdk.NewDec(1)) || weight.LT(sdk.ZeroDec()) {
		k.Logger(ctx).Error("add mission to airdrop campaign weight is not between 0 and 1", "weight", weight)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "add mission to airdrop campaign - weight (%s) is not between 0 and 1 error", weight.String())
	}

	if name == "" {
		k.Logger(ctx).Error("add mission to airdrop campaign: empty name")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "add mission to airdrop campaign - empty name error")
	}

	if description == "" {
		k.Logger(ctx).Error("add mission to airdrop campaign: empty description ")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "add mission to airdrop campaign - mission empty description error")
	}

	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Error("add mission to airdrop campaign owner parsing error", "owner", owner, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "add mission to airdrop campaign - owner parsing error: %s", owner).Error())
	}

	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		k.Logger(ctx).Error("add mission to airdrop campaign not found", "campaignId", campaignId)
		return sdkerrors.Wrapf(c4eerrors.ErrNotExists, "add mission to airdrop campaign - campaign with id %d not found error", campaignId)
	}

	if campaign.Owner != owner {
		k.Logger(ctx).Error("add mission to airdrop you are not the owner of this campaign", "campaignId", campaignId)
		return sdkerrors.Wrapf(sdkerrors.ErrorInvalidSigner, "add mission to airdrop campaign - you are not the owner of the campaign with id %d", campaignId)
	}
	if campaign.Enabled == true {
		k.Logger(ctx).Error("add mission to airdrop campaign is enabled", "campaignId", campaignId)
		return sdkerrors.Wrapf(types.ErrCampaignDisabled, "add mission to airdrop - campaign %d is already enabled error", campaignId)
	}
	if campaign.EndTime.Before(ctx.BlockTime()) {
		k.Logger(ctx).Error("add mission to airdrop campaign is disabled", "campaignId", campaignId)
		return sdkerrors.Wrapf(types.ErrCampaignDisabled, "add mission to airdrop - campaign %d is already disabled error", campaignId)
	}
	_, weightSum := k.AllMissionForCampaign(ctx, campaignId)
	weightSum = weightSum.Add(weight)
	if weightSum.GT(sdk.NewDec(1)) {
		k.Logger(ctx).Error("add mission to airdrop all campaign mission weight sum is >= 1", "weightSum", weightSum)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "add mission to airdrop - all campaign missions weight sum is >= 1 (%s > 1) error", weightSum.String())
	}

	mission := types.Mission{
		CampaignId:     campaignId,
		Name:           name,
		Description:    description,
		MissionType:    missionType,
		Weight:         weight,
		ClaimStartDate: claimStartDate,
	}

	k.AppendNewMission(ctx, campaignId, mission)
	return nil
}

func (k Keeper) missionFirstStep(ctx sdk.Context, campaignId uint64, missionId uint64, claimerAddress string) (*types.Campaign, *types.Mission, *types.UserEntry, error) {
	campaign, campaignFound := k.GetCampaign(ctx, campaignId)
	if !campaignFound {
		k.Logger(ctx).Error("mission first step - camapign not found", "campaignId", campaignId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "camapign not found: campaignId %d", campaignId)
	}
	k.Logger(ctx).Debug("campaignId", campaignId, "missionId", missionId, "blockTime", ctx.BlockTime(), "campaigh start", campaign.StartTime, "campaigh end", campaign.EndTime)

	userEntry, found := k.GetUserEntry(ctx, claimerAddress)
	if !found {
		k.Logger(ctx).Debug("mission first step - claim record not found", "claimerAddress", claimerAddress)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "user airdrop entries not found for address %s", claimerAddress)
	}

	if err := campaign.IsActive(ctx.BlockTime()); err != nil {
		k.Logger(ctx).Error("mission first step - camapign disabled", "campaignId", campaignId, "err", err)
		return nil, nil, nil, err
	}

	mission, missionFound := k.GetMission(ctx, campaignId, missionId)
	if !missionFound {
		k.Logger(ctx).Error("mission first step - mission not found", "campaignId", campaignId, "missionId", missionId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "mission not found - campaignId %d, missionId %d", campaignId, missionId)
	}
	k.Logger(ctx).Debug("mission", mission)
	if err := mission.IsEnabled(ctx.BlockTime()); err != nil {
		k.Logger(ctx).Error("claim mission - mission disabled", "campaignId", campaignId, "missionId", missionId, "err", err)
		return nil, nil, nil, sdkerrors.Wrapf(err, "mission disabled - campaignId %d, missionId %d", campaignId, missionId)
	}

	if !userEntry.HasCampaign(campaignId) {
		k.Logger(ctx).Error("mission first step - campaign record not found", "address", claimerAddress, "campaignId", campaignId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "campaign record with id %d not found for address %s", campaignId, claimerAddress)
	}

	return &campaign, &mission, &userEntry, nil
}

func (k Keeper) missionsWeightGreaterThan1(missions []types.Mission, newMissionWeight sdk.Dec) (bool, sdk.Dec) {
	weightSum := newMissionWeight
	for _, mission := range missions {
		weightSum = weightSum.Add(mission.Weight)
	}
	if weightSum.GT(sdk.NewDec(1)) {
		return true, weightSum
	}
	return false, weightSum
}
