package types

import "fmt"

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
