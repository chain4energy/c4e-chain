package v3

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyDenom = []byte("Denom")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDenom, &p.Denom, func(v interface{}) error {
			return nil
		}),
	}
}

var (
	DefaultDenom string = "uc4e"
) //

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
