package cfevesting_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

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

	cfevesting.InitGenesis(testHelper.Context, testHelper.App.CfevestingKeeper, genesisState, testHelper.App.AccountKeeper, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)
	got := cfevesting.ExportGenesis(testHelper.Context, testHelper.App.CfevestingKeeper)
	require.NotNil(t, got)

	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VestingAccountList, got.VestingAccountList)
	require.Equal(t, genesisState.VestingAccountCount, got.VestingAccountCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

func TestGenesisVestingTypesAndAccounts(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	vestingTypesArray := generateGenesisVestingTypes(10, 1)
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

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)
	got := cfevesting.ExportGenesis(testHelper.Context, k)

	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}

func TestGenesisVestingTypes(t *testing.T) {
	vestingTypesArray := generateGenesisVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params:              types.NewParams("uc4e"),
		VestingAccountList:  []types.VestingAccount{},
		VestingAccountCount: 0,
		VestingTypes:        vestingTypesArray,
	}

	testHelper := testapp.SetupTestApp(t)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)
	got := cfevesting.ExportGenesis(testHelper.Context, k)

	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}

func TestGenesisValidationVestingTypes(t *testing.T) {
	vestingTypesArray := generateGenesisVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params:       types.NewParams("test_denom"),
		VestingTypes: vestingTypesArray,
	}

	err := genesisState.Validate()
	require.Nil(t, err)
}

func TestGenesisValidationVestingAccounts(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	vestingTypesArray := generateGenesisVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),
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

	err := genesisState.Validate()
	require.Nil(t, err)
}

func TestGenesisValidationVestingTypesNameMoreThanOnceError(t *testing.T) {
	vestingTypesArray := generateGenesisVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params:       types.NewParams("test_denom"),
		VestingTypes: vestingTypesArray,
	}

	vestingTypesArray[3].Name = vestingTypesArray[6].Name

	err := genesisState.Validate()
	require.EqualError(t, err,
		"vesting type with name: test-vesting-type-7 defined more than once")
}

func TestGenesisValidationVestingAccountsError(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)
	vestingTypesArray := generateGenesisVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),
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
		VestingAccountCount: 1,
		VestingTypes:        vestingTypesArray,
	}

	err := genesisState.Validate()
	require.EqualError(t, err,
		"vestingAccount id should be lower or equal than the last id")
}

func TestGenesisValidationVestingAccountVestingPools(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)
	vestingTypes := generateGenesisVestingTypesForAccounVestings(accountVestingsListArray)

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        vestingTypes,
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	err := genesisState.Validate()
	require.Nil(t, err)

}

func TestGenesisValidationVestingAccountVestingPoolsNoVestingTypesError(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	err := genesisState.Validate()
	require.EqualError(t, err,
		"vesting with id: 1 defined for account: "+accountVestingsListArray[0].Address+" - vesting type not found: test-vesting-account-1-1")

}

func TestGenesisValidationVestingAccountVestingPoolsOneVestingTypeNotExistError(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)
	vestingTypes := generateGenesisVestingTypesForAccounVestings(accountVestingsListArray)
	accountVestingsListArray[4].VestingPools[7].VestingType = "wrong type"

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        vestingTypes,
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	err := genesisState.Validate()
	require.EqualError(t, err,
		"vesting with id: 8 defined for account: "+accountVestingsListArray[4].Address+" - vesting type not found: "+accountVestingsListArray[4].VestingPools[7].VestingType)

}

func TestGenesisValidationVestingAccountVestingPoolsMoreThanOneIdError(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)
	accountVestingsListArray[4].VestingPools[3].Id = accountVestingsListArray[4].VestingPools[6].Id
	vestingTypes := generateGenesisVestingTypesForAccounVestings(accountVestingsListArray)

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        vestingTypes,
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	err := genesisState.Validate()
	require.EqualError(t, err,
		"vesting with id: 7 defined more than once for account: "+accountVestingsListArray[4].Address)

}

