package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
	"time"
)

const TypeMsgCreateCampaign = "create_claim_campaign"

var _ sdk.Msg = &MsgCreateCampaign{}

func NewMsgCreateCampaign(owner string, name string, description string, campaignType CampaignType, argFeegrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration) *MsgCreateCampaign {
	return &MsgCreateCampaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		CampaignType:           campaignType,
		FeegrantAmount:         argFeegrantAmount,
		InitialClaimFreeAmount: initialClaimFreeAmount,
		StartTime:              startTime,
		EndTime:                endTime,
		LockupPeriod:           lockupPeriod,
		VestingPeriod:          vestingPeriod,
	}
}

func (msg *MsgCreateCampaign) Route() string {
	return RouterKey
}

func (msg *MsgCreateCampaign) Type() string {
	return TypeMsgCreateCampaign
}

func (msg *MsgCreateCampaign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCampaign) ValidateBasic() error {
	return ValidateCampaignCreateParams(msg.Name, msg.Description, msg.StartTime, msg.EndTime, msg.CampaignType, msg.Owner)
}

func ValidateCampaignCreateParams(name string, description string, startTime *time.Time, endTime *time.Time, campaignType CampaignType, owner string) error {
	if err := ValidateCampaignName(name); err != nil {
		return err
	}
	if err := ValidateCampaignDescription(description); err != nil {
		return err
	}
	if err := ValidateCampaignEndTimeAfterStartTime(startTime, endTime); err != nil {
		return err
	}
	if err := ValidateCampaignType(campaignType, owner); err != nil {
		return err
	}
	return ValidateCampaignTypeTeamdrop(campaignType, owner)
}

func ValidateCampaignName(name string) error {
	if name == "" {
		return errors.Wrap(c4eerrors.ErrParam, "campaign name is empty")
	}
	return nil
}

func ValidateCampaignDescription(description string) error {
	if description == "" {
		return errors.Wrap(c4eerrors.ErrParam, "description is empty")
	}
	return nil
}

func ValidateCampaignEndTimeAfterStartTime(startTime *time.Time, endTime *time.Time) error {
	if endTime == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "end time is nil error")
	}
	if startTime.After(*endTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "start time is after end time error (%s > %s)", startTime, endTime)
	}

	return nil
}

func ValidateCampaignType(campaignType CampaignType, owner string) error {
	switch campaignType {
	case CampaignDefault, CampaignSale, CampaignTeamdrop:
		return nil
	}

	return errors.Wrap(sdkerrors.ErrInvalidType, "wrong campaign close action type")
}

func ValidateCampaignTypeTeamdrop(campaignType CampaignType, owner string) error {
	if campaignType == CampaignTeamdrop {
		if !slices.Contains(GetWhitelistedTeamdropAccounts(), owner) {
			return errors.Wrap(sdkerrors.ErrorInvalidSigner, "TeamDrop campaigns can be created only by specific accounts")
		}
	}

	return nil
}
