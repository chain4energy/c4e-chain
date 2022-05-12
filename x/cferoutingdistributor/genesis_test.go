package cferoutingdistributor_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CferoutingdistributorKeeper(t)
	cferoutingdistributor.InitGenesis(ctx, *k, genesisState, nil /* TODO AccountKeeper*/)
	got := cferoutingdistributor.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
