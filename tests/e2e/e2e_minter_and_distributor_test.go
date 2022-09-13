package e2e

import (
	"encoding/json"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	typesDistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	typesMinter "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var denom = "uc4e"

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
			got := runDistributionAndMinting(tt.timeInYears)
			if got.developmentFundCoinsInt.Int64() != tt.want.developmentFundCoinsInt.Int64() {
				t.Errorf("CheckDevelopmentFunds() = %v, want %v", got.developmentFundCoinsInt, tt.want.developmentFundCoinsInt)
			} else if got.governanceBoosterCoinInt.Int64() != tt.want.governanceBoosterCoinInt.Int64() {
				t.Errorf("CheckGovernanceBoosterFunds() = %v, want %v", got.governanceBoosterCoinInt, tt.want.governanceBoosterCoinInt)
			} else if got.lpProviders.Int64() != tt.want.lpProviders.Int64() {
				t.Errorf("CheckLpProvidersFunds() = %v, want %v", got.lpProviders, tt.want.lpProviders)
			} else if got.greenEnergyBoosterCoinInt.Int64() != tt.want.greenEnergyBoosterCoinInt.Int64() {
				t.Errorf("CheckGreenEnergyBoosterFunds() = %v, want %v", got.greenEnergyBoosterCoinInt, tt.want.greenEnergyBoosterCoinInt)
			} else if got.totalSupply.Int64() != tt.want.totalSupply.Int64() {
				t.Errorf("CheckTotoalSupply() = %v, want %v", got.totalSupply, tt.want.totalSupply)
			}
		})
	}
}

func runDistributionAndMinting(timeInYear int) testResult {
	//read distributor params from json
	var distributorParams typesDistributor.Params
	jsonFile, _ := os.Open("test-resources/distributors.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &distributorParams)

	//read minter params from json
	var minterParams typesMinter.Params
	jsonFileMinter, _ := os.Open("test-resources/periodic_reduction_minter.json")
	byteValueMinter, _ := ioutil.ReadAll(jsonFileMinter)
	json.Unmarshal(byteValueMinter, &minterParams)

	startTime := time.Now()
	minterParams.Minter.Start = startTime

	app, ctx := commontestutils.SetupAppWithTime(1, startTime)
	app.CfeminterKeeper.SetParams(ctx, minterParams)
	app.CfedistributorKeeper.SetParams(ctx, distributorParams)

	oneYearDuration := 365 * 24 * time.Hour

	for i := 1; i <= timeInYear; i++ {
		ctx = ctx.WithBlockHeight(int64(i)).WithBlockTime(ctx.BlockTime().Add(oneYearDuration))
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
		app.EndBlocker(ctx, abci.RequestEndBlock{})
	}

	acc, _ := sdk.AccAddressFromBech32("cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")
	developmentFundCoinInt := app.CfedistributorKeeper.GetAccountCoins(ctx, acc).AmountOf(denom)

	governanceBoosterCoinInt := app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "governance_booster_collector").AmountOf(denom)
	greenEnergyBoosterCoinInt := app.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, "green_energy_booster_collector").AmountOf(denom)
	lpAccount, _ := sdk.AccAddressFromBech32("cosmos10mm944ph9jrgdjm5agwrk0csglev44nc4jsqtd")

	lpProviders := app.CfedistributorKeeper.GetAccountCoins(ctx, lpAccount).AmountOf(denom)
	totalSupply := app.BankKeeper.GetSupply(ctx, denom).Amount
	testResult := testResult{developmentFundCoinsInt: developmentFundCoinInt, governanceBoosterCoinInt: governanceBoosterCoinInt,
		greenEnergyBoosterCoinInt: greenEnergyBoosterCoinInt, lpProviders: lpProviders, totalSupply: totalSupply}

	return testResult
}
