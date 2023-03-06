package cfeminter_test

import (
	"github.com/chain4energy/c4e-chain/testutil/app"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
)

func TestOneYearLinear(t *testing.T) {
	totalSupply := sdk.NewInt(40000000000000)

	testHelper := app.SetupTestApp(t)

	yearFromNow := testHelper.InitTime.Add(time.Hour * 24 * 365)
	config, _ := codectypes.NewAnyWithValue(&types.LinearMinting{Amount: totalSupply})

	minters := []*types.Minter{
		{SequenceId: 1, EndTime: &yearFromNow, Config: config},
		{SequenceId: 2},
	}

	genesisState := types.GenesisState{
		Params:      types.NewParams(testenv.DefaultTestDenom, testHelper.InitTime, minters),
		MinterState: types.MinterState{SequenceId: 1, AmountMinted: sdk.ZeroInt(), LastMintBlockTime: testHelper.InitTime},
	}

	testHelper.C4eMinterUtils.InitGenesis(genesisState)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(consToAdd, acountsAddresses[0])

	testHelper.C4eMinterUtils.VerifyInflation(sdk.NewDec(1))
	testHelper.C4eMinterUtils.VerifyMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.InitTime, sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, sdk.ZeroInt())

	numOfHours := 365 * 24
	for i := 1; i <= numOfHours; i++ {
		testHelper.SetContextBlockHeightAndAddTime(int64(i), time.Hour)
		testHelper.BeginBlocker(abci.RequestBeginBlock{})
		testHelper.EndBlocker(abci.RequestEndBlock{})

		expectedToMint := sdk.NewDecFromInt(totalSupply.MulRaw(int64(i))).QuoInt64(int64(numOfHours))
		expectedMinted := expectedToMint.TruncateInt()
		remainder := expectedToMint.Sub(expectedToMint.TruncateDec())
		expectedInflation := sdk.NewDecFromInt(totalSupply).QuoInt(totalSupply.Add(expectedMinted))

		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, expectedMinted)

		if i < numOfHours {
			testHelper.C4eMinterUtils.VerifyInflation(expectedInflation)
			testHelper.C4eMinterUtils.VerifyMinterState(1, expectedMinted, remainder, testHelper.Context.BlockTime(), sdk.ZeroDec())
		} else {
			testHelper.C4eMinterUtils.VerifyInflation(sdk.ZeroDec())
			testHelper.C4eMinterUtils.VerifyMinterState(2, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.Context.BlockTime(), sdk.ZeroDec())
		}

		testHelper.C4eMinterUtils.ExportGenesisAndValidate()
	}

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, totalSupply)
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(totalSupply.MulRaw(2))
}

func TestFewYearsExponentialStepMinting(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(160000000000000)

	testHelper := app.SetupTestApp(t)
	minter := types.ExponentialStepMinting{Amount: startAmountYearly, StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	config, _ := codectypes.NewAnyWithValue(&minter)
	minters := []*types.Minter{
		{SequenceId: 1, Config: config},
	}

	genesisState := types.GenesisState{
		Params:      types.NewParams(testenv.DefaultTestDenom, testHelper.InitTime, minters),
		MinterState: types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(0), LastMintBlockTime: testHelper.InitTime},
	}

	testHelper.C4eMinterUtils.InitGenesis(genesisState)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(consToAdd, acountsAddresses[0])

	testHelper.C4eMinterUtils.VerifyInflation(sdk.MustNewDecFromStr("0.1"))
	testHelper.C4eMinterUtils.VerifyMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.InitTime, sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 4 * year
	amountYearly := startAmountYearly.QuoRaw(4)
	prevPeriodMinted := sdk.ZeroInt()

	for MintersCount := 1; MintersCount <= 5; MintersCount++ {
		for i := 1; i <= numOfHours; i++ {
			testHelper.SetContextBlockHeightAndAddTime(int64(i), 24*time.Hour)
			testHelper.BeginBlocker(abci.RequestBeginBlock{})
			testHelper.EndBlocker(abci.RequestEndBlock{})

			expectedToMint := sdk.NewDecFromInt(amountYearly.MulRaw(int64(i))).QuoInt64(int64(year))
			expectedMinted := expectedToMint.TruncateInt()
			remainder := expectedToMint.Sub(expectedToMint.TruncateDec())
			expectedInflation := sdk.NewDecFromInt(amountYearly).QuoInt(totalSupply.Add(prevPeriodMinted).Add(expectedMinted))

			testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, prevPeriodMinted.Add(expectedMinted))

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
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, expectedMinted)
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(totalSupply.Add(expectedMinted))
}

