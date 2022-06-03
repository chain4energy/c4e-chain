package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStoreSignature = "store_signature"

var _ sdk.Msg = &MsgStoreSignature{}

func NewMsgStoreSignature(creator string, storageKey string, signatureJSON string) *MsgStoreSignature {
	return &MsgStoreSignature{
		Creator:       creator,
		StorageKey:    storageKey,
		SignatureJSON: signatureJSON,
	}
}

func (msg *MsgStoreSignature) Route() string {
	return RouterKey
}

func (msg *MsgStoreSignature) Type() string {
	return TypeMsgStoreSignature
}

func (msg *MsgStoreSignature) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStoreSignature) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStoreSignature) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
