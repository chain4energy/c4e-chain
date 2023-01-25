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
	NoMintingType              string = "NO_MINTING"
	LinearMintingType          string = "LINEAR_MINTING"
	ExponentialStepMintingType string = "EXPONENTIAL_STEP_MINTING"
)

func (params MinterConfig) Validate() error {
	if len(params.Minters) < 1 {
		return fmt.Errorf("no minters defined")
	}

	for i, minter := range params.Minters {
		if minter == nil {
			return fmt.Errorf("minter on position %d cannot be nil", i+1)
		}
	}

	sort.Sort(BySequenceId(params.Minters))
	lastPos := len(params.Minters) - 1
	id := uint32(0)
	for i, minter := range params.Minters {
		minterId, err := params.validateMinterOrderingId(minter, id)
		if err != nil {
			return err
		}
		id = minterId

		err = params.validateEndTimeExistance(minter, i, lastPos)
		if err != nil {
			return err
		}

		err = params.validateMintersEndTimeValue(minter, i, lastPos)
		if err != nil {
			return err
		}

		if err = minter.validate(); err != nil {
			return fmt.Errorf("minter with id %d validation error: %w", minter.SequenceId, err)
		}
	}
	return nil
}

func (params MinterConfig) validateMinterOrderingId(minter *Minter, id uint32) (uint32, error) {
	if id == 0 {
		if minter.SequenceId <= id {
			return 0, fmt.Errorf("first minter sequence id must be bigger than 0, but is %d", minter.SequenceId)
		}
		id = minter.SequenceId
	} else {
		if minter.SequenceId != id+1 {
			return 0, fmt.Errorf("missing minter with sequence id %d", id+1)
		}
		id = minter.SequenceId
	}
	return id, nil
}

func (params MinterConfig) validateEndTimeExistance(minter *Minter, sequenceId int, lastPos int) error {
	if sequenceId == lastPos && minter.EndTime != nil {
		return fmt.Errorf("last minter cannot have EndTime set, but is set to %s", minter.EndTime)
	}
	if sequenceId < lastPos && minter.EndTime == nil {
		return fmt.Errorf("only last minter can have EndTime empty")
	}
	return nil
}

func (params MinterConfig) validateMintersEndTimeValue(minter *Minter, sequenceId int, lastPos int) error {
	if lastPos > 0 {
		if sequenceId == 0 {
			if minter.EndTime.Before(params.StartTime) || minter.EndTime.Equal(params.StartTime) {
				return fmt.Errorf("first minter end must be bigger than minter start")
			}
		} else if sequenceId < lastPos {
			prev := sequenceId - 1
			if minter.EndTime.Before(*params.Minters[prev].EndTime) || minter.EndTime.Equal(*params.Minters[prev].EndTime) {
				return fmt.Errorf("minter with sequence id %d mast have EndTime bigger than minter with sequence id %d", minter.SequenceId, params.Minters[prev].SequenceId)
			}
		}
	}
	return nil
}

func (params MinterConfig) ContainsMinter(sequenceId uint32) bool {
	for _, minter := range params.Minters {
		if sequenceId == minter.SequenceId {
			return true
		}
	}
	return false
}

func (m Minter) validate() error {
	switch m.Type {
	case NoMintingType:
		if m.LinearMinting != nil || m.ExponentialStepMinting != nil {
			return fmt.Errorf("for NO_MINTING type (0) LinearMinting and ExponentialStepMinting cannot be set")
		}
	case LinearMintingType:
		if m.ExponentialStepMinting != nil {
			return fmt.Errorf("for LinearMintingType type (1) ExponentialStepMinting cannot be set")
		}
		if m.EndTime == nil {
			return fmt.Errorf("for LinearMintingType type (1) EndTime must be set")
		}
		if err := m.LinearMinting.validate(); err != nil {
			return fmt.Errorf("LinearMintingType error: %w", err)
		}
	case ExponentialStepMintingType:
		if m.LinearMinting != nil {
			return fmt.Errorf("for ExponentialStepMintingType type (2) LinearMinting cannot be set")
		}
		if err := m.ExponentialStepMinting.validate(); err != nil {
			return fmt.Errorf("ExponentialStepMintingType error: %w", err)
		}
	default:
		return fmt.Errorf("unknow minting configuration type: %s", m.Type)
	}
	return nil
}

func (m *LinearMinting) validate() error {
	if m == nil {
		return fmt.Errorf("for LinearMintingType type (1) LinearMinting must be set")
	}
	if m.Amount.IsNil() {
		return fmt.Errorf("amount cannot be nil")
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("amount cannot be less than 0")
	}
	return nil
}

func (m *ExponentialStepMinting) validate() error {
	if m == nil {
		return fmt.Errorf("for ExponentialStepMintingType type (2) ExponentialStepMinting must be set")
	}
	if m.Amount.IsNil() {
		return fmt.Errorf("amount cannot be nil")
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("amount cannot be less than 0")
	}
	if m.AmountMultiplier.IsNil() {
		return fmt.Errorf("amountMultiplier cannot be nil")
	}
	if m.AmountMultiplier.IsNegative() {
		return fmt.Errorf("amountMultiplier cannot be less than 0")
	}
	if m.StepDuration <= 0 {
		return fmt.Errorf("stepDuration must be bigger than 0")
	}
	return nil
}

