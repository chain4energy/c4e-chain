package v110

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"time"
)

// MigrateParams performs in-place store migrations from v1.0.0 to v1.0.1. The
// migration includes:
//
// - SubDistributor params structure changed.
// - BurnShare and Share now must be set between 0 and 1, not 0 and 100.
func MigrateParams(ctx sdk.Context, storeKey storetypes.StoreKey, paramStore *paramtypes.Subspace) error {
	var oldMinterConfig v101.Minter
	oldMinterConfigRaw := paramStore.GetRaw(ctx, v101.KeyMinter)
	if err := codec.NewLegacyAmino().UnmarshalJSON(oldMinterConfigRaw, &oldMinterConfig); err != nil {
		panic(err)
	}

	var newMinterConfig types.MinterConfig
	newMinterConfig.StartTime = oldMinterConfig.Start
	var newMinters []*types.Minter
	for _, oldMinter := range oldMinterConfig.Periods {
		var linearMinting types.LinearMinting
		var exponentialStepMinting types.ExponentialStepMinting

		if oldMinter.TimeLinearMinter != nil {
			linearMinting.Amount = oldMinter.TimeLinearMinter.Amount
		}
		periodicReductionMinter := oldMinter.PeriodicReductionMinter
		if periodicReductionMinter != nil {
			exponentialStepMinting.Amount = periodicReductionMinter.MintAmount.MulRaw(int64(periodicReductionMinter.ReductionPeriodLength))
			exponentialStepMinting.StepDuration = time.Duration(oldMinter.PeriodicReductionMinter.MintPeriod * 1000000000 * periodicReductionMinter.ReductionPeriodLength)
			exponentialStepMinting.AmountMultiplier = oldMinter.PeriodicReductionMinter.ReductionFactor
		}

		newMinter := types.Minter{
			SequenceId:             uint32(oldMinter.Position),
			EndTime:                oldMinter.PeriodEnd,
			Type:                   oldMinter.Type,
			LinearMinting:          &linearMinting,
			ExponentialStepMinting: &exponentialStepMinting,
		}
		if err := newMinter.Validate(); err != nil {
			return err
		}
		newMinters = append(newMinters, &newMinter)
	}
	newMinterConfig.Minters = newMinters
	err := newMinterConfig.ValidateMinters()
	if err != nil {
		return err
	}
	paramStore.Set(ctx, types.KeyMinterConfig, newMinterConfig)
	return nil
}
