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

func TestCreateVestingPool(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, "v-pool-1", 1000, *usedVestingType, vested, accInitBalance, 0 /*0,*/, accInitBalance-vested /*0,*/, vested)

	verifyAccountVestingPools(t, ctx, app, accAddr, []string{"v-pool-1"}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []int64{vested}, []int64{0})

	createVestingPool(t, ctx, app, accAddr, true, true, "v-pool-2", 1200, *usedVestingType, vested, accInitBalance-vested /*0,*/, vested, accInitBalance-2*vested /*0,*/, 2*vested)

	verifyAccountVestingPools(t, ctx, app, accAddr, []string{"v-pool-1", "v-pool-2"}, []time.Duration{1000, 1200}, []types.VestingType{*usedVestingType, *usedVestingType}, []int64{vested, vested}, []int64{0, 0})

}

func TestCreateVestingPoolUnknownVestingType(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

	setupVestingTypes(ctx, app, 2, 1, 1)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgCreateVestingPool{Creator: accAddr.String(), Name: "pool",
		Amount: sdk.NewInt(vested), Duration: 1000, VestingType: "unknown"}
	_, err := msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting type not found: unknown: not found")

}

func TestCreateVestingPoolNameDuplication(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

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

func TestVestingId(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	const accInitBalance = 10000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

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
