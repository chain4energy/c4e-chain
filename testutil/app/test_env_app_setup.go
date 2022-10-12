package app

import (
	"testing"
	"time"

	c4eapp "github.com/chain4energy/c4e-chain/app"
	testcommon "github.com/chain4energy/c4e-chain/testutil/common"
	testcfevesting "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	abci "github.com/tendermint/tendermint/abci/types"
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

func SetupTestApp(t *testing.T) *TestHelper {
	return SetupTestAppWithHeightAndTime(t, 1, testcommon.TestEnvTime)
}

func SetupTestAppWithHeight(t *testing.T, initBlock int64) *TestHelper {
	return SetupTestAppWithHeightAndTime(t, initBlock, testcommon.TestEnvTime)
}

func SetupTestAppWithHeightAndTime(t *testing.T, initBlock int64, initTime time.Time, balances ...banktypes.Balance) *TestHelper {
	app, ctx, coins := SetupAppWithTime(initBlock, initTime, balances...)
	testHelper := newTestHelper(t, ctx, app, initTime, coins)
	return testHelper
}

type TestHelper struct {
	App                   *c4eapp.App
	Context               sdk.Context
	InitialValidatorsCoin sdk.Coin
	InitTime              time.Time
	BankUtils             *testcommon.ContextBankUtils
	AuthUtils             *testcommon.ContextAuthUtils
	StakingUtils          *testcommon.ContextStakingUtils
	C4eVestingUtils       *testcfevesting.ContextC4eVestingUtils
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

	testHelper := TestHelper{
		App:                   app,
		Context:               ctx,
		InitialValidatorsCoin: initialValidatorsCoin,
		InitTime:              initTime,
	}

	var testHelperP testcommon.TestContext = &testHelper

	bankUtils := testcommon.NewContextBankUtils(t, testHelper, &helperAk, helperBk)

	testHelper.BankUtils = bankUtils
	testHelper.AuthUtils = testcommon.NewContextAuthUtils(testHelper, &helperAk, &bankUtils.BankUtils)
	testHelper.StakingUtils = testcommon.NewContextStakingUtils(t, testHelper, app.StakingKeeper, &bankUtils.BankUtils)
	testHelper.C4eVestingUtils = testcfevesting.NewContextC4eVestingUtils(t, testHelperP, &app.CfevestingKeeper, &bankUtils.BankUtils)
	return &testHelper
}

func (th TestHelper) GetContext() sdk.Context {
	return th.Context
}

func (th *TestHelper) SetContextBlockHeight(height int64) {
	th.Context = th.Context.WithBlockHeight(height)
}

func (th *TestHelper) SetContextBlockTime(time time.Time) {
	th.Context = th.Context.WithBlockTime(time)
}

func (th *TestHelper) SetContextBlockHeightAndTime(height int64, time time.Time) {
	th.Context = th.Context.WithBlockHeight(height).WithBlockTime(time)
}

func (th *TestHelper) IncrementContextBlockHeightAndSetTime(time time.Time) {
	th.Context = th.Context.WithBlockHeight(th.Context.BlockHeight() + 1).WithBlockTime(time)
}

func (th *TestHelper) IncrementContextBlockHeight() {
	th.Context = th.Context.WithBlockHeight(th.Context.BlockHeight() + 1)
}

func (th *TestHelper) SetContextBlockHeightAndAddTime(height int64, durationToAdd time.Duration) {
	th.Context = th.Context.WithBlockHeight(height).WithBlockTime(th.Context.BlockTime().Add(durationToAdd))
}

func (th *TestHelper) BeginBlocker(req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return th.App.BeginBlocker(th.Context, req)
}

func (th *TestHelper) EndBlocker(req abci.RequestEndBlock) abci.ResponseEndBlock {
	return th.App.EndBlocker(th.Context, req)
}
