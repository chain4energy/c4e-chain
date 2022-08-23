package cfeenergybank

import (
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfeenergybanksimulation "github.com/chain4energy/c4e-chain/x/cfeenergybank/simulation"
	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cfeenergybanksimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateTokenParams = "op_weight_msg_create_token_params"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTokenParams int = 100

	opWeightMsgMintToken = "op_weight_msg_mint_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintToken int = 100

	opWeightMsgTransferTokens = "op_weight_msg_transfer_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferTokens int = 100

	opWeightMsgTransferTokensOptimally = "op_weight_msg_transfer_tokens_optimally"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferTokensOptimally int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cfeenergybankGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfeenergybankGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTokenParams int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTokenParams, &weightMsgCreateTokenParams, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTokenParams = defaultWeightMsgCreateTokenParams
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTokenParams,
		cfeenergybanksimulation.SimulateMsgCreateTokenParams(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintToken, &weightMsgMintToken, nil,
		func(_ *rand.Rand) {
			weightMsgMintToken = defaultWeightMsgMintToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintToken,
		cfeenergybanksimulation.SimulateMsgMintToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferTokens int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferTokens, &weightMsgTransferTokens, nil,
		func(_ *rand.Rand) {
			weightMsgTransferTokens = defaultWeightMsgTransferTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferTokens,
		cfeenergybanksimulation.SimulateMsgTransferTokens(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferTokensOptimally int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferTokensOptimally, &weightMsgTransferTokensOptimally, nil,
		func(_ *rand.Rand) {
			weightMsgTransferTokensOptimally = defaultWeightMsgTransferTokensOptimally
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferTokensOptimally,
		cfeenergybanksimulation.SimulateMsgTransferTokensOptimally(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
