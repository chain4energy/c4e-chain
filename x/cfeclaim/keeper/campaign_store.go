package keeper

import (
	"cosmossdk.io/errors"
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CampaignKeyPrefix)
	appendedValue := k.cdc.MustMarshal(&campaign)
	store.Set(types.GetUint64Key(
		campaign.Id,
	), appendedValue)

	// Update vestingAccount count
	k.SetCampaignCount(ctx, count+1)

	return count
}

// SetCampaign set a specific campaignO in the store from its index
func (k Keeper) SetCampaign(ctx sdk.Context, campaign types.Campaign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CampaignKeyPrefix)

	b := k.cdc.MustMarshal(&campaign)
	store.Set(types.GetUint64Key(
		campaign.Id,
	), b)
}

// GetCampaign returns a campaignO from its index
func (k Keeper) GetCampaign(
	ctx sdk.Context,
	campaignId uint64,
) (val types.Campaign, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CampaignKeyPrefix)

	b := store.Get(types.GetUint64Key(
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CampaignKeyPrefix)
	store.Delete(types.GetUint64Key(
		campaignId,
	))
}

// GetAllCampaigns returns all campaignO
func (k Keeper) GetAllCampaigns(ctx sdk.Context) (list []types.Campaign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CampaignKeyPrefix)
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
	byteKey := types.CampaignCountKeyPrefix
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
	byteKey := types.CampaignCountKeyPrefix
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) MustGetCampaign(ctx sdk.Context, campaignId uint64) (*types.Campaign, error) {
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "campaign with id %d not found", campaignId)
	}

	return &campaign, nil
}
