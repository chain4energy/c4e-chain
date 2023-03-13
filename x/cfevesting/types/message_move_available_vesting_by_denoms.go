package types

import (
	"cosmossdk.io/errors"
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
	_, _, err := ValidateMsgMoveAvailableVestingBeDenom(msg.FromAddress, msg.ToAddress)
	if err != nil {
		return err
	}
	denomLen := len(msg.Denoms)
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

func ValidateMsgMoveAvailableVestingBeDenom(fromAddress string, toAddress string) (fromAccAddress sdk.AccAddress, toAccAddress sdk.AccAddress, error error) {
	fromAccAddress, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrap(err, "move available vesting by denoms - from acc address error").Error())
	}
	toAccAddress, err = sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrap(err, "move available vesting by denoms - to acc address error").Error())
	}

	return fromAccAddress, toAccAddress, nil
}
