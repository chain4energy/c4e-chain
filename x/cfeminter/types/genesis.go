package types

import (
	fmt "fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		MinterState: MinterState{
			SequenceId:                  1,
			AmountMinted:                sdk.ZeroInt(),
			RemainderToMint:             sdk.ZeroDec(),
			LastMintBlockTime:           time.Now(),
			RemainderFromPreviousPeriod: sdk.ZeroDec(),
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	if err := gs.MinterState.Validate(); err != nil {
		return err
	}
	if !gs.Params.MinterConfig.ContainsMinter(gs.MinterState.SequenceId) {
		return fmt.Errorf("cfeminter genesis validation error: minter state sequence id %d not found in minters", gs.MinterState.SequenceId)
	}
	return nil
}
