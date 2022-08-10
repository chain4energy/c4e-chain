package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EnergyTokenAll(c context.Context, req *types.QueryAllEnergyTokenRequest) (*types.QueryAllEnergyTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var energyTokens []types.EnergyToken
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	energyTokenStore := prefix.NewStore(store, types.KeyPrefix(types.EnergyTokenKey))

	pageRes, err := query.Paginate(energyTokenStore, req.Pagination, func(key []byte, value []byte) error {
		var energyToken types.EnergyToken
		if err := k.cdc.Unmarshal(value, &energyToken); err != nil {
			return err
		}

		energyTokens = append(energyTokens, energyToken)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEnergyTokenResponse{EnergyToken: energyTokens, Pagination: pageRes}, nil
}

func (k Keeper) EnergyToken(c context.Context, req *types.QueryGetEnergyTokenRequest) (*types.QueryGetEnergyTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	energyToken, found := k.GetEnergyToken(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEnergyTokenResponse{EnergyToken: energyToken}, nil
}
