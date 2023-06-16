package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnergyTransferCompleted = "energy_transfer_completed"

var _ sdk.Msg = &MsgEnergyTransferCompleted{}

func NewMsgEnergyTransferCompleted(creator string, energyTransferId uint64, chargerId string, usedServiceUnits uint64, info string) *MsgEnergyTransferCompleted {
	return &MsgEnergyTransferCompleted{
		Creator:          creator,
		EnergyTransferId: energyTransferId,
		ChargerId:        chargerId,
		UsedServiceUnits: usedServiceUnits,
		Info:             info,
	}
}

func (msg *MsgEnergyTransferCompleted) Route() string {
	return RouterKey
}

func (msg *MsgEnergyTransferCompleted) Type() string {
	return TypeMsgEnergyTransferCompleted
}

func (msg *MsgEnergyTransferCompleted) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnergyTransferCompleted) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnergyTransferCompleted) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.ChargerId == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "charger id cannot be empty")
	}
	return nil
}
