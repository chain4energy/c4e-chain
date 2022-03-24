package cfevesting

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GenerateOneAccountVestingsWithAddressWithRandomVestings(numberOfVestingsPerAccount int,
	accountId int, vestingStartId int) types.AccountVestings {
	return *GenerateAccountVestingsWithRandomVestings(1, numberOfVestingsPerAccount, accountId, vestingStartId)[0]
}

func GenerateAccountVestingsWithRandomVestings(numberOfAccounts int, numberOfVestingsPerAccount int,
	accountStartId int, vestingStartId int) []*types.AccountVestings {
	return generateAccountVestings(numberOfAccounts, numberOfVestingsPerAccount,
		accountStartId, vestingStartId, generateRandomVesting)
}

func GenerateOneAccountVestingsWithAddressWith10BasedVestings(numberOfVestingsPerAccount int,
	accountId int, vestingStartId int) types.AccountVestings {
	return *GenerateAccountVestingsWith10BasedVestings(1, numberOfVestingsPerAccount, accountId, vestingStartId)[0]
}

func GenerateAccountVestingsWith10BasedVestings(numberOfAccounts int, numberOfVestingsPerAccount int,
	accountStartId int, vestingStartId int) []*types.AccountVestings {
	return generateAccountVestings(numberOfAccounts, numberOfVestingsPerAccount,
		accountStartId, vestingStartId, generate10BasedVesting)
}

func generateAccountVestings(numberOfAccounts int, numberOfVestingsPerAccount int,
	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) types.Vesting) []*types.AccountVestings {
	accountVestingsArr := []*types.AccountVestings{}
	for i := 0; i < numberOfAccounts; i++ {
		accountVestings := types.AccountVestings{}
		accountVestings.Address = "test-vesting-account-addr-" + strconv.Itoa(i+accountStartId)
		accountVestings.DelegableAddress = "test-vesting-account-del-addr-" + strconv.Itoa(i+accountStartId)

		vestings := []*types.Vesting{}
		for j := 0; j < numberOfVestingsPerAccount; j++ {
			vesting := generateVesting(i+accountStartId, j+vestingStartId)
			vestings = append(vestings, &vesting)
		}
		accountVestings.Vestings = vestings

		accountVestingsArr = append(accountVestingsArr, &accountVestings)
	}

	return accountVestingsArr
}

func generateRandomVesting(accuntId int, vestingId int) types.Vesting {
	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return types.Vesting{
		Id:                        int32(vestingId),
		VestingType:               "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(accuntId),
		VestingStart:         int64(rgen.Intn(100000)),
		LockEnd:              int64(rgen.Intn(100000)),
		VestingEnd:           int64(rgen.Intn(100000)),
		Vested:                    sdk.NewInt(int64(rgen.Intn(10000000))),
		ReleasePeriod:      int64(rgen.Intn(1000)),
		DelegationAllowed:         rgen.Intn(2) == 1,
		Withdrawn:                 sdk.NewInt(int64(rgen.Intn(10000000))),
		Sent:                      sdk.NewInt(int64(rgen.Intn(10000000))),
		LastModification:     int64(rgen.Intn(100000)),
		LastModificationVested:    sdk.NewInt(int64(rgen.Intn(10000000))),
		LastModificationWithdrawn: sdk.NewInt(int64(rgen.Intn(10000000))),
	}
}

func generate10BasedVesting(accuntId int, vestingId int) types.Vesting {
	return types.Vesting{
		Id:                        int32(vestingId),
		VestingType:               "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(accuntId),
		VestingStart:         1000,
		LockEnd:              10000,
		VestingEnd:           110000,
		Vested:                    sdk.NewInt(1000000),
		ReleasePeriod:      10,
		DelegationAllowed:         true,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:     1000,
		LastModificationVested:    sdk.NewInt(1000000),
		LastModificationWithdrawn: sdk.ZeroInt(),
	}
}

func EqualAccountVestings(t require.TestingT, expected []*types.AccountVestings, actual []*types.AccountVestings) {
	for _, accVest := range expected {
		found := false
		for _, accVestExp := range actual {
			if accVest.Address == accVestExp.Address {
				require.EqualValues(t, accVest, accVestExp)
				found = true
			}
		}
		require.True(t, found, "not found: "+accVest.Address)

	}
}

func ToAccountVestingsPointersArray(src []types.AccountVestings) []*types.AccountVestings {
	result := []*types.AccountVestings{}
	for i := 0; i < len(src); i++ {
		result = append(result, &src[i])
	}
	return result
}

func GetExpectedWithdrawableForVesting(vesting types.Vesting, currentHeight int64) sdk.Int {
	unlockingStartHeight := vesting.LockEnd
	if vesting.VestingStart > unlockingStartHeight {
		unlockingStartHeight = vesting.VestingStart
	}
	if vesting.LastModification > unlockingStartHeight {
		unlockingStartHeight = vesting.LastModification
	}
	expected := GetExpectedWithdrawable(unlockingStartHeight, vesting.VestingEnd, vesting.ReleasePeriod, currentHeight, vesting.LastModificationVested)
	return expected.Sub(vesting.LastModificationWithdrawn)
}

func GetExpectedWithdrawable(unlockingStartHeight int64, vestingEndHeight int64, heightPeriod int64, currentHeight int64, amount sdk.Int) sdk.Int {
	if currentHeight >= vestingEndHeight {
		return amount
	}
	if currentHeight < unlockingStartHeight {
		return sdk.ZeroInt()
	}
	numOfAllPeriodsF := float64(vestingEndHeight-unlockingStartHeight) / float64(heightPeriod)
	numOfAllPeriods := int64(math.Ceil(numOfAllPeriodsF))

	numOfPeriodsF := float64(currentHeight-unlockingStartHeight) / float64(heightPeriod)
	numOfPeriods := int64(math.Floor(numOfPeriodsF))

	amountDec := sdk.NewDecFromInt(amount)

	resultDec := amountDec.MulInt64(numOfPeriods).QuoInt64(numOfAllPeriods)
	return resultDec.TruncateInt()
}
