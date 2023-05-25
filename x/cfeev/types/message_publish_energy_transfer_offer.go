package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPublishEnergyTransferOffer = "publish_energy_transfer_offer"

var _ sdk.Msg = &MsgPublishEnergyTransferOffer{}

func NewMsgPublishEnergyTransferOffer(creator string, chargerId string, tariff int32, location *Location, name string, plugType PlugType) *MsgPublishEnergyTransferOffer {
	return &MsgPublishEnergyTransferOffer{
		Creator:   creator,
		ChargerId: chargerId,
		Tariff:    tariff,
		Location:  location,
		Name:      name,
		PlugType:  plugType,
	}
}

func (msg *MsgPublishEnergyTransferOffer) Route() string {
	return RouterKey
}

func (msg *MsgPublishEnergyTransferOffer) Type() string {
	return TypeMsgPublishEnergyTransferOffer
}

func (msg *MsgPublishEnergyTransferOffer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPublishEnergyTransferOffer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPublishEnergyTransferOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
