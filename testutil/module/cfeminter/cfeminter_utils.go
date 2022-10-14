package cfeminter

import (
	// "context"
	// "strconv"
	"time"

	// "github.com/chain4energy/c4e-chain/testutil/nullify"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	// "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"testing"

	routingdistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/chain4energy/c4e-chain/x/cfeminter"
	cfemintermodulekeeper "github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/stretchr/testify/require"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type C4eMinterUtils struct {
	t                     *testing.T
	helperCfeminterKeeper *cfemintermodulekeeper.Keeper
	helperAccountKeeper   *authkeeper.AccountKeeper
	helperBankKeeper      *bankkeeper.Keeper
	helperStakingKeeper   *stakingkeeper.Keeper
	bankUtils             *commontestutils.BankUtils
	authUtils             *commontestutils.AuthUtils
}

func NewC4eMinterUtils(t *testing.T, helperCfeminterKeeper *cfemintermodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	helperBankKeeper *bankkeeper.Keeper,
	helperStakingKeeper *stakingkeeper.Keeper, bankUtils *commontestutils.BankUtils,
	authUtils *commontestutils.AuthUtils) C4eMinterUtils {
	return C4eMinterUtils{t: t, helperCfeminterKeeper: helperCfeminterKeeper, helperAccountKeeper: helperAccountKeeper,
		helperBankKeeper: helperBankKeeper, helperStakingKeeper: helperStakingKeeper, bankUtils: bankUtils, authUtils: authUtils}
}

func (m *C4eMinterUtils) SetMinterState(ctx sdk.Context, position int32, amountMinted sdk.Int,
	remainderToMint sdk.Dec, lastMintBlockTime time.Time, remainderFromPreviousPeriod sdk.Dec) {

	minterState := cfemintertypes.MinterState{
		Position:                    position,
		AmountMinted:                amountMinted,
		RemainderToMint:             remainderToMint,
		LastMintBlockTime:           lastMintBlockTime,
		RemainderFromPreviousPeriod: remainderFromPreviousPeriod,
	}
	m.helperCfeminterKeeper.SetMinterState(ctx, minterState)
}

func (m *C4eMinterUtils) VerifyMinterState(ctx sdk.Context, expectedMinterStatePosition int32, expectedMinterStateAmountMinted sdk.Int,
	expectedMinterStateRemainderToMint sdk.Dec, expectedMinterStateLastMintBlockTime time.Time, expectedMinterStateRemainderFromPreviousPeriod sdk.Dec) {
	expectedMinterState := cfemintertypes.MinterState{
		Position:                    expectedMinterStatePosition,
		AmountMinted:                expectedMinterStateAmountMinted,
		RemainderToMint:             expectedMinterStateRemainderToMint,
		LastMintBlockTime:           expectedMinterStateLastMintBlockTime,
		RemainderFromPreviousPeriod: expectedMinterStateRemainderFromPreviousPeriod,
	}
	CompareMinterStates(m.t, expectedMinterState, m.helperCfeminterKeeper.GetMinterState(ctx))
}

func (m *C4eMinterUtils) VerifyMinterHistory(ctx sdk.Context, expectedMinterStateHistory ...cfemintertypes.MinterState) {
	history := m.helperCfeminterKeeper.GetAllMinterStateHistory(ctx)
	require.EqualValues(m.t, len(expectedMinterStateHistory), len(history))
	for i, ms := range expectedMinterStateHistory {
		histMS := history[i]
		CompareMinterStates(m.t, ms, histMS)
	}
}

func (m *C4eMinterUtils) Mint(ctx sdk.Context, expectedMintedAmount sdk.Int, expectedMinterStatePosition int32, expectedMinterStateAmountMinted sdk.Int,
	expectedMinterStateRemainderToMint sdk.Dec, expectedMinterStateLastMintBlockTime time.Time, expectedMinterStateRemainderFromPreviousPeriod sdk.Dec,
	expectedMintingReceiverAmount sdk.Int, expectedMinterStateHistory ...cfemintertypes.MinterState) {
	amount, err := m.helperCfeminterKeeper.Mint(ctx)
	require.NoError(m.t, err)
	require.Truef(m.t, expectedMintedAmount.Equal(amount), "expectedMintedAmount %s <> mintedAmount %s", expectedMintedAmount, amount)
	m.VerifyMinterState(ctx, expectedMinterStatePosition, expectedMinterStateAmountMinted, expectedMinterStateRemainderToMint, expectedMinterStateLastMintBlockTime,
		expectedMinterStateRemainderFromPreviousPeriod)
	m.bankUtils.VerifyModuleAccountDefultDenomBalance(ctx, routingdistributortypes.DistributorMainAccount, expectedMintingReceiverAmount)

	m.VerifyMinterHistory(ctx, expectedMinterStateHistory...)
}

func (m *C4eMinterUtils) MintError(ctx sdk.Context, errorMessage string) {
	_, err := m.helperCfeminterKeeper.Mint(ctx)
	require.EqualError(m.t, err, errorMessage)
}

func (m *C4eMinterUtils) InitGenesis(ctx sdk.Context, genState cfemintertypes.GenesisState) {
	cfeminter.InitGenesis(ctx, *m.helperCfeminterKeeper, m.helperAccountKeeper, genState)
}

func (m *C4eMinterUtils) ExportGenesis(ctx sdk.Context, expected cfemintertypes.GenesisState) {
	got := cfeminter.ExportGenesis(ctx, *m.helperCfeminterKeeper)
	require.NotNil(m.t, got)

	require.EqualValues(m.t, expected.Params.MintDenom, got.Params.MintDenom)
	CompareMinters(m.t, expected.Params.Minter, got.Params.Minter)
	CompareMinterStates(m.t, expected.MinterState, got.MinterState)
	require.EqualValues(m.t, len(expected.StateHistory), len(got.StateHistory))

	for i := 0; i < len(expected.StateHistory); i++ {
		CompareMinterStates(m.t, *expected.StateHistory[i], *got.StateHistory[i])
	}
}
func (m *C4eMinterUtils) VerifyInflation(ctx sdk.Context, expectedInflation sdk.Dec) {
	inflation, err := m.helperCfeminterKeeper.GetCurrentInflation(ctx)
	require.NoError(m.t, err)
	require.EqualValues(m.t, expectedInflation, inflation)
}

type ContextC4eMinterUtils struct {
	C4eMinterUtils
	testContext commontestutils.TestContext
}

func NewContextC4eMinterUtils(t *testing.T, testContext commontestutils.TestContext, helperCfeminterKeeper *cfemintermodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	helperBankKeeper *bankkeeper.Keeper,
	helperStakingKeeper *stakingkeeper.Keeper, bankUtils *commontestutils.BankUtils,
	authUtils *commontestutils.AuthUtils) *ContextC4eMinterUtils {
	c4eMinterUtils := NewC4eMinterUtils(t, helperCfeminterKeeper, helperAccountKeeper, helperBankKeeper, helperStakingKeeper, bankUtils, authUtils)
	return &ContextC4eMinterUtils{C4eMinterUtils: c4eMinterUtils, testContext: testContext}
}

func (m *ContextC4eMinterUtils) SetMinterState(position int32, amountMinted sdk.Int,
	remainderToMint sdk.Dec, lastMintBlockTime time.Time, remainderFromPreviousPeriod sdk.Dec) {
	m.C4eMinterUtils.SetMinterState(m.testContext.GetContext(), position, amountMinted, remainderToMint, lastMintBlockTime, remainderFromPreviousPeriod)
}

func (m *ContextC4eMinterUtils) Mint(expectedMintedAmount sdk.Int, expectedMinterStatePosition int32, expectedMinterStateAmountMinted sdk.Int,
	expectedMinterStateRemainderToMint sdk.Dec, expectedMinterStateLastMintBlockTime time.Time, expectedMinterStateRemainderFromPreviousPeriod sdk.Dec,
	expectedMintingReceiverAmount sdk.Int, expectedMinterStateHistory ...cfemintertypes.MinterState) {

	m.C4eMinterUtils.Mint(m.testContext.GetContext(), expectedMintedAmount, expectedMinterStatePosition, expectedMinterStateAmountMinted, expectedMinterStateRemainderToMint, expectedMinterStateLastMintBlockTime,
		expectedMinterStateRemainderFromPreviousPeriod, expectedMintingReceiverAmount, expectedMinterStateHistory...)

}

func (m *ContextC4eMinterUtils) MintError(errorMessage string) {
	m.C4eMinterUtils.MintError(m.testContext.GetContext(), errorMessage)
}

func (m *ContextC4eMinterUtils) InitGenesis(genState cfemintertypes.GenesisState) {
	m.C4eMinterUtils.InitGenesis(m.testContext.GetContext(), genState)
}

func (m *ContextC4eMinterUtils) ExportGenesis(expected cfemintertypes.GenesisState) {
	m.C4eMinterUtils.ExportGenesis(m.testContext.GetContext(), expected)
}

func (m *ContextC4eMinterUtils) VerifyMinterState(expectedMinterStatePosition int32, expectedMinterStateAmountMinted sdk.Int,
	expectedMinterStateRemainderToMint sdk.Dec, expectedMinterStateLastMintBlockTime time.Time, expectedMinterStateRemainderFromPreviousPeriod sdk.Dec) {
	m.C4eMinterUtils.VerifyMinterState(m.testContext.GetContext(), expectedMinterStatePosition, expectedMinterStateAmountMinted, expectedMinterStateRemainderToMint, expectedMinterStateLastMintBlockTime,
		expectedMinterStateRemainderFromPreviousPeriod)

}

func (m *ContextC4eMinterUtils) VerifyInflation(expectedInflation sdk.Dec) {
	m.C4eMinterUtils.VerifyInflation(m.testContext.GetContext(), expectedInflation)
}

func (m *ContextC4eMinterUtils) VerifyMinterHistory(expectedMinterStateHistory ...cfemintertypes.MinterState) {
	m.C4eMinterUtils.VerifyMinterHistory(m.testContext.GetContext(), expectedMinterStateHistory...)
}