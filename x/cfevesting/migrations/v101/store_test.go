package v101_test

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"math/rand"
	"strconv"
	"testing"
	"time"

	v100cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v100"
	v101cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v101"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMigrationManyAccountVestingPoolsWithManyPools(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[0].String(), 10)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[1].String(), 10)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[2].String(), 10)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[3].String(), 10)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[4].String(), 10)
	MigrateV100ToV101(t, testUtil, ctx)
}

func TestMigrationNoAccountVestingPoolsAndNoVestingTypes(t *testing.T) {
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	MigrateV100ToV101(t, testUtil, ctx)
}

func TestMigrationManyAccountVestingPoolsWithNoPools(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[0].String(), 0)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[1].String(), 0)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[2].String(), 0)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[3].String(), 0)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[4].String(), 0)
	MigrateV100ToV101(t, testUtil, ctx)
}

func TestMigrationOneAccountVestingPoolsWithOnePool(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[0].String(), 1)
	MigrateV100ToV101(t, testUtil, ctx)

}

func TestMigrationOneVestingType(t *testing.T) {
	vts := testutils.GenerateVestingTypes(1, 1)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	setV100VestingTypes(ctx, types.VestingTypes{VestingTypes: vts}, testUtil.StoreKey, testUtil.Cdc)
	MigrateV100ToV101(t, testUtil, ctx)
}

func TestMigrationManyVestingType(t *testing.T) {
	vts := testutils.GenerateVestingTypes(10, 1)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	setV100VestingTypes(ctx, types.VestingTypes{VestingTypes: vts}, testUtil.StoreKey, testUtil.Cdc)
	MigrateV100ToV101(t, testUtil, ctx)
}

func TestMigrationAccountVestingPoolsAndVestingTypes(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	vts := testutils.GenerateVestingTypes(10, 1)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[0].String(), 10)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[1].String(), 10)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[2].String(), 10)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[3].String(), 10)
	SetupV100AccountVestingPools(testUtil, ctx, accounts[4].String(), 10)
	setV100VestingTypes(ctx, types.VestingTypes{VestingTypes: vts}, testUtil.StoreKey, testUtil.Cdc)
	MigrateV100ToV101(t, testUtil, ctx)
}

func TestMigrationWrongSentAmount(t *testing.T) {
	accounts, _ := commontestutils.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupV100AccountVestingPoolsWrongSent(testUtil, ctx, accounts[0].String(), 10)
	SetupV100AccountVestingPoolsWrongSent(testUtil, ctx, accounts[1].String(), 10)
	SetupV100AccountVestingPoolsWrongSent(testUtil, ctx, accounts[2].String(), 10)
	SetupV100AccountVestingPoolsWrongSent(testUtil, ctx, accounts[3].String(), 10)
	SetupV100AccountVestingPoolsWrongSent(testUtil, ctx, accounts[4].String(), 10)
	MigrateV100ToV101(t, testUtil, ctx)
}

func SetupV100AccountVestingPools(testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context, address string, numberOfVestingPools int) v100cfevesting.AccountVestingPools {
	accountVestingPools := GenerateOneV100AccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPools, 1, 1)
	accountVestingPools.Address = address
	SetV100AccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc, accountVestingPools)
	return accountVestingPools
}

func SetupV100AccountVestingPoolsWrongSent(testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context, address string, numberOfVestingPools int) v100cfevesting.AccountVestingPools {
	accountVestingPools := GenerateOneV100AccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPools, 1, 1)
	accountVestingPools.Address = address
	for _, vesting := range accountVestingPools.VestingPools {
		vesting.Sent = sdk.NewInt(100)
	}
	SetV100AccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc, accountVestingPools)
	return accountVestingPools
}

func SetV100AccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, accountVestingPools v100cfevesting.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(storeKey), v100cfevesting.AccountVestingPoolsKeyPrefix)
	av := cdc.MustMarshal(&accountVestingPools)
	store.Set([]byte(accountVestingPools.Address), av)
}

