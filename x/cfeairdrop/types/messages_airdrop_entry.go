package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAirdropEntry = "create_airdrop_entry"
	TypeMsgUpdateAirdropEntry = "update_airdrop_entry"
	TypeMsgDeleteAirdropEntry = "delete_airdrop_entry"
)

var _ sdk.Msg = &MsgCreateAirdropEntry{}

func NewMsgCreateAirdropEntry(creator string, address string, amount uint64) *MsgCreateAirdropEntry {
	return &MsgCreateAirdropEntry{
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgCreateAirdropEntry) Route() string {
	return RouterKey
}

func (msg *MsgCreateAirdropEntry) Type() string {
	return TypeMsgCreateAirdropEntry
}

func (msg *MsgCreateAirdropEntry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAirdropEntry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAirdropEntry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAirdropEntry{}

func NewMsgUpdateAirdropEntry(creator string, id uint64, address string, amount uint64) *MsgUpdateAirdropEntry {
	return &MsgUpdateAirdropEntry{
		Id:      id,
		Creator: creator,
		Address: address,
		Amount:  amount,
	}
}

func (msg *MsgUpdateAirdropEntry) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAirdropEntry) Type() string {
	return TypeMsgUpdateAirdropEntry
}

func (msg *MsgUpdateAirdropEntry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAirdropEntry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAirdropEntry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAirdropEntry{}

func NewMsgDeleteAirdropEntry(creator string, id uint64) *MsgDeleteAirdropEntry {
	return &MsgDeleteAirdropEntry{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAirdropEntry) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAirdropEntry) Type() string {
	return TypeMsgDeleteAirdropEntry
}

func (msg *MsgDeleteAirdropEntry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAirdropEntry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAirdropEntry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
