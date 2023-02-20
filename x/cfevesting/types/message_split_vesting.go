package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSplitVesting = "split_vesting"

var _ sdk.Msg = &MsgSplitVesting{}

func NewMsgSplitVesting(fromAddress string, toAddress string, amount sdk.Coins) *MsgSplitVesting {
	return &MsgSplitVesting{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      amount,
	}
}

func (msg *MsgSplitVesting) Route() string {
	return RouterKey
}

func (msg *MsgSplitVesting) Type() string {
	return TypeMsgSplitVesting
}

func (msg *MsgSplitVesting) GetSigners() []sdk.AccAddress {
	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddress}
}

func (msg *MsgSplitVesting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSplitVesting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid fromAddress address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid toAddress address (%s)", err)
	}
	err = msg.Amount.Validate()
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid amount (%s)", err)
	}
	return nil
}
