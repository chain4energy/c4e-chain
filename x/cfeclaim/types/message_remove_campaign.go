package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveCampaign = "remove_claim_campaign"

var _ sdk.Msg = &MsgRemoveCampaign{}

func NewMsgRemoveCampaign(owner string, campaignId uint64) *MsgRemoveCampaign {
	return &MsgRemoveCampaign{
		Owner:      owner,
		CampaignId: campaignId,
	}
}

func (msg *MsgRemoveCampaign) Route() string {
	return RouterKey
}

func (msg *MsgRemoveCampaign) Type() string {
	return TypeMsgRemoveCampaign
}

func (msg *MsgRemoveCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
