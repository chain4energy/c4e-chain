package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

func TestCreateVestingPool(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.CreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []sdk.Int{vested}, []sdk.Int{sdk.ZeroInt()})

	testHelper.C4eVestingUtils.CreateVestingPool(accAddr, true, true, vPool2, 1200, *usedVestingType, vested, accInitBalance.Sub(vested) /*0,*/, vested, accInitBalance.Sub(vested.MulRaw(2)) /*0,*/, vested.MulRaw(2))

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1, vPool2}, []time.Duration{1000, 1200}, []types.VestingType{*usedVestingType, *usedVestingType}, []sdk.Int{vested, vested}, []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt()})

}

func TestCreateVestingPoolUnknownVestingType(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(sdk.NewInt(10000), accAddr)

	testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(testHelper.Context)

	msg := types.MsgCreateVestingPool{Creator: accAddr.String(), Name: "pool",
		Amount: vested, Duration: 1000, VestingType: "unknown"}
	_, err := msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting type not found: unknown: not found")

}

func TestCreateVestingPoolNameDuplication(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.CreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.VerifyAccountVestingPools(accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*usedVestingType}, []sdk.Int{vested}, []sdk.Int{sdk.ZeroInt()})

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(testHelper.Context)

	msg := types.MsgCreateVestingPool{Creator: accAddr.String(), Name: vPool1,
		Amount: vested, Duration: 1000, VestingType: usedVestingType.Name}
	_, err := msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualError(t, err,
		"vesting pool name already exists: "+vPool1+": invalid request")

}

func TestVestingId(t *testing.T) {
	vested := sdk.NewInt(1000)
	accInitBalance := sdk.NewInt(10000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	accAddr := acountsAddresses[0]
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	addr := accAddr.String()

	k := testHelper.App.CfevestingKeeper

	k.SetVestingTypes(testHelper.Context, vestingTypes)
	msgServer, msgServerCtx := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(testHelper.Context)

	msg := types.MsgCreateVestingPool{Creator: addr, Name: vPool1, Amount: vested, Duration: 1000, VestingType: usedVestingType.Name}
	_, error := msgServer.CreateVestingPool(msgServerCtx, &msg)
	require.EqualValues(t, nil, error)

	accVesting, accFound := k.GetAccountVestings(testHelper.Context, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 1, len(accVesting.VestingPools))

	vesting := accVesting.VestingPools[0]
	require.EqualValues(t, 1, vesting.Id)

	msg.Name = vPool2
	_, error = msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualValues(t, nil, error)

	accVesting, accFound = k.GetAccountVestings(testHelper.Context, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 2, len(accVesting.VestingPools))

	vesting = accVesting.VestingPools[0]
	require.EqualValues(t, 1, vesting.Id)

	vesting = accVesting.VestingPools[1]
	require.EqualValues(t, 2, vesting.Id)

	msg.Name = "v-pool-3"
	_, error = msgServer.CreateVestingPool(msgServerCtx, &msg)

	require.EqualValues(t, nil, error)

	accVesting, accFound = k.GetAccountVestings(testHelper.Context, addr)
	require.EqualValues(t, true, accFound)

	require.EqualValues(t, 3, len(accVesting.VestingPools))

	vesting = accVesting.VestingPools[0]
	require.EqualValues(t, 1, vesting.Id)

	vesting = accVesting.VestingPools[1]
	require.EqualValues(t, 2, vesting.Id)

	vesting = accVesting.VestingPools[2]
	require.EqualValues(t, 3, vesting.Id)
}
