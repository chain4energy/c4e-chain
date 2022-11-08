package v101_test

import (
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	v100cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v100"
	v101cfedistributor "github.com/chain4energy/c4e-chain/x/cfedistributor/migrations/v101"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/stretchr/testify/require"
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMigrationStatesBurnFalse(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(commontestutils.DefaultTestDenom, sdk.NewInt(0))
	states := []v100cfedistributor.State{
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
	}
	setOldStates(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV100ToV101(t, testUtil, ctx, false)
}

func TestMigrationStatesBurnTrue(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(commontestutils.DefaultTestDenom, sdk.NewInt(100))
	states := []v100cfedistributor.State{
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
	}
	setOldStates(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV100ToV101(t, testUtil, ctx, false)
}

func TestMigrationStatesBurnTrueAndFalse(t *testing.T) {
	testUtil, ctx := testkeeper.CfedistributorKeeperTestUtilWithCdc(t)
	coin := sdk.NewDecCoin(commontestutils.DefaultTestDenom, sdk.NewInt(100))
	coin2 := sdk.NewDecCoin(commontestutils.DefaultTestDenom, sdk.NewInt(300))
	states := []v100cfedistributor.State{
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin),
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin2),
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin2),
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", true, coin2),
		createOldSubdistributorState("CUSTOM_ID", "CUSTOM_ACC_TYPE", false, coin),
	}
	setOldStates(ctx, testUtil.StoreKey, testUtil.Cdc, states)
	MigrateStoreV100ToV101(t, testUtil, ctx, false)
}

func setOldStates(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, states []v100cfedistributor.State) {
	for _, state := range states {
		store := prefix.NewStore(ctx.KVStore(storeKey), v100cfedistributor.RemainsKeyPrefix)
		av := cdc.MustMarshal(&state)
		store.Set([]byte(OldGetStateKey(state)), av)
	}
}

func OldGetStateKey(state v100cfedistributor.State) string {
	if state.Account != nil && state.Account.Id != "" {
		return state.Account.Id
	} else {
		return "burn_state_key"
	}
}

func MigrateStoreV100ToV101(
	t *testing.T,
	testUtil *testkeeper.ExtendedC4eDistributorKeeperUtils,
	ctx sdk.Context,
	wantError bool,
) {
	store := prefix.NewStore(ctx.KVStore(testUtil.StoreKey), v100cfedistributor.RemainsKeyPrefix)
	oldStates := GetAllOldStates(store, testUtil.Cdc)

	err := v101cfedistributor.MigrateStore(ctx, testUtil.StoreKey, testUtil.Cdc)
	if wantError {
		require.Error(t, err)
		return
	}
	require.NoError(t, err)

	newStates := testUtil.GetC4eDistributorKeeper().GetAllStates(ctx)
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

func GetAllOldStates(store storetypes.KVStore, cdc codec.BinaryCodec) (list []v100cfedistributor.State) {
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val v100cfedistributor.State
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func createOldSubdistributorState(id string, accType string, burn bool, coinState sdk.DecCoin) v100cfedistributor.State {
	return v100cfedistributor.State{
		Account: &v100cfedistributor.Account{
			Id:   id,
			Type: accType,
		},
		CoinsStates: sdk.NewDecCoins(coinState),
		Burn:        burn,
	}
}
