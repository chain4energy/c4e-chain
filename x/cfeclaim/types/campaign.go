package types

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/v2/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

// Campaign types
const (
	UnspecifiedCampaign = CampaignType_CAMPAIGN_TYPE_UNSPECIFIED
	DefaultCampaign     = CampaignType_DEFAULT
	VestingPoolCampaign = CampaignType_VESTING_POOL
)

func CampaignTypeFromString(str string) (CampaignType, error) {
	option, ok := CampaignType_value[str]
	if !ok {
		return UnspecifiedCampaign, fmt.Errorf("'%s' is not a valid campaign type, available options: default/vesting_pool", str)
	}
	return CampaignType(option), nil
}

// NormalizeCampaignType - normalize user specified vote option
func NormalizeCampaignType(option string) string {
	switch option {
	case "VestingPool", "VESTING_POOL", "vesting_pool":
		return VestingPoolCampaign.String()

	case "Default", "default", "DEFAULT":
		return DefaultCampaign.String()

	default:
		return option
	}
}

func (c *Campaign) ValidateIsActive(blockTime time.Time) error {
	if !c.Enabled {
		return errors.Wrapf(ErrCampaignDisabled, "campaign %d", c.Id)
	}
	if blockTime.Before(c.StartTime) {
		return errors.Wrapf(ErrCampaignDisabled, "campaign %d not started yet (%s < startTime %s)", c.Id, blockTime, c.StartTime)
	}
	if blockTime.After(c.EndTime) {
		return errors.Wrapf(ErrCampaignDisabled, "campaign %d has already ended (%s > endTime %s)", c.Id, blockTime, c.EndTime)
	}
	return nil
}

func (c *Campaign) ValidateNotEnabled() error {
	if c.Enabled {
		return ErrCampaignEnabled
	}
	return nil
}

func (c *Campaign) ValidateIsEnabled() error {
	if !c.Enabled {
		return ErrCampaignDisabled
	}
	return nil
}

func ValidateCreateCampaignParams(name string, feegrantAmount math.Int,
	initialClaimFreeAmount math.Int, free sdk.Dec, startTime time.Time, endTime time.Time,
	campaignType CampaignType, lockupPeriod time.Duration, vestingPeriod time.Duration, vestingPoolName string) error {
	if name == "" {
		return errors.Wrap(c4eerrors.ErrParam, "campaign name is empty")
	}
	if lockupPeriod < 0 {
		return errors.Wrap(c4eerrors.ErrParam, "lockup period cannot be negative")
	}
	if vestingPeriod < 0 {
		return errors.Wrap(c4eerrors.ErrParam, "vesting period cannot be negative")
	}
	if err := validateFreeAmount(free); err != nil {
		return err
	}
	if err := validateFeegrantAmount(feegrantAmount); err != nil {
		return err
	}
	if err := validateInitialClaimFreeAmount(initialClaimFreeAmount); err != nil {
		return err
	}
	if err := ValidateCampaignTimes(startTime, endTime); err != nil {
		return err
	}
	return validateCampaignType(campaignType, vestingPoolName)
}

func ValidateCampaignTimes(startTime time.Time, endTime time.Time) error {
	if startTime.After(endTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "start time is after end time (%s > %s)", startTime, endTime)
	}
	if startTime.Equal(endTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "start time is equal to end time (%s = %s)", startTime, endTime)
	}
	return nil
}

func validateCampaignType(campaignType CampaignType, vestingPoolName string) error {
	switch campaignType {
	case VestingPoolCampaign:
		if vestingPoolName == "" {
			return errors.Wrap(c4eerrors.ErrParam, "for VESTING_POOL type campaigns, the vesting pool name must be provided")
		}
		return nil
	case DefaultCampaign:
		if vestingPoolName != "" {
			return errors.Wrap(c4eerrors.ErrParam, "vesting pool name can be set only for VESTING_POOL type campaigns")
		}
		return nil
	}

	return errors.Wrap(sdkerrors.ErrInvalidType, "wrong campaign type")
}

