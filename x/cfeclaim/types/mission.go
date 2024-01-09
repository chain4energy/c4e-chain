package types

import (
	"cosmossdk.io/errors"
	"fmt"
	"time"
)

const (
	MissionEmpty        = MissionType_MISSION_TYPE_UNSPECIFIED
	MissionInitialClaim = MissionType_INITIAL_CLAIM
	MissionDelegate     = MissionType_DELEGATE
	MissionVote         = MissionType_VOTE
	MissionClaim        = MissionType_CLAIM
	MissionToDefine     = MissionType_TO_DEFINE
)

func MissionTypeFromString(str string) (MissionType, error) {
	option, ok := MissionType_value[str]
	if !ok {
		return MissionEmpty, fmt.Errorf("'%s' is not a valid mission type, available options: claim/vote/delegate/to_define", str)
	}
	return MissionType(option), nil
}

func NormalizeMissionType(option string) string {
	switch option {
	case "Delegate", "delegate", "DELEGATE":
		return MissionDelegate.String()

	case "Vote", "vote", "VOTE":
		return MissionVote.String()

	case "Claim", "claim", "CLAIM":
		return MissionClaim.String()

	case "ToDefine", "to_define", "TO_DEFINE":
		return MissionToDefine.String()

	default:
		return option
	}
}

func (c *Mission) IsEnabled(blockTime time.Time) error {
	if c.ClaimStartDate == nil {
		return nil
	}
	if c.ClaimStartDate.After(blockTime) {
		return errors.Wrapf(ErrMissionDisabled, "mission %d not started yet (blocktime %s < mission start time %s)", c.Id, blockTime, c.ClaimStartDate)
	}
	return nil
}
