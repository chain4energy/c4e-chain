package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"strconv"
	"time"
)

func (k Keeper) CreateCampaign(ctx sdk.Context, owner string, name string, description string, campaignType types.CampaignType, feeGrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration, vestingPoolName string) error {
	k.Logger(ctx).Debug("create campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)

	if err := k.ValidateCampaignParams(ctx, name, description, startTime, endTime, campaignType, owner, vestingPoolName, lockupPeriod, vestingPeriod); err != nil {
		return err
	}

	feeGrantAmount, err := getFeeGrantAmount(feeGrantAmount)
	if err != nil {
		return err
	}

	initialClaimFreeAmount, err = getInitialClaimFreeAmount(initialClaimFreeAmount)
	if err != nil {
		return err
	}

	campaign := types.Campaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		CampaignType:           campaignType,
		FeegrantAmount:         *feeGrantAmount,
		InitialClaimFreeAmount: *initialClaimFreeAmount,
		Enabled:                false,
		StartTime:              *startTime,
		EndTime:                *endTime,
		LockupPeriod:           *lockupPeriod,
		VestingPeriod:          *vestingPeriod,
		VestingPoolName:        vestingPoolName,
	}

	campaignId := k.AppendNewCampaign(ctx, campaign)

	missionInitial := types.NewInitialMission(campaignId)
	err = k.AddMissionToCampaign(ctx, owner, campaignId, missionInitial.Name, missionInitial.Description,
		missionInitial.MissionType, missionInitial.Weight, missionInitial.ClaimStartDate)
	if err != nil {
		return err
	}

	return nil
}

func getInitialClaimFreeAmount(initialClaimFreeAmount *math.Int) (*math.Int, error) {
	if initialClaimFreeAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		initialClaimFreeAmount = &zeroInt
	}

	if initialClaimFreeAmount.IsNegative() {
		return nil, errors.Wrapf(c4eerrors.ErrParam, "initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}

	return initialClaimFreeAmount, nil
}

func getFeeGrantAmount(feeGrantAmount *sdk.Int) (*sdk.Int, error) {
	if feeGrantAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		feeGrantAmount = &zeroInt
	}

	if feeGrantAmount.IsNegative() {
		return nil, errors.Wrapf(c4eerrors.ErrParam, "feegrant amount (%s) cannot be negative", feeGrantAmount.String())
	}

	return feeGrantAmount, nil
}

func (k Keeper) EditCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, campaignType types.CampaignType, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration, vestingPoolName string, newOwner string) error {
	k.Logger(ctx).Debug("edit claim campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)

	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		k.Logger(ctx).Debug("edit claim campaign", "err", err.Error())
		return err
	}

	if err = types.ValidateCampaignIsNotEnabled(campaign); err != nil {
		k.Logger(ctx).Debug("edit claim campaign", "err", err.Error())
		return err
	}

	campaign, err = k.updateCampaignWithNewParams(name, description, campaignType, startTime, endTime, lockupPeriod, vestingPeriod, vestingPoolName, newOwner, campaign, ctx)
	k.SetCampaign(ctx, campaign)

	event := &types.EditCampaign{
		Owner:                  campaign.Owner,
		Name:                   campaign.Name,
		Description:            campaign.Description,
		CampaignType:           campaign.CampaignType.String(),
		FeegrantAmount:         campaign.FeegrantAmount.String(),
		InitialClaimFreeAmount: campaign.InitialClaimFreeAmount.String(),
		Enabled:                strconv.FormatBool(campaign.Enabled),
		StartTime:              campaign.StartTime.String(),
		EndTime:                campaign.EndTime.String(),
		LockupPeriod:           campaign.LockupPeriod.String(),
		VestingPeriod:          campaign.VestingPeriod.String(),
		VestingPoolName:        campaign.VestingPoolName,
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Debug("edit campaign emit event error", "event", event, "error", err.Error())
	}
	return nil
}

func (k Keeper) updateCampaignWithNewParams(name string, description string, campaignType types.CampaignType, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration, vestingPoolName string, newOwner string,
	campaign types.Campaign, ctx sdk.Context) (types.Campaign, error) {

	if name != "" {
		campaign.Name = name
	}

	if description != "" {
		campaign.Description = description
	}

	if startTime != nil {
		campaign.StartTime = *startTime
	}

	if endTime != nil {
		campaign.EndTime = *endTime
	}

	if vestingPeriod != nil {
		campaign.VestingPeriod = *vestingPeriod
	}

	if lockupPeriod != nil {
		campaign.LockupPeriod = *lockupPeriod
	}

	if vestingPoolName != "" {
		campaign.VestingPoolName = vestingPoolName
	}

	if newOwner != "" {
		campaign.Owner = newOwner
	}

	if err := types.ValidateCampaignType(campaignType, campaign.Owner, vestingPoolName); err == nil {
		campaign.CampaignType = campaignType
	}

	if err := k.ValidateCampaignParams(ctx, campaign.Name, campaign.Description, &campaign.StartTime, &campaign.EndTime, campaign.CampaignType, campaign.Owner, campaign.VestingPoolName, &campaign.LockupPeriod, &campaign.VestingPeriod); err != nil {
		return types.Campaign{}, err
	}
	return campaign, nil
}

