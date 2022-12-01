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

func (m Params) ValidateMinters() error {
	sort.Sort(BySequenceId(m.Minters))
	if len(m.Minters) < 1 {
		return fmt.Errorf("no minter Minters defined")
	}

	lastPos := len(m.Minters) - 1
	id := int32(0)
	for i, period := range m.Minters {
		periodId, err := m.validatePeriodOrderingId(period, id)
		if err != nil {
			return err
		}
		id = periodId

		err = m.validateEndTimeExistance(period, i, lastPos)
		if err != nil {
			return err
		}

		err = m.validateEndTimeValue(period, i, lastPos)
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

func (m Params) validatePeriodOrderingId(period *Minter, id int32) (int32, error) {
	if id == 0 {
		if period.SequenceId <= id {
			return 0, fmt.Errorf("first period ordering id must be bigger than 0, but is %d", period.SequenceId)
		}
		id = period.SequenceId
	} else {
		if period.SequenceId != id+1 {
			return 0, fmt.Errorf("missing period with ordering id %d", id+1)
		}
		id = period.SequenceId
	}
	return id, nil
}

func (m Params) validateEndTimeExistance(period *Minter, SequenceId int, lastPos int) error {
	if SequenceId == lastPos && period.EndTime != nil {
		return fmt.Errorf("last period cannot have EndTime set, but is set to %s", period.EndTime)
	}
	if SequenceId < lastPos && period.EndTime == nil {
		return fmt.Errorf("only last period can have EndTime empty")
	}
	return nil
}

func (m Params) validateEndTimeValue(period *Minter, SequenceId int, lastPos int) error {
	if lastPos > 0 {
		if SequenceId == 0 {
			if period.EndTime.Before(m.StartTime) || period.EndTime.Equal(m.StartTime) {
				return fmt.Errorf("first period end must be bigger than minter start")
			}
		} else if SequenceId < lastPos {
			prev := SequenceId - 1
			if period.EndTime.Before(*m.Minters[prev].EndTime) || period.EndTime.Equal(*m.Minters[prev].EndTime) {
				return fmt.Errorf("period with Id %d mast have EndTime bigger than period with id %d", period.SequenceId, m.Minters[prev].SequenceId)
			}
		}
	}
	return nil
}

func (m Params) ContainsId(id int32) bool {
	for _, period := range m.Minters {
		if id == period.SequenceId {
			return true
		}
	}
	return false
}

func (m *Minter) AmountToMint(logger log.Logger, state *MinterState, Minterstart time.Time, blockTime time.Time) sdk.Dec {
	switch m.Type {
	case NO_MINTING:
		return sdk.ZeroDec()
	case TIME_LINEAR_MINTER:
		return m.LinearMinting.amountToMint(state, Minterstart, *m.EndTime, blockTime)
	case PERIODIC_REDUCTION_MINTER:
		return m.ExponentialStepMinting.amountToMint(logger, state, Minterstart, m.EndTime, blockTime)
	default:
		return sdk.ZeroDec()
	}
}

func (m Minter) Validate() error {
	switch m.Type {
	case NO_MINTING:
		if m.LinearMinting != nil {
			return fmt.Errorf("period id: %d - for NO_MINTING type (0) LinearMinting must not be set", m.SequenceId)
		}
	case TIME_LINEAR_MINTER:
		if m.LinearMinting == nil {
			return fmt.Errorf("period id: %d - for TIME_LINEAR_MINTER type (1) LinearMinting must be set", m.SequenceId)
		}
		if m.EndTime == nil {
			return fmt.Errorf("period id: %d - for TIME_LINEAR_MINTER type (1) EndTime must be set", m.SequenceId)
		}
		err := m.LinearMinting.validate(m.SequenceId)
		if err != nil {
			return err
		}
	case PERIODIC_REDUCTION_MINTER:
		if m.ExponentialStepMinting == nil {
			return fmt.Errorf("period id: %d - for PERIODIC_REDUCTION_MINTER type (1) ExponentialStepMinting must be set", m.SequenceId)
		}
		err := m.ExponentialStepMinting.validate(m.SequenceId)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("period id: %d - unknow minting period type: %s", m.SequenceId, m.Type)

	}
	return nil
}

func (m *Minter) CalculateInfation(totalSupply sdk.Int, Minterstart time.Time, blockTime time.Time) sdk.Dec {
	if Minterstart.After(blockTime) {
		return sdk.ZeroDec()
	}
	switch m.Type {
	case NO_MINTING:
		return sdk.ZeroDec()
	case TIME_LINEAR_MINTER:
		return m.LinearMinting.calculateInfation(totalSupply, Minterstart, *m.EndTime)
	case PERIODIC_REDUCTION_MINTER:
		return m.ExponentialStepMinting.calculateInfation(totalSupply, Minterstart, m.EndTime, blockTime)
	default:
		return sdk.ZeroDec()
	}
}

type BySequenceId []*Minter

func (a BySequenceId) Len() int           { return len(a) }
func (a BySequenceId) Less(i, j int) bool { return a[i].SequenceId < a[j].SequenceId }
func (a BySequenceId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (m *LinearMinting) amountToMint(state *MinterState, Minterstart time.Time, EndTime time.Time, blockTime time.Time) sdk.Dec {
	if blockTime.After(EndTime) {
		return sdk.NewDecFromInt(m.Amount)
	}
	if blockTime.Before(Minterstart) {
		return sdk.ZeroDec()
	}
	amount := sdk.NewDecFromInt(m.Amount)

	passedTime := blockTime.UnixMilli() - Minterstart.UnixMilli()
	period := EndTime.UnixMilli() - Minterstart.UnixMilli()

	return amount.MulInt64(passedTime).QuoInt64(period)
}

func (m LinearMinting) validate(id int32) error {
	if m.Amount.IsNegative() {
		return fmt.Errorf("period id: %d - LinearMinting amount cannot be less than 0", id)

	}
	return nil
}

func (m *LinearMinting) calculateInfation(totalSupply sdk.Int, Minterstart time.Time, EndTime time.Time) sdk.Dec {
	if totalSupply.LTE(sdk.ZeroInt()) {
		return sdk.ZeroDec()
	}

	amount := m.Amount

	periodDuration := EndTime.Sub(Minterstart)
	mintedYearly := sdk.NewDecFromInt(amount).MulInt64(int64(year)).QuoInt64(int64(periodDuration))
	return mintedYearly.QuoInt(totalSupply)

}

func (m *ExponentialStepMinting) amountToMint(logger log.Logger, state *MinterState, Minterstart time.Time, EndTime *time.Time, blockTime time.Time) sdk.Dec {
	now := blockTime
	if EndTime != nil && blockTime.After(*EndTime) {
		now = *EndTime
	}
	passedTime := int64(now.Sub(Minterstart))
	epoch := int64(m.StepDuration) * int64(5 /* //TODO: change this valuye*/) * int64(time.Second)
	numOfPassedEpochs := passedTime / epoch
	initialEpochAmount := m.Amount.MulRaw(5 /* //TODO: change this valuye*/)

	amountToMint := sdk.ZeroDec()
	epochAmount := sdk.NewDecFromInt(initialEpochAmount)
	for i := int64(0); i < numOfPassedEpochs; i++ {
		if i > 0 {
			epochAmount = epochAmount.Mul(m.AmountMultiplier)
		}
		amountToMint = amountToMint.Add(epochAmount)
	}
	currentEpochStart := Minterstart.Add(time.Duration(numOfPassedEpochs * epoch))
	currentEpochPassedTime := now.Sub(currentEpochStart)
	currentEpochAmount := epochAmount

	logger.Debug("PRMinterMint", "blockTime", blockTime, "now", now, "passedTime", passedTime, "epoch", epoch, "numOfPassedEpochs", numOfPassedEpochs,
		"initialEpochAmount", initialEpochAmount, "epochAmount", epochAmount, "amountToMint", amountToMint, "currentEpochStart", currentEpochStart,
		"currentEpochPassedTime", currentEpochPassedTime, "currentEpochAmount", currentEpochAmount)
	if numOfPassedEpochs > 0 {
		currentEpochAmount = currentEpochAmount.Mul(m.AmountMultiplier)
	}
	currentEpochAmountToMint := currentEpochAmount.MulInt64(int64(currentEpochPassedTime)).QuoInt64(epoch)
	logger.Debug("PRMinterMintCon", "AmountMultiplier", m.AmountMultiplier, "currentEpochAmount", currentEpochAmount, "currentEpochAmountToMint", currentEpochAmountToMint)
	return amountToMint.Add(currentEpochAmountToMint)

}

func (m ExponentialStepMinting) validate(id int32) error {
	if m.Amount.IsNegative() {
		return fmt.Errorf("period id: %d - ExponentialStepMinting Amount cannot be less than 0", id)
	}
	if m.StepDuration <= 0 {
		return fmt.Errorf("period id: %d - ExponentialStepMinting StepDuration must be bigger than 0", id)
	}

	return nil
}

func (m *ExponentialStepMinting) calculateInfation(totalSupply sdk.Int, Minterstart time.Time, EndTime *time.Time, blockTime time.Time) sdk.Dec {
	if totalSupply.LTE(sdk.ZeroInt()) {
		return sdk.ZeroDec()
	}

	if EndTime != nil && (blockTime.Equal(*EndTime) || blockTime.After(*EndTime)) {
		return sdk.ZeroDec()
	}

	passedTime := int64(blockTime.Sub(Minterstart))
	epoch := int64(m.StepDuration) * int64(5 /* //TODO: change this valuye*/) * int64(time.Second)
	numOfPassedEpochs := passedTime / epoch
	initialEpochAmount := m.Amount.MulRaw(5 /* //TODO: change this valuye*/)

	epochAmount := sdk.NewDecFromInt(initialEpochAmount)
	for i := int64(0); i < numOfPassedEpochs; i++ {
		if i > 0 {
			epochAmount = epochAmount.Mul(m.AmountMultiplier)
		}
	}
	if numOfPassedEpochs > 0 {
		epochAmount = epochAmount.Mul(m.AmountMultiplier)
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
