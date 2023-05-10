package keeper

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) SendToPeriodicContinuousVestingAccountFromModule(ctx sdk.Context, moduleName string, userAddress string, amount sdk.Coins, startTime int64, endTime int64) error {
	userAccAddress, err := sdk.AccAddressFromBech32(userAddress)
	if err != nil {
		return errors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: ", userAddress)
	}

	if k.bank.BlockedAddr(userAccAddress) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "account address: %s is not allowed to receive funds error", userAddress)
	}

	if err != nil {
		return err
	}

	periodicContinousVestingAccount, err := k.getOrCreatePeriodicContinousVestingAccount(ctx, userAccAddress, startTime, endTime)
	if err != nil {
		return err
	}

	periodicContinousVestingAccount.AddNewContinousVestingPeriod(startTime, endTime, amount)

	k.account.SetAccount(ctx, periodicContinousVestingAccount)
	return k.bank.SendCoinsFromModuleToAccount(ctx, moduleName, userAccAddress, amount)
}

func (k Keeper) getOrCreatePeriodicContinousVestingAccount(ctx sdk.Context, claimerAddress sdk.AccAddress, startTime, endTime int64) (*types.PeriodicContinuousVestingAccount, error) {
	claimerAccount := k.account.GetAccount(ctx, claimerAddress)
	periodicContinuousVestingAccount, ok := claimerAccount.(*types.PeriodicContinuousVestingAccount)
	if !ok {
		var err error
		claimerAccount, err = k.SetupNewPeriodicContinousVestingAccount(ctx, claimerAddress, startTime, endTime)
		if err != nil {
			return nil, err
		}
		if claimerAccount == nil {
			return nil, errors.Wrapf(c4eerrors.ErrNotExists, "send to claim account - account does not exist: %s", claimerAddress)
		}
		periodicContinuousVestingAccount, ok = claimerAccount.(*types.PeriodicContinuousVestingAccount)
		if !ok {
			return nil, errors.Wrapf(c4eerrors.ErrInvalidAccountType, "send to claim account - expected PeriodicContinuousVestingAccount, got: %T", claimerAccount)
		}
	}

	return periodicContinuousVestingAccount, nil
}

func (k Keeper) SetupNewPeriodicContinousVestingAccount(ctx sdk.Context, address sdk.AccAddress, startTime int64, endTime int64) (*types.PeriodicContinuousVestingAccount, error) {
	baseAccount := k.account.NewAccountWithAddress(ctx, address)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Error("invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
		return nil, errors.Wrapf(c4eerrors.ErrInvalidAccountType, "expected BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), endTime)

	newAcc := types.NewPeriodicContinuousVestingAccountRaw(baseVestingAccount, startTime)
	newAcc.EndTime = endTime
	return newAcc, nil
}
