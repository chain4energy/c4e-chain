package energychain

import (
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	energychainsimulation "github.com/chain4energy/c4e-chain/x/cfefingerprint/simulation"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
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
	_ = energychainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAccount = "op_weight_msg_create_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAccount int = 100

	opWeightMsgCreateReferencePayloadLink = "op_weight_msg_create_reference_payload_link"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateReferencePayloadLink int = 100

	opWeightMsgCreateNewAccount = "op_weight_msg_create_new_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateNewAccount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	energychainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&energychainGenesis)
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

	var weightMsgCreateAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAccount, &weightMsgCreateAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAccount = defaultWeightMsgCreateAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAccount,
		energychainsimulation.SimulateMsgCreateAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateReferencePayloadLink int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateReferencePayloadLink, &weightMsgCreateReferencePayloadLink, nil,
		func(_ *rand.Rand) {
			weightMsgCreateReferencePayloadLink = defaultWeightMsgCreateReferencePayloadLink
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateReferencePayloadLink,
		energychainsimulation.SimulateMsgCreateReferencePayloadLink(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateNewAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateNewAccount, &weightMsgCreateNewAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateNewAccount = defaultWeightMsgCreateNewAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateNewAccount,
		energychainsimulation.SimulateMsgCreateNewAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
