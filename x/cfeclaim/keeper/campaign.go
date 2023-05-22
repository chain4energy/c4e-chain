package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (k Keeper) CreateCampaign(ctx sdk.Context, owner string, name string, description string, campaignType types.CampaignType, removableClaimRecords bool,
	feeGrantAmount *math.Int, initialClaimFreeAmount *math.Int, free *sdk.Dec, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration, vestingPoolName string) (*types.Campaign, error) {
	k.Logger(ctx).Debug("create campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)

	validFree, err := validateFreeAmount(free)
	if err != nil {
		return nil, err
	}

	if err = k.ValidateCampaignParams(ctx, name, description, validFree, startTime, endTime, campaignType, owner, vestingPoolName, lockupPeriod, vestingPeriod); err != nil {
		return nil, err
	}
	if err = ValidateCampaignEndTimeInTheFuture(ctx, endTime); err != nil {
		return nil, err
	}
	validdFeegrantAmount, err := validateFeegrantAmount(feeGrantAmount)
	if err != nil {
		return nil, err
	}
	validInitialClaimFreeAmount, err := validateInitialClaimFreeAmount(initialClaimFreeAmount)
	if err != nil {
		return nil, err
	}

	campaign := types.Campaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		CampaignType:           campaignType,
		RemovableClaimRecords:  removableClaimRecords,
		FeegrantAmount:         validdFeegrantAmount,
		InitialClaimFreeAmount: validInitialClaimFreeAmount,
		Free:                   validFree,
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
		return nil, err
	}

	return &campaign, nil
}

func validateInitialClaimFreeAmount(initialClaimFreeAmount *math.Int) (math.Int, error) {
	if initialClaimFreeAmount == nil {
		return math.ZeroInt(), nil
	}
	if initialClaimFreeAmount.IsNil() {
		return math.ZeroInt(), nil
	}

	if initialClaimFreeAmount.IsNegative() {
		return math.ZeroInt(), errors.Wrapf(c4eerrors.ErrParam, "initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}

	return *initialClaimFreeAmount, nil
}

func validateFreeAmount(free *sdk.Dec) (sdk.Dec, error) {
	if free == nil {
		return sdk.ZeroDec(), nil
	}
	if free.IsNil() {
		return sdk.ZeroDec(), nil
	}

	if free.IsNegative() {
		return sdk.ZeroDec(), errors.Wrapf(c4eerrors.ErrParam, "free amount (%s) cannot be negative", free.String())
	}

	return *free, nil
}

func (k Keeper) CloseCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("close campaign", "owner", owner, "campaignId", campaignId)
	campaign, err := k.ValidateCloseCampaignParams(ctx, campaignId, owner)
	if err != nil {
		return err
	}
	campaign.Enabled = false
	k.SetCampaign(ctx, campaign)
	if err = k.sendCampaignCurrentAmountToOwner(ctx, &campaign, campaign.CampaignCurrentAmount); err != nil {
		return err
	}
	if err = k.closeCampaignSendFeegrant(ctx, &campaign); err != nil {
		return err
	}

	return nil
}

func (k Keeper) sendCampaignCurrentAmountToOwner(ctx sdk.Context, campaign *types.Campaign, amount sdk.Coins) error {
	if amount.IsAnyGT(campaign.CampaignCurrentAmount) {
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

	campaign.CampaignCurrentAmount = campaign.CampaignCurrentAmount.Sub(amount...)
	k.SetCampaign(ctx, *campaign)
	return nil
}

func (k Keeper) EnableCampaign(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) error {
	k.Logger(ctx).Debug("start campaign", "owner", owner, "campaignId", campaignId)

	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return err
	}

	if startTime != nil {
		campaign.StartTime = *startTime
	}
	if endTime != nil {
		campaign.EndTime = *endTime
	}

	err = k.ValidateEnableCampaignParams(ctx, campaign, owner)
	if err != nil {
		return err
	}

	campaign.Enabled = true
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) RemoveCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("remove campaign", "owner", owner, "campaignId", campaignId)

	campaign, err := k.ValidateRemoveCampaignParams(ctx, owner, campaignId)
	if err != nil {
		k.Logger(ctx).Debug("remove campaign", "err", err.Error())
		return err
	}

	if err = k.sendCampaignCurrentAmountToOwner(ctx, campaign, campaign.CampaignCurrentAmount); err != nil {
		return err
	}
	if err = k.closeCampaignSendFeegrant(ctx, campaign); err != nil {
		return err
	}

	k.removeCampaign(ctx, campaignId)
	k.RemoveAllMissionForCampaign(ctx, campaignId)
	return nil
}
func (k Keeper) ValidateCampaignParams(ctx sdk.Context, name string, description string, free sdk.Dec, startTime *time.Time, endTime *time.Time,
	campaignType types.CampaignType, owner string, vestingPoolName string, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	if err := types.ValidateCreateCampaignParams(name, description, startTime, endTime, campaignType, vestingPoolName); err != nil {
		return err
	}

	if campaignType == types.VestingPoolCampaign {
		return k.ValidateCampaignWhenAddedFromVestingPool(ctx, owner, vestingPoolName, lockupPeriod, vestingPeriod, free)
	}
	return nil
}
func (k Keeper) ValidateCloseCampaignParams(ctx sdk.Context, campaignId uint64, owner string) (types.Campaign, error) {
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

	return campaign, nil
}

func (k Keeper) ValidateEnableCampaignParams(ctx sdk.Context, campaign types.Campaign, owner string) error {
	if err := ValidateOwner(campaign, owner); err != nil {
		return err
	}

	if err := types.ValidateCampaignIsNotEnabled(campaign); err != nil {
		return err
	}
	if err := types.ValidateCampaignEndTimeAfterStartTime(&campaign.StartTime, &campaign.EndTime); err != nil {
		return err
	}
	return ValidateCampaignEndTimeInTheFuture(ctx, &campaign.EndTime)
}

func (k Keeper) ValidateRemoveCampaignParams(ctx sdk.Context, owner string, campaignId uint64) (*types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return nil, err
	}

	if err = ValidateOwner(campaign, owner); err != nil {
		return nil, err
	}

	return &campaign, types.ValidateCampaignIsNotEnabled(campaign)
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
		return errors.Wrap(c4eerrors.ErrWrongSigner, "you are not the campaign owner")
	}
	return nil
}

func ValidateCampaignEndTimeInTheFuture(ctx sdk.Context, endTime *time.Time) error {
	if endTime == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "create claim campaign - start time is nil error")
	}
	if endTime.Before(ctx.BlockTime()) {
		return errors.Wrapf(c4eerrors.ErrParam, "end time in the past error (%s < %s)", endTime, ctx.BlockTime())
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
