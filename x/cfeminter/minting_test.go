package cfeminter_test

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	routingdistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
)

func TestOneYearLinear(t *testing.T) {
	totalSupply := sdk.NewInt(40000000000000)

	testHelper := testapp.SetupTestApp(t)

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
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.ZeroInt(), LastMintBlockTime: testHelper.InitTime},
	}

	testHelper.C4eMinterUtils.InitGenesis(genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(consToAdd, acountsAddresses[0])

	testHelper.C4eMinterUtils.VerifyInflation(sdk.NewDec(1))
	testHelper.C4eMinterUtils.VerifyMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.InitTime, sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, sdk.ZeroInt())

	numOfHours := 365 * 24
	for i := 1; i <= numOfHours; i++ {
		testHelper.SetContextBlockHeightAndAddTime(int64(i), time.Hour)
		testHelper.BeginBlocker(abci.RequestBeginBlock{})
		testHelper.EndBlocker(abci.RequestEndBlock{})

		expectedToMint := sdk.NewDecFromInt(totalSupply.MulRaw(int64(i))).QuoInt64(int64(numOfHours))
		expectedMinted := expectedToMint.TruncateInt()
		remainder := expectedToMint.Sub(expectedToMint.TruncateDec())
		expectedInflation := sdk.NewDecFromInt(totalSupply).QuoInt(totalSupply.Add(expectedMinted))

		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, expectedMinted)

		if i < numOfHours {
			testHelper.C4eMinterUtils.VerifyInflation(expectedInflation)
			testHelper.C4eMinterUtils.VerifyMinterState(1, expectedMinted, remainder, testHelper.Context.BlockTime(), sdk.ZeroDec())
		} else {
			testHelper.C4eMinterUtils.VerifyInflation(sdk.ZeroDec())
			testHelper.C4eMinterUtils.VerifyMinterState(2, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.Context.BlockTime(), sdk.ZeroDec())
		}

		testHelper.C4eMinterUtils.ExportGenesisAndValidate()
	}

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, totalSupply)
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(totalSupply.MulRaw(2))
}

func TestFewYearsPeriodicReduction(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(40000000000000)
	testHelper := testapp.SetupTestApp(t)
	pminter := types.PeriodicReductionMinter{MintAmount: startAmountYearly, MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	minter := types.Minter{
		Start: testHelper.InitTime,
		Periods: []*types.MintingPeriod{
			{Position: 1, Type: types.PERIODIC_REDUCTION_MINTER,
				PeriodicReductionMinter: &pminter},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(commontestutils.DefaultTestDenom, minter),
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0), LastMintBlockTime: testHelper.InitTime},
	}

	testHelper.C4eMinterUtils.InitGenesis(genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(consToAdd, acountsAddresses[0])

	testHelper.C4eMinterUtils.VerifyInflation(sdk.MustNewDecFromStr("0.1"))
	testHelper.C4eMinterUtils.VerifyMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.InitTime, sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 4 * year
	amountYearly := startAmountYearly
	prevPeriodMinted := sdk.ZeroInt()

	for periodsCount := 1; periodsCount <= 5; periodsCount++ {
		for i := 1; i <= numOfHours; i++ {
			testHelper.SetContextBlockHeightAndAddTime(int64(i), 24*time.Hour)
			testHelper.BeginBlocker(abci.RequestBeginBlock{})
			testHelper.EndBlocker(abci.RequestEndBlock{})

			expectedToMint := sdk.NewDecFromInt(amountYearly.MulRaw(int64(i))).QuoInt64(int64(year))
			expectedMinted := expectedToMint.TruncateInt()
			remainder := expectedToMint.Sub(expectedToMint.TruncateDec())
			expectedInflation := sdk.NewDecFromInt(amountYearly).QuoInt(totalSupply.Add(prevPeriodMinted).Add(expectedMinted))

			testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, prevPeriodMinted.Add(expectedMinted))

			if i < numOfHours {
				testHelper.C4eMinterUtils.VerifyInflation(expectedInflation)
				testHelper.C4eMinterUtils.VerifyMinterState(1, prevPeriodMinted.Add(expectedMinted), remainder, testHelper.Context.BlockTime(), sdk.ZeroDec())
			} else {
				testHelper.C4eMinterUtils.VerifyInflation(expectedInflation.QuoInt64(2))
				testHelper.C4eMinterUtils.VerifyMinterState(1, prevPeriodMinted.Add(expectedMinted), remainder, testHelper.Context.BlockTime(), sdk.ZeroDec())
				prevPeriodMinted = prevPeriodMinted.Add(expectedMinted)
			}

			testHelper.C4eMinterUtils.ExportGenesisAndValidate()
		}
		amountYearly = amountYearly.QuoRaw(2)
	}
	expectedMinted := sdk.NewInt(310000000000000)
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, expectedMinted)
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(totalSupply.Add(expectedMinted))
}

