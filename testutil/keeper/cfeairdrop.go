package keeper

import (
	appparams "github.com/chain4energy/c4e-chain/app/params"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
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
	cfeclaimtestutils "github.com/chain4energy/c4e-chain/testutil/module/cfeclaim"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// type ExtendedC4eClaimKeeperUtils struct {
// 	cfeclaimtestutils.C4eClaimKeeperUtils
// 	BankUtils *commontestutils.BankUtils
// 	Cdc      *codec.ProtoCodec
// 	StoreKey *storetypes.KVStoreKey
// }

// func NewExtendedC4eClaimKeeperUtils(t *testing.T, helperCfeclaimKeeper *keeper.Keeper,
// 	bankUtils *commontestutils.BankUtils,
// 	cdc *codec.ProtoCodec,
// 	storeKey *storetypes.KVStoreKey) ExtendedC4eClaimKeeperUtils {
// 	return ExtendedC4eClaimKeeperUtils{C4eClaimKeeperUtils: cfeclaimtestutils.NewC4eClaimKeeperUtils(t, helperCfeclaimKeeper),
// 		BankUtils: bankUtils,
// 		Cdc:      cdc,
// 		StoreKey: storeKey}
// }

// func CfeclaimKeeperTestUtil(t *testing.T) (*cfeclaimtestutils.C4eClaimKeeperUtils, *keeper.Keeper, sdk.Context) {
// 	k, ctx := CfeclaimKeeper(t)
// 	utils := cfeclaimtestutils.NewC4eClaimKeeperUtils(t, k)
// 	return &utils, k, ctx
// }

func CfeclaimKeeperTestUtilWithCdc(t *testing.T) (*cfeclaimtestutils.C4eClaimUtils, *keeper.Keeper, sdk.Context) {
	k, ak, bk, ctx, _, _ := cfeclaimKeeperWithBlockHeightAndTime(t, 0, testenv.TestEnvTime)
	bankUtils := commontestutils.NewBankUtils(t, ctx, ak, bk)
	utils := cfeclaimtestutils.NewC4eClaimUtils(t, k, nil, ak, &bankUtils, nil, nil, nil, nil)
	return &utils, k, ctx
}

func cfeclaimKeeperWithBlockHeightAndTime(t testing.TB, blockHeight int64, blockTime time.Time) (*keeper.Keeper, *authkeeper.AccountKeeper, bankkeeper.Keeper, sdk.Context, *codec.ProtoCodec, *storetypes.KVStoreKey) {
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

		accountKeeper,
		bankKeeper,
		feegrantKeeper,
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

	accountKeeper.GetModuleAccount(ctx, types.ModuleName)

	return k, &accountKeeper, &bankKeeper, ctx, cdc, storeKey
}

func CfeclaimKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	k, _, _, ctx, _, _ := cfeclaimKeeperWithBlockHeightAndTime(t, 0, time.Now())
	return k, ctx

}
