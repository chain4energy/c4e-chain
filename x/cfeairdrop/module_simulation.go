package cfeairdrop

import (
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfeairdropsimulation "github.com/chain4energy/c4e-chain/x/cfeairdrop/simulation"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
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
	_ = cfeairdropsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgClaim = "op_weight_msg_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaim int = 100

	opWeightMsgCreateAirdropCampaign = "op_weight_msg_create_airdrop_campaign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAirdropCampaign int = 100

	opWeightMsgAddMissionToAidropCampaign = "op_weight_msg_add_mission_to_aidrop_campaign"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddMissionToAidropCampaign int = 100

	opWeightMsgCreateAirdropEntry = "op_weight_msg_airdrop_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAirdropEntry int = 100

	opWeightMsgUpdateAirdropEntry = "op_weight_msg_airdrop_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAirdropEntry int = 100

	opWeightMsgDeleteAirdropEntry = "op_weight_msg_airdrop_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAirdropEntry int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cfeairdropGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AirdropEntryList: []types.AirdropEntry{
			{
				Address: sample.AccAddress(),
			},
			{
				Address: sample.AccAddress(),
			},
		},
		AirdropEntryCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfeairdropGenesis)
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
		cfeairdropsimulation.SimulateMsgClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateAirdropCampaign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAirdropCampaign, &weightMsgCreateAirdropCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAirdropCampaign = defaultWeightMsgCreateAirdropCampaign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAirdropCampaign,
		cfeairdropsimulation.SimulateMsgCreateAirdropCampaign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddMissionToAidropCampaign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddMissionToAidropCampaign, &weightMsgAddMissionToAidropCampaign, nil,
		func(_ *rand.Rand) {
			weightMsgAddMissionToAidropCampaign = defaultWeightMsgAddMissionToAidropCampaign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddMissionToAidropCampaign,
		cfeairdropsimulation.SimulateMsgAddMissionToAidropCampaign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateAirdropEntry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAirdropEntry, &weightMsgCreateAirdropEntry, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAirdropEntry = defaultWeightMsgCreateAirdropEntry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAirdropEntry,
		cfeairdropsimulation.SimulateMsgCreateAirdropEntry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAirdropEntry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAirdropEntry, &weightMsgUpdateAirdropEntry, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAirdropEntry = defaultWeightMsgUpdateAirdropEntry
		},
	)

	var weightMsgDeleteAirdropEntry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAirdropEntry, &weightMsgDeleteAirdropEntry, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAirdropEntry = defaultWeightMsgDeleteAirdropEntry
		},
	)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
