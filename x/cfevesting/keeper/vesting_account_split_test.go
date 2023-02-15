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

type tcData struct {
	desc            string
	initialAmount   sdk.Int
	lockedBefore    sdk.Int
	blockTime       time.Time
	vAccStartTime   time.Time
	vestingDuration time.Duration
	valid           bool
	errorMassage    string
}

func TestUnlockUnbondedContinuousVestingAccountCoins5555(t *testing.T) {
	initialAmount := sdk.NewInt(8999999999999999999)
	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]

	startTime := testenv.TestEnvTime
	lockedBefore := initialAmount

	for _, tc := range []tcData{
		tcData{desc: "hour before start time", initialAmount: initialAmount, lockedBefore: lockedBefore, blockTime: startTime.Add(-time.Hour), vAccStartTime: startTime, vestingDuration: duration, valid: true, errorMassage: ""},
		tcData{desc: "on start time", initialAmount: initialAmount, lockedBefore: lockedBefore, blockTime: startTime, vAccStartTime: startTime, vestingDuration: duration, valid: true, errorMassage: ""},
		tcData{desc: "hour after start time", initialAmount: initialAmount, lockedBefore: sdk.NewInt(8990999999999999999), blockTime: startTime.Add(time.Hour), vAccStartTime: startTime, vestingDuration: duration, valid: true},
		tcData{desc: "hour before half vesting duration", initialAmount: initialAmount, lockedBefore: sdk.NewInt(4508999999999999999), blockTime: startTime.Add(duration / 2).Add(-time.Hour), vAccStartTime: startTime, vestingDuration: duration, valid: true},
		tcData{desc: "half vesting duration", initialAmount: initialAmount, lockedBefore: initialAmount.QuoRaw(2), blockTime: startTime.Add(duration / 2), vAccStartTime: startTime, vestingDuration: duration, valid: true},
		tcData{desc: "hour after half vesting duration", initialAmount: initialAmount, lockedBefore: sdk.NewInt(4491000000000000000), blockTime: startTime.Add(duration / 2).Add(time.Hour), vAccStartTime: startTime, vestingDuration: duration, valid: true},
		tcData{desc: "hour before vesting end", initialAmount: initialAmount, lockedBefore: sdk.NewInt(9000000000000000), blockTime: startTime.Add(duration - time.Hour), vAccStartTime: startTime, vestingDuration: duration, valid: true},
		
		tcData{desc: "XXXX", initialAmount: initialAmount, lockedBefore: sdk.NewInt(9000000000000000), blockTime: startTime.Add(duration), vAccStartTime: startTime, vestingDuration: duration, valid: false, errorMassage: "dffdscas"},

	} {
		t.Run(tc.desc, func(t *testing.T) {
			if tc.valid {
				testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t, accAddr, tc.initialAmount, tc.lockedBefore, tc.blockTime, tc.vAccStartTime, tc.vestingDuration)
			} else {
				require.Panics(t, func() {
					testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t, accAddr, tc.initialAmount, tc.lockedBefore, tc.blockTime, tc.vAccStartTime, tc.vestingDuration)
				})
			}

		})
	}
}

type MockTestingT struct {
	FailNowCalled bool
	t *testing.T
	require.TestingT
  }

func (m *MockTestingT) FailNow() {
	// register the method is called
	m.FailNowCalled = true
	// exit, as normal behaviour
	panic("")
  }

func testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t *testing.T, accAddr sdk.AccAddress, initialAmount sdk.Int, lockedBefore sdk.Int, blockTime time.Time, vAccStartTime time.Time, vestingDuration time.Duration) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	testHelper.SetContextBlockTime(blockTime)

	require.NoError(t, testHelper.AuthUtils.CreateDefaultDenomVestingAccount(accAddr.String(), initialAmount, vAccStartTime, vAccStartTime.Add(vestingDuration)))

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
