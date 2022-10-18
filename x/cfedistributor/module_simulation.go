package cfedistributor

import (
	sdkrand "github.com/chain4energy/c4e-chain/testutil/simulation"
	"math/rand"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfedistributorsimulation "github.com/chain4energy/c4e-chain/x/cfedistributor/simulation"
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
	_ = cfedistributorsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(MainCollector))
	subdistributors = append(subdistributors, prepareInflationSubDistributor(MainCollector, true))
	subdistributors = append(subdistributors, prepareInflationToPassAcoutSubDistr(MainCollector))
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

type DestinationType int64

const (
	MainCollector DestinationType = iota
	ModuleAccount
	InternalAccount
	BaseAccount
)

func prepareInflationToPassAcoutSubDistr(passThroughAccoutType DestinationType) types.SubDistributor {
	source := types.Account{
		Id:   "c4e",
		Type: "MAIN",
	}

	var address string
	if passThroughAccoutType == BaseAccount {
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = "c4e_distributor"
	}

	var destAccount = types.Account{
		Id: address,
	}

	if passThroughAccoutType == ModuleAccount {
		destAccount.Type = "MODULE_ACCOUNT"
	} else if passThroughAccoutType == InternalAccount {
		destAccount.Type = "INTERNAL_ACCOUNT"
	} else {
		destAccount.Type = "BASE_ACCOUNT"
	}

	if passThroughAccoutType == MainCollector {
		destAccount.Type = "MAIN"
	}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     nil,
		BurnShare: &burnShare,
	}
	return types.SubDistributor{
		Name:        "pass_distributor",
		Sources:     []*types.Account{&source},
		Destination: destination,
	}
}

func prepareInflationSubDistributor(sourceAccoutType DestinationType, toValidators bool) types.SubDistributor {
	var address string
	if sourceAccoutType == BaseAccount {
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = "c4e_distributor"
	}

	var source = types.Account{
		Id: address,
	}

	if sourceAccoutType == ModuleAccount {
		source.Type = "MODULE_ACCOUNT"
	} else if sourceAccoutType == InternalAccount {
		source.Type = "INTERNAL_ACCOUNT"
	} else {
		source.Type = "BASE_ACCOUNT"
	}

	if sourceAccoutType == MainCollector {
		source.Type = "MAIN"
	}

	// source := types.Account{IsMainCollector: true, IsModuleAccount: false, IsInternalAccount: false}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

	var destName string
	if toValidators {
		destName = types.ValidatorsRewardsCollector
	} else {
		destName = "no_validators"
	}

	var destAccount = types.Account{
		Id:   destName,
		Type: "MODULE_ACCOUNT",
	}

	var shareDevelopmentFundAccount = types.Account{
		Id:   "cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag",
		Type: "BASE_ACCOUNT",
	}

	shareDevelopmentFund := types.Share{
		Name:    "development_fund",
		Percent: sdk.MustNewDecFromStr("10.345"),
		Account: shareDevelopmentFundAccount,
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     []*types.Share{&shareDevelopmentFund},
		BurnShare: &burnShare,
	}

	return types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []*types.Account{&source},
		Destination: destination,
	}
}

func randomCollectorName(r *rand.Rand) DestinationType {
	randVal := sdkrand.RandIntBetween(r, 0, 3)
	switch randVal {
	case 0:
		return MainCollector
	case 1:
		return MainCollector
	case 2:
		return MainCollector
	case 3:
		return MainCollector
	}
	return MainCollector
}

func prepareBurningDistributor(destinationType DestinationType) types.SubDistributor {
	var address string
	if destinationType == BaseAccount {
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = "c4e_distributor"
	}

	var destAccount = types.Account{}
	destAccount.Id = address

	if destinationType == ModuleAccount {
		destAccount.Type = "MODULE_ACCOUNT"
	} else if destinationType == InternalAccount {
		destAccount.Type = "INTERNAL_ACCOUNT"
	} else {
		destAccount.Type = "BASE_ACCOUNT"
	}

	if destinationType == MainCollector {
		destAccount.Type = "MAIN"
	}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("51"),
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     nil,
		BurnShare: &burnShare,
	}

	distributor1 := types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []*types.Account{{Id: "fee_collector", Type: "MODULE_ACCOUNT"}},
		Destination: destination,
	}

	return distributor1
}
