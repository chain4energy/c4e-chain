package keeper

import (
	cfemintertestutils "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
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
	"testing"
)

type ExtendedC4eMinterKeeperUtils struct {
	cfemintertestutils.C4eMinterUtils
	Cdc      *codec.ProtoCodec
	StoreKey *storetypes.KVStoreKey
	typesparams.Subspace
}

func NewExtendedC4eMinterKeeperUtils(
	cdc *codec.ProtoCodec,
	storeKey *storetypes.KVStoreKey,
	paramsStore typesparams.Subspace,
) ExtendedC4eMinterKeeperUtils {
	return ExtendedC4eMinterKeeperUtils{
		Cdc:      cdc,
		StoreKey: storeKey,
		Subspace: paramsStore,
	}
}

type AdditionalMinterKeeperData struct {
	*codec.ProtoCodec
	*storetypes.KVStoreKey
	typesparams.Subspace
}

func CfeminterKeeper(t testing.TB) (*keeper.Keeper, sdk.Context, AdditionalDistributorKeeperData) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
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
		paramsSubspace,
		nil,
		nil,
		"test",
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx, AdditionalDistributorKeeperData{
		cdc,
		storeKey,
		paramsSubspace,
	}
}

func CfeminterKeeperTestUtilWithCdc(t *testing.T) (*ExtendedC4eMinterKeeperUtils, sdk.Context, keeper.Keeper) {
	k, ctx, subDistributorKeeperData := CfeminterKeeper(t)

	utils := NewExtendedC4eMinterKeeperUtils(
		subDistributorKeeperData.ProtoCodec,
		subDistributorKeeperData.KVStoreKey,
		subDistributorKeeperData.Subspace,
	)
	return &utils, ctx, *k
}
