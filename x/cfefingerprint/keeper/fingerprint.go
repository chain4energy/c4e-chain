package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/util"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// helper struct to gently store payload links in genesis struct
type PayloadLink struct {
	ReferenceKey   string
	ReferenceValue string
}

// GetAllPayloadLinks returns all payload links
func (k Keeper) GetAllPayloadLinks(ctx sdk.Context) (list []PayloadLink) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PayloadLinkKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val := PayloadLink{
			ReferenceKey:   string(iterator.Key()),
			ReferenceValue: string(iterator.Value()),
		}
		list = append(list, val)
	}
	return
}

func (k Keeper) AppendPayloadLink(ctx sdk.Context, key string, value string) error {
	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PayloadLinkKey))

	store.Set(getStoreKeyBytes(key), []byte(value))

	return nil
}

func (k Keeper) GetPayloadLink(ctx sdk.Context, referenceID string) (string, error) {

	// fetch reference payload link
	referencePayloadLink := util.CalculateHash(referenceID)

	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PayloadLinkKey))
	// get reference payload link value
	referencePayloadLinkValue := store.Get(getStoreKeyBytes(referencePayloadLink))

	// check if there is no document
	if referencePayloadLinkValue == nil {
		return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no payloadlink found")
	}
	// done, fetched link
	return string(referencePayloadLinkValue), nil
}

func (k Keeper) CheckIfPayloadLinkExists(ctx sdk.Context, key string) bool {
	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PayloadLinkKey))
	storedValue := store.Get(getStoreKeyBytes(key))

	return !(storedValue != nil)
}

func getStoreKeyBytes(ID string) []byte {
	return []byte(ID)
}
