package keeper

import (
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AppendNewCampaign appends a campaign in the store with a new id and update the count
func (k Keeper) AppendNewCampaign(
	ctx sdk.Context,
	campaign types.Campaign,
) uint64 {
	// Create the vestingAccount
	count := k.GetCampaignCount(ctx)

	// Set the ID of the appended value
	campaign.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKeyPrefix))
	appendedValue := k.cdc.MustMarshal(&campaign)
	store.Set(types.CampaignKey(
		campaign.Id,
	), appendedValue)

	// Update vestingAccount count
	k.SetCampaignCount(ctx, count+1)

	return count
}

// SetCampaign set a specific campaignO in the store from its index
func (k Keeper) SetCampaign(ctx sdk.Context, campaign types.Campaign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKeyPrefix))

	b := k.cdc.MustMarshal(&campaign)
	store.Set(types.CampaignKey(
		campaign.Id,
	), b)
}

// GetCampaign returns a campaignO from its index
func (k Keeper) GetCampaign(
	ctx sdk.Context,
	campaignId uint64,
) (val types.Campaign, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKeyPrefix))

	b := store.Get(types.CampaignKey(
		campaignId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCampaign removes a campaignO from the store
func (k Keeper) removeCampaign(
	ctx sdk.Context,
	campaignId uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKeyPrefix))
	store.Delete(types.CampaignKey(
		campaignId,
	))
}

// GetCampaigns returns all campaignO
func (k Keeper) GetCampaigns(ctx sdk.Context) (list []types.Campaign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Campaign
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetCampaignCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CampaignCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetCampaignCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CampaignCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetCampaign returns a campaignO from its index
func (k Keeper) GetCampaignTotalAmount(
	ctx sdk.Context,
	campaignId uint64,
) (val types.CampaignTotalAmount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignTotalAmountKeyPrefix))

	b := store.Get(types.CampaignTotalAmountKey(
		campaignId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetCampaigns returns all campaignO
func (k Keeper) GetAllCampaignTotalAmount(ctx sdk.Context) (list []types.CampaignTotalAmount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignTotalAmountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CampaignTotalAmount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCampaign returns a campaignO from its index
func (k Keeper) IncrementCampaignTotalAmount(
	ctx sdk.Context,
	airdropDistrubitions types.CampaignTotalAmount,
) (val types.CampaignTotalAmount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignTotalAmountKeyPrefix))

	b := store.Get(types.CampaignTotalAmountKey(
		airdropDistrubitions.CampaignId,
	))

	if b != nil {
		k.cdc.MustUnmarshal(b, &val)
		val.Amount = val.Amount.Add(airdropDistrubitions.Amount...)
	} else {
		val = airdropDistrubitions
	}

	appendedValue := k.cdc.MustMarshal(&val)
	store.Set(types.CampaignTotalAmountKey(
		val.CampaignId,
	), appendedValue)
	return val
}

// GetCampaign returns a campaignO from its index
func (k Keeper) DecrementCampaignTotalAmount(
	ctx sdk.Context,
	campaignId uint64,
	amount sdk.Coins,
) (val types.CampaignTotalAmount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignTotalAmountKeyPrefix))

	b := store.Get(types.CampaignTotalAmountKey(
		campaignId,
	))

	if b == nil {
		return val
	}
	k.cdc.MustUnmarshal(b, &val)
	val.Amount = val.Amount.Sub(amount)

	appendedValue := k.cdc.MustMarshal(&val)
	store.Set(types.CampaignTotalAmountKey(
		campaignId,
	), appendedValue)
	return val
}

// GetCampaign returns a campaignO from its index
func (k Keeper) GetCampaignAmountLeft(
	ctx sdk.Context,
	campaignId uint64,
) (val types.CampaignAmountLeft, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignAmountLeftPrefix))

	b := store.Get(types.CampaignAmountLeftKey(
		campaignId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetCampaigns returns all campaignO
func (k Keeper) GetAllCampaignAmountLeft(ctx sdk.Context) (list []types.CampaignAmountLeft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignAmountLeftPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CampaignAmountLeft
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCampaign returns a campaignO from its index
func (k Keeper) IncrementCampaignAmountLeft(
	ctx sdk.Context,
	airdropClaimsLeft types.CampaignAmountLeft,
) (val types.CampaignAmountLeft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignAmountLeftPrefix))

	b := store.Get(types.CampaignAmountLeftKey(
		airdropClaimsLeft.CampaignId,
	))

	if b != nil {
		k.cdc.MustUnmarshal(b, &val)
		val.Amount = val.Amount.Add(airdropClaimsLeft.Amount...)
	} else {
		val = airdropClaimsLeft
	}

	appendedValue := k.cdc.MustMarshal(&val)
	store.Set(types.CampaignAmountLeftKey(
		val.CampaignId,
	), appendedValue)
	return val
}

// GetCampaign returns a campaignO from its index
func (k Keeper) DecrementCampaignAmountLeft(
	ctx sdk.Context,
	campaignId uint64,
	amount sdk.Coins,
) (val types.CampaignAmountLeft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignAmountLeftPrefix))

	b := store.Get(types.CampaignAmountLeftKey(
		campaignId,
	))

	if b == nil {
		return val
	}
	k.cdc.MustUnmarshal(b, &val)
	val.Amount = val.Amount.Sub(amount)

	appendedValue := k.cdc.MustMarshal(&val)
	store.Set(types.CampaignAmountLeftKey(
		campaignId,
	), appendedValue)
	return val
}
