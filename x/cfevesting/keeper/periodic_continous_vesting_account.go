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

func (k Keeper) SendToPeriodicContinuousVestingAccountFromModule(ctx sdk.Context, moduleName string, userAddress string, amount sdk.Coins,
	free sdk.Dec, startTime int64, endTime int64) (periodId uint64, periodExists bool, err error) {
	k.Logger(ctx).Debug("send to periodic continous vesting account from module", "moduleName", moduleName,
		"userAddress", userAddress, "amount", amount, "free", free, "startTime", startTime, "endTime", endTime)

	userAccAddress, err := sdk.AccAddressFromBech32(userAddress)
	if err != nil {
		return 0, false, errors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: ", userAddress)
	}

	if k.bank.BlockedAddr(userAccAddress) {
		return 0, false, errors.Wrapf(sdkerrors.ErrUnauthorized, "account address: %s is not allowed to receive funds error", userAddress)
	}
	moduleBalances := k.bank.GetAllBalances(ctx, k.account.GetModuleAddress(moduleName))
	if !amount.IsAllLTE(moduleBalances) {
		return 0, false, errors.Wrapf(sdkerrors.ErrInsufficientFunds, "module balance is too small (%s < %s)", moduleBalances, amount)
	}
	if err = k.bank.IsSendEnabledCoins(ctx, amount...); err != nil {
		return 0, false, err
	}
	periodicContinousVestingAccount, err := k.getOrCreatePeriodicContinousVestingAccount(ctx, userAccAddress, startTime, endTime)
	if err != nil {
		return 0, false, err
	}

	var vestingPeriodCoins sdk.Coins
	for _, coin := range amount {
		decimalAmount := sdk.NewDecFromInt(coin.Amount)
		vestingPeriodAmount := decimalAmount.Sub(decimalAmount.Mul(free)).TruncateInt()
		vestingPeriodCoins = vestingPeriodCoins.Add(sdk.NewCoin(coin.Denom, vestingPeriodAmount))
	}

	if len(vestingPeriodCoins) > 0 {
		periodId = periodicContinousVestingAccount.AddNewContinousVestingPeriod(startTime, endTime, vestingPeriodCoins)
		periodExists = true
	}

	k.account.SetAccount(ctx, periodicContinousVestingAccount)
	if err = k.bank.SendCoinsFromModuleToAccount(ctx, moduleName, userAccAddress, amount); err != nil {
		return 0, false, err
	}

	k.Logger(ctx).Debug("send to periodic continous vesting account from module ret", "periodId", periodId, "periodExists", periodExists)
	return periodId, periodExists, nil
}

func (k Keeper) getOrCreatePeriodicContinousVestingAccount(ctx sdk.Context, userAddress sdk.AccAddress, startTime,
	endTime int64) (*types.PeriodicContinuousVestingAccount, error) {
	account := k.account.GetAccount(ctx, userAddress)

	periodicContinuousVestingAccount, ok := account.(*types.PeriodicContinuousVestingAccount)
	if ok {
		return periodicContinuousVestingAccount, nil
	}

	if account == nil {
		return k.newPeriodicContinousVestingAccount(ctx, userAddress, startTime, endTime)
	} else {
		// If there was a base account previously, migrate it to a periodicVestingAccount. This operation
		// is required because the account to which we want to transfer tokens may have been created earlier,
		// for example, if the feegrant (in the cfeclaim campaign) is set to a value greater than zero.
		baseAccount, ok := account.(*authtypes.BaseAccount)
		if !ok {
			return nil, errors.Wrapf(c4eerrors.ErrInvalidAccountType, "account already exists and is not of PeriodicContinuousVestingAccount nor BaseAccount type, got: %T", account)
		}
		return k.newPeriodicContinousVestingAccountFromBaseAccount(baseAccount, startTime, endTime)
	}
}

func (k Keeper) newPeriodicContinousVestingAccount(ctx sdk.Context, address sdk.AccAddress, startTime int64,
	endTime int64) (*types.PeriodicContinuousVestingAccount, error) {
	account := k.account.NewAccountWithAddress(ctx, address)
	baseAccount, ok := account.(*authtypes.BaseAccount)

	if !ok {
		return nil, errors.Wrapf(c4eerrors.ErrInvalidAccountType, "expected BaseAccount, got: %T", account)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount, sdk.NewCoins(), endTime)
	newAcc := types.NewPeriodicContinuousVestingAccountRaw(baseVestingAccount, startTime)
	newAcc.EndTime = endTime
	return newAcc, nil
}

func (k Keeper) newPeriodicContinousVestingAccountFromBaseAccount(baseAccount *authtypes.BaseAccount, startTime int64,
	endTime int64) (*types.PeriodicContinuousVestingAccount, error) {
	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount, sdk.NewCoins(), endTime)

	newAcc := types.NewPeriodicContinuousVestingAccountRaw(baseVestingAccount, startTime)
	newAcc.EndTime = endTime
	return newAcc, nil
}