func (k Keeper) CloseCampaign(ctx sdk.Context, owner string, campaignId uint64, closeAction types.CloseAction) error {
	k.Logger(ctx).Debug("close campaign", "owner", owner, "campaignId", campaignId, "CloseAction", closeAction)
	campaign, err := k.ValidateCloseCampaignParams(ctx, closeAction, campaignId, owner)
	if err != nil {

		return err
	}

	campaignAmountLeft, _ := k.GetCampaignAmountLeft(ctx, campaign.Id)
	if err = k.closeActionSwitch(ctx, closeAction, &campaign, campaignAmountLeft.Amount); err != nil {
		return err
	}

	campaign.Enabled = false
	k.SetCampaign(ctx, campaign)
	k.DecrementCampaignAmountLeft(ctx, campaignId, campaignAmountLeft.Amount)
	return nil
}

func (k Keeper) closeActionSwitch(ctx sdk.Context, CloseAction types.CloseAction, campaign *types.Campaign, campaignAmountLeft sdk.Coins) error {
	switch CloseAction {
	case types.CloseSendToCommunityPool:
		return k.campaignCloseSendToCommunityPool(ctx, campaign, campaignAmountLeft)
	case types.CampaignCloseBurn:
		return k.campaignCloseBurn(ctx, campaign, campaignAmountLeft)
	case types.CampaignCloseSendToOwner:
		return k.campaignCloseSendToOwner(ctx, campaign, campaignAmountLeft)
	default:
		return errors.Wrap(sdkerrors.ErrInvalidType, "wrong campaign close action type")
	}
}

func (k Keeper) campaignCloseSendToCommunityPool(ctx sdk.Context, campaign *types.Campaign, campaignAmountLeft sdk.Coins) error {
	if err := k.distributionKeeper.FundCommunityPool(ctx, campaignAmountLeft, authtypes.NewModuleAddress(types.ModuleName)); err != nil {
		return err
	}
	if campaign.FeegrantAmount.IsPositive() {
		_, feegrantAccountAddress := FeegrantAccountAddress(campaign.Id)
		feegrantTotalAmount := k.bankKeeper.GetAllBalances(ctx, feegrantAccountAddress)
		if err := k.distributionKeeper.FundCommunityPool(ctx, feegrantTotalAmount, feegrantAccountAddress); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) campaignCloseBurn(ctx sdk.Context, campaign *types.Campaign, coinsToBurn sdk.Coins) error {
	if campaign.FeegrantAmount.IsPositive() {
		_, feegrantAccountAddress := FeegrantAccountAddress(campaign.Id)
		feegrantTotalAmount := k.bankKeeper.GetAllBalances(ctx, feegrantAccountAddress)
		coinsToBurn = coinsToBurn.Add(feegrantTotalAmount...)
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, feegrantAccountAddress, types.ModuleName, feegrantTotalAmount); err != nil {
			return err
		}
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, coinsToBurn); err != nil {
		return err
	}
	return nil
}

func (k Keeper) campaignCloseSendToOwner(ctx sdk.Context, campaign *types.Campaign, campaignAmountLeft sdk.Coins) error {
	if campaign.CampaignType != types.CampaignSale {
		ownerAddress, _ := sdk.AccAddressFromBech32(campaign.Owner)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, campaignAmountLeft); err != nil {
			return err
		}
		if campaign.FeegrantAmount.IsPositive() {
			_, feegrantAccountAddress := FeegrantAccountAddress(campaign.Id)
			feegrantTotalAmount := k.bankKeeper.GetAllBalances(ctx, feegrantAccountAddress)
			if err := k.bankKeeper.SendCoins(ctx, feegrantAccountAddress, ownerAddress, feegrantTotalAmount); err != nil {
				return err
			}
		}
	} else {
		vestingDenom := k.vestingKeeper.Denom(ctx)
		accountVestingPools, _ := k.vestingKeeper.GetAccountVestingPools(ctx, campaign.Owner)
		var vestingPool *cfevestingtypes.VestingPool
		for _, vestPool := range accountVestingPools.VestingPools {
			if vestPool.Name == campaign.VestingPoolName {
				vestingPool = vestPool
				break
			}
		}

		if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, cfevestingtypes.ModuleName, campaignAmountLeft); err != nil {
			return err
		}
		vestingPool.Sent = vestingPool.Sent.Sub(campaignAmountLeft.AmountOf(vestingDenom))

		if campaign.FeegrantAmount.IsPositive() {
			_, feegrantAccountAddress := FeegrantAccountAddress(campaign.Id)
			feegrantTotalAmount := k.bankKeeper.GetAllBalances(ctx, feegrantAccountAddress)
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, feegrantAccountAddress, cfevestingtypes.ModuleName, feegrantTotalAmount); err != nil {
				return err
			}
			vestingPool.Sent = vestingPool.Sent.Sub(feegrantTotalAmount.AmountOf(vestingDenom))
		}

		k.vestingKeeper.SetAccountVestingPools(ctx, accountVestingPools)
	}
	return nil
}

