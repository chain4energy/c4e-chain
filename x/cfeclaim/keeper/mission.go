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

	campaign, err := k.ValidateAddMissionToCampaign(ctx, owner, campaignId, name, description, missionType, weight, claimStartDate)
	if err != nil {
		return err
	}
	if err = types.ValidateCampaignIsNotEnabled(*campaign); err != nil {
		return err
	}
	if err = ValidateCampaignNotEnded(ctx, *campaign); err != nil {
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

	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("add mission to campaign emit event error", "event", event, "error", err.Error())
	}

	return nil
}

func (k Keeper) ValidateAddMissionToCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, // TODO to moze ez do types z przkazanie paramrwo z KV store - dp zastanowienia i nazwa VlidaetNewMission
	missionType types.MissionType, weight sdk.Dec, claimStartDate *time.Time) (*types.Campaign, error) {
	err := types.ValidateAddMissionToCampaign(owner, name, description, missionType, &weight)
	if err != nil {
		return nil, err
	}
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return nil, err
	}
	if err = ValidateOwner(campaign, owner); err != nil {
		return nil, err
	}
	if err = k.ValidateMissionWeightsNotGreaterThan1(ctx, campaignId, weight); err != nil { // TODO do types mission.go z przyjmowanie listy misji 
		return nil, err
	}
	if err = k.ValidateOnlyFirstMissionInitialClaim(ctx, campaignId, missionType); err != nil { // TODO do types mission.go z przyjmowanie listy misji - odciazy to tez z podwnego wyciagane lity misji
		return nil, err
	}
	if err = ValidateMissionClaimStartDate(campaign, claimStartDate); err != nil { 
		return nil, err
	}
	return &campaign, nil
}

func (k Keeper) missionFirstStep(ctx sdk.Context, campaignId uint64, missionId uint64, claimerAddress string) (*types.Campaign, *types.Mission, *types.UserEntry, error) { // TODO nazwa niewiele mowi
	campaign, campaignFound := k.GetCampaign(ctx, campaignId) // TODO mam tez wracenie ze wycigamy to wiele raz w jednym przebiegu - metoda raczej powinno przyjnmowac obiekt a metoda mozna dodac MustGetCampaign na store i wolona w mtodzi nadrzednej
	if !campaignFound {
		return nil, nil, nil, errors.Wrapf(sdkerrors.ErrNotFound, "camapign not found: campaignId %d", campaignId)
	}
	k.Logger(ctx).Debug("campaignId", campaignId, "missionId", missionId, "blockTime", ctx.BlockTime(), "campaigh start", campaign.StartTime, "campaigh end", campaign.EndTime)

	userEntry, found := k.GetUserEntry(ctx, claimerAddress)
	if !found {
		return nil, nil, nil, errors.Wrapf(sdkerrors.ErrNotFound, "user claim entries not found for address %s", claimerAddress)
	}

	if err := campaign.IsActive(ctx.BlockTime()); err != nil {
		return nil, nil, nil, err
	}

	mission, err := k.ValidateMissionExists(ctx, campaignId, missionId)
	if err != nil {
		return nil, nil, nil, err
	}

	k.Logger(ctx).Debug("mission", mission)
	if err := mission.IsEnabled(ctx.BlockTime()); err != nil {
		return nil, nil, nil, errors.Wrapf(err, "mission disabled - campaignId %d, missionId %d", campaignId, missionId)
	}

	if !userEntry.HasCampaign(campaignId) {
		return nil, nil, nil, errors.Wrapf(sdkerrors.ErrNotFound, "campaign record with id %d not found for address %s", campaignId, claimerAddress)
	}

	return &campaign, mission, &userEntry, nil
} // TODO - nowa linia
func (k Keeper) ValidateMissionExists(ctx sdk.Context, campaignId uint64, missionId uint64) (*types.Mission, error) {
	mission, missionFound := k.GetMission(ctx, campaignId, missionId) // TODO nazwa MustGetMission i do mission_store.go
	if !missionFound {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "mission not found - campaignId %d, missionId %d", campaignId, missionId)
	}
	return &mission, nil
}

func (k Keeper) ValidateMissionWeightsNotGreaterThan1(ctx sdk.Context, campaignId uint64, newMissionWeight sdk.Dec) error {  // TODO do types mission.go z przyjmowanie listy misji 
	_, weightSum := k.AllMissionForCampaign(ctx, campaignId)
	weightSum = weightSum.Add(newMissionWeight)
	if weightSum.GT(sdk.NewDec(1)) {
		return errors.Wrapf(c4eerrors.ErrParam, "all campaign missions weight sum is >= 1 (%s > 1) error", weightSum.String())
	}
	return nil
}

func (k Keeper) ValidateOnlyFirstMissionInitialClaim(ctx sdk.Context, campaignId uint64, missionType types.MissionType) error {  // TODO do types mission.go z przyjmowanie listy misji 
	missions, _ := k.AllMissionForCampaign(ctx, campaignId)
	if len(missions) > 0 && missionType == types.MissionInitialClaim {
		return errors.Wrapf(c4eerrors.ErrParam, "there can be only one mission with InitialClaim type and must be first in the campaign")
	} else if len(missions) == 0 && missionType != types.MissionInitialClaim {
		return errors.Wrapf(c4eerrors.ErrParam, "there can be only one mission with InitialClaim type and must be first in the campaign")
	}
	return nil
}

func ValidateMissionClaimStartDate(campaign types.Campaign, claimStartDate *time.Time) error { // TODO metoda Campaign i pod nazwa MustBetweenStartAndEnd - mozna sobie ustalic zasade ze Must jak nie gada zwraca blad jak bymetod czy funkcja IsBetweenStartAndEnd to wtedy wynik (tym przypdku bool)
	if claimStartDate == nil {
		return nil
	}
	if claimStartDate.After(campaign.EndTime) { // TODO or equal zeby bylo spojnie ze end time to juz moment w ktorym nie dziala
		return errors.Wrapf(c4eerrors.ErrParam, "mission claim start date after campaign end time (end time - %s < %s)", campaign.EndTime, claimStartDate)
	}
	if claimStartDate.Before(campaign.StartTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "mission claim start date before campaign start time (start time - %s > %s)", campaign.StartTime, claimStartDate)
	}
	return nil
}
