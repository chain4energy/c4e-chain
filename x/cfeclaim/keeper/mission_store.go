package keeper

import (
	"cosmossdk.io/errors"
	"encoding/binary"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetMission set a specific mission in the store from its index
func (k Keeper) SetMission(ctx sdk.Context, mission types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MissionKeyPrefix)

	b := k.cdc.MustMarshal(&mission)
	store.Set(types.MissionKey(
		mission.CampaignId,
		mission.Id,
	), b)
}

// GetMission returns a mission from its index
func (k Keeper) GetMission(
	ctx sdk.Context,
	campaignId uint64,
	missionId uint64,
) (val types.Mission, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MissionKeyPrefix)

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

// GetMission returns a mission from its index
func (k Keeper) MustGetMission(
	ctx sdk.Context,
	campaignId uint64,
	missionId uint64,
) (*types.Mission, error) {
	mission, found := k.GetMission(
		ctx,
		campaignId,
		missionId,
	)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "mission not found - campaignId %d, missionId %d", campaignId, missionId)
	}
	return &mission, nil
}

// GetAllMission returns all mission
func (k Keeper) GetAllMission(ctx sdk.Context) (list []types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MissionKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mission
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllMissionForCampaign returns all mission
func (k Keeper) AllMissionForCampaign(ctx sdk.Context, campaignId uint64) (list []types.Mission, weightSum sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MissionKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	weightSum = sdk.ZeroDec()
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mission
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.CampaignId == campaignId {
			list = append(list, val)
			weightSum = weightSum.Add(val.Weight)
		}
	}

	return
}

// GetAllMissionForCampaign returns all mission
func (k Keeper) RemoveAllMissionForCampaign(ctx sdk.Context, campaignId uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MissionKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mission
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.CampaignId == campaignId {
			store.Delete(types.MissionKey(
				campaignId,
				val.Id,
			))
		}
	}
}

func (k Keeper) AppendNewMission(
	ctx sdk.Context,
	campaignId uint64,
	mission types.Mission,
) uint64 {
	// Create the vestingAccount
	count := k.GetMissionCount(ctx, campaignId)

	// Set the ID of the appended value
	mission.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MissionKeyPrefix)
	appendedValue := k.cdc.MustMarshal(&mission)
	store.Set(types.MissionKey(
		campaignId,
		mission.Id,
	), appendedValue)

	// Update vestingAccount count
	k.SetMissionCount(ctx, campaignId, count+1)

	return count
}

func (k Keeper) GetMissionCount(ctx sdk.Context, campaignId uint64) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MissionCountKeyPrefix)
	bz := store.Get(types.GetUint64Key(campaignId))

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetMissionCount(ctx sdk.Context, campaignId uint64, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MissionCountKeyPrefix)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(types.GetUint64Key(campaignId), bz)
}
