package common

import (
	"time"

	"github.com/chain4energy/c4e-chain/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var TestEnvTime = time.Now()

func SetupApp(initBlock int64) (*app.App, sdk.Context) {
	return SetupAppWithTime(initBlock, TestEnvTime)
}

func SetupAppWithTime(initBlock int64, initTime time.Time) (*app.App, sdk.Context) {
	app := app.Setup(false)
	header := tmproto.Header{}
	header.Height = initBlock
	header.Time = initTime
	ctx := app.BaseApp.NewContext(false, header)
	return app, ctx
}