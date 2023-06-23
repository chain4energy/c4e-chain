package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveEnergyOffer = "remove_energy_offer"

var _ sdk.Msg = &MsgRemoveEnergyOffer{}

func NewMsgRemoveEnergyOffer(owner string, id uint64) *MsgRemoveEnergyOffer {
	return &MsgRemoveEnergyOffer{
		Owner: owner,
		Id:    id,
	}
}

func (msg *MsgRemoveEnergyOffer) Route() string {
	return RouterKey
}

func (msg *MsgRemoveEnergyOffer) Type() string {
	return TypeMsgRemoveEnergyOffer
}

func (msg *MsgRemoveEnergyOffer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
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
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
