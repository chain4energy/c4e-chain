package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyRoutingDistributor                        = []byte("RoutingDistributor")
	DefaultRoutingDistributor RoutingDistributor = RoutingDistributor{}
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(distributor RoutingDistributor) Params {
	return Params{RoutingDistributor: distributor}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultRoutingDistributor)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyRoutingDistributor, &p.RoutingDistributor, validateRoutingDistributor),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateDenom validates the Denom param
func validateRoutingDistributor(v interface{}) error {
	//routingDistributor, ok := v.(RoutingDistributor)
	//if !ok {
	//	return fmt.Errorf("invalid parameter type: %T", v)
	//} //
	return nil
}
