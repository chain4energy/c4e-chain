package types

import (
	"cosmossdk.io/errors"
	"fmt"
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

func (c *Campaign) IsActive(blockTime time.Time) error {
	if !c.Enabled {
		return errors.Wrapf(ErrCampaignDisabled, "campaign %d error", c.Id)
	}
	if blockTime.Before(c.StartTime) {
		return errors.Wrapf(ErrCampaignDisabled, "campaign %d not started yet (%s < startTime %s) error", c.Id, blockTime, c.StartTime)
	}
	if blockTime.After(c.EndTime) {
		return errors.Wrapf(ErrCampaignDisabled, "campaign %d has already ended (%s > endTime %s) error", c.Id, blockTime, c.EndTime)
	}
	return nil
}

func ValidateCampaignIsNotEnabled(campaign Campaign) error {
	if campaign.Enabled == true {
		return ErrCampaignEnabled
	}
	return nil
}
