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
		return nil, sdkerrors.Wrapf(err, "move available vesting - error parsing from address: %s", msg.FromAddress)
	}
	amount := k.bank.LockedCoins(ctx, from)

	if err := k.splitVestingCoins(ctx, from, msg.ToAddress, amount); err != nil {
		return nil, sdkerrors.Wrap(err, "move available vesting")
	}
	return &types.MsgMoveAvailableVestingResponse{}, nil

}
