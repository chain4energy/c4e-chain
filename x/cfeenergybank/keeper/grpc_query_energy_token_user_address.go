package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EnergyTokenUserAddress(goCtx context.Context, req *types.QueryEnergyTokenUserAddressRequest) (*types.QueryEnergyTokenUserAddressResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var energyTokens []*types.EnergyToken

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)

	energyTokenStore := prefix.NewStore(store, []byte(types.EnergyTokenKey))

	pageRes, err := query.Paginate(energyTokenStore, req.Pagination, func(key []byte, value []byte) error {
		var energyToken types.EnergyToken
		if err := k.cdc.Unmarshal(value, &energyToken); err != nil {
			return err
		}
		if energyToken.UserAddress == req.UserAddress {
			energyTokens = append(energyTokens, &energyToken)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryEnergyTokenUserAddressResponse{EnergyToken: energyTokens, Pagination: pageRes}, nil
}
