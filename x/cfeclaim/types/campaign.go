package types

import (
	"cosmossdk.io/errors"
	"fmt"
	"time"
)

// Campaign types
const (
	UnspecifiedCampaign = CampaignType_CAMPAIGN_TYPE_UNSPECIFIED
	DynamicCampaign     = CampaignType_DYNAMIC
	DefaultCampaign     = CampaignType_DEFAULT
	VestingPoolCampaign = CampaignType_VESTING_POOL
)

// Campaign close action types
const (
	CloseActionUnspecified   = CloseAction_CLOSE_ACTION_UNSPECIFIED
	CloseSendToCommunityPool = CloseAction_SEND_TO_COMMUNITY_POOL
	CloseBurn                = CloseAction_BURN
	CloseSendToOwner         = CloseAction_SEND_TO_OWNER
)

func CampaignTypeFromString(str string) (CampaignType, error) {
	option, ok := CampaignType_value[str]
	if !ok {
		return UnspecifiedCampaign, fmt.Errorf("'%s' is not a valid campaign type, available options: default/dynamic/vesting_pool", str)
	}
	return CampaignType(option), nil
}

// NormalizeCampaignType - normalize user specified vote option
func NormalizeCampaignType(option string) string {
	switch option {
	case "Dynamic", "dynamic", "DYNAMIC":
		return DynamicCampaign.String()

	case "VestingPool", "VESTING_POOL", "vesting_pool":
		return VestingPoolCampaign.String()

	case "Default", "default", "DEFAULT":
		return DefaultCampaign.String()

	default:
		return option
	}
}

// NormalizeCloseAction - normalize user specified vote option
func NormalizeCloseAction(option string) string {
	switch option {
	case "SendToCommunityPool", "sendtocommunitypool":
		return CloseSendToCommunityPool.String()

	case "Burn", "burn":
		return CloseBurn.String()

	case "SendToOwner", "sendtoowner":
		return CloseSendToOwner.String()

	default:
		return option
	}
}

func CloseActionFromString(str string) (CloseAction, error) {
	option, ok := CloseAction_value[str]
	if !ok {
		return CloseAction_CLOSE_ACTION_UNSPECIFIED, fmt.Errorf("'%s' is not a valid mission type, available options: initial_claim/vote/delegation", str)
	}
	return CloseAction(option), nil
}

func GetWhitelistedVestingAccounts() []string {
	return []string{"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj"}
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
