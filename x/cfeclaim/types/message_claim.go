package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaim = "claim"
const InitialMissionId = 0

var _ sdk.Msg = &MsgClaim{}

func NewMsgClaim(claimer string, campaignId uint64, missionId uint64) *MsgClaim {
	return &MsgClaim{
		Claimer:    claimer,
		CampaignId: campaignId,
		MissionId:  missionId,
	}
}

func (msg *MsgClaim) Route() string {
	return RouterKey
}

func (msg *MsgClaim) Type() string {
	return TypeMsgClaim
}

func (msg *MsgClaim) GetSigners() []sdk.AccAddress {
	claimer, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{claimer}
}

func (msg *MsgClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid claimer address (%s)", err)
	}
	return nil
}
