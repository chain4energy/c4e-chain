package cfeclaim

import (
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfeclaimsimulation "github.com/chain4energy/c4e-chain/x/cfeclaim/simulation"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
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
	cfeclaimGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfeclaimGenesis)
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

	var weightMsgClaim = 100
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaim,
		cfeclaimsimulation.SimulateMsgClaim(am.keeper, am.cfevestingKeeper),
	))

	var weightMsgCreateCampaign = 10
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCampaign,
		cfeclaimsimulation.SimulateMsgCreateCampaign(am.keeper, am.cfevestingKeeper),
	))

	var weightMsgAddMissionToCampaign = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddMissionToCampaign,
		cfeclaimsimulation.SimulateMsgAddMissionToCampaign(am.keeper, am.cfevestingKeeper),
	))

	var weightMsgAddClaimRecords = 100
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddClaimRecords,
		cfeclaimsimulation.SimulateMsgAddClaimRecords(am.keeper, am.cfevestingKeeper),
	))

	var weightMsgDeleteClaimRecord = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteClaimRecord,
		cfeclaimsimulation.SimulateMsgDeleteClaimRecord(am.keeper, am.cfevestingKeeper),
	))

	var weightMsgCloseCampaign = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCloseCampaign,
		cfeclaimsimulation.SimulateMsgCloseCampaign(am.keeper, am.cfevestingKeeper),
	))

	var weightMsgEnableCampaign = 20
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnableCampaign,
		cfeclaimsimulation.SimulateMsgEnableCampaign(am.keeper, am.cfevestingKeeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
