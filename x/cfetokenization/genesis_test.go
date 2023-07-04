package cfetokenization_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfetokenization"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CertificateTypeList: []types.CertificateType{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CertificateTypeCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CfetokenizationKeeper(t)
	cfetokenization.InitGenesis(ctx, *k, genesisState)
	got := cfetokenization.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CertificateTypeList, got.CertificateTypeList)
	require.Equal(t, genesisState.CertificateTypeCount, got.CertificateTypeCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
