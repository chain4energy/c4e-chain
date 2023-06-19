package keeper

import (
	"cosmossdk.io/errors"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetAllPayloadLinks returns all payload links
func (k Keeper) GetAllPayloadLinks(ctx sdk.Context) (list []*types.GenesisPayloadLink) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PayloadLinkKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val := &types.GenesisPayloadLink{
			ReferenceKey:   string(iterator.Key()),
			ReferenceValue: string(iterator.Value()),
		}
		list = append(list, val)
	}
	return
}

func (k Keeper) AppendPayloadLink(ctx sdk.Context, referenceKey string, referenceValue string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PayloadLinkKey)
	store.Set(types.GetStringKey(referenceKey), types.GetStringKey(referenceValue))
	return
}

func (k Keeper) GetPayloadLink(ctx sdk.Context, referenceKey string) (string, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PayloadLinkKey)
	payloadLink := store.Get(types.GetStringKey(referenceKey))
	if payloadLink == nil {
		return "", false
	}

	return string(payloadLink), true
}

func (k Keeper) GetPayloadLinkByReferenceId(ctx sdk.Context, referenceId string) (string, bool) {
	referenceKey := types.CalculateHash(referenceId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PayloadLinkKey)
	payloadLink := store.Get(types.GetStringKey(referenceKey))
	if payloadLink == nil {
		return "", false
	}

	return string(payloadLink), true
}

func (k Keeper) MustGetPayloadLinkByReferenceId(ctx sdk.Context, referenceId string) (string, error) {
	payloadLink, found := k.GetPayloadLinkByReferenceId(ctx, referenceId)
	if !found {
		return "", errors.Wrapf(sdkerrors.ErrNotFound, "payload link not found for reference id %s", referenceId)
	}
	return payloadLink, nil
}
