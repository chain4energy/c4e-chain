package app

import (
	"time"

	c4eapp "github.com/chain4energy/c4e-chain/app"
	testcommon "github.com/chain4energy/c4e-chain/testutil/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func Setup(isCheckTx bool) *c4eapp.App {
	app, _ := SetupWithValidatorsAmount(isCheckTx,testcommon.Denom, 1)
	return app
}

func SetupAndGetValidatorsRelatedCoins(isCheckTx bool) (*c4eapp.App, sdk.Coin) {
	return SetupWithValidatorsAmount(isCheckTx,testcommon.Denom, 1)
}

func SetupApp(initBlock int64) (*c4eapp.App, sdk.Context, sdk.Coin) {
	return SetupAppWithTime(initBlock, testcommon.TestEnvTime)
}

func SetupAppWithTime(initBlock int64, initTime time.Time) (*c4eapp.App, sdk.Context, sdk.Coin) {
	app, coins := SetupAndGetValidatorsRelatedCoins(false)
	header := tmproto.Header{}
	header.Height = initBlock
	header.Time = initTime
	ctx := app.BaseApp.NewContext(false, header)
	return app, ctx, coins
}
