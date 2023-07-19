package keeper

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (k Keeper) AddMission(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType types.MissionType,
	weight sdk.Dec, claimStartDate *time.Time) (*types.Mission, error) {
	k.Logger(ctx).Debug("add mission to a campaign", "owner", owner, "campaignId", campaignId, "name", name,
		"description", description, "missionType", missionType, "weight", weight)

	campaign, err := k.ValidateAddMission(ctx, owner, campaignId, name, missionType, weight, claimStartDate)
	if err != nil {
		return nil, err
	}
	if err = campaign.ValidateNotEnabled(); err != nil {
		return nil, err
	}
	if err = campaign.ValidateNotEnded(ctx.BlockTime()); err != nil {
		return nil, err
	}

	mission := types.Mission{
		CampaignId:     campaignId,
		Name:           name,
		Description:    description,
		MissionType:    missionType,
		Weight:         weight,
		ClaimStartDate: claimStartDate,
	}
	mission.Id = k.AppendNewMission(ctx, campaignId, mission)

	event := &types.EventAddMission{
		Id:             mission.Id,
		Owner:          owner,
		CampaignId:     campaignId,
		Name:           name,
		Description:    description,
		MissionType:    missionType,
		Weight:         weight,
		ClaimStartDate: claimStartDate,
	}

	if err = ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("add mission to campaign emit event error", "event", event, "error", err.Error())
	}

	return &mission, nil
}

func (k Keeper) ValidateAddMission(ctx sdk.Context, owner string, campaignId uint64, name string,
	missionType types.MissionType, weight sdk.Dec, claimStartDate *time.Time) (*types.Campaign, error) {
	err := types.ValidateAddMission(owner, name, missionType, weight)
	if err != nil {
		return nil, err
	}
	campaign, err := k.MustGetCampaign(ctx, campaignId)
	if err != nil {
		return nil, err
	}
	if err = campaign.ValidateOwner(owner); err != nil {
		return nil, err
	}
	if err = k.ValidateMissionsWeightAndType(ctx, campaignId, weight, missionType); err != nil {
		return nil, err
	}
	return campaign, nil
}

func (k Keeper) prepareClaimData(ctx sdk.Context, campaignId uint64, missionId uint64, claimerAddress string) (*types.Campaign, *types.Mission, *types.UserEntry, *types.ClaimRecord, error) {
	campaign, err := k.MustGetCampaign(ctx, campaignId)
	if err != nil {
		return missionFirstStepReturnError(err)
	}
	k.Logger(ctx).Debug("prepare claim data", "campaignId", campaignId, "missionId", missionId, "blockTime", ctx.BlockTime(), "campaign start", campaign.StartTime, "campaign end", campaign.EndTime)

	if err = campaign.ValidateIsActive(ctx.BlockTime()); err != nil {
		return missionFirstStepReturnError(err)
	}
	userEntry, err := k.MustGetUserEntry(ctx, claimerAddress)
	if err != nil {
		return missionFirstStepReturnError(err)
	}
	mission, err := k.MustGetMission(ctx, campaignId, missionId)
	if err != nil {
		return missionFirstStepReturnError(err)
	}
	k.Logger(ctx).Debug("prepare claim data", "mission", mission)
	if mission.MissionType == types.MissionToDefine {
		return missionFirstStepReturnError(errors.Wrapf(types.ErrMissionClaiming, "cannot claim mission with type TO_DEFINE"))
	}
	if err = mission.IsEnabled(ctx.BlockTime()); err != nil {
		return missionFirstStepReturnError(err)
	}
	claimRecord, err := userEntry.MustGetClaimRecord(campaignId)
	if err != nil {
		return missionFirstStepReturnError(err)
	}

	return campaign, mission, &userEntry, claimRecord, nil
}

func missionFirstStepReturnError(err error) (*types.Campaign, *types.Mission, *types.UserEntry, *types.ClaimRecord, error) {
	return nil, nil, nil, nil, err
}

func (k Keeper) ValidateMissionsWeightAndType(ctx sdk.Context, campaignId uint64, newMissionWeight sdk.Dec, missionType types.MissionType) error {
	missions, weightSum := k.AllMissionForCampaign(ctx, campaignId)
	weightSum = weightSum.Add(newMissionWeight)
	if weightSum.GT(sdk.NewDec(1)) {
		return errors.Wrapf(c4eerrors.ErrParam, "all campaign missions weight sum is >= 1 (%s > 1) error", weightSum.String())
	}

	if len(missions) > 0 && missionType == types.MissionInitialClaim {
		return errors.Wrapf(c4eerrors.ErrParam, "there can be only one mission with InitialClaim type and must be first in the campaign")
	}

	return nil
}
