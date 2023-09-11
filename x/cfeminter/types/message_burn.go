package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/v2/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurn = "burn"

var _ sdk.Msg = &MsgBurn{}

func NewMsgBurn(address string, amount sdk.Coins) *MsgBurn {
	return &MsgBurn{
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgBurn) Route() string {
	return RouterKey
}

func (msg *MsgBurn) Type() string {
	return TypeMsgBurn
}

func (msg *MsgBurn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	if msg.Amount == nil {
		return errors.Wrap(c4eerrors.ErrParam, "amount is nil")
	}
	return ValidateMsgBurn(msg.Amount)
}

func ValidateMsgBurn(amount sdk.Coins) error {
	if amount.IsAnyNil() {
		return errors.Wrap(c4eerrors.ErrParam, "amount is nil")
	}
	if !amount.IsAllPositive() {
		return errors.Wrap(c4eerrors.ErrParam, "amount is not positive")
	}
	return nil
}
