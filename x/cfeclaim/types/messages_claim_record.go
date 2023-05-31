package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
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

func ValidateClaimRecordEntries(claimRecords []*ClaimRecordEntry) error {
	for i, claimRecord := range claimRecords {
		if err := ValidateClaimRecordEntry(claimRecord); err != nil {
			return errors.Wrapf(err, "claim record entry index %d", i)
		}
	}
	return nil
}

func ValidateClaimRecordEntry(claimRecordEntry *ClaimRecordEntry) error {
	if claimRecordEntry.UserEntryAddress == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record entry empty user entry address")
	}
	if !claimRecordEntry.Amount.IsAllPositive() {
		return errors.Wrapf(c4eerrors.ErrParam, "claim record entry must has at least one coin and all amounts must be positive")
	}
	return nil
}

func ValidateUserEntry(userEntry UserEntry) error {
	if userEntry.Address == "" {
		return errors.Wrapf(c4eerrors.ErrParam, "user entry empty address")
	}
	if err := ValidateUserEntryClaimRecords(userEntry.ClaimRecords); err != nil {
		return err
	}
	return nil
}

func ValidateUserEntryClaimRecords(claimRecords []*ClaimRecord) error {
	for i, claimRecord := range claimRecords {
		if !claimRecord.Amount.IsAllPositive() {
			return errors.Wrapf(c4eerrors.ErrParam, "claim record at index %d must has at least one coin and all amounts must be positive", i)
		}
	}
	return nil
}
