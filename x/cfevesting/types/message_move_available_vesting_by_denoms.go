package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMoveAvailableVestingByDenoms = "move_available_vesting_by_denoms"

var _ sdk.Msg = &MsgMoveAvailableVestingByDenoms{}

func NewMsgMoveAvailableVestingByDenoms(fromAddress string, toAddress string, denoms []string) *MsgMoveAvailableVestingByDenoms {
	return &MsgMoveAvailableVestingByDenoms{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Denoms:      denoms,
	}
}

func (msg *MsgMoveAvailableVestingByDenoms) Route() string {
	return RouterKey
}

func (msg *MsgMoveAvailableVestingByDenoms) Type() string {
	return TypeMsgMoveAvailableVestingByDenoms
}

func (msg *MsgMoveAvailableVestingByDenoms) GetSigners() []sdk.AccAddress {
	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddress}
}

func (msg *MsgMoveAvailableVestingByDenoms) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMoveAvailableVestingByDenoms) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid fromAddress address (%s)", err)
	}
	// TODO check if duplications
	return nil
}
