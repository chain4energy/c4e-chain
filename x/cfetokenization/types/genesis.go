package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CertificateTypeList: []CertificateType{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in certificateType
	certificateTypeIdMap := make(map[uint64]bool)
	certificateTypeCount := gs.GetCertificateTypeCount()
	for _, elem := range gs.CertificateTypeList {
		if _, ok := certificateTypeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for certificateType")
		}
		if elem.Id >= certificateTypeCount {
			return fmt.Errorf("certificateType id should be lower or equal than the last id")
		}
		certificateTypeIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
