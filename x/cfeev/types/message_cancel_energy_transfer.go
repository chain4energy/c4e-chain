package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelEnergyTransferRequest = "cancel_energy_transfer_request"

var _ sdk.Msg = &MsgCancelEnergyTransferRequest{}

func NewMsgCancelEnergyTransferRequest(creator string, energyTransferId uint64, chargerId string, errorInfo string, errorCode string) *MsgCancelEnergyTransferRequest {
	return &MsgCancelEnergyTransferRequest{
		Creator:          creator,
		EnergyTransferId: energyTransferId,
		ChargerId:        chargerId,
		ErrorInfo:        errorInfo,
		ErrorCode:        errorCode,
	}
}

func (msg *MsgCancelEnergyTransferRequest) Route() string {
	return RouterKey
}

func (msg *MsgCancelEnergyTransferRequest) Type() string {
	return TypeMsgCancelEnergyTransferRequest
}

func (msg *MsgCancelEnergyTransferRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelEnergyTransferRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelEnergyTransferRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
