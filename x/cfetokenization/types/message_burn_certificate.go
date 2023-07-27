package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgBurnCertificate = "burn_certificate"
)

var _ sdk.Msg = &MsgBurnCertificate{}

func NewMsgBurnCertificate(owner string, certificateId string) *MsgBurnCertificate {
	return &MsgBurnCertificate{
		Owner:         owner,
		CertificateId: certificateId,
	}
}

func (msg *MsgBurnCertificate) Route() string {
	return RouterKey
}

func (msg *MsgBurnCertificate) Type() string {
	return TypeMsgBurnCertificate
}

func (msg *MsgBurnCertificate) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgBurnCertificate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnCertificate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
