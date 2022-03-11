package app

import (
	"bytes"
	"strconv"

	"github.com/cosmos/cosmos-sdk/simapp"
	dbm "github.com/tendermint/tm-db"

	"encoding/hex"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
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

// CreateTestPubKeys returns a total of numPubKeys public keys in ascending order.
func CreateTestPubKeys(numPubKeys int) []cryptotypes.PubKey {
	var publicKeys []cryptotypes.PubKey
	var buffer bytes.Buffer

	// start at 10 to avoid changing 1 to 01, 2 to 02, etc
	for i := 100; i < (numPubKeys + 100); i++ {
		numString := strconv.Itoa(i)
		buffer.WriteString("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AF") // base pubkey string
		buffer.WriteString(numString)                                                       // adding on final two digits to make pubkeys unique
		publicKeys = append(publicKeys, NewPubKeyFromHex(buffer.String()))
		buffer.Reset()
	}

	return publicKeys
}

// NewPubKeyFromHex returns a PubKey from a hex string.
func NewPubKeyFromHex(pk string) (res cryptotypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}
	if len(pkBytes) != ed25519.PubKeySize {
		panic(errors.Wrap(errors.ErrInvalidPubKey, "invalid pubkey size"))
	}
	return &ed25519.PubKey{Key: pkBytes}
}

func AddMaccPerms(moduleAccountName string, perms []string) {
	maccPerms[moduleAccountName] = perms
}

func (app *App) GetApp() *App {
	return app

}
