package cfevesting

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

)
func GenerateVestingTypes(numberOfVestingTypes int, startId int) []*types.VestingType {
	return Generate10BasedVestingTypes(numberOfVestingTypes, 0, false, startId)
}

func Generate10BasedVestingTypes(numberOfVestingTypes int, amountOf10BasedVestingTypes int, a10BasedVestingTypesDelegationAllowe bool, startId int) []*types.VestingType {
	vestingTypes := []*types.VestingType{}

	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numberOfVestingTypes; i++ {
		if i < amountOf10BasedVestingTypes {
			vestingType := types.VestingType{
				Name:                 "test-vesting-type-" + strconv.Itoa(i+startId),
				LockupPeriod:         1000,
				VestingPeriod:        5000,
				TokenReleasingPeriod: 10,
				DelegationsAllowed:   a10BasedVestingTypesDelegationAllowe,
			}
			vestingTypes = append(vestingTypes, &vestingType)
		} else {
			vestingType := types.VestingType{
				Name:                 "test-vesting-type-" + strconv.Itoa(i+startId),
				LockupPeriod:         int64(rgen.Intn(100000)),
				VestingPeriod:        int64(rgen.Intn(100000)),
				TokenReleasingPeriod: int64(rgen.Intn(1000)),
				DelegationsAllowed:   rgen.Intn(2) == 1,
			}
			vestingTypes = append(vestingTypes, &vestingType)
		}
	}

	return vestingTypes
}
