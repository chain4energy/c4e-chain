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

func TestUnlockUnbondedContinuousVestingAccountCoinsManyDenomError(t *testing.T) {
	denom1 := sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999))
	denom2 := sdk.NewCoin("denon2", sdk.NewInt(123124))
	denom3 := sdk.NewCoin("denon3", sdk.NewInt(10))

	initialAmount := sdk.NewCoins(denom1, denom2, denom3)

	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]

	startTime := testenv.TestEnvTime

	for _, tc := range []struct {
		desc            string
		initialAmount   sdk.Coins
		lockedBefore    sdk.Coins
		toUnlock        sdk.Coins
		blockTime       time.Time
		vAccStartTime   time.Time
		vestingDuration time.Duration
		expectedError   string
	}{
		{
			desc:          "all not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999+1)),
				sdk.NewCoin("denon2", sdk.NewInt(123124+1)),
				sdk.NewCoin("denon3", sdk.NewInt(10+1))),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denon1,123124denon2,10denon3, to unlock: 9000000000000000000denon1,123125denon2,11denon3: entity not exists",
		},
		{
			desc:          "first not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999+1)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denon1,123124denon2,10denon3, to unlock: 9000000000000000000denon1,123124denon2,10denon3: entity not exists",
		},
		{
			desc:          "second not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124+1)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denon1,123124denon2,10denon3, to unlock: 8999999999999999999denon1,123125denon2,10denon3: entity not exists",
		},
		{
			desc:          "thrid not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10+1))),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denon1,123124denon2,10denon3, to unlock: 8999999999999999999denon1,123124denon2,11denon3: entity not exists",
		},
		{
			desc:          "thrid not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10+1))),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denon1,123124denon2,10denon3, to unlock: 8999999999999999999denon1,123124denon2,11denon3: entity not exists",
		},
		{
			desc:          "unknown denom - not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10)),
				sdk.NewCoin("unknown", sdk.NewInt(1)),
			),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denon1,123124denon2,10denon3, to unlock: 8999999999999999999denon1,123124denon2,10denon3,1unknown: entity not exists",
		},
		{
			desc:          "unknown denom only - not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.NewCoins(
				sdk.NewCoin("unknown", sdk.NewInt(1)),
			),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denon1,123124denon2,10denon3, to unlock: 1unknown: entity not exists",
		},
		{
			desc:          "one denom - not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.NewCoins(
				sdk.NewCoin("denon2", sdk.NewInt(123124+1)),
			),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denon1,123124denon2,10denon3, to unlock: 123125denon2: entity not exists",
		},
		{
			desc:          "one denom - not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(123124)),
				sdk.NewCoin("denon3", sdk.NewInt(10))),
			toUnlock: sdk.Coins{
				sdk.NewCoin("denon2", sdk.NewInt(123124+1)),
				sdk.NewCoin("denon2", sdk.NewInt(123124+1)),
			},
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "amount to unlock validation error: duplicate denomination denon2",
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			testHelper := testapp.SetupTestAppWithHeight(t, 1000)
			testHelper.SetContextBlockTime(tc.blockTime)

			require.NoError(t, testHelper.AuthUtils.CreateVestingAccount(accAddr.String(), tc.initialAmount, tc.vAccStartTime, tc.vAccStartTime.Add(tc.vestingDuration)))
			testHelper.C4eVestingUtils.UnlockUnbondedContinuousVestingAccountCoinsError(accAddr, tc.toUnlock, tc.initialAmount, tc.lockedBefore, tc.expectedError)
		})
	}
}

