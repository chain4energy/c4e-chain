package types

import (
	"cosmossdk.io/errors"
)

// C4eCodespace is the codespace for all commen C4E errors
const C4eCodespace = "c4e"

// x/cfeclaim module sentinel errors
var (
	ErrAmount                          = errors.Register(C4eCodespace, 2, "wrong amount value")
	ErrAlreadyExists                   = errors.Register(C4eCodespace, 3, "entity already exists")
	ErrSendCoins                       = errors.Register(C4eCodespace, 4, "failed to send coins")
	ErrAccountNotAllowedToReceiveFunds = errors.Register(C4eCodespace, 5, "account is not allowed to receive")
	ErrInvalidAccountType              = errors.Register(C4eCodespace, 6, "invalid account type")
	ErrParsing                         = errors.Register(C4eCodespace, 7, "failed to parse")
	ErrParam                           = errors.Register(C4eCodespace, 8, "wrong param value")
	ErrNotExists                       = errors.Register(C4eCodespace, 9, "entity does not exist")
)
