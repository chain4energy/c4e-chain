package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EnergyTokenList: []EnergyToken{},
		TokenParamsList: []TokenParams{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in energyToken
	energyTokenIdMap := make(map[uint64]bool)
	energyTokenCount := gs.GetEnergyTokenCount()
	for _, elem := range gs.EnergyTokenList {
		if _, ok := energyTokenIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for energyToken")
		}
		if elem.Id >= energyTokenCount {
			return fmt.Errorf("energyToken id should be lower or equal than the last id")
		}
		energyTokenIdMap[elem.Id] = true
	}
	// Check for duplicated index in tokenParams
	tokenParamsIndexMap := make(map[string]struct{})

	for _, elem := range gs.TokenParamsList {
		index := string(TokenParamsKey(elem.Index))
		if _, ok := tokenParamsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tokenParams")
		}
		tokenParamsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
