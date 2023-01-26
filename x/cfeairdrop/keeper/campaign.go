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
func (k Keeper) RemoveCampaign(
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

func (k Keeper) GetWhitelistedVestingAccounts() []string {
	return []string{"c4e1asgp8qrlznsjs7ww5f60lf64gx04s6nsrte4dv"}
}

// GetCampaign returns a campaignO from its index
func (k Keeper) GetAirdropDistrubitions(
	ctx sdk.Context,
	campaignId uint64,
) (val types.AirdropDistrubitions, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropDistributionsPrefix))

	b := store.Get(types.AirdropDistributionsKey(
		campaignId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetCampaigns returns all campaignO
func (k Keeper) GetAllAirdropDistrubitions(ctx sdk.Context) (list []types.AirdropDistrubitions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropDistributionsPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AirdropDistrubitions
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCampaign returns a campaignO from its index
func (k Keeper) IncrementAirdropDistrubitions(
	ctx sdk.Context,
	airdropDistrubitions types.AirdropDistrubitions,
) (val types.AirdropDistrubitions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropDistributionsPrefix))

	b := store.Get(types.AirdropDistributionsKey(
		airdropDistrubitions.CampaignId,
	))

	if b != nil {
		k.cdc.MustUnmarshal(b, &val)
		val.AirdropCoins = val.AirdropCoins.Add(airdropDistrubitions.AirdropCoins...)
	} else {
		val = airdropDistrubitions
	}

	appendedValue := k.cdc.MustMarshal(&val)
	store.Set(types.AirdropDistributionsKey(
		val.CampaignId,
	), appendedValue)
	return val
}

// GetCampaign returns a campaignO from its index
func (k Keeper) DecrementAirdropDistrubitions(
	ctx sdk.Context,
	campaignId uint64,
	airdropCoins sdk.Coins,
) (val types.AirdropDistrubitions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropDistributionsPrefix))

	b := store.Get(types.AirdropDistributionsKey(
		campaignId,
	))

	if b == nil {
		return val
	}
	k.cdc.MustUnmarshal(b, &val)
	val.AirdropCoins = val.AirdropCoins.Sub(airdropCoins)

	appendedValue := k.cdc.MustMarshal(&val)
	store.Set(types.AirdropDistributionsKey(
		campaignId,
	), appendedValue)
	return val
}

// GetCampaign returns a campaignO from its index
func (k Keeper) GetAirdropClaimsLeft(
	ctx sdk.Context,
	campaignId uint64,
) (val types.AirdropClaimsLeft, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropClaimsLeftPrefix))

	b := store.Get(types.AirdropClaimsLeftKey(
		campaignId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetCampaigns returns all campaignO
func (k Keeper) GetAllAirdropClaimsLeft(ctx sdk.Context) (list []types.AirdropClaimsLeft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropClaimsLeftPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AirdropClaimsLeft
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCampaign returns a campaignO from its index
func (k Keeper) IncrementAirdropClaimsLeft(
	ctx sdk.Context,
	airdropClaimsLeft types.AirdropClaimsLeft,
) (val types.AirdropClaimsLeft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropClaimsLeftPrefix))

	b := store.Get(types.AirdropDistributionsKey(
		airdropClaimsLeft.CampaignId,
	))

	if b != nil {
		k.cdc.MustUnmarshal(b, &val)
		val.AirdropCoins = val.AirdropCoins.Add(airdropClaimsLeft.AirdropCoins...)
	} else {
		val = airdropClaimsLeft
	}

	appendedValue := k.cdc.MustMarshal(&val)
	store.Set(types.AirdropClaimsLeftKey(
		val.CampaignId,
	), appendedValue)
	return val
}

// GetCampaign returns a campaignO from its index
func (k Keeper) DecrementAirdropClaimsLeft(
	ctx sdk.Context,
	campaignId uint64,
	amount sdk.Coins,
) (val types.AirdropClaimsLeft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropClaimsLeftPrefix))

	b := store.Get(types.AirdropClaimsLeftKey(
		campaignId,
	))

	if b == nil {
		return val
	}
	k.cdc.MustUnmarshal(b, &val)
	val.AirdropCoins = val.AirdropCoins.Sub(amount)

	appendedValue := k.cdc.MustMarshal(&val)
	store.Set(types.AirdropDistributionsKey(
		campaignId,
	), appendedValue)
	return val
}
