package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

type PeriodUnit string

const (
	Day    = "day"
	Hour   = "hour"
	Minute = "minute"
	Second = "second"
)

func ConvertVestingTypesToGenesisVestingTypes(vestingTypes *types.VestingTypes) []types.GenesisVestingType {
	gVestingTypes := []types.GenesisVestingType{}

	for _, vestingType := range vestingTypes.VestingTypes {
		lockupPeriodUnit, lockupPeriod := UnitsFromDuration(vestingType.LockupPeriod)
		vestingPeriodUnit, vestingPeriod := UnitsFromDuration(vestingType.VestingPeriod)

		gvt := types.GenesisVestingType{
			Name:              vestingType.Name,
			LockupPeriod:      lockupPeriod,
			LockupPeriodUnit:  string(lockupPeriodUnit),
			VestingPeriod:     vestingPeriod,
			VestingPeriodUnit: string(vestingPeriodUnit),
		}
		gVestingTypes = append(gVestingTypes, gvt)
	}

	return gVestingTypes
}

func DurationFromUnits(ctx sdk.Context, unit PeriodUnit, value int64) time.Duration {
	switch unit {
	case Day:
		return 24 * time.Hour * time.Duration(value)
	case Hour:
		return time.Hour * time.Duration(value)
	case Minute:
		return time.Minute * time.Duration(value)
	case Second:
		return time.Second * time.Duration(value)
	}

	Keeper{}.Logger(ctx).Debug("set account vestings", "vesting utils duration from units unknown periodunit error",
		"unit", unit)
	panic(sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Unknown PeriodUnit: %s", unit))
}

func UnitsFromDuration(duration time.Duration) (unit PeriodUnit, value int64) {
	if duration%(24*time.Hour) == 0 {
		return Day, int64(duration / (24 * time.Hour))
	}
	if duration%(time.Hour) == 0 {
		return Hour, int64(duration / (time.Hour))
	}
	if duration%(time.Minute) == 0 {
		return Minute, int64(duration / (time.Minute))
	}
	return Second, int64(duration / (time.Second))
}
