package cfevesting

import (
	// "math"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func AssertAccountVestingPools(t require.TestingT, expected types.AccountVestingPools, actual types.AccountVestingPools) {

	numOfFields := reflect.TypeOf(types.AccountVestingPools{}).NumField()
	j := 0
	require.EqualValues(t, len(expected.VestingPools), len(actual.VestingPools))
	j++
	require.EqualValues(t, expected.Address, actual.Address)
	j++
	require.EqualValues(t, numOfFields, j)

	numOfFields = reflect.TypeOf(types.VestingPool{}).NumField()
	for i, expectedVesting := range expected.VestingPools {
		actualVesting := actual.VestingPools[i]
		j := 0
		require.EqualValues(t, expectedVesting.Name, actualVesting.Name)
		j++
		require.EqualValues(t, expectedVesting.VestingType, actualVesting.VestingType)
		j++
		require.EqualValues(t, true, expectedVesting.LockStart.Equal(actualVesting.LockStart))
		j++
		require.EqualValues(t, true, expectedVesting.LockEnd.Equal(actualVesting.LockEnd))
		j++
		require.EqualValues(t, expectedVesting.InitiallyLocked, actualVesting.InitiallyLocked)
		j++
		require.EqualValues(t, expectedVesting.Withdrawn, actualVesting.Withdrawn)
		j++
		require.EqualValues(t, expectedVesting.Sent, actualVesting.Sent)
		j++
		require.EqualValues(t, numOfFields, j)

	}

}

func AssertAccountVestingPoolsArrays(t require.TestingT, expected []*types.AccountVestingPools, actual []*types.AccountVestingPools) {
	require.EqualValues(t, len(expected), len(actual))

	for _, accVest := range expected {
		found := false
		for _, accVestExp := range actual {
			if accVest.Address == accVestExp.Address {
				AssertAccountVestingPools(t, *accVest, *accVestExp)
				found = true
			}
		}
		require.True(t, found, "not found: "+accVest.Address)

	}
}

func GenerateOneAccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPoolsPerAccount int,
	accountId int, vestingStartId int) types.AccountVestingPools {
	return *GenerateAccountVestingPoolsWithRandomVestingPools(1, numberOfVestingPoolsPerAccount, accountId, vestingStartId)[0]
}

func GenerateAccountVestingPoolsWithRandomVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int) []*types.AccountVestingPools {
	return generateAccountVestingPools(numberOfAccounts, numberOfVestingPoolsPerAccount,
		accountStartId, vestingStartId, generateRandomVestingPool)
}

func GenerateOneAccountVestingPoolsWithAddressWith10BasedVestingPools(numberOfVestingPoolsPerAccount int,
	accountId int, vestingStartId int) types.AccountVestingPools {
	return *GenerateAccountVestingPoolsWith10BasedVestingPools(1, numberOfVestingPoolsPerAccount, accountId, vestingStartId)[0]
}

func GenerateAccountVestingPoolsWith10BasedVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int) []*types.AccountVestingPools {
	return generateAccountVestingPools(numberOfAccounts, numberOfVestingPoolsPerAccount,
		accountStartId, vestingStartId, generate10BasedVestingPool)
}

func generateAccountVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) types.VestingPool) []*types.AccountVestingPools {
	accountVestingPoolsArr := []*types.AccountVestingPools{}
	accountsAddresses, _ := testcosmos.CreateAccounts(2*numberOfAccounts, 0)

	for i := 0; i < numberOfAccounts; i++ {
		accountVestingPools := types.AccountVestingPools{}
		accountVestingPools.Address = "test-vesting-account-addr-" + strconv.Itoa(i+accountStartId)

		accountVestingPools.Address = accountsAddresses[i].String()

		var vestingPools []*types.VestingPool
		for j := 0; j < numberOfVestingPoolsPerAccount; j++ {
			vesting := generateVesting(i+accountStartId, j+vestingStartId)
			vestingPools = append(vestingPools, &vesting)
		}
		accountVestingPools.VestingPools = vestingPools

		accountVestingPoolsArr = append(accountVestingPoolsArr, &accountVestingPools)
	}

	return accountVestingPoolsArr
}

func generateRandomVestingPool(accuntId int, vestingId int) types.VestingPool {
	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	initiallyLocked := rgen.Intn(10000000)
	withdrawn := rgen.Intn(initiallyLocked)
	sent := rgen.Intn(initiallyLocked - withdrawn)
	return types.VestingPool{
		Name:            "test-vesting-account-name" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		VestingType:     "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		LockStart:       CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		LockEnd:         CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		InitiallyLocked: sdk.NewInt(int64(initiallyLocked)),
		Withdrawn:       sdk.NewInt(int64(withdrawn)),
		Sent:            sdk.NewInt(int64(sent)),
	}
}

func generate10BasedVestingPool(accuntId int, vestingId int) types.VestingPool {
	return types.VestingPool{
		Name:            "test-vesting-account-name" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		VestingType:     "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		LockStart:       CreateTimeFromNumOfHours(1000),
		LockEnd:         CreateTimeFromNumOfHours(110000),
		InitiallyLocked: sdk.NewInt(1000000),
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
	}
}

func ToAccountVestingPoolsPointersArray(src []types.AccountVestingPools) []*types.AccountVestingPools {
	result := []*types.AccountVestingPools{}
	for i := 0; i < len(src); i++ {
		result = append(result, &src[i])
	}
	return result
}

func GetExpectedWithdrawableForVesting(vestingPool types.VestingPool, current time.Time) sdk.Int {
	result := GetExpectedWithdrawable(vestingPool.LockEnd, current, vestingPool.InitiallyLocked.Sub(vestingPool.Sent).Sub(vestingPool.Withdrawn))
	if result.LT(sdk.ZeroInt()) {
		return sdk.ZeroInt()
	}
	return result
}

func GetExpectedWithdrawable(lockEnd time.Time, current time.Time, amount sdk.Int) sdk.Int {
	if current.Equal(lockEnd) || current.After(lockEnd) {
		return amount
	}
	return sdk.ZeroInt()
}

func CreateTimeFromNumOfHours(numOfHours int64) time.Time {
	return testenv.TestEnvTime.Add(time.Hour * time.Duration(numOfHours))
}

func CreateDurationFromNumOfHours(numOfHours int64) time.Duration {
	return time.Hour * time.Duration(numOfHours)
}

func GetVestingPoolByName(vps []*types.VestingPool, name string) (vp *types.VestingPool, found bool) {
	for _, vPool := range vps {
		if vPool.Name == name {
			return vPool, true
		}
	}
	return nil, false
}
