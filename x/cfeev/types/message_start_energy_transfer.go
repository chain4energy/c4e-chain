package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartEnergyTransferRequest = "start_energy_transfer"

var _ sdk.Msg = &MsgStartEnergyTransfer{}

func NewMsgStartEnergyTransfer(creator string, energyTransferOfferId uint64, chargerId string, ownerAccountAddress string, offeredTariff string, collateral *sdk.Coin, energyToTransfer int32) *MsgStartEnergyTransfer {
	return &MsgStartEnergyTransfer{
		Creator:               creator,
		EnergyTransferOfferId: energyTransferOfferId,
		ChargerId:             chargerId,
		OwnerAccountAddress:   ownerAccountAddress,
		OfferedTariff:         offeredTariff,
		Collateral:            collateral,
		EnergyToTransfer:      energyToTransfer,
	}
}

func (msg *MsgStartEnergyTransfer) Route() string {
	return RouterKey
}

func (msg *MsgStartEnergyTransfer) Type() string {
	return TypeMsgStartEnergyTransferRequest
}

func (msg *MsgStartEnergyTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartEnergyTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartEnergyTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.GetChargerId() == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "Charger ID is empty")
	}
	if msg.GetOwnerAccountAddress() == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "CP owner account address cannot be empty")
	}
	if msg.GetOfferedTariff() == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "Offered tariff cannot be empty")
	}
	if msg.GetEnergyToTransfer() == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "Cannot transfer zero [kWh] energy")
	}

	return nil
}
