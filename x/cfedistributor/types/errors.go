package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfedistributor module sentinel errors
var (
	ErrInvalidRequest = sdkerrors.Register(ModuleName, 2, "invalid request error")
)
