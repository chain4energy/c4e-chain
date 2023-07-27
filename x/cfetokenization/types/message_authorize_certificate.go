package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAuthorizeCertificate = "authorize_certificate"
)

var _ sdk.Msg = &MsgAuthorizeCertificate{}

func NewMsgAuthorizeCertificate(authorizer string, userAddress string, certificateId uint64) *MsgAuthorizeCertificate {
	return &MsgAuthorizeCertificate{
		Authorizer:    authorizer,
		UserAddress:   userAddress,
		CertificateId: certificateId,
	}
}

func (msg *MsgAuthorizeCertificate) Route() string {
	return RouterKey
}

func (msg *MsgAuthorizeCertificate) Type() string {
	return TypeMsgAuthorizeCertificate
}

func (msg *MsgAuthorizeCertificate) GetSigners() []sdk.AccAddress {
	authorizer, err := sdk.AccAddressFromBech32(msg.Authorizer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authorizer}
}

func (msg *MsgAuthorizeCertificate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAuthorizeCertificate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authorizer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authorizer address (%s)", err)
	}
	return nil
}
