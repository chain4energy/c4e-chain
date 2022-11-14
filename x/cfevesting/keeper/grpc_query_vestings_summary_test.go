package keeper_test

import (
	"testing"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestVestingsAmountPoolsOnly(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
	}
	accVestingPools := types.AccountVestingPools{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(commontestutils.DefaultTestDenom),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := testapp.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1000000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.ZeroInt(),
		DelegatedVestingAmount:  sdk.ZeroInt(),
	}

	testHelper.C4eVestingUtils.QueryVestings(expected)
}

func TestVestingsAmountPoolsAndAccount(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
	}

	accVestingPools := types.AccountVestingPools{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(commontestutils.DefaultTestDenom),
		VestingAccountList: []types.VestingAccount{
			{
				Id:      0,
				Address: acountsAddresses[1].String(),
			},
		},
		VestingAccountCount: 1,
		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := testapp.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(1).SubRaw(1),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(1).SubRaw(1),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1000000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(0),
		DelegatedVestingAmount:  sdk.NewInt(0),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

}

func TestVestingsAmountPoolsAndAccountWithDelegations(t *testing.T) {
	acountsAddresses, validatorsAddresses := commontestutils.CreateAccounts(2, 3)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
	}

	accVestingPools := types.AccountVestingPools{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(commontestutils.DefaultTestDenom),
		VestingAccountList: []types.VestingAccount{
			{
				Id:      0,
				Address: acountsAddresses[1].String(),
			},
		},
		VestingAccountCount: 1,
		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := testapp.SetupTestApp(t)
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, sdk.NewInt(100))

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)

	testHelper.StakingUtils.MessageDelegate(4, 0, validatorsAddresses[0], acountsAddresses[1], sdk.NewInt(200000))

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(150000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(7750))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1075000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(75000),
		DelegatedVestingAmount:  sdk.NewInt(75000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1000000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(0),
		DelegatedVestingAmount:  sdk.NewInt(0),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

}

func TestVestingsAmountPoolsAndAccountWithUnbondingDelegations(t *testing.T) {
	acountsAddresses, validatorsAddresses := commontestutils.CreateAccounts(2, 3)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
	}

	accVestingPools := types.AccountVestingPools{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(commontestutils.DefaultTestDenom),
		VestingAccountList: []types.VestingAccount{
			{
				Id:      0,
				Address: acountsAddresses[1].String(),
			},
		},
		VestingAccountCount: 1,
		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := testapp.SetupTestApp(t)
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, sdk.NewInt(100))

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)
	testHelper.StakingUtils.MessageDelegate(4, 0, validatorsAddresses[0], acountsAddresses[1], sdk.NewInt(200000))

	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.IncrementContextBlockHeight()
	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	testHelper.StakingUtils.MessageUndelegate(5, 0, validatorsAddresses[0], acountsAddresses[1], sdk.NewInt(100000))
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(150000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(7750))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1075000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(75000),
		DelegatedVestingAmount:  sdk.NewInt(75000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1000000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(0),
		DelegatedVestingAmount:  sdk.NewInt(0),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)
}

func TestVestingsAmountPoolsAndAccountWithUnbondingDelegationsEnded(t *testing.T) {
	acountsAddresses, validatorsAddresses := commontestutils.CreateAccounts(2, 3)

	start := testutils.CreateTimeFromNumOfHours(100000)
	lockEnd := testutils.CreateTimeFromNumOfHours(100000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		VestingType:     "test",
		LockStart:       start,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
	}

	accVestingPools := types.AccountVestingPools{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingPoolsArray := []*types.AccountVestingPools{&accVestingPools}

	genesisState := types.GenesisState{
		Params: types.NewParams(commontestutils.DefaultTestDenom),
		VestingAccountList: []types.VestingAccount{
			{
				Id:      0,
				Address: acountsAddresses[1].String(),
			},
		},
		VestingAccountCount: 1,
		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingPools: accountVestingPoolsArray,
	}

	testHelper := testapp.SetupTestApp(t)
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, sdk.NewInt(100))

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	testHelper.C4eVestingUtils.InitGenesis(genesisState)
	testHelper.StakingUtils.MessageDelegate(4, 0, validatorsAddresses[0], acountsAddresses[1], sdk.NewInt(200000))

	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.IncrementContextBlockHeight()

	testHelper.App.BeginBlock(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	testHelper.StakingUtils.MessageUndelegate(5, 0, validatorsAddresses[0], acountsAddresses[1], sdk.NewInt(100000))
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})

	expected := types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(503))
	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.StakingUtils.VerifyNumberOfUnbondingDelegations(1, acountsAddresses[1])

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(504))

	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.StakingUtils.VerifyNumberOfUnbondingDelegations(0, acountsAddresses[1])

	expected = types.QueryVestingsSummaryResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(100000),
	}
	testHelper.C4eVestingUtils.QueryVestings(expected)

}
