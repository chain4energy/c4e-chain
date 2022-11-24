package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cfeairdrop module sentinel errors
var (
	ErrCampaignDisabled = sdkerrors.Register(ModuleName, 2, "camapaign is disabled")
	ErrMissionCompleted = sdkerrors.Register(ModuleName, 3, "mission already completed")
	ErrMissionClaimed   = sdkerrors.Register(ModuleName, 4, "mission already claimed")
	ErrMissionNotCompleted = sdkerrors.Register(ModuleName, 5, "mission not completed yet")
)
