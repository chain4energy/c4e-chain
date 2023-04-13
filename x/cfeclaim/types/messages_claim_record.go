package types

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAddClaimRecords   = "add_claim_records"
	TypeMsgDeleteClaimRecord = "delete_claim_record"
)

var _ sdk.Msg = &MsgAddClaimRecords{}

func NewMsgAddClaimRecords(onwer string, campaignId uint64, claimEntries []*ClaimRecord) *MsgAddClaimRecords {
	return &MsgAddClaimRecords{
		Owner:        onwer,
		CampaignId:   campaignId,
		ClaimRecords: claimEntries,
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return ValidateClaimRecords(msg.ClaimRecords)
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
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func ValidateClaimRecords(claimRecords []*ClaimRecord) error {
	for i, claimRecord := range claimRecords {
		if err := ValidateClaimRecord(claimRecord); err != nil {
			return WrapClaimRecordIndex(err, i)
		}
	}
	return nil
}

func ValidateClaimRecord(claimRecord *ClaimRecord) error {
	if claimRecord.Address == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record empty address")
	}
	if !claimRecord.Amount.IsAllPositive() {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record must has at least one coin and all amounts must be positive")
	}
	return nil
}

func WrapClaimRecordIndex(err error, index int) error {
	return errors.Wrap(err, fmt.Sprintf("claim records index %d", index))
}
