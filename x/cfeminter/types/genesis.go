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
	err := gs.Params.Validate()
	if err != nil {
		return err
	}

	minterState := gs.MinterState
	err = minterState.Validate()
	if err != nil {
		return err
	}

	if !gs.Params.MinterConfig.ContainsMinter(minterState.SequenceId) {
		return fmt.Errorf("minter state current sequence id not found in minters")
	}
	return nil

}
