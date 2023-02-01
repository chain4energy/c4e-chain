package types

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MessageId uint64

func (cr *UserEntry) GetClaimRecord(camapaignId uint64) *ClaimRecord {
	for _, claimRecordState := range cr.ClaimRecords {
		if claimRecordState.CampaignId == camapaignId {
			return claimRecordState
		}
	}
	return nil
}

func (c *Campaign) IsEnabled(blockTime time.Time) error {
	if !c.Enabled {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaign %d error", c.Id)
	}
	if blockTime.Before(c.StartTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaign %d not started yet (%s < startTime %s) error", c.Id, blockTime, c.StartTime)
	}
	if blockTime.After(c.EndTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaign %d has already ended (%s > endTime %s) error", c.Id, blockTime, c.EndTime)
	}
	return nil
}

func (c *Mission) IsEnabled(blockTime time.Time) error {
	if c.ClaimStartDate == nil {
		return nil
	}
	if c.ClaimStartDate.Before(blockTime) {
		return sdkerrors.Wrapf(ErrMissionDisabled, "mission %d not started yet (%s < startTime %s) error", c.Id, blockTime, c.ClaimStartDate)
	}
	return nil
}

// Validate checks the userEntry is valid
func (m *UserEntry) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return err
	}

	if len(m.ClaimRecords) == 0 {
		return errors.New("at least one campaign record is required")
	}

	campaignIDMap := make(map[uint64]struct{})
	for _, elem := range m.ClaimRecords {
		if _, ok := campaignIDMap[elem.CampaignId]; ok {
			return fmt.Errorf("duplicated campaign id for completed mission")
		}
		campaignIDMap[elem.CampaignId] = struct{}{}
	}

	for _, claimRecord := range m.ClaimRecords {
		if !claimRecord.Amount.IsAllPositive() {
			return errors.New("claimable amount must be positive")
		}

		missionIDMap := make(map[uint64]struct{})
		for _, elem := range claimRecord.CompletedMissions {
			if _, ok := missionIDMap[elem]; ok {
				return fmt.Errorf("duplicated mission id for completed mission")
			}
			missionIDMap[elem] = struct{}{}
		}
	}

	return nil
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m *UserEntry) IsMissionCompleted(campaignId uint64, missionID uint64) bool {
	for _, claimRecord := range m.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			for _, completed := range claimRecord.CompletedMissions {
				if completed == missionID {
					return true
				}
			}
		}
	}
	return false
}

func (m *UserEntry) IsMissionClaimed(campaignId uint64, missionID uint64) bool {
	for _, claimRecord := range m.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			for _, claimed := range claimRecord.ClaimedMissions {
				if claimed == missionID {
					return true
				}
			}
		}
	}
	return false
}

func (m *UserEntry) IsInitialMissionClaimed(campaignId uint64) bool {
	for _, claimRecord := range m.ClaimRecords {
		if claimRecord.CampaignId == campaignId {
			for _, claimed := range claimRecord.ClaimedMissions {
				if claimed == InitialMissionId {
					return true
				}
			}
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

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m *UserEntry) CompleteMission(campaignId uint64, missionID uint64) error {
	claimRecord := m.GetClaimRecord(campaignId)
	if claimRecord == nil {
		return fmt.Errorf("no campaign record with id %d for address %s", campaignId, m.Address)
	}
	claimRecord.CompletedMissions = append(claimRecord.CompletedMissions, missionID)
	return nil
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m *UserEntry) ClaimMission(campaignId uint64, missionID uint64) error {
	claimRecord := m.GetClaimRecord(campaignId)
	if claimRecord == nil {
		return fmt.Errorf("no campaign record with id %d for address %s", campaignId, m.Address)
	}
	claimRecord.ClaimedMissions = append(claimRecord.ClaimedMissions, missionID)
	return nil
}

// ClaimableFromMission returns the amount claimable for this claim record from the provided mission completion
func (m UserEntry) ClaimableFromMission(mission *Mission) (coinSum sdk.Coins) {
	claimRecord := m.GetClaimRecord(mission.CampaignId)
	if claimRecord == nil {
		return sdk.NewCoins() // TODO error ??
	}
	for _, coin := range claimRecord.Amount {
		coinSum = coinSum.Add(sdk.NewCoin(coin.Denom, mission.Weight.Mul(sdk.NewDecFromInt(coin.Amount)).TruncateInt()))
	}
	return
}

func CampaignCloseActionFromString(str string) (CampaignCloseAction, error) {
	option, ok := CampaignCloseAction_value[str]
	if !ok {
		return CampaignCloseAction_CLOSE_ACTION_UNSPECIFIED, fmt.Errorf("'%s' is not a valid mission type, available options: initial_claim/vote/delegation", str)
	}
	return CampaignCloseAction(option), nil
}
