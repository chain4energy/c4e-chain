package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartEnergyTransfer = "start_energy_transfer"

var _ sdk.Msg = &MsgStartEnergyTransfer{}

func NewMsgStartEnergyTransfer(creator string, energyTransferOfferId uint64, offeredTariff uint64,
	energyToTransfer uint64) *MsgStartEnergyTransfer {
	return &MsgStartEnergyTransfer{
		Creator:               creator,
		EnergyTransferOfferId: energyTransferOfferId,
		OfferedTariff:         offeredTariff,
		EnergyToTransfer:      energyToTransfer,
	}
}

func (msg *MsgStartEnergyTransfer) Route() string {
	return RouterKey
}

func (msg *MsgStartEnergyTransfer) Type() string {
	return TypeMsgStartEnergyTransfer
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
	return ValidateStartEnergyTransfer(msg.Creator, msg.OfferedTariff, msg.EnergyToTransfer)
}

func ValidateStartEnergyTransfer(creator string, offeredTarif uint64, energyToTransfer uint64) error {
	_, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if offeredTarif == 0 {
		return errors.Wrapf(c4eerrors.ErrParam, "offered tariff cannot be empty")
	}
	if energyToTransfer == 0 {
		return errors.Wrapf(c4eerrors.ErrParam, "cannot transfer zero [kWh] energy")
	}
	return nil
}
