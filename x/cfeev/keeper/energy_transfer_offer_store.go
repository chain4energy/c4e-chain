package keeper

import (
	"cosmossdk.io/errors"
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/types/util"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetEnergyTransferOfferCount get the total number of energyTransferOffer
func (k Keeper) GetEnergyTransferOfferCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.EnergyTransferOfferCountKey
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
	byteKey := types.EnergyTransferOfferCountKey
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

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferOfferKey)
	appendedValue := k.cdc.MustMarshal(&energyTransferOffer)
	store.Set(util.GetUint64Key(energyTransferOffer.Id), appendedValue)

	// Update energyTransferOffer count
	k.SetEnergyTransferOfferCount(ctx, count+1)

	return count
}

// SetEnergyTransferOffer set a specific energyTransferOffer in the store
func (k Keeper) SetEnergyTransferOffer(ctx sdk.Context, energyTransferOffer types.EnergyTransferOffer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferOfferKey)
	b := k.cdc.MustMarshal(&energyTransferOffer)
	store.Set(util.GetUint64Key(energyTransferOffer.Id), b)
}

// GetEnergyTransferOffer returns a energyTransferOffer from its id
func (k Keeper) GetEnergyTransferOffer(ctx sdk.Context, id uint64) (val types.EnergyTransferOffer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferOfferKey)
	b := store.Get(util.GetUint64Key(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// MustGetEnergyTransferOffer returns a energyTransferOffer from its id
func (k Keeper) MustGetEnergyTransferOffer(ctx sdk.Context, id uint64) (*types.EnergyTransferOffer, error) {
	offer, found := k.GetEnergyTransferOffer(ctx, id)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "energy transfer offer with id %d not found", id)
	}
	return &offer, nil
}

// RemoveEnergyTransferOffer removes a energyTransferOffer from the store
func (k Keeper) RemoveEnergyTransferOffer(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferOfferKey)
	store.Delete(util.GetUint64Key(id))
}

// GetAllEnergyTransferOffers returns all energyTransferOffer
func (k Keeper) GetAllEnergyTransferOffers(ctx sdk.Context) (list []types.EnergyTransferOffer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferOfferKey)
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EnergyTransferOfferKey)
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
