package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPublishEnergyTransferOffer = "publish_energy_transfer_offer"

var _ sdk.Msg = &MsgPublishEnergyTransferOffer{}

func NewMsgPublishEnergyTransferOffer(creator string, chargerId string, tariff uint64, location *Location, name string, plugType PlugType) *MsgPublishEnergyTransferOffer {
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.GetLocation() == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "charger location cannot be nil")
	}
	return ValidatePublishEnergyTransferOffer(msg.GetChargerId(), msg.GetName(), *msg.GetLocation())
}

func ValidatePublishEnergyTransferOffer(chargerId string, name string, location Location) error {
	if chargerId == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "charger id cannot be empty")
	}
	if name == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "charger name cannot be empty")
	}
	return location.Validate()
}
