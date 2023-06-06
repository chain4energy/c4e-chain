package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelEnergyTransferRequest = "cancel_energy_transfer"

var _ sdk.Msg = &MsgCancelEnergyTransfer{}

func NewMsgCancelEnergyTransfer(creator string, energyTransferId uint64, chargerId string, errorInfo string, errorCode string) *MsgCancelEnergyTransfer {
	return &MsgCancelEnergyTransfer{
		Creator:          creator,
		EnergyTransferId: energyTransferId,
		ChargerId:        chargerId,
		ErrorInfo:        errorInfo,
		ErrorCode:        errorCode,
	}
}

func (msg *MsgCancelEnergyTransfer) Route() string {
	return RouterKey
}

func (msg *MsgCancelEnergyTransfer) Type() string {
	return TypeMsgCancelEnergyTransferRequest
}

func (msg *MsgCancelEnergyTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelEnergyTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelEnergyTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
