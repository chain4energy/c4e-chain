package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeySubDistributors     = []byte("SubDistributors")
	DefaultSubDistributors []SubDistributor
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(subDistributors []SubDistributor) Params {
	return Params{SubDistributors: subDistributors}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultSubDistributors)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySubDistributors, &p.SubDistributors, validateSubDistributors),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateSubDistributors(p.SubDistributors); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateSubDistributors validates the SubDistributors param
func validateSubDistributors(v interface{}) error {
	subDistributors, ok := v.([]SubDistributor)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	for _, subDistributor := range subDistributors {
		if error := subDistributor.Validate(); error != nil {
			return error
		}
	}
	return nil
}
