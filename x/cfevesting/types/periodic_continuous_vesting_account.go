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

func (account *PeriodicContinuousVestingAccount) AddNewContinousVestingPeriod(startTime int64, endTime int64, amount sdk.Coins) uint64 {
	vestingPeriodsLen := len(account.VestingPeriods)
	hadPariods := vestingPeriodsLen > 0

	account.VestingPeriods = append(account.VestingPeriods,
		ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: amount})

	account.BaseVestingAccount.OriginalVesting = account.BaseVestingAccount.OriginalVesting.Add(amount...)
	if !hadPariods || endTime > account.BaseVestingAccount.EndTime {
		account.BaseVestingAccount.EndTime = endTime
	}
	if !hadPariods || startTime < account.StartTime {
		account.StartTime = startTime
	}
	return uint64(vestingPeriodsLen)
}

// GetVestedCoins returns the total number of vested coins. If no coins are vested,
// nil is returned.
func (account PeriodicContinuousVestingAccount) GetVestedCoins(blockTime time.Time) sdk.Coins {
	var vestedCoins sdk.Coins
	for _, period := range account.VestingPeriods {
		vestedCoins = vestedCoins.Add(period.GetVestedCoins(blockTime)...)
	}
	return vestedCoins
}

// GetVestingCoinsForSpecyficPeriods returns the total number of vesting coins. If no coins are
// vesting, nil is returned.
func (account PeriodicContinuousVestingAccount) GetVestingCoinsForSpecyficPeriods(blockTime time.Time, periodsToTrace []uint64) sdk.Coins {
	return account.OriginalVesting.Sub(account.GetVestedCoinsForSpecyficPeriods(blockTime, periodsToTrace)...)
}

// GetVestedCoinsForSpecyficPeriods returns the total number of vested coins. If no coins are vested,
// nil is returned.
func (account PeriodicContinuousVestingAccount) GetVestedCoinsForSpecyficPeriods(blockTime time.Time, periodsToTrace []uint64) sdk.Coins {
	var vestedCoins sdk.Coins
	for _, periodId := range periodsToTrace {
		vestedCoins = vestedCoins.Add(account.VestingPeriods[periodId].GetVestedCoins(blockTime)...)
	}
	return vestedCoins
}

// GetVestingCoins returns the total number of vesting coins. If no coins are
// vesting, nil is returned.
func (account PeriodicContinuousVestingAccount) GetVestingCoins(blockTime time.Time) sdk.Coins {
	return account.OriginalVesting.Sub(account.GetVestedCoins(blockTime)...)
}

// GetAllLockedCoins returns the set of coins that are not spendable (i.e. locked),
// defined as the vesting coins that are not delegated.
func (account PeriodicContinuousVestingAccount) LockedCoins(blockTime time.Time) sdk.Coins {
	return account.BaseVestingAccount.LockedCoinsFromVesting(account.GetVestingCoins(blockTime))
}

// GetLockedCoinsFromCoins returns the set of coins that are not spendable (i.e. locked),
// defined as the vesting coins that are not delegated.
func (account PeriodicContinuousVestingAccount) GetLockedCoins(vestingCoins sdk.Coins) sdk.Coins {
	return account.BaseVestingAccount.LockedCoinsFromVesting(vestingCoins)
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (account *PeriodicContinuousVestingAccount) TrackDelegation(blockTime time.Time, balance, amount sdk.Coins) {
	account.BaseVestingAccount.TrackDelegation(balance, account.GetVestingCoins(blockTime), amount)
}

// GetStartTime returns the time when vesting starts for a continuous vesting
// account.

func (account PeriodicContinuousVestingAccount) GetStartTime() int64 {
	return account.StartTime
}

func (account PeriodicContinuousVestingAccount) String() string {
	out, _ := account.MarshalYAML()
	return out.(string)
}

func (account PeriodicContinuousVestingAccount) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &account)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

// Validate checks for errors on the account fields
func (account *PeriodicContinuousVestingAccount) Validate() error {
	if account.GetStartTime() > account.GetEndTime() {
		return fmt.Errorf("vesting end-time (%d) cannot be before start-time (%d)", account.GetEndTime(), account.GetStartTime())
	}
	var vestedCoins sdk.Coins
	var startTime int64 = math.MaxInt64
	var endTime int64 = 0

	for _, period := range account.VestingPeriods {
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
	if !account.BaseVestingAccount.OriginalVesting.IsEqual(vestedCoins) {
		return fmt.Errorf("original vesting (%s) not equal to sum of periods (%s)", account.BaseVestingAccount.OriginalVesting, vestedCoins)
	}
	if len(account.VestingPeriods) > 0 {
		if account.GetStartTime() != startTime {
			return fmt.Errorf("vesting start-time (%d) not eqaul to earliest period start time (%d)", account.GetStartTime(), startTime)
		}
		if account.GetEndTime() != endTime {
			return fmt.Errorf("vesting end-time (%d) not eqaul to lastest period end time (%d)", account.GetEndTime(), endTime)
		}
	}
	return account.BaseVestingAccount.Validate()
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
