package types

import (
	"fmt"
	"sort"
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

func (m MintingPeriod) Validate() error {
	switch m.Type {
	case MintingPeriod_NO_MINTING:
		if m.TimeLinearMinter != nil {
			return fmt.Errorf("period id: %d - for NO_MINTING type (0) TimeLinearMinter must not be set", m.OrderingId)
		}
	case MintingPeriod_TIME_LINEAR_MINTER:
		if m.TimeLinearMinter == nil {
			return fmt.Errorf("period id: %d - for MintingPeriod_TIME_LINEAR_MINTER type (1) TimeLinearMinter must be set", m.OrderingId)
		}
		err := m.TimeLinearMinter.validate(m.OrderingId)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("period id: %d - unknow minting period type: %d", m.OrderingId, m.Type)

	}
	return nil
}

func (m TimeLinearMinter) validate(id int32) error {
	if m.Amount.IsNegative() {
		return fmt.Errorf("period id: %d - TimeLinearMinter amount cannot be less than 0", id)

	}
	return nil
}

func (m Minter) Validate() error {
	sort.Sort(ByOrderingId(m.Periods))
	id := int32(0)
	lastPos := len(m.Periods) - 1
	if len(m.Periods) < 1 {
		return fmt.Errorf("no minter periods defined")
	}
	for i, period := range m.Periods {
		if id == 0 {
			if period.OrderingId <= id {
				return fmt.Errorf("first period ordering id must be bigger than 0, but is %d", period.OrderingId)
			}
			id = period.OrderingId
		} else {
			if period.OrderingId != id+1 {
				return fmt.Errorf("missing period with ordering id %d", id+1)
			}
			id = period.OrderingId
		}
		if i == lastPos && period.PeriodEnd != nil {
			return fmt.Errorf("last period cannot have PeriodEnd set, but is set to %s", period.PeriodEnd)
		}
		if lastPos > 0 {
			if i == 0 {
				if period.PeriodEnd.Before(m.Start) || period.PeriodEnd.Equal(m.Start) {
					return fmt.Errorf("first period end must be bigger than minter start")
				}
			} else if i < lastPos {
				prev := i - 1
				if period.PeriodEnd.Before(*m.Periods[prev].PeriodEnd) || period.PeriodEnd.Equal(*m.Periods[prev].PeriodEnd) {
					return fmt.Errorf("period with Id %d mast have PeriodEnd bigger than period with id %d", period.OrderingId, m.Periods[prev].OrderingId)
				}
			}
		}
		err := period.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Minter) ContainsId(id int32) bool {
	for _, period := range m.Periods {
		if id == period.OrderingId {
			return true
		}
	}
	return false
}


type ByOrderingId []*MintingPeriod

func (a ByOrderingId) Len() int           { return len(a) }
func (a ByOrderingId) Less(i, j int) bool { return a[i].OrderingId < a[j].OrderingId }
func (a ByOrderingId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }


func (m MinterState) Validate() error {
	if m.AmountMinted.IsNegative() {
		return fmt.Errorf("minter state amount cannot be less than 0")

	}
	return nil
}