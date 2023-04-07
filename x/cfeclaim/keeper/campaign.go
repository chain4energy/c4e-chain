package keeper

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/exp/slices"
	"strconv"
	"time"
)

func (k Keeper) CreateCampaign(ctx sdk.Context, owner string, name string, description string, campaignType types.CampaignType, feeGrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	k.Logger(ctx).Debug("create campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)

	log := k.Logger(ctx).With("Create campaign")
	if err := ValidateCampaignCreateParams(log, name, description, startTime, endTime, campaignType, owner, ctx); err != nil {
		return err
	}

	feeGrantAmount, err := GetFeeGrantAmount(log, feeGrantAmount)
	if err != nil {
		return err
	}
	initialClaimFreeAmount, err = GetInitialClaimFreeAmount(log, initialClaimFreeAmount)
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

func ValidateCampaignCreateParams(log log.Logger, name string, description string, startTime *time.Time, endTime *time.Time,
	campaignType types.CampaignType, owner string, ctx sdk.Context) error {
	if err := ValidateCampaignName(log, name); err != nil {
		return err
	}
	if err := ValidateCampaignDescription(log, description); err != nil {
		return err
	}
	if err := ValidateCampaignStartTime(log, startTime, ctx); err != nil {
		return err
	}
	if err := ValidateCampaignEndTime(log, startTime, endTime); err != nil {
		return err
	}
	if err := ValidateCampaignType(log, campaignType, owner); err != nil {
		return err
	}
	return nil
}

func ValidateCampaignEditParams(log log.Logger, name string, description string, startTime *time.Time, endTime *time.Time,
	campaignType types.CampaignType, owner string, ctx sdk.Context) error {
	if err := ValidateCampaignName(log, name); err != nil {
		return err
	}
	if err := ValidateCampaignDescription(log, description); err != nil {
		return err
	}
	if err := ValidateCampaignStartTime(log, startTime, ctx); err != nil {
		return err
	}
	if err := ValidateCampaignEndTime(log, startTime, endTime); err != nil {
		return err
	}
	if err := ValidateCampaignType(log, campaignType, owner); err != nil {
		return err
	}
	return nil
}

func GetFeeGrantAmount(logger log.Logger, feeGrantAmount *sdk.Int) (*sdk.Int, error) {
	if feeGrantAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		feeGrantAmount = &zeroInt
	}

	if feeGrantAmount.IsNegative() {
		logger.Debug("initial feegrant amount cannot be negative", "initialClaimFreeAmount", feeGrantAmount)
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParam, "feegrant amount (%s) cannot be negative", feeGrantAmount.String())
	}

	return feeGrantAmount, nil
}

func SwapToNewParams(log log.Logger, name string, description string, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration,
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

	if err := ValidateCampaignCreateParams(log, name, description, startTime, endTime,
		campaign.CampaignType, campaign.Owner, ctx); err != nil {
		return types.Campaign{}, err
	}

	return campaign, nil
}

func (k Keeper) EditCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	k.Logger(ctx).Debug("edit claim campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)

	logger := ctx.Logger().With("Edit campaign")

	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return err
	}

	if err = ValidateCampaignIsNotEnabled(logger, campaign); err != nil {
		return err
	}

	campaign, err = SwapToNewParams(logger, name, description, startTime, endTime, lockupPeriod, vestingPeriod, campaign, ctx)
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
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("edit campaign emit event error", "event", event, "error", err.Error())
	}
	return nil
}

func (k Keeper) CloseCampaign(ctx sdk.Context, owner string, campaignId uint64, campaignCloseAction types.CampaignCloseAction) error {
	logger := ctx.Logger().With("close claim campaign", "owner", owner, "campaignId", campaignId, "campaignCloseAction", campaignCloseAction)

	campaign, validationResult := k.ValidateCloseCampaignParams(logger, campaignId, ctx, owner)
	if validationResult != nil {
		return validationResult
	}

	campaignAmountLeft, _ := k.GetCampaignAmountLeft(ctx, campaign.Id)

	if err := k.campaignCloseActionSwitch(ctx, campaignCloseAction, &campaign, campaignAmountLeft.Amount); err != nil {
		return err
	}

	campaign.Enabled = false
	k.SetCampaign(ctx, campaign)
	k.DecrementCampaignAmountLeft(ctx, campaignId, campaignAmountLeft.Amount)
	return nil
}

