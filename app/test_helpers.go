package app

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	dbm "github.com/tendermint/tm-db"

	"encoding/json"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"


)

type SimApp interface {
	cosmoscmd.App
	GetApp() *App
	GetBaseApp() *baseapp.BaseApp
	AppCodec() codec.Codec
	SimulationManager() *module.SimulationManager
	ModuleAccountAddrs() map[string]bool
	Name() string
	LegacyAmino() *codec.LegacyAmino
	BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock
	EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock
	InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain
}

// Setup initializes a new chainforenergyApp
func Setup(isCheckTx bool) *App {
	db := dbm.NewMemDB()
	encoding := cosmoscmd.MakeEncodingConfig(ModuleBasics)

	app := New(
		log.TestingLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		0,
		encoding,
		simapp.EmptyAppOptions{},
	)

	if !isCheckTx {
		genesisState := NewDefaultGenesisState(encoding.Marshaler)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app.(SimApp).GetApp()
}

func AddMaccPerms(moduleAccountName string, perms []string) {
	maccPerms[moduleAccountName] = perms
}

func (app *App) GetApp() *App {
	return app

}
