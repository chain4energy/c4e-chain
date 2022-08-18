package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferTokens = "transfer_tokens"

var _ sdk.Msg = &MsgTransferTokens{}

func NewMsgTransferTokens(creator string, addressFrom string, addressTo string, amount uint64, tokenId uint64) *MsgTransferTokens {
	return &MsgTransferTokens{
		Creator:     creator,
		AddressFrom: addressFrom,
		AddressTo:   addressTo,
		Amount:      amount,
		TokenId:     tokenId,
	}
}

func (msg *MsgTransferTokens) Route() string {
	return RouterKey
}

func (msg *MsgTransferTokens) Type() string {
	return TypeMsgTransferTokens
}

func (msg *MsgTransferTokens) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferTokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
