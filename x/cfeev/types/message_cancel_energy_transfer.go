package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelEnergyTransfer = "cancel_energy_transfer"

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
	return TypeMsgCancelEnergyTransfer
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.ChargerId == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "charger id cannot be empty")
	}
	return nil
}
