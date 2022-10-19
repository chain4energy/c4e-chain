package cfevesting

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfevestingsimulation "github.com/chain4energy/c4e-chain/x/cfevesting/simulation"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
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
// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	cfevestingGenesis := types.GenesisState{
		Params: types.NewParams("stake"),
		VestingTypes: []types.GenesisVestingType{
			{
				Name:              "New vesting0",
				VestingPeriod:     helpers.RandomInt(simState.Rand, 10000000),
				VestingPeriodUnit: "second",
				LockupPeriod:      helpers.RandomInt(simState.Rand, 10000000),
				LockupPeriodUnit:  "second",
			},
			{
				Name:              "New vesting1",
				VestingPeriod:     helpers.RandomInt(simState.Rand, 1000),
				VestingPeriodUnit: "second",
				LockupPeriod:      helpers.RandomInt(simState.Rand, 1000),
				LockupPeriodUnit:  "second",
			},
			{
				Name:              "New vesting2",
				VestingPeriod:     helpers.RandomInt(simState.Rand, 10),
				VestingPeriodUnit: "second",
				LockupPeriod:      helpers.RandomInt(simState.Rand, 10),
				LockupPeriodUnit:  "second",
			},
			{
				Name:              "New vesting3",
				VestingPeriod:     helpers.RandomInt(simState.Rand, 5),
				VestingPeriodUnit: "second",
				LockupPeriod:      helpers.RandomInt(simState.Rand, 5),
				LockupPeriodUnit:  "second",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfevestingGenesis)
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
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
	// No decoder
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightSimulateVestingOperations = 100
	operations = append(operations, simulation.NewWeightedOperation(
		weightSimulateVestingOperations,
		cfevestingsimulation.SimulateVestingOperations(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateVestingAccount = 50
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVestingAccount,
		cfevestingsimulation.SimulateMsgCreateVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
