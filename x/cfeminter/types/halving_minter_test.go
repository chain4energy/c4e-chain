package types_test

import (
	"fmt"
	"testing"

	"github.com/chain4energy/c4e-chain/app"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestHalvingMinter_NextCointCount2(t *testing.T) {
	//app := simapp.NewSimApp()

}

func TestHalvingMinter_NextCointCount(t *testing.T) {

	halvingMinter := types.InitialHalvingMinter()

	tests := []struct {
		blockHeight   int64
		expProvisions int64
	}{
		//{1999, 100},
		//{2000, 50},
		{3999, 50},
		//{4000, 25},
	}

	for i, tc := range tests {

		newCoinCount := halvingMinter.NextCointCount(tc.blockHeight)

		require.True(t, tc.expProvisions == newCoinCount,
			"test: %v\n\tExp: %v\n\tGot: %v\n",
			i, tc.expProvisions, newCoinCount)
	}
}

type MintKeeperTestSuite struct {
	suite.Suite

	app              *simapp.SimApp
	ctx              sdk.Context
	legacyQuerierCdc *codec.AminoCodec
}

//func (suite *MintKeeperTestSuite) SetupTest() {
//	app := simapp.Setup(suite.T(), true)
//	ctx := app.BaseApp.NewContext(true, tmproto.Header{})
//
//	app.MintKeeper.SetParams(ctx, types.DefaultParams())
//	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())
//
//	legacyQuerierCdc := codec.NewAminoCodec(app.LegacyAmino())
//
//	suite.app = app
//	suite.ctx = ctx
//	suite.legacyQuerierCdc = legacyQuerierCdc
//
//	height := ctx.BlockHeight()
//}

func TestGenesis(t *testing.T) {
	//genesisState := types.GenesisState{
	//	Params: types.DefaultParams(),
	//
	//	// this line is used by starport scaffolding # genesis/test/state
	//}
	//
	//k, ctx := keepertest.EnergyminterKeeper(t)
	//energyminter.InitGenesis(ctx, *k, genesisState)
	//got := energyminter.ExportGenesis(ctx, *k)
	//require.NotNil(t, got)
	//
	//nullify.Fill(&genesisState)
	//nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert

	app := app.Setup(false)

	//ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	for i := 0; i < 10; i++ {
		header := tmproto.Header{Height: app.LastBlockHeight() + 1}
		app.BeginBlock(abci.RequestBeginBlock{Header: header})
		fmt.Println(app.LastBlockHeight())

	}

	fmt.Println(app.LastBlockHeight())

	//app.BankKeeper.GetSupply(ctx, "stake")

	//require.Equal(t, app.BankKeeper.GetSupply(ctx, "stake").Amount.String(), 2, "asd")
}
