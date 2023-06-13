package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"time"
)

func (k Keeper) CreateCampaign(ctx sdk.Context, owner string, name string, description string, campaignType types.CampaignType, removableClaimRecords bool,
	feegrantAmount math.Int, initialClaimFreeAmount math.Int, free sdk.Dec, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration, vestingPoolName string) (*types.Campaign, error) {
	k.Logger(ctx).Debug("create campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)

	if err := k.ValidateCampaignParams(ctx, name, description, feegrantAmount, initialClaimFreeAmount, free, startTime, endTime, campaignType, owner, vestingPoolName, lockupPeriod, vestingPeriod); err != nil {
		return nil, err
	}

	if err := validateEndTimeAfterBlockTime(endTime, ctx.BlockTime()); err != nil {
		return nil, err
	}

	campaign := types.Campaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		CampaignType:           campaignType,
		RemovableClaimRecords:  removableClaimRecords,
		FeegrantAmount:         feegrantAmount,
		InitialClaimFreeAmount: initialClaimFreeAmount,
		Free:                   free,
		Enabled:                false,
		StartTime:              startTime,
		EndTime:                endTime,
		LockupPeriod:           lockupPeriod,
		VestingPeriod:          vestingPeriod,
		VestingPoolName:        vestingPoolName,
	}

	campaign.Id = k.AppendNewCampaign(ctx, campaign)
	// Adding the inititalClaim mission to a campaign is done automatically as this mission is required for every campaign
	k.AppendNewMission(ctx, campaign.Id, *types.NewInitialMission(campaign.Id))
	k.Logger(ctx).Debug("create campaign ret", "campaignId", campaign.Id)
	event := &types.EventNewCampaign{
		Id:                     strconv.FormatUint(campaign.Id, 10),
		Owner:                  campaign.Owner,
		Name:                   campaign.Name,
		Description:            campaign.Description,
		CampaignType:           campaign.CampaignType.String(),
		FeegrantAmount:         campaign.FeegrantAmount.String(),
		InitialClaimFreeAmount: campaign.InitialClaimFreeAmount.String(),
		Enabled:                "false",
		StartTime:              campaign.StartTime.String(),
		EndTime:                campaign.EndTime.String(),
		LockupPeriod:           campaign.LockupPeriod.String(),
		VestingPeriod:          campaign.VestingPeriod.String(),
		VestingPoolName:        campaign.VestingPoolName,
	}
	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("create campaign emit event error", "event", event, "error", err.Error())
	}

	k.Logger(ctx).Debug("create campaign ret", "campaignId", campaign.Id)
	return &campaign, nil
}

func (k Keeper) CloseCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("close campaign", "owner", owner, "campaignId", campaignId)
	campaign, err := k.MustGetCampaign(ctx, campaignId)
	if err != nil {
		return err
	}
	if err = k.validateCloseCampaignParams(ctx, campaign, owner); err != nil {
		return err
	}
	if err = k.returnAllToOwner(ctx, campaign); err != nil {
		return err
	}
	campaign.DecrementCampaignCurrentAmount(campaign.CampaignCurrentAmount)
	campaign.Enabled = false
	k.SetCampaign(ctx, *campaign)
	k.Logger(ctx).Debug("closed campaign", "campaignId", campaignId, "decrementedCampaignCurrentAmount", campaign.CampaignCurrentAmount)
	return nil
}

func (k Keeper) RemoveCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("remove campaign", "owner", owner, "campaignId", campaignId)

	campaign, err := k.MustGetCampaign(ctx, campaignId)
	if err != nil {
		return err
	}
	if err = campaign.ValidateRemoveCampaignParams(owner); err != nil {
		return err
	}

	if err = k.returnAllToOwner(ctx, campaign); err != nil {
		return err
	}

	k.removeCampaign(ctx, campaignId)
	k.RemoveAllMissionForCampaign(ctx, campaignId)
	return nil
}

