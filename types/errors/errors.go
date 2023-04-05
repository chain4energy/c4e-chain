package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// C4eCodespace is the codespace for all commen C4E errors
const C4eCodespace = "c4e"

// x/cfeairdrop module sentinel errors
var (
	ErrAmount                          = sdkerrors.Register(C4eCodespace, 2, "wrong amount value")
	ErrAlreadyExists                   = sdkerrors.Register(C4eCodespace, 3, "entity already exists")
	ErrSendCoins                       = sdkerrors.Register(C4eCodespace, 4, "failed to send coins")
	ErrAccountNotAllowedToReceiveFunds = sdkerrors.Register(C4eCodespace, 5, "account is not allowed to receive")
	ErrInvalidAccountType              = sdkerrors.Register(C4eCodespace, 6, "invalid account type")
	ErrParsing                         = sdkerrors.Register(C4eCodespace, 7, "failed to parse")
	ErrParam                           = sdkerrors.Register(C4eCodespace, 8, "wrong param value")
	ErrNotExists                       = sdkerrors.Register(C4eCodespace, 9, "entity does not exist")
)
