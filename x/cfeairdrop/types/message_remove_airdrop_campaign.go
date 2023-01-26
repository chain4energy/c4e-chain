package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveAirdropCampaign = "remove_airdrop_campaign"

var _ sdk.Msg = &MsgRemoveAirdropCampaign{}

func NewMsgRemoveAirdropCampaign(owner string, campaignId uint64) *MsgRemoveAirdropCampaign {
	return &MsgRemoveAirdropCampaign{
		Owner:      owner,
		CampaignId: campaignId,
	}
}

func (msg *MsgRemoveAirdropCampaign) Route() string {
	return RouterKey
}

func (msg *MsgRemoveAirdropCampaign) Type() string {
	return TypeMsgRemoveAirdropCampaign
}

func (msg *MsgRemoveAirdropCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveAirdropCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveAirdropCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
