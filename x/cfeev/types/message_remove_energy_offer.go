package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveEnergyOffer = "remove_energy_offer"

var _ sdk.Msg = &MsgRemoveEnergyOffer{}

func NewMsgRemoveEnergyOffer(creator string, id uint64) *MsgRemoveEnergyOffer {
	return &MsgRemoveEnergyOffer{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRemoveEnergyOffer) Route() string {
	return RouterKey
}

func (msg *MsgRemoveEnergyOffer) Type() string {
	return TypeMsgRemoveEnergyOffer
}

func (msg *MsgRemoveEnergyOffer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveEnergyOffer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveEnergyOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
