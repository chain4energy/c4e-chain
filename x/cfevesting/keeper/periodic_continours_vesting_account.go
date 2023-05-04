package keeper

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) SetupNewPeriodicContinousVestingAccount(ctx sdk.Context, address sdk.AccAddress, startTime int64, endTime int64) (*cfevestingtypes.PeriodicContinuousVestingAccount, error) {
	baseAccount := k.account.NewAccountWithAddress(ctx, address)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Error("invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
		return nil, errors.Wrapf(c4eerrors.ErrInvalidAccountType, "expected BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), endTime)

	newAcc := cfevestingtypes.NewPeriodicContinuousVestingAccountRaw(baseVestingAccount, startTime)
	newAcc.EndTime = endTime
	return newAcc, nil
}
