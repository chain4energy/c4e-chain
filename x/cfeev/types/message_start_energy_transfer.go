package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartEnergyTransfer = "start_energy_transfer"

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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.GetChargerId() == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "Charger ID is empty")
	}
	if msg.GetOwnerAccountAddress() == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "CP owner account address cannot be empty")
	}
	if msg.GetOfferedTariff() == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "Offered tariff cannot be empty")
	}
	if msg.GetEnergyToTransfer() == 0 {
		return errors.Wrapf(c4eerrors.ErrParam, "Cannot transfer zero [kWh] energy")
	}

	return nil
}
