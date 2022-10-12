package keeper_test

import (
	// "fmt"
	"testing"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestVestingsAmountPoolsOnly(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		Id:                        1,
		VestingType:               "test",
		LockStart:                 start,
		LockEnd:                   lockEnd,
		Vested:                    amount,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          start,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	accVestings := types.AccountVestings{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingsListArray := []*types.AccountVestings{&accVestings}

	genesisState := types.GenesisState{
		Params: types.NewParams(commontestutils.DefaultTestDenom),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	testHelper := testapp.SetupTestApp(t)

	wctx := sdk.WrapSDKContext(testHelper.Context)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	request := types.QueryVestingsRequest{}

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1000000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.ZeroInt(),
		DelegatedVestingAmount:  sdk.ZeroInt(),
	}

	require.Equal(t, expected, *resp)

}

func TestVestingsAmountPoolsAndAccount(t *testing.T) {
	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		Id:                        1,
		VestingType:               "test",
		LockStart:                 start,
		LockEnd:                   lockEnd,
		Vested:                    amount,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          start,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	accVestings := types.AccountVestings{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingsListArray := []*types.AccountVestings{&accVestings}

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
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	testHelper := testapp.SetupTestApp(t)

	wctx := sdk.WrapSDKContext(testHelper.Context)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	request := types.QueryVestingsRequest{}

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(1).SubRaw(1),
	}
	require.Equal(t, expected, *resp)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))
	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(1).SubRaw(1),
	}
	require.Equal(t, expected, *resp)
	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))
	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1000000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(0),
		DelegatedVestingAmount:  sdk.NewInt(0),
	}
	require.Equal(t, expected, *resp)

}

func TestVestingsAmountPoolsAndAccountWithDelegations(t *testing.T) {
	acountsAddresses, validatorsAddresses := commontestutils.CreateAccounts(2, 3)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		Id:                        1,
		VestingType:               "test",
		LockStart:                 start,
		LockEnd:                   lockEnd,
		Vested:                    amount,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          start,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	accVestings := types.AccountVestings{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingsListArray := []*types.AccountVestings{&accVestings}

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
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}


	testHelper := testapp.SetupTestApp(t)
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, sdk.NewInt(100))

	wctx := sdk.WrapSDKContext(testHelper.Context)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	request := types.QueryVestingsRequest{}

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)
	validator, _ := testHelper.StakingUtils.GetValidator(validatorsAddresses[0])

	require.Equal(t, 4, len(testHelper.App.StakingKeeper.GetAllDelegations(testHelper.Context)))
	require.Equal(t, 0, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))
	testHelper.App.StakingKeeper.Delegate(testHelper.Context, acountsAddresses[1], sdk.NewInt(200000), stakingtypes.Unbonded,
		validator, true)

	require.Equal(t, 5, len(testHelper.App.StakingKeeper.GetAllDelegations(testHelper.Context)))
	require.Equal(t, 0, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	require.Equal(t, expected, *resp)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))
	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(150000),
	}
	require.Equal(t, expected, *resp)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(7750))
	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1075000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(75000),
		DelegatedVestingAmount:  sdk.NewInt(75000),
	}
	require.Equal(t, expected, *resp)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))
	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1000000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(0),
		DelegatedVestingAmount:  sdk.NewInt(0),
	}
	require.Equal(t, expected, *resp)

}

