package types

import (
	"fmt"

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
	_, err = sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid toAddress address (%s)", err)
	}
	denomLen := len(msg.Denoms)
	fmt.Printf("XXXXX: %d\n", denomLen)
	if denomLen == 0 {
		return sdkerrors.Wrap(ErrParam, "no denominations")
	}

	if denomLen == 1 {
		if len(msg.Denoms[0]) == 0 {
			return sdkerrors.Wrap(ErrParam, "empty denomination")
		}
	}

	if denomLen > 1 {
		if len(msg.Denoms[0]) == 0 {
			return sdkerrors.Wrap(ErrParam, "empty denomination at position 0")
		}

		firstDenom := msg.Denoms[0]
		seenDenoms := make(map[string]bool)
		seenDenoms[firstDenom] = true

		for i, denom := range msg.Denoms[1:] {
			if len(denom) == 0 {
				return sdkerrors.Wrapf(ErrParam, "empty denomination at position %d", i+1)
			}

			if seenDenoms[denom] {
				return sdkerrors.Wrapf(ErrParam, "duplicate denomination %s", denom)
			}

			seenDenoms[denom] = true
		}

	}
	return nil
}
