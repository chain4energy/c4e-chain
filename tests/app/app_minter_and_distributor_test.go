package app

import (
	"cosmossdk.io/math"
	testgenesis "github.com/chain4energy/c4e-chain/tests/app/genesis"
	"github.com/chain4energy/c4e-chain/testutil/app"
	distributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
	"time"
)

const oneYearDuration = time.Hour * 24 * 365

type testResult struct {
	developmentFundCoinsInt   math.Int
	governanceBoosterCoinInt  math.Int
	greenEnergyBoosterCoinInt math.Int
	lpProviders               math.Int
	totalSupply               math.Int
}

func TestMinterWithDistributor(t *testing.T) {
	tests := []struct {
		name        string
		timeInYears int
		want        testResult
	}{
		{"One year test", 1, testResult{developmentFundCoinsInt: math.NewInt(2000000000000),
			governanceBoosterCoinInt: math.NewInt(4620000000000), greenEnergyBoosterCoinInt: math.NewInt(4760000000000),
			lpProviders: math.NewInt(4620000000000), totalSupply: math.NewInt(40000000000000)}},

		{"Two years test", 2, testResult{developmentFundCoinsInt: math.NewInt(2 * 2000000000000),
			governanceBoosterCoinInt: math.NewInt(2 * 4620000000000), greenEnergyBoosterCoinInt: math.NewInt(2 * 4760000000000),
			lpProviders: math.NewInt(2 * 4620000000000), totalSupply: math.NewInt(2 * 40000000000000)}},

		{"Four years test", 4, testResult{developmentFundCoinsInt: math.NewInt(4 * 2000000000000),
			governanceBoosterCoinInt: math.NewInt(4 * 4620000000000), greenEnergyBoosterCoinInt: math.NewInt(4 * 4760000000000),
			lpProviders: math.NewInt(4 * 4620000000000), totalSupply: math.NewInt(4 * 40000000000000)}},

		{"8 years test", 8, testResult{developmentFundCoinsInt: math.NewInt(6 * 2000000000000),
			governanceBoosterCoinInt: math.NewInt(6 * 4620000000000), greenEnergyBoosterCoinInt: math.NewInt(6 * 4760000000000),
			lpProviders: math.NewInt(6 * 4620000000000), totalSupply: math.NewInt(6 * 40000000000000)}},

		{"16 years test", 16, testResult{developmentFundCoinsInt: math.NewInt(7.5 * 2000000000000),
			governanceBoosterCoinInt: math.NewInt(7.5 * 4620000000000), greenEnergyBoosterCoinInt: math.NewInt(7.5 * 4760000000000),
			lpProviders: math.NewInt(7.5 * 4620000000000), totalSupply: math.NewInt(7.5 * 40000000000000)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runDistributionAndMinting(t, tt.timeInYears, tt.want)
		})
	}
}

func runDistributionAndMinting(t *testing.T, timeInYear int, expectedResults testResult) {
	var distributorParams = testgenesis.CfeDistributorParams
	var minterParams = testgenesis.CfeMinterrParams()

	testHelper := app.SetupTestAppWithHeightAndTime(t, 1, minterParams.StartTime)
	testHelper.C4eMinterUtils.SetParams(minterParams)
	testHelper.C4eDistributorUtils.SetParams(distributorParams)

	for i := 1; i <= timeInYear; i++ {
		testHelper.SetContextBlockHeightAndAddTime(int64(i), oneYearDuration)
		testHelper.BeginBlocker(abci.RequestBeginBlock{})
		testHelper.EndBlocker(abci.RequestEndBlock{})
	}

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(testgenesis.DevelopmentFundAddr, expectedResults.developmentFundCoinsInt)
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(distributortypes.GovernanceBoosterCollector, expectedResults.governanceBoosterCoinInt)
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(distributortypes.GreenEnergyBoosterCollector, expectedResults.greenEnergyBoosterCoinInt)

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(testgenesis.LpAccountAddr, expectedResults.lpProviders)
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(expectedResults.totalSupply.Add(testHelper.InitialValidatorsCoin.Amount))
}
