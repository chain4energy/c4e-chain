package cfefingerprint_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	energychain "github.com/chain4energy/c4e-chain/x/cfefingerprint"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CfefingerprintKeeper(t)
	energychain.InitGenesis(ctx, *k, genesisState)
	got := energychain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
