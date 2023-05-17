package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	for _, claimRecord := range m.ClaimRecords {
		if !claimRecord.Amount.IsAllPositive() {
			return fmt.Errorf("claimable amount must be positive")
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
func (m UserEntry) ClaimableFromMission(mission *Mission) (sdk.Coins, error) {
	claimRecord := m.GetClaimRecord(mission.CampaignId)

	if claimRecord == nil {
		return nil, fmt.Errorf("no campaign record with id %d for address %s", mission.CampaignId, m.Address)
	}

	var coinSum sdk.Coins
	for _, coin := range claimRecord.Amount {
		coinSum = coinSum.Add(sdk.NewCoin(coin.Denom, mission.Weight.Mul(sdk.NewDecFromInt(coin.Amount)).TruncateInt()))
	}
	return coinSum, nil
}

func (cr *UserEntry) GetClaimRecord(camapaignId uint64) *ClaimRecord {
	for _, claimRecord := range cr.ClaimRecords {
		if claimRecord.CampaignId == camapaignId {
			return claimRecord
		}
	}
	return nil
}
