package keeper

import (
	"context"
	"time"

	"github.com/chain4energy/c4e-chain/x/compatibility/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CfevestingV1VestingPools(goCtx context.Context, req *types.QueryCfevestingV1VestingPoolsRequest) (*types.QueryCfevestingV1VestingPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	vestings, found := k.GetCfevestingV1AccountVestings(ctx, req.Address)
	if !found {
		return &types.QueryCfevestingV1VestingPoolsResponse{}, nil
	}

	denom := "uc4e" // TODO: make configurable

	result := types.QueryCfevestingV1VestingPoolsResponse{}
	for _, vesting := range vestings.VestingPools {
		coin := sdk.Coin{Denom: denom, Amount: vesting.Vested}
		withdrawable := CfevestingV1CalculateWithdrawable(ctx.BlockTime(), *vesting)
		current := vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
		vestingInfo := types.CfevestingV1VestingPoolInfo{
			Id:                  vesting.Id,
			Name:                vesting.Name,
			VestingType:         vesting.VestingType,
			LockStart:           vesting.LockStart,
			LockEnd:             vesting.LockEnd,
			Withdrawable:        withdrawable.String(),
			Vested:              &coin,
			CurrentVestedAmount: current.String(),
			SentAmount:          vesting.Sent.String(),
		}
		result.VestingPools = append(result.VestingPools, &vestingInfo)

	}
	return &result, nil

	// return &types.QueryCfevestingV1VestingPoolsResponse{}, nil
}

// get the vesting types
func (k Keeper) GetCfevestingV1AccountVestings(ctx sdk.Context, accountAddress string) (accountVestings types.C4EVestingV1AccountVestings, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.c4eVestingStoreKey), types.C4eVestingV1AccountVestingsKeyPrefix)

	b := store.Get([]byte(accountAddress))
	if b == nil {
		found = false
		return
	}
	found = true
	k.cdc.MustUnmarshal(b, &accountVestings)
	return
}

func CfevestingV1CalculateWithdrawable(current time.Time, vesting types.C4EVestingV1VestingPool) sdk.Int {
	if current.Equal(vesting.LockEnd) || current.After(vesting.LockEnd) {
		return vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	}
	return sdk.ZeroInt()
}
