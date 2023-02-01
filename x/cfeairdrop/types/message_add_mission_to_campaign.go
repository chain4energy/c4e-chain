package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgAddMissionToAidropCampaign = "add_mission_to_aidrop_campaign"

var _ sdk.Msg = &MsgAddMissionToCampaign{}

func NewMsgAddMissionToAidropCampaign(owner string, campaignId uint64, name string, description string, missionType MissionType, weight *sdk.Dec, claimStartDate *time.Time) *MsgAddMissionToCampaign {
	return &MsgAddMissionToCampaign{
		Owner:          owner,
		Name:           name,
		Description:    description,
		CampaignId:     campaignId,
		MissionType:    missionType,
		Weight:         weight,
		ClaimStartDate: claimStartDate,
	}
}

func NewInitialMission(campaignId uint64) *Mission {
	return &Mission{
		CampaignId:  campaignId,
		Name:        "Initial mission",
		Description: "Initial mission - basic mission that must be claimed first",
		MissionType: MissionInitialClaim,
		Weight:      sdk.ZeroDec(),
	}
}

func (msg *MsgAddMissionToCampaign) Route() string {
	return RouterKey
}

func (msg *MsgAddMissionToCampaign) Type() string {
	return TypeMsgAddMissionToAidropCampaign
}

func (msg *MsgAddMissionToCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddMissionToCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMissionToCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
