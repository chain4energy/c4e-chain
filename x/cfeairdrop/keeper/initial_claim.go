package keeper

import (
	"encoding/binary"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetCampaignIdBytes returns the byte representation of the ID
func GetCampaignIdBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// SetInitialClaim set a specific initialClaim in the store from its index
func (k Keeper) SetInitialClaim(ctx sdk.Context, initialClaim types.InitialClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKeyPrefix))
	b := k.cdc.MustMarshal(&initialClaim)
	store.Set(GetCampaignIdBytes(
		initialClaim.CampaignId,
	), b)
}

// GetInitialClaim returns a initialClaim from its index
func (k Keeper) GetInitialClaim(
	ctx sdk.Context,
	campaignId uint64,

) (val types.InitialClaim, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKeyPrefix))

	b := store.Get(GetCampaignIdBytes(
		campaignId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInitialClaim removes a initialClaim from the store
func (k Keeper) RemoveInitialClaim(
	ctx sdk.Context,
	campaignId uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKeyPrefix))
	store.Delete(GetCampaignIdBytes(
		campaignId,
	))
}

// GetAllInitialClaim returns all initialClaim
func (k Keeper) GetAllInitialClaim(ctx sdk.Context) (list []types.InitialClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InitialClaim
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ClaimInitial(ctx sdk.Context, campaignId uint64, missionId uint64, claimer string) error {
	// retrieve initial claim information
	initialClaim, found := k.GetInitialClaim(ctx, campaignId)
	if !found {
		return nil //types.ErrInitialClaimNotFound
	}
	if !initialClaim.Enabled {
		return nil // types.ErrInitialClaimNotEnabled
	}
	if err := k.CompleteMission(ctx, true, campaignId, initialClaim.MissionId, claimer); err != nil {
		return nil // errors.Wrap(types.ErrMissionCompleteFailure, err.Error())
	}
	return nil
}



