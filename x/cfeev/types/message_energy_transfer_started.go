package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnergyTransferStartedRequest = "energy_transfer_started_request"

var _ sdk.Msg = &MsgEnergyTransferStartedRequest{}

func NewMsgEnergyTransferStartedRequest(creator string, energyTransferId uint64, chargerId string, info string) *MsgEnergyTransferStartedRequest {
	return &MsgEnergyTransferStartedRequest{
		Creator:          creator,
		EnergyTransferId: energyTransferId,
		ChargerId:        chargerId,
		Info:             info,
	}
}

func (msg *MsgEnergyTransferStartedRequest) Route() string {
	return RouterKey
}

func (msg *MsgEnergyTransferStartedRequest) Type() string {
	return TypeMsgEnergyTransferStartedRequest
}

func (msg *MsgEnergyTransferStartedRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnergyTransferStartedRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnergyTransferStartedRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
