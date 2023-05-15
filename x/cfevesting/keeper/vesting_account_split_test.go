package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
)

var TestDenomPrefix = "denom"

func TestUnlockUnbondedContinuousVestingAccountCoinsManyDenomError(t *testing.T) {
	denom1 := math.NewInt(8999999999999999999)
	denom2 := math.NewInt(123124)
	denom3 := math.NewInt(10)
	initialAmount := createDenomCoins([]math.Int{denom1, denom2, denom3})

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
			desc:            "all not enough to unlock - before start",
			initialAmount:   initialAmount,
			lockedBefore:    initialAmount,
			toUnlock:        createDenomCoins([]math.Int{denom1.AddRaw(1), denom2.AddRaw(1), denom3.AddRaw(1)}),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denom1,123124denom2,10denom3, to unlock: 9000000000000000000denom1,123125denom2,11denom3: insufficient funds",
		},
		{
			desc:            "first not enough to unlock - before start",
			initialAmount:   initialAmount,
			lockedBefore:    initialAmount,
			toUnlock:        createDenomCoins([]math.Int{denom1.AddRaw(1), denom2, denom3}),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denom1,123124denom2,10denom3, to unlock: 9000000000000000000denom1,123124denom2,10denom3: insufficient funds",
		},
		{
			desc:            "second not enough to unlock - before start",
			initialAmount:   initialAmount,
			lockedBefore:    initialAmount,
			toUnlock:        createDenomCoins([]math.Int{denom1, denom2.AddRaw(1), denom3}),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denom1,123124denom2,10denom3, to unlock: 8999999999999999999denom1,123125denom2,10denom3: insufficient funds",
		},
		{
			desc:            "thrid not enough to unlock - before start",
			initialAmount:   initialAmount,
			lockedBefore:    initialAmount,
			toUnlock:        createDenomCoins([]math.Int{denom1, denom2, denom3.AddRaw(1)}),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denom1,123124denom2,10denom3, to unlock: 8999999999999999999denom1,123124denom2,11denom3: insufficient funds",
		},
		{
			desc:            "unknown denom - not enough to unlock - before start",
			initialAmount:   initialAmount,
			lockedBefore:    initialAmount,
			toUnlock:        createDenomCoins([]math.Int{denom1, denom2, denom3, math.NewInt(1)}),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denom1,123124denom2,10denom3, to unlock: 8999999999999999999denom1,123124denom2,10denom3,1denom4: insufficient funds",
		},
		{
			desc:            "unknown denom only - not enough to unlock - before start",
			initialAmount:   initialAmount,
			lockedBefore:    initialAmount,
			toUnlock:        sdk.NewCoins(sdk.NewCoin("unknown", math.NewInt(1))),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denom1,123124denom2,10denom3, to unlock: 1unknown: insufficient funds",
		},
		{
			desc:            "one denom - not enough to unlock - before start",
			initialAmount:   initialAmount,
			lockedBefore:    initialAmount,
			toUnlock:        sdk.NewCoins(sdk.NewCoin("denom2", denom2.AddRaw(1))),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999denom1,123124denom2,10denom3, to unlock: 123125denom2: insufficient funds",
		},

		{
			desc:          "denom duplication - not enough to unlock - before start",
			initialAmount: initialAmount,
			lockedBefore:  initialAmount,
			toUnlock: sdk.Coins{
				sdk.NewCoin("denom2", denom2.AddRaw(1)),
				sdk.NewCoin("denom2", denom2.AddRaw(1)),
			},
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "amount to unlock validation error: duplicate denomination denom2",
		},
		{
			desc:            "second not enough to unlock -  half vesting",
			initialAmount:   initialAmount,
			lockedBefore:    createDenomCoins([]math.Int{denom1.QuoRaw(2), denom2.QuoRaw(2), denom3.QuoRaw(2)}),
			toUnlock:        createDenomCoins([]math.Int{denom1.QuoRaw(2), denom2.QuoRaw(2).AddRaw(1), denom3.QuoRaw(2)}),
			blockTime:       startTime.Add(duration / 2),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "account " + accAddr.String() + ": not enough to unlock. locked: 4499999999999999999denom1,61562denom2,5denom3, to unlock: 4499999999999999999denom1,61563denom2,5denom3: insufficient funds",
		},
		{
			desc:            "one coin is zore - before start",
			initialAmount:   initialAmount,
			lockedBefore:    initialAmount,
			toUnlock:        createDenomCoins([]math.Int{denom1, math.ZeroInt(), denom3}),
			blockTime:       startTime.Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
			expectedError:   "amount to unlock validation error: coin denom2 amount is not positive",
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

	initialAmount := createDenomCoins(createArrayOfInt(math.NewInt(8999999999999999999), 8))

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
			desc:            "hour after start time",
			initialAmount:   initialAmount,
			lockedBefore:    createDenomCoins(createArrayOfInt(math.NewInt(8990999999999999999), 8)),
			blockTime:       startTime.Add(time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:            "hour before half vesting duration",
			initialAmount:   initialAmount,
			lockedBefore:    createDenomCoins(createArrayOfInt(math.NewInt(4508999999999999999), 8)),
			blockTime:       startTime.Add(duration / 2).Add(-time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:            "half vesting duration",
			initialAmount:   initialAmount,
			lockedBefore:    createDenomCoins(createArrayOfInt(math.NewInt(4499999999999999999), 8)),
			blockTime:       startTime.Add(duration / 2),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:            "hour after half vesting duration",
			initialAmount:   initialAmount,
			lockedBefore:    createDenomCoins(createArrayOfInt(math.NewInt(4491000000000000000), 8)),
			blockTime:       startTime.Add(duration / 2).Add(time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
		{
			desc:            "hour after half vesting duration",
			initialAmount:   initialAmount,
			lockedBefore:    createDenomCoins(createArrayOfInt(math.NewInt(9000000000000000), 8)),
			blockTime:       startTime.Add(duration - time.Hour),
			vAccStartTime:   startTime,
			vestingDuration: duration,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			toUnlock := sdk.NewCoins(sdk.NewCoin(tc.lockedBefore[0].Denom, math.NewInt(1)))
			toUnlock = toUnlock.Add(sdk.NewCoin(tc.lockedBefore[1].Denom, math.NewInt(300)))
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
	initialAmount := math.NewInt(8999999999999999999)
	duration := 1000 * time.Hour

	acountsAddresses, valAddrs := testcosmos.CreateAccounts(2, 1)
	accAddr := acountsAddresses[0]

	delegationAmount := initialAmount.QuoRaw(4)

	startTime := testenv.TestEnvTime
	type AccType int
	const (
		Vesting AccType = iota
		Base
		None
	)

	for _, tc := range []struct {
		desc            string
		initialAmount   math.Int
		lockedBefore    math.Int
		toUnlock        math.Int
		blockTime       time.Time
		vAccStartTime   time.Time
		vestingDuration time.Duration
		expectedError   string
		accountType     AccType
		delegation      bool
	}{
		{desc: "before vesting start - not enought to unlock", initialAmount: initialAmount, lockedBefore: initialAmount, blockTime: startTime.Add(-time.Hour), toUnlock: initialAmount.AddRaw(1),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999uc4e, to unlock: 9000000000000000000uc4e: insufficient funds"},
		{desc: "on vesting start - not enought to unlock", initialAmount: initialAmount, lockedBefore: initialAmount, blockTime: startTime, toUnlock: initialAmount.AddRaw(1),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + ": not enough to unlock. locked: 8999999999999999999uc4e, to unlock: 9000000000000000000uc4e: insufficient funds"},
		{desc: "on half vesting - not enought to unlock", initialAmount: initialAmount, lockedBefore: initialAmount.QuoRaw(2), blockTime: startTime.Add(duration / 2), toUnlock: initialAmount.QuoRaw(2).AddRaw(1),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + ": not enough to unlock. locked: 4499999999999999999uc4e, to unlock: 4500000000000000000uc4e: insufficient funds"},
		{desc: "on vesting end - not enought to unlock", initialAmount: initialAmount, lockedBefore: math.ZeroInt(), blockTime: startTime.Add(duration), toUnlock: math.NewInt(1),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + ": not enough to unlock. locked: , to unlock: 1uc4e: insufficient funds"},
		{desc: "no account", initialAmount: math.ZeroInt(), lockedBefore: math.ZeroInt(), blockTime: startTime, toUnlock: math.NewInt(1),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + " doesn't exist: entity does not exist", accountType: None},
		{desc: "wrong account type", initialAmount: initialAmount, lockedBefore: math.ZeroInt(), blockTime: startTime.Add(duration), toUnlock: math.NewInt(1),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + " is not ContinuousVestingAccount: invalid type", accountType: Base},
		{desc: "on half vesting - not enought to unlock with delegation", initialAmount: initialAmount, lockedBefore: initialAmount.QuoRaw(2), blockTime: startTime.Add(duration / 2), toUnlock: initialAmount.QuoRaw(4).AddRaw(2),
			vAccStartTime: startTime, vestingDuration: duration, expectedError: "account " + accAddr.String() + ": not enough to unlock. locked: 2250000000000000000uc4e, to unlock: 2250000000000000001uc4e: insufficient funds", delegation: true},
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
			if tc.delegation {
				testHelper.StakingUtils.SetupValidators(valAddrs, math.NewInt(1))
				testHelper.StakingUtils.MessageDelegate(2, 0, valAddrs[0], accAddr, delegationAmount)
				testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoinsError(accAddr, tc.toUnlock, tc.initialAmount.Sub(delegationAmount), tc.lockedBefore.Sub(delegationAmount), tc.expectedError)

			} else {
				testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoinsError(accAddr, tc.toUnlock, tc.initialAmount, tc.lockedBefore, tc.expectedError)
			}
		})
	}
}

func TestUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t *testing.T) {
	testUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t, false)
}

func TestUnlockUnbondedContinuousVestingAccountCoinsSingleDenomWithDelegations(t *testing.T) {
	testUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t, true)
}

func testUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t *testing.T, useDelegations bool) {
	initialAmount := math.NewInt(8999999999999999999)
	duration := 1000 * time.Hour

	acountsAddresses, valAddrs := testcosmos.CreateAccounts(2, 1)
	accAddr := acountsAddresses[0]

	startTime := testenv.TestEnvTime
	lockedBefore := initialAmount

	for _, tc := range []struct {
		desc            string
		initialAmount   math.Int
		lockedBefore    math.Int
		blockTime       time.Time
		vAccStartTime   time.Time
		vestingDuration time.Duration
	}{
		{desc: "hour before start time", initialAmount: initialAmount, lockedBefore: lockedBefore, blockTime: startTime.Add(-time.Hour), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "on start time", initialAmount: initialAmount, lockedBefore: lockedBefore, blockTime: startTime, vAccStartTime: startTime, vestingDuration: duration},
		{desc: "hour after start time", initialAmount: initialAmount, lockedBefore: math.NewInt(8990999999999999999), blockTime: startTime.Add(time.Hour), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "hour before half vesting duration", initialAmount: initialAmount, lockedBefore: math.NewInt(4508999999999999999), blockTime: startTime.Add(duration / 2).Add(-time.Hour), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "half vesting duration", initialAmount: initialAmount, lockedBefore: initialAmount.QuoRaw(2), blockTime: startTime.Add(duration / 2), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "hour after half vesting duration", initialAmount: initialAmount, lockedBefore: math.NewInt(4491000000000000000), blockTime: startTime.Add(duration / 2).Add(time.Hour), vAccStartTime: startTime, vestingDuration: duration},
		{desc: "hour before vesting end", initialAmount: initialAmount, lockedBefore: math.NewInt(9000000000000000), blockTime: startTime.Add(duration - time.Hour), vAccStartTime: startTime, vestingDuration: duration},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			testSingleTimeUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t, accAddr, tc.initialAmount, tc.lockedBefore, tc.blockTime, tc.vAccStartTime, tc.vestingDuration, useDelegations, valAddrs)

		})
	}
}

