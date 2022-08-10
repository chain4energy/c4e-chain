package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateTokenParams = "create_token_params"

var _ sdk.Msg = &MsgCreateTokenParams{}

func NewMsgCreateTokenParams(creator string, name string, tradingCompany string, burningTime uint64, burningType string, sendPrice uint64) *MsgCreateTokenParams {
	return &MsgCreateTokenParams{
		Creator:        creator,
		Name:           name,
		TradingCompany: tradingCompany,
		BurningTime:    burningTime,
		BurningType:    burningType,
		SendPrice:      sendPrice,
	}
}

func (msg *MsgCreateTokenParams) Route() string {
	return RouterKey
}

func (msg *MsgCreateTokenParams) Type() string {
	return TypeMsgCreateTokenParams
}

func (msg *MsgCreateTokenParams) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTokenParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTokenParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
