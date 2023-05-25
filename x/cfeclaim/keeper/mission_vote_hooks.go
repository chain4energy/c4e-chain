package keeper

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type MissionVoteHooks struct {
	k Keeper
}

// NewMissionVoteHooks returns a GovHooks that triggers mission completion on voting for a proposal
func (k Keeper) NewMissionVoteHooks() MissionVoteHooks {
	return MissionVoteHooks{k}
}

var _ govtypes.GovHooks = MissionVoteHooks{}

// AfterProposalVote completes mission when a vote is cast
func (h MissionVoteHooks) AfterProposalVote(ctx sdk.Context, _ uint64, voterAddr sdk.AccAddress) {
	missions := h.k.GetAllMission(ctx)
	for _, mission := range missions {
		if mission.MissionType == types.MissionVote {
			if err := h.k.CompleteMissionFromHook(ctx, mission.CampaignId, mission.Id, voterAddr.String()); err != nil {
				ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName)).Debug("mission vote hook error", "error", err)
			}
		}
	}
}

// AfterProposalSubmission implements GovHooks
func (h MissionVoteHooks) AfterProposalSubmission(_ sdk.Context, _ uint64) {
}

// AfterProposalDeposit implements GovHooks
func (h MissionVoteHooks) AfterProposalDeposit(_ sdk.Context, _ uint64, _ sdk.AccAddress) {
}

// AfterProposalFailedMinDeposit implements GovHooks
func (h MissionVoteHooks) AfterProposalFailedMinDeposit(_ sdk.Context, _ uint64) {
}

// AfterProposalVotingPeriodEnded implements GovHooks
func (h MissionVoteHooks) AfterProposalVotingPeriodEnded(_ sdk.Context, _ uint64) {
}
