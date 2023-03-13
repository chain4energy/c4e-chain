package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	_, _, err := ValidateMsgSplitVesting(msg.FromAddress, msg.ToAddress, msg.Amount)
	if err != nil {
		return err
	}
	return nil
}

func ValidateMsgSplitVesting(fromAddress string, toAddress string,
	amount sdk.Coins) (fromAccAddress sdk.AccAddress, toAccAddress sdk.AccAddress, error error) {
	if amount == nil {
		return nil, nil, errors.Wrapf(ErrParam, "split vesting - amount cannot be nil")
	}
	if amount.IsAnyNil() {
		return nil, nil, errors.Wrapf(ErrParam, "split vesting - amount cannot be nil")
	}
	err := amount.Validate()
	if err != nil {
		return nil, nil, errors.Wrapf(ErrParam, "split vesting - invalid amount (%s)", err)
	}
	fromAccAddress, err = sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrap(err, "split vesting - from acc address error").Error())
	}
	toAccAddress, err = sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return nil, nil, errors.Wrap(ErrParsing, errors.Wrap(err, "split vesting - to acc address error").Error())
	}

	return fromAccAddress, toAccAddress, nil
}
