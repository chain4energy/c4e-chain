package v120_test

import (
	v110 "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v110"
	v120 "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v120"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMigrationManyAccountVestingPoolsWithManyPools(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[0].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[1].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[2].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[3].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[4].String(), 10)
	MigrateV110ToV120(t, testUtil, ctx)
}

func TestMigrationNoAccountVestingPoolsAndNoVestingTypes(t *testing.T) {
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	MigrateV110ToV120(t, testUtil, ctx)
}

func TestMigrationManyAccountVestingPoolsWithNoPools(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[0].String(), 0)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[1].String(), 0)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[2].String(), 0)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[3].String(), 0)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[4].String(), 0)
	MigrateV110ToV120(t, testUtil, ctx)
}

func TestMigrationOneAccountVestingPoolsWithOnePool(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[0].String(), 1)
	MigrateV110ToV120(t, testUtil, ctx)

}

func TestMigrationAccountVestingPools(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[0].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[1].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[2].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[3].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[4].String(), 10)
	MigrateV110ToV120(t, testUtil, ctx)
}

func SetupOldAccountVestingPools(testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context, address string, numberOfVestingPools int) v110.AccountVestingPools {
	accountVestingPools := generateOneOldAccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPools, 1, 1)
	accountVestingPools.Address = address
	setOldAccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc, accountVestingPools)
	return accountVestingPools
}

func setOldAccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, accountVestingPools v110.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(storeKey), v110.AccountVestingPoolsKeyPrefix)
	av := cdc.MustMarshal(&accountVestingPools)
	store.Set([]byte(accountVestingPools.Address), av)
}

func MigrateV110ToV120(t *testing.T, testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context) {
	oldAccPools := getAllOldAccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc)
	err := v120.MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
	require.NoError(t, err)

	newAccPools := testUtil.GetC4eVestingKeeper().GetAllAccountVestingPools(ctx)

	require.EqualValues(t, len(oldAccPools), len(newAccPools))
	for i := 0; i < len(oldAccPools); i++ {
		require.EqualValues(t, oldAccPools[i].Address, newAccPools[i].Owner)
		require.EqualValues(t, len(oldAccPools[i].VestingPools), len(newAccPools[i].VestingPools))
		for j := 0; j < len(oldAccPools[i].VestingPools); j++ {
			oldVestingPool := oldAccPools[i].VestingPools[j]
			newVestingPool := newAccPools[i].VestingPools[j]
			require.EqualValues(t, oldVestingPool, newVestingPool)
		}
	}
}

func getAllOldAccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (list []v110.AccountVestingPools) {
	prefixStore := prefix.NewStore(ctx.KVStore(storeKey), v110.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v110.AccountVestingPools
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func generateOneOldAccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPoolsPerAccount int,
	accountId int, vestingStartId int) v110.AccountVestingPools {
	return *generateOldAccountVestingPoolsWithRandomVestingPools(1, numberOfVestingPoolsPerAccount, accountId, vestingStartId)[0]
}

func generateOldAccountVestingPoolsWithRandomVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int) []*v110.AccountVestingPools {
	return generateOldAccountVestingPools(numberOfAccounts, numberOfVestingPoolsPerAccount,
		accountStartId, vestingStartId, generateRandomOldVestingPool)
}

func generateOldAccountVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) v110.VestingPool) []*v110.AccountVestingPools {
	var accountVestingPoolsArr []*v110.AccountVestingPools
	accountsAddresses, _ := testcosmos.CreateAccounts(2*numberOfAccounts, 0)

	for i := 0; i < numberOfAccounts; i++ {
		accountVestingPools := v110.AccountVestingPools{}
		accountVestingPools.Address = accountsAddresses[i].String()

		var vestingPools []*v110.VestingPool
		for j := 0; j < numberOfVestingPoolsPerAccount; j++ {
			vesting := generateVesting(i+accountStartId, j+vestingStartId)
			vestingPools = append(vestingPools, &vesting)
		}
		accountVestingPools.VestingPools = vestingPools

		accountVestingPoolsArr = append(accountVestingPoolsArr, &accountVestingPools)
	}

	return accountVestingPoolsArr
}

func generateRandomOldVestingPool(accuntId int, vestingId int) v110.VestingPool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	vested := int(helpers.RandIntBetweenWith0(r, 1, 10000000))
	withdrawn := r.Intn(vested)
	sent := helpers.RandIntWith0(r, vested-withdrawn)

	return v110.VestingPool{
		Name:            "test-vesting-account-name" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		VestingType:     "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		LockStart:       testutils.CreateTimeFromNumOfHours(int64(r.Intn(100000))),
		LockEnd:         testutils.CreateTimeFromNumOfHours(int64(r.Intn(100000))),
		InitiallyLocked: sdk.NewInt(int64(vested)),
		Withdrawn:       sdk.NewInt(int64(withdrawn)),
		Sent:            sdk.NewInt(int64(sent)),
	}
}
