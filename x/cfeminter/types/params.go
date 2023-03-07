package types

import (
	fmt "fmt"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	DefaultMintDenom = "uc4e"
	DefaultMinters   = MinterConfig{
		StartTime: time.Now(),
		Minters: []*Minter{
			{
				SequenceId: 1,
				Type:       NoMintingType,
			},
		},
	}
) //

// NewParams creates a new Params instance
func NewParams(denom string, minterConfig MinterConfig) Params {
	return Params{MintDenom: denom, MinterConfig: minterConfig}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMintDenom, DefaultMinters)
}

// Validate validates the set of params
func (params Params) Validate() error {
	if err := validateDenom(params.MintDenom); err != nil {
		return err
	}
	if err := validateMinters(params.MinterConfig); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (params Params) String() string {
	out, _ := yaml.Marshal(params)
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

// validateMinters validates Minters
func validateMinters(v interface{}) error {
	minterConfig, ok := v.(MinterConfig)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", minterConfig)
	}
	err := minterConfig.Validate()
	if err != nil {
		return err
	}

	return nil
}
