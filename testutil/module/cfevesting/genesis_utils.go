package cfevesting

import (
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
)

func GenerateGenesisVestingTypes(numberOfVestingTypes int, startId int) []types.GenesisVestingType {
	vts := GenerateVestingTypes(numberOfVestingTypes, startId)
	result := []types.GenesisVestingType{}
	for _, vt := range vts {

		gvt := types.GenesisVestingType{
			Name:              vt.Name,
			LockupPeriod:      vt.LockupPeriod.Nanoseconds() / int64(time.Hour),
			LockupPeriodUnit:  types.Day,
			VestingPeriod:     vt.VestingPeriod.Nanoseconds() / int64(time.Hour),
			VestingPeriodUnit: types.Day,
		}
		result = append(result, gvt)
	}
	return result
}

func GenerateGenesisVestingTypesForAccounVestingPools(accountVestingPools []*types.AccountVestingPools) []types.GenesisVestingType {
	vt := GenerateVestingTypes(1, 1)[0]
	m := make(map[string]types.GenesisVestingType)
	result := []types.GenesisVestingType{}
	for _, av := range accountVestingPools {
		for _, v := range av.VestingPools {
			gvt := types.GenesisVestingType{
				Name:              v.VestingType,
				LockupPeriod:      vt.LockupPeriod.Nanoseconds() / int64(time.Hour),
				LockupPeriodUnit:  types.Day,
				VestingPeriod:     vt.VestingPeriod.Nanoseconds() / int64(time.Hour),
				VestingPeriodUnit: types.Day,
			}
			m[v.VestingType] = gvt

		}
	}
	for _, gvt := range m {
		result = append(result, gvt)
	}

	return result
}
