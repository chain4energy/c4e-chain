package cfeminter

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"
	"time"
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
const SecondsInYear = int32(3600 * 24 * 365)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	randAmount := helpers.RandomInt(simState.Rand, 40000000000000)
	randStepDuration := helpers.RandomInt(simState.Rand, int(31536000*time.Second*4))
	randomIntBetween := helpers.RandIntBetween(simState.Rand, 1, 100)
	amountMultiplierFloat := float64(randomIntBetween) / float64(100)
	randAmountMultiplierFactor := fmt.Sprintf("%f", amountMultiplierFloat)
	now := simState.GenTimestamp

	exponentialStepMinting := types.ExponentialStepMinting{
		Amount:           sdk.NewInt(randAmount),
		StepDuration:     time.Duration(randStepDuration),
		AmountMultiplier: sdk.MustNewDecFromStr(randAmountMultiplierFactor),
	}
	config, _ := codectypes.NewAnyWithValue(&exponentialStepMinting)
	minters := []*types.Minter{
		{
			SequenceId: 1,
			Config:     config,
		},
	}

	genesisState := types.GenesisState{
		Params: types.NewParams("stake", now, minters),
		MinterState: types.MinterState{
			SequenceId:   1,
			AmountMinted: sdk.NewInt(0),
		},
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&genesisState)
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
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
