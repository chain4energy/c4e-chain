package subspace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	ParamSet = paramtypes.ParamSet
	// Subspace defines an interface that implements the legacy x/params Subspace
	// type.
	//
	// NOTE: This is used solely for migration of x/params managed parameters.
	Subspace interface {
		GetParamSet(ctx sdk.Context, ps paramtypes.ParamSet)
		GetRaw(ctx sdk.Context, key []byte) []byte
		Set(ctx sdk.Context, key []byte, value interface{})
		HasKeyTable() bool
		WithKeyTable(paramtypes.KeyTable) paramtypes.Subspace
	}
)
