package types

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/yaml.v2"
	"sort"
	"time"
)

const year = time.Hour * 24 * 365

func (params Params) Validate() error {
	if err := params.ValidateParamsMintDenom(); err != nil {
		return err
	}
	if err := params.ValidateParamsMinters(); err != nil {
		return err
	}
	return nil
}

func (params Params) ValidateParamsMintDenom() error {
	if len(params.MintDenom) == 0 {
		return fmt.Errorf("denom cannot be empty")
	}
	return nil
}

func (params Params) ValidateParamsMinters() error {
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

func (params Params) validateMinterOrderingId(minter *Minter, id uint32) (uint32, error) {
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

func (params Params) validateEndTimeExistance(minter *Minter, sequenceId int, lastPos int) error {
	if sequenceId == lastPos && minter.EndTime != nil {
		return fmt.Errorf("last minter cannot have EndTime set, but is set to %s", minter.EndTime)
	}
	if sequenceId < lastPos && minter.EndTime == nil {
		return fmt.Errorf("only last minter can have EndTime empty")
	}
	return nil
}

func (params Params) validateMintersEndTimeValue(minter *Minter, sequenceId int, lastPos int) error {
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

func (params Params) ContainsMinter(sequenceId uint32) bool {
	for _, minter := range params.Minters {
		if sequenceId == minter.SequenceId {
			return true
		}
	}
	return false
}

func (m *Minter) GetMinterConfig() (MinterConfigI, error) {
	if m.Config == nil {
		return nil, fmt.Errorf("minter config is nil")
	}
	minterConfigI, ok := m.Config.GetCachedValue().(MinterConfigI)
	if !ok {
		return nil, fmt.Errorf("expected %T, got %T", (MinterConfigI)(nil), m.Config.GetCachedValue())
	}
	return minterConfigI, nil
}

func (m *Minter) validate() error {
	minterConfig, err := m.GetMinterConfig()
	if err != nil {
		return err
	}

	_, ok := m.Config.GetCachedValue().(*LinearMinting)
	if ok {
		if m.EndTime == nil {
			return fmt.Errorf("for LinearMinting EndTime must be set")
		}
	}
	if err := minterConfig.Validate(); err != nil {
		return fmt.Errorf("minter config validation error: %w", err)
	}

	return nil
}

type MinterJSON struct {
	SequenceId uint32     `json:"sequence_id"`
	EndTime    *time.Time `json:"end_time"`

	// custom fields based on concrete vesting type which can be omitted
	Config string `json:"config,omitempty"`
	Type   string `json:"type,omitempty"`
}

func (m *Minter) GetMinterJSON() MinterJSON {
	if m == nil {
		return MinterJSON{}
	}
	minterConfig, _ := m.GetMinterConfig()
	var config string
	if minterConfig != nil {
		config = minterConfig.String()
	}
	return MinterJSON{
		SequenceId: m.SequenceId,
		EndTime:    m.EndTime,
		Type:       m.Config.GetTypeUrl(),
		Config:     config,
	}
}

func (m *LinearMinting) Validate() error {
	if m == nil {
		return fmt.Errorf("LinearMinting must be set")
	}
	if m.Amount.IsNil() {
		return fmt.Errorf("amount cannot be nil")
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("amount cannot be less than 0")
	}
	return nil
}

func (m *ExponentialStepMinting) Validate() error {
	if m == nil {
		return fmt.Errorf("ExponentialStepMintingType must be set")
	}
	if m.Amount.IsNil() {
		return fmt.Errorf("amount cannot be nil")
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("amount cannot be less than 0")
	}
	if !m.Amount.IsPositive() {
		return fmt.Errorf("amount must be positive")
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
	if m.RemainderFromPreviousMinter.IsNil() {
		return fmt.Errorf("minter state validation error: remainderFromPreviousMinter cannot be nil")
	}
	if m.RemainderFromPreviousMinter.IsNegative() {
		return fmt.Errorf("minter state validation error: remainderFromPreviousMinter cannot be less than 0")
	}
	if m.RemainderToMint.IsNil() {
		return fmt.Errorf("minter state validation error: remainderToMint cannot be nil")
	}
	if m.RemainderToMint.IsNegative() {
		return fmt.Errorf("minter state validation error: remainderToMint cannot be less than 0")
	}
	return nil
}

func (m *NoMinting) Validate() error {
	return nil
}

type BySequenceId []*Minter

func (a BySequenceId) Len() int           { return len(a) }
func (a BySequenceId) Less(i, j int) bool { return a[i].SequenceId < a[j].SequenceId }
func (a BySequenceId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (m *Minter) CalculateInflation(totalSupply math.Int, startTime time.Time, blockTime time.Time) sdk.Dec {
	if startTime.After(blockTime) {
		return sdk.ZeroDec()
	}
	minterConfig, _ := m.GetMinterConfig()
	return minterConfig.CalculateInflation(totalSupply, startTime, m.EndTime, blockTime)
}

func (m *LinearMinting) CalculateInflation(totalSupply math.Int, minterStart time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	if totalSupply.LTE(math.ZeroInt()) {
		return sdk.ZeroDec()
	}

	periodDuration := endTime.Sub(minterStart)
	mintedYearly := sdk.NewDecFromInt(m.Amount).MulInt64(int64(year)).QuoInt64(int64(periodDuration))
	return mintedYearly.QuoInt(totalSupply)
}

func (m *ExponentialStepMinting) CalculateInflation(totalSupply math.Int, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	if totalSupply.LTE(math.ZeroInt()) {
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

func (m *NoMinting) CalculateInflation(totalSupply math.Int, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	return sdk.ZeroDec()
}

func (m *Minter) AmountToMint(logger log.Logger, startTime time.Time, blockTime time.Time) sdk.Dec {
	minterConfig, _ := m.GetMinterConfig()
	return minterConfig.AmountToMint(logger, startTime, m.EndTime, blockTime)
}

func (m *LinearMinting) AmountToMint(logger log.Logger, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	if blockTime.After(*endTime) {
		return sdk.NewDecFromInt(m.Amount)
	}
	if blockTime.Before(startTime) {
		return sdk.ZeroDec()
	}
	amount := sdk.NewDecFromInt(m.Amount)

	passedTime := blockTime.UnixMilli() - startTime.UnixMilli()
	period := endTime.UnixMilli() - startTime.UnixMilli()

	return amount.MulInt64(passedTime).QuoInt64(period)
}

func (m *ExponentialStepMinting) AmountToMint(logger log.Logger, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	now := blockTime
	if endTime != nil && blockTime.After(*endTime) {
		now = *endTime
	}
	passedTime := int64(now.Sub(startTime))
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
	currentEpochStart := startTime.Add(time.Duration(numOfPassedEpochs * epoch))
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

func (m *NoMinting) AmountToMint(logger log.Logger, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	return sdk.ZeroDec()
}

func (acc *Minter) String() string {
	out, _ := yaml.Marshal(acc)
	return string(out)
}
