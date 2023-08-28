package types

import (
	fmt "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	DefaultActionTimeWindow        time.Duration = 1 * time.Hour
	DefaultMarketplaceFee          sdk.Dec       = sdk.MustNewDecFromStr("0.05")
	DefaultAuthorityFee            sdk.Dec       = sdk.MustNewDecFromStr("0.05")
	DefaultMarketplaceOwnerAddress string        = "c4e1yd97thfrcftqcd8htn2v8z90h502xl3f6rf6p7"
) //

// NewParams creates a new Params instance
func NewParams(actionTimeWindow time.Duration, marketplaceFee sdk.Dec, authorityFee sdk.Dec, marketplaceOwner string) Params {
	return Params{
		ActionTimeWindow:        actionTimeWindow,
		MarketplaceFee:          marketplaceFee,
		AuthorityFee:            authorityFee,
		MarketplaceOwnerAddress: marketplaceOwner,
	}

}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultActionTimeWindow,
		DefaultMarketplaceFee,
		DefaultAuthorityFee,
		DefaultMarketplaceOwnerAddress,
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
