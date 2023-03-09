package v2_test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"

	"github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v1"
	"github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v2"
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
	MigrateV1ToV2(t, testUtil, ctx)
}

func TestMigrationNoAccountVestingPoolsAndNoVestingTypes(t *testing.T) {
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	MigrateV1ToV2(t, testUtil, ctx)
}

func TestMigrationManyAccountVestingPoolsWithNoPools(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[0].String(), 0)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[1].String(), 0)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[2].String(), 0)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[3].String(), 0)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[4].String(), 0)
	MigrateV1ToV2(t, testUtil, ctx)
}

func TestMigrationOneAccountVestingPoolsWithOnePool(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[0].String(), 1)
	MigrateV1ToV2(t, testUtil, ctx)

}

func TestMigrationOneVestingType(t *testing.T) {
	vts := testutils.GenerateVestingTypes(1, 1)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	setOldVestingTypes(ctx, types.VestingTypes{VestingTypes: vts}, testUtil.StoreKey, testUtil.Cdc)
	MigrateV1ToV2(t, testUtil, ctx)
}

func TestMigrationManyVestingType(t *testing.T) {
	vts := generateOldVestingTypes(10, 1)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	setOldVestingTypes(ctx, types.VestingTypes{VestingTypes: vts}, testUtil.StoreKey, testUtil.Cdc)
	MigrateV1ToV2(t, testUtil, ctx)
}

func TestMigrationValidatorsVestingType(t *testing.T) {
	vts := generateOldVestingTypes(10, 1)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	vts[3].Name = "Validators"
	setOldVestingTypes(ctx, types.VestingTypes{VestingTypes: vts}, testUtil.StoreKey, testUtil.Cdc)
	MigrateV1ToV2(t, testUtil, ctx)
}

func TestMigrationAccountVestingPoolsAndVestingTypes(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(5, 0)
	vts := testutils.GenerateVestingTypes(10, 1)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[0].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[1].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[2].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[3].String(), 10)
	SetupOldAccountVestingPools(testUtil, ctx, accounts[4].String(), 10)
	setOldVestingTypes(ctx, types.VestingTypes{VestingTypes: vts}, testUtil.StoreKey, testUtil.Cdc)
	MigrateV1ToV2(t, testUtil, ctx)
}

func TestMigrationWrongSentAmount(t *testing.T) {
	accounts, _ := testcosmos.CreateAccounts(5, 0)
	testUtil, _, ctx := testkeeper.CfevestingKeeperTestUtilWithCdc(t)
	SetupOldAccountVestingPoolsWrongSent(testUtil, ctx, accounts[0].String(), 10)
	SetupOldAccountVestingPoolsWrongSent(testUtil, ctx, accounts[1].String(), 10)
	SetupOldAccountVestingPoolsWrongSent(testUtil, ctx, accounts[2].String(), 10)
	SetupOldAccountVestingPoolsWrongSent(testUtil, ctx, accounts[3].String(), 10)
	SetupOldAccountVestingPoolsWrongSent(testUtil, ctx, accounts[4].String(), 10)
	MigrateV1ToV2(t, testUtil, ctx)
}

func SetupOldAccountVestingPools(testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context, address string, numberOfVestingPools int) v1.AccountVestingPools {
	accountVestingPools := generateOneOldAccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPools, 1, 1)
	accountVestingPools.Address = address
	setOldAccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc, accountVestingPools)
	return accountVestingPools
}

func SetupOldAccountVestingPoolsWrongSent(testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context, address string, numberOfVestingPools int) v1.AccountVestingPools {
	accountVestingPools := generateOneOldAccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPools, 1, 1)
	accountVestingPools.Address = address
	for _, vesting := range accountVestingPools.VestingPools {
		vesting.Sent = sdk.NewInt(100)
	}
	setOldAccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc, accountVestingPools)
	return accountVestingPools
}

func setOldAccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, accountVestingPools v1.AccountVestingPools) {
	store := prefix.NewStore(ctx.KVStore(storeKey), v1.AccountVestingPoolsKeyPrefix)
	av := cdc.MustMarshal(&accountVestingPools)
	store.Set([]byte(accountVestingPools.Address), av)
}

