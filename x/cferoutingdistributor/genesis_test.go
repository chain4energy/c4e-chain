package cferoutingdistributor_test

import (
	"testing"
)

func TestGenesis(t *testing.T) {
	//genesisState := types.GenesisState{
	//	Params:             types.DefaultParams(),
	//	RoutingDistributor: prepareBurningDistributor(),
	//}
	//
	//k, ctx := keepertest.CferoutingdistributorKeeper(t)
	//app := testapp.Setup(false)
	//
	//cferoutingdistributor.InitGenesis(ctx, *k, genesisState, app.AccountKeeper)
	//got := cferoutingdistributor.ExportGenesis(ctx, *k)
	//
	//require.EqualValues(t, sdk.MustNewDecFromStr("51"), got.RoutingDistributor.SubDistributor[0].Destination.BurnShare.Percent)
	//require.EqualValues(t, "remains", got.RoutingDistributor.RemainsCoinModuleAccount)
	//require.EqualValues(t, "c4e_distributor", got.RoutingDistributor.SubDistributor[0].Destination.Account.Address)
	//require.NotNil(t, got)
	//nullify.Fill(&genesisState)
	//nullify.Fill(got)
}
