package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (av AccountVestingPools) Validate() error {
	vs := av.VestingPools
	_, err := sdk.AccAddressFromBech32(av.Address)
	if err != nil {
		return fmt.Errorf("account vesting pools address: %s: %s", av.Address, err.Error())
	}
	for _, v := range vs {
		if err = v.Validate(av.Address); err != nil {
			return err
		}
		err = av.checkDuplications(vs, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (av AccountVestingPools) checkDuplications(vs []*VestingPool, v *VestingPool) error {
	numOfNames := 0
	for _, vCheck := range vs {
		if v.Name == vCheck.Name {
			numOfNames++
		}
		if numOfNames > 1 {
			return fmt.Errorf("vesting pool with name: %s defined more than once for account: %s", v.Name, av.Address)
		}
	}

	return nil
}

func (av AccountVestingPools) ValidateAgainstVestingTypes(vestingTypes []GenesisVestingType) error {
	vs := av.VestingPools
	for _, v := range vs {
		found := false
		for _, vtCheck := range vestingTypes {
			if v.VestingType == vtCheck.Name {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("vesting pool with name: %s defined for account: %s - vesting type not found: %s", v.Name, av.Address, v.VestingType)
		}
	}
	return nil
}

func (m *VestingPool) GetCurrentlyLocked() sdk.Int {
	return m.InitiallyLocked.Sub(m.Sent).Sub(m.Withdrawn)
}

func (m *VestingPool) Validate(accountAdd string) error {
	if len(m.Name) == 0 {
		return fmt.Errorf("vesting pool defined for account: %s has no name", accountAdd)
	}
	return nil
}

