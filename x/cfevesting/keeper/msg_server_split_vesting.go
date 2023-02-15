package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SplitVesting(goCtx context.Context, msg *types.MsgSplitVesting) (*types.MsgSplitVestingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	if err := k.splitVestingCoins(ctx, from, msg.ToAddress, msg.Amount); err != nil {
		return nil, err
	}
	return &types.MsgSplitVestingResponse{}, nil
}

func (k msgServer) splitVestingCoins(ctx sdk.Context, from sdk.AccAddress, toAddress string,
	amount sdk.Coins) error {

	if amount.IsZero() || amount.IsAnyNegative() || amount.IsAnyNil() {
		return sdkerrors.Wrapf(types.ErrParam, "split vesting coins - amount must be positive: %s", amount)
	}

	if amount.IsAnyNil() {
		return sdkerrors.Wrapf(types.ErrParam, "split vesting coins - all coins of amount must not be null: %s", amount)
	}

	if err := k.bank.IsSendEnabledCoins(ctx, amount...); err != nil {
		return err
	}

	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return err
	}

	if k.bank.BlockedAddr(to) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := k.account.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Debug("split vesting coins - to account already exists error", "toAddress", to)
		return sdkerrors.Wrapf(types.ErrAlreadyExists, "split vesting coins - account address: %s", to)
	}

	vestingAcc, err := k.UnlockUnbondedContinuousVestingAccountCoins(ctx, from, amount)
	if err != nil {
		return err
	}

	startTime := ctx.BlockTime().Unix()
	if vestingAcc.StartTime > startTime {
		startTime = vestingAcc.StartTime
	}

	if _, err = k.newContinuousVestingAccount(ctx, to, amount, startTime, vestingAcc.EndTime); err != nil {
		return err
	}

	k.bank.SendCoins(ctx, from, to, amount)

	return nil
}
