package types

import (
	fmt "fmt"
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyMintDenom            = []byte("MintDenom")
	KeyMinter               = []byte("Minter")
	DefaultMintDenom string = "uc4e"
	DefaultMinter    Minter = Minter{Start: time.Now(), Periods: []*MintingPeriod{{Position: 1, Type: MintingPeriod_NO_MINTING}}}
) //

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(denom string, minter Minter) Params {
	return Params{MintDenom: denom, Minter: minter}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMintDenom, DefaultMinter)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateDenom),
		paramtypes.NewParamSetPair(KeyMinter, &p.Minter, validateMinter),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateMinter(p.Minter); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateDenom validates the Denom param
func validateDenom(v interface{}) error {
	denom, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if len(denom) == 0 {
		return fmt.Errorf("denom cannot be empty")
	}

	return nil
}

// validateDenom validates the Denom param
func validateMinter(v interface{}) error {
	minter, ok := v.(Minter)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	return minter.Validate()
}
