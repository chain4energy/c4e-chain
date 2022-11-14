package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func (k Keeper) CreateAirdropAccount(ctx sdk.Context, fromAddress string, toAddress string,
	amount sdk.Coins, startTime int64, endTime int64) error {
	k.Logger(ctx).Debug("create airdrop account", "fromAddress", fromAddress, "toAddress", toAddress,
		"amount", amount, "startTime", startTime, "endTime", endTime)
	ak := k.accountKeeper
	bk := k.bankKeeper

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Error("create vesting account send coins disabled", "error", err.Error())
		return sdkerrors.Wrap(err, "create vesting account - send coins disabled")
	}

	from, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		k.Logger(ctx).Error("create vesting account from-address parsing error", "fromAddress", fromAddress, "error", err.Error())
		return sdkerrors.Wrap(types.ErrSample /* TODO */, sdkerrors.Wrapf(err, "create vesting account - from-address parsing error: %s", fromAddress).Error())
	}
	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		k.Logger(ctx).Error("create vesting account to-address parsing error", "toAddress", toAddress, "error", err.Error())
		return sdkerrors.Wrap(types.ErrSample /* TODO */, sdkerrors.Wrapf(err, "create vesting account - to-address parsing error: %s", toAddress).Error())
	}

	if bk.BlockedAddr(to) {
		k.Logger(ctx).Error("create vesting account account is not allowed to receive funds error", "toAddress", toAddress)
		return sdkerrors.Wrapf(types.ErrSample /* TODO */, "create vesting account - account address: %s", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Error("create vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(types.ErrSample /* TODO */, "create vesting account - account address: %s", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Error("create vesting account invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
		return sdkerrors.Wrapf(types.ErrSample /* TODO */, "create vesting account - expected BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), amount.Sort(), endTime)

	acc := types.NewAirdropVestingAccountRaw(baseVestingAccount, startTime)

	ak.SetAccount(ctx, acc)
	k.Logger(ctx).Debug("create vesting account", "baseAccount", baseVestingAccount.BaseAccount, "originalVesting",
		baseVestingAccount.OriginalVesting, "delegatedFree", baseVestingAccount.DelegatedFree, "delegatedVesting",
		baseVestingAccount.DelegatedVesting, "endTime", baseVestingAccount.EndTime, "startTime", startTime)
	// err = ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
	// 	Address: acc.Address,
	// })
	// if err != nil {
	// 	k.Logger(ctx).Error("new vestig account emit event error", "error", err.Error())
	// }

	err = bk.SendCoins(ctx, from, to, amount)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account send coins to vesting account error", "fromAddress", fromAddress, "toAddress", toAddress,
			"amount", amount, "error", err.Error())
		return sdkerrors.Wrap(types.ErrSample /* TODO */, sdkerrors.Wrapf(err,
			"create vesting account - send coins to vesting account error (from: %s, to: %s, amount: %s)", fromAddress, toAddress, amount).Error())
	}
	return nil
}
