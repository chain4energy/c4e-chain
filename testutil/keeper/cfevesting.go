package keeper

import (
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"testing"
	"time"

	testenv "github.com/chain4energy/c4e-chain/testutil/env"

	cfevestingtestutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
)

type ExtendedC4eVestingKeeperUtils struct {
	cfevestingtestutils.C4eVestingKeeperUtils
	Cdc         *codec.ProtoCodec
	StoreKey    *storetypes.KVStoreKey
	ParamsStore typesparams.Subspace
}

func NewExtendedC4eVestingKeeperUtils(
	t *testing.T,
	helperCfevestingKeeper *keeper.Keeper,
	cdc *codec.ProtoCodec,
	storeKey *storetypes.KVStoreKey,
	paramsStore typesparams.Subspace) ExtendedC4eVestingKeeperUtils {
	return ExtendedC4eVestingKeeperUtils{
		C4eVestingKeeperUtils: cfevestingtestutils.NewC4eVestingKeeperUtils(t, helperCfevestingKeeper),
		Cdc:                   cdc,
		StoreKey:              storeKey,
		ParamsStore:           paramsStore,
	}
}

type AdditionalVestingKeeperData struct {
	*codec.ProtoCodec
	*storetypes.KVStoreKey
	typesparams.Subspace
}

func CfevestingKeeperWithBlockHeightAndTimeAndStore(t *testing.T, blockHeight int64, blockTime time.Time, db *tmdb.MemDB,
	stateStore storetypes.CommitMultiStore) (*keeper.Keeper, sdk.Context, AdditionalVestingKeeperData) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"Cfevesting",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		nil,
		nil,
		nil,
		nil,
		nil,
		appparams.GetAuthority(),
	)

	header := tmproto.Header{}
	header.Height = blockHeight
	header.Time = blockTime
	ctx := sdk.NewContext(stateStore, header, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx, AdditionalVestingKeeperData{
		ProtoCodec: cdc,
		KVStoreKey: storeKey,
		Subspace:   paramsSubspace,
	}
}

func CfevestingKeeperWithBlockHeightAndTime(t *testing.T, blockHeight int64, blockTime time.Time) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	k, ctx, _ := CfevestingKeeperWithBlockHeightAndTimeAndStore(t, blockHeight, blockTime, db, stateStore)
	return k, ctx
}

func CfevestingKeeperWithBlockHeight(t *testing.T, blockHeight int64) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	k, ctx, _ := CfevestingKeeperWithBlockHeightAndTimeAndStore(t, blockHeight, testenv.TestEnvTime, db, stateStore)
	return k, ctx
}

func CfevestingKeeper(t *testing.T) (*keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	k, ctx, _ := CfevestingKeeperWithBlockHeightAndTimeAndStore(t, 0, testenv.TestEnvTime, db, stateStore)
	return k, ctx
}

func CfevestingKeeperTestUtil(t *testing.T) (*cfevestingtestutils.C4eVestingKeeperUtils, *keeper.Keeper, sdk.Context) {
	k, ctx := CfevestingKeeper(t)
	utils := cfevestingtestutils.NewC4eVestingKeeperUtils(t, k)
	return &utils, k, ctx
}

func CfevestingKeeperTestUtilWithCdc(t *testing.T) (*ExtendedC4eVestingKeeperUtils, *keeper.Keeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	k, ctx, additionalData := CfevestingKeeperWithBlockHeightAndTimeAndStore(t, 0, testenv.TestEnvTime, db, stateStore)
	utils := NewExtendedC4eVestingKeeperUtils(t, k, additionalData.ProtoCodec, additionalData.KVStoreKey, additionalData.Subspace)
	return &utils, k, ctx
}
