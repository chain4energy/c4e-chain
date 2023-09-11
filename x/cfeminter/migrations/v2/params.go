package v2

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/v2/types/subspace"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/migrations/v1"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// MigrateParams performs in-place store migrations from v1.0.1 to v1.1.0
// The migration includes:
// - Cfeminter params structure changed
// - Remove ReductionPeriodLength from PeriodicReducstionMinter
func MigrateParams(ctx sdk.Context, paramStore subspace.Subspace) error {
	var oldMinterConfig v1.Minter
	if !paramStore.HasKeyTable() {
		paramStore.WithKeyTable(ParamKeyTable())
	}
	oldMinterConfigRaw := paramStore.GetRaw(ctx, v1.KeyMinter)
	if err := codec.NewLegacyAmino().UnmarshalJSON(oldMinterConfigRaw, &oldMinterConfig); err != nil {
		panic(err)
	}

	var newMinterConfig MinterConfig
	newMinterConfig.StartTime = oldMinterConfig.Start
	var newMinters []*Minter
	for _, oldMinter := range oldMinterConfig.Periods {
		var linearMinting *LinearMinting
		var exponentialStepMinting *ExponentialStepMinting

		if oldMinter.TimeLinearMinter != nil {
			linearMinting = &LinearMinting{
				Amount: oldMinter.TimeLinearMinter.Amount,
			}
		}
		periodicReductionMinter := oldMinter.PeriodicReductionMinter
		if periodicReductionMinter != nil {
			exponentialStepMinting = &ExponentialStepMinting{
				Amount:           periodicReductionMinter.MintAmount.MulRaw(int64(periodicReductionMinter.ReductionPeriodLength)),
				StepDuration:     time.Duration(oldMinter.PeriodicReductionMinter.MintPeriod*periodicReductionMinter.ReductionPeriodLength) * time.Second,
				AmountMultiplier: oldMinter.PeriodicReductionMinter.ReductionFactor,
			}
		}

		var newType string
		switch oldMinter.Type {
		case "TIME_LINEAR_MINTER":
			newType = "LINEAR_MINTING"
		case "PERIODIC_REDUCTION_MINTER":
			newType = "EXPONENTIAL_STEP_MINTING"
		case "NO_MINTING":
			newType = "NO_MINTING"
		default:
			return fmt.Errorf("wrong minting period type")
		}

		newMinter := Minter{
			SequenceId:             uint32(oldMinter.Position),
			EndTime:                oldMinter.PeriodEnd,
			Type:                   newType,
			LinearMinting:          linearMinting,
			ExponentialStepMinting: exponentialStepMinting,
		}
		newMinters = append(newMinters, &newMinter)
	}
	newMinterConfig.Minters = newMinters
	paramStore.Set(ctx, KeyMinterConfig, newMinterConfig)

	return nil
}
