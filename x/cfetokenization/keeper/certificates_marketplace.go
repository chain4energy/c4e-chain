package keeper

import (
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetUserDevicesCount get the total number of userDevices
func (k Keeper) GetMarketplaceCertificatesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MarketplaceCertificatesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetUserDevicesCount set the total number of userDevices
func (k Keeper) SetMarketplaceCertificatesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MarketplaceCertificatesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendUserDevices appends a userDevices in the store with a new id and update the count
func (k Keeper) AppendMarketplaceCertificate(
	ctx sdk.Context,
	certificateOffer types.CertificateOffer,
) uint64 {
	// Create the userDevices
	count := k.GetMarketplaceCertificatesCount(ctx)

	// Set the ID of the appended value
	certificateOffer.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketplaceCertificatesKey))
	appendedValue := k.cdc.MustMarshal(&certificateOffer)
	store.Set(GetCertificateOfferIDBytes(certificateOffer.Id), appendedValue)

	// Update userDevices count
	k.SetMarketplaceCertificatesCount(ctx, count+1)

	return count
}

// SetUserDevices set a specific userDevices in the store
func (k Keeper) SetMarketplaceCertificate(ctx sdk.Context, userDevices types.CertificateOffer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketplaceCertificatesKey))
	b := k.cdc.MustMarshal(&userDevices)
	store.Set(GetCertificateOfferIDBytes(userDevices.Id), b)
}

// GetUserDevices returns a userDevices from its id
func (k Keeper) GetMarketplaceCertificate(ctx sdk.Context, id uint64) (val types.CertificateOffer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketplaceCertificatesKey))
	b := store.Get(GetCertificateOfferIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllUserDevices returns all userDevices
func (k Keeper) GetMarketplaceCertificates(ctx sdk.Context) (list []types.CertificateOffer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketplaceCertificatesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CertificateOffer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCertificateOfferIDBytes returns the byte representation of the ID
func GetCertificateOfferIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
