package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/yaml.v2"
)

var maccPerms map[string][]string

func SetMaccPerms(perms map[string][]string) {
	maccPerms = perms
}

var (
	KeySubDistributors     = []byte("SubDistributors")
	DefaultSubDistributors = []SubDistributor{
		{
			Name: "default_distributor",
			Destinations: Destinations{
				PrimaryShare: Account{
					Id:   ValidatorsRewardsCollector,
					Type: ModuleAccount,
				},
				BurnShare: sdk.ZeroDec(),
			},
			Sources: []*Account{
				{
					Id:   "",
					Type: Main,
				},
			},
		},
	}
)

// NewParams creates a new Params instance
func NewParams(subDistributors []SubDistributor) Params {
	return Params{SubDistributors: subDistributors}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultSubDistributors)
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
