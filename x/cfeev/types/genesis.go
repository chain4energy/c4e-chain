package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EnergyTransferOfferList: []EnergyTransferOffer{},
		EnergyTransferList:      []EnergyTransfer{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in energyTransferOffer
	energyTransferOfferIdMap := make(map[uint64]bool)
	energyTransferOfferCount := gs.GetEnergyTransferOfferCount()
	for _, elem := range gs.EnergyTransferOfferList {
		if _, ok := energyTransferOfferIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for energyTransferOffer")
		}
		if elem.Id >= energyTransferOfferCount {
			return fmt.Errorf("energyTransferOffer id should be lower or equal than the last id")
		}
		energyTransferOfferIdMap[elem.Id] = true
	}
	// Check for duplicated ID in energyTransfer
	energyTransferIdMap := make(map[uint64]bool)
	energyTransferCount := gs.GetEnergyTransferCount()
	for _, elem := range gs.EnergyTransferList {
		if _, ok := energyTransferIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for energyTransfer")
		}
		if elem.Id >= energyTransferCount {
			return fmt.Errorf("energyTransfer id should be lower or equal than the last id")
		}
		energyTransferIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
