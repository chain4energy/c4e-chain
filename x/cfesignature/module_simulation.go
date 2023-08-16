package cfesignature

import (
	"math/rand"

	cfesignaturesimulation "github.com/chain4energy/c4e-chain/x/cfesignature/simulation"
	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

const (
	opWeightMsgStoreSignature = "op_weight_msg_store_signature"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStoreSignature int = 100

	opWeightMsgPublishReferencePayloadLink = "op_weight_msg_publish_reference_payload_link"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPublishReferencePayloadLink int = 100

	opWeightMsgCreateAccount = "op_weight_msg_create_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAccount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
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

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
	// No decoder
}

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
		cfesignaturesimulation.SimulateMsgStoreSignature(am.keeper),
	))

	var weightMsgPublishReferencePayloadLink int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPublishReferencePayloadLink, &weightMsgPublishReferencePayloadLink, nil,
		func(_ *rand.Rand) {
			weightMsgPublishReferencePayloadLink = defaultWeightMsgPublishReferencePayloadLink
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPublishReferencePayloadLink,
		cfesignaturesimulation.SimulateMsgPublishReferencePayloadLink(am.keeper),
	))

	var weightMsgCreateAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAccount, &weightMsgCreateAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAccount = defaultWeightMsgCreateAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAccount,
		cfesignaturesimulation.SimulateMsgCreateAccount(am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
