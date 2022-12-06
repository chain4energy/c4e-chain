package v110_test

import (
	"github.com/chain4energy/c4e-chain/testutil/common"
	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	v101 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	v110 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v110"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationCorrectMinterState(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV101MinterState(1, sdk.ZeroDec(), sdk.ZeroDec(), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, "")
}

func TestMigrationWrongMinterStateNegativeAmount(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV101MinterState(1, sdk.ZeroDec(), sdk.ZeroDec(), time.Now(), sdk.NewInt(-10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, "minter state amount cannot be less than 0")
}

func TestMigrationWrongMinterStateNegativeRemainder(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV101MinterState(1, sdk.MustNewDecFromStr("-100"), sdk.ZeroDec(), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, "minter remainder to mint amount cannot be less than 0")
}

func TestMigrationWrongMinterStateNegativePreviousRemainder(t *testing.T) {
	k, ctx, keeperData := testkeeper.CfeminterKeeper(t)
	state := createV101MinterState(1, sdk.ZeroDec(), sdk.MustNewDecFromStr("-100"), time.Now(), sdk.NewInt(10000))
	setV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc, state)
	MigrateStoreV100ToV101(t, ctx, *k, &keeperData, "minter remainder from previous period amount cannot be less than 0")
}

func MigrateStoreV100ToV101(
	t *testing.T,
	ctx sdk.Context,
	keeper keeper.Keeper,
	keeperData *common.AdditionalKeeperData,
	errorMessage string,
) {
	oldState := getV101MinterState(ctx, keeperData.StoreKey, keeperData.Cdc)
	err := v110.MigrateStore(ctx, keeperData.StoreKey, keeperData.Cdc)

	if len(errorMessage) > 0 {
		require.Equal(t, err.Error(), errorMessage)
		return
	}
	require.NoError(t, err)

	newState := keeper.GetMinterState(ctx)
	require.Equal(t, newState.AmountMinted, oldState.AmountMinted)
	require.Equal(t, newState.RemainderFromPreviousPeriod, oldState.RemainderFromPreviousPeriod)
	require.Equal(t, newState.RemainderToMint, oldState.RemainderToMint)
	require.Equal(t, newState.LastMintBlockTime, oldState.LastMintBlockTime)
	require.EqualValues(t, newState.SequenceId, oldState.Position)
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
	if b == nil {
		panic("stored minter state should not have been nil")
	}

	cdc.MustUnmarshal(b, &minterState)
	return
}
