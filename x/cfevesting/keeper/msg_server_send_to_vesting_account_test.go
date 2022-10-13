package keeper_test

import (
	"testing"
	// "time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	"github.com/stretchr/testify/require"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

func TestSendVestingAccount(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, 0 /*0,*/, accInitBalance-vested /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)

	account := app.AccountKeeper.GetAccount(ctx, accAddr2)

	bal := app.BankKeeper.GetBalance(ctx, accAddr2, commontestutils.Denom)
	require.Equal(t, sdk.NewInt(100), bal.Amount)

	require.Equal(t, uint64(1), app.CfevestingKeeper.GetVestingAccountCount(ctx))
	vaccFromList, found := app.CfevestingKeeper.GetVestingAccount(ctx, uint64(0))
	require.Equal(t, true, found)
	require.Equal(t, accAddr2.String(), vaccFromList.Address)

	vacc, ok := account.(vestexported.VestingAccount)
	require.Equal(t, true, ok)
	locked := vacc.LockedCoins(ctx.BlockTime())
	require.Equal(t, commontestutils.Denom, locked[0].Denom)
	require.Equal(t, sdk.NewInt(100), locked[0].Amount)

	require.Equal(t, (ctx.BlockTime().UnixNano()+int64(usedVestingType.VestingPeriod+usedVestingType.LockupPeriod))/1000000000, vacc.GetEndTime())

	require.Equal(t, (ctx.BlockTime().UnixNano()+int64(usedVestingType.LockupPeriod))/1000000000, vacc.GetStartTime())

}

func TestSendVestingAccountVestingPoolNotExistsForAddress(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 2, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err, "no vestings found: failed to withdraw all available")

	require.Equal(t, uint64(0), app.CfevestingKeeper.GetVestingAccountCount(ctx))

}

func TestSendVestingAccountVestingPoolNotFound(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, 0 /*0,*/, accInitBalance-vested /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 2, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err, "vesting pool with id 2 not found: vesting pool not found")

	require.Equal(t, uint64(0), app.CfevestingKeeper.GetVestingAccountCount(ctx))

}

func TestSendVestingAccounNotEnoughToSend(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, 0 /*0,*/, accInitBalance-vested /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(1100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting available: 1000 is smaller than 1100: vesting available is smaller than amount")

	require.Equal(t, uint64(0), app.CfevestingKeeper.GetVestingAccountCount(ctx))
}

func TestSendVestingAccountNotEnoughToSendAferSuccesfulSend(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, 0 /*0,*/, accInitBalance-vested /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)

	msg = types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(950), RestartVesting: true}
	_, err = msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting available: 900 is smaller than 950: vesting available is smaller than amount")

	require.Equal(t, uint64(1), app.CfevestingKeeper.GetVestingAccountCount(ctx))
	vaccFromList, found := app.CfevestingKeeper.GetVestingAccount(ctx, uint64(0))
	require.Equal(t, true, found)
	require.Equal(t, accAddr2.String(), vaccFromList.Address)
}

func TestSendVestingAccountAlreadyExists(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000
	app, ctx := commontestutils.SetupApp(1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	const accInitBalance = 10000
	commontestutils.AddCoinsToAccount(accInitBalance, ctx, app, accAddr)

	vestingTypes := setupVestingTypes(ctx, app, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	createVestingPool(t, ctx, app, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, 0 /*0,*/, accInitBalance-vested /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)

	msg = types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(300), RestartVesting: true}
	_, err = msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"account address: "+accAddr2.String()+": account account already exists")

	require.Equal(t, uint64(1), app.CfevestingKeeper.GetVestingAccountCount(ctx))
	vaccFromList, found := app.CfevestingKeeper.GetVestingAccount(ctx, uint64(0))
	require.Equal(t, true, found)
	require.Equal(t, accAddr2.String(), vaccFromList.Address)
}
