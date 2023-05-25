package keeper

import (
	"fmt"
	"testing"

	"github.com/chain4energy/c4e-chain/app"
	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
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

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	cfeevmoduletypes "github.com/chain4energy/c4e-chain/x/cfeev/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
)

func CfeevKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("c4e", "c4e"+"pub")
	config.Seal()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	// auth module store key
	authStoreKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	// bank module store key
	bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(authStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(bankStoreKey, storetypes.StoreTypeIAVL, db)

	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	encodingConfig := app.MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CfeevParams",
	)

	// create account keeper
	var accountKeeper banktypes.AccountKeeper
	accountKeeper, _ = createAccountKeeper(appCodec, authStoreKey, memStoreKey)

	// create bank keeper
	var bankKeeper types.BankKeeper
	bankKeeper, _ = createBankKeeper(appCodec, bankStoreKey, accountKeeper, memStoreKey, storeKey)
	_ = bankKeeper

	// the app's keeper
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		bankKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	// Populate the cfeev account with some coins
	coins := sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(100000)))
	err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	err = bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.ModuleName, coins)

	if err != nil {
		panic(fmt.Errorf("Minting initial coins failed %e", err))
	}

	return k, ctx
}

func createBankKeeper(codec codec.Codec, bankStoreKey storetypes.StoreKey, accountKeeper banktypes.AccountKeeper, memStoreKey *storetypes.MemoryStoreKey, storeKey *storetypes.KVStoreKey) (types.BankKeeper, error) {
	var bankKeeper types.BankKeeper

	paramsSubspace := typesparams.NewSubspace(codec,
		types.Amino,
		storeKey,
		memStoreKey,
		"CfeevchainParams",
	)

	bankKeeper = bankkeeper.NewBaseKeeper(
		codec,
		bankStoreKey,
		accountKeeper,
		// app.GetSubspace(banktypes.ModuleName),
		paramsSubspace,
		// app.BlockedModuleAccountAddrs(),
		nil,
	)

	return bankKeeper, nil

}

func createAccountKeeper(codec codec.Codec, authStoreKey storetypes.StoreKey, memStoreKey *storetypes.MemoryStoreKey) (banktypes.AccountKeeper, error) {
	var accountKeeper banktypes.AccountKeeper

	paramsAuthSubspace := typesparams.NewSubspace(codec,
		types.Amino,
		authStoreKey,
		memStoreKey,
		"CfeevchainParams",
	)

	// module account permissions
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		icatypes.ModuleName:            nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		cfeevmoduletypes.ModuleName:    {authtypes.Minter, authtypes.Burner, authtypes.Staking},
	}

	accountKeeper = authkeeper.NewAccountKeeper(
		codec,
		authStoreKey,
		paramsAuthSubspace,
		authtypes.ProtoBaseAccount,
		maccPerms,
		// sdk.Bech32PrefixAccAddr,
		"c4e",
	)

	return accountKeeper, nil
}
