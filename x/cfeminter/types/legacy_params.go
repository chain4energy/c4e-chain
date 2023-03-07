package types

import (
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
		paramtypes.NewParamSetPair(KeyMintDenom, &params.MintDenom, validateDenom),
		paramtypes.NewParamSetPair(KeyMinterConfig, &params.MinterConfig, validateMinters),
	}
}
