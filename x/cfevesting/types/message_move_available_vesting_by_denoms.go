package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	_, _, err := ValidateMsgMoveAvailableVestingByDenom(msg.FromAddress, msg.ToAddress, msg.Denoms)
	return err
}

func ValidateMsgMoveAvailableVestingByDenom(fromAddress string, toAddress string, denoms []string) (fromAccAddress sdk.AccAddress, toAccAddress sdk.AccAddress, error error) {
	fromAccAddress, toAccAddress, err := ValidateAccountAddresses(fromAddress, toAddress)
	if err != nil {
		return nil, nil, errors.Wrap(err, "move available vesting by denoms")
	}

	if len(denoms) == 0 {
		return nil, nil, errors.Wrap(ErrParam, "move available vesting by denoms - no denominations")
	}
	seenDenoms := make(map[string]bool)
	for i, denom := range denoms {
		if len(denom) == 0 {
			return nil, nil, errors.Wrapf(ErrParam, "move available vesting by denoms - empty denomination at position %d", i)
		}

		if seenDenoms[denom] {
			return nil, nil, errors.Wrapf(ErrParam, "move available vesting by denoms - duplicate denomination %s", denom)
		}
		seenDenoms[denom] = true
	}

	return fromAccAddress, toAccAddress, nil
}
