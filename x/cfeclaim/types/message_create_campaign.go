package types

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgCreateCampaign = "create_claim_campaign"

var _ sdk.Msg = &MsgCreateCampaign{}

func NewMsgCreateCampaign(owner string, name string, description string, campaignType CampaignType, removableClaimRecords bool, argFeegrantAmount *math.Int,
	initialClaimFreeAmount *math.Int, free *sdk.Dec, startTime *time.Time, endTime *time.Time, lockupPeriod *time.Duration,
	vestingPeriod *time.Duration, vestingPoolName string) *MsgCreateCampaign {
	return &MsgCreateCampaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		CampaignType:           campaignType,
		RemovableClaimRecords:  removableClaimRecords,
		FeegrantAmount:         argFeegrantAmount,
		InitialClaimFreeAmount: initialClaimFreeAmount,
		Free:                   free,
		StartTime:              startTime,
		EndTime:                endTime,
		LockupPeriod:           lockupPeriod,
		VestingPeriod:          vestingPeriod,
		VestingPoolName:        vestingPoolName,
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
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return ValidateCreateCampaignParams(msg.Name, msg.Description, msg.StartTime, msg.EndTime, msg.CampaignType, msg.VestingPoolName)
}

func ValidateCreateCampaignParams(name string, description string, startTime *time.Time, endTime *time.Time, campaignType CampaignType, vestingPoolName string) error {
	if err := ValidateCampaignName(name); err != nil { // TODO bardzo duze rodrobnienie - jest to potrzebne? 
		return err
	}
	if err := ValidateCampaignDescription(description); err != nil { // TODO bardzo duze rodrobnienie - jest to potrzebne? 
		return err
	}
	if err := ValidateCampaignEndTimeAfterStartTime(startTime, endTime); err != nil {
		return err
	}
	return ValidateCampaignType(campaignType, vestingPoolName)
	// TODO dodanie validacji periodow i free gdy kampabni typy vestig pool
}

func ValidateCampaignName(name string) error { // TODO raczej za duze rozdrobnienie - jest jakis powod?
	if name == "" {
		return errors.Wrap(c4eerrors.ErrParam, "campaign name is empty")
	}
	return nil
}

func ValidateCampaignDescription(description string) error { // TODO raczej za duze rozdrobnienie - jest jakis powod? czy to wlasciwe jest? description moze byc puste
	if description == "" {
		return errors.Wrap(c4eerrors.ErrParam, "description is empty")
	}
	return nil
}

func ValidateCampaignEndTimeAfterStartTime(startTime *time.Time, endTime *time.Time) error { // TODO nazwa ValidateCampaignTimes
	if endTime == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "end time is nil error")
	}
	if startTime.After(*endTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "start time is after end time (%s > %s)", startTime, endTime)
	}
	if startTime.Equal(*endTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "start time is equal to end time (%s = %s)", startTime, endTime)
	}
	return nil
}

func ValidateCampaignType(campaignType CampaignType, vestingPoolName string) error { // TODO raczej prywatna funkcja
	switch campaignType {
	case VestingPoolCampaign:
		if vestingPoolName == "" {
			return errors.Wrap(c4eerrors.ErrParam, "for VESTING_POOL type campaigns, the vesting pool name must be provided")
		}
		return nil
	case DefaultCampaign:
		if vestingPoolName != "" {
			return errors.Wrap(c4eerrors.ErrParam, "vesting pool name can be set only for VESTING_POOL type campaigns")
		}
		return nil
	}

	return errors.Wrap(sdkerrors.ErrInvalidType, "wrong campaign type")
}