func testSingleTimeUnlockUnbondedContinuousVestingAccountCoinsSingleDenom(t require.TestingT, accToCreateAddr sdk.AccAddress,
	initialAmount math.Int, lockedBefore math.Int, blockTime time.Time,
	vAccStartTime time.Time, vestingDuration time.Duration, useDelegation bool, valAddrs []sdk.ValAddress) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	testHelper.SetContextBlockTime(blockTime)

	delegationAmount := math.NewInt(100000)
	restoreAmount := initialAmount
	if useDelegation {
		testHelper.StakingUtils.SetupValidators(valAddrs, math.NewInt(1))
	}
	require.NoError(t, testHelper.AuthUtils.CreateDefaultDenomVestingAccount(accToCreateAddr.String(), initialAmount, vAccStartTime, vAccStartTime.Add(vestingDuration)))
	if useDelegation {
		testHelper.StakingUtils.MessageDelegate(2, 0, valAddrs[0], accToCreateAddr, delegationAmount)
		initialAmount = initialAmount.Sub(delegationAmount)
		lockedBefore = lockedBefore.Sub(delegationAmount)
	}

	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, math.NewInt(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, restoreAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, math.NewInt(300), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, restoreAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.QuoRaw(2).SubRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, restoreAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.QuoRaw(2), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, restoreAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.QuoRaw(2).AddRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, restoreAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.SubRaw(300), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, restoreAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore.SubRaw(1), initialAmount, lockedBefore)

	require.NoError(t, testHelper.AuthUtils.ModifyVestingAccountOriginalVesting(accToCreateAddr.String(), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, restoreAmount))))
	testHelper.C4eVestingUtils.UnlockUnbondedDefaultDenomContinuousVestingAccountCoins(accToCreateAddr, lockedBefore, initialAmount, lockedBefore)
}

func testSingleTimeUnlockUnbondedContinuousVestingAccountCoins(t require.TestingT, accToCreateAddr sdk.AccAddress, initialAmount sdk.Coins, lockedBefore sdk.Coins, amountToUnlock sdk.Coins, blockTime time.Time, vAccStartTime time.Time, vestingDuration time.Duration) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	testHelper.SetContextBlockTime(blockTime)

	require.NoError(t, testHelper.AuthUtils.CreateVestingAccount(accToCreateAddr.String(), initialAmount, vAccStartTime, vAccStartTime.Add(vestingDuration)))

	testHelper.C4eVestingUtils.UnlockUnbondedContinuousVestingAccountCoins(accToCreateAddr, amountToUnlock, initialAmount, lockedBefore)
}

func createDenomCoins(amounts []math.Int) sdk.Coins {
	result := sdk.Coins{}
	for i, amount := range amounts {
		coin := sdk.Coin{
			Denom:  fmt.Sprintf("%s%d", TestDenomPrefix, i+1),
			Amount: amount,
		}
		result = append(result, coin)

		// result = result.Add(sdk.NewCoin(fmt.Sprintf("denom%d", i+1), amount))
	}
	return result
}

func createArrayOfInt(amount math.Int, count int) []math.Int {
	result := []math.Int{}
	for i := 0; i < count; i++ {
		result = append(result, amount)
	}
	return result
}
