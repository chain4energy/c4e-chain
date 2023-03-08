package v2

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyMintDenom    = []byte("MintDenom")
	KeyMinterConfig = []byte("MinterConfig")
) //

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (params *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &params.MintDenom, validateLegacyDenom),
		paramtypes.NewParamSetPair(KeyMinterConfig, &params.MinterConfig, validateLegacyMinters),
	}
}

func (params Params) Validate() error {
	if err := validateLegacyDenom(params.MintDenom); err != nil {
		return err
	}
	if err := validateLegacyMinters(params.MinterConfig); err != nil {
		return err
	}

	return nil
}

func DefaultParams() Params {
	return NewParams(DefaultMintDenom, DefaultMinters)
}
func NewParams(denom string, minterConfig MinterConfig) Params {
	return Params{MintDenom: denom, MinterConfig: minterConfig}
}

// validateDenom validates the Denom param
func validateLegacyDenom(v interface{}) error {
	denom, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if len(denom) == 0 {
		return fmt.Errorf("denom cannot be empty")
	}

	return nil
}

// validateMinters validates Minters
func validateLegacyMinters(v interface{}) error {
	minterConfig, ok := v.(MinterConfig)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", minterConfig)
	}
	err := minterConfig.Validate()
	if err != nil {
		return err
	}

	return nil
}
