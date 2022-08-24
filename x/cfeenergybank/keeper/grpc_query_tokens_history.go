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

func (k Keeper) TokensHistoryAll(c context.Context, req *types.QueryAllTokensHistoryRequest) (*types.QueryAllTokensHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokensHistorys []types.TokensHistory
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tokensHistoryStore := prefix.NewStore(store, types.KeyPrefix(types.TokensHistoryKey))

	pageRes, err := query.Paginate(tokensHistoryStore, req.Pagination, func(key []byte, value []byte) error {
		var tokensHistory types.TokensHistory
		if err := k.cdc.Unmarshal(value, &tokensHistory); err != nil {
			return err
		}

		tokensHistorys = append(tokensHistorys, tokensHistory)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokensHistoryResponse{TokensHistory: tokensHistorys, Pagination: pageRes}, nil
}

func (k Keeper) TokensHistory(c context.Context, req *types.QueryGetTokensHistoryRequest) (*types.QueryGetTokensHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tokensHistory, found := k.GetTokensHistory(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTokensHistoryResponse{TokensHistory: tokensHistory}, nil
}

func (k Keeper) TokenHistoryUserAddress(c context.Context, userBlockchainAddress string) (list []types.TokensHistory) {

	ctx := sdk.UnwrapSDKContext(c)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensHistoryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokensHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.UserAddress == userBlockchainAddress {
			list = append(list, val)
		}
	}
	return
}
