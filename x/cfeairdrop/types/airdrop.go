package types

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MessageId uint64

func (cr *UserEntry) GetAidropEntry(camapaignId uint64) *ClaimRecord {
	for _, airdropEntryState := range cr.ClaimRecords {
		if airdropEntryState.CampaignId == camapaignId {
			return airdropEntryState
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

	for _, airdropEntry := range m.ClaimRecords {
		if !airdropEntry.AirdropCoins.IsAllPositive() {
			return errors.New("claimable amount must be positive")
		}

		missionIDMap := make(map[uint64]struct{})
		for _, elem := range airdropEntry.CompletedMissions {
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
	for _, airdropEntry := range m.ClaimRecords {
		if airdropEntry.CampaignId == campaignId {
			for _, completed := range airdropEntry.CompletedMissions {
				if completed == missionID {
					return true
				}
			}
		}
	}
	return false
}

func (m *UserEntry) IsMissionClaimed(campaignId uint64, missionID uint64) bool {
	for _, airdropEntry := range m.ClaimRecords {
		if airdropEntry.CampaignId == campaignId {
			for _, claimed := range airdropEntry.ClaimedMissions {
				if claimed == missionID {
					return true
				}
			}
		}
	}
	return false
}

func (m *UserEntry) IsInitialMissionClaimed(campaignId uint64) bool {
	for _, airdropEntry := range m.ClaimRecords {
		if airdropEntry.CampaignId == campaignId {
			for _, claimed := range airdropEntry.ClaimedMissions {
				if claimed == InitialMissionId {
					return true
				}
			}
		}
	}
	return false
}

// HasCampaign checks if the specified reccord for campignId ID exists
func (m UserEntry) HasCampaign(campaignId uint64) bool {
	for _, airdropEntry := range m.ClaimRecords {
		if airdropEntry.CampaignId == campaignId {
			return true
		}
	}
	return false
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m *UserEntry) CompleteMission(campaignId uint64, missionID uint64) error {
	airdropEntry := m.GetAidropEntry(campaignId)
	if airdropEntry == nil {
		return fmt.Errorf("no campaign record with id %d for address %s", campaignId, m.Address)
	}
	airdropEntry.CompletedMissions = append(airdropEntry.CompletedMissions, missionID)
	return nil
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m *UserEntry) ClaimMission(campaignId uint64, missionID uint64) error {
	airdropEntry := m.GetAidropEntry(campaignId)
	if airdropEntry == nil {
		return fmt.Errorf("no campaign record with id %d for address %s", campaignId, m.Address)
	}
	airdropEntry.ClaimedMissions = append(airdropEntry.ClaimedMissions, missionID)
	return nil
}

// ClaimableFromMission returns the amount claimable for this claim record from the provided mission completion
func (m UserEntry) ClaimableFromMission(mission *Mission) (coinSum sdk.Coins) {
	airdropEntry := m.GetAidropEntry(mission.CampaignId)
	if airdropEntry == nil {
		return sdk.NewCoins() // TODO error ??
	}
	for _, coin := range airdropEntry.AirdropCoins {
		coinSum = coinSum.Add(sdk.NewCoin(coin.Denom, mission.Weight.Mul(sdk.NewDecFromInt(coin.Amount)).TruncateInt()))
	}
	return
}

func AirdropCloseActionFromString(str string) (AirdropCloseAction, error) {
	option, ok := AirdropCloseAction_value[str]
	if !ok {
		return AirdropCloseAction_AIRDROP_CLOSE_ACTION_UNSPECIFIED, fmt.Errorf("'%s' is not a valid mission type, available options: initial_claim/vote/delegation", str)
	}
	return AirdropCloseAction(option), nil
}
