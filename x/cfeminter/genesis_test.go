package cfeminter_test

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	"github.com/chain4energy/c4e-chain/x/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	routingdistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
)

const iterationText = "iterarion %d"

const PeriodDuration = time.Duration(345600000000 * 1000000)
const Year = time.Hour * 24 * 365
const SecondsInYear = int32(3600 * 24 * 365)

func TestGenesis(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	mintTime, _ := time.Parse(layout, str)
	genesisState := types.GenesisState{
		Params: types.NewParams("myc4e", createMinter(time.Now())),
		MinterState: types.MinterState{
			Position:                    9,
			AmountMinted:                sdk.NewInt(12312),
			RemainderToMint:             sdk.MustNewDecFromStr("1233.546"),
			RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("7654.423"),
			LastMintBlockTime:           mintTime,
		},

		// this line is used by starport scaffolding # genesis/test/state

	}
	testHelper, ctx := testapp.SetupTestApp(t)

	cfeminter.InitGenesis(ctx, testHelper.App.CfeminterKeeper, testHelper.App.AccountKeeper, genesisState)
	got := cfeminter.ExportGenesis(ctx, testHelper.App.CfeminterKeeper)
	require.NotNil(t, got)

	require.EqualValues(t, genesisState.Params.MintDenom, got.Params.MintDenom)
	testminter.CompareMinters(t, genesisState.Params.Minter, got.Params.Minter)
	require.EqualValues(t, genesisState.MinterState, got.MinterState)

	// this line is used by starport scaffolding # genesis/test/assert
}

func TestGenesisWithHistory(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	mintTime, _ := time.Parse(layout, str)
	genesisState := types.GenesisState{
		Params: types.NewParams("myc4e", createMinter(time.Now())),
		MinterState: types.MinterState{
			Position:                    9,
			AmountMinted:                sdk.NewInt(12312),
			RemainderToMint:             sdk.MustNewDecFromStr("1233.546"),
			RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("7654.423"),
			LastMintBlockTime:           mintTime,
		},
		StateHistory: createHistory(),
		// this line is used by starport scaffolding # genesis/test/state

	}

	testHelper, ctx := testapp.SetupTestApp(t)

	cfeminter.InitGenesis(ctx, testHelper.App.CfeminterKeeper, testHelper.App.AccountKeeper, genesisState)
	got := cfeminter.ExportGenesis(ctx, testHelper.App.CfeminterKeeper)
	require.NotNil(t, got)

	require.EqualValues(t, genesisState.Params.MintDenom, got.Params.MintDenom)
	testminter.CompareMinters(t, genesisState.Params.Minter, got.Params.Minter)
	require.EqualValues(t, genesisState.MinterState, got.MinterState)
	require.EqualValues(t, len(genesisState.StateHistory), len(got.StateHistory))

	for i := 0; i < len(genesisState.StateHistory); i++ {
		require.EqualValues(t, genesisState.StateHistory[i], got.StateHistory[i])
	}

	// this line is used by starport scaffolding # genesis/test/assert
}

func createHistory() []*types.MinterState {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	mintTime, _ := time.Parse(layout, str)

	history := make([]*types.MinterState, 0)
	state1 := types.MinterState{
		Position:                    0,
		AmountMinted:                sdk.NewInt(324),
		RemainderToMint:             sdk.MustNewDecFromStr("1243.221"),
		LastMintBlockTime:           mintTime,
		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("3124.543"),
	}
	str = "2016-06-12T11:35:46.371Z"
	mintTime, _ = time.Parse(layout, str)
	state2 := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(432),
		RemainderToMint:             sdk.MustNewDecFromStr("12433.221"),
		LastMintBlockTime:           mintTime,
		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("3284.543"),
	}
	return append(history, &state1, &state2)
}

