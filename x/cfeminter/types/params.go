package types

import (
	fmt "fmt"
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyMintDenom     = []byte("MintDenom")
	KeyMinters       = []byte("Minters")
	KeyStartTime     = []byte("StartTime")
	DefaultMintDenom = "uc4e"
	DefaultStartTime = time.Now()
	DefaultMinters   = []*Minter{{
		SequenceId: 1,
		Type:       NO_MINTING,
	},
	}
) //

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(denom string, startTime time.Time, minters []*Minter) Params {
	return Params{MintDenom: denom, StartTime: startTime, Minters: minters}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMintDenom, DefaultStartTime, DefaultMinters)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateDenom),
		paramtypes.NewParamSetPair(KeyStartTime, &p.StartTime, validateStartTime),
		paramtypes.NewParamSetPair(KeyMinters, &p, validateMinters),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinters(p); err != nil {
		return err
	}
	if err := validateDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateStartTime(p.StartTime); err != nil {
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

// validateMinters validates the Denom param
func validateMinters(v interface{}) error {
	params, ok := v.(Params)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	return params.ValidateMinters()
}

// validateStartTime validates the StartTime param //TODO: add logic here
func validateStartTime(v interface{}) error {
	minter, ok := v.(Minter)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	return minter.Validate()
}