func (k Keeper) campaignCloseActionSwitch(ctx sdk.Context, campaignCloseAction types.CampaignCloseAction, campaign *types.Campaign, campaignAmountLeft sdk.Coins) error {
	switch campaignCloseAction {
	case types.CampaignCloseSendToCommunityPool:
		return k.campaignCloseSendToCommunityPool(ctx, campaign, campaignAmountLeft)
	case types.CampaignCloseBurn:
		return k.campaignCloseBurn(ctx, campaign, campaignAmountLeft)
	case types.CampaignCloseSendToOwner:
		return k.campaignCloseSendToOwner(ctx, campaign, campaignAmountLeft)
	default:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "wrong campaign close action type")
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
	return nil
}

func (k Keeper) ValidateCloseCampaignParams(logger log.Logger, campaignId uint64, ctx sdk.Context, owner string) (types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return types.Campaign{}, err
	}

	if err = ValidateOwner(logger, campaign, owner); err != nil {
		return types.Campaign{}, err
	}

	if err = ValidateCampaignEnded(logger, ctx, campaign); err != nil {
		return types.Campaign{}, err
	}

	return campaign, nil
}

func ValidateCampaignNotEnded(logger log.Logger, ctx sdk.Context, campaign types.Campaign) error {
	if ctx.BlockTime().After(campaign.EndTime) {
		logger.Debug("close claim campaign campaign is not over yet", "startTime", campaign.StartTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "campaign with id %d campaign is over (end time - %s < %s)", campaign.Id, campaign.EndTime, ctx.BlockTime())
	}
	return nil
}

func ValidateCampaignEnded(logger log.Logger, ctx sdk.Context, campaign types.Campaign) error {
	if ctx.BlockTime().Before(campaign.EndTime) {
		logger.Debug("campaign is over", "startTime", campaign.StartTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "campaign with id %d campaign is not over yet (endtime - %s < %s)", campaign.Id, campaign.EndTime, ctx.BlockTime())
	}
	return nil
}

func ValidateCampaignStarted(logger log.Logger, ctx sdk.Context, campaign types.Campaign) error {
	if !campaign.EndTime.After(ctx.BlockTime()) {
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "campaign with id %d has not started yet (startTime - %s < %s)", campaign.Id, campaign.StartTime, ctx.BlockTime())
	}
	return nil
}

func (k Keeper) StartCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	logger := ctx.Logger().With("start claim campaign", "owner", owner, "campaignId", campaignId)

	campaign, validationResult := k.ValidateStartCampaignParams(logger, campaignId, ctx, owner)
	if validationResult != nil {
		return validationResult
	}

	campaign.Enabled = true
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) RemoveCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("Remove claim campaign", "owner", owner, "campaignId", campaignId)
	validationResult := k.ValidateRemoveCampaignParams(ctx, owner, campaignId)
	if validationResult != nil {
		return validationResult
	}

	k.removeCampaign(ctx, campaignId)
	k.RemoveAllMissionForCampaign(ctx, campaignId)
	return nil
}

func (k Keeper) ValidateStartCampaignParams(logger log.Logger, campaignId uint64, ctx sdk.Context, owner string) (types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return types.Campaign{}, err
	}

	if err = ValidateOwner(logger, campaign, owner); err != nil {
		return types.Campaign{}, err
	}

	if err = ValidateCampaignIsNotEnabled(logger, campaign); err != nil {
		return types.Campaign{}, err
	}

	if err = ValidateCampaignStart(ctx, campaign, logger); err != nil {
		return types.Campaign{}, err
	}

	return campaign, nil
}

