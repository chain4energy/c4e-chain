package cfeclaim

import (
	"cosmossdk.io/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	setCampaigns(ctx, k, genState.Campaigns, genState.CampaignCount)
	setMissions(ctx, k, genState.Missions, genState.MissionCounts)
	setUsersEntries(ctx, k, genState.UsersEntries)
}

func setCampaigns(ctx sdk.Context, k keeper.Keeper, campaigns []types.Campaign, campaignCount uint64) {
	for _, campaign := range campaigns {
		if err := k.ValidateCampaignParams(ctx, campaign.Name, campaign.Description, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount,
			campaign.Free, campaign.StartTime, campaign.EndTime, campaign.CampaignType, campaign.Owner,
			campaign.VestingPoolName, campaign.LockupPeriod, campaign.VestingPeriod); err != nil {
			panic(err)
		}
		k.SetCampaign(ctx, campaign)
	}
	k.SetCampaignCount(ctx, campaignCount)
}

func setMissions(ctx sdk.Context, k keeper.Keeper, missions []types.Mission, missionCounts []*types.MissionCount) {
	for _, mission := range missions {
		campaign, err := k.MustGetCampaign(ctx, mission.CampaignId)
		if err != nil {
			panic(errors.Wrapf(err, "mission %s", mission.Name))
		}
		if _, err = k.ValidateAddMission(ctx, campaign.Owner, mission.CampaignId, mission.Name,
			mission.Description, mission.MissionType, mission.Weight, mission.ClaimStartDate); err != nil {
			panic(err)
		}
		k.SetMission(ctx, mission)
	}
	for _, misisonCount := range missionCounts {
		k.SetMissionCount(ctx, misisonCount.CampaignId, misisonCount.Count)
	}
}

func setUsersEntries(ctx sdk.Context, k keeper.Keeper, usersEntries []types.UserEntry) {
	for userEntryIndex, usersEntry := range usersEntries {
		if err := usersEntry.Validate(); err != nil {
			panic(errors.Wrapf(err, "userEntry index: %d", userEntryIndex))
		}
		validateClaimRecords(ctx, k, usersEntry.ClaimRecords, int64(userEntryIndex))
		k.SetUserEntry(ctx, usersEntry)
	}
}

func validateClaimRecords(ctx sdk.Context, k keeper.Keeper, claimRecords []*types.ClaimRecord, userEntryIndex int64) {
	for claimRecordIndex, claimRecord := range claimRecords {
		_, err := k.MustGetCampaign(ctx, claimRecord.CampaignId)
		if err != nil {
			panic(errors.Wrapf(err, "userEntry index: %d, claimRecord index: %d", userEntryIndex, claimRecordIndex))
		}
		for i, missionId := range claimRecord.ClaimedMissions {
			_, err = k.MustGetMission(ctx, claimRecord.CampaignId, missionId)
			if err != nil {
				panic(errors.Wrapf(err, "userEntry index: %d, claimRecord index: %d, claimed mission index: %d", userEntryIndex, claimRecordIndex, i))
			}
		}
		for i, missionId := range claimRecord.CompletedMissions {
			_, err = k.MustGetMission(ctx, claimRecord.CampaignId, missionId)
			if err != nil {
				panic(errors.Wrapf(err, "userEntry index: %d, claimRecord index: %d, completed mission index: %d", userEntryIndex, claimRecordIndex, i))
			}
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.UsersEntries = k.GetAllUsersEntries(ctx)
	genesis.Missions = k.GetAllMission(ctx)
	genesis.Campaigns = k.GetAllCampaigns(ctx)
	genesis.CampaignCount = k.GetCampaignCount(ctx)
	genesis.MissionCounts = k.GetAllMissionCount(ctx)

	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
