package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates the set of params

func (s State) Validate() error {
	if s.Burn && s.Account != nil {
		return fmt.Errorf("when burn is set to true account cannot exist")
	}
	if !s.Burn && s.Account == nil {
		return fmt.Errorf("when burn is set to false account must exist")
	}
	if err := s.IsNegative(); err != nil {
		return err
	}

	return nil
}

func (s State) IsNegative() error {
	for _, coinState := range s.Remains {
		if coinState.IsNegative() {
			return fmt.Errorf("\tnegative coin state %s in state %s", coinState, s.StateIdString())
		}
	}

	return nil
}

func StateSumIsInteger(states []State) (error, sdk.Coins) {
	statesSum := sdk.NewDecCoins()
	for _, state := range states {
		statesSum = statesSum.Add(state.Remains...)
	}

	remainsSum, change := statesSum.TruncateDecimal()
	if !change.IsZero() {
		return fmt.Errorf("\tthe sum of the states should be integer: sum: %v", statesSum), nil
	}

	return nil, remainsSum
}
