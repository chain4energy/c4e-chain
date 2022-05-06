package types

import (
	// this line is used by starport scaffolding # genesis/types/import
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	// TODO validate accouns vestings
	// testmapIndexMap := make(map[string]struct{})

	// for _, elem := range gs.TestmapList {
	// 	index := string(TestmapKey(elem.Index))
	// 	if _, ok := testmapIndexMap[index]; ok {
	// 		return fmt.Errorf("duplicated index for testmap")
	// 	}
	// 	testmapIndexMap[index] = struct{}{}
	// }

	err := gs.validateVestingTypes()
	if err != nil {
		return err
	}
	err = gs.validateAccountsVestings()
	if err != nil {
		return err
	}
	return gs.Params.Validate()
}

func (gs GenesisState) validateVestingTypes() error {
	vts := gs.VestingTypes
	for _, vt := range vts {
		numOfNames := 0
		for _, vtCheck := range vts {
			if vt.Name == vtCheck.Name {
				numOfNames++
			}
			if numOfNames > 1 {
				return fmt.Errorf("vesting type with name: %s defined more than once", vt.Name)
			}
		}
	}
	return nil
}

func (gs GenesisState) validateAccountsVestings() error {
	avts := gs.AccountVestingsList.Vestings
	vts := gs.VestingTypes
	for _, avt := range avts {
		err := avt.Validate()
		if err != nil {
			return err
		}
		numOfAddress := 0

		for _, avtCheck := range avts {
			if avt.Address == avtCheck.Address {
				numOfAddress++
			}
			if numOfAddress > 1 {
				return fmt.Errorf("account vestings with address: %s defined more than once", avt.Address)
			}
		}
		err = avt.ValidateAgainstVestingTypes(vts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (av AccountVestings) Validate() error {
	vs := av.VestingPools
	_, err := sdk.AccAddressFromBech32(av.Address)
	if err != nil {
		return fmt.Errorf("account vestings address: %s: %s", av.Address, err.Error())
	}
	for _, v := range vs {
		numOfIds := 0
		numOfNames := 0
		for _, vCheck := range vs {
			if v.Id == vCheck.Id {
				numOfIds++
			}
			if numOfIds > 1 {
				return fmt.Errorf("vesting with id: %d defined more than once for account: %s", v.Id, av.Address)
			}

			if v.Name == vCheck.Name {
				numOfNames++
			}
			if numOfNames > 1 {
				return fmt.Errorf("vesting with name: %s defined more than once for account: %s", v.Name, av.Address)
			}
		}
	}
	return nil
}

func (av AccountVestings) ValidateAgainstVestingTypes(vestingTypes []GenesisVestingType) error {
	vs := av.VestingPools
	for _, v := range vs {
		found := false
		for _, vtCheck := range vestingTypes {
			if v.VestingType == vtCheck.Name {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("vesting with id: %d defined for account: %s - vesting type not found: %s", v.Id, av.Address, v.VestingType)
		}
	}
	return nil
}