func (k Keeper) EnableCampaign(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) error {
	k.Logger(ctx).Debug("enable campaign", "owner", owner, "campaignId", campaignId, "startTime", &startTime, "endTime", &endTime)

	campaign, err := k.MustGetCampaign(ctx, campaignId)
	if err != nil {
		return err
	}

	if startTime != nil {
		campaign.StartTime = *startTime
	}
	if endTime != nil {
		campaign.EndTime = *endTime
	}

	err = k.validateEnableCampaignParams(ctx, campaign, owner)
	if err != nil {
		return err
	}

	campaign.Enabled = true
	k.SetCampaign(ctx, *campaign)
	k.Logger(ctx).Debug("enabled campaign", "campaignId", campaignId, "startTime", campaign.StartTime, "endTime", campaign.EndTime)
	return nil
}

func (k Keeper) returnAllToOwner(ctx sdk.Context, campaign *types.Campaign) error {
	k.Logger(ctx).Debug("return all to owner", "campaignId", campaign.Id)
	if err := k.sendCampaignCurrentAmountToOwner(ctx, campaign, campaign.CampaignCurrentAmount); err != nil {
		return err
	}
	return k.sendCampaignFeegrantToOwner(ctx, campaign)
}

func (k Keeper) sendCampaignCurrentAmountToOwner(ctx sdk.Context, campaign *types.Campaign, amount sdk.Coins) error {
	if amount.IsZero() {
		return nil
	}
	if !amount.IsAllLTE(campaign.CampaignCurrentAmount) {
		return errors.Wrapf(c4eerrors.ErrAmount,
			"cannot send campaign current amount to owner, campaign current amount is lower than amount (%s < %s)", campaign.CampaignCurrentAmount, amount)
	}
	if campaign.CampaignType == types.VestingPoolCampaign {
		if err := k.vestingKeeper.RemoveVestingPoolReservation(ctx, campaign.Owner, campaign.VestingPoolName, campaign.Id,
			amount.AmountOf(k.vestingKeeper.Denom(ctx))); err != nil {
			return err
		}
	} else {
		ownerAddress, _ := sdk.AccAddressFromBech32(campaign.Owner)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, amount); err != nil {
			return err
		}
	}

	k.Logger(ctx).Debug("send campaign current amount to owner", "campaignId", campaign.Id, "owner", campaign.Owner, "amount", amount)
	return nil
}

func (k Keeper) ValidateCampaignParams(ctx sdk.Context, name string, description string, feegrantAmount math.Int,
	inititalClaimFreeAmount math.Int, free sdk.Dec, startTime time.Time, endTime time.Time,
	campaignType types.CampaignType, owner string, vestingPoolName string, lockupPeriod time.Duration, vestingPeriod time.Duration) error {

	if err := types.ValidateCreateCampaignParams(name, description, feegrantAmount, inititalClaimFreeAmount,
		free, startTime, endTime, campaignType, lockupPeriod, vestingPeriod, vestingPoolName); err != nil {
		return err
	}
	if campaignType == types.VestingPoolCampaign {
		return k.ValidateVestingPoolCampaign(ctx, owner, vestingPoolName, lockupPeriod, vestingPeriod, free)
	}

	return nil
}

func (k Keeper) ValidateVestingPoolCampaign(ctx sdk.Context, owner string, vestingPoolName string,
	lockupPeriod time.Duration, vestingPeriod time.Duration, free sdk.Dec) error {
	vestingType, err := k.vestingKeeper.MustGetVestingTypeForVestingPool(ctx, owner, vestingPoolName)
	if err != nil {
		return err
	}
	if err = vestingType.ValidateVestingPeriods(lockupPeriod, vestingPeriod); err != nil {
		return err
	}
	return vestingType.ValidateVestingFree(free)
}

func validateEndTimeAfterBlockTime(endTime time.Time, blockTime time.Time) error {
	if endTime.Before(blockTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "end time in the past error (%s < %s)", endTime, blockTime)
	}
	return nil
}

func (k Keeper) validateEnableCampaignParams(ctx sdk.Context, campaign *types.Campaign, owner string) error {
	if err := campaign.ValidateEnableCampaignParams(owner); err != nil {
		return err
	}
	return campaign.ValidateEndTimeAfterBlockTime(ctx.BlockTime())
}

func (k Keeper) validateCloseCampaignParams(ctx sdk.Context, campaign *types.Campaign, owner string) error {
	if err := campaign.ValidateCloseCampaignParams(owner); err != nil {
		return err
	}
	return campaign.ValidateEnded(ctx.BlockTime())
}
