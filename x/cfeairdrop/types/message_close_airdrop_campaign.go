package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCloseAirdropCampaign = "close_airdrop_campaign"

var _ sdk.Msg = &MsgCloseAirdropCampaign{}

func NewMsgCloseAirdropCampaign(owner string, campaignId uint64, airdropCloseAction AirdropCloseAction) *MsgCloseAirdropCampaign {
	return &MsgCloseAirdropCampaign{
		Owner:              owner,
		CampaignId:         campaignId,
		AirdropCloseAction: airdropCloseAction,
	}
}

func (msg *MsgCloseAirdropCampaign) Route() string {
	return RouterKey
}

func (msg *MsgCloseAirdropCampaign) Type() string {
	return TypeMsgCloseAirdropCampaign
}

func (msg *MsgCloseAirdropCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCloseAirdropCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCloseAirdropCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
