package types

import (
	fmt "fmt"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	DefaultActionTimeWindow time.Duration = 1 * time.Hour
) //

// NewParams creates a new Params instance
func NewParams(actionTimeWindow time.Duration) Params {
	return Params{ActionTimeWindow: actionTimeWindow}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultActionTimeWindow,
	)
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateDenom(p.ActionTimeWindow); err != nil {
		return err
	} //
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateDenom validates the ActionTimeWindow param
func validateDenom(v interface{}) error {
	_, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return nil
}
