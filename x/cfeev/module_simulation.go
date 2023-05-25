package cfeev

import (
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfeevsimulation "github.com/chain4energy/c4e-chain/x/cfeev/simulation"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
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
	_ = cfeevsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgPublishEnergyTransferOffer = "op_weight_msg_publish_energy_transfer_offer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPublishEnergyTransferOffer int = 100

	opWeightMsgStartEnergyTransferRequest = "op_weight_msg_start_energy_transfer_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartEnergyTransferRequest int = 100

	opWeightMsgEnergyTransferStartedRequest = "op_weight_msg_energy_transfer_started_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEnergyTransferStartedRequest int = 100

	opWeightMsgEnergyTransferCompletedRequest = "op_weight_msg_energy_transfer_completed_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEnergyTransferCompletedRequest int = 100

	opWeightMsgCancelEnergyTransferRequest = "op_weight_msg_cancel_energy_transfer_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelEnergyTransferRequest int = 100

	opWeightMsgRemoveEnergyOffer = "op_weight_msg_remove_energy_offer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveEnergyOffer int = 100

	opWeightMsgRemoveTransfer = "op_weight_msg_remove_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveTransfer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cfeevGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfeevGenesis)
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

	var weightMsgPublishEnergyTransferOffer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPublishEnergyTransferOffer, &weightMsgPublishEnergyTransferOffer, nil,
		func(_ *rand.Rand) {
			weightMsgPublishEnergyTransferOffer = defaultWeightMsgPublishEnergyTransferOffer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPublishEnergyTransferOffer,
		cfeevsimulation.SimulateMsgPublishEnergyTransferOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStartEnergyTransferRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartEnergyTransferRequest, &weightMsgStartEnergyTransferRequest, nil,
		func(_ *rand.Rand) {
			weightMsgStartEnergyTransferRequest = defaultWeightMsgStartEnergyTransferRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartEnergyTransferRequest,
		cfeevsimulation.SimulateMsgStartEnergyTransferRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEnergyTransferStartedRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnergyTransferStartedRequest, &weightMsgEnergyTransferStartedRequest, nil,
		func(_ *rand.Rand) {
			weightMsgEnergyTransferStartedRequest = defaultWeightMsgEnergyTransferStartedRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnergyTransferStartedRequest,
		cfeevsimulation.SimulateMsgEnergyTransferStartedRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEnergyTransferCompletedRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnergyTransferCompletedRequest, &weightMsgEnergyTransferCompletedRequest, nil,
		func(_ *rand.Rand) {
			weightMsgEnergyTransferCompletedRequest = defaultWeightMsgEnergyTransferCompletedRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnergyTransferCompletedRequest,
		cfeevsimulation.SimulateMsgEnergyTransferCompletedRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelEnergyTransferRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelEnergyTransferRequest, &weightMsgCancelEnergyTransferRequest, nil,
		func(_ *rand.Rand) {
			weightMsgCancelEnergyTransferRequest = defaultWeightMsgCancelEnergyTransferRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelEnergyTransferRequest,
		cfeevsimulation.SimulateMsgCancelEnergyTransferRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveEnergyOffer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveEnergyOffer, &weightMsgRemoveEnergyOffer, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveEnergyOffer = defaultWeightMsgRemoveEnergyOffer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveEnergyOffer,
		cfeevsimulation.SimulateMsgRemoveEnergyOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveTransfer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveTransfer, &weightMsgRemoveTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveTransfer = defaultWeightMsgRemoveTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveTransfer,
		cfeevsimulation.SimulateMsgRemoveTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
