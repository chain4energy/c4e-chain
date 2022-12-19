package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

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
					Type: MODULE_ACCOUNT,
				},
				BurnShare: sdk.ZeroDec(),
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
	fmt.Println("SUBDISTRIBUTORS PRINT")
	fmt.Println(subDistributors)
	fmt.Println(len(subDistributors))
	fmt.Println(&subDistributors == nil)
	fmt.Println(subDistributors)
	fmt.Println(&subDistributors[0])
	fmt.Println(subDistributors[0])
	fmt.Println(subDistributors[1].Name)
	fmt.Println(subDistributors[1].Name == "")
	fmt.Println(subDistributors[1].Destinations)
	fmt.Println(subDistributors[1].Destinations.BurnShare)
	fmt.Println(subDistributors[1].Sources)
	//if subDistributors[0] == nil {
	//
	//}
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
