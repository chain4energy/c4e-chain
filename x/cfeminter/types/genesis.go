package types

import (
	"cosmossdk.io/math"
	fmt "fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// this line is used by starport scaffolding # genesis/types/import

var _ codectypes.UnpackInterfacesMessage = GenesisState{}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		MinterState: MinterState{
			SequenceId:                  1,
			AmountMinted:                math.ZeroInt(),
			RemainderToMint:             sdk.ZeroDec(),
			LastMintBlockTime:           time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			RemainderFromPreviousMinter: sdk.ZeroDec(),
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/Validate
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	if err := gs.MinterState.Validate(); err != nil {
		return err
	}
	if !gs.Params.ContainsMinter(gs.MinterState.SequenceId) {
		return fmt.Errorf("cfeminter genesis validation error: minter state sequence id %d not found in minters", gs.MinterState.SequenceId)
	}
	return nil
}

func (gs GenesisState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return gs.Params.UnpackInterfaces(unpacker)
}
