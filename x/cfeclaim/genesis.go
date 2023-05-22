package cfeclaim

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.Logger(ctx).Debug("init cfeclaim genesis", "genState", genState)
	// Set all the campaigns
	for _, campaign := range genState.Campaigns {
		//if err := types.ValidateCreateCampaignParams(name, description, startTime, endTime, campaignType, vestingPoolName); err != nil {
		//	return err
		//}
		//
		//if campaignType == types.VestingPoolCampaign {
		//	return k.ValidateCampaignWhenAddedFromVestingPool(ctx, owner, vestingPoolName, lockupPeriod, vestingPeriod, free)
		//}
		k.SetCampaign(ctx, campaign)
	}
	// Set all the missions
	for _, mission := range genState.Missions {
		campaign, found := k.GetCampaign(ctx, mission.CampaignId)
		if !found {
			panic(fmt.Sprintf("Campaign %d not found for mission %s", mission.CampaignId, mission.Name))
		}
		if err := k.ValidateAddMissionToCampaign(ctx, campaign.Owner, mission.CampaignId, mission.Name,
			mission.Description, mission.MissionType, mission.Weight, mission.ClaimStartDate); err != nil {
			panic(err)
		}
		k.SetMission(ctx, mission)
	}
	// Set all the claimRecords
	for _, usersEntry := range genState.UsersEntries {
		if err := types.ValidateUserEntry(usersEntry); err != nil {
			panic(err)
		}
		k.SetUserEntry(ctx, usersEntry)
	}
}

// TODO : add valitations
// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.UsersEntries = k.GetAllUsersEntries(ctx)
	genesis.Missions = k.GetAllMission(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