func MigrateV100ToV101(t *testing.T, testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context) {
	oldAccPools := getAllV100AccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc)
	oldVestingTypes := getAllV100VestingType(ctx, testUtil.StoreKey, testUtil.Cdc)
	err := v101cfevesting.MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
	require.NoError(t, err)
	require.EqualValues(t, 0, len(getAllV100VestingType(ctx, testUtil.StoreKey, testUtil.Cdc).VestingTypes))

	newAccPools := testUtil.GetC4eVestingKeeper().GetAllAccountVestingPools(ctx)
	newVestingTypes := testUtil.GetC4eVestingKeeper().GetAllVestingTypes(ctx)

	require.EqualValues(t, len(oldAccPools), len(newAccPools))
	for i := 0; i < len(oldAccPools); i++ {
		require.EqualValues(t, oldAccPools[i].Address, newAccPools[i].Address)
		require.EqualValues(t, len(oldAccPools[i].VestingPools), len(newAccPools[i].VestingPools))
		for j := 0; j < len(oldAccPools[i].VestingPools); j++ {
			oldVestingPool := oldAccPools[i].VestingPools[j]
			newVestingPool := newAccPools[i].VestingPools[j]
			require.EqualValues(t, oldVestingPool.Name, newVestingPool.Name)
			require.EqualValues(t, oldVestingPool.VestingType, newVestingPool.VestingType)
			require.EqualValues(t, oldVestingPool.LockStart, newVestingPool.LockStart)
			require.EqualValues(t, oldVestingPool.LockEnd, newVestingPool.LockEnd)
			require.EqualValues(t, oldVestingPool.Vested, newVestingPool.InitiallyLocked)
			require.EqualValues(t, oldVestingPool.Withdrawn, newVestingPool.Withdrawn)
			oldSentCalculated := oldVestingPool.LastModificationWithdrawn.Add(oldVestingPool.Vested).Sub(oldVestingPool.Withdrawn).Sub(oldVestingPool.LastModificationVested)
			require.True(t, oldSentCalculated.Equal(newAccPools[i].VestingPools[j].Sent))
		}
	}
	require.ElementsMatch(t, oldVestingTypes.VestingTypes, newVestingTypes.VestingTypes)
}

func getAllV100AccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (list []v100cfevesting.AccountVestingPools) {
	prefixStore := prefix.NewStore(ctx.KVStore(storeKey), v100cfevesting.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v100cfevesting.AccountVestingPools
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func GenerateOneV100AccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPoolsPerAccount int,
	accountId int, vestingStartId int) v100cfevesting.AccountVestingPools {
	return *GenerateV100AccountVestingPoolsWithRandomVestingPools(1, numberOfVestingPoolsPerAccount, accountId, vestingStartId)[0]
}

func GenerateV100AccountVestingPoolsWithRandomVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int) []*v100cfevesting.AccountVestingPools {
	return generateV100AccountVestingPools(numberOfAccounts, numberOfVestingPoolsPerAccount,
		accountStartId, vestingStartId, generateRandomV100VestingPool)
}

func generateV100AccountVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) v100cfevesting.VestingPool) []*v100cfevesting.AccountVestingPools {
	accountVestingPoolsArr := []*v100cfevesting.AccountVestingPools{}
	accountsAddresses, _ := commontestutils.CreateAccounts(2*numberOfAccounts, 0)

	for i := 0; i < numberOfAccounts; i++ {
		accountVestingPools := v100cfevesting.AccountVestingPools{}
		accountVestingPools.Address = "test-vesting-account-addr-" + strconv.Itoa(i+accountStartId)

		accountVestingPools.Address = accountsAddresses[i].String()

		var vestingPools []*v100cfevesting.VestingPool
		for j := 0; j < numberOfVestingPoolsPerAccount; j++ {
			vesting := generateVesting(i+accountStartId, j+vestingStartId)
			vestingPools = append(vestingPools, &vesting)
		}
		accountVestingPools.VestingPools = vestingPools

		accountVestingPoolsArr = append(accountVestingPoolsArr, &accountVestingPools)
	}

	return accountVestingPoolsArr
}

func generateRandomV100VestingPool(accuntId int, vestingId int) v100cfevesting.VestingPool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	vested := int(helpers.RandIntBetweenWith0(r, 1, 10000000))
	withdrawn := r.Intn(vested)
	sent := helpers.RandIntWith0(r, vested-withdrawn)
	randWith := helpers.RandIntWith0(r, withdrawn)
	lastModificationVested := vested - sent - randWith
	lastModificationWithdrawn := withdrawn - randWith

	return v100cfevesting.VestingPool{
		Id:                        int32(vestingId),
		Name:                      "test-vesting-account-name" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		VestingType:               "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
		LockStart:                 testutils.CreateTimeFromNumOfHours(int64(r.Intn(100000))),
		LockEnd:                   testutils.CreateTimeFromNumOfHours(int64(r.Intn(100000))),
		Vested:                    sdk.NewInt(int64(vested)),
		Withdrawn:                 sdk.NewInt(int64(withdrawn)),
		Sent:                      sdk.NewInt(int64(sent)),
		LastModification:          testutils.CreateTimeFromNumOfHours(int64(r.Intn(100000))),
		LastModificationVested:    sdk.NewInt(int64(lastModificationVested)),
		LastModificationWithdrawn: sdk.NewInt(int64(lastModificationWithdrawn)),
	}
}

func setV100VestingTypes(ctx sdk.Context, vestingTypes types.VestingTypes, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
	store := ctx.KVStore(storeKey)
	b := cdc.MustMarshal(&vestingTypes)
	store.Set(v100cfevesting.VestingTypesKey, b)
}

func getAllV100VestingType(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (vestingTypes types.VestingTypes) {
	store := ctx.KVStore(storeKey)
	b := store.Get(v100cfevesting.VestingTypesKey)
	if b == nil {
		return vestingTypes
	}

	cdc.MustUnmarshal(b, &vestingTypes)
	return

}
