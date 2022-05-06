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
	vestings, found := k.GetAccountVestings(ctx, req.Address)
	if !found {
		return &types.QueryVestingPoolsResponse{}, nil
	}

	result := types.QueryVestingPoolsResponse{}
	for _, vesting := range vestings.VestingPools {
		coin := sdk.Coin{Denom: k.GetParams(ctx).Denom, Amount: vesting.Vested}
		withdrawable := CalculateWithdrawable(ctx.BlockTime(), *vesting)
		current := vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
		vestingInfo := types.VestingPoolInfo{
			Id:                  vesting.Id,
			Name:                vesting.Name,
			VestingType:         vesting.VestingType,
			LockStart:           vesting.LockStart,
			LockEnd:             vesting.LockEnd,
			Withdrawable:        withdrawable.String(),
			Vested:              &coin,
			CurrentVestedAmount: current.String(),
			SentAmount:          vesting.Sent.String(),
			TransferAllowed:     vesting.TransferAllowed,
		}
		result.VestingPools = append(result.VestingPools, &vestingInfo)

	}
	return &result, nil
}
