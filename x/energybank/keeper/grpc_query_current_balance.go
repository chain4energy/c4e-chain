package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"google.golang.org/grpc/status"
	"time"

	"github.com/chain4energy/c4e-chain/x/energybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
)

func (k Keeper) CurrentBalance(goCtx context.Context, req *types.QueryCurrentBalanceRequest) (*types.QueryCurrentBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	tokenParams, isFound := k.GetTokenParams(ctx, req.TokenName)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Token with this name doesn't exist.")
	}
	store := ctx.KVStore(k.storeKey)

	energyTokenStore := prefix.NewStore(store, []byte(types.EnergyTokenKey))

	iterator := sdk.KVStorePrefixIterator(energyTokenStore, []byte{})

	defer iterator.Close()
	var balance uint64 = 0

	if tokenParams.BurningType == "linear" {

		for ; iterator.Valid(); iterator.Next() {
			var val types.EnergyToken
			k.cdc.MustUnmarshal(iterator.Value(), &val)
			if val.UserAddress == req.UserAddress && val.Name == req.TokenName {

				timeNow := time.Now().Unix()
				timeDiffInDays := (int(timeNow) - int(val.CreatedAt)) / 86400
				burningPercentage := 1 / float32(tokenParams.BurningTime)
				currentVal := float32(val.Amount) - (float32(val.Amount)*burningPercentage)*float32(timeDiffInDays)

				if currentVal > 0 {
					balance += uint64(currentVal)
				}
			}
		}
	}

	if tokenParams.BurningType == "step" {
		for ; iterator.Valid(); iterator.Next() {
			var val types.EnergyToken
			k.cdc.MustUnmarshal(iterator.Value(), &val)

			if val.UserAddress == req.UserAddress && val.Name == req.TokenName {

				deadline := int(val.CreatedAt + tokenParams.BurningTime*86400)

				secondsDiff := deadline - (int(time.Now().Unix()))

				if secondsDiff > 0 {
					balance += val.Amount
				}
			}
		}
	}

	return &types.QueryCurrentBalanceResponse{Balance: balance}, nil
}
