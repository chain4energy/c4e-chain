package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnergyTransferCompletedRequest = "energy_transfer_completed_request"

var _ sdk.Msg = &MsgEnergyTransferCompletedRequest{}

func NewMsgEnergyTransferCompletedRequest(creator string, energyTransferId uint64, chargerId string, usedServiceUnits int32, info string) *MsgEnergyTransferCompletedRequest {
	return &MsgEnergyTransferCompletedRequest{
		Creator:          creator,
		EnergyTransferId: energyTransferId,
		ChargerId:        chargerId,
		UsedServiceUnits: usedServiceUnits,
		Info:             info,
	}
}

func (msg *MsgEnergyTransferCompletedRequest) Route() string {
	return RouterKey
}

func (msg *MsgEnergyTransferCompletedRequest) Type() string {
	return TypeMsgEnergyTransferCompletedRequest
}

func (msg *MsgEnergyTransferCompletedRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnergyTransferCompletedRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnergyTransferCompletedRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
