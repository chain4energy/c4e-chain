package v110_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	v101 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	v110 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v110"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationCorrectMinterState(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV101MinterState(1, sdk.ZeroDec(), sdk.ZeroDec(), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationWrongMinterStateNegativeAmount(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV101MinterState(1, sdk.ZeroDec(), sdk.ZeroDec(), time.Now(), sdk.NewInt(-10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, true, "minter state validation error: amountMinted cannot be less than 0")
}

func TestMigrationWrongMinterStateNegativeRemainder(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV101MinterState(1, sdk.MustNewDecFromStr("-100"), sdk.ZeroDec(), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, true, "minter state validation error: remainderToMint cannot be less than 0")
}

func TestMigrationWrongMinterStateNegativePreviousRemainder(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV101MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("-100"), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, true, "minter state validation error: remainderFromPreviousPeriod cannot be less than 0")
}

func TestMigrationNoMinterStates(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, true, "stored minter state should not have been nil")
}

func TestMigrationMinterStateHistory(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	stateHistory := []v101.MinterState{
		createV101MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), sdk.NewInt(10000)),
		createV101MinterState(2, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), sdk.NewInt(10001)),
	}
	minterState := createV101MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, minterState)
	setOldMinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc, stateHistory)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, false, "")
}

func TestMigrationWrongMinterStateHistory(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	stateHistory := []v101.MinterState{
		createV101MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("-100"), time.Now(), sdk.NewInt(10000)),
		createV101MinterState(2, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), sdk.NewInt(10001)),
	}
	minterState := createV101MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, minterState)
	setOldMinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc, stateHistory)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, true, "minter state validation error: remainderFromPreviousPeriod cannot be less than 0")
}

func TestMigrationNoStateHistory(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	stateHistory := []v101.MinterState{}
	minterState := createV101MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("100"), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, minterState)
	setOldMinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc, stateHistory)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, false, "")
}

func MigrateStoreV100ToV101(
	t *testing.T,
	ctx sdk.Context,
	keeper keeper.Keeper,
	keeperData *cosmossdk.AdditionalKeeperData,
	expectError bool, errorMessage string,
) {
	oldState := getV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc)
	oldMinterHistory := getV101MinterStateHistory(ctx, keeperData.StoreKey, keeperData.Cdc)

	err := v110.MigrateStore(ctx, keeperData.StoreKey, keeperData.Cdc)
	if expectError {
		require.EqualError(t, err, errorMessage)
		return
	}
	require.NoError(t, err)

	newState := keeper.GetMinterState(ctx)
	require.Equal(t, newState.AmountMinted, oldState.AmountMinted)
	require.Equal(t, newState.RemainderFromPreviousPeriod, oldState.RemainderFromPreviousPeriod)
	require.Equal(t, newState.RemainderToMint, oldState.RemainderToMint)
	require.Equal(t, newState.LastMintBlockTime, oldState.LastMintBlockTime)
	require.EqualValues(t, newState.SequenceId, oldState.Position)

	newMinterStateHistory := keeper.GetAllMinterStateHistory(ctx)
	require.Equal(t, len(oldMinterHistory), len(newMinterStateHistory))
	for i, oldMinterHistory := range oldMinterHistory {
		require.Equal(t, newMinterStateHistory[i].AmountMinted, oldMinterHistory.AmountMinted)
		require.Equal(t, newMinterStateHistory[i].RemainderFromPreviousPeriod, oldMinterHistory.RemainderFromPreviousPeriod)
		require.Equal(t, newMinterStateHistory[i].RemainderToMint, oldMinterHistory.RemainderToMint)
		require.Equal(t, newMinterStateHistory[i].LastMintBlockTime, oldMinterHistory.LastMintBlockTime)
		require.EqualValues(t, newMinterStateHistory[i].SequenceId, oldMinterHistory.Position)
	}

}

func createV101MinterState(
	position int32,
	remainderToMint,
	remainderFromPreviousPeriod sdk.Dec,
	lastMintBlockTime time.Time,
	amountMinted sdk.Int,
) v101.MinterState {
	return v101.MinterState{
		Position:                    position,
		RemainderToMint:             remainderToMint,
		RemainderFromPreviousPeriod: remainderFromPreviousPeriod,
		LastMintBlockTime:           lastMintBlockTime,
		AmountMinted:                amountMinted,
	}
}

func setV101MinterState(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, state v101.MinterState) {
	store := ctx.KVStore(storeKey)
	b := cdc.MustMarshal(&state)
	store.Set(types.MinterStateKey, b)
}

func getV101MinterState(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (minterState v101.MinterState) {
	store := ctx.KVStore(storeKey)
	b := store.Get(types.MinterStateKey)
	cdc.MustUnmarshal(b, &minterState)
	return
}

func getV101MinterStateHistory(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (minterStateList []*v101.MinterState) {
	store := ctx.KVStore(storeKey)
	prefixStore := prefix.NewStore(store, v101.MinterStateHistoryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(prefixStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v101.MinterState
		cdc.MustUnmarshal(iterator.Value(), &val)
		minterStateList = append(minterStateList, &val)
	}

	return
}

func setOldMinterStateHistory(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, minterStateList []v101.MinterState) {
	store := ctx.KVStore(storeKey)
	prefixStore := prefix.NewStore(store, types.MinterStateHistoryKeyPrefix)
	for _, V101MinterState := range minterStateList {
		av := cdc.MustMarshal(&V101MinterState)
		prefixStore.Set([]byte(strconv.FormatInt(int64(V101MinterState.Position), 10)), av)
	}
}
