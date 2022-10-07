package keeper_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/stretchr/testify/require"
)

func TestWithdrawAllAvailableOnLockStart(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx, _ := testapp.SetupAppWithTime(1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	addCoinsToCfevestingModule(vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), 1, vested, 0)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, 0, vested)
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyVestingsOnLockStart(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx, _ := testapp.SetupAppWithTime(1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addCoinsToCfevestingModule(3*vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), 3, vested, 0)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, 3*vested, 0, 3*vested)
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableDuringLock(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000000
	const withdrawable = 0
	app, ctx, _ := testapp.SetupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	addCoinsToCfevestingModule(vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), 1, vested, 0)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, withdrawable, vested-withdrawable)
	accountVestings.VestingPools[0].Withdrawn = sdk.NewInt(withdrawable)
	accountVestings.VestingPools[0].LastModificationWithdrawn = sdk.NewInt(withdrawable)
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableManyLockedDuringLock(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000000
	const withdrawable = 0
	app, ctx, _ := testapp.SetupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addCoinsToCfevestingModule(3*vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), 3, vested, 0)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, 3*vested, 3*withdrawable, 3*vested-3*withdrawable)
	for _, vesting := range accountVestings.VestingPools {
		vesting.Withdrawn = sdk.NewInt(withdrawable)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawable)
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestWithdrawAllAvailableAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000000
	const withdrawable = vested
	const withdrawn = 300

	app, ctx, _ := testapp.SetupAppWithTime(110000, testutils.CreateTimeFromNumOfHours(110000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	addCoinsToCfevestingModule(vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), 1, vested, withdrawn)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, withdrawable-withdrawn, vested-withdrawable+withdrawn)
	accountVestings.VestingPools[0].Withdrawn = sdk.NewInt(withdrawable)
	accountVestings.VestingPools[0].LastModificationWithdrawn = sdk.NewInt(withdrawable)
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)

}

func TestWithdrawAllAvailableManyVestedAllToWithdrawAndSomeWithdrawn(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000000
	const withdrawable = vested
	const withdrawn = 300

	app, ctx, _ := testapp.SetupAppWithTime(110000, testutils.CreateTimeFromNumOfHours(110000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addCoinsToCfevestingModule(3*vested, ctx, app)

	accAddr := acountsAddresses[0]
	accountVestings := setupAccountsVestings(ctx, app, accAddr.String(), 3, vested, withdrawn)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, 3*vested, 3*withdrawable-3*withdrawn, 3*vested-3*withdrawable+3*withdrawn)
	for _, vesting := range accountVestings.VestingPools {
		vesting.Withdrawn = sdk.NewInt(withdrawable)
		vesting.LastModificationWithdrawn = sdk.NewInt(withdrawable)
	}
	compareStoredAcountVestings(t, ctx, app, accAddr, accountVestings)
}

func TestVestAndWithdrawAllAvailable(t *testing.T) {
	commontestutils.AddHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx, _ := testapp.SetupAppWithTime(1000, testutils.CreateTimeFromNumOfHours(1000))

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	accAddr := acountsAddresses[0]
	commontestutils.AddCoinsToAccount(vested, ctx, app, accAddr)

	modifyVestingType := func(vt *types.VestingType) {
		vt.LockupPeriod = testutils.CreateDurationFromNumOfHours(9000)
		vt.VestingPeriod = testutils.CreateDurationFromNumOfHours(100000)
	}
	vestingTypes := setupVestingTypesWithModification(ctx, app, modifyVestingType, 1, 1, 1)

	createVestingPool(t, ctx, app, accAddr, false, true, vPool1, 1000, *vestingTypes.VestingTypes[0], vested, vested /*0,*/, 0, 0 /*0,*/, vested)

	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, 0, vested)

	verifyAccountVestingPools(t, ctx, app, accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []int64{vested}, []int64{0})

	oldCtx := ctx
	ctx = ctx.WithBlockHeight(int64(110000)).WithBlockTime(testutils.CreateTimeFromNumOfHours(110000))

	const withdrawn = vested
	withdrawAllAvailable(t, ctx, app, accAddr, 0, vested, withdrawn, vested-withdrawn)

	verifyAccountVestingPools(t, oldCtx, app, accAddr, []string{vPool1}, []time.Duration{1000}, []types.VestingType{*vestingTypes.VestingTypes[0]}, []int64{vested}, []int64{withdrawn})

}

func TestWithdrawAllAvailableBadAddress(t *testing.T) {

	app, ctx, _ := testapp.SetupAppWithTime(10100, testutils.CreateTimeFromNumOfHours(10100))

	msgServer, msgServerCtx := keeper.NewMsgServerImpl(app.CfevestingKeeper), sdk.WrapSDKContext(ctx)

	msg := types.MsgWithdrawAllAvailable{Creator: "badaddress"}
	_, err := msgServer.WithdrawAllAvailable(msgServerCtx, &msg)

	require.EqualError(t, err,
		"decoding bech32 failed: invalid separator index -1")

}
