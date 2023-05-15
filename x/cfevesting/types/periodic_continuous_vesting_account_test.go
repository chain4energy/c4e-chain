package types_test

import (
	"cosmossdk.io/math"
	"fmt"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"

	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func TestSinglePeriod(t *testing.T) {
	startTime := time.Now()
	endTime := time.Now().Add(100 * 100 * time.Hour)
	periods := types.ContinuousVestingPeriods{
		{
			Amount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000000))),
			StartTime: startTime.Unix(),
			EndTime:   endTime.Unix(),
		},
	}
	sum := periodsSum(&periods)
	acc := createClaimAccout(sum, periods[0].StartTime, periods[0].EndTime, periods)

	for i := -10; i <= 150; i++ {
		var checkTime time.Time
		var vested sdk.Coins
		checkTime = startTime.Add(time.Duration(i*100) * time.Hour)
		if i <= 0 {
			vested = sdk.NewCoins()
		} else if i >= 100 {
			vested = sdk.NewCoins(periods[0].Amount...)
		} else {
			vested = sdk.NewCoins(sdk.NewCoin(periods[0].Amount[0].Denom, periods[0].Amount[0].Amount))
			vested[0].Amount = vested[0].Amount.MulRaw(int64(i)).QuoRaw(100)
		}
		require.True(t, acc.GetVestedCoins(checkTime).IsEqual(vested))
		require.True(t, acc.GetVestingCoins(checkTime).IsEqual(acc.OriginalVesting.Sub(vested...)))
		require.True(t, acc.LockedCoins(checkTime).IsEqual(acc.OriginalVesting.Sub(vested...)))
		require.True(t, acc.DelegatedFree.IsZero())
		require.True(t, acc.DelegatedVesting.IsZero())
	}

	checkTime := startTime.Add(time.Duration(30*100) * time.Hour)
	vested := sdk.NewCoins(sdk.NewCoin(periods[0].Amount[0].Denom, periods[0].Amount[0].Amount))
	vested[0].Amount = vested[0].Amount.MulRaw(int64(30)).QuoRaw(100)
	acc.TrackDelegation(checkTime, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000000))), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(5000000))))

	require.True(t, acc.GetVestedCoins(checkTime).IsEqual(vested))
	require.True(t, acc.GetVestingCoins(checkTime).IsEqual(acc.OriginalVesting.Sub(vested...)))
	require.True(t, acc.LockedCoins(checkTime).IsEqual(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(2000000)))))
	require.True(t, acc.DelegatedVesting.IsEqual(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(5000000)))))
	require.True(t, acc.DelegatedFree.IsZero())

	checkTime = startTime.Add(time.Duration(40*100) * time.Hour)
	vested = sdk.NewCoins(sdk.NewCoin(periods[0].Amount[0].Denom, periods[0].Amount[0].Amount))
	vested[0].Amount = vested[0].Amount.MulRaw(int64(40)).QuoRaw(100)
	acc.TrackDelegation(checkTime, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(5000000))), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(3000000))))

	require.True(t, acc.GetVestedCoins(checkTime).IsEqual(vested))
	require.True(t, acc.GetVestingCoins(checkTime).IsEqual(acc.OriginalVesting.Sub(vested...)))
	require.True(t, acc.LockedCoins(checkTime).IsZero())
	require.True(t, acc.DelegatedVesting.IsEqual(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(6000000)))))
	require.True(t, acc.DelegatedFree.IsEqual(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(2000000)))))

	checkTime = startTime.Add(time.Duration(50*100) * time.Hour)
	vested = sdk.NewCoins(sdk.NewCoin(periods[0].Amount[0].Denom, periods[0].Amount[0].Amount))
	vested[0].Amount = vested[0].Amount.MulRaw(int64(50)).QuoRaw(100)
	acc.TrackDelegation(checkTime, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(2000000))), sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(2000000))))
	require.True(t, acc.GetVestedCoins(checkTime).IsEqual(vested))
	require.True(t, acc.GetVestingCoins(checkTime).IsEqual(acc.OriginalVesting.Sub(vested...)))
	require.True(t, acc.LockedCoins(checkTime).IsZero())
	require.True(t, acc.DelegatedVesting.IsEqual(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(6000000)))))
	require.True(t, acc.DelegatedFree.IsEqual(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(4000000)))))

}

