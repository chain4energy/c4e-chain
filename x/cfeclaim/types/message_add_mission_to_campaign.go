package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgAddMissionToAidropCampaign = "add_mission_to_aidrop_campaign"

var _ sdk.Msg = &MsgAddMission{}

func NewMsgAddMission(owner string, campaignId uint64, name string, description string, missionType MissionType, weight *sdk.Dec, claimStartDate *time.Time) *MsgAddMission {
	return &MsgAddMission{
		Owner:          owner,
		Name:           name,
		Description:    description,
		CampaignId:     campaignId,
		MissionType:    missionType,
		Weight:         weight,
		ClaimStartDate: claimStartDate,
	}
}

func NewInitialMission(campaignId uint64) *Mission {
	return &Mission{
		CampaignId:  campaignId,
		Name:        "Initial mission",
		Description: "Initial mission - basic mission that must be claimed first",
		MissionType: MissionInitialClaim,
		Weight:      sdk.ZeroDec(),
	}
}

func (msg *MsgAddMission) Route() string {
	return RouterKey
}

func (msg *MsgAddMission) Type() string {
	return TypeMsgAddMissionToAidropCampaign
}

func (msg *MsgAddMission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddMission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMission) ValidateBasic() error {
	if msg.Weight == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "weight cannot be nil")
	}
	return ValidateAddMission(msg.Owner, msg.Name, msg.Description, msg.MissionType, *msg.Weight)
}

func ValidateAddMission(owner string, name string, description string, missionType MissionType, weight sdk.Dec) error {
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if name == "" {
		return errors.Wrap(c4eerrors.ErrParam, "empty name error")
	}
	if description == "" {
		return errors.Wrap(c4eerrors.ErrParam, "mission empty description error")
	}
	if err = ValidateMissionWeight(weight, missionType); err != nil {
		return err
	}
	return ValidateMissionType(missionType)
}

func ValidateMissionWeight(weight sdk.Dec, missionType MissionType) error {
	if weight.IsNil() {
		return errors.Wrapf(c4eerrors.ErrParam, "add mission to claim campaign weight is nil error")
	}
	if weight.GT(sdk.NewDec(1)) || weight.LT(sdk.ZeroDec()) {
		return errors.Wrapf(c4eerrors.ErrParam, "weight (%s) is not between 0 and 1 error", weight.String())
	}
	if missionType != MissionInitialClaim {
		if weight.Equal(sdk.ZeroDec()) {
			return errors.Wrapf(c4eerrors.ErrParam, "weight (%s) cannot be 0 for non-initial mission", weight.String())
		}
	}

	return nil
}

func ValidateMissionType(missionType MissionType) error {
	switch missionType {
	case MissionClaim, MissionDelegate, MissionVote, MissionInitialClaim:
		return nil
	}

	return errors.Wrap(sdkerrors.ErrInvalidType, "wrong mission type")
}
