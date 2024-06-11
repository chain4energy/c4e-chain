package keeper_test

import (
	"cosmossdk.io/math"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/app"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

	abci "github.com/cometbft/cometbft/abci/types"
)

func TestVestingsAmountPoolsOnly(t *testing.T) {
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

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1000000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.ZeroInt(),
		DelegatedVestingAmount:  math.ZeroInt(),
	}

	testHelper.C4eVestingUtils.QueryVestings(expected)
}

func TestVestingsAmountPoolsAndAccount(t *testing.T) {
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

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1300000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(300000),
		DelegatedVestingAmount:  math.NewInt(1).SubRaw(1),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1150000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(150000),
		DelegatedVestingAmount:  math.NewInt(1).SubRaw(1),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1000000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(0),
		DelegatedVestingAmount:  math.NewInt(0),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

}

func TestVestingsAmountPoolsAndAccountWithDelegations(t *testing.T) {
	acountsAddresses, validatorsAddresses := testcosmos.CreateAccounts(2, 3)

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
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, math.NewInt(100))

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), math.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	testHelper.StakingUtils.MessageDelegate(4, 0, validatorsAddresses[0], acountsAddresses[1], math.NewInt(200000))

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1300000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(300000),
		DelegatedVestingAmount:  math.NewInt(200000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1150000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(150000),
		DelegatedVestingAmount:  math.NewInt(150000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(7750))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1075000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(75000),
		DelegatedVestingAmount:  math.NewInt(75000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1000000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(0),
		DelegatedVestingAmount:  math.NewInt(0),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

}

func TestVestingsAmountPoolsAndAccountWithUnbondingDelegations(t *testing.T) {
	acountsAddresses, validatorsAddresses := testcosmos.CreateAccounts(2, 3)

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
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, math.NewInt(100))

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), math.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)
	testHelper.StakingUtils.MessageDelegate(4, 0, validatorsAddresses[0], acountsAddresses[1], math.NewInt(200000))

	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.IncrementContextBlockHeight()
	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	testHelper.StakingUtils.MessageUndelegate(5, 0, validatorsAddresses[0], acountsAddresses[1], math.NewInt(100000))
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1300000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(300000),
		DelegatedVestingAmount:  math.NewInt(200000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1150000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(150000),
		DelegatedVestingAmount:  math.NewInt(150000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(7750))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1075000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(75000),
		DelegatedVestingAmount:  math.NewInt(75000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1000000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(0),
		DelegatedVestingAmount:  math.NewInt(0),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)
}

func TestVestingsAmountPoolsAndAccountWithUnbondingDelegationsEnded(t *testing.T) {
	acountsAddresses, validatorsAddresses := testcosmos.CreateAccounts(2, 3)

	start := testutils.CreateTimeFromNumOfHours(100000)
	lockEnd := testutils.CreateTimeFromNumOfHours(100000)
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
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, math.NewInt(100))

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(math.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), math.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)
	testHelper.StakingUtils.MessageDelegate(4, 0, validatorsAddresses[0], acountsAddresses[1], math.NewInt(200000))

	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.IncrementContextBlockHeight()

	testHelper.App.BeginBlock(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	testHelper.StakingUtils.MessageUndelegate(5, 0, validatorsAddresses[0], acountsAddresses[1], math.NewInt(100000))
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1300000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(300000),
		DelegatedVestingAmount:  math.NewInt(200000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(503))
	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.StakingUtils.VerifyNumberOfUnbondingDelegations(1, acountsAddresses[1])

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1300000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(300000),
		DelegatedVestingAmount:  math.NewInt(200000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(504))

	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.StakingUtils.VerifyNumberOfUnbondingDelegations(0, acountsAddresses[1])

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        math.NewInt(1300000),
		VestingInPoolsAmount:    math.NewInt(1000000),
		VestingInAccountsAmount: math.NewInt(300000),
		DelegatedVestingAmount:  math.NewInt(100000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

}
