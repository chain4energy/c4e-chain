package types

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate checks the userEntry is valid
func (m *UserEntry) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return err
	}

	if len(m.ClaimRecords) == 0 {
		return fmt.Errorf("at least one campaign record is required")
	}

	campaignIDMap := make(map[uint64]struct{})
	for _, elem := range m.ClaimRecords {
		if _, ok := campaignIDMap[elem.CampaignId]; ok {
			return fmt.Errorf("duplicated campaign id for completed mission")
		}
		campaignIDMap[elem.CampaignId] = struct{}{}
	}

	for i, claimRecord := range m.ClaimRecords {
		if err := claimRecord.Validate(); err != nil {
			return errors.Wrapf(err, "claim record index %d", i)
		}
		if !claimRecord.Amount.IsAllPositive() {
			return fmt.Errorf("claimable amount must be positive")
		}

		missionIdMap := make(map[uint64]struct{})
		for _, elem := range claimRecord.CompletedMissions {
			if _, ok := missionIdMap[elem]; ok {
				return fmt.Errorf("duplicated mission id for completed mission")
			}
			missionIdMap[elem] = struct{}{}
		}
	}

	return nil
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m *ClaimRecord) IsMissionCompleted(missionId uint64) bool {
	for _, completed := range m.CompletedMissions {
		if completed == missionId {
			return true
		}
	}

	return false
}

func (m *ClaimRecord) IsMissionClaimed(missionId uint64) bool {
	for _, claimed := range m.ClaimedMissions {
		if claimed == missionId {
			return true
		}
	}

	return false
}

func (m *ClaimRecord) IsInitialMissionClaimed() bool {
	for _, claimed := range m.ClaimedMissions {
		if claimed == InitialMissionId {
			return true
		}
	}

	return false
}

// HasCampaign checks if the specified reccord for campignId ID exists
func (m *UserEntry) HasCampaign(campaignId uint64) bool {
	for _, claimRecord := range m.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			return true
		}
	}
	return false
}

// CompleteMission checks if the specified mission ID is completed for the claim record
func (m *ClaimRecord) CompleteMission(campaignId uint64, missionId uint64) error {
	if m.IsMissionCompleted(missionId) {
		return errors.Wrapf(ErrMissionCompleted, "campaignId: %d, missionId: %d", campaignId, missionId)
	}
	m.CompletedMissions = append(m.CompletedMissions, missionId)
	return nil
}

// CalculateInitialClaimClaimableAmount calculates the initial claimable amount for the claim record
func (m *ClaimRecord) CalculateInitialClaimClaimableAmount(weightSum sdk.Dec) sdk.Coins {
	allMissionsAmountSum := sdk.NewCoins()
	for _, amount := range m.Amount {
		allMissionsAmountSum = allMissionsAmountSum.Add(sdk.NewCoin(amount.Denom, weightSum.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()))
	}
	return m.Amount.Sub(allMissionsAmountSum...)
}

// ClaimMission claims the specified mission ID for the claim record
func (m *ClaimRecord) ClaimMission(campaignId uint64, missionId uint64) error {
	if !m.IsMissionCompleted(missionId) {
		return errors.Wrapf(ErrMissionNotCompleted, "campaignId: %d, missionId: %d", campaignId, missionId)
	}

	if m.IsMissionClaimed(missionId) {
		return errors.Wrapf(ErrMissionClaimed, "campaignId: %d, missionId: %d", campaignId, missionId)
	}
	m.ClaimedMissions = append(m.ClaimedMissions, missionId)
	return nil
}

// ClaimableFromMission returns the amount claimable for this claim record from the provided mission completion
func (m ClaimRecord) ClaimableFromMission(mission *Mission) (sdk.Coins, error) {
	var coinSum sdk.Coins
	for _, coin := range m.Amount {
		coinSum = coinSum.Add(sdk.NewCoin(coin.Denom, mission.Weight.Mul(sdk.NewDecFromInt(coin.Amount)).TruncateInt()))
	}
	return coinSum, nil
}

func (cr *UserEntry) GetClaimRecord(campaignId uint64) *ClaimRecord {
	for _, claimRecord := range cr.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			return claimRecord
		}
	}
	return nil
}

// DeleteClaimRecord deletes the claim record for the specified campaign ID
func (cr *UserEntry) DeleteClaimRecord(campaignId uint64) {
	for i, claimRecord := range cr.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			cr.ClaimRecords = append(cr.ClaimRecords[:i], cr.ClaimRecords[i+1:]...)
		}
	}
}

func (cr *UserEntry) MustGetClaimRecord(camapaignId uint64) (*ClaimRecord, error) {
	for _, claimRecord := range cr.ClaimRecords {
		if claimRecord.CampaignId == camapaignId {
			return claimRecord, nil
		}
	}
	return nil, errors.Wrapf(sdkerrors.ErrNotFound, "claim record with campaign id %d not found for address %s", camapaignId, cr.Address)
}

func (claimRecord *ClaimRecord) Validate() error {
	if claimRecord.Amount == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record amount cannot be nil")
	}
	if claimRecord.Amount.IsAnyNil() {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record amount cannot be nil")
	}
	if err := claimRecord.Amount.Validate(); err != nil {
		return errors.Wrapf(c4eerrors.ErrParam, "wrong claim record amount (%s)", err)
	}
	return nil
}

func (claimRecordEntry *ClaimRecordEntry) Validate() error {
	_, err := sdk.AccAddressFromBech32(claimRecordEntry.UserEntryAddress)
	if err != nil {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record entry user entry address parsing error (%s)", err)
	}
	if claimRecordEntry.Amount == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record entry amount cannot be nil")
	}
	if claimRecordEntry.Amount.IsAnyNil() {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record entry amount cannot be nil")
	}
	if err := claimRecordEntry.Amount.Validate(); err != nil {
		return errors.Wrapf(c4eerrors.ErrParam, "wrong claim record entry amount (%s)", err)
	}
	if !claimRecordEntry.Amount.IsAllPositive() {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record amount must be positive")
	}
	return nil
}

func ValidateClaimRecordEntries(claimRecordEntries []*ClaimRecordEntry) error {
	for i, claimRecordEntry := range claimRecordEntries {
		if err := claimRecordEntry.Validate(); err != nil {
			return errors.Wrapf(err, "claim record entry index %d", i)
		}
	}
	return nil
}

func ValidateVestingPoolCampaignClaimRecordEntries(claimRecordEntries []*ClaimRecordEntry, vestingDenom string) error {
	for i, claimRecordEntry := range claimRecordEntries {
		if err := claimRecordEntry.Validate(); err != nil {
			return errors.Wrapf(err, "claim record entry index %d", i)
		}
		if claimRecordEntry.Amount.Len() > 1 || claimRecordEntry.Amount.AmountOf(vestingDenom).IsZero() {
			return errors.Wrapf(c4eerrors.ErrParam, "claim record entry index %d: for vesting pool campaigns,"+
				" the claim record entry must have only one coin with the denomination currently used by the cfevesting module (%s)", i, vestingDenom)
		}
	}
	return nil
}
