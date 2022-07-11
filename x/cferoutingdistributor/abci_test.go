package cferoutingdistributor_test

import (
	testapp "github.com/chain4energy/c4e-chain/app"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func prepareBurningDistributor() types.RoutingDistributor {
	destAccount := types.Account{
		Address:         "c4e_distributor",
		IsModuleAccount: true,
		LeftoverCoin:    sdk.MustNewDecFromStr("0"),
	}

	burnShare := types.BurnShare{
		Percent:      sdk.MustNewDecFromStr("51"),
		LeftoverCoin: sdk.MustNewDecFromStr("0"),
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     nil,
		BurnShare: burnShare,
	}

	distributor1 := types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []string{"fee_collector"},
		Destination: destination,
		Order:       0,
	}

	routingDistributor := types.RoutingDistributor{
		SubDistributor:           []types.SubDistributor{distributor1},
		ModuleAccounts:           nil,
		RemainsCoinModuleAccount: "remains",
	}

	return routingDistributor
}

func TestBurningDistributor(t *testing.T) {

	perms := []string{authtypes.Minter, authtypes.Burner}
	collector := "fee_collector"
	denom := "uc4e"
	testapp.AddMaccPerms(collector, perms)
	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	//prepare module account with coin to distribute fee_collector 1017
	cointToMint := sdk.NewCoin(denom, sdk.NewInt(1017))
	app.BankKeeper.MintCoins(ctx, collector, sdk.NewCoins(cointToMint))

	app.CferoutingdistributorKeeper.SetRoutingDistributor(ctx, prepareBurningDistributor())
	ctx = ctx.WithBlockHeight(int64(2))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{})

	//app.EndBlocker(ctx, abci.RequestEndBlock{})

	//
	//coin on "c4e_distributor" should be equal 498, remains: 1 and 0.33 on remains
	coinRemains := app.CferoutingdistributorKeeper.GetRoutingDistributorr(ctx).SubDistributor[0].Destination.Account.LeftoverCoin
	require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains)

	coinOnRemainAccount := app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "remains")
	require.EqualValues(t, sdk.NewInt(1), coinOnRemainAccount.AmountOf(denom))

	coinAfterDistribution :=
		app.CferoutingdistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "c4e_distributor")

	require.EqualValues(t, sdk.NewInt(498), coinAfterDistribution.AmountOf(denom))

}
