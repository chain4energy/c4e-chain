package types

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var OneToken = sdk.NewInt(1000000)

type MessageId uint64

func (cr *UserAirdropEntries) GetAidropEntry(camapaignId uint64) *AirdropEntry {
	for _, airdropEntryState := range cr.AirdropEntries {
		if airdropEntryState.CampaignId == camapaignId {
			return airdropEntryState
		}
	}
	return nil
}

func (c *Campaign) IsEnabled(blockTime time.Time) error {
	if !c.Enabled {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaignId %d", c.Id)
	}
	if blockTime.Before(*c.StartTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaignId %d not started: time %s < startTime %s", c.Id, blockTime, c.StartTime)
	}
	if c.EndTime != nil && blockTime.After(*c.EndTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaignId %d ended: time %s > endTime %s", c.Id, blockTime, c.EndTime)
	}
	return nil
}

func (c *Mission) IsEnabled(blockTime time.Time) error {
	if c.ClaimStartDate == nil {
		return nil
	}
	if c.ClaimStartDate.Before(blockTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "missionId %d not started: time %s < startTime %s", c.Id, blockTime, c.ClaimStartDate)
	}
	return nil
}

func NewAirdropCampaign(owner string, name string, description string, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) *Campaign {

	return &Campaign{
		Id:            0,
		Owner:         owner,
		Name:          name,
		Description:   description,
		Enabled:       false,
		StartTime:     &startTime,
		EndTime:       &endTime,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
	}
}

// Validate checks the userAirdropEntries is valid
func (m *UserAirdropEntries) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return err
	}

	if len(m.AirdropEntries) == 0 {
		return errors.New("at least one campaign record is required")
	}

	campaignIDMap := make(map[uint64]struct{})
	for _, elem := range m.AirdropEntries {
		if _, ok := campaignIDMap[elem.CampaignId]; ok {
			return fmt.Errorf("duplicated campaign id for completed mission")
		}
		campaignIDMap[elem.CampaignId] = struct{}{}
	}

	for _, airdropEntry := range m.AirdropEntries {
		if !airdropEntry.Amount.IsAllPositive() {
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
func (m *UserAirdropEntries) IsMissionCompleted(campaignId uint64, missionID uint64) bool {
	for _, airdropEntry := range m.AirdropEntries {
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

func (m *UserAirdropEntries) IsMissionClaimed(campaignId uint64, missionID uint64) bool {
	for _, airdropEntry := range m.AirdropEntries {
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

func (m *UserAirdropEntries) IsInitialMissionClaimed(campaignId uint64) bool {
	for _, airdropEntry := range m.AirdropEntries {
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
func (m UserAirdropEntries) HasCampaign(campaignId uint64) bool {
	for _, airdropEntry := range m.AirdropEntries {
		if airdropEntry.CampaignId == campaignId {
			return true
		}
	}
	return false
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m *UserAirdropEntries) CompleteMission(campaignId uint64, missionID uint64) error {
	airdropEntry := m.GetAidropEntry(campaignId)
	if airdropEntry == nil {
		return fmt.Errorf("no campaign record with id %d for address %s", campaignId, m.Address)
	}
	airdropEntry.CompletedMissions = append(airdropEntry.CompletedMissions, missionID)
	return nil
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m *UserAirdropEntries) ClaimMission(campaignId uint64, missionID uint64) error {
	airdropEntry := m.GetAidropEntry(campaignId)
	if airdropEntry == nil {
		return fmt.Errorf("no campaign record with id %d for address %s", campaignId, m.Address)
	}
	airdropEntry.ClaimedMissions = append(airdropEntry.ClaimedMissions, missionID)
	return nil
}

// ClaimableFromMission returns the amount claimable for this claim record from the provided mission completion
func (m UserAirdropEntries) ClaimableFromMission(mission *Mission) (coinSum sdk.Coins) {
	airdropEntry := m.GetAidropEntry(mission.CampaignId)
	if airdropEntry == nil {
		return sdk.NewCoins() // TODO error ??
	}
	for _, amount := range airdropEntry.Amount {
		coinSum = coinSum.Add(sdk.NewCoin(amount.Denom, mission.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()))
	}
	return
}
