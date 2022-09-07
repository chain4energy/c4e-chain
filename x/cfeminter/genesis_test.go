package cfeminter_test

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/app"

	"github.com/chain4energy/c4e-chain/x/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	routingdistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
)

const iterationText = "iterarion %d"

const PeriodDuration = time.Duration(345600000000 * 1000000)
const Year = time.Hour * 24 * 365
const SecondsInYear = int32(3600 * 24 * 365)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:      types.NewParams("myc4e", createMinter(time.Now())),
		MinterState: types.MinterState{CurrentPosition: 9, AmountMinted: sdk.NewInt(12312)},

		// this line is used by starport scaffolding # genesis/test/state

	}

	app := testapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	cfeminter.InitGenesis(ctx, app.CfeminterKeeper, app.AccountKeeper, genesisState)
	got := cfeminter.ExportGenesis(ctx, app.CfeminterKeeper)
	require.NotNil(t, got)

	require.EqualValues(t, genesisState.Params.MintDenom, got.Params.MintDenom)
	testminter.CompareMinters(t, genesisState.Params.Minter, got.Params.Minter)
	require.EqualValues(t, genesisState.MinterState, got.MinterState)

	// this line is used by starport scaffolding # genesis/test/assert
}

func TestOneYearLinear(t *testing.T) {
	totalSupply := int64(40000000000000)
	commontestutils.AddHelperModuleAccountPerms()
	now := time.Now()
	yearFromNow := now.Add(time.Hour * 24 * 365)
	minter := types.Minter{
		Start: now,
		Periods: []*types.MintingPeriod{
			{Position: 1, PeriodEnd: &yearFromNow, Type: types.TIME_LINEAR_MINTER,
				TimeLinearMinter: &types.TimeLinearMinter{Amount: sdk.NewInt(totalSupply)}},
			{Position: 2, Type: types.NO_MINTING},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(commontestutils.Denom, minter),
		MinterState: types.MinterState{CurrentPosition: 1, AmountMinted: sdk.NewInt(0)},
	}

	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: now})
	cfeminter.InitGenesis(ctx, app.CfeminterKeeper, app.AccountKeeper, genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	commontestutils.AddCoinsToAccount(uint64(totalSupply), ctx, app, acountsAddresses[0])

	inflation, err := app.CfeminterKeeper.GetCurrentInflation(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewDec(1), inflation)
	state := app.CfeminterKeeper.GetMinterState(ctx)
	require.EqualValues(t, int32(1), state.CurrentPosition)
	require.EqualValues(t, sdk.ZeroInt(), state.AmountMinted)
	commontestutils.VerifyModuleAccountBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, sdk.ZeroInt())

	numOfHours := 365 * 24
	for i := 1; i <= numOfHours; i++ {
		ctx = ctx.WithBlockHeight(int64(i)).WithBlockTime(ctx.BlockTime().Add(time.Hour))
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
		app.EndBlocker(ctx, abci.RequestEndBlock{})

		expectedMinted := totalSupply * int64(i) / int64(numOfHours)
		expectedInflation := sdk.NewDec(totalSupply).QuoInt64(totalSupply + expectedMinted)

		commontestutils.VerifyModuleAccountBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, sdk.NewInt(expectedMinted))

		inflation, err := app.CfeminterKeeper.GetCurrentInflation(ctx)
		require.NoError(t, err)
		if i < numOfHours {
			require.EqualValuesf(t, expectedInflation, inflation, iterationText, i)
			state := app.CfeminterKeeper.GetMinterState(ctx)
			require.EqualValues(t, int32(1), state.CurrentPosition)
			require.EqualValues(t, sdk.NewInt(expectedMinted), state.AmountMinted)
		} else {
			require.EqualValuesf(t, sdk.ZeroDec(), inflation, iterationText, i)
			state := app.CfeminterKeeper.GetMinterState(ctx)
			require.EqualValues(t, int32(2), state.CurrentPosition)
			require.EqualValues(t, sdk.ZeroInt(), state.AmountMinted)
		}

	}

	commontestutils.VerifyModuleAccountBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, sdk.NewInt(totalSupply))
	supp := app.BankKeeper.GetSupply(ctx, commontestutils.Denom)
	require.EqualValues(t, sdk.NewInt(2*totalSupply), supp.Amount)

}

