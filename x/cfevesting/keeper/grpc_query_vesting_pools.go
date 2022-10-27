package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VestingPools(goCtx context.Context, req *types.QueryVestingPoolsRequest) (*types.QueryVestingPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	accountVestingPools, found := k.GetAccountVestingPools(ctx, req.Address)
	if !found {
		return nil, status.Error(codes.NotFound, "vesting pools not found")
	}

	result := types.QueryVestingPoolsResponse{}
	for _, vesting := range accountVestingPools.VestingPools {
		coin := sdk.Coin{Denom: k.GetParams(ctx).Denom, Amount: vesting.InitiallyLocked}
		withdrawable := CalculateWithdrawable(ctx.BlockTime(), *vesting)
		current := vesting.GetCurrentlyLocked()
		vestingInfo := types.VestingPoolInfo{
			Name:            vesting.Name,
			VestingType:     vesting.VestingType,
			LockStart:       vesting.LockStart,
			LockEnd:         vesting.LockEnd,
			Withdrawable:    withdrawable.String(),
			InitiallyLocked: &coin,
			CurrentlyLocked: current.String(),
			SentAmount:      vesting.Sent.String(),
		}
		result.VestingPools = append(result.VestingPools, &vestingInfo)

	}
	return &result, nil
}
