package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type MissionDelegationHooks struct {
	k Keeper
}

func (h MissionDelegationHooks) AfterUnbondingInitiated(ctx sdk.Context, id uint64) error {
	return nil
}

// NewMissionDelegationHooks returns a StakingHooks that triggers mission completion on delegation for an account
func (k Keeper) NewMissionDelegationHooks() MissionDelegationHooks {
	return MissionDelegationHooks{k}
}

var _ stakingtypes.StakingHooks = MissionDelegationHooks{}

// BeforeDelegationCreated completes mission when a delegation is performed
func (h MissionDelegationHooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, _ sdk.ValAddress) {

	missions := h.k.GetAllMission(ctx)
	for _, mission := range missions {
		// TODO error handling
		if mission.Id == uint64(types.DELEGATION) {
			_ = h.k.CompleteMission(ctx, mission.CampaignId, mission.Id, delAddr.String(), true)
		}
	}

}

// AfterValidatorCreated implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorCreated(_ sdk.Context, _ sdk.ValAddress) {
}

// AfterValidatorRemoved implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorRemoved(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

// BeforeDelegationSharesModified implements StakingHooks
func (h MissionDelegationHooks) BeforeDelegationSharesModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}

// AfterDelegationModified implements StakingHooks
func (h MissionDelegationHooks) AfterDelegationModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}

// BeforeValidatorSlashed implements StakingHooks
func (h MissionDelegationHooks) BeforeValidatorSlashed(_ sdk.Context, _ sdk.ValAddress, _ sdk.Dec) {
}

// BeforeValidatorModified implements StakingHooks
func (h MissionDelegationHooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) {
}

// AfterValidatorBonded implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorBonded(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

// AfterValidatorBeginUnbonding implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

// BeforeDelegationRemoved implements StakingHooks
func (h MissionDelegationHooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}
