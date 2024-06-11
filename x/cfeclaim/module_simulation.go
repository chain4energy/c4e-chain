package cfeclaim

import (
	cfeclaimsimulation "github.com/chain4energy/c4e-chain/x/cfeclaim/simulation"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cfeclaimGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfeclaimGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgClaim = 100
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaim,
		cfeclaimsimulation.SimulateMsgClaim(am.keeper, am.accountKeeper, am.bankKeeper, am.cfevestingKeeper),
	))

	var weightMsgCreateCampaign = 10
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCampaign,
		cfeclaimsimulation.SimulateMsgCreateCampaign(am.keeper, am.accountKeeper, am.bankKeeper, am.cfevestingKeeper),
	))
	//
	var weightMsgAddMission = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddMission,
		cfeclaimsimulation.SimulateMsgAddMission(am.keeper, am.accountKeeper, am.bankKeeper, am.cfevestingKeeper),
	))

	var weightMsgAddClaimRecords = 100
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddClaimRecords,
		cfeclaimsimulation.SimulateMsgAddClaimRecords(am.keeper, am.accountKeeper, am.bankKeeper, am.cfevestingKeeper),
	))

	var weightMsgDeleteClaimRecord = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteClaimRecord,
		cfeclaimsimulation.SimulateMsgDeleteClaimRecord(am.keeper, am.accountKeeper, am.bankKeeper, am.cfevestingKeeper),
	))

	var weightMsgCloseCampaign = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCloseCampaign,
		cfeclaimsimulation.SimulateMsgCloseCampaign(am.keeper, am.accountKeeper, am.bankKeeper, am.cfevestingKeeper),
	))

	var weightMsgEnableCampaign = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnableCampaign,
		cfeclaimsimulation.SimulateMsgEnableCampaign(am.keeper, am.accountKeeper, am.bankKeeper, am.cfevestingKeeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