func validateFeegrantAmount(feeGrantAmount math.Int) error {
	if feeGrantAmount.IsNil() {
		return nil
	}

	if feeGrantAmount.IsNegative() {
		return errors.Wrapf(c4eerrors.ErrParam, "feegrant amount (%s) cannot be negative", feeGrantAmount.String())
	}

	return nil
}

func validateFreeAmount(free sdk.Dec) error {
	if free.IsNil() {
		return nil
	}

	if free.IsNegative() {
		return errors.Wrapf(c4eerrors.ErrParam, "free amount (%s) cannot be negative", free.String())
	}

	return nil
}

func validateInitialClaimFreeAmount(initialClaimFreeAmount math.Int) error {
	if initialClaimFreeAmount.IsNil() {
		return nil
	}

	if initialClaimFreeAmount.IsNegative() {
		return errors.Wrapf(c4eerrors.ErrParam, "initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}

	return nil
}

func (c *Campaign) ValidateOwner(owner string) error {
	if c.Owner != owner {
		return errors.Wrapf(sdkerrors.ErrorInvalidSigner, "address %s is not owner of campaign with id %d", owner, c.Id)
	}
	return nil
}

func (c *Campaign) ValidateRemoveCampaignParams(owner string) error {
	if err := c.ValidateOwner(owner); err != nil {
		return err
	}
	return c.ValidateNotEnabled()
}

func (c *Campaign) ValidateCloseCampaignParams(owner string) error {
	if err := c.ValidateOwner(owner); err != nil {
		return err
	}
	return c.ValidateIsEnabled()
}

func (c *Campaign) ValidateEnableCampaignParams(owner string) error {
	if err := c.ValidateOwner(owner); err != nil {
		return err
	}
	if err := c.ValidateNotEnabled(); err != nil {
		return err
	}
	if err := ValidateCampaignTimes(c.StartTime, c.EndTime); err != nil {
		return err
	}
	return nil
}

func (c *Campaign) ValidateEndTimeAfterBlockTime(blockTime time.Time) error {
	if c.EndTime.Before(blockTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "end time in the past error (%s < %s)", c.EndTime, blockTime)
	}
	return nil
}

func (c *Campaign) ValidateNotEnded(blockTime time.Time) error {
	if blockTime.After(c.EndTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "campaign with id %d campaign is over (end time - %s < %s)", c.Id, c.EndTime, blockTime)
	}
	return nil
}

func (c *Campaign) ValidateEnded(blockTime time.Time) error {
	if blockTime.Before(c.EndTime) || blockTime.Equal(c.EndTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "campaign with id %d campaign is not over yet (endtime - %s < %s)", c.Id, c.EndTime, blockTime)
	}
	return nil
}

func (c *Campaign) DecrementCampaignCurrentAmount(coins sdk.Coins) {
	c.CampaignCurrentAmount = c.CampaignCurrentAmount.Sub(coins...)
}

func (c *Campaign) IncrementCampaignCurrentAmount(coins sdk.Coins) {
	c.CampaignCurrentAmount = c.CampaignCurrentAmount.Add(coins...)
}

func (c *Campaign) DecrementCampaignTotalAmount(coins sdk.Coins) {
	c.CampaignTotalAmount = c.CampaignTotalAmount.Sub(coins...)
}

func (c *Campaign) IncrementCampaignTotalAmount(coins sdk.Coins) {
	c.CampaignTotalAmount = c.CampaignTotalAmount.Add(coins...)
}

func (c *Campaign) ValidateRemovableClaimRecords() error {
	if c.Enabled && !c.RemovableClaimRecords {
		return errors.Wrap(sdkerrors.ErrInvalidType, "campaign must have RemovableClaimRecords flag set to true to be able to delete its entries")
	}
	return nil
}
