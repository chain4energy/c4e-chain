package v2_test

import (
	"cosmossdk.io/math"
	"testing"

	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"

	v1 "github.com/chain4energy/c4e-chain/v2/x/cfedistributor/migrations/v1"
	v2 "github.com/chain4energy/c4e-chain/v2/x/cfedistributor/migrations/v2"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/chain4energy/c4e-chain/v2/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationStatesBurnFalse(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(testenv.DefaultTestDenom, math.NewInt(0))
	states := []v1.State{
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
	}
	setV1States(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV1ToV2(t, testUtil, ctx)
}

func TestMigrationStatesBurnTrue(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(testenv.DefaultTestDenom, math.NewInt(100))
	states := []v1.State{
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
	}
	setV1States(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV1ToV2(t, testUtil, ctx)
}

func TestMigrationStatesBurnTrueAndFalse(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(testenv.DefaultTestDenom, math.NewInt(100))
	coin2 := sdk.NewDecCoin(testenv.DefaultTestDenom, math.NewInt(300))
	states := []v1.State{
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin2),
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin2),
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin2),
		createV1SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
	}
	setV1States(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV1ToV2(t, testUtil, ctx)
}

func setV1States(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, states []v1.State) {
	for _, state := range states {
		store := prefix.NewStore(ctx.KVStore(storeKey), v1.RemainsKeyPrefix)
		av := cdc.MustMarshal(&state)
		store.Set([]byte(OldGetStateKey(state)), av)
	}
}

func OldGetStateKey(state v1.State) string {
	if state.Account != nil && state.Account.Id != "" {
		return state.Account.Id
	} else {
		return "burn_state_key"
	}
}

func MigrateStoreV1ToV2(
	t *testing.T,
	testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils,
	ctx sdk.Context,
) {
	store := prefix.NewStore(ctx.KVStore(testUtil.StoreKey), v1.RemainsKeyPrefix)
	oldStates := GetAllV1States(store, testUtil.Cdc)

	err := v2.MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
	require.NoError(t, err)

	newStates := testUtil.GetC4eDistributorKeeper().GetAllStates(ctx)

	require.EqualValues(t, len(newStates), len(oldStates))
	for i, oldState := range oldStates {
		if oldState.Burn == true {
			require.Nil(t, newStates[i].Account)
		} else {
			require.EqualValues(t, newStates[i].Account.Id, oldState.Account.Id)
			require.EqualValues(t, newStates[i].Account.Type, oldState.Account.Type)
		}
		require.EqualValues(t, newStates[i].Burn, oldState.Burn)
		require.ElementsMatch(t, newStates[i].Remains, oldState.CoinsStates)
	}
}

func GetAllV1States(store storetypes.KVStore, cdc codec.BinaryCodec) (list []v1.State) {
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val v1.State
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func createV1SubdistributorState(id string, accType string, burn bool, coinState sdk.DecCoin) v1.State {
	return v1.State{
		Account: &v1.Account{
			Id:   id,
			Type: accType,
		},
		CoinsStates: sdk.NewDecCoins(coinState),
		Burn:        burn,
	}
}
