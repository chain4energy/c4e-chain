package cfevesting

import (
	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	cfevestingpoolsimulation "github.com/chain4energy/c4e-chain/x/cfevesting/simulation"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"
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
	cfevestingGenesis := types.GenesisState{
		Params: types.NewParams("stake"),
		VestingTypes: []types.GenesisVestingType{
			{
				Name:              "New vesting0",
				VestingPeriod:     utils.RandInt64(simState.Rand, 10000),
				VestingPeriodUnit: "second",
				LockupPeriod:      utils.RandInt64(simState.Rand, 10000),
				LockupPeriodUnit:  "second",
				Free:              sdk.MustNewDecFromStr("0.6"),
			},
			{
				Name:              "New vesting1",
				VestingPeriod:     utils.RandInt64(simState.Rand, 1000),
				VestingPeriodUnit: "second",
				LockupPeriod:      utils.RandInt64(simState.Rand, 1000),
				LockupPeriodUnit:  "second",
				Free:              sdk.MustNewDecFromStr("0.5"),
			},
			{
				Name:              "New vesting2",
				VestingPeriod:     utils.RandInt64(simState.Rand, 100),
				VestingPeriodUnit: "second",
				LockupPeriod:      utils.RandInt64(simState.Rand, 100),
				LockupPeriodUnit:  "second",
				Free:              sdk.MustNewDecFromStr("0.4"),
			},
			{
				Name:              "New vesting3",
				VestingPeriod:     utils.RandInt64(simState.Rand, 10),
				VestingPeriodUnit: "second",
				LockupPeriod:      utils.RandInt64(simState.Rand, 10),
				LockupPeriodUnit:  "second",
				Free:              sdk.MustNewDecFromStr("0.3"),
			},
		},
		AccountVestingPools: []*types.AccountVestingPools{},
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
	var weightSimulateSendToVestingAccount = 100
	operations = append(operations, simulation.NewWeightedOperation(
		weightSimulateSendToVestingAccount,
		cfevestingpoolsimulation.SimulateSendToVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	var weightSimulateCreateVestingPool = 30
	operations = append(operations, simulation.NewWeightedOperation(
		weightSimulateCreateVestingPool,
		cfevestingpoolsimulation.SimulateCreateVestingPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	var weightSimulateVestingOperations = 30
	operations = append(operations, simulation.NewWeightedOperation(
		weightSimulateVestingOperations,
		cfevestingpoolsimulation.SimulateVestingOperations(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateVestingAccount = 10
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVestingAccount,
		cfevestingpoolsimulation.SimulateMsgCreateVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	var weightSimulateVestingMultiOperations = 100
	operations = append(operations, simulation.NewWeightedOperation(
		weightSimulateVestingMultiOperations,
		cfevestingpoolsimulation.SimulateVestingMultiOperations(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	var weightSimulateWithdrawAllAvailable = 50
	operations = append(operations, simulation.NewWeightedOperation(
		weightSimulateWithdrawAllAvailable,
		cfevestingpoolsimulation.SimulateWithdrawAllAvailable(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSplitVesting = 50
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSplitVesting,
		cfevestingpoolsimulation.SimulateMsgSplitVesting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMoveAvailableVesting = 50
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMoveAvailableVesting,
		cfevestingpoolsimulation.SimulateMsgMoveAvailableVesting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMoveAvailableVestingByDenoms = 50
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMoveAvailableVestingByDenoms,
		cfevestingpoolsimulation.SimulateMsgMoveAvailableVestingByDenoms(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
