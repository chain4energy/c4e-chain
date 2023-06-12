package types

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
	//States []*State
	var states []State
	for _, state := range gs.States {
		if err := state.Validate(); err != nil {
			return err
		}
		states = append(states, *state)
	}
	if _, err := StateSumIsInteger(states); err != nil {
		return err
	}

	return gs.Params.Validate()
}