func TestUnlockUnbondedContinuousVestingAccountCoinsManyDenom(t *testing.T) {
	denom1 := sdk.NewCoin("denon1", sdk.NewInt(8999999999999999999))
	denom2 := sdk.NewCoin("denon2", sdk.NewInt(8999999999999999999))
	denom3 := sdk.NewCoin("denon3", sdk.NewInt(8999999999999999999))
	denom4 := sdk.NewCoin("denon4", sdk.NewInt(8999999999999999999))
	denom5 := sdk.NewCoin("denon5", sdk.NewInt(8999999999999999999))
	denom6 := sdk.NewCoin("denon6", sdk.NewInt(8999999999999999999))
	denom7 := sdk.NewCoin("denon7", sdk.NewInt(8999999999999999999))
	denom8 := sdk.NewCoin("denon8", sdk.NewInt(8999999999999999999))

	initialAmount := sdk.NewCoins(denom1, denom2, denom3, denom4, denom5, denom6, denom7, denom8)

	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]

	startTime := testenv.TestEnvTime
	lockedBefore := initialAmount

	for _, tc := range []struct {
		desc            string
		initialAmount   sdk.Coins
		lockedBefore    sdk.Coins
		blockTime       time.Time
		vAccStartTime   time.Time
		vestingDuration time.Duration
	}{
		{
			desc:            "hour before start time",
			initialAmount:   initialAmount,
			lockedBefore:    lockedBefore,
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:            "on start time",
			initialAmount:   initialAmount,
			lockedBefore:    lockedBefore,
			blockTime:       startTime,
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:          "hour after start time",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(8990999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(8990999999999999999)),
				sdk.NewCoin("denon3", sdk.NewInt(8990999999999999999)),
				sdk.NewCoin("denon4", sdk.NewInt(8990999999999999999)),
				sdk.NewCoin("denon5", sdk.NewInt(8990999999999999999)),
				sdk.NewCoin("denon6", sdk.NewInt(8990999999999999999)),
				sdk.NewCoin("denon7", sdk.NewInt(8990999999999999999)),
				sdk.NewCoin("denon8", sdk.NewInt(8990999999999999999))),
			blockTime:       startTime.Add(time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:          "hour before half vesting duration",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon3", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon4", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon5", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon6", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon7", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon8", sdk.NewInt(4508999999999999999))),
			blockTime:       startTime.Add(duration / 2).Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:          "half vesting duration",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon3", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon4", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon5", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon6", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon7", sdk.NewInt(4508999999999999999)),
				sdk.NewCoin("denon8", sdk.NewInt(4508999999999999999))),
			blockTime:       startTime.Add(duration / 2).Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:          "half vesting duration",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(4499999999999999999)),
				sdk.NewCoin("denon2", sdk.NewInt(4499999999999999999)),
				sdk.NewCoin("denon3", sdk.NewInt(4499999999999999999)),
				sdk.NewCoin("denon4", sdk.NewInt(4499999999999999999)),
				sdk.NewCoin("denon5", sdk.NewInt(4499999999999999999)),
				sdk.NewCoin("denon6", sdk.NewInt(4499999999999999999)),
				sdk.NewCoin("denon7", sdk.NewInt(4499999999999999999)),
				sdk.NewCoin("denon8", sdk.NewInt(4499999999999999999))),
			blockTime:       startTime.Add(duration / 2),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:          "hour after half vesting duration",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(4491000000000000000)),
				sdk.NewCoin("denon2", sdk.NewInt(4491000000000000000)),
				sdk.NewCoin("denon3", sdk.NewInt(4491000000000000000)),
				sdk.NewCoin("denon4", sdk.NewInt(4491000000000000000)),
				sdk.NewCoin("denon5", sdk.NewInt(4491000000000000000)),
				sdk.NewCoin("denon6", sdk.NewInt(4491000000000000000)),
				sdk.NewCoin("denon7", sdk.NewInt(4491000000000000000)),
				sdk.NewCoin("denon8", sdk.NewInt(4491000000000000000))),
			blockTime:       startTime.Add(duration / 2).Add(time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:          "hour after half vesting duration",
			initialAmount: initialAmount,
			lockedBefore: sdk.NewCoins(sdk.NewCoin("denon1", sdk.NewInt(9000000000000000)),
				sdk.NewCoin("denon2", sdk.NewInt(9000000000000000)),
				sdk.NewCoin("denon3", sdk.NewInt(9000000000000000)),
				sdk.NewCoin("denon4", sdk.NewInt(9000000000000000)),
				sdk.NewCoin("denon5", sdk.NewInt(9000000000000000)),
				sdk.NewCoin("denon6", sdk.NewInt(9000000000000000)),
				sdk.NewCoin("denon7", sdk.NewInt(9000000000000000)),
				sdk.NewCoin("denon8", sdk.NewInt(9000000000000000))),
			blockTime:       startTime.Add(duration - time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			toUnlock := sdk.NewCoins(sdk.NewCoin(tc.lockedBefore[0].Denom, sdk.NewInt(1)))
			toUnlock = toUnlock.Add(sdk.NewCoin(tc.lockedBefore[1].Denom, sdk.NewInt(300)))
			toUnlock = toUnlock.Add(sdk.NewCoin(tc.lockedBefore[2].Denom, tc.lockedBefore[2].Amount.QuoRaw(2).SubRaw(1)))
			toUnlock = toUnlock.Add(sdk.NewCoin(tc.lockedBefore[3].Denom, tc.lockedBefore[3].Amount.QuoRaw(2)))
			toUnlock = toUnlock.Add(sdk.NewCoin(tc.lockedBefore[4].Denom, tc.lockedBefore[4].Amount.QuoRaw(2).AddRaw(1)))
			toUnlock = toUnlock.Add(sdk.NewCoin(tc.lockedBefore[5].Denom, tc.lockedBefore[5].Amount.SubRaw(300)))
			toUnlock = toUnlock.Add(sdk.NewCoin(tc.lockedBefore[6].Denom, tc.lockedBefore[6].Amount.SubRaw(1)))
			toUnlock = toUnlock.Add(sdk.NewCoin(tc.lockedBefore[7].Denom, tc.lockedBefore[7].Amount))
			testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t, accAddr, tc.initialAmount, tc.lockedBefore, toUnlock, tc.blockTime, tc.vAccStartTime, tc.vestingDuration)

		})
	}
}

