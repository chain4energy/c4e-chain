package keeper

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Vesting(goCtx context.Context, req *types.QueryVestingRequest) (*types.QueryVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	vestings, found := k.GetAccountVestings(ctx, req.Address)
	if !found {
		return &types.QueryVestingResponse{}, nil
	}

	result := types.QueryVestingResponse{}
	for _, vesting := range vestings.Vestings {
		coin := sdk.Coin{Denom: k.GetParams(ctx).Denom, Amount: sdk.NewIntFromUint64(vesting.Vested)}
		withdrawable := CalculateWithdrawable(ctx.BlockHeight(), *vesting)
		current := vesting.Vested - vesting.Withdrawn
		vestingInfo := types.VestingInfo{
			VestingType: vesting.VestingType,
			VestingStartHeight: vesting.VestingStartBlock,
			LockEndHeight: vesting.LockEndBlock,
			VestingEndHeight: vesting.VestingEndBlock,
			Withdrawable: withdrawable.String(),
			DelegationAllowed: vesting.DelegationAllowed,
			Vested: &coin,
			CurrentVestedAmount: strconv.FormatUint(current, 10)}
		result.Vestings = append(result.Vestings, &vestingInfo)
		
	}
	if len(vestings.DelegableAddress) > 0 {
		result.DelegableAddress = vestings.DelegableAddress
	}
	return &result, nil
}
