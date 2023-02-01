package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAirdropEntry = "create_airdrop_entry"
	TypeMsgUpdateAirdropEntry = "update_airdrop_entry"
	TypeMsgDeleteClaimRecord  = "delete_airdrop_entry"
)

var _ sdk.Msg = &MsgAddClaimRecords{}

func NewMsgCreateAirdropEntry(onwer string, campaignId uint64, airdropEntries []*ClaimRecord) *MsgAddClaimRecords {
	return &MsgAddClaimRecords{
		Owner:        onwer,
		CampaignId:   campaignId,
		ClaimRecords: airdropEntries,
	}
}

func (msg *MsgAddClaimRecords) Route() string {
	return RouterKey
}

func (msg *MsgAddClaimRecords) Type() string {
	return TypeMsgCreateAirdropEntry
}

func (msg *MsgAddClaimRecords) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddClaimRecords) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddClaimRecords) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteClaimRecord{}

func NewMsgDeleteClaimRecord(onwer string, campaignId uint64, userAddress string) *MsgDeleteClaimRecord {
	return &MsgDeleteClaimRecord{
		CampaignId:  campaignId,
		Owner:       onwer,
		UserAddress: userAddress,
	}
}
func (msg *MsgDeleteClaimRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClaimRecord) Type() string {
	return TypeMsgDeleteClaimRecord
}

func (msg *MsgDeleteClaimRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClaimRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClaimRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
