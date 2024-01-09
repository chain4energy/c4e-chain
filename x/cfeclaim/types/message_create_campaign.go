package types

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgCreateCampaign = "create_campaign"

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
	if msg.StartTime == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "start time cannot be nil")
	}
	if msg.EndTime == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "end time cannot be nil")
	}
	if msg.FeegrantAmount == nil {
		return errors.Wrap(c4eerrors.ErrParam, "feegrant amount cannot be nil")
	}
	if msg.InitialClaimFreeAmount == nil {
		return errors.Wrap(c4eerrors.ErrParam, "initital claim free amount cannot be nil")
	}
	if msg.Free == nil {
		return errors.Wrap(c4eerrors.ErrParam, "free decimal cannot be nil")
	}
	if msg.LockupPeriod == nil {
		return errors.Wrap(c4eerrors.ErrParam, "lockup period cannot be nil")
	}
	if msg.VestingPeriod == nil {
		return errors.Wrap(c4eerrors.ErrParam, "vesting period cannot be nil")
	}
	return ValidateCreateCampaignParams(msg.Name, *msg.FeegrantAmount,
		*msg.InitialClaimFreeAmount, *msg.Free, *msg.StartTime, *msg.EndTime, msg.CampaignType, *msg.LockupPeriod, *msg.VestingPeriod, msg.VestingPoolName)
}