func TestFewYearsPeriodicReductionInOneBlock(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(40000000000000)
	testHelper := testapp.SetupTestApp(t)

	pminter := types.PeriodicReductionMinter{MintAmount: startAmountYearly, MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	minter := types.Minter{
		Start: testHelper.InitTime,
		Periods: []*types.MintingPeriod{
			{Position: 1, Type: types.PERIODIC_REDUCTION_MINTER,
				PeriodicReductionMinter: &pminter},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(commontestutils.DefaultTestDenom, minter),
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0), LastMintBlockTime: testHelper.InitTime},
	}

	testHelper.C4eMinterUtils.InitGenesis(genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(consToAdd, acountsAddresses[0])

	testHelper.C4eMinterUtils.VerifyInflation(sdk.MustNewDecFromStr("0.1"))
	testHelper.C4eMinterUtils.VerifyMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.InitTime, sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 301 * year

	for i := 1; i <= numOfHours; i++ {
		testHelper.SetContextBlockHeightAndAddTime(int64(i), 24*time.Hour)
	}

	testHelper.BeginBlocker(abci.RequestBeginBlock{})
	testHelper.EndBlocker(abci.RequestEndBlock{})
	expectedMinted := sdk.NewInt(320000000000000)
	expectedRemainder := sdk.MustNewDecFromStr("0.000000004235164732")
	testHelper.C4eMinterUtils.VerifyMinterState(1, expectedMinted, expectedRemainder, testHelper.Context.BlockTime(), sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, expectedMinted)

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(totalSupply.Add(expectedMinted))

	testHelper.C4eMinterUtils.ExportGenesisAndValidate()
}

func TestFewYearsLinearAndPeriodicReductionInOneBlock(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(40000000000000)

	testHelper := testapp.SetupTestApp(t)

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
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0), LastMintBlockTime: testHelper.InitTime},
	}

	testHelper.C4eMinterUtils.InitGenesis(genesisState)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(consToAdd, acountsAddresses[0])

	testHelper.C4eMinterUtils.VerifyInflation(sdk.MustNewDecFromStr("0.05"))
	testHelper.C4eMinterUtils.VerifyMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.InitTime, sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 321 * year

	for i := 1; i <= numOfHours; i++ {
		testHelper.SetContextBlockHeightAndAddTime(int64(i), 24*time.Hour)
	}

	testHelper.BeginBlocker(abci.RequestBeginBlock{})
	testHelper.EndBlocker(abci.RequestEndBlock{})

	expectedMintedPosition3 := sdk.NewInt(320000000000000)
	expectedRemainderPosition3 := sdk.MustNewDecFromStr("0.000000004235164732")

	expectedMinted := sdk.NewInt(620000000000000)

	testHelper.C4eMinterUtils.VerifyMinterState(3, expectedMintedPosition3, expectedRemainderPosition3, testHelper.Context.BlockTime(), sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(routingdistributortypes.DistributorMainAccount, expectedMinted)

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(totalSupply.Add(expectedMinted))

	expectedHist1 := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(200000000000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           testHelper.Context.BlockTime(),
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}

	expectedHist2 := types.MinterState{
		Position:                    2,
		AmountMinted:                sdk.NewInt(100000000000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           testHelper.Context.BlockTime(),
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
	}

	testHelper.C4eMinterUtils.VerifyMinterHistory(expectedHist1, expectedHist2)

	testHelper.C4eMinterUtils.ExportGenesisAndValidate()
}
