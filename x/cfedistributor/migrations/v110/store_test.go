package v110_test

import (
	"testing"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v101"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v110"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationStatesBurnFalse(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(testcosmos.DefaultTestDenom, sdk.NewInt(0))
	states := []v101.State{
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
	}
	setV101States(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV101ToV110(t, testUtil, ctx)
}

func TestMigrationStatesBurnTrue(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(testcosmos.DefaultTestDenom, sdk.NewInt(100))
	states := []v101.State{
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
	}
	setV101States(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV101ToV110(t, testUtil, ctx)
}

func TestMigrationStatesBurnTrueAndFalse(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(testcosmos.DefaultTestDenom, sdk.NewInt(100))
	coin2 := sdk.NewDecCoin(testcosmos.DefaultTestDenom, sdk.NewInt(300))
	states := []v101.State{
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin2),
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin2),
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin2),
		createV101SubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
	}
	setV101States(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV101ToV110(t, testUtil, ctx)
}

func setV101States(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, states []v101.State) {
	for _, state := range states {
		store := prefix.NewStore(ctx.KVStore(storeKey), v101.RemainsKeyPrefix)
		av := cdc.MustMarshal(&state)
		store.Set([]byte(OldGetStateKey(state)), av)
	}
}

func OldGetStateKey(state v101.State) string {
	if state.Account != nil && state.Account.Id != "" {
		return state.Account.Id
	} else {
		return "burn_state_key"
	}
}

func MigrateStoreV101ToV110(
	t *testing.T,
	testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils,
	ctx sdk.Context,
) {
	store := prefix.NewStore(ctx.KVStore(testUtil.StoreKey), v101.RemainsKeyPrefix)
	oldStates := GetAllV101States(store, testUtil.Cdc)

	err := v110.MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
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

func GetAllV101States(store storetypes.KVStore, cdc codec.BinaryCodec) (list []v101.State) {
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val v101.State
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func createV101SubdistributorState(id string, accType string, burn bool, coinState sdk.DecCoin) v101.State {
	return v101.State{
		Account: &v101.Account{
			Id:   id,
			Type: accType,
		},
		CoinsStates: sdk.NewDecCoins(coinState),
		Burn:        burn,
	}
}
