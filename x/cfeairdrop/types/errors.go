package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfeairdrop module sentinel errors
var (
	ErrCampaignDisabled    = sdkerrors.Register(ModuleName, 2, "campaign is disabled")
	ErrMissionCompleted    = sdkerrors.Register(ModuleName, 3, "mission already completed")
	ErrMissionClaimed      = sdkerrors.Register(ModuleName, 4, "mission already claimed")
	ErrMissionNotCompleted = sdkerrors.Register(ModuleName, 5, "mission not completed yet")
	ErrMissionCompletion   = sdkerrors.Register(ModuleName, 6, "mission completion error")
	ErrMissionClaiming     = sdkerrors.Register(ModuleName, 7, "mission claiming error")
	ErrMissionDisabled     = sdkerrors.Register(ModuleName, 8, "mission is disabled")
)
