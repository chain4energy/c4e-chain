package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawAllAvailable(goCtx context.Context, msg *types.MsgWithdrawAllAvailable) (*types.MsgWithdrawAllAvailableResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper
	err := keeper.WithdrawAllAvailable(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawAllAvailableResponse{}, nil
}
