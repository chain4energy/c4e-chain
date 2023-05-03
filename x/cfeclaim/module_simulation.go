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

const (
	opWeightMsgClaim = "op_weight_msg_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaim int = 400

	opWeightMsgCreateCampaign = "op_weight_msg_create_claim_campaign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCampaign int = 100

	opWeightMsgAddMissionToAidropCampaign = "op_weight_msg_add_mission_to_aidrop_campaign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddMissionToAidropCampaign int = 100

	opWeightMsgCreateEntry = "op_weight_msg_claim_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEntry int = 100

	opWeightMsgUpdateEntry = "op_weight_msg_claim_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEntry int = 100

	opWeightMsgDeleteClaimRecord = "op_weight_msg_claim_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteClaimRecord int = 100

	opWeightMsgCloseCampaign = "op_weight_msg_close_claim_campaign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCloseCampaign int = 100

	opWeightMsgStartCampaign = "op_weight_msg_start_claim_campaign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartCampaign int = 100

	// this line is used by starport scaffolding # simapp/module/const
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

	var weightMsgClaim int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaim, &weightMsgClaim, nil,
		func(_ *rand.Rand) {
			weightMsgClaim = defaultWeightMsgClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaim,
		cfeclaimsimulation.SimulateMsgClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateCampaign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCampaign, &weightMsgCreateCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCampaign = defaultWeightMsgCreateCampaign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCampaign,
		cfeclaimsimulation.SimulateMsgCreateCampaign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddMissionToAidropCampaign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddMissionToAidropCampaign, &weightMsgAddMissionToAidropCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgAddMissionToAidropCampaign = defaultWeightMsgAddMissionToAidropCampaign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddMissionToAidropCampaign,
		cfeclaimsimulation.SimulateMsgAddMissionToAidropCampaign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateEntry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateEntry, &weightMsgCreateEntry, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEntry = defaultWeightMsgCreateEntry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEntry,
		cfeclaimsimulation.SimulateMsgAddClaimRecords(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEntry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateEntry, &weightMsgUpdateEntry, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEntry = defaultWeightMsgUpdateEntry
		},
	)

	var weightMsgDeleteClaimRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteClaimRecord, &weightMsgDeleteClaimRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteClaimRecord = defaultWeightMsgDeleteClaimRecord
		},
	)

	var weightMsgCloseCampaign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCloseCampaign, &weightMsgCloseCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgCloseCampaign = defaultWeightMsgCloseCampaign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCloseCampaign,
		cfeclaimsimulation.SimulateMsgCloseCampaign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStartCampaign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartCampaign, &weightMsgStartCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgStartCampaign = defaultWeightMsgStartCampaign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartCampaign,
		cfeclaimsimulation.SimulateMsgStartCampaign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
