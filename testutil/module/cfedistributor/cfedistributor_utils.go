package cfedistributor

import (
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"testing"

	"github.com/chain4energy/c4e-chain/x/cfedistributor"
	cfedistributormodulekeeper "github.com/chain4energy/c4e-chain/x/cfedistributor/keeper"
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/stretchr/testify/require"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

type C4eDistributorKeeperUtils struct {
	t                          *testing.T
	helperCfedistributorKeeper *cfedistributormodulekeeper.Keeper
}

func NewC4eDistributorKeeperUtils(t *testing.T, helperCfedistributorKeeper *cfedistributormodulekeeper.Keeper) C4eDistributorKeeperUtils {
	return C4eDistributorKeeperUtils{t: t, helperCfedistributorKeeper: helperCfedistributorKeeper}
}

func (d *C4eDistributorKeeperUtils) SetSubDistributorsParams(ctx sdk.Context, subdistributors []cfedistributortypes.SubDistributor) {
	d.helperCfedistributorKeeper.SetParams(ctx, cfedistributortypes.NewParams(subdistributors))
}

func (d *C4eDistributorKeeperUtils) SetState(ctx sdk.Context, state cfedistributortypes.State) {
	d.helperCfedistributorKeeper.SetState(ctx, state)
}

func (h *C4eDistributorKeeperUtils) CheckNonNegativeCoinStateInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfedistributormodulekeeper.NonNegativeCoinStateInvariant(*h.helperCfedistributorKeeper)
	commontestutils.CheckInvariant(h.t, ctx, invariant, failed, message)
}

func (h *C4eDistributorKeeperUtils) GetC4eDistributorKeeper() *cfedistributormodulekeeper.Keeper {
	return h.helperCfedistributorKeeper
}

func (h *C4eDistributorKeeperUtils) CheckStateSumBalanceCheckInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfedistributormodulekeeper.StateSumBalanceCheckInvariant(*h.helperCfedistributorKeeper)
	commontestutils.CheckInvariant(h.t, ctx, invariant, failed, message)
}

type C4eDistributorUtils struct {
	C4eDistributorKeeperUtils
	helperAccountKeeper *authkeeper.AccountKeeper
}

func NewC4eDistributorUtils(t *testing.T, helperCfedistributorKeeper *cfedistributormodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper) C4eDistributorUtils {
	return C4eDistributorUtils{C4eDistributorKeeperUtils: NewC4eDistributorKeeperUtils(t, helperCfedistributorKeeper), helperAccountKeeper: helperAccountKeeper}
}

func (d *C4eDistributorUtils) VerifyStateAmount(ctx sdk.Context, stateName string, denom string, expectedRemains sdk.Dec) {
	state, _ := d.helperCfedistributorKeeper.GetState(ctx, stateName)

	coinRemains := state.Remains
	require.EqualValues(d.t, expectedRemains, coinRemains.AmountOf(denom))
}

func (d *C4eDistributorUtils) VerifyDefaultDenomStateAmount(ctx sdk.Context, account cfedistributortypes.Account, expectedRemains sdk.Dec) {
	d.VerifyStateAmount(ctx, account.GetAccounteKey(), commontestutils.DefaultTestDenom, expectedRemains)
}

func (d *C4eDistributorUtils) VerifyBurnStateAmount(ctx sdk.Context, denom string, expectedRemains sdk.Dec) {
	d.VerifyStateAmount(ctx, cfedistributortypes.BurnStateKey, denom, expectedRemains)

}

func (d *C4eDistributorUtils) VerifyDefaultDenomBurnStateAmount(ctx sdk.Context, expectedRemains sdk.Dec) {
	d.VerifyBurnStateAmount(ctx, commontestutils.DefaultTestDenom, expectedRemains)
}

func (d *C4eDistributorUtils) VerifyNumberOfStates(ctx sdk.Context, expectedNumberOfStates int) {
	require.EqualValues(d.t, expectedNumberOfStates, len(d.helperCfedistributorKeeper.GetAllStates(ctx)))
}

func (d *C4eDistributorUtils) InitGenesis(ctx sdk.Context, genState cfedistributortypes.GenesisState) {
	cfedistributor.InitGenesis(ctx, *d.helperCfedistributorKeeper, genState, d.helperAccountKeeper)
}

func (m *C4eDistributorUtils) ExportGenesis(ctx sdk.Context, expected cfedistributortypes.GenesisState) {
	got := cfedistributor.ExportGenesis(ctx, *m.helperCfedistributorKeeper)
	require.NotNil(m.t, got)

	require.ElementsMatch(m.t, expected.Params.SubDistributors, got.Params.SubDistributors)
	require.ElementsMatch(m.t, expected.States, got.States)
}

func (m *C4eDistributorUtils) ExportGenesisAndValidate(ctx sdk.Context) {
	exportedGenesis := cfedistributor.ExportGenesis(ctx, *m.helperCfedistributorKeeper)
	err := exportedGenesis.Validate()
	require.NoError(m.t, err)
}

