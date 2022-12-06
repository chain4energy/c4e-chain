package v110_test

import (
	v101 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v101"
	v110 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v110"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationStatesBurnFalse(t *testing.T) {
	testUtil, ctx, _ := testkeeper.CfeminterKeeperTestUtilWithCdc(t)

	state := createV101MinterState(1, sdk.ZeroDec(), sdk.ZeroDec(), time.Now(), sdk.NewInt(10000))
	setOldState(ctx, testUtil.StoreKey, testUtil.Cdc, state)
	MigrateStoreV100ToV101(t, testUtil, ctx, false)
}

func setOldState(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, state v101.MinterState) {
	store := ctx.KVStore(storeKey)
	b := cdc.MustMarshal(&state)
	store.Set(types.MinterStateKey, b)
}

func MigrateStoreV100ToV101(
	t *testing.T,
	testUtil *testkeeper.ExtendedC4eMinterKeeperUtils,
	ctx sdk.Context,
	wantError bool,
) {

	store := prefix.NewStore(ctx.KVStore(testUtil.StoreKey), v101.MinterStateKey)
	oldStates := GetV101MinterState(store, testUtil.Cdc)

	err := v110.MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
	if wantError {
		require.Error(t, err)
		return
	}
	require.NoError(t, err)

	newStates := testUtil.GetC4eMinterKeeper().GetMinterState(ctx)

	require.EqualValues(t, newStates, oldStates)

}

func GetV101MinterState(store storetypes.KVStore, cdc codec.BinaryCodec) (minterState v101.MinterState) {
	b := store.Get(types.MinterStateKey)
	if b == nil {
		panic("stored minter state should not have been nil")
	}

	cdc.MustUnmarshal(b, &minterState)
	return
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
