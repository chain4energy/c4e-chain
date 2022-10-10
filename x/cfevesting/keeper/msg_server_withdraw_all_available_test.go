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

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, testHelper.App, accAddr.String(), 1, vested, sdk.ZeroInt())

	withdrawAllAvailable(t, ctx, testHelper.App, accAddr, sdk.ZeroInt(), vested, sdk.ZeroInt(), vested)
	compareStoredAcountVestings(t, ctx, testHelper.App, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestingsOnLockStart(t *testing.T) {
	vested := sdk.NewInt(1000000)
	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, testHelper.App, accAddr.String(), 3, vested, sdk.ZeroInt())

	withdrawAllAvailable(t, ctx, testHelper.App, accAddr, sdk.ZeroInt(), vested.MulRaw(3), sdk.ZeroInt(), vested.MulRaw(3))
	compareStoredAcountVestings(t, ctx, testHelper.App, accAddr, accountVestings)
}

func TestWithdrawAllAvailableDuringLock(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := sdk.ZeroInt()
	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, testHelper.App, accAddr.String(), 1, vested, sdk.ZeroInt())

	withdrawAllAvailable(t, ctx, testHelper.App, accAddr, sdk.ZeroInt(), vested, withdrawable, vested.Sub(withdrawable))
	accountVestings.VestingPools[0].Withdrawn = withdrawable
	accountVestings.VestingPools[0].LastModificationWithdrawn = withdrawable
	compareStoredAcountVestings(t, ctx, testHelper.App, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyLockedDuringLock(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := sdk.ZeroInt()
	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, testHelper.App, accAddr.String(), 3, vested, sdk.ZeroInt())

	withdrawAllAvailable(t, ctx, testHelper.App, accAddr, sdk.ZeroInt(), vested.MulRaw(3), withdrawable.MulRaw(3), vested.MulRaw(3).Sub(withdrawable.MulRaw(3)))
	for _, vesting := range accountVestings.VestingPools {
		vesting.Withdrawn = withdrawable
		vesting.LastModificationWithdrawn = withdrawable
	}
	compareStoredAcountVestings(t, ctx, testHelper.App, accAddr, accountVestings)
}

func TestWithdrawAllAvailableAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := vested
	withdrawn := sdk.NewInt(300)

	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 110000, testutils.CreateTimeFromNumOfHours(110000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested, types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, testHelper.App, accAddr.String(), 1, vested, withdrawn)

	withdrawAllAvailable(t, ctx, testHelper.App, accAddr, sdk.ZeroInt(), vested, withdrawable.Sub(withdrawn), vested.Sub(withdrawable).Add(withdrawn))
	accountVestings.VestingPools[0].Withdrawn = withdrawable
	accountVestings.VestingPools[0].LastModificationWithdrawn = withdrawable
	compareStoredAcountVestings(t, ctx, testHelper.App, accAddr, accountVestings)

}

func TestWithdrawAllAvailableManyVestedAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	vested := sdk.NewInt(1000000)
	withdrawable := vested
	withdrawn := sdk.NewInt(300)

	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 110000, testutils.CreateTimeFromNumOfHours(110000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, vested.MulRaw(3), types.ModuleName)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, testHelper.App, accAddr.String(), 3, vested, withdrawn)

	withdrawAllAvailable(t, ctx, testHelper.App, accAddr, sdk.ZeroInt(), vested.MulRaw(3), withdrawable.MulRaw(3).Sub(withdrawn.MulRaw(3)), vested.MulRaw(3).Sub(withdrawable.MulRaw(3)).Add(withdrawn.MulRaw(3)))
	for _, vesting := range accountVestings.VestingPools {
		vesting.Withdrawn = withdrawable
		vesting.LastModificationWithdrawn = withdrawable
	}
	compareStoredAcountVestings(t, ctx, testHelper.App, accAddr, accountVestings)
}

func TestVestAndWithdrawAllAvailable(t *testing.T) {
	vested := sdk.NewInt(1000000)
	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	accAddr := acountsAddresses[0]

	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(ctx, vested, accAddr)

	modifyVestingType := func(vt *types.VestingType) {
		vt.LockupPeriod = testutils.CreateDurationFromNumOfHours(9000)
		vt.VestingPeriod = testutils.CreateDurationFromNumOfHours(100000)
	}
	vestingTypes := setupVestingTypesWithModification(ctx, testHelper.App, modifyVestingType, 1, 1, 1)

	createVestingPool(t, ctx, testHelper.App, accAddr, false, true, vPool1, 1000, *vestingTypes.VestingTypes[0], vested, vested /*0,*/, sdk.ZeroInt(), sdk.ZeroInt() /*0,*/, vested)

	withdrawAllAvailable(t, ctx, testHelper.App, accAddr, sdk.ZeroInt(), vested, sdk.ZeroInt(), vested)

	verifyAccountVestingPools(t, ctx, testHelper.App, accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []sdk.Int{vested}, []sdk.Int{sdk.ZeroInt()})

	oldCtx := ctx
	ctx = ctx.WithBlockHeight(int64(110000)).WithBlockTime(testutils.CreateTimeFromNumOfHours(110000))

	withdrawn := vested
	withdrawAllAvailable(t, ctx, testHelper.App, accAddr, sdk.ZeroInt(), vested, withdrawn, vested.Sub(withdrawn))

	verifyAccountVestingPools(t, oldCtx, testHelper.App, accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []sdk.Int{vested}, []sdk.Int{withdrawn})

}

func TestWithdrawAllAvailableBadAddress(t *testing.T) {

	testHelper, ctx := testapp.SetupTestAppWithHeightAndTime(t, 10100, testutils.CreateTimeFromNumOfHours(10100))

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(testHelper.App.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: "badaddress"}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)

	require.EqualError(t, err,
		"decoding bech32 failed: invalid separator index -1")

}
