package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfevesting module sentinel errors
var (
	ErrAmount                          = sdkerrors.Register(ModuleName, 2, "wrong amount value")
	ErrAlreadyExists                   = sdkerrors.Register(ModuleName, 3, "entity already exists")
	ErrSendCoins                       = sdkerrors.Register(ModuleName, 4, "failed to send coins")
	ErrIdenticalAccountsAddresses      = sdkerrors.Register(ModuleName, 5, "account addresses cannot be identical")
	ErrGetVestingType                  = sdkerrors.Register(ModuleName, 6, "failed to get vesting type")
	ErrAccountNotAllowedToReceiveFunds = sdkerrors.Register(ModuleName, 7, "account is not allowed to receive")
	ErrInvalidAccountType              = sdkerrors.Register(ModuleName, 8, "invalid account type")
	ErrParsing                         = sdkerrors.Register(ModuleName, 9, "failed to parse")
	ErrParam                           = sdkerrors.Register(ModuleName, 10, "wrong param value")
	ErrStartTimeAfterEndTime           = sdkerrors.Register(ModuleName, 11, "start time cannot be after end time")
)
