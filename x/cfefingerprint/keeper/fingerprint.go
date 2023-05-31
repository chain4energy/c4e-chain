package keeper

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"

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

func (k Keeper) CreatePayloadLink(ctx sdk.Context, payloadHash string) (string, error) {
	// get transaction bytes
	inputBytes := ctx.TxBytes()

	// create a TxHash commonly used as a transaction ID
	hash := sha256.Sum256(inputBytes)
	txHash := hex.EncodeToString(hash[:])

	// create referenceId
	referenceId, err := createReferenceID(64, txHash)
	if err != nil {
		return "", sdkerrors.Wrap(sdkerrors.ErrLogic, "failed to generate referenceID")
	}
	ctx.Logger().Debug("calculated referenceID = %s", referenceId)

	// create reference payload link
	referenceKey := util.CalculateHash(referenceId)
	referenceValue := util.CalculateHash(util.HashConcat(referenceId, payloadHash))

	ctx.Logger().Debug("calculated referenceKey = %s / referenceValue = %s", referenceKey, referenceValue)

	// publish reference payload link

	// Check if a Payload Link was already stored at the given key
	if !(k.CheckIfPayloadLinkExists(ctx, referenceKey)) {
		return "", sdkerrors.Wrap(types.ErrAlreadyExists, "data was found at the given key, cannot overwrite present payloadlinks")
	}

	// store payload link
	k.AppendPayloadLink(ctx, referenceKey, referenceValue)

	/*
		As long as referenceId cannot be correlated with a specific account address it can be included in the emitted event.
	*/

	// emit related newPayloadLinkEvent
	newPayloadLinkEvent := &types.NewPayloadLink{
		ReferenceId:    referenceId,
		ReferenceKey:   referenceKey,
		ReferenceValue: referenceKey,
	}
	err = ctx.EventManager().EmitTypedEvent(newPayloadLinkEvent)

	return referenceId, nil
}

func (k Keeper) AppendPayloadLink(ctx sdk.Context, referenceKey string, referenceValue string) {
	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PayloadLinkKey))

	store.Set(getStoreKeyBytes(referenceKey), []byte(referenceValue))
	return
}

func (k Keeper) VerifyPayloadLink(ctx sdk.Context, referenceId, payloadHash string) (bool, error) {
	// fetch data published on ledger
	ledgerPayloadLinkValue, err := k.GetPayloadLink(ctx, referenceId)
	if err != nil {
		k.Logger(ctx).Error("VerifyPayloadLink - failed to get payloadLink from KV store", "error", err)
		return false, err
	}

	// calculate expeced data based on payload hash
	// so called reference value
	expectedPayloadLinkValue := util.CalculateHash(util.HashConcat(referenceId, payloadHash))

	// verify ledger matches the payloadhash
	if expectedPayloadLinkValue == ledgerPayloadLinkValue {
		return true, nil
	} else {
		return false, nil
	}
}

func (k Keeper) GetPayloadLink(ctx sdk.Context, referenceID string) (string, error) {

	// fetch reference payload link key
	referencePayloadLinkKey := util.CalculateHash(referenceID)

	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PayloadLinkKey))
	// get reference payload link value
	referencePayloadLinkValue := store.Get(getStoreKeyBytes(referencePayloadLinkKey))

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

func createReferenceID(length int, txHash string) (string, error) {

	// convert to bytes
	dataBytes := []byte(txHash)
	seed := binary.BigEndian.Uint64(dataBytes)

	rand.Seed(int64(seed))
	b := make([]byte, length+2)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b)[2 : length+2], nil
}
