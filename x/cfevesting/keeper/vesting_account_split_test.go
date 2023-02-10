package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
)

func TestUnlockUnbondedContinuousVestingAccountCoins(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	initialAmount := sdk.NewInt(8999999999999999999)
	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]

	startTime := testHelper.Context.BlockTime()
	lockedBefore := initialAmount

	require.NoError(t, testHelper.AuthUtils.CreateDefaultDenomVestingAccount(accAddr.String(), initialAmount, startTime, startTime.Add(duration)))

	// ----- On start time
	testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t, testHelper, accAddr, initialAmount, lockedBefore)

	// ----- Hour after start time

	testHelper.SetContextBlockTime(startTime.Add(time.Hour))
	lockedBefore = sdk.NewInt(8990999999999999999)
	testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t, testHelper, accAddr, initialAmount, lockedBefore)

	// ----- Half vesting duration

	testHelper.SetContextBlockTime(startTime.Add(duration / 2))
	lockedBefore = initialAmount.QuoRaw(2)
	testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t, testHelper, accAddr, initialAmount, lockedBefore)

	// ----- Hour before vesting end

	testHelper.SetContextBlockTime(startTime.Add(duration - time.Hour))
	lockedBefore = sdk.NewInt(9000000000000000)
	testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t, testHelper, accAddr, initialAmount, lockedBefore)
}

func testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t *testing.T, testHelper *testapp.TestHelper, accAddr sdk.AccAddress, initialAmount sdk.Int, lockedBefore sdk.Int) {
	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accAddr, sdk.NewInt(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accAddr, sdk.NewInt(300), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accAddr, lockedBefore.QuoRaw(2).SubRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accAddr, lockedBefore.QuoRaw(2), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accAddr, lockedBefore.QuoRaw(2).AddRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accAddr, lockedBefore.SubRaw(300), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accAddr, lockedBefore.SubRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accAddr, lockedBefore, initialAmount, lockedBefore)
}
