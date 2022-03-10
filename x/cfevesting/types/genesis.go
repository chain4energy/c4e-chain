package types

import (
// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	// TODO validate accouns vestings
	// testmapIndexMap := make(map[string]struct{})

	// for _, elem := range gs.TestmapList {
	// 	index := string(TestmapKey(elem.Index))
	// 	if _, ok := testmapIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for testmap")
	// 	}
	// 	testmapIndexMap[index] = struct{}{}
	// }

	return gs.Params.Validate()
}
