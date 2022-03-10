package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBeginRedelegate = "begin_redelegate"

var _ sdk.Msg = &MsgBeginRedelegate{}

func NewMsgBeginRedelegate(delegatorAddress string, validatorSrcAddress string, validatorDstAddress string, amount sdk.Coin) *MsgBeginRedelegate {
	return &MsgBeginRedelegate{
		DelegatorAddress:    delegatorAddress,
		ValidatorSrcAddress: validatorSrcAddress,
		ValidatorDstAddress: validatorDstAddress,
		Amount:              amount,
	}
}

func (msg *MsgBeginRedelegate) Route() string {
	return RouterKey
}

func (msg *MsgBeginRedelegate) Type() string {
	return TypeMsgBeginRedelegate
}

func (msg *MsgBeginRedelegate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBeginRedelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBeginRedelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
