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
