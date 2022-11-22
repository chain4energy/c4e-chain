package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfeairdrop module sentinel errors
var (
	ErrInitialClaimNotFound = sdkerrors.Register(ModuleName, 3, "initial claim not found")
	ErrInitialClaimNotEnabled = sdkerrors.Register(ModuleName, 4, "initial claim not enabled")
	ErrMissionNotFound = sdkerrors.Register(ModuleName, 5, "mission not found")

)


