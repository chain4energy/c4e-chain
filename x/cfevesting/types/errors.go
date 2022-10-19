package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfevesting module sentinel errors
var (
	ErrInvalidRequest                    = sdkerrors.Register(ModuleName, 2, "invalid request error")
	ErrAmount							 = sdkerrors.Register(ModuleName, 3, "wrong amount value")
	ErrAlreadyExists      				 = sdkerrors.Register(ModuleName, 4, "entity already exists")
	ErrSendCoins						 = sdkerrors.Register(ModuleName, 5, "failed to send coins")
	ErrIdenticalAccountsAddresses        = sdkerrors.Register(ModuleName, 6, "account addresses cannot be identical")
	ErrGetVestingType                    = sdkerrors.Register(ModuleName, 7, "failed to get vesting type")
	ErrAccountNotAllowedToReceiveFunds   = sdkerrors.Register(ModuleName, 8, "account is not allowed to receive")
	ErrInvalidAccountType                = sdkerrors.Register(ModuleName, 9, "invalid account type")
	ErrParsing                           = sdkerrors.Register(ModuleName, 10, "failed to parse")
)
