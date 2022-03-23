package cfevesting

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"

)

func GenerateVestingTypes(numberOfVestingTypes int, startId int) []*types.VestingType {
	vestingTypes := []*types.VestingType{}

	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numberOfVestingTypes; i++ {
		vestingType := types.VestingType{
			Name:                 "test-vesting-type-" + strconv.Itoa(i+startId),
			LockupPeriod:         int64(rgen.Intn(100000)),
			VestingPeriod:        int64(rgen.Intn(100000)),
			TokenReleasingPeriod: int64(rgen.Intn(1000)),
			DelegationsAllowed:   rgen.Intn(2) == 1,
		}
		vestingTypes = append(vestingTypes, &vestingType)
	}

	return vestingTypes
}
