package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAddClaimRecords   = "add_claim_records"
	TypeMsgDeleteClaimRecord = "delete_claim_record"
)

var _ sdk.Msg = &MsgAddClaimRecords{}

func NewMsgAddClaimRecords(onwer string, campaignId uint64, claimRecordEntries []*ClaimRecordEntry) *MsgAddClaimRecords {
	return &MsgAddClaimRecords{
		Owner:              onwer,
		CampaignId:         campaignId,
		ClaimRecordEntries: claimRecordEntries,
	}
}

func (msg *MsgAddClaimRecords) Route() string {
	return RouterKey
}

func (msg *MsgAddClaimRecords) Type() string {
	return TypeMsgAddClaimRecords
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return ValidateClaimRecordEntries(msg.ClaimRecordEntries)
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
