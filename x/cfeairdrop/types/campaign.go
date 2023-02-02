package types

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const (
	CampaignUnsecified = CampaignType_CAMPAIGN_TYPE_UNSPECIFIED
	CampaignTeamdrop   = CampaignType_TEAMDROP
	CampaignDefault    = CampaignType_DEFAULT
	CampaignSale       = CampaignType_SALE
)

func CampaignTypeFromString(str string) (CampaignType, error) {
	option, ok := MissionType_value[str]
	if !ok {
		return CampaignUnsecified, fmt.Errorf("'%s' is not a valid campaign type, available options: teamdrop/sale/default", str)
	}
	return CampaignType(option), nil
}

// NormalizeCampaignType - normalize user specified vote option
func NormalizeCampaignType(option string) string {
	switch option {
	case "Teamdrop", "teamdrop":
		return CampaignTeamdrop.String()

	case "Sale", "sale":
		return CampaignSale.String()

	case "Default", "default":
		return CampaignDefault.String()

	default:
		return option
	}
}

func CampaignCloseActionFromString(str string) (CampaignCloseAction, error) {
	option, ok := CampaignCloseAction_value[str]
	if !ok {
		return CampaignCloseAction_CLOSE_ACTION_UNSPECIFIED, fmt.Errorf("'%s' is not a valid mission type, available options: initial_claim/vote/delegation", str)
	}
	return CampaignCloseAction(option), nil
}
func GetWhitelistedVestingAccounts() []string {
	return []string{"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj"}
}

func GetTeamdropAccounts() []string {
	return []string{"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj"}
}

func (c *Campaign) IsEnabled(blockTime time.Time) error {
	if !c.Enabled {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaign %d error", c.Id)
	}
	if blockTime.Before(c.StartTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaign %d not started yet (%s < startTime %s) error", c.Id, blockTime, c.StartTime)
	}
	if blockTime.After(c.EndTime) {
		return sdkerrors.Wrapf(ErrCampaignDisabled, "campaign %d has already ended (%s > endTime %s) error", c.Id, blockTime, c.EndTime)
	}
	return nil
}
