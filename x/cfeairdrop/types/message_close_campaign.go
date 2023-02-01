package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCloseCampaign = "close_airdrop_campaign"

var _ sdk.Msg = &MsgCloseCampaign{}

func NewMsgCloseCampaign(owner string, campaignId uint64, campaignCloseAction CampaignCloseAction) *MsgCloseCampaign {
	return &MsgCloseCampaign{
		Owner:               owner,
		CampaignId:          campaignId,
		CampaignCloseAction: campaignCloseAction,
	}
}

func (msg *MsgCloseCampaign) Route() string {
	return RouterKey
}

func (msg *MsgCloseCampaign) Type() string {
	return TypeMsgCloseCampaign
}

func (msg *MsgCloseCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCloseCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCloseCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
