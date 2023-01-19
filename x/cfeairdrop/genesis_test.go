package cfeairdrop_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		Campaigns: []types.Campaign{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		Missions: []types.Mission{
			{
				CampaignId: 0,
				Id:         0,
			},
			{
				CampaignId: 1,
				Id:         1,
			},
		},
		UserAirdropEntries: []types.UserAirdropEntries{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CfeairdropKeeper(t)
	cfeairdrop.InitGenesis(ctx, *k, genesisState)
	got := cfeairdrop.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UserAirdropEntries, got.UserAirdropEntries)
	require.ElementsMatch(t, genesisState.Missions, got.Missions)
	require.ElementsMatch(t, genesisState.Campaigns, got.Campaigns)
	require.ElementsMatch(t, genesisState.AirdropDistrubitions, got.AirdropDistrubitions)
	require.ElementsMatch(t, genesisState.AirdropClaimsLeft, got.AirdropClaimsLeft)
	// this line is used by starport scaffolding # genesis/test/assert
}
