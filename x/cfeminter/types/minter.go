package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *MintingPeriod) AmountToMint(state *MinterState, periodStart time.Time, blockTime time.Time) sdk.Int {
	switch m.Type {
	case MintingPeriod_NO_MINTING:
		return sdk.ZeroInt()
	case MintingPeriod_TIME_LINEAR_MINTER:
		return m.TimeLinearMinter.amountToMint(state, periodStart, *m.PeriodEnd, blockTime)
	default:
		return sdk.ZeroInt()
	}
}

func (m *TimeLinearMinter) amountToMint(state *MinterState, periodStart time.Time, periodEnd time.Time, blockTime time.Time) sdk.Int {
	amount := m.Amount
	if blockTime.After(periodEnd) {
		return amount.Sub(state.AmountMinted)
	}
	if blockTime.Before(periodStart) {
		return sdk.ZeroInt()
	}
	passedTime := blockTime.UnixMilli() - periodStart.UnixMilli()
	period := periodEnd.UnixMilli() - periodStart.UnixMilli()
	return amount.MulRaw(passedTime).QuoRaw(period).Sub(state.AmountMinted)
}
