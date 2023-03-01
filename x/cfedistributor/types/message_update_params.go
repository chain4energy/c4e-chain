package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateAllSubDistributors = "update_all_subdistributors"

var _ sdk.Msg = &MsgUpdateAllSubDistributors{}

func (msg *MsgUpdateAllSubDistributors) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAllSubDistributors) Type() string {
	return TypeMsgUpdateAllSubDistributors
}

func (msg *MsgUpdateAllSubDistributors) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAllSubDistributors) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAllSubDistributors) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgUpdateSingleSubdistributor = "update_single_subdistributor"

var _ sdk.Msg = &MsgUpdateSingleSubdistributor{}

func (msg *MsgUpdateSingleSubdistributor) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSingleSubdistributor) Type() string {
	return TypeMsgUpdateSingleSubdistributor
}

func (msg *MsgUpdateSingleSubdistributor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSingleSubdistributor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSingleSubdistributor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
