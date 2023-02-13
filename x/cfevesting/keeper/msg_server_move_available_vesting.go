package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MoveAvailableVesting(goCtx context.Context, msg *types.MsgMoveAvailableVesting) (*types.MsgMoveAvailableVestingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	to, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return nil, err
	}

	amount := k.bank.LockedCoins(ctx, from)
	if err := k.bank.IsSendEnabledCoins(ctx, amount...); err != nil {
		return nil, err
	}

	if k.bank.BlockedAddr(to) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", msg.ToAddress)
	}

	if acc := k.account.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Debug("new vesting account account already exists error", "toAddress", to)
		return nil, sdkerrors.Wrapf(types.ErrAlreadyExists, "new vesting account - account address: %s", to)
	}

	vestingAcc, err := k.UnlockUnbondedContinuousVestingAccountCoins(ctx, from, amount)
	if err != nil {
		return nil, err
	}

	startTime := ctx.BlockTime().Unix()
	if vestingAcc.StartTime > startTime {
		startTime = vestingAcc.StartTime
	}

	if _, err = k.newContinuousVestingAccount(ctx, to, amount, startTime, vestingAcc.EndTime); err != nil {
		return nil, err
	}

	k.bank.SendCoins(ctx, from, to, amount)

	return &types.MsgMoveAvailableVestingResponse{}, nil
}
