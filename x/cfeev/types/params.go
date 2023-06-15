package types

import (
	fmt "fmt"

	"gopkg.in/yaml.v2"
)

var (
	DefaultDenom = "uc4e"
)

// NewParams creates a new Params instance
func NewParams(denom string) Params {
	return Params{Denom: denom}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultDenom,
	)
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateDenom(p.Denom); err != nil {
		return err
	} //
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
