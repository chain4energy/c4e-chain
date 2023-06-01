package cfedistributor

import (
	"github.com/chain4energy/c4e-chain/testutil/utils"
	"math/rand"

	subdistributortestutils "github.com/chain4energy/c4e-chain/testutil/module/cfedistributor/subdistributor"
	"github.com/chain4energy/c4e-chain/testutil/sample"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
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
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	randDistinationType := RandomCollectorName(simState.Rand)
	subdistributors = append(subdistributors, subdistributortestutils.PrepareBurningDistributor(randDistinationType))
	if randDistinationType != subdistributortestutils.MainCollector {
		subdistributors = append(subdistributors, subdistributortestutils.PrepareInflationToPassAcoutSubDistr(randDistinationType))
	}
	subdistributors = append(subdistributors, subdistributortestutils.PrepareInflationSubDistributor(randDistinationType, true))

	genesisState.Params.SubDistributors = subdistributors
	cfedistributorGenesis := types.GenesisState{
		Params: types.NewParams(subdistributors),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cfedistributorGenesis)
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

func RandomCollectorName(r *rand.Rand) subdistributortestutils.DestinationType {
	randVal := utils.RandIntBetween(r, 0, 3)
	switch randVal {
	case 0:
		return subdistributortestutils.MainCollector
	case 1:
		return subdistributortestutils.ModuleAccount
	case 2:
		return subdistributortestutils.InternalAccount
	case 3:
		return subdistributortestutils.BaseAccount
	}
	return subdistributortestutils.MainCollector
}
