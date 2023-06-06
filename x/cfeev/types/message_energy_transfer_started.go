package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnergyTransferStartedRequest = "energy_transfer_started"

var _ sdk.Msg = &MsgEnergyTransferStarted{}

func NewMsgEnergyTransferStarted(creator string, energyTransferId uint64, chargerId string, info string) *MsgEnergyTransferStarted {
	return &MsgEnergyTransferStarted{
		Creator:          creator,
		EnergyTransferId: energyTransferId,
		ChargerId:        chargerId,
		Info:             info,
	}
}

func (msg *MsgEnergyTransferStarted) Route() string {
	return RouterKey
}

func (msg *MsgEnergyTransferStarted) Type() string {
	return TypeMsgEnergyTransferStartedRequest
}

func (msg *MsgEnergyTransferStarted) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnergyTransferStarted) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnergyTransferStarted) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
