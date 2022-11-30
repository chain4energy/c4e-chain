package cfevesting

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
)

func GenerateVestingTypes(numberOfVestingTypes int, startId int) []*types.VestingType {
	return Generate10BasedVestingTypes(numberOfVestingTypes, 0, startId)
}

func Generate10BasedVestingTypes(numberOfVestingTypes int, amountOf10BasedVestingTypes int, startId int) []*types.VestingType {
	vestingTypes := []*types.VestingType{}

	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numberOfVestingTypes; i++ {
		if i < amountOf10BasedVestingTypes {
			vestingType := types.VestingType{
				Name:          "test-vesting-type-" + strconv.Itoa(i+startId),
				LockupPeriod:  CreateDurationFromNumOfHours(1000),
				VestingPeriod: CreateDurationFromNumOfHours(5000),
				InitialBonus:  sdk.ZeroDec(),
			}
			vestingTypes = append(vestingTypes, &vestingType)
		} else {
			vestingType := types.VestingType{
				Name:          "test-vesting-type-" + strconv.Itoa(i+startId),
				LockupPeriod:  CreateDurationFromNumOfHours(int64(rgen.Intn(100000))),
				VestingPeriod: CreateDurationFromNumOfHours(int64(rgen.Intn(100000))),
				InitialBonus:  sdk.ZeroDec(),
			}
			vestingTypes = append(vestingTypes, &vestingType)
		}
	}

	return vestingTypes
}

func GenerateVestingTypesForAccountVestingPools(accountVestingPools []types.AccountVestingPools) []*types.VestingType {

	m := make(map[string]*types.VestingType)
	result := []*types.VestingType{}
	for _, av := range accountVestingPools {
		for _, v := range av.VestingPools {
			vt := GenerateVestingTypes(1, 1)[0]
			vt.Name = v.VestingType
			m[v.VestingType] = vt

		}
	}
	for _, vt := range m {
		result = append(result, vt)
	}

	return result
}
