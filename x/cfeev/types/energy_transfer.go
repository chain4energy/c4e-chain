package types

import (
	"cosmossdk.io/math"
)

const SafeAmountToExceedByCharger = 4

func (m *EnergyTransfer) GetCollateral() math.Int {
	if m != nil {
		return m.Collateral
	}
	return math.ZeroInt()
}
