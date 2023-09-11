package keeper

import (
	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
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
	if _, found := h.k.GetUserEntry(ctx, voterAddr.String()); found {
		missions := h.k.GetAllMission(ctx)
		for _, mission := range missions {
			if mission.MissionType == types.MissionVote {
				if err := h.k.CompleteMissionFromHook(ctx, mission.CampaignId, mission.Id, voterAddr.String()); err != nil {
					h.k.Logger(ctx).Debug("mission vote hook unsuccessful", "info", err)
				}
			}
		}
	}
}

// Below are the other hooks used by GovHooks interface, they are not used by this module

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
