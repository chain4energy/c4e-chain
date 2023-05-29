package v3_test

import (
	"cosmossdk.io/math"
	"encoding/binary"
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v2"
	"github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v3"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
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

func TestMigrationVestingAccountTraces(t *testing.T) {
	oldTraces := generateOldVestingAccountTraces(10)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetOldVestingAccountTraceCount(ctx, testUtil.StoreKey, testUtil.Cdc, uint64(len(oldTraces)))
	for _, oldTrace := range oldTraces {
		SetOldVestingTraceAccount(ctx, testUtil.StoreKey, testUtil.Cdc, oldTrace)
	}

	MigrateV110ToV120(t, testUtil, ctx)
}

func SetupOldAccountVestingPools(testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context, address string, numberOfVestingPools int) v2.AccountVestingPools {
	accountVestingPools := generateOneOldAccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPools, 1, 1)
	accountVestingPools.Address = address
	setOldAccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc, accountVestingPools)
	return accountVestingPools
}

func setOldAccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, accountVestingPools v2.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(storeKey), v2.AccountVestingPoolsKeyPrefix)
	av := cdc.MustMarshal(&accountVestingPools)
	store.Set([]byte(accountVestingPools.Address), av)
}

func MigrateV110ToV120(t *testing.T, testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context) {
	oldAccPools := getAllOldAccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc)
	oldVestingAccountTraces := GetAllOldVestingAccountTraces(ctx, testUtil.StoreKey, testUtil.Cdc)
	oldVestingAccountTracesCount := GetOldVestingAccountTraceCount(ctx, testUtil.StoreKey, testUtil.Cdc)
	err := v3.MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
	require.NoError(t, err)

	newAccPools := testUtil.GetC4eVestingKeeper().GetAllAccountVestingPools(ctx)

	require.EqualValues(t, len(oldAccPools), len(newAccPools))
	for i := 0; i < len(oldAccPools); i++ {
		require.EqualValues(t, oldAccPools[i].Address, newAccPools[i].Owner)
		require.EqualValues(t, len(oldAccPools[i].VestingPools), len(newAccPools[i].VestingPools))
		for j := 0; j < len(oldAccPools[i].VestingPools); j++ {
			oldVestingPool := oldAccPools[i].VestingPools[j]
			expectedVestingPool := types.VestingPool{
				Name:            oldVestingPool.Name,
				VestingType:     oldVestingPool.VestingType,
				LockStart:       oldVestingPool.LockStart,
				LockEnd:         oldVestingPool.LockEnd,
				InitiallyLocked: oldVestingPool.InitiallyLocked,
				Withdrawn:       oldVestingPool.Withdrawn,
				Sent:            oldVestingPool.Sent,
				GenesisPool:     false,
			}
			newVestingPool := newAccPools[i].VestingPools[j]
			require.EqualValues(t, &expectedVestingPool, newVestingPool)
		}
	}

	expected := []types.VestingAccountTrace{}
	for _, oldVestingAccountTrace := range oldVestingAccountTraces {
		expected = append(expected,
			types.VestingAccountTrace{
				Id:                 oldVestingAccountTrace.Id,
				Address:            oldVestingAccountTrace.Address,
				Genesis:            false,
				FromGenesisPool:    false,
				FromGenesisAccount: false,
			},
		)
	}

	require.Equal(t, oldVestingAccountTracesCount, testUtil.GetC4eVestingKeeper().GetVestingAccountTraceCount(ctx))
	require.ElementsMatch(t, expected, testUtil.GetC4eVestingKeeper().GetAllVestingAccountTrace(ctx))

}

func getAllOldAccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (list []v2.AccountVestingPools) {
	prefixStore := prefix.NewStore(ctx.KVStore(storeKey), v2.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v2.AccountVestingPools
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func generateOneOldAccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPoolsPerAccount int,
	accountId int, vestingStartId int) v2.AccountVestingPools {
	return *generateOldAccountVestingPoolsWithRandomVestingPools(1, numberOfVestingPoolsPerAccount, accountId, vestingStartId)[0]
}

func generateOldAccountVestingPoolsWithRandomVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int) []*v2.AccountVestingPools {
	return generateOldAccountVestingPools(numberOfAccounts, numberOfVestingPoolsPerAccount,
		accountStartId, vestingStartId, generateRandomOldVestingPool)
}

func generateOldAccountVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) v2.VestingPool) []*v2.AccountVestingPools {
	var accountVestingPoolsArr []*v2.AccountVestingPools
	accountsAddresses, _ := testcosmos.CreateAccounts(2*numberOfAccounts, 0)

	for i := 0; i < numberOfAccounts; i++ {
		accountVestingPools := v2.AccountVestingPools{}
		accountVestingPools.Address = accountsAddresses[i].String()

		var vestingPools []*v2.VestingPool
		for j := 0; j < numberOfVestingPoolsPerAccount; j++ {
			vesting := generateVesting(i+accountStartId, j+vestingStartId)
			vestingPools = append(vestingPools, &vesting)
		}
		accountVestingPools.VestingPools = vestingPools

		accountVestingPoolsArr = append(accountVestingPoolsArr, &accountVestingPools)
	}

	return accountVestingPoolsArr
}

func generateRandomOldVestingPool(accuntId int, vestingId int) v2.VestingPool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	vested := int(helpers.RandIntBetweenWith0(r, 1, 10000000))
	withdrawn := r.Intn(vested)
	sent := helpers.RandIntWith0(r, vested-withdrawn)

	return v2.VestingPool{
		Name:            "test-vesting-account-name" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		VestingType:     "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		LockStart:       testutils.CreateTimeFromNumOfHours(int64(r.Intn(100000))),
		LockEnd:         testutils.CreateTimeFromNumOfHours(int64(r.Intn(100000))),
		InitiallyLocked: math.NewInt(int64(vested)),
		Withdrawn:       math.NewInt(int64(withdrawn)),
		Sent:            math.NewInt(int64(sent)),
	}
}

func generateOldVestingAccountTraces(amount int) []v2.VestingAccount {
	var traces []v2.VestingAccount
	for i := 0; i < amount; i++ {
		traces = append(traces, v2.VestingAccount{Id: uint64(i), Address: fmt.Sprintf("Address-%d", i)})
	}
	return traces
}

func SetOldVestingAccountTraceCount(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, count uint64) {
	store := prefix.NewStore(ctx.KVStore(storeKey), []byte{})
	byteKey := []byte(v2.VestingAccountCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func SetOldVestingTraceAccount(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, vestingAccount v2.VestingAccount) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(v2.VestingAccountKey))
	b := cdc.MustMarshal(&vestingAccount)
	store.Set(GetOldVestingAccountTraceIDBytes(vestingAccount.Id), b)
}

func GetOldVestingAccountTraceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetAllOldVestingAccountTraces(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (list []v2.VestingAccount) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(v2.VestingAccountKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()
	list = []v2.VestingAccount{}
	for ; iterator.Valid(); iterator.Next() {
		var val v2.VestingAccount
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetOldVestingAccountTraceCount(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) uint64 {
	store := prefix.NewStore(ctx.KVStore(storeKey), []byte{})
	byteKey := types.KeyPrefix(v2.VestingAccountCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}
