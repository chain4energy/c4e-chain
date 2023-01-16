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

var _ sdk.Msg = &MsgAddAirdropEntries{}

func NewMsgCreateAirdropEntry(onwer string, campaignId uint64, airdropEntries []*AirdropEntry) *MsgAddAirdropEntries {
	return &MsgAddAirdropEntries{
		Owner:          onwer,
		CampaignId:     campaignId,
		AirdropEntries: airdropEntries,
	}
}

func (msg *MsgAddAirdropEntries) Route() string {
	return RouterKey
}

func (msg *MsgAddAirdropEntries) Type() string {
	return TypeMsgCreateAirdropEntry
}

func (msg *MsgAddAirdropEntries) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAirdropEntries) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAirdropEntries) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
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