func (k Keeper) StartCampaign(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) error {
	k.Logger(ctx).Debug("start campaign", "owner", owner, "campaignId", campaignId)

	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return err
	}
	if campaign.CampaignType == types.CampaignSale {
		if startTime != nil {
			campaign.StartTime = *startTime
		}
		if endTime != nil {
			campaign.EndTime = *endTime
		}
	}

	err = k.ValidateStartCampaignParams(ctx, campaign, owner)
	if err != nil {
		return err
	}

	campaign.Enabled = true
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) RemoveCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("remove campaign", "owner", owner, "campaignId", campaignId)

	err := k.ValidateRemoveCampaignParams(ctx, owner, campaignId)
	if err != nil {
		k.Logger(ctx).Debug("remove campaign", "err", err.Error())
		return err
	}

	k.removeCampaign(ctx, campaignId)
	k.RemoveAllMissionForCampaign(ctx, campaignId)
	return nil
}
func (k Keeper) ValidateCampaignParams(ctx sdk.Context, name string, description string, startTime *time.Time, endTime *time.Time,
	campaignType types.CampaignType, owner string, vestingPoolName string, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	if err := types.ValidateCampaignCreateParams(name, description, startTime, endTime, campaignType, owner, vestingPoolName); err != nil {
		return err
	}

	if campaignType != types.CampaignSale && startTime != nil {
		if err := ValidateCampaignStartTimeInTheFuture(ctx, startTime); err != nil {
			return err
		}
	}
	if campaignType == types.CampaignSale {
		return k.ValidateCampaignWhenAddedFromVestingPool(ctx, owner, vestingPoolName, lockupPeriod, vestingPeriod)
	}
	return nil
}
func (k Keeper) ValidateCloseCampaignParams(ctx sdk.Context, action types.CloseAction, campaignId uint64, owner string) (types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return types.Campaign{}, err
	}
	if err = ValidateOwner(campaign, owner); err != nil {
		return types.Campaign{}, err
	}
	if err = ValidateCampaignEnded(ctx, campaign); err != nil {
		return types.Campaign{}, err
	}
	if err = types.ValidateCloseAction(action); err != nil {
		return types.Campaign{}, err
	}

	return campaign, nil
}

func (k Keeper) ValidateStartCampaignParams(ctx sdk.Context, campaign types.Campaign, owner string) error {
	if err := ValidateOwner(campaign, owner); err != nil {
		return err
	}

	if err := types.ValidateCampaignIsNotEnabled(campaign); err != nil {
		return err
	}
	if err := types.ValidateCampaignEndTimeAfterStartTime(&campaign.StartTime, &campaign.EndTime); err != nil {
		return err
	}
	return ValidateCampaignStartTimeInTheFuture(ctx, &campaign.StartTime)
}

func (k Keeper) ValidateRemoveCampaignParams(ctx sdk.Context, owner string, campaignId uint64) error {
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return err
	}

	if err = ValidateOwner(campaign, owner); err != nil {
		return err
	}

	return types.ValidateCampaignIsNotEnabled(campaign)
}

func (k Keeper) ValidateCampaignExists(ctx sdk.Context, campaignId uint64) (types.Campaign, error) {
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		return types.Campaign{}, errors.Wrapf(c4eerrors.ErrNotExists, "campaign with id %d not found", campaignId)
	}
	return campaign, nil
}

func ValidateOwner(campaign types.Campaign, owner string) error {
	if campaign.Owner != owner {
		return errors.Wrap(sdkerrors.ErrorInvalidSigner, "you are not the campaign owner")
	}
	return nil
}

func ValidateCampaignIsNotDisabled(campaign types.Campaign) error {
	if campaign.Enabled == false {
		return errors.Wrap(c4eerrors.ErrAlreadyExists, "campaign is disabled")
	}
	return nil
}

func ValidateCampaignStartTimeInTheFuture(ctx sdk.Context, startTime *time.Time) error {
	if startTime == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "create claim campaign - start time is nil error")
	}
	if startTime.Before(ctx.BlockTime()) {
		return errors.Wrapf(c4eerrors.ErrParam, "start time in the past error (%s < %s)", startTime, ctx.BlockTime())
	}
	return nil
}

func (k Keeper) ValidateVestingPool(ctx sdk.Context, owner string, vestingPoolName string) error {
	_, found := k.vestingKeeper.GetAccountVestingPool(ctx, owner, vestingPoolName)
	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pool %s not found for address %s", vestingPoolName, owner)
	}
	return nil
}

func ValidateCampaignNotEnded(ctx sdk.Context, campaign types.Campaign) error {
	if ctx.BlockTime().After(campaign.EndTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "campaign with id %d campaign is over (end time - %s < %s)", campaign.Id, campaign.EndTime, ctx.BlockTime())
	}
	return nil
}

func ValidateCampaignEnded(ctx sdk.Context, campaign types.Campaign) error {
	if ctx.BlockTime().Before(campaign.EndTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "campaign with id %d campaign is not over yet (endtime - %s < %s)", campaign.Id, campaign.EndTime, ctx.BlockTime())
	}
	return nil
}
