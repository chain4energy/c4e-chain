package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO tests when campaigns in params are nil

// SetMission set a specific mission in the store from its index
func (k Keeper) SetMission(ctx sdk.Context, mission types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKeyPrefix))

	b := k.cdc.MustMarshal(&mission)
	store.Set(types.MissionKey(
		mission.CampaignId,
		mission.MissionId,
	), b)
}

// GetMission returns a mission from its index
func (k Keeper) GetMission(
	ctx sdk.Context,
	campaignId uint64,
	missionId uint64,

) (val types.Mission, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKeyPrefix))

	b := store.Get(types.MissionKey(
		campaignId,
		missionId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMission removes a mission from the store
func (k Keeper) RemoveMission(
	ctx sdk.Context,
	campaignId uint64,
	missionId uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKeyPrefix))
	store.Delete(types.MissionKey(
		campaignId,
		missionId,
	))
}

// GetAllMission returns all mission
func (k Keeper) GetAllMission(ctx sdk.Context) (list []types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mission
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
