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

func (k Keeper) CfevestingV2VestingPools(goCtx context.Context, req *types.QueryCfevestingV2VestingPoolsRequest) (*types.QueryCfevestingV2VestingPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	accountVestingPools, found := k.GetCfevestingV2AccountVestingPools(ctx, req.Address)
	if !found {
		return nil, status.Error(codes.NotFound, "vesting pools not found")
	}

	denom := "uc4e" // TODO: make configurable

	result := types.QueryCfevestingV2VestingPoolsResponse{}
	for _, vesting := range accountVestingPools.VestingPools {
		coin := sdk.Coin{Denom: denom, Amount: vesting.InitiallyLocked}
		withdrawable := CalculateWithdrawable(ctx.BlockTime(), *vesting)
		current := vesting.GetCurrentlyLocked()
		vestingInfo := types.CfevestingV2VestingPoolInfo{
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

	// return &types.QueryCfevestingV2VestingPoolsResponse{}, nil
}

func (k Keeper) GetCfevestingV2AccountVestingPools(ctx sdk.Context, accountAddress string) (accountVestingPools types.C4EVestingV2AccountVestingPools, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.c4eVestingStoreKey), types.C4eVestingV2AccountVestingPoolsKeyPrefix)

	b := store.Get([]byte(accountAddress))
	if b == nil {
		found = false
		return
	}
	found = true
	k.cdc.MustUnmarshal(b, &accountVestingPools)
	return
}

func CalculateWithdrawable(current time.Time, vestingPool types.C4EVestingV2VestingPool) sdk.Int {
	if current.Equal(vestingPool.LockEnd) || current.After(vestingPool.LockEnd) {
		return vestingPool.GetCurrentlyLocked()
	}
	return sdk.ZeroInt()
}
