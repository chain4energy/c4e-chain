package v2_test

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v1"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v2"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationCorrectMinterState(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV1MinterState(1, sdk.ZeroDec(), sdk.ZeroDec(), time.Now(), math.NewInt(10000))
	setV1MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV1(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationWrongMinterStateNegativeAmount(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV1MinterState(1, sdk.ZeroDec(), sdk.ZeroDec(), time.Now(), math.NewInt(-10000))
	setV1MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV1(t, ctx, *k, &keeperData, true, "minter state validation error: amountMinted cannot be less than 0")
}

func TestMigrationWrongMinterStateNegativeRemainder(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV1MinterState(1, sdk.MustNewDecFromStr("-100"), sdk.ZeroDec(), time.Now(), math.NewInt(10000))
	setV1MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV1(t, ctx, *k, &keeperData, true, "minter state validation error: remainderToMint cannot be less than 0")
}

func TestMigrationWrongMinterStateNegativePreviousRemainder(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV1MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("-100"), time.Now(), math.NewInt(10000))
	setV1MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV1(t, ctx, *k, &keeperData, true, "minter state validation error: remainderFromPreviousPeriod cannot be less than 0")
}

func TestMigrationNoMinterStates(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	MigrateStoreV100ToV1(t, ctx, *k, &keeperData, true, "stored minter state should not have been nil")
}

func TestMigrationMinterStateHistory(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	stateHistory := []v1.MinterState{
		createV1MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), math.NewInt(10000)),
		createV1MinterState(2, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), math.NewInt(10001)),
	}
	minterState := createV1MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), math.NewInt(10000))
	setV1MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, minterState)
	setOldMinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc, stateHistory)
	MigrateStoreV100ToV1(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationWrongMinterStateHistory(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	stateHistory := []v1.MinterState{
		createV1MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("-100"), time.Now(), math.NewInt(10000)),
		createV1MinterState(2, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), math.NewInt(10001)),
	}
	minterState := createV1MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), math.NewInt(10000))
	setV1MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, minterState)
	setOldMinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc, stateHistory)
	MigrateStoreV100ToV1(t, ctx, *k, &keeperData, true, "minter state validation error: remainderFromPreviousPeriod cannot be less than 0")
}

func TestMigrationNoStateHistory(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	stateHistory := []v1.MinterState{}
	minterState := createV1MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), math.NewInt(10000))
	setV1MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, minterState)
	setOldMinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc, stateHistory)
	MigrateStoreV100ToV1(t, ctx, *k, &keeperData, false, "")
}

func MigrateStoreV100ToV1(
	t *testing.T,
	ctx sdk.Context,
	keeper keeper.Keeper,
	keeperData *testenv.AdditionalKeeperData,
	expectError bool, errorMessage string,
) {
	oldState := getV1MinterState(ctx, keeperData.StoreKey, keeperData.Cdc)
	oldMinterHistory := getV1MinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc)

	err := v2.MigrateStore(ctx, keeperData.StoreKey, keeperData.Cdc)
	if expectError {
		require.EqualError(t, err, errorMessage)
		return
	}
	require.NoError(t, err)

	newState := getV2MinterState(ctx, keeperData.StoreKey, keeperData.Cdc)
	require.Equal(t, newState.AmountMinted, oldState.AmountMinted)
	require.Equal(t, newState.RemainderFromPreviousPeriod, oldState.RemainderFromPreviousPeriod)
	require.Equal(t, newState.RemainderToMint, oldState.RemainderToMint)
	require.Equal(t, newState.LastMintBlockTime, oldState.LastMintBlockTime)
	require.EqualValues(t, newState.SequenceId, oldState.Position)

	newMinterStateHistory := getV2MinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc)
	require.Equal(t, len(oldMinterHistory), len(newMinterStateHistory))
	for i, oldMinterHistory := range oldMinterHistory {
		require.Equal(t, newMinterStateHistory[i].AmountMinted, oldMinterHistory.AmountMinted)
		require.Equal(t, newMinterStateHistory[i].RemainderFromPreviousPeriod, oldMinterHistory.RemainderFromPreviousPeriod)
		require.Equal(t, newMinterStateHistory[i].RemainderToMint, oldMinterHistory.RemainderToMint)
		require.Equal(t, newMinterStateHistory[i].LastMintBlockTime, oldMinterHistory.LastMintBlockTime)
		require.EqualValues(t, newMinterStateHistory[i].SequenceId, oldMinterHistory.Position)
	}

}

func createV1MinterState(
	position int32,
	remainderToMint,
	remainderFromPreviousPeriod sdk.Dec,
	lastMintBlockTime time.Time,
	amountMinted math.Int,
) v1.MinterState {
	return v1.MinterState{
		Position:                    position,
		RemainderToMint:             remainderToMint,
		RemainderFromPreviousPeriod: remainderFromPreviousPeriod,
		LastMintBlockTime:           lastMintBlockTime,
		AmountMinted:                amountMinted,
	}
}

func setV1MinterState(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, state v1.MinterState) {
	store := ctx.KVStore(storeKey)
	b := cdc.MustMarshal(&state)
	store.Set(types.MinterStateKey, b)
}

func getV1MinterState(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (minterState v1.MinterState) {
	store := ctx.KVStore(storeKey)
	b := store.Get(types.MinterStateKey)
	cdc.MustUnmarshal(b, &minterState)
	return
}

func getV2MinterState(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (minterState types.LegacyMinterState) {
	store := ctx.KVStore(storeKey)
	b := store.Get(types.MinterStateKey)
	cdc.MustUnmarshal(b, &minterState)
	return
}

func getV1MinterStateHistory(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (minterStateList []*v1.MinterState) {
	store := ctx.KVStore(storeKey)
	prefixStore := prefix.NewStore(store, v1.MinterStateHistoryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v1.MinterState
		cdc.MustUnmarshal(iterator.Value(), &val)
		minterStateList = append(minterStateList, &val)
	}

	return
}

func getV2MinterStateHistory(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (minterStateList []*types.LegacyMinterState) {
	store := ctx.KVStore(storeKey)
	prefixStore := prefix.NewStore(store, v1.MinterStateHistoryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LegacyMinterState
		cdc.MustUnmarshal(iterator.Value(), &val)
		minterStateList = append(minterStateList, &val)
	}

	return
}

func setOldMinterStateHistory(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, minterStateList []v1.MinterState) {
	store := ctx.KVStore(storeKey)
	prefixStore := prefix.NewStore(store, types.MinterStateHistoryKeyPrefix)
	for _, V1MinterState := range minterStateList {
		av := cdc.MustMarshal(&V1MinterState)
		prefixStore.Set([]byte(strconv.FormatInt(int64(V1MinterState.Position), 10)), av)
	}
}
