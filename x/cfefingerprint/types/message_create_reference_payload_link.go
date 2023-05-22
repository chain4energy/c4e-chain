package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateReferencePayloadLink = "create_reference_payload_link"

var _ sdk.Msg = &MsgCreateReferencePayloadLink{}

func NewMsgCreateReferencePayloadLink(creator string, payloadHash string) *MsgCreateReferencePayloadLink {
	return &MsgCreateReferencePayloadLink{
		Creator:     creator,
		PayloadHash: payloadHash,
	}
}

func (msg *MsgCreateReferencePayloadLink) Route() string {
	return RouterKey
}

func (msg *MsgCreateReferencePayloadLink) Type() string {
	return TypeMsgCreateReferencePayloadLink
}

func (msg *MsgCreateReferencePayloadLink) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateReferencePayloadLink) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateReferencePayloadLink) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
