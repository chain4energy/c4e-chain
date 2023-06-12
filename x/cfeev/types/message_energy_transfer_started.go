package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnergyTransferStarted = "energy_transfer_started"

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
	return TypeMsgEnergyTransferStarted
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
