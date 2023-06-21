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
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
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

	var weightMsgPublishEnergyTransferOffer = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPublishEnergyTransferOffer,
		cfeevsimulation.SimulateMsgPublishEnergyTransferOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStartEnergyTransfer = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartEnergyTransfer,
		cfeevsimulation.SimulateMsgStartEnergyTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	//var weightMsgEnergyTransferStarted int
	//operations = append(operations, simulation.NewWeightedOperation(
	//	weightMsgEnergyTransferStarted,
	//	cfeevsimulation.SimulateMsgEnergyTransferStarted(am.accountKeeper, am.bankKeeper, am.keeper),
	//))
	//
	//var weightMsgEnergyTransferCompleted int
	//operations = append(operations, simulation.NewWeightedOperation(
	//	weightMsgEnergyTransferCompleted,
	//	cfeevsimulation.SimulateMsgEnergyTransferCompleted(am.accountKeeper, am.bankKeeper, am.keeper),
	//))
	//
	//var weightMsgCancelEnergyTransfer int
	//operations = append(operations, simulation.NewWeightedOperation(
	//	weightMsgCancelEnergyTransfer,
	//	cfeevsimulation.SimulateMsgCancelEnergyTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	//))
	//
	var weightMsgRemoveEnergyOffer = 10
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveEnergyOffer,
		cfeevsimulation.SimulateMsgRemoveEnergyOffer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveTransfer = 10
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveTransfer,
		cfeevsimulation.SimulateMsgRemoveTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
