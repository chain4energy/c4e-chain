package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *VestingPool) GetCurrentlyLocked() sdk.Int {
	return m.InitiallyLocked.Sub(m.Sent).Sub(m.Withdrawn)
}
