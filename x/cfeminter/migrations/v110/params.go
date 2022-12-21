package v110

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"time"
)

// MigrateParams performs in-place store migrations from v1.0.1 to v1.1.0
// The migration includes:
// - Cfeminter params structure changed
// - Remove ReductionPeriodLength from PeriodicReducstionMinter
func MigrateParams(ctx sdk.Context, paramStore *paramtypes.Subspace) error {
	var oldMinterConfig v101.Minter
	oldMinterConfigRaw := paramStore.GetRaw(ctx, v101.KeyMinter)
	if err := codec.NewLegacyAmino().UnmarshalJSON(oldMinterConfigRaw, &oldMinterConfig); err != nil {
		panic(err)
	}

	var newMinterConfig types.MinterConfig
	newMinterConfig.StartTime = oldMinterConfig.Start
	var newMinters []*types.Minter
	for _, oldMinter := range oldMinterConfig.Periods {
		var linearMinting *types.LinearMinting
		var exponentialStepMinting *types.ExponentialStepMinting

		if oldMinter.TimeLinearMinter != nil {
			linearMinting = &types.LinearMinting{
				Amount: oldMinter.TimeLinearMinter.Amount,
			}
		}
		periodicReductionMinter := oldMinter.PeriodicReductionMinter
		if periodicReductionMinter != nil {
			exponentialStepMinting = &types.ExponentialStepMinting{
				Amount:           periodicReductionMinter.MintAmount.MulRaw(int64(periodicReductionMinter.ReductionPeriodLength)),
				StepDuration:     time.Duration(oldMinter.PeriodicReductionMinter.MintPeriod*periodicReductionMinter.ReductionPeriodLength) * time.Second,
				AmountMultiplier: oldMinter.PeriodicReductionMinter.ReductionFactor,
			}
		}

		var newType string
		switch oldMinter.Type {
		case "TIME_LINEAR_MINTER":
			newType = "LINEAR_MINTING"
			break
		case "PERIODIC_REDUCTION_MINTER":
			newType = "EXPONENTIAL_STEP_MINTING"
			break
		case "NO_MINTING":
			newType = "NO_MINTING"
			break
		default:
			return fmt.Errorf("wrong minting period type")
		}

		newMinter := types.Minter{
			SequenceId:             uint32(oldMinter.Position),
			EndTime:                oldMinter.PeriodEnd,
			Type:                   newType,
			LinearMinting:          linearMinting,
			ExponentialStepMinting: exponentialStepMinting,
		}
		newMinters = append(newMinters, &newMinter)
	}
	newMinterConfig.Minters = newMinters
	if err := newMinterConfig.Validate(); err != nil {
		return err
	}
	paramStore.Set(ctx, types.KeyMinterConfig, newMinterConfig)

	return nil
}
