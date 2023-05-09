package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"golang.org/x/exp/slices"
	"time"
)

func (k Keeper) CreateCampaign(ctx sdk.Context, owner string, name string, description string, campaignType types.CampaignType, feeGrantAmount *math.Int, initialClaimFreeAmount *math.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration, vestingPoolName string) (*types.Campaign, error) {
	k.Logger(ctx).Debug("create campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)

	if err := k.ValidateCampaignParams(ctx, name, description, startTime, endTime, campaignType, owner, vestingPoolName, lockupPeriod, vestingPeriod); err != nil {
		return nil, err
	}

	feeGrantAmount, err := validateFeegrantAmount(feeGrantAmount)
	if err != nil {
		return nil, err
	}

	initialClaimFreeAmount, err = validateInitialClaimFreeAmount(initialClaimFreeAmount)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &campaign, nil
}

func validateInitialClaimFreeAmount(initialClaimFreeAmount *math.Int) (*math.Int, error) {
	if initialClaimFreeAmount == nil {
		zeroInt := sdk.ZeroInt()
		initialClaimFreeAmount = &zeroInt
	}
	if initialClaimFreeAmount.IsNil() {
		zeroInt := sdk.ZeroInt()
		initialClaimFreeAmount = &zeroInt
	}

	if initialClaimFreeAmount.IsNegative() {
		return nil, errors.Wrapf(c4eerrors.ErrParam, "initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}

	return initialClaimFreeAmount, nil
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

	if err = k.closeCampaignSendFeegrant(ctx, closeAction, &campaign); err != nil {
		return err
	}

	campaign.Enabled = false
	k.SetCampaign(ctx, campaign)
	k.DecrementCampaignAmountLeft(ctx, campaignId, campaignAmountLeft.Amount)
	return nil
}

func (k Keeper) closeActionSwitch(ctx sdk.Context, CloseAction types.CloseAction, campaign *types.Campaign, amount sdk.Coins) error {
	switch CloseAction {
	case types.CloseSendToCommunityPool:
		if campaign.CampaignType == types.VestingPoolCampaign || slices.Contains(types.GetWhitelistedVestingAccounts(), campaign.Owner) {
			return errors.Wrap(sdkerrors.ErrInvalidType, "in the case of sale campaigns and campaigns created from whitelist vesting accounts, it is not possible to use sendToCommunityPool close action")
		}
		return k.distributionKeeper.FundCommunityPool(ctx, amount, authtypes.NewModuleAddress(types.ModuleName))
	case types.CloseBurn:
		return k.bankKeeper.BurnCoins(ctx, types.ModuleName, amount)
	case types.CloseSendToOwner:
		return k.closeSendToOwner(ctx, campaign, amount)
	}
	return errors.Wrap(sdkerrors.ErrInvalidType, "wrong campaign close action type")
}

func (k Keeper) closeSendToOwner(ctx sdk.Context, campaign *types.Campaign, campaignAmountLeft sdk.Coins) error {
	if campaign.CampaignType == types.VestingPoolCampaign {
		return k.vestingKeeper.SendFromModuleToVestingPool(ctx, campaign.Owner, campaign.VestingPoolName, campaignAmountLeft, types.ModuleName)
	} else {
		ownerAddress, _ := sdk.AccAddressFromBech32(campaign.Owner)
		if slices.Contains(types.GetWhitelistedVestingAccounts(), campaign.Owner) { // TODO: probably delete
			ownerAccount := k.accountKeeper.GetAccount(ctx, ownerAddress)
			if ownerAccount == nil {
				return errors.Wrapf(c4eerrors.ErrNotExists, "account %s doesn't exist", ownerAddress)
			}
			vestingAcc := ownerAccount.(*vestingtypes.ContinuousVestingAccount)
			vestingAcc.OriginalVesting = vestingAcc.OriginalVesting.Add(campaignAmountLeft...)

			k.accountKeeper.SetAccount(ctx, vestingAcc)
			return k.bankKeeper.BurnCoins(ctx, types.ModuleName, campaignAmountLeft)
		}

		return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, campaignAmountLeft)
	}
}

func (k Keeper) StartCampaign(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) error {
	k.Logger(ctx).Debug("start campaign", "owner", owner, "campaignId", campaignId)

	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return err
	}
	if campaign.CampaignType == types.VestingPoolCampaign {
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
	if err := types.ValidateCreateCampaignParams(name, description, startTime, endTime, campaignType, owner, vestingPoolName); err != nil {
		return err
	}

	if campaignType == types.VestingPoolCampaign {
		return k.ValidateCampaignWhenAddedFromVestingPool(ctx, owner, vestingPoolName, lockupPeriod, vestingPeriod)
	} else {
		return ValidateCampaignStartTimeInTheFuture(ctx, startTime)
	}
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
