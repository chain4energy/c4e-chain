package cfetokenization

import (
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfetokenizationsimulation "github.com/chain4energy/c4e-chain/x/cfetokenization/simulation"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
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
	_ = cfetokenizationsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateUserDevices = "op_weight_msg_user_devices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUserDevices int = 100

	opWeightMsgUpdateUserDevices = "op_weight_msg_user_devices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUserDevices int = 100

	opWeightMsgDeleteUserDevices = "op_weight_msg_user_devices"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUserDevices int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cfetokenizationGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		UserDevicesList: []types.UserDevices{
			{
				Id:    0,
				Owner: sample.AccAddress(),
			},
			{
				Id:    1,
				Owner: sample.AccAddress(),
			},
		},
		UserDevicesCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfetokenizationGenesis)
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
	//
	//var weightMsgCreateUserDevices int
	//simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateUserDevices, &weightMsgCreateUserDevices, nil,
	//	func(_ *rand.Rand) {
	//		weightMsgCreateUserDevices = defaultWeightMsgCreateUserDevices
	//	},
	//)
	//operations = append(operations, simulation.NewWeightedOperation(
	//	weightMsgCreateUserDevices,
	//	cfetokenizationsimulation.SimulateMsgCreateUserDevices(am.accountKeeper, am.bankKeeper, am.keeper),
	//))
	//
	//var weightMsgUpdateUserDevices int
	//simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateUserDevices, &weightMsgUpdateUserDevices, nil,
	//	func(_ *rand.Rand) {
	//		weightMsgUpdateUserDevices = defaultWeightMsgUpdateUserDevices
	//	},
	//)
	//operations = append(operations, simulation.NewWeightedOperation(
	//	weightMsgUpdateUserDevices,
	//	cfetokenizationsimulation.SimulateMsgUpdateUserDevices(am.accountKeeper, am.bankKeeper, am.keeper),
	//))
	//
	//var weightMsgDeleteUserDevices int
	//simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteUserDevices, &weightMsgDeleteUserDevices, nil,
	//	func(_ *rand.Rand) {
	//		weightMsgDeleteUserDevices = defaultWeightMsgDeleteUserDevices
	//	},
	//)
	//operations = append(operations, simulation.NewWeightedOperation(
	//	weightMsgDeleteUserDevices,
	//	cfetokenizationsimulation.SimulateMsgDeleteUserDevices(am.accountKeeper, am.bankKeeper, am.keeper),
	//))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
