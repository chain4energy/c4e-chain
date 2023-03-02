package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = Params{SubDistributors: msg.SubDistributors}.Validate()
	if err != nil {
		errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}
	return nil
}

const TypeMsgUpdateSubDistributor = "update_single_subdistributor"

var _ sdk.Msg = &MsgUpdateSubDistributor{}

func (msg *MsgUpdateSubDistributor) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSubDistributor) Type() string {
	return TypeMsgUpdateSubDistributor
}

func (msg *MsgUpdateSubDistributor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSubDistributor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSubDistributor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = msg.SubDistributor.Validate()
	if err != nil {
		errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}
	return nil
}

const TypeMsgUpdateSubDistributorBurnShare = "update_sub_distributor_burn_share"

var _ sdk.Msg = &MsgUpdateSubDistributorBurnShare{}

func (msg *MsgUpdateSubDistributorBurnShare) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSubDistributorBurnShare) Type() string {
	return TypeMsgUpdateSubDistributorBurnShare
}

func (msg *MsgUpdateSubDistributorBurnShare) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSubDistributorBurnShare) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSubDistributorBurnShare) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.SubDistributorName == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "empty destination name")
	}
	if msg.BurnShare.IsNil() {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "burn share cannot be nil")
	}
	if msg.BurnShare.GTE(sdk.NewDec(maxShare)) || msg.BurnShare.IsNegative() {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "burn share must be between 0 and 1")
	}
	return nil
}

const TypeMsgUpdateSubDistributorDestinationShare = "update_sub_distributor_destination_share"

var _ sdk.Msg = &MsgUpdateSubDistributorBurnShare{}

func (msg *MsgUpdateSubDistributorDestinationShare) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSubDistributorDestinationShare) Type() string {
	return TypeMsgUpdateSubDistributorDestinationShare
}

func (msg *MsgUpdateSubDistributorDestinationShare) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSubDistributorDestinationShare) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSubDistributorDestinationShare) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: (%s)", err)
	}
	if msg.DestinationName == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "empty destination name")
	}
	if msg.DestinationName == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "empty destination name")
	}
	if msg.Share.IsNil() {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "share cannot be nil")
	}
	if msg.Share.GTE(sdk.NewDec(maxShare)) || msg.Share.IsNegative() {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "share must be between 0 and 1")
	}
	return nil
}
