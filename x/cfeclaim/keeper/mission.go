package keeper

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
	"time"
)

func (k Keeper) AddMissionToCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType types.MissionType,
	weight sdk.Dec, claimStartDate *time.Time) error {
	k.Logger(ctx).Debug("add mission to claim campaign", "owner", owner, "campaignId", campaignId, "name", name,
		"description", description, "missionType", missionType, "weight", weight)

	err := types.ValidateAddMissionToCampaign(owner, name, description, missionType, &weight)
	if err != nil {
		return err
	}

	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return err
	}
	if err = types.ValidateCampaignIsNotEnabled(campaign); err != nil {
		return err
	}
	if err = ValidateOwner(campaign, owner); err != nil {
		return err
	}
	if err = k.ValidateMissionWeightsNotGreaterThan1(ctx, campaignId, weight); err != nil {
		return err
	}

	if err = ValidateCampaignNotEnded(ctx, campaign); err != nil {
		return err
	}
	if err = ValidateMissionClaimStartDate(campaign, claimStartDate); err != nil {
		return err
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

	eventClaimStartDate := ""
	if claimStartDate != nil {
		eventClaimStartDate = claimStartDate.String()
	}

	event := &types.AddMissionToCampaign{
		Owner:          owner,
		CampaignId:     strconv.FormatUint(campaignId, 10),
		Name:           name,
		Description:    description,
		MissionType:    missionType.String(),
		Weight:         weight.String(),
		ClaimStartDate: eventClaimStartDate,
	}

	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("add mission to campaign emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) missionFirstStep(ctx sdk.Context, campaignId uint64, missionId uint64, claimerAddress string) (*types.Campaign, *types.Mission, *types.UserEntry, error) {
	campaign, campaignFound := k.GetCampaign(ctx, campaignId)
	if !campaignFound {
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "camapign not found: campaignId %d", campaignId)
	}
	k.Logger(ctx).Debug("campaignId", campaignId, "missionId", missionId, "blockTime", ctx.BlockTime(), "campaigh start", campaign.StartTime, "campaigh end", campaign.EndTime)

	userEntry, found := k.GetUserEntry(ctx, claimerAddress)
	if !found {
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "user claim entries not found for address %s", claimerAddress)
	}

	if err := campaign.IsActive(ctx.BlockTime()); err != nil {
		return nil, nil, nil, err
	}

	mission, missionFound := k.GetMission(ctx, campaignId, missionId)
	if !missionFound {
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "mission not found - campaignId %d, missionId %d", campaignId, missionId)
	}
	k.Logger(ctx).Debug("mission", mission)
	if err := mission.IsEnabled(ctx.BlockTime()); err != nil {
		return nil, nil, nil, sdkerrors.Wrapf(err, "mission disabled - campaignId %d, missionId %d", campaignId, missionId)
	}

	if !userEntry.HasCampaign(campaignId) {
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "campaign record with id %d not found for address %s", campaignId, claimerAddress)
	}

	return &campaign, &mission, &userEntry, nil
}

func (k Keeper) ValidateMissionWeightsNotGreaterThan1(ctx sdk.Context, campaignId uint64, newMissionWeight sdk.Dec) error {
	_, weightSum := k.AllMissionForCampaign(ctx, campaignId)
	weightSum = weightSum.Add(newMissionWeight)
	if weightSum.GT(sdk.NewDec(1)) {
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "add mission to claim - all campaign missions weight sum is >= 1 (%s > 1) error", weightSum.String())
	}
	return nil
}

func ValidateMissionClaimStartDate(campaign types.Campaign, claimStartDate *time.Time) error {
	if claimStartDate == nil {
		return nil
	}
	if claimStartDate.After(campaign.EndTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "mission claim start date after campaign end time (end time - %s < %s)", campaign.EndTime, claimStartDate)
	}
	if claimStartDate.Before(campaign.StartTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "mission claim start date before campaign start time (start time - %s > %s)", campaign.StartTime, claimStartDate)
	}
	return nil
}
