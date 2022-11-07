package v101_test

//
//import (
//	v100cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v100"
//	v100cfevesting "github.com/chain4energy/c4e-chain/x/cfevesting/migrations/v100"
//
//	"math/rand"
//	"strconv"
//	"testing"
//	"time"
//
//	"github.com/cosmos/cosmos-sdk/store/prefix"
//
//	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
//	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
//	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
//	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
//	"github.com/cosmos/cosmos-sdk/codec"
//	storetypes "github.com/cosmos/cosmos-sdk/store/types"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/stretchr/testify/require"
//)
//
//func TestMigrationManyAccountVestingPoolsWithManyPools(t *testing.T) {
//	accounts, _ := commontestutils.CreateAccounts(5, 0)
//	testUtil, _, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
//	SetupV100AccountVestingPools(testUtil, ctx, accounts[0].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
//	SetupV100AccountVestingPools(testUtil, ctx, accounts[1].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
//	SetupV100AccountVestingPools(testUtil, ctx, accounts[2].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
//	SetupV100AccountVestingPools(testUtil, ctx, accounts[3].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
//	SetupV100AccountVestingPools(testUtil, ctx, accounts[4].String(), 10, sdk.ZeroInt(), sdk.ZeroInt())
//	MigrateV100ToV101(t, testUtil, ctx)
//}
//
//func SetupV100AccountVestingPools(testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context, address string, numberOfVestingPools int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) v100cfevesting.AccountVestingPools {
//	return SetupV100AccountVestingPoolsWithModification(testUtil, ctx, func(*v100cfevesting.VestingPool) { /*do not modify*/ }, address, numberOfVestingPools, vestingAmount, withdrawnAmount)
//}
//
//func SetupV100AccountVestingPoolsWithModification(testUtil *testkeeper.ExtendedC4eVestingKeeperUtils, ctx sdk.Context, modifyVesting func(*v100cfevesting.VestingPool), address string, numberOfVestingPools int, vestingAmount sdk.Int, withdrawnAmount sdk.Int) v100cfevesting.AccountVestingPools {
//	accountVestingPools := GenerateOneV100AccountVestingPoolsWithAddressWithRandomVestingPools(numberOfVestingPools, 1, 1)
//	SetV100AccountVestingPools(ctx, testUtil.StoreKey, testUtil.Cdc, accountVestingPools)
//	return accountVestingPools
//}
//
//func SetV100AccountVestingPools(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, accountVestingPools v100cfevesting.AccountVestingPools) {
//	store := prefix.NewStore(ctx.KVStore(storeKey), v100cfevesting.AccountVestingPoolsKeyPrefix)
//	av := cdc.MustMarshal(&accountVestingPools)
//	store.Set([]byte(accountVestingPools.Address), av)
//}
//
//func MigrateV100ToV101(t *testing.T, testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils, ctx sdk.Context) {
//	oldStates, _ := getAllv100DistributorStates(ctx, testUtil.StoreKey, testUtil.Cdc)
//	err := MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
//	require.NoError(t, err)
//	newStates := testUtil.GetC4eDistributorKeeper().GetAllStates(ctx)
//	require.EqualValues(t, len(oldStates), len(newStates))
//	for i := 0; i < len(oldStates); i++ {
//		require.EqualValues(t, oldStates[i].CoinsStates, newStates[i].Remains)
//		require.EqualValues(t, oldStates[i].Burn, newStates[i].Burn)
//		require.EqualValues(t, oldStates[i].Account, newStates[i].Account)
//	}
//}
//
//func getAllv100DistributorStates(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (list []v100cfedistributor.State, err error) {
//
//	prefixStore := prefix.NewStore(ctx.KVStore(storeKey), v100cfedistributor.RemainsKeyPrefix)
//	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})
//	defer iterator.Close()
//
//	for ; iterator.Valid(); iterator.Next() {
//		var val v100cfedistributor.State
//		err := cdc.Unmarshal(iterator.Value(), &val)
//		if err != nil {
//			return nil, err
//		}
//		list = append(list, val)
//	}
//	return list, nil
//
//}
//
//func generateV100AccountVestingPools(numberOfAccounts int, numberOfVestingPoolsPerAccount int,
//	accountStartId int, vestingStartId int, generateVesting func(accuntId int, vestingId int) v100cfevesting.VestingPool) []*v100cfevesting.AccountVestingPools {
//	accountVestingPoolsArr := []*v100cfevesting.AccountVestingPools{}
//	accountsAddresses, _ := commontestutils.CreateAccounts(2*numberOfAccounts, 0)
//
//	for i := 0; i < numberOfAccounts; i++ {
//		accountVestingPools := v100cfevesting.AccountVestingPools{}
//		accountVestingPools.Address = "test-vesting-account-addr-" + strconv.Itoa(i+accountStartId)
//
//		accountVestingPools.Address = accountsAddresses[i].String()
//
//		var vestingPools []*v100cfevesting.VestingPool
//		for j := 0; j < numberOfVestingPoolsPerAccount; j++ {
//			vesting := generateVesting(i+accountStartId, j+vestingStartId)
//			vestingPools = append(vestingPools, &vesting)
//		}
//		accountVestingPools.VestingPools = vestingPools
//
//		accountVestingPoolsArr = append(accountVestingPoolsArr, &accountVestingPools)
//	}
//
//	return accountVestingPoolsArr
//}
//
//func generateRandomV100VestingPool(accuntId int, vestingId int) v100cfevesting.VestingPool {
//	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
//	initiallyLocked := rgen.Intn(10000000)
//	withdrawn := rgen.Intn(initiallyLocked)
//	sent := rgen.Intn(initiallyLocked - withdrawn)
//	lastModificationVested := rgen.Intn(10000000)
//	lastModificationWithdrawn := rgen.Intn(lastModificationVested)
//	return v100cfevesting.VestingPool{
//		Id:                        int32(vestingId),
//		Name:                      "test-vesting-account-name" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
//		VestingType:               "test-vesting-account-" + strconv.Itoa(accuntId) + "-" + strconv.Itoa(vestingId),
//		LockStart:                 testutils.CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
//		LockEnd:                   testutils.CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
//		Vested:                    sdk.NewInt(int64(initiallyLocked)),
//		Withdrawn:                 sdk.NewInt(int64(withdrawn)),
//		Sent:                      sdk.NewInt(int64(sent)),
//		LastModification:          testutils.CreateTimeFromNumOfHours(int64(rgen.Intn(100000))),
//		LastModificationVested:    sdk.NewInt(int64(lastModificationVested)),
//		LastModificationWithdrawn: sdk.NewInt(int64(lastModificationWithdrawn)),
//	}
//}
//
//func setV100VestingTypes(ctx sdk.Context, vestingTypes types.VestingTypes, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
//	store := ctx.KVStore(storeKey)
//	b := cdc.MustMarshal(&vestingTypes)
//	store.Set(v100cfevesting.VestingTypesKey, b)
//}
//
//func getAllV100VestingType(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (vestingTypes types.VestingTypes) {
//	store := ctx.KVStore(storeKey)
//	b := store.Get(v100cfevesting.VestingTypesKey)
//	if b == nil {
//		return vestingTypes
//	}
//
//	cdc.MustUnmarshal(b, &vestingTypes)
//	return
//
//}
