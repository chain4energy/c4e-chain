package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgBuyCertificate = "buy_certificate"
)

var _ sdk.Msg = &MsgBuyCertificate{}

func NewMsgBuyCertificate(buyer string, certificateId uint64) *MsgBuyCertificate {
	return &MsgBuyCertificate{
		Buyer:                    buyer,
		MarketplaceCertificateId: certificateId,
	}
}

func (msg *MsgBuyCertificate) Route() string {
	return RouterKey
}

func (msg *MsgBuyCertificate) Type() string {
	return TypeMsgBuyCertificate
}

func (msg *MsgBuyCertificate) GetSigners() []sdk.AccAddress {
	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{buyer}
}

func (msg *MsgBuyCertificate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyCertificate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid buyer address (%s)", err)
	}
	return nil
}
