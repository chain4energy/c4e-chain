package cfeminter

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/testutil/sample"
	cfemintersimulation "github.com/chain4energy/c4e-chain/x/cfeminter/simulation"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
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
	_ = cfemintersimulation.FindAccount
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
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}

	startAmountYearly := int64(40000000000000)
	//AddHelperModuleAccountPerms()
	now := time.Now()
	pminter := types.PeriodicReductionMinter{MintAmount: sdk.NewInt(startAmountYearly), MintPeriod: SecondsInYear, ReductionPeriodLength: 4, ReductionFactor: sdk.MustNewDecFromStr("0.5")}

	minter := types.Minter{
		Start: now,
		Periods: []*types.MintingPeriod{
			{Position: 1, Type: types.PERIODIC_REDUCTION_MINTER,
				PeriodicReductionMinter: &pminter},
		}}

	genesisState := types.GenesisState{
		Params:      types.NewParams("uc4e", minter),
		MinterState: types.MinterState{Position: 1, AmountMinted: sdk.NewInt(0)},
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
func CreateIncrementalAccounts(accNum int, genInitNumber int) []sdk.AccAddress {
	var addresses []sdk.AccAddress
	var buffer bytes.Buffer

	// start at 100 so we can make up to 999 test addresses with valid test addresses
	for i := 100; i < (accNum + 100); i++ {
		numString := strconv.Itoa(i + genInitNumber)
		buffer.WriteString("A58856F0FD53BF058B4909A21AEC019107BA6") // base address string

		buffer.WriteString(numString) // adding on final two digits to make addresses unique
		res, _ := sdk.AccAddressFromHex(buffer.String())
		bech := res.String()
		addr, _ := TestAddr(buffer.String(), bech)

		addresses = append(addresses, addr)
		buffer.Reset()
	}

	return addresses
}
func TestAddr(addr string, bech string) (sdk.AccAddress, error) {
	res, err := sdk.AccAddressFromHex(addr)
	if err != nil {
		return nil, err
	}
	bechexpected := res.String()
	if bech != bechexpected {
		return nil, fmt.Errorf("bech encoding doesn't match reference")
	}

	bechres, err := sdk.AccAddressFromBech32(bech)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(bechres, res) {
		return nil, err
	}

	return res, nil
}

const helperModuleAccount = "helperTestAcc"

//func AddHelperModuleAccountPerms() {
//	perms := []string{"minter"}
//	app.AddMaccPerms(helperModuleAccount, perms)
//}