func TestMultiplePeriods(t *testing.T) {
	shift := 50
	timeShift := time.Duration(shift*100) * time.Hour
	p1StartTime := time.Now()
	p1EndTime := time.Now().Add(100 * 100 * time.Hour)
	p2StartTime := p1StartTime.Add(timeShift)
	p2EndTime := p1EndTime.Add(timeShift)
	p3StartTime := p1StartTime.Add(2 * timeShift)
	p3EndTime := p1EndTime.Add(2 * timeShift)
	periods := types.ContinuousVestingPeriods{
		{
			Amount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000000))),
			StartTime: p1StartTime.Unix(),
			EndTime:   p1EndTime.Unix(),
		},
		{
			Amount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(100000000))),
			StartTime: p2StartTime.Unix(),
			EndTime:   p2EndTime.Unix(),
		},
		{
			Amount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1000000000))),
			StartTime: p3StartTime.Unix(),
			EndTime:   p3EndTime.Unix(),
		},
	}
	sum := periodsSum(&periods)
	acc := createClaimAccout(sum, periods[0].StartTime, periods[0].EndTime, periods)

	for i := -10; i <= 150; i++ {
		var checkTime time.Time
		var vested sdk.Coins
		checkTime = p1StartTime.Add(time.Duration(i*100) * time.Hour)
		if i >= 100 {
			vested = vested.Add(periods[0].Amount...)
		} else if i > 0 {
			periodVested := sdk.NewCoins(sdk.NewCoin(periods[0].Amount[0].Denom, periods[0].Amount[0].Amount))
			periodVested[0].Amount = periodVested[0].Amount.MulRaw(int64(i)).QuoRaw(100)
			vested = vested.Add(periodVested...)
		}

		if i >= 100+shift {
			vested = vested.Add(periods[1].Amount...)
		} else if i > shift {
			periodVested := sdk.NewCoins(sdk.NewCoin(periods[1].Amount[0].Denom, periods[1].Amount[0].Amount))
			periodVested[0].Amount = periodVested[0].Amount.MulRaw(int64(i - shift)).QuoRaw(100)
			vested = vested.Add(periodVested...)
		}

		if i >= 100+2*shift {
			vested = vested.Add(periods[2].Amount...)
		} else if i > 2*shift {
			periodVested := sdk.NewCoins(sdk.NewCoin(periods[2].Amount[0].Denom, periods[2].Amount[0].Amount))
			periodVested[0].Amount = periodVested[0].Amount.MulRaw(int64(i - 2*shift)).QuoRaw(100)
			vested = vested.Add(periodVested...)
		}

		require.True(t, acc.GetVestedCoins(checkTime).IsEqual(vested))
		require.True(t, acc.GetVestingCoins(checkTime).IsEqual(acc.OriginalVesting.Sub(vested...)))
		require.True(t, acc.LockedCoins(checkTime).IsEqual(acc.OriginalVesting.Sub(vested...)))
		require.True(t, acc.DelegatedFree.IsZero())
		require.True(t, acc.DelegatedVesting.IsZero())
	}
}

func TestValidateClaimAccount(t *testing.T) {
	for _, tc := range []ClaimAccountTc{
		correctClaimVestingAccount(),
		wrongOriginalVestingClaimVestingAccount(),
		wrongStartTimeClaimVestingAccount(),
		wrongEndTimeClaimVestingAccount(),
		endLessThanStartClaimVestingAccount(),
		periodEndLessThanStartClaimVestingAccount(),
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.account.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.message)
			}
		})
	}
}

func correctClaimVestingAccount() ClaimAccountTc {
	return ClaimAccountTc{
		desc:    "valid account",
		account: createCorrectClaimAccout(),
		valid:   true,
	}
}

