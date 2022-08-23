package cfeenergybank_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeenergybank"
	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EnergyTokenList: []types.EnergyToken{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EnergyTokenCount: 2,
		TokenParamsList: []types.TokenParams{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CfeEnergybankKeeper(t)
	cfeenergybank.InitGenesis(ctx, *k, genesisState)
	got := cfeenergybank.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EnergyTokenList, got.EnergyTokenList)
	require.Equal(t, genesisState.EnergyTokenCount, got.EnergyTokenCount)
	require.ElementsMatch(t, genesisState.TokenParamsList, got.TokenParamsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
