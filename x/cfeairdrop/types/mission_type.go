package types

import "fmt"

const (
	MissionEmpty        = MissionType_UNSPECIFIED
	MissionInitialClaim = MissionType_INITIAL_CLAIM
	MissionDelegation   = MissionType_DELEGATION
	MissionVote         = MissionType_VOTE
	//VOTE
	end
)

func MissionTypeFromString(str string) (MissionType, error) {
	option, ok := MissionType_value[str]
	if !ok {
		return MissionEmpty, fmt.Errorf("'%s' is not a valid mission type, available options: initial_claim/vote/delegation", str)
	}
	return MissionType(option), nil
}

// NormalizeVoteOption - normalize user specified vote option
func NormalizeMissionType(option string) string {
	switch option {
	case "InitialClaim", "initial_claim", "INITIAL_CLAIM":
		return MissionInitialClaim.String()

	case "Delegation", "delegation", "DELEGATIOn":
		return MissionDelegation.String()

	case "Vote", "vote", "VOTE":
		return MissionVote.String()

	default:
		return option
	}
}
