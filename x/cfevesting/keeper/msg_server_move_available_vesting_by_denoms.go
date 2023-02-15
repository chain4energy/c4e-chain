package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MoveAvailableVestingByDenoms(goCtx context.Context, msg *types.MsgMoveAvailableVestingByDenoms) (*types.MsgMoveAvailableVestingByDenomsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}
	locked := k.bank.LockedCoins(ctx, from)
	amount := sdk.NewCoins()
	for _, denom := range msg.Denoms {
		denAmount := locked.AmountOf(denom)
		if denAmount.IsPositive() {
			amount.Add(sdk.NewCoin(denom, denAmount))
		}
	}

	if err := k.splitVestingCoins(ctx, from, msg.ToAddress, amount); err != nil {
		return nil, err
	}
	return &types.MsgMoveAvailableVestingByDenomsResponse{}, nil

}
