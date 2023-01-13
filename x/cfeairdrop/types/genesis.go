package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ClaimRecords:     []ClaimRecord{},
		InitialClaims:    []InitialClaim{},
		Missions:         []Mission{},
		AirdropEntryList: []AirdropEntry{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in claimRecordXX
	claimRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClaimRecords {
		index := string(ClaimRecordKey(elem.Address))
		if _, ok := claimRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for claimRecordXX")
		}
		claimRecordIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in initialClaim
	initialClaimIndexMap := make(map[uint64]struct{})

	for _, elem := range gs.InitialClaims {
		index := elem.CampaignId
		if _, ok := initialClaimIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for initialClaim")
		}
		initialClaimIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in mission
	missionIndexMap := make(map[string]struct{})

	for _, elem := range gs.Missions {
		index := string(MissionKey(elem.CampaignId, elem.MissionId))
		if _, ok := missionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mission")
		}
		missionIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in airdropEntry
	airdropEntryIdMap := make(map[uint64]bool)
	airdropEntryCount := gs.GetAirdropEntryCount()
	for _, elem := range gs.AirdropEntryList {
		if _, ok := airdropEntryIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for airdropEntry")
		}
		if elem.Id >= airdropEntryCount {
			return fmt.Errorf("airdropEntry id should be lower or equal than the last id")
		}
		airdropEntryIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
