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
	// Set all the mission
	for _, elem := range genState.Missions {
		k.SetMission(ctx, elem)
	}
	// Set all the claimRecordXX
	for _, elem := range genState.UsersEntries {
		k.SetUserEntry(ctx, elem)
	}
	for _, elem := range genState.AirdropClaimsLeft {
		k.IncrementAirdropClaimsLeft(ctx, elem)
	}
	for _, elem := range genState.AirdropDistrubitions {
		k.IncrementAirdropDistrubitions(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.UsersEntries = k.GetUsersEntries(ctx)
	genesis.Missions = k.GetAllMission(ctx)
	genesis.AirdropDistrubitions = k.GetAllAirdropDistrubitions(ctx)
	genesis.AirdropClaimsLeft = k.GetAllAirdropClaimsLeft(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
