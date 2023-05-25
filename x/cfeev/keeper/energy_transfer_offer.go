package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetEnergyTransferOfferCount get the total number of energyTransferOffer
func (k Keeper) GetEnergyTransferOfferCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EnergyTransferOfferCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEnergyTransferOfferCount set the total number of energyTransferOffer
func (k Keeper) SetEnergyTransferOfferCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EnergyTransferOfferCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEnergyTransferOffer appends a energyTransferOffer in the store with a new id and update the count
func (k Keeper) AppendEnergyTransferOffer(
	ctx sdk.Context,
	energyTransferOffer types.EnergyTransferOffer,
) uint64 {
	// Create the energyTransferOffer
	count := k.GetEnergyTransferOfferCount(ctx)

	// Set the ID of the appended value
	energyTransferOffer.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTransferOfferKey))
	appendedValue := k.cdc.MustMarshal(&energyTransferOffer)
	store.Set(GetEnergyTransferOfferIDBytes(energyTransferOffer.Id), appendedValue)

	// Update energyTransferOffer count
	k.SetEnergyTransferOfferCount(ctx, count+1)

	return count
}

// SetEnergyTransferOffer set a specific energyTransferOffer in the store
func (k Keeper) SetEnergyTransferOffer(ctx sdk.Context, energyTransferOffer types.EnergyTransferOffer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTransferOfferKey))
	b := k.cdc.MustMarshal(&energyTransferOffer)
	store.Set(GetEnergyTransferOfferIDBytes(energyTransferOffer.Id), b)
}

// GetEnergyTransferOffer returns a energyTransferOffer from its id
func (k Keeper) GetEnergyTransferOffer(ctx sdk.Context, id uint64) (val types.EnergyTransferOffer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTransferOfferKey))
	b := store.Get(GetEnergyTransferOfferIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEnergyTransferOffer removes a energyTransferOffer from the store
func (k Keeper) RemoveEnergyTransferOffer(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTransferOfferKey))
	store.Delete(GetEnergyTransferOfferIDBytes(id))
}

// GetAllEnergyTransferOffer returns all energyTransferOffer
func (k Keeper) GetAllEnergyTransferOffer(ctx sdk.Context) (list []types.EnergyTransferOffer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTransferOfferKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EnergyTransferOffer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetTransferOfferByChargerId(ctx sdk.Context, chargerId string) (val types.EnergyTransferOffer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTransferOfferKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EnergyTransferOffer
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if val.ChargerId == chargerId {
			return val, true
		}

	}

	return val, false
}

// GetEnergyTransferOfferIDBytes returns the byte representation of the ID
func GetEnergyTransferOfferIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetEnergyTransferOfferIDFromBytes returns ID in uint64 format from a byte array
func GetEnergyTransferOfferIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
