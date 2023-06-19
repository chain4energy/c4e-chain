package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.PayloadHash == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "payload hash cannot be empty")
	}
	return nil
}
