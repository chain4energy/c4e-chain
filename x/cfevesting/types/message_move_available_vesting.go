package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgMoveAvailableVesting = "move_available_vesting"

var _ sdk.Msg = &MsgMoveAvailableVesting{}

func NewMsgMoveAvailableVesting(fromAddress string, toAddress string) *MsgMoveAvailableVesting {
	return &MsgMoveAvailableVesting{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
	}
}

func (msg *MsgMoveAvailableVesting) Route() string {
	return RouterKey
}

func (msg *MsgMoveAvailableVesting) Type() string {
	return TypeMsgMoveAvailableVesting
}

func (msg *MsgMoveAvailableVesting) GetSigners() []sdk.AccAddress {
	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddress}
}

func (msg *MsgMoveAvailableVesting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMoveAvailableVesting) ValidateBasic() error {
	_, _, err := ValidateMsgMoveAvailableVesting(msg.FromAddress, msg.ToAddress)
	if err != nil {
		return err
	}
	return nil
}

func ValidateMsgMoveAvailableVesting(fromAddress string, toAddress string) (fromAccAddress sdk.AccAddress, toAccAddress sdk.AccAddress, error error) {
	fromAccAddress, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrap(err, "move available vesting - from acc address error").Error())
	}
	toAccAddress, err = sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrap(err, "move available vesting - to acc address error").Error())
	}

	return fromAccAddress, toAccAddress, nil
}