func (m MinterState) Validate() error {
	if m.AmountMinted.IsNil() {
		return fmt.Errorf("minter state validation error: amountMinted cannot be nil")
	}
	if m.AmountMinted.IsNegative() {
		return fmt.Errorf("minter state validation error: amountMinted cannot be less than 0")
	}
	if m.RemainderFromPreviousPeriod.IsNil() {
		return fmt.Errorf("minter state validation error: remainderFromPreviousPeriod cannot be nil")
	}
	if m.RemainderFromPreviousPeriod.IsNegative() {
		return fmt.Errorf("minter state validation error: remainderFromPreviousPeriod cannot be less than 0")
	}
	if m.RemainderToMint.IsNil() {
		return fmt.Errorf("minter state validation error: remainderToMint cannot be nil")
	}
	if m.RemainderToMint.IsNegative() {
		return fmt.Errorf("minter state validation error: remainderToMint cannot be less than 0")
	}
	return nil
}

type BySequenceId []*Minter

func (a BySequenceId) Len() int           { return len(a) }
func (a BySequenceId) Less(i, j int) bool { return a[i].SequenceId < a[j].SequenceId }
func (a BySequenceId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (m *Minter) CalculateInflation(totalSupply sdk.Int, startTime time.Time, blockTime time.Time) sdk.Dec {
	if startTime.After(blockTime) {
		return sdk.ZeroDec()
	}
	switch m.Type {
	case NoMintingType:
		return sdk.ZeroDec()
	case LinearMintingType:
		return m.LinearMinting.calculateInflation(totalSupply, startTime, *m.EndTime)
	case ExponentialStepMintingType:
		return m.ExponentialStepMinting.calculateInflation(totalSupply, startTime, m.EndTime, blockTime)
	default:
		return sdk.ZeroDec()
	}
}

func (m *LinearMinting) calculateInflation(totalSupply sdk.Int, minterStart time.Time, endTime time.Time) sdk.Dec {
	if totalSupply.LTE(sdk.ZeroInt()) {
		return sdk.ZeroDec()
	}

	periodDuration := endTime.Sub(minterStart)
	mintedYearly := sdk.NewDecFromInt(m.Amount).MulInt64(int64(year)).QuoInt64(int64(periodDuration))
	return mintedYearly.QuoInt(totalSupply)
}

func (m *ExponentialStepMinting) calculateInflation(totalSupply sdk.Int, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	if totalSupply.LTE(sdk.ZeroInt()) {
		return sdk.ZeroDec()
	}

	if endTime != nil && (blockTime.Equal(*endTime) || blockTime.After(*endTime)) {
		return sdk.ZeroDec()
	}

	passedTime := int64(blockTime.Sub(startTime))
	epoch := int64(m.StepDuration)
	numOfPassedEpochs := passedTime / epoch
	epochAmount := sdk.NewDecFromInt(m.Amount)

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

func (m *Minter) AmountToMint(logger log.Logger, state *MinterState, minterStart time.Time, blockTime time.Time) sdk.Dec {
	switch m.Type {
	case NoMintingType:
		return sdk.ZeroDec()
	case LinearMintingType:
		return m.LinearMinting.amountToMint(minterStart, *m.EndTime, blockTime)
	case ExponentialStepMintingType:
		return m.ExponentialStepMinting.amountToMint(logger, minterStart, m.EndTime, blockTime)
	default:
		return sdk.ZeroDec()
	}
}

func (m *LinearMinting) amountToMint(minterStart time.Time, EndTime time.Time, blockTime time.Time) sdk.Dec {
	if blockTime.After(EndTime) {
		return sdk.NewDecFromInt(m.Amount)
	}
	if blockTime.Before(minterStart) {
		return sdk.ZeroDec()
	}
	amount := sdk.NewDecFromInt(m.Amount)

	passedTime := blockTime.UnixMilli() - minterStart.UnixMilli()
	period := EndTime.UnixMilli() - minterStart.UnixMilli()

	return amount.MulInt64(passedTime).QuoInt64(period)
}

func (m *ExponentialStepMinting) amountToMint(logger log.Logger, startTIme time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	now := blockTime
	if endTime != nil && blockTime.After(*endTime) {
		now = *endTime
	}
	passedTime := int64(now.Sub(startTIme))
	epoch := int64(m.StepDuration)
	numOfPassedEpochs := passedTime / epoch

	amountToMint := sdk.ZeroDec()
	epochAmount := sdk.NewDecFromInt(m.Amount)
	for i := int64(0); i < numOfPassedEpochs; i++ {
		if i > 0 {
			epochAmount = epochAmount.Mul(m.AmountMultiplier)
		}
		amountToMint = amountToMint.Add(epochAmount)
	}
	currentEpochStart := startTIme.Add(time.Duration(numOfPassedEpochs * epoch))
	currentEpochPassedTime := now.Sub(currentEpochStart)
	currentEpochAmount := epochAmount

	logger.Debug("ESMintingMint", "blockTime", blockTime, "now", now, "passedTime", passedTime, "epoch", epoch, "numOfPassedEpochs", numOfPassedEpochs,
		"Amount", m.Amount, "epochAmount", epochAmount, "amountToMint", amountToMint, "currentEpochStart", currentEpochStart,
		"currentEpochPassedTime", currentEpochPassedTime, "currentEpochAmount", currentEpochAmount)
	if numOfPassedEpochs > 0 {
		currentEpochAmount = currentEpochAmount.Mul(m.AmountMultiplier)
	}

	currentEpochAmountToMint := currentEpochAmount.MulInt64(int64(currentEpochPassedTime)).QuoInt64(epoch)
	logger.Debug("ESMintingMintCon", "AmountMultiplier", m.AmountMultiplier, "currentEpochAmount", currentEpochAmount, "currentEpochAmountToMint", currentEpochAmountToMint)
	return amountToMint.Add(currentEpochAmountToMint)
}
