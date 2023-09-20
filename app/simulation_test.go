package app

import (
	"encoding/json"
	"fmt"
	"github.com/CosmWasm/wasmd/app"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	cfedistributortypes "github.com/chain4energy/c4e-chain/v2/x/cfedistributor/types"
	cfemintertypes "github.com/chain4energy/c4e-chain/v2/x/cfeminter/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	dbm "github.com/cometbft/cometbft-db"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/stretchr/testify/require"
	"os"
	"testing"

	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

var emptyWasmOpts []wasmkeeper.Option

// Get flags every time the simulator is run
func init() {
	simcli.GetSimulatorFlags()
}

type StoreKeysPrefixes struct {
	A        storetypes.StoreKey
	B        storetypes.StoreKey
	Prefixes [][]byte
}

// BenchmarkSimulation run the chain simulation
// Running as go benchmark test:
func BenchmarkSimulation(b *testing.B) {
	_, _, cleanup := setupSimulation(b, "goleveldb-app-sim", "Simulation")

	defer func() {
		cleanup()
	}()
}

func BenchmarkSimTest(b *testing.B) {
	app, _, cleanup1 := setupSimulation(b, "goleveldb-app-sim", "Simulation")

	fmt.Printf("exporting genesis...\n")
	exported, err := app.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(b, err)

	fmt.Printf("importing genesis...\n")

	var genesisState GenesisState

	err = json.Unmarshal(exported.AppState, &genesisState)
	require.NoError(b, err)

	newApp, _, _, _, _, cleanup2 := BaseSimulationSetup(b, "goleveldb-app-sim-2", "Simulation-2")

	defer func() {
		cleanup1()
		cleanup2()
	}()

	ctxA := app.NewContext(true, tmproto.Header{Height: app.LastBlockHeight()})
	ctxB := newApp.NewContext(true, tmproto.Header{Height: app.LastBlockHeight()})

	newApp.mm.InitGenesis(ctxB, app.AppCodec(), genesisState)
	newApp.StoreConsensusParams(ctxB, exported.ConsensusParams)

	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{app.keys[authtypes.StoreKey], newApp.keys[authtypes.StoreKey], [][]byte{}},
		{
			app.keys[stakingtypes.StoreKey], newApp.keys[stakingtypes.StoreKey],
			[][]byte{
				stakingtypes.UnbondingQueueKey, stakingtypes.RedelegationQueueKey, stakingtypes.ValidatorQueueKey,
				stakingtypes.HistoricalInfoKey,
			},
		},
		{app.keys[slashingtypes.StoreKey], newApp.keys[slashingtypes.StoreKey], [][]byte{}},
		{app.keys[distrtypes.StoreKey], newApp.keys[distrtypes.StoreKey], [][]byte{}},
		{app.keys[banktypes.StoreKey], newApp.keys[banktypes.StoreKey], [][]byte{banktypes.BalancesPrefix}},
		{app.keys[paramtypes.StoreKey], newApp.keys[paramtypes.StoreKey], [][]byte{}},
		{app.keys[govtypes.StoreKey], newApp.keys[govtypes.StoreKey], [][]byte{}},
		{app.keys[evidencetypes.StoreKey], newApp.keys[evidencetypes.StoreKey], [][]byte{}},
		{app.keys[capabilitytypes.StoreKey], newApp.keys[capabilitytypes.StoreKey], [][]byte{}},
		{app.keys[authzkeeper.StoreKey], newApp.keys[authzkeeper.StoreKey], [][]byte{authzkeeper.GrantKey, authzkeeper.GrantQueuePrefix}},

		// IBC
		{app.keys[ibctransfertypes.StoreKey], newApp.keys[ibctransfertypes.StoreKey], [][]byte{}},

		// OUR MODULES
		{app.keys[cfevestingtypes.StoreKey], newApp.keys[cfevestingtypes.StoreKey], [][]byte{}},
		{app.keys[cfedistributortypes.StoreKey], newApp.keys[cfedistributortypes.StoreKey], [][]byte{}},
		{app.keys[cfemintertypes.StoreKey], newApp.keys[cfemintertypes.StoreKey], [][]byte{}},
		{app.keys[cfeclaimtypes.StoreKey], newApp.keys[cfeclaimtypes.StoreKey], [][]byte{}},
	}

	for _, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(b, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf("compared %d different key/value pairs between %s and %s\n", len(failedKVAs), skp.A, skp.B)
		simLog := simtestutil.GetSimulationLog(
			skp.A.Name(),
			app.SimulationManager().StoreDecoders,
			failedKVAs,
			failedKVBs,
		)
		require.Equal(b, len(failedKVAs), 0, simLog)
	}
}

func setupSimulation(tb testing.TB, dirPrevix string, dbName string) (c4eapp2 *App, simParams simulation.Params, cleanup func()) {
	c4eapp, _, config, db, _, cleanup := BaseSimulationSetup(tb, dirPrevix, dbName)

	weightedOperations := simtestutil.SimulationOperations(c4eapp, c4eapp.AppCodec(), config)
	defer func() {
		if r := recover(); r != nil {
			cleanup()
		}
	}()
	_, simParams, simErr := simulation.SimulateFromSeed(
		tb,
		os.Stdout,
		c4eapp.BaseApp,
		simtestutil.AppStateFn(
			c4eapp.AppCodec(),
			c4eapp.SimulationManager(),
			app.NewDefaultGenesisState(c4eapp.AppCodec()),
		),
		simulationtypes.RandomAccounts,
		weightedOperations,
		app.ModuleAccountAddrs(),
		config,
		c4eapp.AppCodec(),
	)
	require.NoError(tb, simErr)

	err := simtestutil.CheckExportSimulation(c4eapp, config, simParams)

	require.NoError(tb, err)
	if config.Commit {
		simtestutil.PrintStats(db)
	}

	return c4eapp, simParams, cleanup
}

func BaseSimulationSetup(tb testing.TB, dirPrevix string, dbName string) (*App, GenesisState, simulationtypes.Config, dbm.DB, string, func()) {
	config := simcli.NewConfigFromFlags()
	config.ChainID = "c4e-chain"
	db, dir, logger, _, err := simtestutil.SetupSimulation(config, dirPrevix, dbName, simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	require.NoError(tb, err, "simulation setup failed")

	encoding := MakeEncodingConfig()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = app.DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue
	c4eapp := New(
		logger,
		db,
		nil,
		true,
		wasmtypes.EnableAllProposals,
		map[int64]bool{},
		DefaultNodeHome,
		simcli.FlagPeriodValue,
		encoding,
		appOptions,
		emptyWasmOpts,
	)
	genesisState := NewDefaultGenesisState(encoding.Codec)
	cleanup := func() {
		require.NoError(tb, db.Close())
		require.NoError(tb, os.RemoveAll(dir))
	}

	return c4eapp, genesisState, config, db, dir, cleanup
}
