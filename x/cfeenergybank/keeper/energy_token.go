package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetEnergyTokenCount get the total number of energyToken
func (k Keeper) GetEnergyTokenCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EnergyTokenCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEnergyTokenCount set the total number of energyToken
func (k Keeper) SetEnergyTokenCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EnergyTokenCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEnergyToken appends a energyToken in the store with a new id and update the count
func (k Keeper) AppendEnergyToken(
	ctx sdk.Context,
	energyToken types.EnergyToken,
) uint64 {
	// Create the energyToken
	count := k.GetEnergyTokenCount(ctx)

	// Set the ID of the appended value
	energyToken.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTokenKey))
	appendedValue := k.cdc.MustMarshal(&energyToken)
	store.Set(GetEnergyTokenIDBytes(energyToken.Id), appendedValue)

	// Update energyToken count
	k.SetEnergyTokenCount(ctx, count+1)

	return count
}

// SetEnergyToken set a specific energyToken in the store
func (k Keeper) SetEnergyToken(ctx sdk.Context, energyToken types.EnergyToken) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTokenKey))
	b := k.cdc.MustMarshal(&energyToken)
	store.Set(GetEnergyTokenIDBytes(energyToken.Id), b)
}

// GetEnergyToken returns a energyToken from its id
func (k Keeper) GetEnergyToken(ctx sdk.Context, id uint64) (val types.EnergyToken, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTokenKey))
	b := store.Get(GetEnergyTokenIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEnergyToken removes a energyToken from the store
func (k Keeper) RemoveEnergyToken(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTokenKey))
	store.Delete(GetEnergyTokenIDBytes(id))
}

// GetAllEnergyToken returns all energyToken
func (k Keeper) GetAllEnergyToken(ctx sdk.Context) (list []types.EnergyToken) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyTokenKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EnergyToken
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetEnergyTokenIDBytes returns the byte representation of the ID
func GetEnergyTokenIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetEnergyTokenIDFromBytes returns ID in uint64 format from a byte array
func GetEnergyTokenIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
