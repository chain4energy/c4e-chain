package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Inflation(goCtx context.Context, req *types.QueryInflationRequest) (*types.QueryInflationResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	inflation, err := k.GetCurrentInflation(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrGetCurrentInflatio, err.Error())
	}

	return &types.QueryInflationResponse{Inflation: inflation}, nil
}