func TestGenesisValidationVestingAccountVestingPoolsMoreThanOneNameError(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)
	accountVestingsListArray[4].VestingPools[3].Name = accountVestingsListArray[4].VestingPools[6].Name
	vestingTypes := generateGenesisVestingTypesForAccounVestings(accountVestingsListArray)

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        vestingTypes,
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	err := genesisState.Validate()
	require.EqualError(t, err,
		"vesting with name: "+accountVestingsListArray[4].VestingPools[3].Name+" defined more than once for account: "+accountVestingsListArray[4].Address)

}

func TestGenesisValidationVestingAccountVestingPoolsMoreThanOneAddressError(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)
	accountVestingsListArray[3].Address = accountVestingsListArray[7].Address
	vestingTypes := generateGenesisVestingTypesForAccounVestings(accountVestingsListArray)

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        vestingTypes,
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	err := genesisState.Validate()
	require.EqualError(t, err,
		"account vestings with address: "+accountVestingsListArray[3].Address+" defined more than once")

}

func TestGenesisVestingTypesUnitsSecondsToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*60*24, keeper.Second, keeper.Day)
}

func TestGenesisVestingTypesUnitsSecondsToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*60, keeper.Second, keeper.Hour)
}

func TestGenesisVestingTypesUnitsSecondsToMinutes(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60, keeper.Second, keeper.Minute)
}

func TestGenesisVestingTypesUnitsSecondsToSeconds(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, keeper.Second, keeper.Second)
}

func TestGenesisVestingTypesUnitsMinutesToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60*24, keeper.Minute, keeper.Day)
}

func TestGenesisVestingTypesUnitsMinutesToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 60, keeper.Minute, keeper.Hour)
}

func TestGenesisVestingTypesUnitsMinutesToMinutes(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, keeper.Minute, keeper.Minute)
}

func TestGenesisVestingTypesUnitsHoursToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 24, keeper.Hour, keeper.Day)
}

func TestGenesisVestingTypesUnitsHoursToHours(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, keeper.Hour, keeper.Hour)
}

func TestGenesisVestingTypesUnitsDaysToDays(t *testing.T) {
	genesisVestingTypesUnitsTest(t, 1, keeper.Day, keeper.Day)
}

func genesisVestingTypesUnitsTest(t *testing.T, multiplier int64, srcUnits string, dstUnits string) {
	vestingTypesArray := generateGenesisVestingTypes(1, 1)
	vestingTypesArray[0].LockupPeriod = 234 * multiplier
	vestingTypesArray[0].LockupPeriodUnit = srcUnits

	vestingTypesArray[0].VestingPeriod = 345 * multiplier
	vestingTypesArray[0].VestingPeriodUnit = srcUnits

	genesisState := types.GenesisState{
		Params:              types.NewParams("uc4e"),
		VestingAccountList:  []types.VestingAccount{},
		VestingAccountCount: 0,
		VestingTypes:        vestingTypesArray,
	}

	testHelper := testapp.SetupTestApp(t)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)

	vestingTypesArray[0].LockupPeriod = 234
	vestingTypesArray[0].LockupPeriodUnit = dstUnits

	vestingTypesArray[0].VestingPeriod = 345
	vestingTypesArray[0].VestingPeriodUnit = dstUnits

	got := cfevesting.ExportGenesis(testHelper.Context, k)

	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}

func getUndelegableAmount(accvestings []*types.AccountVestings) sdk.Int {
	result := sdk.ZeroInt()
	for _, accV := range accvestings {
		for _, v := range accV.VestingPools {
			result = result.Add(v.LastModificationVested).Sub(v.LastModificationWithdrawn)
		}
	}
	return result
}

func TestGenesisAccountVestingsList(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)

	genesisState := types.GenesisState{
		Params: types.NewParams(commontestutils.DefaultTestDenom),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	testHelper := testapp.SetupTestApp(t)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	mintUndelegableCoinsToModule(testHelper, genesisState, getUndelegableAmount(accountVestingsListArray))
	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)
	got := cfevesting.ExportGenesis(testHelper.Context, k)
	require.NotNil(t, got)
	require.EqualValues(t, genesisState.Params, got.GetParams())
	require.EqualValues(t, genesisState.VestingTypes, (*got).VestingTypes)
	require.EqualValues(t, len(accountVestingsListArray), len((*got).AccountVestingsList.Vestings))

	testutils.AssertAccountVestingsArrays(t, accountVestingsListArray, (*got).AccountVestingsList.Vestings)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

}

