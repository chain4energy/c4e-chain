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
	// Set all the campaigns
	for _, campaign := range genState.Campaigns {
		if err := k.ValidateCampaignParams(ctx, campaign.Name, campaign.Description, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount,
			campaign.Free, campaign.StartTime, campaign.EndTime, campaign.CampaignType, campaign.Owner,
			campaign.VestingPoolName, campaign.LockupPeriod, campaign.VestingPeriod); err != nil {
			panic(err)
		}
		k.SetCampaign(ctx, campaign)
	}
	// Set all the missions
	for _, mission := range genState.Missions {
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
	// Set all user entries
	for userEntryIndex, usersEntry := range genState.UsersEntries {
		if err := types.ValidateUserEntry(usersEntry); err != nil {
			panic(errors.Wrapf(err, "userEntry index: %d", userEntryIndex))
		}
		for claimRecordIndex, claimRecord := range usersEntry.ClaimRecords {
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
		k.SetUserEntry(ctx, usersEntry)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.UsersEntries = k.GetAllUsersEntries(ctx)
	genesis.Missions = k.GetAllMission(ctx)
	genesis.Campaigns = k.GetAllCampaigns(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
