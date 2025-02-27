package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
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
func (h MissionDelegationHooks) BeforeDelegationCreated(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

// Below are the other hooks used by StakingHooks interface, they are not used by this module

// AfterValidatorCreated implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) error {
	return nil
}

// AfterValidatorRemoved implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorRemoved(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// BeforeDelegationSharesModified implements StakingHooks
func (h MissionDelegationHooks) BeforeDelegationSharesModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

// AfterDelegationModified implements StakingHooks
func (h MissionDelegationHooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, _ sdk.ValAddress) error {
	if _, found := h.k.GetUserEntry(ctx, delAddr.String()); found {
		missions := h.k.GetAllMission(ctx)
		for _, mission := range missions {
			if mission.MissionType == types.MissionDelegate {
				if err := h.k.CompleteMissionFromHook(ctx, mission.CampaignId, mission.Id, delAddr.String()); err != nil {
					h.k.Logger(ctx).Debug("mission delegation hook unsuccessful", "info", err)
				}
			}
		}
	}
	return nil
}

// BeforeValidatorSlashed implements StakingHooks
func (h MissionDelegationHooks) BeforeValidatorSlashed(_ sdk.Context, _ sdk.ValAddress, _ sdk.Dec) error {
	return nil
}

// BeforeValidatorModified implements StakingHooks
func (h MissionDelegationHooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) error {
	return nil
}

// AfterValidatorBonded implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorBonded(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// AfterValidatorBeginUnbonding implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// BeforeDelegationRemoved implements StakingHooks
func (h MissionDelegationHooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}
