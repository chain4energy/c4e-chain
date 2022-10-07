package app

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	dbm "github.com/tendermint/tm-db"

	"encoding/json"

	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	c4eapp "github.com/chain4energy/c4e-chain/app"
)

// Setup initializes a new chainforenergyApp
func Setup(isCheckTx bool) *c4eapp.App {
	db := dbm.NewMemDB()

	encoding := cosmoscmd.MakeEncodingConfig(c4eapp.ModuleBasics)

	app := c4eapp.New(
		log.TestingLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		c4eapp.DefaultNodeHome,
		0,
		encoding,
		simapp.EmptyAppOptions{},
	)

	if !isCheckTx {
		genesisState := c4eapp.NewDefaultGenesisState(encoding.Marshaler)
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
	return app.(*c4eapp.App)
}



