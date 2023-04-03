package app

import (
	"cosmossdk.io/math"
	testgenesis "github.com/chain4energy/c4e-chain/tests/app/genesis"
	"github.com/chain4energy/c4e-chain/testutil/app"
	distributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		{"One year test", 1, testResult{developmentFundCoinsInt: sdk.NewInt(2000000000000),
			governanceBoosterCoinInt: sdk.NewInt(4620000000000), greenEnergyBoosterCoinInt: sdk.NewInt(4760000000000),
			lpProviders: sdk.NewInt(4620000000000), totalSupply: sdk.NewInt(40000000000000)}},

		{"Two years test", 2, testResult{developmentFundCoinsInt: sdk.NewInt(2 * 2000000000000),
			governanceBoosterCoinInt: sdk.NewInt(2 * 4620000000000), greenEnergyBoosterCoinInt: sdk.NewInt(2 * 4760000000000),
			lpProviders: sdk.NewInt(2 * 4620000000000), totalSupply: sdk.NewInt(2 * 40000000000000)}},

		{"Four years test", 4, testResult{developmentFundCoinsInt: sdk.NewInt(4 * 2000000000000),
			governanceBoosterCoinInt: sdk.NewInt(4 * 4620000000000), greenEnergyBoosterCoinInt: sdk.NewInt(4 * 4760000000000),
			lpProviders: sdk.NewInt(4 * 4620000000000), totalSupply: sdk.NewInt(4 * 40000000000000)}},

		{"8 years test", 8, testResult{developmentFundCoinsInt: sdk.NewInt(6 * 2000000000000),
			governanceBoosterCoinInt: sdk.NewInt(6 * 4620000000000), greenEnergyBoosterCoinInt: sdk.NewInt(6 * 4760000000000),
			lpProviders: sdk.NewInt(6 * 4620000000000), totalSupply: sdk.NewInt(6 * 40000000000000)}},

		{"16 years test", 16, testResult{developmentFundCoinsInt: sdk.NewInt(7.5 * 2000000000000),
			governanceBoosterCoinInt: sdk.NewInt(7.5 * 4620000000000), greenEnergyBoosterCoinInt: sdk.NewInt(7.5 * 4760000000000),
			lpProviders: sdk.NewInt(7.5 * 4620000000000), totalSupply: sdk.NewInt(7.5 * 40000000000000)}},
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