func wrongOriginalVestingClaimVestingAccount() ClaimAccountTc {
	acc := createCorrectClaimAccout()
	acc.OriginalVesting[0].Amount = acc.OriginalVesting[0].Amount.Add(math.NewInt(300))
	return ClaimAccountTc{
		desc:    "wrong original vesting claim account",
		account: acc,
		valid:   false,
		message: "original vesting (922100uc4e) not equal to sum of periods (921800uc4e)",
	}
}

func wrongStartTimeClaimVestingAccount() ClaimAccountTc {
	acc := createCorrectClaimAccout()
	acc.StartTime = acc.StartTime - 100
	return ClaimAccountTc{
		desc:    "wrong start time claim account",
		account: acc,
		valid:   false,
		message: fmt.Sprintf("vesting start-time (%d) not eqaul to earliest period start time (%d)", acc.StartTime, acc.VestingPeriods[1].StartTime),
	}
}

func wrongEndTimeClaimVestingAccount() ClaimAccountTc {
	acc := createCorrectClaimAccout()
	acc.EndTime = acc.EndTime - 100
	return ClaimAccountTc{
		desc:    "wrong end time claim account",
		account: acc,
		valid:   false,
		message: fmt.Sprintf("vesting end-time (%d) not eqaul to lastest period end time (%d)", acc.EndTime, acc.VestingPeriods[2].EndTime),
	}
}

func endLessThanStartClaimVestingAccount() ClaimAccountTc {
	acc := createCorrectClaimAccout()
	acc.EndTime = acc.StartTime - 100
	return ClaimAccountTc{
		desc:    "wrong end time claim account",
		account: acc,
		valid:   false,
		message: fmt.Sprintf("vesting end-time (%d) cannot be before start-time (%d)", acc.EndTime, acc.StartTime),
	}
}

func periodEndLessThanStartClaimVestingAccount() ClaimAccountTc {
	acc := createCorrectClaimAccout()

	acc.VestingPeriods[3].EndTime = acc.VestingPeriods[3].StartTime - 100
	return ClaimAccountTc{
		desc:    "wrong end time claim account",
		account: acc,
		valid:   false,
		message: fmt.Sprintf("vesting period end-time (%d) cannot be before start-time (%d)", acc.VestingPeriods[3].EndTime, acc.VestingPeriods[3].StartTime),
	}
}

func createCorrectClaimAccout() *types.PeriodicContinuousVestingAccount {
	periods := types.ContinuousVestingPeriods{
		{
			Amount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1000))),
			StartTime: time.Now().Add(-24 * 100 * time.Hour).Unix(),
			EndTime:   time.Now().Add(24 * 100 * time.Hour).Unix(),
		},
		{
			Amount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(800))),
			StartTime: time.Now().Add(-24 * 300 * time.Hour).Unix(),
			EndTime:   time.Now().Add(24 * 150 * time.Hour).Unix(),
		},
		{
			Amount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(900000))),
			StartTime: time.Now().Add(-24 * 32 * time.Hour).Unix(),
			EndTime:   time.Now().Add(24 * 400 * time.Hour).Unix(),
		},
		{
			Amount:    sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(20000))),
			StartTime: time.Now().Add(-24 * 200 * time.Hour).Unix(),
			EndTime:   time.Now().Add(24 * 150 * time.Hour).Unix(),
		},
	}
	sum := periodsSum(&periods)
	return createClaimAccout(sum, periods[1].StartTime, periods[2].EndTime, periods)

}

func createClaimAccout(originalVesting sdk.Coins, startTime int64, endTime int64, periods types.ContinuousVestingPeriods) *types.PeriodicContinuousVestingAccount {
	return types.NewRepeatedContinuousVestingAccount(&authtypes.BaseAccount{}, originalVesting, startTime, endTime, periods)
}

type ClaimAccountTc struct {
	desc    string
	account *types.PeriodicContinuousVestingAccount
	valid   bool
	message string
}

func periodsSum(periods *types.ContinuousVestingPeriods) (sum sdk.Coins) {
	for _, period := range *periods {
		sum = sum.Add(period.Amount...)
	}
	return
}
