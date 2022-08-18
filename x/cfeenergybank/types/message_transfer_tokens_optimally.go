package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferTokensOptimally = "transfer_tokens_optimally"

var _ sdk.Msg = &MsgTransferTokensOptimally{}

func NewMsgTransferTokensOptimally(creator string, addressFrom string, addressTo string, amount string, tokenName string) *MsgTransferTokensOptimally {
	return &MsgTransferTokensOptimally{
		Creator:     creator,
		AddressFrom: addressFrom,
		AddressTo:   addressTo,
		Amount:      amount,
		TokenName:   tokenName,
	}
}

func (msg *MsgTransferTokensOptimally) Route() string {
	return RouterKey
}

func (msg *MsgTransferTokensOptimally) Type() string {
	return TypeMsgTransferTokensOptimally
}

func (msg *MsgTransferTokensOptimally) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferTokensOptimally) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferTokensOptimally) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
