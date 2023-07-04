package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetCertificateTypeCount get the total number of certificateType
func (k Keeper) GetCertificateTypeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CertificateTypeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCertificateTypeCount set the total number of certificateType
func (k Keeper) SetCertificateTypeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CertificateTypeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCertificateType appends a certificateType in the store with a new id and update the count
func (k Keeper) AppendCertificateType(
	ctx sdk.Context,
	certificateType types.CertificateType,
) uint64 {
	// Create the certificateType
	count := k.GetCertificateTypeCount(ctx)

	// Set the ID of the appended value
	certificateType.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateTypeKey))
	appendedValue := k.cdc.MustMarshal(&certificateType)
	store.Set(GetCertificateTypeIDBytes(certificateType.Id), appendedValue)

	// Update certificateType count
	k.SetCertificateTypeCount(ctx, count+1)

	return count
}

// SetCertificateType set a specific certificateType in the store
func (k Keeper) SetCertificateType(ctx sdk.Context, certificateType types.CertificateType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateTypeKey))
	b := k.cdc.MustMarshal(&certificateType)
	store.Set(GetCertificateTypeIDBytes(certificateType.Id), b)
}

// GetCertificateType returns a certificateType from its id
func (k Keeper) GetCertificateType(ctx sdk.Context, id uint64) (val types.CertificateType, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateTypeKey))
	b := store.Get(GetCertificateTypeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCertificateType removes a certificateType from the store
func (k Keeper) RemoveCertificateType(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateTypeKey))
	store.Delete(GetCertificateTypeIDBytes(id))
}

// GetAllCertificateType returns all certificateType
func (k Keeper) GetAllCertificateType(ctx sdk.Context) (list []types.CertificateType) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateTypeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CertificateType
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCertificateTypeIDBytes returns the byte representation of the ID
func GetCertificateTypeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCertificateTypeIDFromBytes returns ID in uint64 format from a byte array
func GetCertificateTypeIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
