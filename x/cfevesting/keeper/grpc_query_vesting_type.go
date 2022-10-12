package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) VestingType(goCtx context.Context, req *types.QueryVestingTypeRequest) (*types.QueryVestingTypeResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	vestingTypes := k.GetVestingTypes(ctx)
	return &types.QueryVestingTypeResponse{VestingTypes: ConvertVestingTypesToGenesisVestingTypes(&vestingTypes)}, nil
}
