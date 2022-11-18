package types

import (
	"fmt"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

const year = time.Hour * 24 * 365

const ( // MintingPeriod types
	NO_MINTING                string = "NO_MINTING"
	TIME_LINEAR_MINTER        string = "TIME_LINEAR_MINTER"
	PERIODIC_REDUCTION_MINTER string = "PERIODIC_REDUCTION_MINTER"
)

func (m Minter) Validate() error {
	sort.Sort(ByPosition(m.Periods))
	if len(m.Periods) < 1 {
		return fmt.Errorf("no minter periods defined")
	}
	return m.validatePeriods()
}

func (m Minter) validatePeriods() error {
	lastPos := len(m.Periods) - 1
	id := int32(0)
	for i, period := range m.Periods {
		periodId, err := m.validatePeriodOrderingId(period, id)
		if err != nil {
			return err
		}
		id = periodId

		err = m.validatePeriodEndExistance(period, i, lastPos)
		if err != nil {
			return err
		}

		err = m.validatePeriodEndValue(period, i, lastPos)
		if err != nil {
			return err
		}

		err = period.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Minter) validatePeriodOrderingId(period *MintingPeriod, id int32) (int32, error) {
	if id == 0 {
		if period.Position <= id {
			return 0, fmt.Errorf("first period ordering id must be bigger than 0, but is %d", period.Position)
		}
		id = period.Position
	} else {
		if period.Position != id+1 {
			return 0, fmt.Errorf("missing period with ordering id %d", id+1)
		}
		id = period.Position
	}
	return id, nil
}

func (m Minter) validatePeriodEndExistance(period *MintingPeriod, position int, lastPos int) error {
	if position == lastPos && period.PeriodEnd != nil {
		return fmt.Errorf("last period cannot have PeriodEnd set, but is set to %s", period.PeriodEnd)
	}
	if position < lastPos && period.PeriodEnd == nil {
		return fmt.Errorf("only last period can have PeriodEnd empty")
	}
	return nil
}

func (m Minter) validatePeriodEndValue(period *MintingPeriod, position int, lastPos int) error {
	if lastPos > 0 {
		if position == 0 {
			if period.PeriodEnd.Before(m.Start) || period.PeriodEnd.Equal(m.Start) {
				return fmt.Errorf("first period end must be bigger than minter start")
			}
		} else if position < lastPos {
			prev := position - 1
			if period.PeriodEnd.Before(*m.Periods[prev].PeriodEnd) || period.PeriodEnd.Equal(*m.Periods[prev].PeriodEnd) {
				return fmt.Errorf("period with Id %d mast have PeriodEnd bigger than period with id %d", period.Position, m.Periods[prev].Position)
			}
		}
	}
	return nil
}

func (m Minter) ContainsId(id int32) bool {
	for _, period := range m.Periods {
		if id == period.Position {
			return true
		}
	}
	return false
}

func (m *MintingPeriod) AmountToMint(logger log.Logger, state *MinterState, periodStart time.Time, blockTime time.Time) sdk.Dec {
	switch m.Type {
	case NO_MINTING:
		return sdk.ZeroDec()
	case TIME_LINEAR_MINTER:
		return m.TimeLinearMinter.amountToMint(state, periodStart, *m.PeriodEnd, blockTime)
	case PERIODIC_REDUCTION_MINTER:
		return m.PeriodicReductionMinter.amountToMint(logger, state, periodStart, m.PeriodEnd, blockTime)
	default:
		return sdk.ZeroDec()
	}
}

func (m MintingPeriod) Validate() error {
	switch m.Type {
	case NO_MINTING:
		if m.TimeLinearMinter != nil {
			return fmt.Errorf("period id: %d - for NO_MINTING type (0) TimeLinearMinter must not be set", m.Position)
		}
	case TIME_LINEAR_MINTER:
		if m.TimeLinearMinter == nil {
			return fmt.Errorf("period id: %d - for TIME_LINEAR_MINTER type (1) TimeLinearMinter must be set", m.Position)
		}
		if m.PeriodEnd == nil {
			return fmt.Errorf("period id: %d - for TIME_LINEAR_MINTER type (1) PeriodEnd must be set", m.Position)
		}
		err := m.TimeLinearMinter.validate(m.Position)
		if err != nil {
			return err
		}
	case PERIODIC_REDUCTION_MINTER:
		if m.PeriodicReductionMinter == nil {
			return fmt.Errorf("period id: %d - for PERIODIC_REDUCTION_MINTER type (1) PeriodicReductionMinter must be set", m.Position)
		}
		err := m.PeriodicReductionMinter.validate(m.Position)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("period id: %d - unknow minting period type: %s", m.Position, m.Type)

	}
	return nil
}

func (m *MintingPeriod) CalculateInfation(totalSupply sdk.Int, periodStart time.Time, blockTime time.Time) sdk.Dec {
	if periodStart.After(blockTime) {
		return sdk.ZeroDec()
	}
	switch m.Type {
	case NO_MINTING:
		return sdk.ZeroDec()
	case TIME_LINEAR_MINTER:
		return m.TimeLinearMinter.calculateInfation(totalSupply, periodStart, *m.PeriodEnd)
	case PERIODIC_REDUCTION_MINTER:
		return m.PeriodicReductionMinter.calculateInfation(totalSupply, periodStart, m.PeriodEnd, blockTime)
	default:
		return sdk.ZeroDec()
	}
}

type ByPosition []*MintingPeriod

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (m *TimeLinearMinter) amountToMint(state *MinterState, periodStart time.Time, periodEnd time.Time, blockTime time.Time) sdk.Dec {
	if blockTime.After(periodEnd) {
		return sdk.NewDecFromInt(m.Amount)
	}
	if blockTime.Before(periodStart) {
		return sdk.ZeroDec()
	}
	amount := sdk.NewDecFromInt(m.Amount)

	passedTime := blockTime.UnixMilli() - periodStart.UnixMilli()
	period := periodEnd.UnixMilli() - periodStart.UnixMilli()

	return amount.MulInt64(passedTime).QuoInt64(period)
}

func (m TimeLinearMinter) validate(id int32) error {
	if m.Amount.IsNegative() {
		return fmt.Errorf("period id: %d - TimeLinearMinter amount cannot be less than 0", id)

	}
	return nil
}

func (m *TimeLinearMinter) calculateInfation(totalSupply sdk.Int, periodStart time.Time, periodEnd time.Time) sdk.Dec {
	if totalSupply.LTE(sdk.ZeroInt()) {
		return sdk.ZeroDec()
	}

	amount := m.Amount

	periodDuration := periodEnd.Sub(periodStart)
	mintedYearly := sdk.NewDecFromInt(amount).MulInt64(int64(year)).QuoInt64(int64(periodDuration))
	return mintedYearly.QuoInt(totalSupply)

}

func (m *PeriodicReductionMinter) amountToMint(logger log.Logger, state *MinterState, periodStart time.Time, periodEnd *time.Time, blockTime time.Time) sdk.Dec {
	now := blockTime
	if periodEnd != nil && blockTime.After(*periodEnd) {
		now = *periodEnd
	}
	passedTime := int64(now.Sub(periodStart))
	epoch := int64(m.MintPeriod) * int64(m.ReductionPeriodLength) * int64(time.Second)
	numOfPassedEpochs := passedTime / epoch
	initialEpochAmount := m.MintAmount.MulRaw(int64(m.ReductionPeriodLength))

	amountToMint := sdk.ZeroDec()
	epochAmount := sdk.NewDecFromInt(initialEpochAmount)
	for i := int64(0); i < numOfPassedEpochs; i++ {
		if i > 0 {
			epochAmount = epochAmount.Mul(m.ReductionFactor)
		}
		amountToMint = amountToMint.Add(epochAmount)
	}
	currentEpochStart := periodStart.Add(time.Duration(numOfPassedEpochs * epoch))
	currentEpochPassedTime := now.Sub(currentEpochStart)
	currentEpochAmount := epochAmount

	logger.Debug("PRMinterMint", "blockTime", blockTime, "now", now, "passedTime", passedTime, "epoch", epoch, "numOfPassedEpochs", numOfPassedEpochs,
		"initialEpochAmount", initialEpochAmount, "epochAmount", epochAmount, "amountToMint", amountToMint, "currentEpochStart", currentEpochStart,
		"currentEpochPassedTime", currentEpochPassedTime, "currentEpochAmount", currentEpochAmount)
	if numOfPassedEpochs > 0 {
		currentEpochAmount = currentEpochAmount.Mul(m.ReductionFactor)
	}
	currentEpochAmountToMint := currentEpochAmount.MulInt64(int64(currentEpochPassedTime)).QuoInt64(epoch)
	logger.Debug("PRMinterMintCon", "ReductionFactor", m.ReductionFactor, "currentEpochAmount", currentEpochAmount, "currentEpochAmountToMint", currentEpochAmountToMint)
	return amountToMint.Add(currentEpochAmountToMint)

}

func (m PeriodicReductionMinter) validate(id int32) error {
	if m.MintAmount.IsNegative() {
		return fmt.Errorf("period id: %d - PeriodicReductionMinter MintAmount cannot be less than 0", id)
	}
	if m.MintPeriod <= 0 {
		return fmt.Errorf("period id: %d - PeriodicReductionMinter MintPeriod must be bigger than 0", id)
	}
	if m.ReductionPeriodLength <= 0 {
		return fmt.Errorf("period id: %d - PeriodicReductionMinter ReductionPeriodLength must be bigger than 0", id)
	}
	return nil
}

func (m *PeriodicReductionMinter) calculateInfation(totalSupply sdk.Int, periodStart time.Time, periodEnd *time.Time, blockTime time.Time) sdk.Dec {
	if totalSupply.LTE(sdk.ZeroInt()) {
		return sdk.ZeroDec()
	}

	if periodEnd != nil && (blockTime.Equal(*periodEnd) || blockTime.After(*periodEnd)) {
		return sdk.ZeroDec()
	}

	passedTime := int64(blockTime.Sub(periodStart))
	epoch := int64(m.MintPeriod) * int64(m.ReductionPeriodLength) * int64(time.Second)
	numOfPassedEpochs := passedTime / epoch
	initialEpochAmount := m.MintAmount.MulRaw(int64(m.ReductionPeriodLength))

	epochAmount := sdk.NewDecFromInt(initialEpochAmount)
	for i := int64(0); i < numOfPassedEpochs; i++ {
		if i > 0 {
			epochAmount = epochAmount.Mul(m.ReductionFactor)
		}
	}
	if numOfPassedEpochs > 0 {
		epochAmount = epochAmount.Mul(m.ReductionFactor)
	}
	mintedYearly := epochAmount.MulInt64(int64(year)).QuoInt64(epoch)
	return mintedYearly.QuoInt(totalSupply)
}

func (m MinterState) Validate() error {
	if m.AmountMinted.IsNegative() {
		return fmt.Errorf("minter state amount cannot be less than 0")

	}
	if m.RemainderFromPreviousPeriod.IsNegative() {
		return fmt.Errorf("minter remainder from previous period amount cannot be less than 0")

	}
	if m.RemainderToMint.IsNegative() {
		return fmt.Errorf("minter remainder to mint amount cannot be less than 0")

	}
	return nil
}
