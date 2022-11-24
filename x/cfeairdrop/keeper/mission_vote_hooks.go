package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
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
		// TODO error handling
		if mission.MissionId == uint64(types.VOTE) {
			_ = h.k.CompleteMission(ctx, mission.CampaignId, mission.MissionId, voterAddr.String())
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
