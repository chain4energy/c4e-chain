package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeySubDistributors     = []byte("SubDistributors")
	DefaultSubDistributors = []SubDistributor{
		{
			Name: "default_distributor",
			Destination: Destination{
				Account: Account{
					Id:   ValidatorsRewardsCollector,
					Type: MODULE_ACCOUNT,
				},
				BurnShare: &BurnShare{
					Percent: sdk.MustNewDecFromStr("0"),
				},
			},
			Sources: []*Account{
				{
					Id:   "",
					Type: MAIN,
				},
			},
		},
	}
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
		if err := subDistributor.Validate(); err != nil {
			return err
		}
	}
	err := ValidateSubDistributors(subDistributors)
	if err != nil {
		return err
	}

	return nil
}
