package app

import (
	"time"
	"github.com/cosmos/cosmos-sdk/simapp"
	dbm "github.com/tendermint/tm-db"

	"encoding/json"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"

	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	c4eapp "github.com/chain4energy/c4e-chain/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

)

// Setup initializes a new chainforenergyApp
// func Setup(isCheckTx bool) *c4eapp.App {
// 	db := dbm.NewMemDB()

// 	encoding := cosmoscmd.MakeEncodingConfig(c4eapp.ModuleBasics)

// 	app := c4eapp.New(
// 		log.TestingLogger(),
// 		db,
// 		nil,
// 		true,
// 		map[int64]bool{},
// 		c4eapp.DefaultNodeHome,
// 		0,
// 		encoding,
// 		simapp.EmptyAppOptions{},
// 	)

// 	if !isCheckTx {
// 		genesisState := c4eapp.NewDefaultGenesisState(encoding.Marshaler)
// 		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
// 		if err != nil {
// 			panic(err)
// 		}

// 		app.InitChain(
// 			abci.RequestInitChain{
// 				Validators:      []abci.ValidatorUpdate{},
// 				ConsensusParams: simapp.DefaultConsensusParams,
// 				AppStateBytes:   stateBytes,
// 			},
// 		)
// 	}
// 	return app.(*c4eapp.App)
// }












// Setup initializes a new chainforenergyApp
func Setup(isCheckTx bool) *c4eapp.App {
	app, genesisState := BaseSetup()
	if !isCheckTx {
		AddValidatorToAppGenesis(app, genesisState)
	}
	return app
}

func BaseSetup() (*c4eapp.App, c4eapp.GenesisState) {
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
	genesisState := c4eapp.NewDefaultGenesisState(encoding.Marshaler)

	return app.(*c4eapp.App), genesisState
}

func AddValidatorToAppGenesis(app *c4eapp.App, genesisState c4eapp.GenesisState) {
	valPrivKey := secp256k1.GenPrivKey()  // TODO to PV after migration to cosmos sdk 0.46
	pubKey, _ := cryptocodec.ToTmPubKeyInterface(valPrivKey.PubKey())
	validator := &tmtypes.Validator{
		Address:          pubKey.Address(),
		PubKey:           pubKey,
		VotingPower:      1,
		ProposerPriority: 0,
	}

	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
	}

	genesisState = genesisStateWithValSet(app, genesisState, valSet, []authtypes.GenesisAccount{acc}, balance)
	stateBytes, _ := json.MarshalIndent(genesisState, "", " ")
	app.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simapp.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)
}

func genesisStateWithValSet(
	app *c4eapp.App, genesisState c4eapp.GenesisState,
	valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) c4eapp.GenesisState {
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, _ := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		pkAny, _ := codectypes.NewAnyWithValue(pk)

		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdk.OneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			MinSelfDelegation: sdk.ZeroInt(),
		}

		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))
	}

	stakingGenesis := stakingtypes.NewGenesisState(stakingtypes.DefaultParams(), validators, delegations)
	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		totalSupply = totalSupply.Add(b.Coins...)
	}

	for range delegations {
		totalSupply = totalSupply.Add(sdk.NewCoin(sdk.DefaultBondDenom, bondAmt))
	}

	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, bondAmt)},
	})

	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	return genesisState
}
