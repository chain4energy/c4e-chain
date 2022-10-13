package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfevesting module sentinel errors
var (
	ErrInvalidRequest                    = sdkerrors.Register(ModuleName, 2, "invalid request error")
	ErrWithdrawAllAvailable              = sdkerrors.Register(ModuleName, 3, "failed to withdraw all available")
	ErrVestingTypeNotFound               = sdkerrors.Register(ModuleName, 4, "vesting type not found")
	ErrVestingPoolAountEqualsZero        = sdkerrors.Register(ModuleName, 5, "vesting pool amount equals zero")
	ErrVestingPoolNameAlreadyExists      = sdkerrors.Register(ModuleName, 6, "vesting pool name already exists")
	ErrSendigCoinsToVestingPool          = sdkerrors.Register(ModuleName, 7, "failed to send coins to vesting pool")
	ErrSendigCoinsToVestingAccount       = sdkerrors.Register(ModuleName, 8, "failed to send coins to vesting account")
	ErrSendigCoinsFromModuleToAccount    = sdkerrors.Register(ModuleName, 9, "failed to send coins from module to account")
	ErrVestingPoolsNotFound              = sdkerrors.Register(ModuleName, 10, "no vestings found")
	ErrIdenticalAccountsAddresses        = sdkerrors.Register(ModuleName, 11, "account addresses cannot be identical")
	ErrVestingPoolNotFound               = sdkerrors.Register(ModuleName, 12, "vesting pool not found")
	ErrVestingAvailableSmallerThanAmount = sdkerrors.Register(ModuleName, 13, "vesting available is smaller than amount")
	ErrGetVestingType                    = sdkerrors.Register(ModuleName, 14, "failed to get vesting type")
	ErrSendEnabledCoins                  = sdkerrors.Register(ModuleName, 15, "send enabled coins error")
	ErrAccountNotAllowedToReceiveFunds   = sdkerrors.Register(ModuleName, 16, "account is not allowed to receive")
	ErrVestingAccountExists              = sdkerrors.Register(ModuleName, 17, "account account already exists")
	ErrInvalidAccountType                = sdkerrors.Register(ModuleName, 18, "invalid account type")
	ErrNoVestingPoolsFound               = sdkerrors.Register(ModuleName, 19, "no vesting pools found")
	ErrParsing                           = sdkerrors.Register(ModuleName, 20, "failed to parse address")
)