func TestGenesisAccountVestingsListWrongAmountInModuleAccount(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)

	genesisState := types.GenesisState{
		Params: types.NewParams("uc4e"),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	testHelper := testapp.SetupTestApp(t)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	undelegableAmount := getUndelegableAmount(accountVestingsListArray)
	wrongAcountAmount := getUndelegableAmount(accountVestingsListArray).SubRaw(10)
	mintUndelegableCoinsToModule(testHelper, genesisState, wrongAcountAmount)

	require.PanicsWithError(t, fmt.Sprintf("module: cfevesting account balance of denom uc4e not equal of sum of undelegable vestings: %s <> %s", wrongAcountAmount.String(), undelegableAmount.String()),
		func() { cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper) }, "")

}

func mintUndelegableCoinsToModule(testHelper *testapp.TestHelper, genesisState types.GenesisState, amount sdk.Int) {
	mintedCoin := sdk.NewCoin(genesisState.Params.Denom, amount)
	testHelper.BankUtils.AddCoinsToModule(mintedCoin, types.ModuleName)
}

func TestDurationFromUnits(t *testing.T) {
	amount := int64(456)
	require.EqualValues(t, amount*int64(time.Second), keeper.DurationFromUnits(keeper.Second, amount))
	require.EqualValues(t, amount*int64(time.Minute), keeper.DurationFromUnits(keeper.Minute, amount))
	require.EqualValues(t, amount*int64(time.Hour), keeper.DurationFromUnits(keeper.Hour, amount))
	require.EqualValues(t, amount*int64(time.Hour*24), keeper.DurationFromUnits(keeper.Day, amount))

}

func TestDurationFromUnitsWrongUnit(t *testing.T) {
	require.PanicsWithError(t, "Unknown PeriodUnit: das: invalid type", func() { keeper.DurationFromUnits("das", 234) }, "")

}

func TestUnitsFromDuration(t *testing.T) {
	unit, amount := keeper.UnitsFromDuration(234 * time.Second)
	require.EqualValues(t, keeper.Second, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = keeper.UnitsFromDuration(234 * 60 * time.Second)
	require.EqualValues(t, keeper.Minute, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = keeper.UnitsFromDuration(234 * 60 * 60 * time.Second)
	require.EqualValues(t, keeper.Hour, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = keeper.UnitsFromDuration(234 * 60 * 60 * 24 * time.Second)
	require.EqualValues(t, keeper.Day, unit)
	require.EqualValues(t, 234, amount)
}

func generateGenesisVestingTypes(numberOfVestingTypes int, startId int) []types.GenesisVestingType {
	vts := testutils.GenerateVestingTypes(numberOfVestingTypes, startId)
	result := []types.GenesisVestingType{}
	for _, vt := range vts {

		gvt := types.GenesisVestingType{
			Name:              vt.Name,
			LockupPeriod:      vt.LockupPeriod.Nanoseconds() / int64(time.Hour),
			LockupPeriodUnit:  keeper.Day,
			VestingPeriod:     vt.VestingPeriod.Nanoseconds() / int64(time.Hour),
			VestingPeriodUnit: keeper.Day,
		}
		result = append(result, gvt)
	}
	return result
}

func generateGenesisVestingTypesForAccounVestings(vestings []*types.AccountVestings) []types.GenesisVestingType {
	vt := testutils.GenerateVestingTypes(1, 1)[0]
	m := make(map[string]types.GenesisVestingType)
	result := []types.GenesisVestingType{}
	for _, av := range vestings {
		for _, v := range av.VestingPools {
			gvt := types.GenesisVestingType{
				Name:              v.VestingType,
				LockupPeriod:      vt.LockupPeriod.Nanoseconds() / int64(time.Hour),
				LockupPeriodUnit:  keeper.Day,
				VestingPeriod:     vt.VestingPeriod.Nanoseconds() / int64(time.Hour),
				VestingPeriodUnit: keeper.Day,
			}
			m[v.VestingType] = gvt

		}
	}
	for _, gvt := range m {
		result = append(result, gvt)
	}

	return result
}
