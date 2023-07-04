package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateUserCertificates = "create_user_certificates"
	TypeMsgUpdateUserCertificates = "update_user_certificates"
	TypeMsgDeleteUserCertificates = "delete_user_certificates"
)

var _ sdk.Msg = &MsgCreateUserCertificates{}

func NewMsgCreateUserCertificates(owner string, certificates *Certificate) *MsgCreateUserCertificates {
	return &MsgCreateUserCertificates{
		Owner:        owner,
		Certificates: certificates,
	}
}

func (msg *MsgCreateUserCertificates) Route() string {
	return RouterKey
}

func (msg *MsgCreateUserCertificates) Type() string {
	return TypeMsgCreateUserCertificates
}

func (msg *MsgCreateUserCertificates) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgCreateUserCertificates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUserCertificates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUserCertificates{}

func NewMsgUpdateUserCertificates(owner string, id uint64, certificates *Certificate) *MsgUpdateUserCertificates {
	return &MsgUpdateUserCertificates{
		Id:           id,
		Owner:        owner,
		Certificates: certificates,
	}
}

func (msg *MsgUpdateUserCertificates) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUserCertificates) Type() string {
	return TypeMsgUpdateUserCertificates
}

func (msg *MsgUpdateUserCertificates) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgUpdateUserCertificates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUserCertificates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUserCertificates{}

func NewMsgDeleteUserCertificates(owner string, id uint64) *MsgDeleteUserCertificates {
	return &MsgDeleteUserCertificates{
		Id:    id,
		Owner: owner,
	}
}
func (msg *MsgDeleteUserCertificates) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUserCertificates) Type() string {
	return TypeMsgDeleteUserCertificates
}

func (msg *MsgDeleteUserCertificates) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgDeleteUserCertificates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUserCertificates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
