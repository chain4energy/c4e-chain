package keeper

import (
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"testing"

	cfedistributortestutils "github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
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

type ExtendedC4eDistributorKeeperUtils struct {
	cfedistributortestutils.C4eDistributorKeeperUtils
	Cdc      *codec.ProtoCodec
	StoreKey *storetypes.KVStoreKey
	typesparams.Subspace
}

func NewExtendedC4eDistributorKeeperUtils(
	t *testing.T,
	helperCfedistributorKeeper *keeper.Keeper,
	cdc *codec.ProtoCodec,
	storeKey *storetypes.KVStoreKey,
	paramsStore typesparams.Subspace,
) ExtendedC4eDistributorKeeperUtils {
	return ExtendedC4eDistributorKeeperUtils{
		C4eDistributorKeeperUtils: cfedistributortestutils.NewC4eDistributorKeeperUtils(t, helperCfedistributorKeeper),
		Cdc:                       cdc,
		StoreKey:                  storeKey,
		Subspace:                  paramsStore,
	}
}

type AdditionalDistributorKeeperData struct {
	*codec.ProtoCodec
	*storetypes.KVStoreKey
	typesparams.Subspace
}

func CfedistributorKeeper(t testing.TB) (*keeper.Keeper, sdk.Context, AdditionalDistributorKeeperData) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsStore := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"Cfedistributor",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsStore,
		nil,
		nil,
		nil,
		appparams.GetAuthority(),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	cfedistributortestutils.SetTestMaccPerms()

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx, AdditionalDistributorKeeperData{
		cdc,
		storeKey,
		paramsStore,
	}
}

func CfedistributorKeeperTestUtil(t *testing.T) (*cfedistributortestutils.C4eDistributorKeeperUtils, *keeper.Keeper, sdk.Context) {
	k, ctx, _ := CfedistributorKeeper(t)
	utils := cfedistributortestutils.NewC4eDistributorKeeperUtils(t, k)
	return &utils, k, ctx
}

func CfedistributorKeeperTestUtilWithCdc(t *testing.T) (*ExtendedC4eDistributorKeeperUtils, sdk.Context) {
	k, ctx, subDistributorKeeperData := CfedistributorKeeper(t)
	utils := NewExtendedC4eDistributorKeeperUtils(
		t,
		k,
		subDistributorKeeperData.ProtoCodec,
		subDistributorKeeperData.KVStoreKey,
		subDistributorKeeperData.Subspace,
	)
	keyTable := types.ParamKeyTable() //nolint:staticcheck
	utils.Subspace.WithKeyTable(keyTable)

	return &utils, ctx
}
