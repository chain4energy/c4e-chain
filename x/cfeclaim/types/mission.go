package types

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const (
	MissionEmpty        = MissionType_MISSION_TYPE_UNSPECIFIED
	MissionInitialClaim = MissionType_INITIAL_CLAIM
	MissionDelegation   = MissionType_DELEGATION
	MissionVote         = MissionType_VOTE
	MissionClaim        = MissionType_CLAIM
)

func MissionTypeFromString(str string) (MissionType, error) {
	option, ok := MissionType_value[str]
	if !ok {
		return MissionEmpty, fmt.Errorf("'%s' is not a valid mission type, available options: initial_claim/vote/delegation", str)
	}
	return MissionType(option), nil
}

func NormalizeMissionType(option string) string {
	switch option {
	case "InitialClaim", "initial_claim", "INITIAL_CLAIM":
		return MissionInitialClaim.String()

	case "Delegation", "delegation", "DELEGATION":
		return MissionDelegation.String()

	case "Vote", "vote", "VOTE":
		return MissionVote.String()

	default:
		return option
	}
}

func (c *Mission) IsEnabled(blockTime time.Time) error {
	if c.ClaimStartDate == nil {
		return nil
	}
	if c.ClaimStartDate.Before(blockTime) {
		return sdkerrors.Wrapf(ErrMissionDisabled, "mission %d not started yet (%s < startTime %s) error", c.Id, blockTime, c.ClaimStartDate)
	}
	return nil
}
