package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MoveAvailableVesting(goCtx context.Context, msg *types.MsgMoveAvailableVesting) (*types.MsgMoveAvailableVestingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}
	amount := k.bank.LockedCoins(ctx, from)

	if err := k.splitVestingCoins(ctx, from, msg.ToAddress, amount); err != nil {
		return nil, err
	}
	return &types.MsgMoveAvailableVestingResponse{}, nil

}
