package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/cfevesting module sentinel errors
var (
	ErrAmount                          = errors.Register(ModuleName, 2, "wrong amount value")
	ErrAlreadyExists                   = errors.Register(ModuleName, 3, "entity already exists")
	ErrSendCoins                       = errors.Register(ModuleName, 4, "failed to send coins")
	ErrIdenticalAccountsAddresses      = errors.Register(ModuleName, 5, "account addresses cannot be identical")
	ErrGetVestingType                  = errors.Register(ModuleName, 6, "failed to get vesting type")
	ErrAccountNotAllowedToReceiveFunds = errors.Register(ModuleName, 7, "account is not allowed to receive")
	ErrInvalidAccountType              = errors.Register(ModuleName, 8, "invalid account type")
	ErrParsing                         = errors.Register(ModuleName, 9, "failed to parse")
	ErrParam                           = errors.Register(ModuleName, 10, "wrong param value")
)
