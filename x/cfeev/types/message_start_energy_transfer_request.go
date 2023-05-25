package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartEnergyTransferRequest = "start_energy_transfer_request"

var _ sdk.Msg = &MsgStartEnergyTransferRequest{}

func NewMsgStartEnergyTransferRequest(creator string, energyTransferOfferId uint64, chargerId string, ownerAccountAddress string, offeredTariff string, collateral *sdk.Coin, energyToTransfer int32) *MsgStartEnergyTransferRequest {
	return &MsgStartEnergyTransferRequest{
		Creator:               creator,
		EnergyTransferOfferId: energyTransferOfferId,
		ChargerId:             chargerId,
		OwnerAccountAddress:   ownerAccountAddress,
		OfferedTariff:         offeredTariff,
		Collateral:            collateral,
		EnergyToTransfer:      energyToTransfer,
	}
}

func (msg *MsgStartEnergyTransferRequest) Route() string {
	return RouterKey
}

func (msg *MsgStartEnergyTransferRequest) Type() string {
	return TypeMsgStartEnergyTransferRequest
}

func (msg *MsgStartEnergyTransferRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartEnergyTransferRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartEnergyTransferRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
