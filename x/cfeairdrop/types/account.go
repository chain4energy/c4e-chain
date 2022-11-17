package types

import (
	"errors"
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
func NewAirdropVestingAccount(baseAcc *authtypes.BaseAccount, originalVesting sdk.Coins, startTime int64, endTime int64, periods ContinuousVestingPeriods) *AirdropVestingAccount {
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
	for _, period := range cva.VestingPeriods {
		vestedCoins = vestedCoins.Add(period.GetVestedCoins(blockTime)...)
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
	if cva.GetStartTime() != startTime {
		return fmt.Errorf("vesting start-time (%d) not eqaul to earliest period start time (%d)", cva.GetStartTime(), startTime)
	}
	if cva.GetEndTime() != endTime {
		return fmt.Errorf("vesting end-time (%d) not eqaul to lastest period end time (%d)", cva.GetEndTime(), endTime)
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
		vestedAmt := ovc.Amount.ToDec().Mul(s).RoundInt()
		vestedCoins = append(vestedCoins, sdk.NewCoin(ovc.Denom, vestedAmt))
	}

	return vestedCoins
}

// Validate checks for errors on the account fields
func (cva ContinuousVestingPeriod) Validate() error {
	if cva.GetStartTime() >= cva.GetEndTime() {
		return fmt.Errorf("vesting period end-time (%d) cannot be before start-time (%d)", cva.GetEndTime(), cva.GetStartTime())
	}
	return nil
}

// Validate checks the claimRecord is valid
func (m ClaimRecord) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return err
	}

	if len(m.CampaignRecords) == 0 {
		return errors.New("at least one campaign record is required")
	}

	campaignIDMap := make(map[uint64]struct{})
	for _, elem := range m.CampaignRecords {
		if _, ok := campaignIDMap[elem.CampaignId]; ok {
			return fmt.Errorf("duplicated campaign id for completed mission")
		}
		campaignIDMap[elem.CampaignId] = struct{}{}
	}

	for _, campaignRecord := range m.CampaignRecords {
		if !campaignRecord.Claimable.IsPositive() {
			return errors.New("claimable amount must be positive")
		}

		missionIDMap := make(map[uint64]struct{})
		for _, elem := range campaignRecord.CompletedMissions {
			if _, ok := missionIDMap[elem]; ok {
				return fmt.Errorf("duplicated mission id for completed mission")
			}
			missionIDMap[elem] = struct{}{}
		}
	}

	return nil
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m ClaimRecord) IsMissionCompleted(campaignId uint64, missionID uint64) bool {
	for _, campaignRecord := range m.CampaignRecords {
		if campaignRecord.CampaignId == campaignId {
			for _, completed := range campaignRecord.CompletedMissions {
				if completed == missionID {
					return true
				}
			}
		}
	}
	return false
}

// HasCampaign checks if the specified reccord for campignId ID exists
func (m ClaimRecord) HasCampaign(campaignId uint64) bool {
	for _, campaignRecord := range m.CampaignRecords {
		if campaignRecord.CampaignId == campaignId {
			return true
		}
	}
	return false
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m ClaimRecord) CompleteMission(campaignId uint64, missionID uint64) error {
	for _, campaignRecord := range m.CampaignRecords {
		campaignRecord.CompletedMissions = append(campaignRecord.CompletedMissions, missionID)
		return nil
	}
	return fmt.Errorf("no campaign with id %d", campaignId)
}

// ClaimableFromMission returns the amount claimable for this claim record from the provided mission completion
func (m ClaimRecord) ClaimableFromMission(mission Mission) sdk.Int {
	for _, campaignRecord := range m.CampaignRecords {
		return mission.Weight.Mul(sdk.NewDecFromInt(campaignRecord.Claimable)).TruncateInt()
	}
	return sdk.ZeroInt() // TODO panic ??
}
