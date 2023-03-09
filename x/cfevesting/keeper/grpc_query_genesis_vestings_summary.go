package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GenesisVestingsSummary(goCtx context.Context, req *types.QueryGenesisVestingsSummaryRequest) (*types.QueryGenesisVestingsSummaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	summary, err := k.createVestingsSummary(ctx, true)
	result := types.QueryGenesisVestingsSummaryResponse(*summary)
	return &result, err
}
