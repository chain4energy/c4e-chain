package keeper_test

import (
	"cosmossdk.io/math"
	"testing"

	"github.com/chain4energy/c4e-chain/v2/testutil/app"

	testcosmos "github.com/chain4energy/c4e-chain/v2/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"

	testutils "github.com/chain4energy/c4e-chain/v2/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	// abci "github.com/cometbft/cometbft/abci/types"
)

func TestVestingsAmountPoolsOnlyNoGenesis(t *testing.T) {
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := math.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
	}
	accVestingPools := types.AccountVestingPools{
		Owner:        acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(testenv.DefaultTestDenom),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := app.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1000000), types.ModuleName)
	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	expected := types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.ZeroInt(),
		VestingInPoolsAmount:    math.ZeroInt(),
		VestingInAccountsAmount: math.ZeroInt(),
		DelegatedVestingAmount:  math.ZeroInt(),
	}

	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)
}

func TestVestingsAmountPoolsOnlyGenesis(t *testing.T) {
	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := math.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}
	accVestingPools := types.AccountVestingPools{
		Owner:        acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(testenv.DefaultTestDenom),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := app.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1000000), types.ModuleName)
	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	expected := types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1000000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.ZeroInt(),
		DelegatedVestingAmount:  math.ZeroInt(),
	}

	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)
}

func TestVestingsAmountPoolsAndAccountNoGenesis(t *testing.T) {
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := math.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
	}

	accVestingPools := types.AccountVestingPools{
		Owner:        acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(testenv.DefaultTestDenom),
		VestingAccountTraces: []types.VestingAccountTrace{
			{
				Id:      0,
				Address: acountsAddresses[1].String(),
			},
		},
		VestingAccountTraceCount: 1,
		VestingTypes:             []types.GenesisVestingType{},
		AccountVestingPools:      accountVestingPoolsArray,
	}

	testHelper := app.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), math.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	expected := types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.ZeroInt(),
		VestingInPoolsAmount:    math.ZeroInt(),
		VestingInAccountsAmount: math.ZeroInt(),
		DelegatedVestingAmount:  math.ZeroInt(),
	}
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	// expected = types.QueryGenesisVestingsSummaryResponse{
	// 	VestingAllAmount:        math.NewInt(1150000),
	// 	VestingInPoolsAmount:    math.NewInt(1000000),
	// 	VestingInAccountsAmount: math.NewInt(150000),
	// 	DelegatedVestingAmount:  math.NewInt(1).SubRaw(1),
	// }
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	// expected = types.QueryGenesisVestingsSummaryResponse{
	// 	VestingAllAmount:        math.NewInt(1000000),
	// 	VestingInPoolsAmount:    math.NewInt(1000000),
	// 	VestingInAccountsAmount: math.NewInt(0),
	// 	DelegatedVestingAmount:  math.NewInt(0),
	// }
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

}

func TestVestingsAmountPoolsAndAccountGenesisAccount(t *testing.T) {
	testVestingsAmountPoolsAndAccountGenesis(t, true, false, false)
}

func TestVestingsAmountPoolsAndAccountAccountFromGenesisPool(t *testing.T) {
	testVestingsAmountPoolsAndAccountGenesis(t, false, true, false)
}

func TestVestingsAmountPoolsAndAccountAccountFromGenesisAccount(t *testing.T) {
	testVestingsAmountPoolsAndAccountGenesis(t, false, false, true)
}

func testVestingsAmountPoolsAndAccountGenesis(t *testing.T, genesis bool, fromGenesisPool bool, fromGenesisAccount bool) {
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := math.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}

	accVestingPools := types.AccountVestingPools{
		Owner:        acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(testenv.DefaultTestDenom),
		VestingAccountTraces: []types.VestingAccountTrace{
			{
				Id:                 0,
				Address:            acountsAddresses[1].String(),
				Genesis:            genesis,
				FromGenesisPool:    fromGenesisPool,
				FromGenesisAccount: fromGenesisAccount,
			},
		},
		VestingAccountTraceCount: 1,
		VestingTypes:             []types.GenesisVestingType{},
		AccountVestingPools:      accountVestingPoolsArray,
	}

	testHelper := app.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), math.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	expected := types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1300000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(300000),
		DelegatedVestingAmount:  math.NewInt(1).SubRaw(1),
	}
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	expected = types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1150000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(150000),
		DelegatedVestingAmount:  math.NewInt(1).SubRaw(1),
	}
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	expected = types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1000000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(0),
		DelegatedVestingAmount:  math.NewInt(0),
	}
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

}

func TestVestingsAmountPoolsAndAccountMixed(t *testing.T) {
	acountsAddresses, _ := testcosmos.CreateAccounts(6, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := math.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
	}

	vestingPoolNoGenesis := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     false,
	}

	accVestingPools := types.AccountVestingPools{
		Owner:        acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool, &vestingPool, &vestingPool, &vestingPoolNoGenesis, &vestingPoolNoGenesis},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(testenv.DefaultTestDenom),
		VestingAccountTraces: []types.VestingAccountTrace{
			{
				Id:                 0,
				Address:            acountsAddresses[1].String(),
				Genesis:            true,
				FromGenesisPool:    false,
				FromGenesisAccount: false,
			},
			{
				Id:                 2,
				Address:            acountsAddresses[2].String(),
				Genesis:            false,
				FromGenesisPool:    true,
				FromGenesisAccount: false,
			},
			{
				Id:                 3,
				Address:            acountsAddresses[3].String(),
				Genesis:            false,
				FromGenesisPool:    false,
				FromGenesisAccount: true,
			},
			{
				Id:                 4,
				Address:            acountsAddresses[4].String(),
				Genesis:            false,
				FromGenesisPool:    false,
				FromGenesisAccount: false,
			},
			{
				Id:                 5,
				Address:            acountsAddresses[5].String(),
				Genesis:            false,
				FromGenesisPool:    false,
				FromGenesisAccount: false,
			},
		},
		VestingAccountTraceCount: 6,
		VestingTypes:             []types.GenesisVestingType{},
		AccountVestingPools:      accountVestingPoolsArray,
	}

	testHelper := app.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(5*1000000), types.ModuleName)
	for i := 1; i < 6; i++ {
		testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[i].String(), math.NewInt(300000), start, lockEnd)
	}
	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	expected := types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(3 * 1300000),
		VestingInPoolsAmount:    math.NewInt(3 * 1000000),
		VestingInAccountsAmount: math.NewInt(3 * 300000),
		DelegatedVestingAmount:  math.NewInt(1).SubRaw(1),
	}
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	expected = types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(3 * 1150000),
		VestingInPoolsAmount:    math.NewInt(3 * 1000000),
		VestingInAccountsAmount: math.NewInt(3 * 150000),
		DelegatedVestingAmount:  math.NewInt(1).SubRaw(1),
	}
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	expected = types.QueryGenesisVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(3 * 1000000),
		VestingInPoolsAmount:    math.NewInt(3 * 1000000),
		VestingInAccountsAmount: math.NewInt(0),
		DelegatedVestingAmount:  math.NewInt(0),
	}
	testHelper.C4eVestingUtils.QueryGenesisVestings(expected)

}
