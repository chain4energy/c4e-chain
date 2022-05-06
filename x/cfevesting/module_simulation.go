package cfevesting

import (
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
	_ = cfevestingsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateVestingPool = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVestingPool int = 100

	opWeightMsgWithdrawAllAvailable = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawAllAvailable int = 100

	opWeightMsgDelegate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDelegate int = 100

	opWeightMsgUndelegate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUndelegate int = 100

	opWeightMsgBeginRedelegate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBeginRedelegate int = 100

	opWeightMsgWithdrawDelegatorReward = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawDelegatorReward int = 100

	opWeightMsgSendVesting = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendVesting int = 100

	opWeightMsgVote = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVote int = 100

	opWeightMsgVoteWeighted = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVoteWeighted int = 100

	opWeightMsgCreateVestingAccount = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVestingAccount int = 100

	opWeightMsgSendToVestingAccount = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendToVestingAccount int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cfevestingGenesis := types.GenesisState{
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
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateVestingPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVestingPool, &weightMsgCreateVestingPool, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVestingPool = defaultWeightMsgCreateVestingPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVestingPool,
		cfevestingsimulation.SimulateMsgCreateVestingPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawAllAvailable int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawAllAvailable, &weightMsgWithdrawAllAvailable, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawAllAvailable = defaultWeightMsgWithdrawAllAvailable
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawAllAvailable,
		cfevestingsimulation.SimulateMsgWithdrawAllAvailable(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// var weightMsgDelegate int
	// simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDelegate, &weightMsgDelegate, nil,
	// 	func(_ *rand.Rand) {
	// 		weightMsgDelegate = defaultWeightMsgDelegate
	// 	},
	// )
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgDelegate,
	// 	cfevestingsimulation.SimulateMsgDelegate(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	// var weightMsgUndelegate int
	// simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUndelegate, &weightMsgUndelegate, nil,
	// 	func(_ *rand.Rand) {
	// 		weightMsgUndelegate = defaultWeightMsgUndelegate
	// 	},
	// )
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgUndelegate,
	// 	cfevestingsimulation.SimulateMsgUndelegate(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	// var weightMsgBeginRedelegate int
	// simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBeginRedelegate, &weightMsgBeginRedelegate, nil,
	// 	func(_ *rand.Rand) {
	// 		weightMsgBeginRedelegate = defaultWeightMsgBeginRedelegate
	// 	},
	// )
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgBeginRedelegate,
	// 	cfevestingsimulation.SimulateMsgBeginRedelegate(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	// var weightMsgWithdrawDelegatorReward int
	// simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawDelegatorReward, &weightMsgWithdrawDelegatorReward, nil,
	// 	func(_ *rand.Rand) {
	// 		weightMsgWithdrawDelegatorReward = defaultWeightMsgWithdrawDelegatorReward
	// 	},
	// )
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgWithdrawDelegatorReward,
	// 	cfevestingsimulation.SimulateMsgWithdrawDelegatorReward(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	// var weightMsgSendVesting int
	// simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendVesting, &weightMsgSendVesting, nil,
	// 	func(_ *rand.Rand) {
	// 		weightMsgSendVesting = defaultWeightMsgSendVesting
	// 	},
	// )
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgSendVesting,
	// 	cfevestingsimulation.SimulateMsgSendVesting(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	// var weightMsgVote int
	// simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVote, &weightMsgVote, nil,
	// 	func(_ *rand.Rand) {
	// 		weightMsgVote = defaultWeightMsgVote
	// 	},
	// )
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgVote,
	// 	cfevestingsimulation.SimulateMsgVote(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	// var weightMsgVoteWeighted int
	// simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVoteWeighted, &weightMsgVoteWeighted, nil,
	// 	func(_ *rand.Rand) {
	// 		weightMsgVoteWeighted = defaultWeightMsgVoteWeighted
	// 	},
	// )
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgVoteWeighted,
	// 	cfevestingsimulation.SimulateMsgVoteWeighted(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	var weightMsgCreateVestingAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVestingAccount, &weightMsgCreateVestingAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVestingAccount = defaultWeightMsgCreateVestingAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVestingAccount,
		cfevestingsimulation.SimulateMsgCreateVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendToVestingAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendToVestingAccount, &weightMsgSendToVestingAccount, nil,
		func(_ *rand.Rand) {
			weightMsgSendToVestingAccount = defaultWeightMsgSendToVestingAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendToVestingAccount,
		cfevestingsimulation.SimulateMsgSendToVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
