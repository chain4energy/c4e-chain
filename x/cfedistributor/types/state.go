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
			return fmt.Errorf("\tnegative coin state %s in state %s", coinState, s.stateIdString())
		}
	}

	return nil
}

func (s State) stateIdString() string {
	if s.Burn {
		return Burn
	} else if s.Account != nil && s.Account.Type == Main {
		return Main
	} else if s.Account != nil {
		return s.Account.Type + "-" + s.Account.Id
	} else {
		return UnknownAccount
	}
}

func StateSumIsInteger(states []State) (sdk.Coins, error) {
	statesSum := sdk.NewDecCoins()
	for _, state := range states {
		statesSum = statesSum.Add(state.Remains...)
	}

	remainsSum, change := statesSum.TruncateDecimal()
	if !change.IsZero() {
		return nil, fmt.Errorf("\tthe sum of the states should be integer: sum: %v", statesSum)
	}

	return remainsSum, nil
}

func (state State) GetStateKey() string {
	if state.Account != nil && state.Account.Id != "" && state.Account.Type != "" {
		return state.Account.GetAccountKey()
	} else {
		return BurnStateKey
	}
}
