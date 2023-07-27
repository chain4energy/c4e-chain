package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAddCertificateToMarketplace = "add_certificate_to_marketplace"
)

var _ sdk.Msg = &MsgAddCertificateToMarketplace{}

func NewMsgAddCertificateToMarketplace(owner string, certificateId uint64, price sdk.Coins) *MsgAddCertificateToMarketplace {
	return &MsgAddCertificateToMarketplace{
		Owner:         owner,
		CertificateId: certificateId,
		Price:         price,
	}
}

func (msg *MsgAddCertificateToMarketplace) Route() string {
	return RouterKey
}

func (msg *MsgAddCertificateToMarketplace) Type() string {
	return TypeMsgAddCertificateToMarketplace
}

func (msg *MsgAddCertificateToMarketplace) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgAddCertificateToMarketplace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddCertificateToMarketplace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