func (m *C4eDistributorUtils) ValidateInvariants(ctx sdk.Context) {
	invariants := []sdk.Invariant{
		cfedistributormodulekeeper.StateSumBalanceCheckInvariant(*m.helperCfedistributorKeeper),
		cfedistributormodulekeeper.NonNegativeCoinStateInvariant(*m.helperCfedistributorKeeper),
	}
	commontestutils.ValidateManyInvariants(m.t, ctx, invariants)
}

func (m *C4eDistributorUtils) SetParams(ctx sdk.Context, params cfedistributortypes.Params) {
	m.helperCfedistributorKeeper.SetParams(ctx, params)
}

type ContextC4eDistributorUtils struct {
	C4eDistributorUtils
	testContext commontestutils.TestContext
}

func NewContextC4eDistributorUtils(t *testing.T, testContext commontestutils.TestContext, helperCfedistributorKeeper *cfedistributormodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper) *ContextC4eDistributorUtils {
	c4eDistributorUtils := NewC4eDistributorUtils(t, helperCfedistributorKeeper, helperAccountKeeper)
	return &ContextC4eDistributorUtils{C4eDistributorUtils: c4eDistributorUtils, testContext: testContext}
}

func (d *ContextC4eDistributorUtils) SetSubDistributorsParams(subdistributors []cfedistributortypes.SubDistributor) {
	d.C4eDistributorUtils.SetSubDistributorsParams(d.testContext.GetContext(), subdistributors)
}

func (d *ContextC4eDistributorUtils) VerifyStateAmount(stateName string, denom string, expectedRemains sdk.Dec) {
	d.C4eDistributorUtils.VerifyStateAmount(d.testContext.GetContext(), stateName, denom, expectedRemains)

}

func (d *ContextC4eDistributorUtils) VerifyDefaultDenomStateAmount(account cfedistributortypes.Account, expectedRemains sdk.Dec) {
	d.C4eDistributorUtils.VerifyDefaultDenomStateAmount(d.testContext.GetContext(), account, expectedRemains)
}

func (d *ContextC4eDistributorUtils) VerifyBurnStateAmount(denom string, expectedRemains sdk.Dec) {
	d.C4eDistributorUtils.VerifyBurnStateAmount(d.testContext.GetContext(), denom, expectedRemains)

}

func (d *ContextC4eDistributorUtils) VerifyDefaultDenomBurnStateAmount(expectedRemains sdk.Dec) {
	d.C4eDistributorUtils.VerifyDefaultDenomBurnStateAmount(d.testContext.GetContext(), expectedRemains)
}

func (d *ContextC4eDistributorUtils) VerifyNumberOfStates(expectedNumberOfStates int) {
	d.C4eDistributorUtils.VerifyNumberOfStates(d.testContext.GetContext(), expectedNumberOfStates)
}

func (m *ContextC4eDistributorUtils) InitGenesis(genState cfedistributortypes.GenesisState) {
	m.C4eDistributorUtils.InitGenesis(m.testContext.GetContext(), genState)
}

func (m *ContextC4eDistributorUtils) ExportGenesis(expected cfedistributortypes.GenesisState) {
	m.C4eDistributorUtils.ExportGenesis(m.testContext.GetContext(), expected)
}

func (m *ContextC4eDistributorUtils) ValidateGenesisAndInvariants() {
	m.C4eDistributorUtils.ExportGenesisAndValidate(m.testContext.GetContext())
	m.C4eDistributorUtils.ValidateInvariants(m.testContext.GetContext())
}

func (m *ContextC4eDistributorUtils) SetState(state cfedistributortypes.State) {
	m.C4eDistributorUtils.SetState(m.testContext.GetContext(), state)
}

func (m *ContextC4eDistributorUtils) SetParams(params cfedistributortypes.Params) {
	m.C4eDistributorUtils.SetParams(m.testContext.GetContext(), params)
}

func (m *ContextC4eDistributorUtils) CheckStateSumBalanceCheckInvariant(failed bool, message string) {
	m.C4eDistributorUtils.CheckStateSumBalanceCheckInvariant(m.testContext.GetContext(), failed, message)

}

func (h *ContextC4eDistributorUtils) InitGenesisError(genState cfedistributortypes.GenesisState, errorMessage string) {
	h.C4eDistributorUtils.InitGenesisError(h.testContext.GetContext(), genState, errorMessage)
}

func (h *C4eDistributorUtils) InitGenesisError(ctx sdk.Context, genState cfedistributortypes.GenesisState, errorMessage string) {
	require.PanicsWithValue(h.t, errorMessage,
		func() {
			cfedistributor.InitGenesis(ctx, *h.helperCfedistributorKeeper, genState, h.helperAccountKeeper)
		}, "")
}
