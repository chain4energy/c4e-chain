package keeper

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
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

func (k Keeper) missionIntialStep(ctx sdk.Context, log string, campaignId uint64, missionId uint64, address string) (*types.Campaign, *types.Mission, *types.ClaimRecord, error) {
	campaignConfig := k.Campaign(ctx, campaignId)
	if err := campaignConfig.IsEnabled(ctx.BlockTime()); err != nil {
		return nil, nil, nil, sdkerrors.Wrapf(err, "claim mission - campaignId %d", campaignId)
	}
	k.Logger(ctx).Debug(log, "campaignId", campaignId, "missionId", missionId, "blockTime", ctx.BlockTime(), "campaigh start", campaignConfig.StartTime, "campaigh end", campaignConfig.EndTime)

	mission, found := k.GetMission(ctx, campaignId, missionId)
	if !found {
		k.Logger(ctx).Error(log+" - mission not found", "campaignId", campaignId, "missionId", missionId)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "mission not found - campaignId %d, missionId %d", campaignId, missionId)
	}
	claimRecord, found := k.GetClaimRecord(ctx, address)
	if !found {
		k.Logger(ctx).Error(log+" - claim record not found", "address", address)
		return nil, nil, nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "claim record not found for address %s", address)
	}

	return campaignConfig, &mission, &claimRecord, nil
}

func (k Keeper) ClaimInitialMission(ctx sdk.Context, campaignId uint64, missionId uint64, address string) error {
	campaignConfig, mission, claimRecord, err := k.missionIntialStep(ctx, "claim initial mission", campaignId, missionId, address)
	if err != nil {
		return err
	}
	claimRecord, err = k.completeMission(ctx, true, mission, claimRecord)
	if err != nil {
		return err
	}
	claimRecord, err = k.claimMission(ctx, true, campaignConfig, mission, claimRecord)
	if err != nil {
		return err
	}
	k.SetClaimRecord(ctx, *claimRecord)
	return nil
}

func (k Keeper) ClaimMission(ctx sdk.Context, initialClaim bool, campaignId uint64, missionId uint64, address string) error {
	campaignConfig, mission, claimRecord, err := k.missionIntialStep(ctx, "claim mission", campaignId, missionId, address)
	if err != nil {
		return err
	}
	claimRecord, err = k.claimMission(ctx, false, campaignConfig, mission, claimRecord)
	if err != nil {
		return err
	}

	k.SetClaimRecord(ctx, *claimRecord)
	return nil
}

func (k Keeper) claimMission(ctx sdk.Context, initialClaim bool, campaignConfig *types.Campaign, mission *types.Mission, claimRecord *types.ClaimRecord) (*types.ClaimRecord, error) {
	campaignId := mission.CampaignId
	missionId := mission.MissionId
	address := claimRecord.Address
	if !claimRecord.IsMissionCompleted(campaignId, missionId) {
		k.Logger(ctx).Error("claim mission - mission not completed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionNotCompleted, "mission not completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if claimRecord.IsMissionClaimed(campaignId, missionId) {
		k.Logger(ctx).Error("claim mission - mission already claimed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionClaimed, "mission already claimed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	claimRecord.ClaimMission(campaignId, missionId)

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
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParsing, "wrong claiming address %s: "+err.Error(), sendTo)

		// return errors.Criticalf("invalid claimer address %s", err.Error())
	}
	start := ctx.BlockTime().Add(campaignConfig.LockupPeriod)
	end := start.Add(campaignConfig.VestingPeriod)
	if err := k.SendToAirdropAccount(ctx, claimer, claimable, start.Unix(), end.Unix(), initialClaim); err != nil {
		return nil, sdkerrors.Wrapf(c4eerrors.ErrSendCoins, "send to claiming address %s error: "+err.Error(), sendTo)
	}
	return claimRecord, nil

}

func (k Keeper) CompleteMission(ctx sdk.Context, campaignId uint64, missionId uint64, address string) error {
	_, mission, claimRecord, err := k.missionIntialStep(ctx, "complete mission", campaignId, missionId, address)
	if err != nil {
		return err
	}
	claimRecord, err = k.completeMission(ctx, false, mission, claimRecord)
	if err != nil {
		return err
	}
	k.SetClaimRecord(ctx, *claimRecord)
	return nil
}

// CompleteMission triggers the completion of the mission and distribute the claimable portion of airdrop to the user
// the method fails if the mission has already been completed
func (k Keeper) completeMission(ctx sdk.Context, initialClaim bool, mission *types.Mission, claimRecord *types.ClaimRecord) (*types.ClaimRecord, error) {
	campaignId := mission.CampaignId
	missionId := mission.MissionId
	address := claimRecord.Address
	// check if the mission is already complted for the claim record
	if claimRecord.IsMissionCompleted(campaignId, missionId) {
		k.Logger(ctx).Error("complete mission - mission already completed", "address", address, "campaignId", campaignId, "missionId", missionId)
		return nil, sdkerrors.Wrapf(types.ErrMissionCompleted, "mission already completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
	}

	if !initialClaim {
		initialClaim, found := k.GetInitialClaim(ctx, campaignId)
		if !found {
			k.Logger(ctx).Error("complete mission - initial claim not found", "campaignId", campaignId)
			return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "initial claim not found - campaignId %d", campaignId)
		}
		if !claimRecord.IsMissionClaimed(initialClaim.CampaignId, initialClaim.MissionId) {
			k.Logger(ctx).Error("complete mission - mission already completed", "address", address, "campaignId", campaignId, "missionId", missionId)
			return nil, sdkerrors.Wrapf(types.ErrMissionClaimed, "mission already completed: address %s, campaignId: %d, missionId: %d", address, campaignId, missionId)
		}
	}

	claimRecord.CompleteMission(campaignId, missionId)

	// k.SetClaimRecord(ctx, claimRecord)

	// err = ctx.EventManager().EmitTypedEvent(&types.EventMissionCompleted{
	// 	MissionID: missionID,
	// 	Claimer:   address,
	// })

	return claimRecord, nil
}
