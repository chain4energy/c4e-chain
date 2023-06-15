package types

import (
	"cosmossdk.io/math"
)

func (m *EnergyTransfer) GetCollateral() math.Int {
	if m != nil {
		return m.Collateral
	}
	return math.ZeroInt()
}