func TestVestingsAmountPoolsAndAccountWithUnbondingDelegations(t *testing.T) {
	acountsAddresses, validatorsAddresses := commontestutils.CreateAccounts(2, 3)

	start := testutils.CreateTimeFromNumOfHours(1000)
	lockEnd := testutils.CreateTimeFromNumOfHours(10000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		Id:                        1,
		VestingType:               "test",
		LockStart:                 start,
		LockEnd:                   lockEnd,
		Vested:                    amount,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          start,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	accVestings := types.AccountVestings{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingsListArray := []*types.AccountVestings{&accVestings}

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
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	testHelper := testapp.SetupTestApp(t)
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, sdk.NewInt(100))

	wctx := sdk.WrapSDKContext(testHelper.Context)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	request := types.QueryVestingsRequest{}

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)
	validator, _ := testHelper.StakingUtils.GetValidator(validatorsAddresses[0])

	require.Equal(t, 4, len(testHelper.App.StakingKeeper.GetAllDelegations(testHelper.Context)))
	require.Equal(t, 0, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))
	testHelper.App.StakingKeeper.Delegate(testHelper.Context, acountsAddresses[1], sdk.NewInt(200000), stakingtypes.Unbonded,
		validator, true)

	require.Equal(t, 5, len(testHelper.App.StakingKeeper.GetAllDelegations(testHelper.Context)))
	require.Equal(t, 0, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))

	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.IncrementContextBlockHeight()
	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	testHelper.App.StakingKeeper.Undelegate(testHelper.Context, acountsAddresses[1], validatorsAddresses[0], sdk.NewDec(100000))

	require.Equal(t, 5, len(testHelper.App.StakingKeeper.GetAllDelegations(testHelper.Context)))
	require.Equal(t, 1, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	require.Equal(t, expected, *resp)
	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(5500))

	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(150000),
	}
	require.Equal(t, expected, *resp)
	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(7750))
	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1075000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(75000),
		DelegatedVestingAmount:  sdk.NewInt(75000),
	}
	require.Equal(t, expected, *resp)

	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(10000))
	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1000000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(0),
		DelegatedVestingAmount:  sdk.NewInt(0),
	}
	require.Equal(t, expected, *resp)

}

func TestVestingsAmountPoolsAndAccountWithUnbondingDelegationsEnded(t *testing.T) {
	acountsAddresses, validatorsAddresses := commontestutils.CreateAccounts(2, 3)

	start := testutils.CreateTimeFromNumOfHours(100000)
	lockEnd := testutils.CreateTimeFromNumOfHours(100000)
	amount := sdk.NewInt(1000000)

	vestingPool := types.VestingPool{
		Id:                        1,
		VestingType:               "test",
		LockStart:                 start,
		LockEnd:                   lockEnd,
		Vested:                    amount,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          start,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}

	accVestings := types.AccountVestings{
		Address:      acountsAddresses[0].String(),
		VestingPools: []*types.VestingPool{&vestingPool},
	}

	accountVestingsListArray := []*types.AccountVestings{&accVestings}

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
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	testHelper := testapp.SetupTestApp(t)
	testHelper.StakingUtils.SetupValidators(validatorsAddresses, sdk.NewInt(100))

	wctx := sdk.WrapSDKContext(testHelper.Context)

	k := testHelper.App.CfevestingKeeper
	ak := testHelper.App.AccountKeeper

	request := types.QueryVestingsRequest{}

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(1000000), types.ModuleName)
	testHelper.AuthUtils.CreateDefaultDenomVestingAccount(acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	cfevesting.InitGenesis(testHelper.Context, k, genesisState, ak, testHelper.App.BankKeeper, testHelper.App.StakingKeeper)
	validator, _ := testHelper.StakingUtils.GetValidator(validatorsAddresses[0])

	require.Equal(t, 4, len(testHelper.App.StakingKeeper.GetAllDelegations(testHelper.Context)))
	require.Equal(t, 0, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))
	testHelper.App.StakingKeeper.Delegate(testHelper.Context, acountsAddresses[1], sdk.NewInt(200000), stakingtypes.Unbonded,
		validator, true)

	require.Equal(t, 5, len(testHelper.App.StakingKeeper.GetAllDelegations(testHelper.Context)))
	require.Equal(t, 0, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))

	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	testHelper.IncrementContextBlockHeight()

	testHelper.App.BeginBlock(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})

	testHelper.App.StakingKeeper.Undelegate(testHelper.Context, acountsAddresses[1], validatorsAddresses[0], sdk.NewDec(100000))

	require.Equal(t, 5, len(testHelper.App.StakingKeeper.GetAllDelegations(testHelper.Context)))
	require.Equal(t, 1, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	require.Equal(t, expected, *resp)
	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(503))
	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	require.Equal(t, 1, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))

	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	require.Equal(t, expected, *resp)
	testHelper.IncrementContextBlockHeightAndSetTime(testutils.CreateTimeFromNumOfHours(504))

	testHelper.BeginBlocker(abci.RequestBeginBlock{Header: testHelper.Context.BlockHeader()})
	testHelper.EndBlocker(abci.RequestEndBlock{Height: testHelper.Context.BlockHeight()})
	require.Equal(t, 0, len(testHelper.App.StakingKeeper.GetAllUnbondingDelegations(testHelper.Context, acountsAddresses[1])))

	wctx = sdk.WrapSDKContext(testHelper.Context)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(100000),
	}
	require.Equal(t, expected, *resp)

}
