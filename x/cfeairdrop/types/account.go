package types

import (
	"errors"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var (
	_ vestexported.VestingAccount = (*AirdropVestingAccount)(nil)
	_ authtypes.GenesisAccount    = (*AirdropVestingAccount)(nil)
)

type ContinuousVestingPeriods []ContinuousVestingPeriod

// NewAirdropVestingAccountRaw creates a new AirdropVestingAccount object from BaseVestingAccount
func NewAirdropVestingAccountRaw(bva *vestingtypes.BaseVestingAccount, startTime int64) *AirdropVestingAccount {
	return &AirdropVestingAccount{
		BaseVestingAccount: bva,
		StartTime:          startTime,
	}
}

// NewAirdropVestingAccount returns a new AirdropVestingAccount
func NewAirdropVestingAccount(baseAcc *authtypes.BaseAccount, originalVesting sdk.Coins, startTime int64, periods ContinuousVestingPeriods) *AirdropVestingAccount {
	endTime := startTime
	// for _, p := range periods {
	// 	endTime += p.Length
	// }
	baseVestingAcc := &vestingtypes.BaseVestingAccount{
		BaseAccount:     baseAcc,
		OriginalVesting: originalVesting,
		EndTime:         endTime,
	}

	return &AirdropVestingAccount{
		BaseVestingAccount: baseVestingAcc,
		StartTime:          startTime,
		VestingPeriods:     periods,
	}
}

// GetVestedCoins returns the total number of vested coins. If no coins are vested,
// nil is returned.
func (cva AirdropVestingAccount) GetVestedCoins(blockTime time.Time) sdk.Coins {
	var vestedCoins sdk.Coins

	// We must handle the case where the start time for a vesting account has
	// been set into the future or when the start of the chain is not exactly
	// known.
	if blockTime.Unix() <= cva.StartTime {
		return vestedCoins
	} else if blockTime.Unix() >= cva.EndTime {
		return cva.OriginalVesting
	}

	// calculate the vesting scalar
	x := blockTime.Unix() - cva.StartTime
	y := cva.EndTime - cva.StartTime
	s := sdk.NewDec(x).Quo(sdk.NewDec(y))

	for _, ovc := range cva.OriginalVesting {
		vestedAmt := sdk.NewDecFromInt(ovc.Amount).Mul(s).RoundInt()
		vestedCoins = append(vestedCoins, sdk.NewCoin(ovc.Denom, vestedAmt))
	}

	return vestedCoins
}

// GetVestingCoins returns the total number of vesting coins. If no coins are
// vesting, nil is returned.
func (cva AirdropVestingAccount) GetVestingCoins(blockTime time.Time) sdk.Coins {
	return cva.OriginalVesting.Sub(cva.GetVestedCoins(blockTime))
}

// LockedCoins returns the set of coins that are not spendable (i.e. locked),
// defined as the vesting coins that are not delegated.
func (cva AirdropVestingAccount) LockedCoins(blockTime time.Time) sdk.Coins {
	return cva.BaseVestingAccount.LockedCoinsFromVesting(cva.GetVestingCoins(blockTime))
}

// TrackDelegation tracks a desired delegation amount by setting the appropriate
// values for the amount of delegated vesting, delegated free, and reducing the
// overall amount of base coins.
func (cva *AirdropVestingAccount) TrackDelegation(blockTime time.Time, balance, amount sdk.Coins) {
	cva.BaseVestingAccount.TrackDelegation(balance, cva.GetVestingCoins(blockTime), amount)
}

// GetStartTime returns the time when vesting starts for a continuous vesting
// account.

func (acc AirdropVestingAccount) GetStartTime() int64 {
	return acc.StartTime
}

func (acc AirdropVestingAccount) String() string {
	out, _ := acc.MarshalYAML()
	return out.(string)
}

func (acc AirdropVestingAccount) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &acc)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

// Validate checks for errors on the account fields
func (cva AirdropVestingAccount) Validate() error {
	if cva.GetStartTime() >= cva.GetEndTime() {
		return errors.New("vesting start-time cannot be before end-time")
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