func TestFewYearsExponentialStepMintingInOneBlock(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(160000000000000)
	testHelper := app.SetupTestApp(t)

	minter := types.ExponentialStepMinting{Amount: startAmountYearly, StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	config, _ := codectypes.NewAnyWithValue(&minter)

	minters := []*types.Minter{{SequenceId: 1, Config: config}}

	genesisState := types.GenesisState{
		Params:      types.NewParams(testenv.DefaultTestDenom, testHelper.InitTime, minters),
		MinterState: types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(0), LastMintBlockTime: testHelper.InitTime},
	}

	testHelper.C4eMinterUtils.InitGenesis(genesisState)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(consToAdd, acountsAddresses[0])

	testHelper.C4eMinterUtils.VerifyInflation(sdk.MustNewDecFromStr("0.1"))
	testHelper.C4eMinterUtils.VerifyMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.InitTime, sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, sdk.ZeroInt())

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

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, expectedMinted)

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(totalSupply.Add(expectedMinted))

	testHelper.C4eMinterUtils.ExportGenesisAndValidate()
}

func TestFewYearsLinearMintingAndExponentialStepMintingInOneBlock(t *testing.T) {
	totalSupply := sdk.NewInt(400000000000000)
	startAmountYearly := sdk.NewInt(160000000000000)

	testHelper := app.SetupTestApp(t)

	tenYears := Year * 10
	endTime1 := testHelper.InitTime.Add(tenYears)
	endTime2 := endTime1.Add(tenYears)

	linearMinting1 := types.LinearMinting{Amount: sdk.NewInt(200000000000000)}
	linearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000000000000)}
	exponentialStepMinting := types.ExponentialStepMinting{Amount: startAmountYearly, StepDuration: NanoSecondsInFourYears, AmountMultiplier: sdk.MustNewDecFromStr("0.5")}
	config3, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)
	config1, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config1}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3, Config: config3}

	minters := []*types.Minter{
		&minter1,
		&minter2,
		&minter3,
	}

	genesisState := types.GenesisState{
		Params:      types.NewParams(testenv.DefaultTestDenom, testHelper.InitTime, minters),
		MinterState: types.MinterState{SequenceId: 1, AmountMinted: sdk.NewInt(0), LastMintBlockTime: testHelper.InitTime},
	}

	testHelper.C4eMinterUtils.InitGenesis(genesisState)

	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
	consToAdd := totalSupply.Sub(testHelper.InitialValidatorsCoin.Amount)

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(consToAdd, acountsAddresses[0])

	testHelper.C4eMinterUtils.VerifyInflation(sdk.MustNewDecFromStr("0.05"))
	testHelper.C4eMinterUtils.VerifyMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), testHelper.InitTime, sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, sdk.ZeroInt())

	year := 365 //* 24
	numOfHours := 321 * year

	for i := 1; i <= numOfHours; i++ {
		testHelper.SetContextBlockHeightAndAddTime(int64(i), 24*time.Hour)
	}

	testHelper.BeginBlocker(abci.RequestBeginBlock{})
	testHelper.EndBlocker(abci.RequestEndBlock{})

	expectedMintedSequenceId3 := sdk.NewInt(320000000000000)
	expectedRemainderSequenceId3 := sdk.MustNewDecFromStr("0.000000004235164732")

	expectedMinted := sdk.NewInt(620000000000000)

	testHelper.C4eMinterUtils.VerifyMinterState(3, expectedMintedSequenceId3, expectedRemainderSequenceId3, testHelper.Context.BlockTime(), sdk.ZeroDec())

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(testenv.DefaultDistributionDestination, expectedMinted)

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(totalSupply.Add(expectedMinted))

	expectedHist1 := types.MinterState{
		SequenceId:                  1,
		AmountMinted:                sdk.NewInt(200000000000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           testHelper.Context.BlockTime(),
		RemainderFromPreviousMinter: sdk.ZeroDec(),
	}

	expectedHist2 := types.MinterState{
		SequenceId:                  2,
		AmountMinted:                sdk.NewInt(100000000000000),
		RemainderToMint:             sdk.ZeroDec(),
		LastMintBlockTime:           testHelper.Context.BlockTime(),
		RemainderFromPreviousMinter: sdk.ZeroDec(),
	}

	testHelper.C4eMinterUtils.VerifyMinterHistory(expectedHist1, expectedHist2)

	testHelper.C4eMinterUtils.ExportGenesisAndValidate()
}