func MigrateV1ToV2(t *testing.T, testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context) {
	oldAccPools := getAllOldAccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc)
	oldVestingTypes := getAllOldVestingType(ctx, testUtil.StoreKey, testUtil.Cdc)
	err := v2.MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
	require.NoError(t, err)
	require.EqualValues(t, 0, len(getAllOldVestingType(ctx, testUtil.StoreKey, testUtil.Cdc).VestingTypes))

	newAccPools := testUtil.GetC4eVestingKeeper().GetAllAccountVestingPools(ctx)
	newVestingTypes := testUtil.GetC4eVestingKeeper().GetAllVestingTypes(ctx)

	for _, oldVestingType := range oldVestingTypes.VestingTypes {
		foundNewVestingType := false
		for _, newVesting := range newVestingTypes.VestingTypes {
			if oldVestingType.Name == newVesting.Name {
				require.EqualValues(t, oldVestingType.Name, newVesting.Name)
				require.EqualValues(t, oldVestingType.LockupPeriod, newVesting.LockupPeriod)
				require.EqualValues(t, oldVestingType.VestingPeriod, newVesting.VestingPeriod)
				if newVesting.Name == "Validators" {
					require.EqualValues(t, sdk.MustNewDecFromStr("0.05"), newVesting.Free)
				} else {
					require.EqualValues(t, sdk.ZeroDec(), newVesting.Free)
				}
				foundNewVestingType = true
				break
			}

		}
		require.True(t, foundNewVestingType)
	}

	require.EqualValues(t, len(oldAccPools), len(newAccPools))
	for i := 0; i < len(oldAccPools); i++ {
		require.EqualValues(t, oldAccPools[i].Address, newAccPools[i].Owner)
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
}

func getAllOldAccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (list []v1.AccountVestingPools) {
	prefixStore := prefix.NewStore(ctx.KVStore(storeKey), v1.AccountVestingPoolsKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v1.AccountVestingPools
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func generateOneOldAccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPoolsPerAccount int,
	accountId int, vestingStartId int) v1.AccountVestingPools {
	return *generateOldAccountVestingPoolsWithRandomVestingPools(1, numberOfVestingPoolsPerAccount, accountId, vestingStartId)[0]
}

func generateOldAccountVestingPoolsWithRandomVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int) []*v1.AccountVestingPools {
	return generateOldAccountVestingPools(numberOfAccounts, numberOfVestingPoolsPerAccount,
		accountStartId, vestingStartId, generateRandomOldVestingPool)
}

func generateOldAccountVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) v1.VestingPool) []*v1.AccountVestingPools {
	accountVestingPoolsArr := []*v1.AccountVestingPools{}
	accountsAddresses, _ := testcosmos.CreateAccounts(2*numberOfAccounts, 0)

	for i := 0; i < numberOfAccounts; i++ {
		accountVestingPools := v1.AccountVestingPools{}
		accountVestingPools.Address = accountsAddresses[i].String()

		var vestingPools []*v1.VestingPool
		for j := 0; j < numberOfVestingPoolsPerAccount; j++ {
			vesting := generateVesting(i+accountStartId, j+vestingStartId)
			vestingPools = append(vestingPools, &vesting)
		}
		accountVestingPools.VestingPools = vestingPools

		accountVestingPoolsArr = append(accountVestingPoolsArr, &accountVestingPools)
	}

	return accountVestingPoolsArr
}

func generateOldVestingTypes(numberOfVestingTypes int, startId int) []*types.VestingType {
	var vestingTypes []*types.VestingType

	for i := 0; i < numberOfVestingTypes; i++ {
		vestingType := types.VestingType{
			Name:          "test-vesting-type-" + strconv.Itoa(i+startId),
			LockupPeriod:  testutils.CreateDurationFromNumOfHours(1000),
			VestingPeriod: testutils.CreateDurationFromNumOfHours(5000),
			Free:          sdk.MustNewDecFromStr("0.5"),
		}
		vestingTypes = append(vestingTypes, &vestingType)
	}

	return vestingTypes
}

func generateRandomOldVestingPool(accuntId int, vestingId int) v1.VestingPool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	vested := int(helpers.RandIntBetweenWith0(r, 1, 10000000))
	withdrawn := r.Intn(vested)
	sent := helpers.RandIntWith0(r, vested-withdrawn)
	randWith := helpers.RandIntWith0(r, withdrawn)
	lastModificationVested := vested - sent - randWith
	lastModificationWithdrawn := withdrawn - randWith

	return v1.VestingPool{
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

func setOldVestingTypes(ctx sdk.Context, vestingTypes types.VestingTypes, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
	store := ctx.KVStore(storeKey)
	b := cdc.MustMarshal(&vestingTypes)
	store.Set(v1.VestingTypesKey, b)
}

func getAllOldVestingType(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (vestingTypes types.VestingTypes) {
	store := ctx.KVStore(storeKey)
	b := store.Get(v1.VestingTypesKey)
	if b == nil {
		return vestingTypes
	}

	cdc.MustUnmarshal(b, &vestingTypes)
	return

}
