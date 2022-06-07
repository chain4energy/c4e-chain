package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

// func TestVestDelegationNotAllowed(t *testing.T) {
// 	vestTest(t)
// }

// func TestVestDelegationAllowed(t *testing.T) {
// 	vestTest(t, true)
// }

// func TestVestLockupZeroDelegationNotAllowed(t *testing.T) {
// 	vestTestLockupZero(t)
// }

// func TestVestLockupZeroDelegationAllowed(t *testing.T) {
// 	vestTestLockupZero(t, true)
// }

func TestCreateVestingPool(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, "v-pool-1", 1000, *usedVestingType, vested, accInitBalance, 0 /*0,*/, accInitBalance-vested /*0,*/, vested)

	verifyAccountVestingPools(t, ctx, app, accAddr, []string{"v-pool-1"}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{0})

	createVestingPool(t, ctx, app, accAddr, true, true, "v-pool-2", 1200, *usedVestingType, vested, accInitBalance-vested /*0,*/, vested, accInitBalance-2*vested /*0,*/, 2*vested)

	verifyAccountVestingPools(t, ctx, app, accAddr, []string{"v-pool-1", "v-pool-2"}, []time.Duration{1000, 1200}, []types.VestingType{*usedVestingType, *usedVestingType}, []int64{vested, vested}, []int64{0, 0})

}

func TestCreateVestingPoolUnknownVestingType(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	setupVestingTypes(ctx, app, 2, 1, 1)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgCreateVestingPool{Creator: accAddr.String(), Name: "pool",
		Amount: sdk.NewInt(vested), Duration: 1000, VestingType: "unknown"}
	_, err := msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting type not found: unknown: not found")

}

func TestCreateVestingPoolNameDuplication(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, "v-pool-1", 1000, *usedVestingType, vested, accInitBalance, 0 /*0,*/, accInitBalance-vested /*0,*/, vested)

	verifyAccountVestingPools(t, ctx, app, accAddr, []string{"v-pool-1"}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{0})

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgCreateVestingPool{Creator: accAddr.String(), Name: "v-pool-1",
		Amount: sdk.NewInt(vested), Duration: 1000, VestingType: usedVestingType.Name}
	_, err := msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting pool name already exists: v-pool-1: invalid request")

}

// func vestTestLockupZero(t *testing.T) {
// 	addHelperModuleAccountPerms()
// 	const vested = 1000
// 	app, ctx := setupApp(1000)

// 	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

// 	accAddr := acountsAddresses[0]

// 	const accInitBalance = 10000
// 	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

// 	modifyVestingType := func(vt *types.VestingType) {
// 		vt.LockupPeriod = 0
// 	}

// 	vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 2, 1, 1)
// 	usedVestingType := vestingTypes.VestingTypes[0]

// 	createVestingPool(t, ctx, app, accAddr, false, true, /*false, false,*/ *usedVestingType, vested, accInitBalance, /*0,*/ 0, accInitBalance-vested, /*0,*/ vested)

// 	verifyAccountVestingPools(t, ctx, app, accAddr, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{0})

// 	createVestingPool(t, ctx, app, accAddr, true, true, /*false, false,*/ *usedVestingType, vested, accInitBalance-vested, /*0,*/ vested, accInitBalance-2*vested, /*0,*/ 2*vested)

// 	verifyAccountVestingPools(t, ctx, app, accAddr, []types.VestingType{*usedVestingType, *usedVestingType}, []int64{vested, vested}, []int64{0, 0})

// }

// func TestVestFirstDelagtionNotAllowedSecondAllowed(t *testing.T) {
// 	addHelperModuleAccountPerms()
// 	const vested = 1000
// 	app, ctx := setupApp(1000)

// 	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

// 	accAddr := acountsAddresses[0]

// 	const accInitBalance = 10000
// 	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

// 	// vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 1, 1, false, 1)

// 	// vestingTypes := setupVestingTypes(ctx, app, 2, 2, delegationAllowed, 1)

// 	i := 1
// 	modifyVestingType := func(vt *types.VestingType) {
// 		if i == 1 {
// 			vt.DelegationsAllowed = true
// 		}
// 		i++
// 	}

// 	vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 2, 2, false, 1)
// 	delegableVestingType := vestingTypes.VestingTypes[0]
// 	nonDelegableVestingType := vestingTypes.VestingTypes[1]

// 	makeVesting(t, ctx, app, accAddr, false, true, false, false, *nonDelegableVestingType, vested, accInitBalance, 0, 0, accInitBalance-vested, 0, vested)

// 	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*nonDelegableVestingType}, []int64{vested}, []int64{0})

// 	makeVesting(t, ctx, app, accAddr, true, true, true, true, *delegableVestingType, vested, accInitBalance-vested, 0, vested, accInitBalance-2*vested, vested, vested)

// 	verifyAccountVestings(t, ctx, app, accAddr, []types.VestingType{*nonDelegableVestingType, *delegableVestingType}, []int64{vested, vested}, []int64{0, 0})

// }

func TestVestingId(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	const accInitBalance = 10000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	addr := accAddr.String()

	k := app.CfevestingKeeper

	k.SetVestingTypes(ctx, vestingTypes)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)

	msg := types.MsgCreateVestingPool{Creator: addr, Name: "v-pool-1", Amount: sdk.NewInt(vested), Duration: 1000, VestingType: usedVestingType.Name}
	_, error := msgServer.CreateVestingPool(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVesting, accFound := k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 1, len(accVesting.VestingPools))

	vesting := accVesting.VestingPools[0]
	require.EqualValues(t, 1, vesting.Id)

	msg.Name = "v-pool-2"
	_, error = msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualValues(t, nil, error)

	accVesting, accFound = k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 2, len(accVesting.VestingPools))

	vesting = accVesting.VestingPools[0]
	require.EqualValues(t, 1, vesting.Id)

	vesting = accVesting.VestingPools[1]
	require.EqualValues(t, 2, vesting.Id)

	msg.Name = "v-pool-3"
	_, error = msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualValues(t, nil, error)

	accVesting, accFound = k.GetAccountVestings(ctx, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 3, len(accVesting.VestingPools))

	vesting = accVesting.VestingPools[0]
	require.EqualValues(t, 1, vesting.Id)

	vesting = accVesting.VestingPools[1]
	require.EqualValues(t, 2, vesting.Id)

	vesting = accVesting.VestingPools[2]
	require.EqualValues(t, 3, vesting.Id)
}
