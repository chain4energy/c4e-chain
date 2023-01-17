package cfeairdrop

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the campaignXX
	for _, elem := range genState.Campaigns {
		k.SetCampaign(ctx, elem)
	}
	// Set all the claimRecordXX
	for _, elem := range genState.UserAirdropEntries {
		k.SetUserAirdropEntries(ctx, elem)
	}
	// Set all the mission
	for _, elem := range genState.Missions {
		k.SetMission(ctx, elem)
	}

	// Set airdropEntry count
	k.SetAirdropEntryCount(ctx, genState.AirdropEntryCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.UserAirdropEntries = k.GetUsersAirdropEntries(ctx)
	genesis.Missions = k.GetAllMission(ctx)
	genesis.AirdropEntryList = k.GetAllAirdropEntry(ctx)
	genesis.AirdropEntryCount = k.GetAirdropEntryCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
