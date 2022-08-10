package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/energybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CurrentBalance(goCtx context.Context, req *types.QueryCurrentBalanceRequest) (*types.QueryCurrentBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryCurrentBalanceResponse{}, nil
}
