package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPublishReferencePayloadLink = "publish_reference_payload_link"

var _ sdk.Msg = &MsgPublishReferencePayloadLink{}

func NewMsgPublishReferencePayloadLink(creator string, key string, value string) *MsgPublishReferencePayloadLink {
	return &MsgPublishReferencePayloadLink{
		Creator: creator,
		Key:     key,
		Value:   value,
	}
}

func (msg *MsgPublishReferencePayloadLink) Route() string {
	return RouterKey
}

func (msg *MsgPublishReferencePayloadLink) Type() string {
	return TypeMsgPublishReferencePayloadLink
}

func (msg *MsgPublishReferencePayloadLink) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPublishReferencePayloadLink) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPublishReferencePayloadLink) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
