package keeper

import (
	appparams "github.com/chain4energy/c4e-chain/app/params"
	cfemintertestutils "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v2"
	"testing"

	testenv "github.com/chain4energy/c4e-chain/testutil/env"

	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
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

type ExtendedC4eMinterKeeperUtils struct {
	cfemintertestutils.C4eMinterKeeperUtils
	Cdc      *codec.ProtoCodec
	StoreKey *storetypes.KVStoreKey
	typesparams.Subspace
}

func NewExtendedC4eMinterKeeperUtils(
	t *testing.T,
	helperCfedistributorKeeper *keeper.Keeper,
	cdc *codec.ProtoCodec,
	storeKey *storetypes.KVStoreKey,
	paramsStore typesparams.Subspace,
) ExtendedC4eMinterKeeperUtils {
	return ExtendedC4eMinterKeeperUtils{
		C4eMinterKeeperUtils: cfemintertestutils.NewC4eMinterKeeperUtils(t, helperCfedistributorKeeper),
		Cdc:                  cdc,
		StoreKey:             storeKey,
		Subspace:             paramsStore,
	}
}

type AdditionalMinterKeeperData struct {
	*codec.ProtoCodec
	*storetypes.KVStoreKey
	typesparams.Subspace
}

func CfeminterKeeper(t testing.TB) (*keeper.Keeper, sdk.Context, testenv.AdditionalKeeperData) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CfeminterParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		nil,
		nil,
		"test",
		appparams.GetAuthority(),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	k.SetParams(ctx, types.DefaultParams())
	keyTable := v2.ParamKeyTable()
	paramsSubspace.WithKeyTable(keyTable)

	return k, ctx, testenv.AdditionalKeeperData{
		Cdc:      cdc,
		StoreKey: storeKey,
		Subspace: paramsSubspace,
	}
}

func CfeminterKeeperTestUtil(t *testing.T) (*cfemintertestutils.C4eMinterKeeperUtils, *keeper.Keeper, sdk.Context) {
	k, ctx, _ := CfeminterKeeper(t)
	utils := cfemintertestutils.NewC4eMinterKeeperUtils(t, k)
	return &utils, k, ctx
}

func CfeminterKeeperTestUtilWithCdc(t *testing.T) (*ExtendedC4eMinterKeeperUtils, sdk.Context) {
	k, ctx, additionalKeeperData := CfeminterKeeper(t)
	utils := NewExtendedC4eMinterKeeperUtils(
		t,
		k,
		additionalKeeperData.Cdc,
		additionalKeeperData.StoreKey,
		additionalKeeperData.Subspace,
	)

	return &utils, ctx
}
