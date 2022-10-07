package app

import (
	"time"

	"github.com/chain4energy/c4e-chain/app"
	testcommon "github.com/chain4energy/c4e-chain/testutil/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func SetupApp(initBlock int64) (*app.App, sdk.Context) {
	return SetupAppWithTime(initBlock, testcommon.TestEnvTime)
}

func SetupAppWithTime(initBlock int64, initTime time.Time) (*app.App, sdk.Context) {
	app := Setup(false)
	header := tmproto.Header{}
	header.Height = initBlock
	header.Time = initTime
	ctx := app.BaseApp.NewContext(false, header)
	return app, ctx
}
