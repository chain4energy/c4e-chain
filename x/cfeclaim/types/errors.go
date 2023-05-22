package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/cfeclaim module sentinel errors
var (
	ErrCampaignDisabled    = errors.Register(ModuleName, 2, "campaign is disabled")
	ErrMissionCompleted    = errors.Register(ModuleName, 3, "mission already completed")
	ErrMissionClaimed      = errors.Register(ModuleName, 4, "mission already claimed")
	ErrMissionNotCompleted = errors.Register(ModuleName, 5, "mission not completed yet")
	ErrMissionCompletion   = errors.Register(ModuleName, 6, "mission completion error")
	ErrMissionClaiming     = errors.Register(ModuleName, 7, "mission claiming error")
	ErrMissionDisabled     = errors.Register(ModuleName, 8, "mission is disabled")
	ErrCampaignEnabled     = errors.Register(ModuleName, 9, "campaign is enabled")
)
