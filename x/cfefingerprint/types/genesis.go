package types

import "fmt"

// this line is used by starport scaffolding # genesis/types/import

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs GenesisState) Validate() error {
	payloadLinkMap := make(map[string]bool)
	for _, payloadLink := range gs.PayloadLinks {
		if _, ok := payloadLinkMap[payloadLink.ReferenceKey]; ok {
			return fmt.Errorf("duplicated reference key %s", payloadLink.ReferenceKey)
		}
		payloadLinkMap[payloadLink.ReferenceKey] = true
		if err := payloadLink.Validate(); err != nil {
			return err
		}
	}
	return nil
}
