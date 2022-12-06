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
	NO_MINTING               string = "NO_MINTING"
	LINEAR_MINTING           string = "LINEAR_MINTING"
	EXPONENTIAL_STEP_MINTING string = "EXPONENTIAL_STEP_MINTING"
)

type Minters []*Minter

func (params MinterConfig) ValidateMinters() error {
	sort.Sort(BySequenceId(params.Minters))
	if len(params.Minters) < 1 {
		return fmt.Errorf("no minters defined")
	}

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

		err = minter.Validate()
		if err != nil {
			return err
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

func (params MinterConfig) validateEndTimeExistance(minter *Minter, SequenceId int, lastPos int) error {
	if SequenceId == lastPos && minter.EndTime != nil {
		return fmt.Errorf("last minter cannot have EndTime set, but is set to %s", minter.EndTime)
	}
	if SequenceId < lastPos && minter.EndTime == nil {
		return fmt.Errorf("only last minter can have EndTime empty")
	}
	return nil
}

func (params MinterConfig) validateMintersEndTimeValue(minter *Minter, SequenceId int, lastPos int) error {
	if lastPos > 0 {
		if SequenceId == 0 {
			if minter.EndTime.Before(params.StartTime) || minter.EndTime.Equal(params.StartTime) {
				return fmt.Errorf("first minter end must be bigger than minter start")
			}
		} else if SequenceId < lastPos {
			prev := SequenceId - 1
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

func (m *Minter) AmountToMint(logger log.Logger, state *MinterState, startTime time.Time, blockTime time.Time) sdk.Dec {
	switch m.Type {
	case NO_MINTING:
		return sdk.ZeroDec()
	case LINEAR_MINTING:
		return m.LinearMinting.amountToMint(startTime, *m.EndTime, blockTime)
	case EXPONENTIAL_STEP_MINTING:
		return m.ExponentialStepMinting.amountToMint(logger, startTime, m.EndTime, blockTime)
	default:
		return sdk.ZeroDec()
	}
}

func (m Minter) Validate() error {
	switch m.Type {
	case NO_MINTING:
		if m.LinearMinting != nil {
			return fmt.Errorf("minter sequence id: %d - for NO_MINTING type (0) LinearMinting must not be set", m.SequenceId)
		}
	case LINEAR_MINTING:
		if m.LinearMinting == nil {
			return fmt.Errorf("minter sequence id: %d - for LINEAR_MINTING type (1) LinearMinting must be set", m.SequenceId)
		}
		if m.EndTime == nil {
			return fmt.Errorf("minter sequence id: %d - for LINEAR_MINTING type (1) EndTime must be set", m.SequenceId)
		}
		err := m.LinearMinting.validate(m.SequenceId)
		if err != nil {
			return err
		}
	case EXPONENTIAL_STEP_MINTING:
		if m.ExponentialStepMinting == nil {
			return fmt.Errorf("minter sequence id: %d - for EXPONENTIAL_STEP_MINTING type (1) ExponentialStepMinting must be set", m.SequenceId)
		}
		err := m.ExponentialStepMinting.validate(m.SequenceId)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("minter sequence id: %d - unknow minting configuration type: %s", m.SequenceId, m.Type)

	}
	return nil
}

func (m *Minter) CalculateInflation(totalSupply sdk.Int, startTime time.Time, blockTime time.Time) sdk.Dec {
	if startTime.After(blockTime) {
		return sdk.ZeroDec()
	}
	switch m.Type {
	case NO_MINTING:
		return sdk.ZeroDec()
	case LINEAR_MINTING:
		return m.LinearMinting.calculateInfation(totalSupply, startTime, *m.EndTime)
	case EXPONENTIAL_STEP_MINTING:
		return m.ExponentialStepMinting.calculateInfation(totalSupply, startTime, m.EndTime, blockTime)
	default:
		return sdk.ZeroDec()
	}
}

type BySequenceId []*Minter

func (a BySequenceId) Len() int           { return len(a) }
func (a BySequenceId) Less(i, j int) bool { return a[i].SequenceId < a[j].SequenceId }
func (a BySequenceId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (m *LinearMinting) amountToMint(Minterstart time.Time, EndTime time.Time, blockTime time.Time) sdk.Dec {
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

func (m LinearMinting) validate(id uint32) error {
	if m.Amount.IsNegative() {
		return fmt.Errorf("minter sequence id: %d - LinearMinting amount cannot be less than 0", id)

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

	logger.Debug("PRMinterMint", "blockTime", blockTime, "now", now, "passedTime", passedTime, "epoch", epoch, "numOfPassedEpochs", numOfPassedEpochs,
		"Amount", m.Amount, "epochAmount", epochAmount, "amountToMint", amountToMint, "currentEpochStart", currentEpochStart,
		"currentEpochPassedTime", currentEpochPassedTime, "currentEpochAmount", currentEpochAmount)
	if numOfPassedEpochs > 0 {
		currentEpochAmount = currentEpochAmount.Mul(m.AmountMultiplier)
	}
	currentEpochAmountToMint := currentEpochAmount.MulInt64(int64(currentEpochPassedTime)).QuoInt64(epoch)
	logger.Debug("PRMinterMintCon", "AmountMultiplier", m.AmountMultiplier, "currentEpochAmount", currentEpochAmount, "currentEpochAmountToMint", currentEpochAmountToMint)
	return amountToMint.Add(currentEpochAmountToMint)

}

func (m ExponentialStepMinting) validate(id uint32) error {
	if m.Amount.IsNegative() {
		return fmt.Errorf("minter sequence id: %d - ExponentialStepMinting Amount cannot be less than 0", id)
	}
	if m.StepDuration <= 0 {
		return fmt.Errorf("minter sequence id: %d - ExponentialStepMinting StepDuration must be bigger than 0", id)
	}

	return nil
}

func (m *ExponentialStepMinting) calculateInfation(totalSupply sdk.Int, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
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
