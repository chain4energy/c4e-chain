package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgAddMissionToAidropCampaign = "add_mission_to_aidrop_campaign"

var _ sdk.Msg = &MsgAddMissionToAidropCampaign{}

func NewMsgAddMissionToAidropCampaign(owner string, campaignId uint64, name string, description string, missionType MissionType, weight sdk.Dec, claimStartDate time.Time) *MsgAddMissionToAidropCampaign {
	return &MsgAddMissionToAidropCampaign{
		Owner:          owner,
		Name:           name,
		Description:    description,
		CampaignId:     campaignId,
		MissionType:    missionType,
		Weight:         weight,
		ClaimStartDate: &claimStartDate,
	}
}

func (msg *MsgAddMissionToAidropCampaign) Route() string {
	return RouterKey
}

func (msg *MsgAddMissionToAidropCampaign) Type() string {
	return TypeMsgAddMissionToAidropCampaign
}

func (msg *MsgAddMissionToAidropCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddMissionToAidropCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMissionToAidropCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