func TestOneYearLinear(t *testing.T) {
	totalSupply := sdk.NewInt(40000000000000)

	testHelper, ctx := testapp.SetupTestApp(t)

	yearFromNow := testHelper.InitTime.Add(time.Hour * 24 * 365)
	minter := types.Minter{
		Start: testHelper.InitTime,
		Periods: []*types.MintingPeriod{
			{Position: 1, PeriodEnd: &yearFromNow, Type: types.TIME_LINEAR_MINTER,
				TimeLinearMinter: &types.TimeLinearMinter{Amount: totalSupply}},
			{Position: 2, Type: types.NO_MINTING},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(commontestutils.DefaultTestDenom, minter),
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0)},
	}

	cfeminter.InitGenesis(ctx, testHelper.App.CfeminterKeeper, testHelper.App.AccountKeeper, genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, consToAdd, acountsAddresses[0])

	inflation, err := testHelper.App.CfeminterKeeper.GetCurrentInflation(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewDec(1), inflation)
	state := testHelper.App.CfeminterKeeper.GetMinterState(ctx)
	require.EqualValues(t, int32(1), state.Position)
	require.EqualValues(t, sdk.ZeroInt(), state.AmountMinted)

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, sdk.ZeroInt())

	numOfHours := 365 * 24
	for i := 1; i <= numOfHours; i++ {
		ctx = ctx.WithBlockHeight(int64(i)).WithBlockTime(ctx.BlockTime().Add(time.Hour))
		testHelper.App.BeginBlocker(ctx, abci.RequestBeginBlock{})
		testHelper.App.EndBlocker(ctx, abci.RequestEndBlock{})

		expectedMinted := totalSupply.MulRaw(int64(i)).QuoRaw(int64(numOfHours))
		expectedInflation := sdk.NewDecFromInt(totalSupply).QuoInt(totalSupply.Add(expectedMinted))

		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, expectedMinted)

		inflation, err := testHelper.App.CfeminterKeeper.GetCurrentInflation(ctx)
		require.NoError(t, err)
		if i < numOfHours {
			require.EqualValuesf(t, expectedInflation, inflation, iterationText, i)
			state := testHelper.App.CfeminterKeeper.GetMinterState(ctx)
			require.EqualValues(t, int32(1), state.Position)
			require.EqualValues(t, expectedMinted, state.AmountMinted)
		} else {
			require.EqualValuesf(t, sdk.ZeroDec(), inflation, iterationText, i)
			state := testHelper.App.CfeminterKeeper.GetMinterState(ctx)
			require.EqualValues(t, int32(2), state.Position)
			require.EqualValues(t, sdk.ZeroInt(), state.AmountMinted)
		}

	}

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, totalSupply)

	supp := testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom)
	require.EqualValues(t, totalSupply.MulRaw(2), supp.Amount)

}

func TestFewYearsPeriodicReduction(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(40000000000000)
	testHelper, ctx := testapp.SetupTestApp(t)
	pminter := types.PeriodicReductionMinter{MintAmount: startAmountYearly, MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	minter := types.Minter{
		Start: testHelper.InitTime,
		Periods: []*types.MintingPeriod{
			{Position: 1, Type: types.PERIODIC_REDUCTION_MINTER,
				PeriodicReductionMinter: &pminter},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(commontestutils.DefaultTestDenom, minter),
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0)},
	}

	cfeminter.InitGenesis(ctx, testHelper.App.CfeminterKeeper, testHelper.App.AccountKeeper, genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, consToAdd, acountsAddresses[0])

	inflation, err := testHelper.App.CfeminterKeeper.GetCurrentInflation(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.MustNewDecFromStr("0.1"), inflation)
	state := testHelper.App.CfeminterKeeper.GetMinterState(ctx)
	require.EqualValues(t, int32(1), state.Position)
	require.EqualValues(t, sdk.ZeroInt(), state.AmountMinted)

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 4 * year
	amountYearly := startAmountYearly
	prevPeriodMinted := sdk.ZeroInt()

	for periodsCount := 1; periodsCount <= 5; periodsCount++ {
		for i := 1; i <= numOfHours; i++ {
			ctx = ctx.WithBlockHeight(int64(i)).WithBlockTime(ctx.BlockTime().Add(24 * time.Hour))
			testHelper.App.BeginBlocker(ctx, abci.RequestBeginBlock{})
			testHelper.App.EndBlocker(ctx, abci.RequestEndBlock{})

			expectedMinted := amountYearly.MulRaw(int64(i)).QuoRaw(int64(year))
			expectedInflation := sdk.NewDecFromInt(amountYearly).QuoInt(totalSupply.Add(prevPeriodMinted).Add(expectedMinted))

			testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, prevPeriodMinted.Add(expectedMinted))

			inflation, err := testHelper.App.CfeminterKeeper.GetCurrentInflation(ctx)
			require.NoError(t, err)
			if i < numOfHours {
				require.EqualValuesf(t, expectedInflation, inflation, iterationText, i)
				state := testHelper.App.CfeminterKeeper.GetMinterState(ctx)
				require.EqualValues(t, int32(1), state.Position)
				require.EqualValues(t, prevPeriodMinted.Add(expectedMinted), state.AmountMinted)
			} else {
				require.EqualValuesf(t, expectedInflation.QuoInt64(2), inflation, iterationText, i)
				state := testHelper.App.CfeminterKeeper.GetMinterState(ctx)
				require.EqualValues(t, int32(1), state.Position)
				require.EqualValues(t, prevPeriodMinted.Add(expectedMinted), state.AmountMinted)
				prevPeriodMinted = prevPeriodMinted.Add(expectedMinted)
			}

		}
		amountYearly = amountYearly.QuoRaw(2)
	}
	expectedMinted := sdk.NewInt(310000000000000)
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, expectedMinted)
	supp := testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom)
	require.EqualValues(t, totalSupply.Add(expectedMinted), supp.Amount)

}

