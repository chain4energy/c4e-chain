package types_test

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDurationFromUnits(t *testing.T) {
	amount := int64(456)
	duration, err := types.DurationFromUnits(types.Second, amount)
	require.NoError(t, err)
	require.EqualValues(t, amount*int64(time.Second), duration)
	duration, err = types.DurationFromUnits(types.Minute, amount)
	require.NoError(t, err)
	require.EqualValues(t, amount*int64(time.Minute), duration)
	duration, err = types.DurationFromUnits(types.Hour, amount)
	require.NoError(t, err)
	require.EqualValues(t, amount*int64(time.Hour), duration)
	duration, err = types.DurationFromUnits(types.Day, amount)
	require.NoError(t, err)
	require.EqualValues(t, amount*int64(time.Hour*24), duration)

}

func TestDurationFromUnitsWrongUnit(t *testing.T) {
	_, err := types.DurationFromUnits("das", 234)
	require.EqualError(t, err, "Unknown PeriodUnit: das: invalid type")
}

func TestUnitsFromDuration(t *testing.T) {
	unit, amount := types.UnitsFromDuration(234 * time.Second)
	require.EqualValues(t, types.Second, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = types.UnitsFromDuration(234 * 60 * time.Second)
	require.EqualValues(t, types.Minute, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = types.UnitsFromDuration(234 * 60 * 60 * time.Second)
	require.EqualValues(t, types.Hour, unit)
	require.EqualValues(t, 234, amount)

	unit, amount = types.UnitsFromDuration(234 * 60 * 60 * 24 * time.Second)
	require.EqualValues(t, types.Day, unit)
	require.EqualValues(t, 234, amount)
}
