package cfeev_test

import (
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeev"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EnergyTransferOffers: []types.EnergyTransferOffer{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EnergyTransferOfferCount: 2,
		EnergyTransfers: []types.EnergyTransfer{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EnergyTransferCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.CfeevKeeper(t)
	cfeev.InitGenesis(ctx, *k, genesisState)
	got := cfeev.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EnergyTransferOffers, got.EnergyTransferOffers)
	require.Equal(t, genesisState.EnergyTransferOfferCount, got.EnergyTransferOfferCount)
	require.ElementsMatch(t, genesisState.EnergyTransfers, got.EnergyTransfers)
	require.Equal(t, genesisState.EnergyTransferCount, got.EnergyTransferCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
