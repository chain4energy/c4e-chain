package types

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type PeriodUnit string

const (
	Day    = "day"
	Hour   = "hour"
	Minute = "minute"
	Second = "second"
)

func ConvertVestingTypesToGenesisVestingTypes(vestingTypes *VestingTypes) []GenesisVestingType {
	gVestingTypes := []GenesisVestingType{}

	for _, vestingType := range vestingTypes.VestingTypes {
		lockupPeriodUnit, lockupPeriod := UnitsFromDuration(vestingType.LockupPeriod)
		vestingPeriodUnit, vestingPeriod := UnitsFromDuration(vestingType.VestingPeriod)

		gvt := GenesisVestingType{
			Name:              vestingType.Name,
			LockupPeriod:      lockupPeriod,
			LockupPeriodUnit:  string(lockupPeriodUnit),
			VestingPeriod:     vestingPeriod,
			VestingPeriodUnit: string(vestingPeriodUnit),
			InitialBonus:      vestingType.InitialBonus,
		}
		gVestingTypes = append(gVestingTypes, gvt)
	}

	return gVestingTypes
}

func DurationFromUnits(unit PeriodUnit, value int64) (time.Duration, error) {
	switch unit {
	case Day:
		return 24 * time.Hour * time.Duration(value), nil
	case Hour:
		return time.Hour * time.Duration(value), nil
	case Minute:
		return time.Minute * time.Duration(value), nil
	case Second:
		return time.Second * time.Duration(value), nil
	}
	return time.Duration(0), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Unknown PeriodUnit: %s", unit)
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
