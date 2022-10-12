package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfeminter module sentinel errors
var (
	ErrInvalidRequest     = sdkerrors.Register(ModuleName, 2, "invalid request error")
	ErrGetCurrentInflatio = sdkerrors.Register(ModuleName, 3, "failed to get current inflatio")
)