func TestFewYearsPeriodicReductionInOneBlock(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(40000000000000)
	testHelper, ctx := testapp.SetupTestApp(t)

	pminter := types.PeriodicReductionMinter{MintAmount: startAmountYearly, MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	minter := types.Minter{
		Start: testHelper.InitTime,
		Periods: []*types.MintingPeriod{
			{Position: 1, Type: types.PERIODIC_REDUCTION_MINTER,
				PeriodicReductionMinter: &pminter},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(commontestutils.DefaultTestDenom, minter),
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0)},
	}

	cfeminter.InitGenesis(ctx, testHelper.App.CfeminterKeeper, testHelper.App.AccountKeeper, genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, consToAdd, acountsAddresses[0])

	inflation, err := testHelper.App.CfeminterKeeper.GetCurrentInflation(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.MustNewDecFromStr("0.1"), inflation)
	state := testHelper.App.CfeminterKeeper.GetMinterState(ctx)
	require.EqualValues(t, int32(1), state.Position)
	require.EqualValues(t, sdk.ZeroInt(), state.AmountMinted)

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 301 * year

	for i := 1; i <= numOfHours; i++ {
		ctx = ctx.WithBlockHeight(int64(i)).WithBlockTime(ctx.BlockTime().Add(24 * time.Hour))
	}

	testHelper.App.BeginBlocker(ctx, abci.RequestBeginBlock{})
	testHelper.App.EndBlocker(ctx, abci.RequestEndBlock{})

	expectedMinted := sdk.NewInt(320000000000000)
	state = testHelper.App.CfeminterKeeper.GetMinterState(ctx)
	require.EqualValues(t, int32(1), state.Position)
	require.EqualValues(t, expectedMinted, state.AmountMinted)

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, expectedMinted)

	supp := testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom)
	require.EqualValues(t, totalSupply.Add(expectedMinted), supp.Amount)

}

func TestFewYearsLinearAndPeriodicReductionInOneBlock(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(40000000000000)

	testHelper, ctx := testapp.SetupTestApp(t)

	tenYears := time.Duration(int64(time.Second) * int64(SecondsInYear) * 10)
	endTime1 := testHelper.InitTime.Add(tenYears)
	endTime2 := endTime1.Add(tenYears)

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(200000000000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000000000000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	pminter := types.PeriodicReductionMinter{MintAmount: startAmountYearly, MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	minter := types.Minter{
		Start: testHelper.InitTime,
		Periods: []*types.MintingPeriod{
			&period1,
			&period2,
			{Position: 3, Type: types.PERIODIC_REDUCTION_MINTER,
				PeriodicReductionMinter: &pminter},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(commontestutils.DefaultTestDenom, minter),
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0)},
	}

	cfeminter.InitGenesis(ctx, testHelper.App.CfeminterKeeper, testHelper.App.AccountKeeper, genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, consToAdd, acountsAddresses[0])

	inflation, err := testHelper.App.CfeminterKeeper.GetCurrentInflation(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.MustNewDecFromStr("0.05"), inflation)
	state := testHelper.App.CfeminterKeeper.GetMinterState(ctx)
	require.EqualValues(t, int32(1), state.Position)
	require.EqualValues(t, sdk.ZeroInt(), state.AmountMinted)

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 321 * year

	for i := 1; i <= numOfHours; i++ {
		ctx = ctx.WithBlockHeight(int64(i)).WithBlockTime(ctx.BlockTime().Add(24 * time.Hour))
	}

	testHelper.App.BeginBlocker(ctx, abci.RequestBeginBlock{})
	testHelper.App.EndBlocker(ctx, abci.RequestEndBlock{})

	expectedMintedPosition3 := sdk.NewInt(320000000000000)
	expectedMinted := sdk.NewInt(620000000000000)
	state = testHelper.App.CfeminterKeeper.GetMinterState(ctx)
	require.EqualValues(t, int32(3), state.Position)
	require.EqualValues(t, expectedMintedPosition3, state.AmountMinted)

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, expectedMinted)

	supp := testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom)
	require.EqualValues(t, totalSupply.Add(expectedMinted), supp.Amount)

	history := testHelper.App.CfeminterKeeper.GetAllMinterStateHistory(ctx)
	require.EqualValues(t, 2, len(history))

	expectedHist1 := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(200000000000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           ctx.BlockTime(),
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}

	expectedHist2 := types.MinterState{
		Position:                    2,
		AmountMinted:                sdk.NewInt(100000000000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           ctx.BlockTime(),
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}

	require.EqualValues(t, expectedHist1, history[0])
	require.EqualValues(t, expectedHist2, history[1])

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
