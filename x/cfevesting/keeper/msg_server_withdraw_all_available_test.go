package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/stretchr/testify/require"
)

func TestWithdrawAllAvailableOnLockStart(t *testing.T) {
	vested := sdk.NewInt(1000000)

	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := vestingTestHelper.SetupAccountsVestings(ctx, accAddr.String(), 1, vested, sdk.ZeroInt())

	vestingTestHelper.WithdrawAllAvailable(ctx, accAddr, sdk.ZeroInt(), vested, sdk.ZeroInt(), vested)
	vestingTestHelper.CompareStoredAcountVestings(ctx, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestingsOnLockStart(t *testing.T) {
	vested := sdk.NewInt(1000000)
	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := vestingTestHelper.SetupAccountsVestings(ctx, accAddr.String(), 3, vested, sdk.ZeroInt())

	vestingTestHelper.WithdrawAllAvailable(ctx, accAddr, sdk.ZeroInt(), vested.MulRaw(3), sdk.ZeroInt(), vested.MulRaw(3))
	vestingTestHelper.CompareStoredAcountVestings(ctx, accAddr, accountVestings)
}

func TestWithdrawAllAvailableDuringLock(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := sdk.ZeroInt()
	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := vestingTestHelper.SetupAccountsVestings(ctx, accAddr.String(), 1, vested, sdk.ZeroInt())

	vestingTestHelper.WithdrawAllAvailable(ctx, accAddr, sdk.ZeroInt(), vested, withdrawable, vested.Sub(withdrawable))
	accountVestings.VestingPools[0].Withdrawn = withdrawable
	accountVestings.VestingPools[0].LastModificationWithdrawn = withdrawable
	vestingTestHelper.CompareStoredAcountVestings(ctx, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyLockedDuringLock(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := sdk.ZeroInt()
	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := vestingTestHelper.SetupAccountsVestings(ctx, accAddr.String(), 3, vested, sdk.ZeroInt())

	vestingTestHelper.WithdrawAllAvailable(ctx, accAddr, sdk.ZeroInt(), vested.MulRaw(3), withdrawable.MulRaw(3), vested.MulRaw(3).Sub(withdrawable.MulRaw(3)))
	for _, vesting := range accountVestings.VestingPools {
		vesting.Withdrawn = withdrawable
		vesting.LastModificationWithdrawn = withdrawable
	}
	vestingTestHelper.CompareStoredAcountVestings(ctx, accAddr, accountVestings)
}

func TestWithdrawAllAvailableAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := vested
	withdrawn := sdk.NewInt(300)

	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 110000, testutils.CreateTimeFromNumOfHours(110000))
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := vestingTestHelper.SetupAccountsVestings(ctx, accAddr.String(), 1, vested, withdrawn)

	vestingTestHelper.WithdrawAllAvailable(ctx, accAddr, sdk.ZeroInt(), vested, withdrawable.Sub(withdrawn), vested.Sub(withdrawable).Add(withdrawn))
	accountVestings.VestingPools[0].Withdrawn = withdrawable
	accountVestings.VestingPools[0].LastModificationWithdrawn = withdrawable
	vestingTestHelper.CompareStoredAcountVestings(ctx, accAddr, accountVestings)

}

func TestWithdrawAllAvailableManyVestedAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := vested
	withdrawn := sdk.NewInt(300)

	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 110000, testutils.CreateTimeFromNumOfHours(110000))
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := vestingTestHelper.SetupAccountsVestings(ctx, accAddr.String(), 3, vested, withdrawn)

	vestingTestHelper.WithdrawAllAvailable(ctx, accAddr, sdk.ZeroInt(), vested.MulRaw(3), withdrawable.MulRaw(3).Sub(withdrawn.MulRaw(3)), vested.MulRaw(3).Sub(withdrawable.MulRaw(3)).Add(withdrawn.MulRaw(3)))
	for _, vesting := range accountVestings.VestingPools {
		vesting.Withdrawn = withdrawable
		vesting.LastModificationWithdrawn = withdrawable
	}
	vestingTestHelper.CompareStoredAcountVestings(ctx, accAddr, accountVestings)
}

func TestVestAndWithdrawAllAvailable(t *testing.T) {
	vested := sdk.NewInt(1000000)
	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))
	vestingTestHelper := NewVestingTestHelper(t, testHelper)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	accAddr := acountsAddresses[0]

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, vested, accAddr)

	modifyVestingType := func(vt *types.VestingType) {
		vt.LockupPeriod = testutils.CreateDurationFromNumOfHours(9000)
		vt.VestingPeriod = testutils.CreateDurationFromNumOfHours(100000)
	}
	vestingTypes := vestingTestHelper.SetupVestingTypesWithModification(ctx, modifyVestingType, 1, 1, 1)

	vestingTestHelper.CreateVestingPool(ctx, accAddr, false, true, vPool1, 1000, *vestingTypes.VestingTypes[0], vested, vested /*0,*/, sdk.ZeroInt(), sdk.ZeroInt() /*0,*/, vested)

	vestingTestHelper.WithdrawAllAvailable(ctx, accAddr, sdk.ZeroInt(), vested, sdk.ZeroInt(), vested)

	vestingTestHelper.VerifyAccountVestingPools(ctx, accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []sdk.Int{vested}, []sdk.Int{sdk.ZeroInt()})

	oldCtx := ctx
	ctx = ctx.WithBlockHeight(int64(110000)).WithBlockTime(testutils.CreateTimeFromNumOfHours(110000))

	withdrawn := vested
	vestingTestHelper.WithdrawAllAvailable(ctx, accAddr, sdk.ZeroInt(), vested, withdrawn, vested.Sub(withdrawn))

	vestingTestHelper.VerifyAccountVestingPools(oldCtx, accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []sdk.Int{vested}, []sdk.Int{withdrawn})

}

func TestWithdrawAllAvailableBadAddress(t *testing.T) {

	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: "badaddress"}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)

	require.EqualError(t, err,
		"decoding bech32 failed: invalid separator index -1")

}
