package types

import (
	"cosmossdk.io/math"
)

const SAFE_AMOUNT_TO_EXCEED_BY_CHARGER = 4

func (m *EnergyTransfer) GetCollateral() math.Int {
	if m != nil {
		return m.Collateral
	}
	return math.ZeroInt()
}
