package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCampaign set a specific campaignO in the store from its index
func (k Keeper) SetCampaign(ctx sdk.Context, campaign types.Campaign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKeyPrefix))
	b := k.cdc.MustMarshal(&campaign)
	store.Set(types.CampaignKey(
		campaign.CampaignId,
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

// GetAllCampaign returns all campaignO
func (k Keeper) GetAllCampaign(ctx sdk.Context) (list []types.Campaign) {
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
