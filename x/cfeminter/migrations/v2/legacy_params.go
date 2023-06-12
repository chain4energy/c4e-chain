package v2

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyMintDenom    = []byte("MintDenom")
	KeyMinterConfig = []byte("MinterConfig")
)       //
const ( // MintingPeriod types
	NoMintingType              string = "NO_MINTING"
	LinearMintingType          string = "LINEAR_MINTING"
	ExponentialStepMintingType string = "EXPONENTIAL_STEP_MINTING"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (params *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &params.MintDenom, func(interface{}) error { return nil }),
		paramtypes.NewParamSetPair(KeyMinterConfig, &params.MinterConfig, func(interface{}) error { return nil }),
	}
}