func (k Keeper) ValidateRemoveCampaignParams(ctx sdk.Context, owner string, campaignId uint64) error {
	logger := ctx.Logger().With("Remove campaign validation")

	campaign, err := k.ValidateCampaignExists(logger, campaignId, ctx)
	if err != nil {
		return err
	}

	if err = ValidateOwner(logger, campaign, owner); err != nil {
		return err
	}

	if err = ValidateCampaignIsNotEnabled(logger, campaign); err != nil {
		return err
	}

	return nil
}

func (k Keeper) ValidateCampaignExists(log log.Logger, campaignId uint64, ctx sdk.Context) (types.Campaign, error) {
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		log.Debug("campaign not exist", "campaignId", campaignId)
		return types.Campaign{}, sdkerrors.Wrapf(c4eerrors.ErrNotExists, "campaign with id %d not found", campaignId)
	}
	return campaign, nil
}

func ValidateOwner(log log.Logger, campaign types.Campaign, owner string) error {
	if campaign.Owner != owner {
		log.Debug("you are not campaign owner")
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "you are not the campaign owner")
	}
	return nil
}

func ValidateCampaignIsNotEnabled(log log.Logger, campaign types.Campaign) error {
	if campaign.Enabled == true {
		log.Debug("campaign is enabled")
		return sdkerrors.Wrap(c4eerrors.ErrAlreadyExists, "campaign is enabled")
	}
	return nil
}

func ValidateCampaignIsNotDisabled(log log.Logger, campaign types.Campaign) error {
	if campaign.Enabled == false {
		log.Debug("campaign is disabled")
		return sdkerrors.Wrap(c4eerrors.ErrAlreadyExists, "campaign is disabled")
	}
	return nil
}

func ValidateCampaignName(log log.Logger, name string) error {
	if name == "" {
		log.Debug("param err, campaign name is empty")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "campaign name is empty")
	}
	return nil
}

func ValidateCampaignDescription(log log.Logger, description string) error {
	if description == "" {
		log.Debug("param err, description is empty")
		return sdkerrors.Wrap(c4eerrors.ErrParam, "description is empty")
	}
	return nil
}

func ValidateCampaignStartTime(log log.Logger, startTime *time.Time, ctx sdk.Context) error {
	if startTime == nil {
		log.Debug("param err, start time is nil")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create claim campaign - start time is nil error")
	}
	if startTime.Before(ctx.BlockTime()) {
		log.Debug("param err, start time in the past", "startTime", startTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "start time in the past error (%s < %s)", startTime, ctx.BlockTime())
	}
	return nil
}

func ValidateCampaignEndTime(log log.Logger, startTime *time.Time, endTime *time.Time) error {
	if endTime == nil {
		log.Debug("param err,  end time is nil")
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "end time is nil error")
	}
	if startTime.After(*endTime) {
		log.Debug("param err,  start time is after end time", "startTime", startTime, "endTime", endTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "start time is after end time error (%s > %s)", startTime, endTime)
	}

	return nil
}

func ValidateCampaignType(log log.Logger, campaignType types.CampaignType, owner string) error {
	if campaignType == types.CampaignTeamdrop {
		if !slices.Contains(types.GetTeamdropAccounts(), owner) {
			log.Debug("param err, this campaign type can be created only by specific accounts", "owner", owner)
			return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "TeamDrop campaigns can be created only by specific accounts")
		}
	}

	return nil
}

func ValidateCampaignStart(ctx sdk.Context, campaign types.Campaign, logger log.Logger) error {
	if campaign.StartTime.Before(ctx.BlockTime()) {
		logger.Debug("Campaign start time in the past", "startTime", campaign.StartTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "campaign with id %d start time in the past error (%s < %s)", campaign.Id, campaign.StartTime, ctx.BlockTime())
	}
	return nil
}

func GetInitialClaimFreeAmount(log log.Logger, initialClaimFreeAmount *sdk.Int) (*sdk.Int, error) {
	if initialClaimFreeAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		initialClaimFreeAmount = &zeroInt
	}

	if initialClaimFreeAmount.IsNegative() {
		log.Debug("initial claim free amount cannot be negative", "initialClaimFreeAmount", initialClaimFreeAmount)
		return nil, sdkerrors.Wrapf(c4eerrors.ErrParam, "initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}

	return initialClaimFreeAmount, nil
}
