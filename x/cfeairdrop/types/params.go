package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var KeyCampaigns = []byte("Campaigns")

var (
	KeyDenom            = []byte("Denom")
	DefaultDenom string = "uc4e"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(denom string, campaigns []*Campaign) Params {
	return Params{Denom: denom, Campaigns: campaigns}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultDenom, nil)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDenom, &p.Denom, validateDenom),
		paramtypes.NewParamSetPair(KeyCampaigns, &p.Campaigns, validateCampaigns),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateDenom(p.Denom); err != nil {
		return err
	}
	if err := validateCampaigns(p.Campaigns); err != nil {
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

func validateCampaigns(v interface{}) error {
	campaigns, ok := v.([]*Campaign)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	for _, campaign := range campaigns {
		_ = campaign
	}

	return nil
}
