package cfevesting_test

import (
	"fmt"
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
)

func TestGenesisWholeApp(t *testing.T) {
	genesisState := types.GenesisState{
		Params:              types.NewParams("uc4e"),
		VestingAccountList:  []types.VestingAccount{},
		VestingAccountCount: 0,
		// this line is used by starport scaffolding # genesis/test/state
		VestingTypes: []types.GenesisVestingType{},
	}

	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eVestingUtils.InitGenesis(genesisState)
	testHelper.C4eVestingUtils.ExportGenesis(genesisState)
}

func TestGenesisVestingTypesAndAccounts(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	vestingTypesArray := testutils.GenerateGenesisVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params: types.NewParams("uc4e"),
		VestingAccountList: []types.VestingAccount{
			{
				Id:      0,
				Address: acountsAddresses[0].String(),
			},
			{
				Id:      1,
				Address: acountsAddresses[1].String(),
			},
		},
		VestingAccountCount: 2,
		VestingTypes:        vestingTypesArray,
	}

	testHelper := testapp.SetupTestApp(t)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)
	testHelper.C4eVestingUtils.ExportGenesis(genesisState)
}

func TestGenesisVestingTypes(t *testing.T) {
	vestingTypesArray := testutils.GenerateGenesisVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params:              types.NewParams("uc4e"),
		VestingAccountList:  []types.VestingAccount{},
		VestingAccountCount: 0,
		VestingTypes:        vestingTypesArray,
	}

	testHelper := testapp.SetupTestApp(t)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)
	testHelper.C4eVestingUtils.ExportGenesis(genesisState)
}

func TestGenesisVestingTypesUnitsSecondsToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*60*24, types.Second, types.Day)
}

func TestGenesisVestingTypesUnitsSecondsToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*60, types.Second, types.Hour)
}

func TestGenesisVestingTypesUnitsSecondsToMinutes(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60, types.Second, types.Minute)
}

func TestGenesisVestingTypesUnitsSecondsToSeconds(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, types.Second, types.Second)
}

func TestGenesisVestingTypesUnitsMinutesToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*24, types.Minute, types.Day)
}

func TestGenesisVestingTypesUnitsMinutesToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60, types.Minute, types.Hour)
}

func TestGenesisVestingTypesUnitsMinutesToMinutes(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, types.Minute, types.Minute)
}

func TestGenesisVestingTypesUnitsHoursToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 24, types.Hour, types.Day)
}

func TestGenesisVestingTypesUnitsHoursToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, types.Hour, types.Hour)
}

func TestGenesisVestingTypesUnitsDaysToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, types.Day, types.Day)
}

func genesisVestingTypesUnitsTest(t *testing.T, multiplier int64, srcUnits string, dstUnits string) {
	vestingTypesArray := testutils.GenerateGenesisVestingTypes(1, 1)
	vestingTypesArray[0].LockupPeriod = 234 * multiplier
	vestingTypesArray[0].LockupPeriodUnit = srcUnits

	vestingTypesArray[0].VestingPeriod = 345 * multiplier
	vestingTypesArray[0].VestingPeriodUnit = srcUnits
	vestingTypesArray[0].Free = sdk.ZeroDec()
	genesisState := types.GenesisState{
		Params:              types.NewParams("uc4e"),
		VestingAccountList:  []types.VestingAccount{},
		VestingAccountCount: 0,
		VestingTypes:        vestingTypesArray,
	}

	testHelper := testapp.SetupTestApp(t)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	vestingTypesArray[0].LockupPeriod = 234
	vestingTypesArray[0].LockupPeriodUnit = dstUnits

	vestingTypesArray[0].VestingPeriod = 345
	vestingTypesArray[0].VestingPeriodUnit = dstUnits

	testHelper.C4eVestingUtils.ExportGenesis(genesisState)
}

func getVestingPoolsAmount(accVestingPools []*types.AccountVestingPools) sdk.Int {
	result := sdk.ZeroInt()
	for _, accV := range accVestingPools {
		for _, v := range accV.VestingPools {
			result = result.Add(v.GetCurrentlyLocked())
		}
	}
	return result
}

func TestGenesisAccountVestingPools(t *testing.T) {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)

	genesisState := types.GenesisState{
		Params: types.NewParams(commontestutils.DefaultTestDenom),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := testapp.SetupTestApp(t)

	mintUndelegableCoinsToModule(testHelper, genesisState, getVestingPoolsAmount(accountVestingPoolsArray))
	testHelper.C4eVestingUtils.InitGenesis(genesisState)
	testHelper.C4eVestingUtils.ExportGenesis(genesisState)
}

func TestGenesisAccountVestingPoolsWrongAmountInModuleAccount(t *testing.T) {
	accountVestingPoolsArray := testutils.GenerateAccountVestingPoolsWithRandomVestingPools(10, 10, 1, 1)

	genesisState := types.GenesisState{
		Params: types.NewParams("uc4e"),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := testapp.SetupTestApp(t)

	VestingPoolsAmount := getVestingPoolsAmount(accountVestingPoolsArray)
	wrongAcountAmount := getVestingPoolsAmount(accountVestingPoolsArray).SubRaw(10)
	mintUndelegableCoinsToModule(testHelper, genesisState, wrongAcountAmount)
	testHelper.C4eVestingUtils.InitGenesisError(genesisState, fmt.Sprintf("module: cfevesting account balance of denom uc4e not equal of sum of vesting pools: %s <> %s", wrongAcountAmount.String(), VestingPoolsAmount.String()))

}

func mintUndelegableCoinsToModule(testHelper *testapp.TestHelper, genesisState types.GenesisState, amount sdk.Int) {
	mintedCoin := sdk.NewCoin(genesisState.Params.Denom, amount)
	testHelper.BankUtils.AddCoinsToModule(mintedCoin, types.ModuleName)
}