func TestFewYearsPeriodicReduction(t *testing.T) {
	totalSupply := int64(400000000000000)
	startAmountYearly := int64(40000000000000)
	commontestutils.AddHelperModuleAccountPerms()
	now := time.Now()
	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(startAmountYearly), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	minter := types.Minter{
		Start: now,
		Periods: []*types.MintingPeriod{
			{Position: 1, Type: types.PERIODIC_REDUCTION_MINTER,
				PeriodicReductionMinter: &pminter},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(commontestutils.Denom, minter),
		MinterState: types.MinterState{CurrentPosition: 1, AmountMinted: sdk.NewInt(0)},
	}

	app := testapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: now})
	cfeminter.InitGenesis(ctx, app.CfeminterKeeper, app.AccountKeeper, genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	commontestutils.AddCoinsToAccount(uint64(totalSupply), ctx, app, acountsAddresses[0])

	inflation, err := app.CfeminterKeeper.GetCurrentInflation(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.MustNewDecFromStr("0.1"), inflation)
	state := app.CfeminterKeeper.GetMinterState(ctx)
	require.EqualValues(t, int32(1), state.CurrentPosition)
	require.EqualValues(t, sdk.ZeroInt(), state.AmountMinted)
	commontestutils.VerifyModuleAccountBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 4 * year
	amountYearly := startAmountYearly
	prevPeriodMinted := int64(0)

	for periodsCount := 1; periodsCount <= 5; periodsCount++ {
		for i := 1; i <= numOfHours; i++ {
			ctx = ctx.WithBlockHeight(int64(i)).WithBlockTime(ctx.BlockTime().Add(24 * time.Hour))
			app.BeginBlocker(ctx, abci.RequestBeginBlock{})
			app.EndBlocker(ctx, abci.RequestEndBlock{})

			expectedMinted := amountYearly * int64(i) / int64(year)
			expectedInflation := sdk.NewDec(amountYearly).QuoInt64(totalSupply + prevPeriodMinted + expectedMinted)

			commontestutils.VerifyModuleAccountBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, sdk.NewInt(prevPeriodMinted+expectedMinted))

			inflation, err := app.CfeminterKeeper.GetCurrentInflation(ctx)
			require.NoError(t, err)
			if i < numOfHours {
				require.EqualValuesf(t, expectedInflation, inflation, iterationText, i)
				state := app.CfeminterKeeper.GetMinterState(ctx)
				require.EqualValues(t, int32(1), state.CurrentPosition)
				require.EqualValues(t, sdk.NewInt(prevPeriodMinted+expectedMinted), state.AmountMinted)
			} else {
				require.EqualValuesf(t, expectedInflation.QuoInt64(2), inflation, iterationText, i)
				state := app.CfeminterKeeper.GetMinterState(ctx)
				require.EqualValues(t, int32(1), state.CurrentPosition)
				require.EqualValues(t, sdk.NewInt(prevPeriodMinted+expectedMinted), state.AmountMinted)
				prevPeriodMinted += expectedMinted
			}

		}
		amountYearly = amountYearly / 2
	}
	expectedMinted := int64(310000000000000)
	commontestutils.VerifyModuleAccountBalanceByName(routingdistributortypes.DistributorMainAccount, ctx, app, t, sdk.NewInt(expectedMinted))
	supp := app.BankKeeper.GetSupply(ctx, commontestutils.Denom)
	require.EqualValues(t, sdk.NewInt(totalSupply+expectedMinted), supp.Amount)

}

func createMinter(startTime time.Time) types.Minter {
	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	minter := types.Minter{Start: startTime, Periods: periods}
	return minter
}
