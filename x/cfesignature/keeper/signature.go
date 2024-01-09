package keeper

import (
	"cosmossdk.io/errors"
	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/chain4energy/c4e-chain/x/cfesignature/util"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) AppendSignature(ctx sdk.Context, storageKey string, signature types.Signature) string {
	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SignatureKey))

	// Marshal the signature into a slice of bytes.
	appendedValue := k.cdc.MustMarshal(&signature)

	// Insert the signature bytes using storageKey as a key.
	store.Set(getStoreKeyBytes(storageKey), appendedValue)

	return signature.Timestamp
}

func (k Keeper) GetSignature(ctx sdk.Context, storageKey string) (*types.Signature, error) {
	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SignatureKey))

	var signature types.Signature

	signatureBytes := store.Get(getStoreKeyBytes(storageKey)) // Get returns nil if key doesn't exist.

	if signatureBytes == nil {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "failed to get signature")
	}

	k.cdc.MustUnmarshal(signatureBytes, &signature)

	return &signature, nil
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
		return "", errors.Wrap(sdkerrors.ErrInvalidRequest, "no payloadlink found")
	}
	// done, fetched link
	return string(referencePayloadLinkValue), nil
}

func (k Keeper) checkIfPayloadLinkExists(ctx sdk.Context, key string) bool {
	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PayloadLinkKey))
	storedValue := store.Get(getStoreKeyBytes(key))

	return storedValue == nil
}

func getStoreKeyBytes(ID string) []byte {
	return []byte(ID)
}
