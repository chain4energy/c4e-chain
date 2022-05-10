package keeper_test

import (
	"testing"
	// "time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

func TestSendVestingAccount(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, "v-pool-1", 1000, *usedVestingType, vested, accInitBalance, 0, /*0,*/ accInitBalance-vested, /*0,*/ vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(), 
		VestingId: 1, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)
	
	account := app.AccountKeeper.GetAccount(ctx, accAddr2)

	bal := app.BankKeeper.GetBalance(ctx, accAddr2, denom)
	require.Equal(t, sdk.NewInt(100), bal.Amount)

	vacc, ok := account.(vestexported.VestingAccount)
	require.Equal(t, true, ok)
	locked := vacc.LockedCoins(ctx.BlockTime())
	require.Equal(t, denom, locked[0].Denom)
	require.Equal(t, sdk.NewInt(100), locked[0].Amount)

	require.Equal(t, (ctx.BlockTime().UnixNano() + int64(usedVestingType.VestingPeriod + usedVestingType.LockupPeriod))/1000000000, vacc.GetEndTime())

	require.Equal(t, (ctx.BlockTime().UnixNano() + int64(usedVestingType.LockupPeriod))/1000000000, vacc.GetStartTime())

}

func TestSendVestingAccountVestingPoolNotExistsForAddress(t *testing.T) {
	addHelperModuleAccountPerms()
	// const vested = 1000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	// vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	// usedVestingType := vestingTypes.VestingTypes[0]

	// createVestingPool(t, ctx, app, accAddr, false, true, "v-pool-1", 1000, *usedVestingType, vested, accInitBalance, 0, /*0,*/ accInitBalance-vested, /*0,*/ vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(), 
		VestingId: 2, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"rpc error: code = NotFound desc = No vestings")

}

func TestSendVestingAccountVestingPoolNotFound(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, "v-pool-1", 1000, *usedVestingType, vested, accInitBalance, 0, /*0,*/ accInitBalance-vested, /*0,*/ vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(), 
		VestingId: 2, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting pool with id 2 not found: not found")

}


func TestSendVestingAccounNotEnoughToSend(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, "v-pool-1", 1000, *usedVestingType, vested, accInitBalance, 0, /*0,*/ accInitBalance-vested, /*0,*/ vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(), 
		VestingId: 1, Amount: sdk.NewInt(1100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting available: 1000 is smaller than 1100: insufficient funds")

}

func TestSendVestingAccountNotEnoughToSendAferSuccesfulSend(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := setupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	addCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, "v-pool-1", 1000, *usedVestingType, vested, accInitBalance, 0, /*0,*/ accInitBalance-vested, /*0,*/ vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(), 
		VestingId: 1, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)

	msg = types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(), 
		VestingId: 1, Amount: sdk.NewInt(950), RestartVesting: true}
	_, err = msgServer.SendToVestingAccount(msgServerCtx, &msg)
	
	require.EqualError(t, err,
		"vesting available: 900 is smaller than 950: insufficient funds")
	// account := app.AccountKeeper.GetAccount(ctx, accAddr2)

	// bal := app.BankKeeper.GetBalance(ctx, accAddr2, denom)
	// require.Equal(t, sdk.NewInt(100), bal.Amount)

	// vacc, ok := account.(vestexported.VestingAccount)
	// require.Equal(t, true, ok)
	// locked := vacc.LockedCoins(ctx.BlockTime())
	// require.Equal(t, denom, locked[0].Denom)
	// require.Equal(t, sdk.NewInt(100), locked[0].Amount)

	// require.Equal(t, (ctx.BlockTime().UnixNano() + int64(usedVestingType.VestingPeriod + usedVestingType.LockupPeriod))/1000000000, vacc.GetEndTime())

	// require.Equal(t, (ctx.BlockTime().UnixNano() + int64(usedVestingType.LockupPeriod))/1000000000, vacc.GetStartTime())

}