package app

import (
	"context"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"time"

	"github.com/stretchr/testify/require"

	c4eapp "github.com/chain4energy/c4e-chain/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	testcfeclaim "github.com/chain4energy/c4e-chain/testutil/module/cfeclaim"
	testcfedistributor "github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	testcfeminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	testcfevesting "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	abci "github.com/cometbft/cometbft/abci/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func Setup(isCheckTx bool) *c4eapp.App {
	app, _ := SetupWithValidatorsAmount(isCheckTx, testenv.DefaultTestDenom, 1)
	return app
}

func SetupAndGetValidatorsRelatedCoins(isCheckTx bool, balances ...banktypes.Balance) (*c4eapp.App, sdk.Coin) {
	return SetupWithValidatorsAmount(isCheckTx, testenv.DefaultTestDenom, 1, balances...)
}

func SetupApp(initBlock int64) (*c4eapp.App, sdk.Context, sdk.Coin) {
	return SetupAppWithTime(initBlock, testenv.TestEnvTime)
}

func SetupAppWithTime(initBlock int64, initTime time.Time, balances ...banktypes.Balance) (*c4eapp.App, sdk.Context, sdk.Coin) {
	app, coins := SetupAndGetValidatorsRelatedCoins(false, balances...)
	header := tmproto.Header{}
	header.Height = initBlock
	header.Time = initTime
	ctx := app.BaseApp.NewContext(false, header)
	return app, ctx, coins
}

func SetupTestApp(t require.TestingT) *TestHelper {
	return SetupTestAppWithHeightAndTime(t, 1, testenv.TestEnvTime)
}

func SetupTestAppWithHeight(t require.TestingT, initBlock int64) *TestHelper {
	return SetupTestAppWithHeightAndTime(t, initBlock, testenv.TestEnvTime)
}

func SetupTestAppWithHeightAndTime(t require.TestingT, initBlock int64, initTime time.Time, balances ...banktypes.Balance) *TestHelper {
	app, ctx, coins := SetupAppWithTime(initBlock, initTime, balances...)
	testHelper := newTestHelper(t, ctx, app, initTime, coins)
	return testHelper
}

type TestHelper struct {
	App                   *c4eapp.App
	Context               sdk.Context
	WrappedContext        context.Context
	InitialValidatorsCoin sdk.Coin
	InitTime              time.Time
	BankUtils             *testcosmos.ContextBankUtils
	AuthUtils             *testcosmos.ContextAuthUtils
	StakingUtils          *testcosmos.ContextStakingUtils
	GovUtils              *testcosmos.ContextGovUtils
	DistributionUtils     *testcosmos.ContextDistributionUtils
	FeegrantUtils         *testcosmos.ContextFeegrantUtils
	C4eVestingUtils       *testcfevesting.ContextC4eVestingUtils
	C4eMinterUtils        *testcfeminter.ContextC4eMinterUtils
	C4eDistributorUtils   *testcfedistributor.ContextC4eDistributorUtils
	C4eClaimUtils         *testcfeclaim.ContextC4eClaimUtils
}

