package keeper_test

import (
	"testing"
	// "time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	"github.com/stretchr/testify/require"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

func TestSendVestingAccount(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper, ctx := testapp.SetupTestAppWithHeight(t, 1000)
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, accInitBalance, accAddr)

	vestingTypes := vestingTestHelper.SetupVestingTypes(ctx, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	vestingTestHelper.CreateVestingPool(ctx, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)

	account := testHelper.App.AccountKeeper.GetAccount(ctx, accAddr2)

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(ctx, accAddr2, sdk.NewInt(100))

	require.Equal(t, uint64(1), testHelper.App.CfevestingKeeper.GetVestingAccountCount(ctx))
	vaccFromList, found := testHelper.App.CfevestingKeeper.GetVestingAccount(ctx, uint64(0))
	require.Equal(t, true, found)
	require.Equal(t, accAddr2.String(), vaccFromList.Address)

	vacc, ok := account.(vestexported.VestingAccount)
	require.Equal(t, true, ok)
	locked := vacc.LockedCoins(ctx.BlockTime())
	require.Equal(t, commontestutils.DefaultTestDenom, locked[0].Denom)
	require.Equal(t, sdk.NewInt(100), locked[0].Amount)

	require.Equal(t, (ctx.BlockTime().UnixNano()+int64(usedVestingType.VestingPeriod+usedVestingType.LockupPeriod))/1000000000, vacc.GetEndTime())

	require.Equal(t, (ctx.BlockTime().UnixNano()+int64(usedVestingType.LockupPeriod))/1000000000, vacc.GetStartTime())

}

func TestSendVestingAccountVestingPoolNotExistsForAddress(t *testing.T) {
	testHelper, ctx := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, accInitBalance, accAddr)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 2, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"rpc error: code = NotFound desc = No vestings")

	require.Equal(t, uint64(0), testHelper.App.CfevestingKeeper.GetVestingAccountCount(ctx))

}

func TestSendVestingAccountVestingPoolNotFound(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper, ctx := testapp.SetupTestAppWithHeight(t, 1000)
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, accInitBalance, accAddr)

	vestingTypes := vestingTestHelper.SetupVestingTypes(ctx, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	vestingTestHelper.CreateVestingPool(ctx, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 2, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting pool with id 2 not found: not found")

	require.Equal(t, uint64(0), testHelper.App.CfevestingKeeper.GetVestingAccountCount(ctx))

}

func TestSendVestingAccounNotEnoughToSend(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper, ctx := testapp.SetupTestAppWithHeight(t, 1000)
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, accInitBalance, accAddr)

	vestingTypes := vestingTestHelper.SetupVestingTypes(ctx, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	vestingTestHelper.CreateVestingPool(ctx, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(1100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting available: 1000 is smaller than 1100: insufficient funds")

	require.Equal(t, uint64(0), testHelper.App.CfevestingKeeper.GetVestingAccountCount(ctx))
}

func TestSendVestingAccountNotEnoughToSendAferSuccesfulSend(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper, ctx := testapp.SetupTestAppWithHeight(t, 1000)
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, accInitBalance, accAddr)

	vestingTypes := vestingTestHelper.SetupVestingTypes(ctx, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	vestingTestHelper.CreateVestingPool(ctx, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)

	msg = types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(950), RestartVesting: true}
	_, err = msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting available: 900 is smaller than 950: insufficient funds")

	require.Equal(t, uint64(1), testHelper.App.CfevestingKeeper.GetVestingAccountCount(ctx))
	vaccFromList, found := testHelper.App.CfevestingKeeper.GetVestingAccount(ctx, uint64(0))
	require.Equal(t, true, found)
	require.Equal(t, accAddr2.String(), vaccFromList.Address)
}

func TestSendVestingAccountAlreadyExists(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper, ctx := testapp.SetupTestAppWithHeight(t, 1000)
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, accInitBalance, accAddr)

	vestingTypes := vestingTestHelper.SetupVestingTypes(ctx, 2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	vestingTestHelper.CreateVestingPool(ctx, accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(100), RestartVesting: true}
	_, err := msgServer.SendToVestingAccount(msgServerCtx, &msg)
	require.EqualValues(t, nil, err)

	msg = types.MsgSendToVestingAccount{FromAddress: accAddr.String(), ToAddress: accAddr2.String(),
		VestingId: 1, Amount: sdk.NewInt(300), RestartVesting: true}
	_, err = msgServer.SendToVestingAccount(msgServerCtx, &msg)

	require.EqualError(t, err,
		"account "+accAddr2.String()+" already exists: invalid request")

	require.Equal(t, uint64(1), testHelper.App.CfevestingKeeper.GetVestingAccountCount(ctx))
	vaccFromList, found := testHelper.App.CfevestingKeeper.GetVestingAccount(ctx, uint64(0))
	require.Equal(t, true, found)
	require.Equal(t, accAddr2.String(), vaccFromList.Address)
}
