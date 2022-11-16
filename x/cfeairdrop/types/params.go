package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var KeyCampaigns = []byte("Campaigns")

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(campaigns []*Campaign) Params {
	return Params{Campaigns: campaigns}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams([]*Campaign{})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCampaigns, &p.Campaigns, validateCampaigns),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
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
