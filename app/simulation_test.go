package app

import (
	"encoding/json"
	"fmt"
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfesignaturetypes "github.com/chain4energy/c4e-chain/x/cfesignature/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

func init() {
	simapp.GetSimulatorFlags()
}

type StoreKeysPrefixes struct {
	A        storetypes.StoreKey
	B        storetypes.StoreKey
	Prefixes [][]byte
}

var RegexExcludedOperations = regexp.MustCompile(`authz`)

// BenchmarkSimulation run the chain simulation
// Running as go benchmark test:
func BenchmarkSimulation(b *testing.B) {
	_, _ = setupSimulation(b, "goleveldb-app-sim", "Simulation")
}

func BenchmarkSimTest(b *testing.B) {
	c4eapp1, _ := setupSimulation(b, "goleveldb-app-sim", "Simulation")

	fmt.Printf("exporting genesis...\n")
	exported, err := c4eapp1.ExportAppStateAndValidators(false, []string{})
	require.NoError(b, err)

	fmt.Printf("importing genesis...\n")

	var genesisState GenesisState
	fmt.Println(genesisState)
	err = json.Unmarshal(exported.AppState, &genesisState)
	require.NoError(b, err)

	c4eapp2, _, _, _, _ := BaseSimulationSetup(b, "goleveldb-app-sim-2", "Simulation-2")
	ctxA := c4eapp1.NewContext(true, tmproto.Header{Height: c4eapp1.LastBlockHeight()})
	ctxB := c4eapp2.NewContext(true, tmproto.Header{Height: c4eapp1.LastBlockHeight()})
	c4eapp2.mm.InitGenesis(ctxB, c4eapp1.AppCodec(), genesisState)
	c4eapp2.StoreConsensusParams(ctxB, exported.ConsensusParams)

	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{c4eapp1.keys[authtypes.StoreKey], c4eapp2.keys[authtypes.StoreKey], [][]byte{}},
		{
			c4eapp1.keys[stakingtypes.StoreKey], c4eapp2.keys[stakingtypes.StoreKey],
			[][]byte{
				stakingtypes.UnbondingQueueKey, stakingtypes.RedelegationQueueKey, stakingtypes.ValidatorQueueKey,
				stakingtypes.HistoricalInfoKey,
			},
		},
		{c4eapp1.keys[slashingtypes.StoreKey], c4eapp2.keys[slashingtypes.StoreKey], [][]byte{}},
		{c4eapp1.keys[distrtypes.StoreKey], c4eapp2.keys[distrtypes.StoreKey], [][]byte{}},
		{c4eapp1.keys[banktypes.StoreKey], c4eapp2.keys[banktypes.StoreKey], [][]byte{banktypes.BalancesPrefix}},
		{c4eapp1.keys[paramtypes.StoreKey], c4eapp2.keys[paramtypes.StoreKey], [][]byte{}},
		{c4eapp1.keys[govtypes.StoreKey], c4eapp2.keys[govtypes.StoreKey], [][]byte{}},
		{c4eapp1.keys[evidencetypes.StoreKey], c4eapp2.keys[evidencetypes.StoreKey], [][]byte{}},
		{c4eapp1.keys[capabilitytypes.StoreKey], c4eapp2.keys[capabilitytypes.StoreKey], [][]byte{}},
		{c4eapp1.keys[authzkeeper.StoreKey], c4eapp2.keys[authzkeeper.StoreKey], [][]byte{}},

		// IBC
		{c4eapp1.keys[ibchost.StoreKey], c4eapp2.keys[ibchost.StoreKey], [][]byte{}},
		{c4eapp1.keys[ibctransfertypes.StoreKey], c4eapp2.keys[ibctransfertypes.StoreKey], [][]byte{}},

		// OUR MODULES
		{c4eapp1.keys[cfevestingtypes.StoreKey], c4eapp2.keys[cfevestingtypes.StoreKey], [][]byte{
			cfevestingtypes.AccountVestingPoolsKeyPrefix, cfevestingtypes.VestingTypesKeyPrefix, cfevestingtypes.ParamsKey,
		}},
		{c4eapp1.keys[cfedistributortypes.StoreKey], c4eapp2.keys[cfedistributortypes.StoreKey], [][]byte{
			cfedistributortypes.StateKeyPrefix,
		}},
		{c4eapp1.keys[cfemintertypes.StoreKey], c4eapp2.keys[cfemintertypes.StoreKey], [][]byte{
			cfemintertypes.MinterStateHistoryKeyPrefix, cfemintertypes.IsGenesisKey, cfemintertypes.MinterStateKey,
		}},
		{c4eapp1.keys[cfesignaturetypes.StoreKey], c4eapp2.keys[cfesignaturetypes.StoreKey], [][]byte{}},
	}

	for _, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(b, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf("compared %d different key/value pairs between %s and %s\n", len(failedKVAs), skp.A, skp.B)
		require.Equal(b, len(failedKVAs), 0, simapp.GetSimulationLog(skp.A.Name(), c4eapp1.SimulationManager().StoreDecoders, failedKVAs, failedKVBs))
	}
}

func setupSimulation(tb testing.TB, dirPrevix string, dbName string) (c4eapp *App, simParams simulation.Params) {
	app, _, config, db, dir := BaseSimulationSetup(tb, dirPrevix, dbName)

	defer func() {
		err := db.Close()
		require.NoError(tb, err)
		err = os.RemoveAll(dir)
		require.NoError(tb, err)
	}()

	weightedOperations := simapp.SimulationOperations(app, app.AppCodec(), config)

	for i, operation := range weightedOperations {
		operationName := runtime.FuncForPC(reflect.ValueOf(operation.Op()).Pointer()).Name()
		if RegexExcludedOperations.MatchString(operationName) {
			weightedOperations = append(weightedOperations[:i], weightedOperations[i+1:]...)
			fmt.Println("Excluded operation: " + operationName)
		}
	}

	_, simParams, simErr := simulation.SimulateFromSeed(
		tb,
		os.Stdout,
		app.BaseApp,
		simapp.AppStateFn(app.AppCodec(), app.SimulationManager()),
		simulationtypes.RandomAccounts,
		weightedOperations,
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)

	err := simapp.CheckExportSimulation(app, config, simParams)
	require.NoError(tb, err)
	require.NoError(tb, simErr)
	if config.Commit {
		simapp.PrintStats(db)
	}
	return app, simParams
}

func BaseSimulationSetup(tb testing.TB, dirPrevix string, dbName string) (*App, GenesisState, simulationtypes.Config, dbm.DB, string) {
	config, db, dir, _, _, err := simapp.SetupSimulation(dirPrevix, dbName)
	require.NoError(tb, err, "simulation setup failed")

	encoding := MakeEncodingConfig()
	app := New(
		log.TestingLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		20,
		encoding,
		simapp.EmptyAppOptions{},
	)
	genesisState := NewDefaultGenesisState(encoding.Marshaler)

	return app, genesisState, config, db, dir

}
