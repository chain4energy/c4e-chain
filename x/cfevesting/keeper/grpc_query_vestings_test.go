package keeper_test

import (
	// "fmt"
	"testing"

	"github.com/chain4energy/c4e-chain/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestVestingsAmountPoolsOnly(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
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
		Params: types.NewParams(commontestutils.Denom),

		VestingTypes:        []types.GenesisVestingType{},
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 0, Time: commontestutils.TestEnvTime})
	wctx := sdk.WrapSDKContext(ctx)

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	request := types.QueryVestingsRequest{}

	commontestutils.AddCoinsToModuleByName(1000000, types.ModuleName, ctx, app)
	cfevesting.InitGenesis(ctx, k, genesisState, ak, app.BankKeeper, app.StakingKeeper)

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
	commontestutils.AddHelperModuleAccountPerms()
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
		Params: types.NewParams(commontestutils.Denom),
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

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 0, Time: commontestutils.TestEnvTime})
	wctx := sdk.WrapSDKContext(ctx)

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	request := types.QueryVestingsRequest{}

	commontestutils.AddCoinsToModuleByName(1000000, types.ModuleName, ctx, app)
	commontestutils.CreateVestingAccount(ctx, app, acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	cfevesting.InitGenesis(ctx, k, genesisState, ak, app.BankKeeper, app.StakingKeeper)

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(1).SubRaw(1),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(5500))
	// ctx = app.BaseApp.NewContext(false, tmproto.Header{Height: 0, Time: testutils.CreateTimeFromNumOfHours(5500)})
	wctx = sdk.WrapSDKContext(ctx)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(1).SubRaw(1),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(10000))
	wctx = sdk.WrapSDKContext(ctx)

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
	commontestutils.AddHelperModuleAccountPerms()
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
		Params: types.NewParams(commontestutils.Denom),
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

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 0, Time: commontestutils.TestEnvTime})
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = commontestutils.Denom
	app.StakingKeeper.SetParams(ctx, stakingParams)
	setupValidators(t, ctx, app, validatorsAddresses, uint64(100))

	wctx := sdk.WrapSDKContext(ctx)

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	request := types.QueryVestingsRequest{}

	commontestutils.AddCoinsToModuleByName(1000000, types.ModuleName, ctx, app)
	commontestutils.CreateVestingAccount(ctx, app, acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	cfevesting.InitGenesis(ctx, k, genesisState, ak, app.BankKeeper, app.StakingKeeper)
	validator, _ := app.StakingKeeper.GetValidator(ctx, validatorsAddresses[0])

	require.Equal(t, 3, len(app.StakingKeeper.GetAllDelegations(ctx)))
	require.Equal(t, 0, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))
	app.StakingKeeper.Delegate(ctx, acountsAddresses[1], sdk.NewInt(200000), stakingtypes.Unbonded,
		validator, true)

	require.Equal(t, 4, len(app.StakingKeeper.GetAllDelegations(ctx)))
	require.Equal(t, 0, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(5500))
	wctx = sdk.WrapSDKContext(ctx)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(150000),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(7750))
	wctx = sdk.WrapSDKContext(ctx)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1075000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(75000),
		DelegatedVestingAmount:  sdk.NewInt(75000),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(10000))
	wctx = sdk.WrapSDKContext(ctx)

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
	commontestutils.AddHelperModuleAccountPerms()
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
		Params: types.NewParams(commontestutils.Denom),
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

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 0, Time: commontestutils.TestEnvTime})
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = commontestutils.Denom
	app.StakingKeeper.SetParams(ctx, stakingParams)
	setupValidators(t, ctx, app, validatorsAddresses, uint64(100))

	wctx := sdk.WrapSDKContext(ctx)

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	request := types.QueryVestingsRequest{}

	commontestutils.AddCoinsToModuleByName(1000000, types.ModuleName, ctx, app)
	commontestutils.CreateVestingAccount(ctx, app, acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	cfevesting.InitGenesis(ctx, k, genesisState, ak, app.BankKeeper, app.StakingKeeper)
	validator, _ := app.StakingKeeper.GetValidator(ctx, validatorsAddresses[0])

	require.Equal(t, 3, len(app.StakingKeeper.GetAllDelegations(ctx)))
	require.Equal(t, 0, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))
	app.StakingKeeper.Delegate(ctx, acountsAddresses[1], sdk.NewInt(200000), stakingtypes.Unbonded,
		validator, true)

	require.Equal(t, 4, len(app.StakingKeeper.GetAllDelegations(ctx)))
	require.Equal(t, 0, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))

	app.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()})
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	app.BeginBlocker(ctx, abci.RequestBeginBlock{Header: ctx.BlockHeader()})

	app.StakingKeeper.Undelegate(ctx, acountsAddresses[1], validatorsAddresses[0], sdk.NewDec(100000))

	require.Equal(t, 4, len(app.StakingKeeper.GetAllDelegations(ctx)))
	require.Equal(t, 1, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))
	app.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()})

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	require.Equal(t, expected, *resp)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(5500))

	app.BeginBlocker(ctx, abci.RequestBeginBlock{Header: ctx.BlockHeader()})

	wctx = sdk.WrapSDKContext(ctx)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1150000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(150000),
		DelegatedVestingAmount:  sdk.NewInt(150000),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(7750))
	wctx = sdk.WrapSDKContext(ctx)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1075000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(75000),
		DelegatedVestingAmount:  sdk.NewInt(75000),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(10000))
	wctx = sdk.WrapSDKContext(ctx)

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
	commontestutils.AddHelperModuleAccountPerms()
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
		Params: types.NewParams(commontestutils.Denom),
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

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 0, Time: commontestutils.TestEnvTime})
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = commontestutils.Denom
	app.StakingKeeper.SetParams(ctx, stakingParams)
	setupValidators(t, ctx, app, validatorsAddresses, uint64(100))

	wctx := sdk.WrapSDKContext(ctx)

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	request := types.QueryVestingsRequest{}

	commontestutils.AddCoinsToModuleByName(1000000, types.ModuleName, ctx, app)
	commontestutils.CreateVestingAccount(ctx, app, acountsAddresses[1].String(), sdk.NewInt(300000), start, lockEnd)

	cfevesting.InitGenesis(ctx, k, genesisState, ak, app.BankKeeper, app.StakingKeeper)
	validator, _ := app.StakingKeeper.GetValidator(ctx, validatorsAddresses[0])

	require.Equal(t, 3, len(app.StakingKeeper.GetAllDelegations(ctx)))
	require.Equal(t, 0, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))
	app.StakingKeeper.Delegate(ctx, acountsAddresses[1], sdk.NewInt(200000), stakingtypes.Unbonded,
		validator, true)

	require.Equal(t, 4, len(app.StakingKeeper.GetAllDelegations(ctx)))
	require.Equal(t, 0, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))

	app.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()})
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

	app.BeginBlock(abci.RequestBeginBlock{Header: ctx.BlockHeader()})

	app.StakingKeeper.Undelegate(ctx, acountsAddresses[1], validatorsAddresses[0], sdk.NewDec(100000))

	require.Equal(t, 4, len(app.StakingKeeper.GetAllDelegations(ctx)))
	require.Equal(t, 1, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))
	app.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()})

	resp, err := k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected := types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(503))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{Header: ctx.BlockHeader()})
	app.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()})
	require.Equal(t, 1, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))

	wctx = sdk.WrapSDKContext(ctx)

	resp, err = k.Vestings(wctx, &request)
	require.NoError(t, err)

	expected = types.QueryVestingsResponse{
		VestingAllAmount:        sdk.NewInt(1300000),
		VestingInPoolsAmount:    sdk.NewInt(1000000),
		VestingInAccountsAmount: sdk.NewInt(300000),
		DelegatedVestingAmount:  sdk.NewInt(200000),
	}
	require.Equal(t, expected, *resp)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(testutils.CreateTimeFromNumOfHours(504))
	app.BeginBlocker(ctx, abci.RequestBeginBlock{Header: ctx.BlockHeader()})
	app.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()})
	require.Equal(t, 0, len(app.StakingKeeper.GetAllUnbondingDelegations(ctx, acountsAddresses[1])))

	wctx = sdk.WrapSDKContext(ctx)

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