func newTestHelper(t require.TestingT, ctx sdk.Context, app *c4eapp.App, initTime time.Time, initialValidatorsCoin sdk.Coin) *TestHelper {
	maccPerms := testcosmos.AddHelperModuleAccountPermissions(c4eapp.GetMaccPerms())

	helperAk := authkeeper.NewAccountKeeper(
		app.AppCodec(), app.GetKey(authtypes.StoreKey), authtypes.ProtoBaseAccount, maccPerms, c4eapp.AccountAddressPrefix, appparams.GetAuthority(),
	)

	moduleAccAddrs := testcosmos.AddHelperModuleAccountAddr(app.ModuleAccountAddrs())

	helperBk := bankkeeper.NewBaseKeeper(
		app.AppCodec(), app.GetKey(banktypes.StoreKey), helperAk, moduleAccAddrs, appparams.GetAuthority(),
	)

	testHelper := TestHelper{
		App:                   app,
		Context:               ctx,
		WrappedContext:        sdk.WrapSDKContext(ctx),
		InitialValidatorsCoin: initialValidatorsCoin,
		InitTime:              initTime,
	}

	var testHelperP testenv.TestContext = &testHelper

	bankUtils := testcosmos.NewContextBankUtils(t, testHelperP, &helperAk, helperBk)

	testHelper.BankUtils = bankUtils
	testHelper.AuthUtils = testcosmos.NewContextAuthUtils(t, testHelper, &helperAk, &bankUtils.BankUtils)
	testHelper.StakingUtils = testcosmos.NewContextStakingUtils(t, testHelper, *app.StakingKeeper, &bankUtils.BankUtils)
	testHelper.GovUtils = testcosmos.NewContextGovUtils(t, &testHelper, &app.GovKeeper)
	testHelper.FeegrantUtils = testcosmos.NewContextFeegrantUtils(t, &testHelper, &app.FeeGrantKeeper)
	testHelper.DistributionUtils = testcosmos.NewContextDistributionUtils(t, &testHelper, &app.DistrKeeper)

	testHelper.C4eVestingUtils = testcfevesting.NewContextC4eVestingUtils(t, testHelperP, &app.CfevestingKeeper, &app.AccountKeeper, &app.BankKeeper, app.StakingKeeper, &bankUtils.BankUtils, &testHelper.AuthUtils.AuthUtils)
	testHelper.C4eMinterUtils = testcfeminter.NewContextC4eMinterUtils(t, testHelperP, &app.CfeminterKeeper, &app.AccountKeeper, &bankUtils.BankUtils)
	testHelper.C4eDistributorUtils = testcfedistributor.NewContextC4eDistributorUtils(t, testHelperP, &app.CfedistributorKeeper, &app.AccountKeeper, &bankUtils.BankUtils)
	testHelper.C4eClaimUtils = testcfeclaim.NewContextC4eClaimUtils(t, &testHelper, &app.CfeclaimKeeper, &app.CfevestingKeeper, &app.AccountKeeper, &bankUtils.BankUtils, &testHelper.StakingUtils.StakingUtils, &testHelper.GovUtils.GovUtils, &testHelper.FeegrantUtils.FeegrantUtils, &testHelper.DistributionUtils.DistributionUtils)

	return &testHelper
}

func (th TestHelper) GetContext() sdk.Context {
	return th.Context
}

func (th TestHelper) GetWrappedContext() context.Context {
	return th.WrappedContext
}

func (th *TestHelper) SetContextBlockHeight(height int64) {
	th.setContext(th.Context.WithBlockHeight(height))
}

func (th *TestHelper) SetContextBlockTime(time time.Time) {
	th.setContext(th.Context.WithBlockTime(time))
}

func (th *TestHelper) SetContextBlockHeightAndTime(height int64, time time.Time) {
	th.setContext(th.Context.WithBlockHeight(height).WithBlockTime(time))
}

func (th *TestHelper) IncrementContextBlockHeightAndSetTime(time time.Time) {
	th.setContext(th.Context.WithBlockHeight(th.Context.BlockHeight() + 1).WithBlockTime(time))
}

func (th *TestHelper) IncrementContextBlockHeight() {
	th.setContext(th.Context.WithBlockHeight(th.Context.BlockHeight() + 1))
}

func (th *TestHelper) SetContextBlockHeightAndAddTime(height int64, durationToAdd time.Duration) {
	th.setContext(th.Context.WithBlockHeight(height).WithBlockTime(th.Context.BlockTime().Add(durationToAdd)))
}

func (th *TestHelper) BeginBlocker(req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return th.App.BeginBlocker(th.Context, req)
}

func (th *TestHelper) EndBlocker(req abci.RequestEndBlock) abci.ResponseEndBlock {
	return th.App.EndBlocker(th.Context, req)
}

func (th *TestHelper) setContext(ctx sdk.Context) {
	th.Context = ctx
	th.WrappedContext = sdk.WrapSDKContext(ctx)
}
