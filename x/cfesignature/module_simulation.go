package cfesignature

import (
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfesignaturesimulation "github.com/chain4energy/c4e-chain/x/cfesignature/simulation"
	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
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
	_ = cfesignaturesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgStoreSignature = "op_weight_msg_store_signature"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStoreSignature int = 100

	opWeightMsgPublishReferencePayloadLink = "op_weight_msg_publish_reference_payload_link"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPublishReferencePayloadLink int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cfesignatureGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfesignatureGenesis)
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

	var weightMsgStoreSignature int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStoreSignature, &weightMsgStoreSignature, nil,
		func(_ *rand.Rand) {
			weightMsgStoreSignature = defaultWeightMsgStoreSignature
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStoreSignature,
		cfesignaturesimulation.SimulateMsgStoreSignature(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPublishReferencePayloadLink int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPublishReferencePayloadLink, &weightMsgPublishReferencePayloadLink, nil,
		func(_ *rand.Rand) {
			weightMsgPublishReferencePayloadLink = defaultWeightMsgPublishReferencePayloadLink
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPublishReferencePayloadLink,
		cfesignaturesimulation.SimulateMsgPublishReferencePayloadLink(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
