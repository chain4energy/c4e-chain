package types

import (
	fmt "fmt"
	"time"

	"math"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var (
	_ vestexported.VestingAccount = (*PeriodicContinuousVestingAccount)(nil)
	_ authtypes.GenesisAccount    = (*PeriodicContinuousVestingAccount)(nil)
)

type ContinuousVestingPeriods []ContinuousVestingPeriod

// NewPeriodicContinuousVestingAccountRaw creates a new VestingAccount object from BaseVestingAccount
func NewPeriodicContinuousVestingAccountRaw(bva *vestingtypes.BaseVestingAccount, startTime int64) *PeriodicContinuousVestingAccount {
	return &PeriodicContinuousVestingAccount{
		BaseVestingAccount: bva,
		StartTime:          startTime,
	}
}

// NewRepeatedContinuousVestingAccount returns a new VestingAccount
func NewRepeatedContinuousVestingAccount(baseAcc *authtypes.BaseAccount, originalVesting sdk.Coins, startTime int64, endTime int64, periods ContinuousVestingPeriods) *PeriodicContinuousVestingAccount {
	baseVestingAcc := &vestingtypes.BaseVestingAccount{
		BaseAccount:     baseAcc,
		OriginalVesting: originalVesting,
		EndTime:         endTime,
	}

	return &PeriodicContinuousVestingAccount{
		BaseVestingAccount: baseVestingAcc,
		StartTime:          startTime,
		VestingPeriods:     periods,
	}
}

func AddNewContinousVestingPeriods(claimerAccount *PeriodicContinuousVestingAccount, startTime int64, endTime int64, amount sdk.Coins) *PeriodicContinuousVestingAccount {
	hadPariods := len(claimerAccount.VestingPeriods) > 0

	claimerAccount.VestingPeriods = append(claimerAccount.VestingPeriods,
		ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: amount})

	claimerAccount.BaseVestingAccount.OriginalVesting = claimerAccount.BaseVestingAccount.OriginalVesting.Add(amount...)
	if !hadPariods || endTime > claimerAccount.BaseVestingAccount.EndTime {
		claimerAccount.BaseVestingAccount.EndTime = endTime
	}
	if !hadPariods || startTime < claimerAccount.StartTime {
		claimerAccount.StartTime = startTime
	}
	return claimerAccount
}

// GetVestedCoins returns the total number of vested coins. If no coins are vested,
// nil is returned.
func (cva PeriodicContinuousVestingAccount) GetVestedCoins(blockTime time.Time) sdk.Coins {
	var vestedCoins sdk.Coins
	for _, period := range cva.VestingPeriods {
		vestedCoins = vestedCoins.Add(period.GetVestedCoins(blockTime)...)
	}
	return vestedCoins
}

// GetVestingCoins returns the total number of vesting coins. If no coins are
// vesting, nil is returned.
func (cva PeriodicContinuousVestingAccount) GetVestingCoins(blockTime time.Time) sdk.Coins {
	return cva.OriginalVesting.Sub(cva.GetVestedCoins(blockTime)...)
}

// LockedCoins returns the set of coins that are not spendable (i.e. locked),
// defined as the vesting coins that are not delegated.
func (cva PeriodicContinuousVestingAccount) LockedCoins(blockTime time.Time) sdk.Coins {
	return cva.BaseVestingAccount.LockedCoinsFromVesting(cva.GetVestingCoins(blockTime))
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (cva *PeriodicContinuousVestingAccount) TrackDelegation(blockTime time.Time, balance, amount sdk.Coins) {
	cva.BaseVestingAccount.TrackDelegation(balance, cva.GetVestingCoins(blockTime), amount)
}

// GetStartTime returns the time when vesting starts for a continuous vesting
// account.

func (acc PeriodicContinuousVestingAccount) GetStartTime() int64 {
	return acc.StartTime
}

func (acc PeriodicContinuousVestingAccount) String() string {
	out, _ := acc.MarshalYAML()
	return out.(string)
}

func (acc PeriodicContinuousVestingAccount) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &acc)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

// Validate checks for errors on the account fields
func (cva *PeriodicContinuousVestingAccount) Validate() error {
	if cva.GetStartTime() > cva.GetEndTime() {
		return fmt.Errorf("vesting end-time (%d) cannot be before start-time (%d)", cva.GetEndTime(), cva.GetStartTime())
	}
	var vestedCoins sdk.Coins
	var startTime int64 = math.MaxInt64
	var endTime int64 = 0

	for _, period := range cva.VestingPeriods {
		if err := period.Validate(); err != nil {
			return err
		}
		if period.StartTime < startTime {
			startTime = period.StartTime
		}
		if period.EndTime > endTime {
			endTime = period.EndTime
		}
		vestedCoins = vestedCoins.Add(period.Amount...)
	}
	if !cva.BaseVestingAccount.OriginalVesting.IsEqual(vestedCoins) {
		return fmt.Errorf("original vesting (%s) not equal to sum of periods (%s)", cva.BaseVestingAccount.OriginalVesting, vestedCoins)
	}
	if len(cva.VestingPeriods) > 0 {
		if cva.GetStartTime() != startTime {
			return fmt.Errorf("vesting start-time (%d) not eqaul to earliest period start time (%d)", cva.GetStartTime(), startTime)
		}
		if cva.GetEndTime() != endTime {
			return fmt.Errorf("vesting end-time (%d) not eqaul to lastest period end time (%d)", cva.GetEndTime(), endTime)
		}
	}
	return cva.BaseVestingAccount.Validate()
}

func (cvp ContinuousVestingPeriod) String() string {
	out, _ := cvp.MarshalYAML()
	return out.(string)
}

func (cvp ContinuousVestingPeriod) MarshalYAML() (interface{}, error) {
	// TODO checkit with cosmos sdk 0.45.9
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &cvp)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

func (cva ContinuousVestingPeriod) GetVestedCoins(blockTime time.Time) sdk.Coins {
	var vestedCoins sdk.Coins

	// We must handle the case where the start time for a vesting account has
	// been set into the future or when the start of the chain is not exactly
	// known.
	if blockTime.Unix() <= cva.StartTime {
		return vestedCoins
	} else if blockTime.Unix() >= cva.EndTime {
		return cva.Amount
	}

	// calculate the vesting scalar
	x := blockTime.Unix() - cva.StartTime
	y := cva.EndTime - cva.StartTime
	s := sdk.NewDec(x).Quo(sdk.NewDec(y))

	for _, ovc := range cva.Amount {
		vestedAmt := sdk.NewDecFromInt(ovc.Amount).Mul(s).RoundInt()
		vestedCoins = append(vestedCoins, sdk.NewCoin(ovc.Denom, vestedAmt))
	}

	return vestedCoins
}

// Validate checks for errors on the account fields
func (cva *ContinuousVestingPeriod) Validate() error {
	if cva.GetStartTime() >= cva.GetEndTime() {
		return fmt.Errorf("vesting period end-time (%d) cannot be before start-time (%d)", cva.GetEndTime(), cva.GetStartTime())
	}
	return nil
}
