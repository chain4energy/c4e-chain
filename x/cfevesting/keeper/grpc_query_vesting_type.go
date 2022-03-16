package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VestingType(goCtx context.Context, req *types.QueryVestingTypeRequest) (*types.QueryVestingTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	vestingTypes := k.GetVestingTypes(ctx)

	return &types.QueryVestingTypeResponse{VestingTypes: vestingTypes}, nil
}
