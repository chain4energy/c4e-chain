package cferoutingdistributor_test

import (
	testapp "github.com/chain4energy/c4e-chain/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:             types.DefaultParams(),
		RoutingDistributor: prepareBurningDistributor(),
	}

	k, ctx := keepertest.CferoutingdistributorKeeper(t)
	app := testapp.Setup(false)

	cferoutingdistributor.InitGenesis(ctx, *k, genesisState, app.AccountKeeper)
	got := cferoutingdistributor.ExportGenesis(ctx, *k)

	require.EqualValues(t, sdk.MustNewDecFromStr("51"), got.RoutingDistributor.SubDistributor[0].Destination.BurnShare.Percent)
	require.EqualValues(t, "remains", got.RoutingDistributor.RemainsCoinModuleAccount)
	require.EqualValues(t, "c4e_distributor", got.RoutingDistributor.SubDistributor[0].Destination.Account.Address)
	require.NotNil(t, got)
	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
