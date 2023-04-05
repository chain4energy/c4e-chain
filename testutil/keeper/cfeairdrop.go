package keeper

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
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

	commontestutils "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	cfeairdroptestutils "github.com/chain4energy/c4e-chain/testutil/module/cfeairdrop"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// type ExtendedC4eAirdropKeeperUtils struct {
// 	cfeairdroptestutils.C4eAirdropKeeperUtils
// 	BankUtils *commontestutils.BankUtils
// 	Cdc      *codec.ProtoCodec
// 	StoreKey *storetypes.KVStoreKey
// }

// func NewExtendedC4eAirdropKeeperUtils(t *testing.T, helperCfeairdropKeeper *keeper.Keeper,
// 	bankUtils *commontestutils.BankUtils,
// 	cdc *codec.ProtoCodec,
// 	storeKey *storetypes.KVStoreKey) ExtendedC4eAirdropKeeperUtils {
// 	return ExtendedC4eAirdropKeeperUtils{C4eAirdropKeeperUtils: cfeairdroptestutils.NewC4eAirdropKeeperUtils(t, helperCfeairdropKeeper),
// 		BankUtils: bankUtils,
// 		Cdc:      cdc,
// 		StoreKey: storeKey}
// }

// func CfeairdropKeeperTestUtil(t *testing.T) (*cfeairdroptestutils.C4eAirdropKeeperUtils, *keeper.Keeper, sdk.Context) {
// 	k, ctx := CfeairdropKeeper(t)
// 	utils := cfeairdroptestutils.NewC4eAirdropKeeperUtils(t, k)
// 	return &utils, k, ctx
// }

func CfeairdropKeeperTestUtilWithCdc(t *testing.T) (*cfeairdroptestutils.C4eAirdropUtils, *keeper.Keeper, sdk.Context) {
	k, ak, bk, ctx, _, _ := cfeairdropKeeperWithBlockHeightAndTime(t, 0, testenv.TestEnvTime)
	bankUtils := commontestutils.NewBankUtils(t, ctx, ak, bk)
	utils := cfeairdroptestutils.NewC4eAirdropUtils(t, k, ak, &bankUtils, nil, nil, nil, nil)
	return &utils, k, ctx
}

func cfeairdropKeeperWithBlockHeightAndTime(t testing.TB, blockHeight int64, blockTime time.Time) (*keeper.Keeper, *authkeeper.AccountKeeper, bankkeeper.Keeper, sdk.Context, *codec.ProtoCodec, *storetypes.KVStoreKey) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	authStoreKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	feegrantStoreKey := sdk.NewKVStoreKey(feegrant.StoreKey)
	stakingStoreKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(authStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(bankStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(feegrantStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(stakingStoreKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CfeairdropParams",
	)

	accParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		authStoreKey,
		authStoreKey,
		"acc",
	)

	bankParamsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		bankStoreKey,
		bankStoreKey,
		"bankParams",
	)

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc, authStoreKey, accParamsSubspace, authtypes.ProtoBaseAccount, commontestutils.AddHelperModuleAccountPermissions(map[string][]string{types.ModuleName: nil}), testenv.DefaultBechPrefix,
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc, bankStoreKey, accountKeeper, bankParamsSubspace, map[string]bool{},
	)
	feegrantKeeper := feegrantkeeper.NewKeeper(
		cdc, feegrantStoreKey, accountKeeper,
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		accountKeeper,
		bankKeeper,
		feegrantKeeper,
		nil,
		nil,
	)

	header := tmproto.Header{}
	header.Height = blockHeight
	header.Time = blockTime
	ctx := sdk.NewContext(stateStore, header, false, log.NewNopLogger())
	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	accountKeeper.GetModuleAccount(ctx, types.ModuleName)

	return k, &accountKeeper, &bankKeeper, ctx, cdc, storeKey
}

func CfeairdropKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	k, _, _, ctx, _, _ := cfeairdropKeeperWithBlockHeightAndTime(t, 0, time.Now())
	return k, ctx

}
