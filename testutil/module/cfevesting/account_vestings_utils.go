package cfevesting

import (
	// "math"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func AssertAccountVestings(t *testing.T, expected types.AccountVestings, actual types.AccountVestings) {

	numOfFields := reflect.TypeOf(types.AccountVestings{}).NumField()
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
		require.EqualValues(t, expectedVesting.Id, actualVesting.Id)
		j++
		require.EqualValues(t, expectedVesting.Name, actualVesting.Name)
		j++
		require.EqualValues(t, expectedVesting.VestingType, actualVesting.VestingType)
		j++
		require.EqualValues(t, true, expectedVesting.LockStart.Equal(actualVesting.LockStart))
		j++
		require.EqualValues(t, true, expectedVesting.LockEnd.Equal(actualVesting.LockEnd))
		j++
		require.EqualValues(t, expectedVesting.Vested, actualVesting.Vested)
		j++
		require.EqualValues(t, expectedVesting.Withdrawn, actualVesting.Withdrawn)
		j++
		require.EqualValues(t, expectedVesting.Sent, actualVesting.Sent)
		j++
		require.EqualValues(t, true, expectedVesting.LastModification.Equal(actualVesting.LastModification))
		j++
		require.EqualValues(t, expectedVesting.LastModificationVested, actualVesting.LastModificationVested)
		j++
		require.EqualValues(t, expectedVesting.LastModificationWithdrawn, actualVesting.LastModificationWithdrawn)
		j++
		require.EqualValues(t, numOfFields, j)

	}

}

func AssertAccountVestingsArrays(t *testing.T, expected []*types.AccountVestings, actual []*types.AccountVestings) {
	require.EqualValues(t, len(expected), len(actual))

	for _, accVest := range expected {
		found := false
		for _, accVestExp := range actual {
			if accVest.Address == accVestExp.Address {
				AssertAccountVestings(t, *accVest, *accVestExp)
				found = true
			}
		}
		require.True(t, found, "not found: "+accVest.Address)

	}
}

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
	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) types.VestingPool) []*types.AccountVestings {
	accountVestingsArr := []*types.AccountVestings{}
	accountsAddresses, _ := commontestutils.CreateAccounts(2*numberOfAccounts, 0)

	for i := 0; i < numberOfAccounts; i++ {
		accountVestings := types.AccountVestings{}
		accountVestings.Address = "test-vesting-account-addr-" + strconv.Itoa(i+accountStartId)

		accountVestings.Address = accountsAddresses[i].String()

		vestings := []*types.VestingPool{}
		for j := 0; j < numberOfVestingsPerAccount; j++ {
			vesting := generateVesting(i+accountStartId, j+vestingStartId)
			vestings = append(vestings, &vesting)
		}
		accountVestings.VestingPools = vestings

		accountVestingsArr = append(accountVestingsArr, &accountVestings)
	}

	return accountVestingsArr
}

func generateRandomVesting(accuntId int, vestingId int) types.VestingPool {
	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	lastModificationVested := rgen.Intn(10000000)
	lastModificationWithdrawn := rgen.Intn(lastModificationVested)
	return types.VestingPool{
		Id:                        int32(vestingId),
		Name:                      "test-vesting-account-name" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		VestingType:               "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		LockStart:                 CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		LockEnd:                   CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		Vested:                    sdk.NewInt(int64(rgen.Intn(10000000))),
		Withdrawn:                 sdk.NewInt(int64(rgen.Intn(10000000))),
		Sent:                      sdk.NewInt(int64(rgen.Intn(10000000))),
		LastModification:          CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
		LastModificationVested:    sdk.NewInt(int64(lastModificationVested)),
		LastModificationWithdrawn: sdk.NewInt(int64(lastModificationWithdrawn)),
	}
}

func generate10BasedVesting(accuntId int, vestingId int) types.VestingPool {
	return types.VestingPool{
		Id:                        int32(vestingId),
		Name:                      "test-vesting-account-name" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		VestingType:               "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		LockStart:                 CreateTimeFromNumOfHours(1000),
		LockEnd:                   CreateTimeFromNumOfHours(110000),
		Vested:                    sdk.NewInt(1000000),
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          CreateTimeFromNumOfHours(1000),
		LastModificationVested:    sdk.NewInt(1000000),
		LastModificationWithdrawn: sdk.ZeroInt(),
	}
}

func ToAccountVestingsPointersArray(src []types.AccountVestings) []*types.AccountVestings {
	result := []*types.AccountVestings{}
	for i := 0; i < len(src); i++ {
		result = append(result, &src[i])
	}
	return result
}

func GetExpectedWithdrawableForVesting(vesting types.VestingPool, current time.Time) sdk.Int {
	unlockingStart := vesting.LockEnd
	if vesting.LockStart.After(unlockingStart) {
		unlockingStart = vesting.LockStart
	}
	if vesting.LastModification.After(unlockingStart) {
		unlockingStart = vesting.LastModification
	}
	expected := GetExpectedWithdrawable(unlockingStart, vesting.LockEnd, current, vesting.LastModificationVested)
	result := expected.Sub(vesting.LastModificationWithdrawn)
	if result.LT(sdk.ZeroInt()) {
		return sdk.ZeroInt()
	}
	return result
}

func GetExpectedWithdrawable(unlockingStart time.Time, vestingEnd time.Time, current time.Time, amount sdk.Int) sdk.Int {
	if current.Equal(vestingEnd) || current.After(vestingEnd) {
		return amount
	}
	return sdk.ZeroInt()
}

func CreateTimeFromNumOfHours(numOfHours int64) time.Time {
	return commontestutils.TestEnvTime.Add(time.Hour * time.Duration(numOfHours))
}

func CreateDurationFromNumOfHours(numOfHours int64) time.Duration {
	return time.Hour * time.Duration(numOfHours)
}
