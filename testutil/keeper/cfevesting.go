package keeper

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"

)

func CfevestingKeeperWithBlockHeightAndTimeAndStore(t testing.TB, blockHeight int64, blockTime time.Time, db *tmdb.MemDB, stateStore storetypes.CommitMultiStore) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	// stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CfevestingParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		nil,
		nil,
		nil,
		nil,
	)

	header := tmproto.Header{}
	header.Height = blockHeight
	header.Time = blockTime
	ctx := sdk.NewContext(stateStore, header, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

func CfevestingKeeperWithBlockHeightAndTime(t testing.TB, blockHeight int64, blockTime time.Time) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	return CfevestingKeeperWithBlockHeightAndTimeAndStore(t, blockHeight, blockTime, db, stateStore)
}

func CfevestingKeeperWithBlockHeight(t testing.TB, blockHeight int64) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	return CfevestingKeeperWithBlockHeightAndTimeAndStore(t, blockHeight, testutils.TestEnvTime, db, stateStore)
}

func CfevestingKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	return CfevestingKeeperWithBlockHeightAndTimeAndStore(t, 0, testutils.TestEnvTime, db, stateStore)
}

func CfevestingKeeperWithStore(t testing.TB) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	return CfevestingKeeperWithBlockHeightAndTimeAndStore(t, 0, testutils.TestEnvTime, db, stateStore)
}

func AccountKeeperWithBlockHeight(t testing.TB, ctx sdk.Context, stateStore storetypes.CommitMultiStore, db *tmdb.MemDB) (*authkeeper.AccountKeeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey + "mem")

	// db := tmdb.NewMemDB()
	// stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"accountParams",
	)

	maccPerms := map[string][]string{
		types.ModuleName: nil,
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	k := authkeeper.NewAccountKeeper(
		cdc, sdk.NewKVStoreKey(authtypes.StoreKey), paramsSubspace, authtypes.ProtoBaseAccount, maccPerms,
	)

	// Initialize params
	k.SetParams(ctx, authtypes.DefaultParams())

	return &k, ctx
}
