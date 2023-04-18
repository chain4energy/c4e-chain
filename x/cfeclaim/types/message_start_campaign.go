package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartCampaign = "start_claim_campaign"

var _ sdk.Msg = &MsgStartCampaign{}

func NewMsgStartCampaign(owner string, campaignId uint64) *MsgStartCampaign {
	return &MsgStartCampaign{
		Owner:      owner,
		CampaignId: campaignId,
	}
}

func (msg *MsgStartCampaign) Route() string {
	return RouterKey
}

func (msg *MsgStartCampaign) Type() string {
	return TypeMsgStartCampaign
}

func (msg *MsgStartCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
