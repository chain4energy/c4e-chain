package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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

// CompleteMission triggers the completion of the mission and distribute the claimable portion of airdrop to the user
// the method fails if the mission has already been completed
func (k Keeper) CompleteMission(ctx sdk.Context, initialClaim bool, campaignId uint64, missionId uint64, address string) error {
	campaignConfig := k.Campaign(ctx, campaignId)
	if !campaignConfig.Enabled {
		return nil
	}
	k.Logger(ctx).Debug("complete mission", "campaignId", campaignId, "blockTime", ctx.BlockTime(), "campaigh start", campaignConfig.StartTime, "campaigh end", campaignConfig.EndTime)
	if ctx.BlockTime().Before(campaignConfig.StartTime) || (campaignConfig.EndTime != nil && ctx.BlockTime().After(*campaignConfig.EndTime)) {
		k.Logger(ctx).Debug("complete mission campaign not enabled due time", "campaignId", campaignId)

		return nil
	}
	// airdropSupply, found := k.GetAirdropSupply(ctx)
	// if !found {
	// return errors.Wrapf(types.ErrAirdropSupplyNotFound, "airdrop supply is not defined")
	// }

	// retrieve mission
	mission, found := k.GetMission(ctx, campaignId, missionId)
	if !found {
		k.Logger(ctx).Error("mission not found", "campaignId", campaignId, "missionId", missionId)
		return sdkerrors.Wrapf(types.ErrMissionNotFound, "campaignId %d, missionId %d", campaignId, missionId)
		// return errors.Wrapf(types.ErrMissionNotFound, "mission %d not found", missionID)
	}

	// retrieve claim record of the user
	claimRecord, found := k.GetClaimRecord(ctx, address)
	if !found {
		return nil
		// return errors.Wrapf(types.ErrClaimRecordNotFound, "claim record not found for address %s", address)
	}

	// check if the mission is already complted for the claim record
	if claimRecord.IsMissionCompleted(campaignId, missionId) {
		// return errors.Wrapf(
		// 	types.ErrMissionCompleted,
		// 	"mission %d completed for address %s",
		// 	missionID,
		// 	address,
		// )
		return nil
	}

	if !initialClaim {
		initialClaim, found := k.GetInitialClaim(ctx, campaignId)
		if !found {
			return nil //types.ErrInitialClaimNotFound
		}
		if !claimRecord.IsMissionCompleted(initialClaim.CampaignId, initialClaim.MissionId) {
			return nil //types.ErrInitialClaimNotFound
		}
	}
	// claimRecord.CompletedMissions = append(claimRecord.CompletedMissions, missionID)

	claimRecord.CompleteMission(campaignId, missionId)
	// calculate claimable from mission weight and claim
	claimableAmount := claimRecord.ClaimableFromMission(mission)
	// claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Denom, claimableAmount))
	claimable := sdk.NewCoins(sdk.NewCoin("uc4e", claimableAmount)) // TODO - uc4e to param

	// calculate claimable after decay factor
	// decayInfo := k.GetParams(ctx).DecayInformation
	// claimable = decayInfo.ApplyDecayFactor(claimable, ctx.BlockTime())

	// check final claimable non-zero
	// if claimable.Empty() {
	// 	return types.ErrNoClaimable
	// }

	// decrease airdrop supply
	// airdropSupply.Amount = airdropSupply.Amount.Sub(claimable.AmountOf(airdropSupply.Denom))
	// if airdropSupply.Amount.IsNegative() {
	// 	return errors.Critical("airdrop supply is lower than total claimable")
	// }

	// send claimable to the user
	sendTo := address
	if len(claimRecord.ClaimAddress) > 0 {
		sendTo = claimRecord.ClaimAddress
	}
	claimer, err := sdk.AccAddressFromBech32(sendTo)
	if err != nil {
		return nil
		// return errors.Criticalf("invalid claimer address %s", err.Error())
	}
	start := ctx.BlockTime().Add(campaignConfig.LockupPeriod)
	end := start.Add(campaignConfig.VestingPeriod)
	if err := k.SendToAirdropAccount(ctx, claimer, claimable, start.Unix(), end.Unix(), initialClaim); err != nil {
		return nil
		// return errors.Criticalf("can't send claimable coins %s", err.Error())
	}

	// update store
	// k.SetAirdropSupply(ctx, airdropSupply)
	k.SetClaimRecord(ctx, claimRecord)

	// err = ctx.EventManager().EmitTypedEvent(&types.EventMissionCompleted{
	// 	MissionID: missionID,
	// 	Claimer:   address,
	// })

	return err
}
