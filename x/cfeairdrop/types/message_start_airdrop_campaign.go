package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartAirdropCampaign = "start_airdrop_campaign"

var _ sdk.Msg = &MsgStartAirdropCampaign{}

func NewMsgStartAirdropCampaign(owner string, campaignId uint64) *MsgStartAirdropCampaign {
	return &MsgStartAirdropCampaign{
		Owner:      owner,
		CampaignId: campaignId,
	}
}

func (msg *MsgStartAirdropCampaign) Route() string {
	return RouterKey
}

func (msg *MsgStartAirdropCampaign) Type() string {
	return TypeMsgStartAirdropCampaign
}

func (msg *MsgStartAirdropCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartAirdropCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartAirdropCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
