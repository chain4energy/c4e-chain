package e2e

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcommon "github.com/chain4energy/c4e-chain/testutil/common"
	distributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	mintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const developmentFundAddrParamName = "development_fund_address"
const lpAccountAddrParamName = "lp_address"

var accountsAddresses, _ = testcommon.CreateAccounts(2, 0)
var developmentFundAddr = accountsAddresses[0]
var lpAccountAddr = accountsAddresses[1]

var distributorsJsonParams = map[string]string{
	developmentFundAddrParamName: developmentFundAddr.String(),
	lpAccountAddrParamName:       lpAccountAddr.String(),
}

type testResult struct {
	developmentFundCoinsInt   sdk.Int
	governanceBoosterCoinInt  sdk.Int
	greenEnergyBoosterCoinInt sdk.Int
	lpProviders               sdk.Int
	totalSupply               sdk.Int
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
	//read distributor params from json
	var distributorParams distributortypes.Params
	testcommon.UnmarshalJsonFileWithParams("test-resources/distributors.json", &distributorParams, distributorsJsonParams)

	//read minter params from json
	var minterParams mintertypes.Params
	testcommon.UnmarshalJsonFile("test-resources/exponential_step_minting.json", &minterParams)

	startTime := time.Now()
	minterParams.StartTime = startTime

	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1, startTime)
	testHelper.C4eMinterUtils.SetParams(minterParams)
	testHelper.C4eDistributorUtils.SetParams(distributorParams)

	oneYearDuration := 365 * 24 * time.Hour

	for i := 1; i <= timeInYear; i++ {
		testHelper.SetContextBlockHeightAndAddTime(int64(i), oneYearDuration)
		testHelper.BeginBlocker(abci.RequestBeginBlock{})
		testHelper.EndBlocker(abci.RequestEndBlock{})
	}

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(developmentFundAddr, expectedResults.developmentFundCoinsInt)
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(distributortypes.GovernanceBoosterCollector, expectedResults.governanceBoosterCoinInt)
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(distributortypes.GreenEnergyBoosterCollector, expectedResults.greenEnergyBoosterCoinInt)

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(lpAccountAddr, expectedResults.lpProviders)
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(expectedResults.totalSupply.Add(testHelper.InitialValidatorsCoin.Amount))
}
