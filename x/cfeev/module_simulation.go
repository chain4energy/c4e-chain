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

	opWeightMsgStartEnergyTransfer = "op_weight_msg_start_energy_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartEnergyTransfer int = 100

	opWeightMsgEnergyTransferStarted = "op_weight_msg_energy_transfer_started"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEnergyTransferStarted int = 100

	opWeightMsgEnergyTransferCompleted = "op_weight_msg_energy_transfer_completed"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEnergyTransferCompleted int = 100

	opWeightMsgCancelEnergyTransfer = "op_weight_msg_cancel_energy_transfer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelEnergyTransfer int = 100

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

	var weightMsgStartEnergyTransfer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartEnergyTransfer, &weightMsgStartEnergyTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgStartEnergyTransfer = defaultWeightMsgStartEnergyTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartEnergyTransfer,
		cfeevsimulation.SimulateMsgStartEnergyTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEnergyTransferStarted int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnergyTransferStarted, &weightMsgEnergyTransferStarted, nil,
		func(_ *rand.Rand) {
			weightMsgEnergyTransferStarted = defaultWeightMsgEnergyTransferStarted
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnergyTransferStarted,
		cfeevsimulation.SimulateMsgEnergyTransferStarted(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEnergyTransferCompleted int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnergyTransferCompleted, &weightMsgEnergyTransferCompleted, nil,
		func(_ *rand.Rand) {
			weightMsgEnergyTransferCompleted = defaultWeightMsgEnergyTransferCompleted
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnergyTransferCompleted,
		cfeevsimulation.SimulateMsgEnergyTransferCompleted(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelEnergyTransfer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelEnergyTransfer, &weightMsgCancelEnergyTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgCancelEnergyTransfer = defaultWeightMsgCancelEnergyTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelEnergyTransfer,
		cfeevsimulation.SimulateMsgCancelEnergyTransfer(am.accountKeeper, am.bankKeeper, am.keeper),
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
