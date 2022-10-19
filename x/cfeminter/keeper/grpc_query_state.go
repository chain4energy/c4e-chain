package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) State(goCtx context.Context, req *types.QueryStateRequest) (*types.QueryStateResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryStateResponse{MinterState: k.GetMinterState(ctx), StateHistory: k.ConvertMinterStateHistory(ctx)}, nil
}