func TestUnlockUnbondedContinuousVestingAccountCoinsSingleDenomError(t *testing.T) {
	initialAmount := sdk.NewInt(8999999999999999999)
	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]

	startTime := testenv.TestEnvTime
	type AccType int
	const (
		Vesting AccType = iota
		Base
		None
	)

	for _, tc := range []struct {
		desc            string
		initialAmount   sdk.Int
		lockedBefore    sdk.Int
		blockTime       time.Time
		vAccStartTime   time.Time
		vestingDuration time.Duration
		expectedError   string
		accountType     AccType
	}{
		{desc: "on vesting end - not enought to unlock", initialAmount: initialAmount, lockedBefore: sdk.ZeroInt(), blockTime: startTime.Add(duration),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + ": not enough to unlock. locked: , to unlock: 1uc4e: entity not exists"},
		{desc: "no account", initialAmount: sdk.ZeroInt(), lockedBefore: sdk.ZeroInt(), blockTime: startTime,
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + " doesn't exist: entity not exists", accountType: None},
		{desc: "wrong account type", initialAmount: initialAmount, lockedBefore: sdk.ZeroInt(), blockTime: startTime.Add(duration),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + " is not ContinuousVestingAccount: entity not exists", accountType: Base},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			testHelper := testapp.SetupTestAppWithHeight(t, 1000)
			testHelper.SetContextBlockTime(tc.blockTime)

			switch tc.accountType {
			case Vesting:
				require.NoError(t, testHelper.AuthUtils.CreateDefaultDenomVestingAccount(accAddr.String(), tc.initialAmount, tc.vAccStartTime, tc.vAccStartTime.Add(tc.vestingDuration)))
			case Base:
				require.NoError(t, testHelper.AuthUtils.CreateDefaultDenomBaseAccount(accAddr.String(), tc.initialAmount))
			}
			testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoinsError(accAddr, sdk.NewInt(1), tc.initialAmount, tc.lockedBefore, tc.expectedError)

		})
	}
}

func TestUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t *testing.T) {
	initialAmount := sdk.NewInt(8999999999999999999)
	duration := 1000 * time.Hour

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	accAddr := acountsAddresses[0]

	startTime := testenv.TestEnvTime
	lockedBefore := initialAmount

	for _, tc := range []struct {
		desc            string
		initialAmount   sdk.Int
		lockedBefore    sdk.Int
		blockTime       time.Time
		vAccStartTime   time.Time
		vestingDuration time.Duration
	}{
		{desc: "hour before start time", initialAmount: initialAmount, lockedBefore: lockedBefore, blockTime: startTime.Add(-time.Hour), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "on start time", initialAmount: initialAmount, lockedBefore: lockedBefore, blockTime: startTime, vAccStartTime: startTime, vestingDuration: duration},
		{desc: "hour after start time", initialAmount: initialAmount, lockedBefore: sdk.NewInt(8990999999999999999), blockTime: startTime.Add(time.Hour), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "hour before half vesting duration", initialAmount: initialAmount, lockedBefore: sdk.NewInt(4508999999999999999), blockTime: startTime.Add(duration / 2).Add(-time.Hour), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "half vesting duration", initialAmount: initialAmount, lockedBefore: initialAmount.QuoRaw(2), blockTime: startTime.Add(duration / 2), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "hour after half vesting duration", initialAmount: initialAmount, lockedBefore: sdk.NewInt(4491000000000000000), blockTime: startTime.Add(duration / 2).Add(time.Hour), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "hour before vesting end", initialAmount: initialAmount, lockedBefore: sdk.NewInt(9000000000000000), blockTime: startTime.Add(duration - time.Hour), vAccStartTime: startTime, vestingDuration: duration},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			testSingleTimeUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t, accAddr, tc.initialAmount, tc.lockedBefore, tc.blockTime, tc.vAccStartTime, tc.vestingDuration)

		})
	}
}

func testSingleTimeUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t require.TestingT, accToCreateAddr sdk.AccAddress, initialAmount sdk.Int, lockedBefore sdk.Int, blockTime time.Time, vAccStartTime time.Time, vestingDuration time.Duration) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	testHelper.SetContextBlockTime(blockTime)

	require.NoError(t, testHelper.AuthUtils.CreateDefaultDenomVestingAccount(accToCreateAddr.String(), initialAmount, vAccStartTime, vAccStartTime.Add(vestingDuration)))

	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, sdk.NewInt(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, sdk.NewInt(300), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.QuoRaw(2).SubRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.QuoRaw(2), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.QuoRaw(2).AddRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.SubRaw(300), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.SubRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, initialAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore, initialAmount, lockedBefore)
}

func testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t require.TestingT, accToCreateAddr sdk.AccAddress, initialAmount sdk.Coins, lockedBefore sdk.Coins, amountToUnlock sdk.Coins, blockTime time.Time, vAccStartTime time.Time, vestingDuration time.Duration) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	testHelper.SetContextBlockTime(blockTime)

	require.NoError(t, testHelper.AuthUtils.CreateVestingAccount(accToCreateAddr.String(), initialAmount, vAccStartTime, vAccStartTime.Add(vestingDuration)))

	testHelper.C4eVestingUtils.UnlockUnbondedContinuousVestingAccountCoins(accToCreateAddr, amountToUnlock, initialAmount, lockedBefore)
}
