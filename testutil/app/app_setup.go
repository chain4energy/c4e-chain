package app

import (
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	dbm "github.com/tendermint/tm-db"

	"encoding/json"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"

	c4eapp "github.com/chain4energy/c4e-chain/app"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"

	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Setup initializes a new chainforenergyApp
func SetupWithBondDenom(isCheckTx bool, bondDenom string) (*c4eapp.App, sdk.Coin) {
	return SetupWithValidatorsAmount(isCheckTx, bondDenom, 1)
}

func SetupWithValidatorsAmount(isCheckTx bool, bondDenom string, validatorsAmount int, balances ...banktypes.Balance) (*c4eapp.App, sdk.Coin) {
	app, genesisState := BaseSetup()
	if !isCheckTx {
		coin := AddValidatorsToAppGenesis(app, genesisState, bondDenom, createValidatorSet(validatorsAmount), balances...)
		return app, coin
	}
	return app, sdk.NewCoin(bondDenom, sdk.ZeroInt())
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

func createValidatorSet(validatorsAmount int) *tmtypes.ValidatorSet {
	var vals []*tmtypes.Validator

	for i := 0; i < validatorsAmount; i++ {
		valPrivKey := secp256k1.GenPrivKey() // TODO to PV after migration to cosmos sdk 0.46
		pubKey, _ := cryptocodec.ToTmPubKeyInterface(valPrivKey.PubKey())
		validator := &tmtypes.Validator{
			Address:          pubKey.Address(),
			PubKey:           pubKey,
			VotingPower:      1,
			ProposerPriority: 0,
		}
		vals = append(vals, validator)
	}

	return tmtypes.NewValidatorSet(vals)
}

func AddValidatorsToAppGenesis(app *c4eapp.App, genesisState c4eapp.GenesisState, bondDenom string, valSet *tmtypes.ValidatorSet, balances ...banktypes.Balance) sdk.Coin {
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	senderCoin := sdk.NewCoin(bondDenom, sdk.NewInt(1000000))
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(senderCoin),
	}

	var allbalances []banktypes.Balance
	allbalances = append(allbalances, balance)
	allbalances = append(allbalances, balances...)
	genesisState, delegated := genesisStateWithValSet(app, genesisState, bondDenom, valSet, []authtypes.GenesisAccount{acc}, allbalances...)
	stateBytes, _ := json.MarshalIndent(genesisState, "", " ")
	app.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simapp.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	app.Commit()
	app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
		Height:             app.LastBlockHeight() + 1,
		AppHash:            app.LastCommitID().Hash,
		ValidatorsHash:     valSet.Hash(),
		NextValidatorsHash: valSet.Hash(),
	}})

	return delegated.Add(senderCoin)
}

func genesisStateWithValSet(
	app *c4eapp.App, genesisState c4eapp.GenesisState, bondDenom string,
	valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) (c4eapp.GenesisState, sdk.Coin) {
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
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = bondDenom
	stakingGenesis := stakingtypes.NewGenesisState(stakingParams, validators, delegations)
	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		totalSupply = totalSupply.Add(b.Coins...)
	}

	delegationsSum := sdk.NewCoin(bondDenom, sdk.ZeroInt())
	for range delegations {
		totalSupply = totalSupply.Add(sdk.NewCoin(bondDenom, bondAmt))
		delegationsSum = delegationsSum.Add(sdk.NewCoin(bondDenom, bondAmt))

	}

	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(bondDenom, bondAmt)},
	})

	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	vestingGenesis := cfevestingtypes.DefaultGenesis()
	vestingGenesis.Params.Denom = testenv.DefaultTestDenom
	genesisState[cfevestingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(vestingGenesis)

	distributorGenesis := cfedistributortypes.DefaultGenesis()
	distributorGenesis.Params.SubDistributors[0].Destinations.PrimaryShare.Id = testenv.DefaultDistributionDestination
	genesisState[cfedistributortypes.ModuleName] = app.AppCodec().MustMarshalJSON(distributorGenesis)

	govGenesis := govtypes.DefaultGenesisState()
	govGenesis.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, govtypes.DefaultMinDepositTokens))
	genesisState[govtypes.ModuleName] = app.AppCodec().MustMarshalJSON(govGenesis)

	return genesisState, delegationsSum
}
