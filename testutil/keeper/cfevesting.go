package keeper

import (
	"testing"
	"time"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	cfevestingtestutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func CfevestingKeeperWithBlockHeightAndTimeAndStore(t *testing.T, blockHeight int64, blockTime time.Time, db *tmdb.MemDB, stateStore storetypes.CommitMultiStore) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

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

func CfevestingKeeperWithBlockHeightAndTime(t *testing.T, blockHeight int64, blockTime time.Time) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	return CfevestingKeeperWithBlockHeightAndTimeAndStore(t, blockHeight, blockTime, db, stateStore)
}

func CfevestingKeeperWithBlockHeight(t *testing.T, blockHeight int64) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	return CfevestingKeeperWithBlockHeightAndTimeAndStore(t, blockHeight, commontestutils.TestEnvTime, db, stateStore)
}

func CfevestingKeeper(t *testing.T) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	return CfevestingKeeperWithBlockHeightAndTimeAndStore(t, 0, commontestutils.TestEnvTime, db, stateStore)
}

func CfevestingKeeperTestUtil(t *testing.T) (*cfevestingtestutils.C4eVestingKeeperUtils, *keeper.Keeper, sdk.Context) {
	k, ctx := CfevestingKeeper(t)
	utils := cfevestingtestutils.NewC4eVestingKeeperUtils(t, k)
	return &utils, k, ctx
}
