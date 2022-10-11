package app

import (
	"testing"
	"time"

	c4eapp "github.com/chain4energy/c4e-chain/app"
	testcommon "github.com/chain4energy/c4e-chain/testutil/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func Setup(isCheckTx bool) *c4eapp.App {
	app, _ := SetupWithValidatorsAmount(isCheckTx, testcommon.DefaultTestDenom, 1)
	return app
}

func SetupAndGetValidatorsRelatedCoins(isCheckTx bool, balances ...banktypes.Balance) (*c4eapp.App, sdk.Coin) {
	return SetupWithValidatorsAmount(isCheckTx, testcommon.DefaultTestDenom, 1, balances...)
}

func SetupApp(initBlock int64) (*c4eapp.App, sdk.Context, sdk.Coin) {
	return SetupAppWithTime(initBlock, testcommon.TestEnvTime)
}

func SetupAppWithTime(initBlock int64, initTime time.Time, balances ...banktypes.Balance) (*c4eapp.App, sdk.Context, sdk.Coin) {
	app, coins := SetupAndGetValidatorsRelatedCoins(false, balances...)
	header := tmproto.Header{}
	header.Height = initBlock
	header.Time = initTime
	ctx := app.BaseApp.NewContext(false, header)
	return app, ctx, coins
}

func SetupTestApp(t *testing.T) (*TestHelper, sdk.Context) {
	return SetupTestAppWithHeightAndTime(t, 1, testcommon.TestEnvTime)
}

func SetupTestAppWithHeight(t *testing.T, initBlock int64) (*TestHelper, sdk.Context) {
	return SetupTestAppWithHeightAndTime(t, initBlock, testcommon.TestEnvTime)
}

func SetupTestAppWithHeightAndTime(t *testing.T, initBlock int64, initTime time.Time, balances ...banktypes.Balance) (*TestHelper, sdk.Context) {
	app, ctx, coins := SetupAppWithTime(initBlock, initTime, balances...)
	testHelper := newTestHelper(t, ctx, app, initTime, coins)
	return testHelper, ctx
}

type TestHelper struct {
	App                   *c4eapp.App
	InitialValidatorsCoin sdk.Coin
	InitTime              time.Time
	BankUtils             *testcommon.BankUtils
	AuthUtils             *testcommon.AuthUtils
	StakingUtils          *testcommon.StakingUtils
}

func newTestHelper(t *testing.T, ctx sdk.Context, app *c4eapp.App, initTime time.Time, initialValidatorsCoin sdk.Coin) *TestHelper {
	maccPerms := testcommon.AddHelperModuleAccountPermissions(c4eapp.GetMaccPerms())

	helperAk := authkeeper.NewAccountKeeper(
		app.AppCodec(), app.GetKey(authtypes.StoreKey), app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)

	moduleAccAddrs := testcommon.AddHelperModuleAccountAddr(app.ModuleAccountAddrs())

	helperBk := bankkeeper.NewBaseKeeper(
		app.AppCodec(), app.GetKey(banktypes.StoreKey), helperAk, app.GetSubspace(banktypes.ModuleName), moduleAccAddrs,
	)
	bankUtils := testcommon.NewBankUtils(t, ctx, &helperAk, helperBk)

	testHelper := TestHelper{
		App:                   app,
		InitialValidatorsCoin: initialValidatorsCoin,
		InitTime:              initTime,
		BankUtils:             bankUtils,
		AuthUtils:             testcommon.NewAuthUtils(&helperAk, bankUtils),
		StakingUtils:          testcommon.NewStakingUtils(t, app.StakingKeeper, bankUtils),
	}
	return &testHelper
}
