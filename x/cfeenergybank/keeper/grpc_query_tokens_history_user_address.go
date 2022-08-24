package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) TokensHistoryUserAddress(goCtx context.Context, req *types.QueryTokensHistoryUserAddressRequest) (*types.QueryTokensHistoryUserAddressResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokensHistoryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	var tokenHistory []types.TokensHistory
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokensHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.UserAddress == req.UserBlockchainAddress {
			tokenHistory = append(tokenHistory, val)
		}
	}
	return &types.QueryTokensHistoryUserAddressResponse{TokensHistory: tokenHistory}, nil
}
