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

type C4eDistributorUtils struct {
	t                     *testing.T
	helperCfedistributorKeeper *cfedistributormodulekeeper.Keeper
	helperAccountKeeper   *authkeeper.AccountKeeper
	bankUtils             *commontestutils.BankUtils
}

func NewC4eDistributorUtils(t *testing.T, helperCfedistributorKeeper *cfedistributormodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *commontestutils.BankUtils) C4eDistributorUtils {
	return C4eDistributorUtils{t: t, helperCfedistributorKeeper: helperCfedistributorKeeper, helperAccountKeeper: helperAccountKeeper,
		bankUtils: bankUtils}
}

func (d *C4eDistributorUtils) SetSubDistributorsParams(ctx sdk.Context, subdistributors []cfedistributortypes.SubDistributor) {
	d.helperCfedistributorKeeper.SetParams(ctx, cfedistributortypes.NewParams(subdistributors))
}

func (d *C4eDistributorUtils) SetState(ctx sdk.Context, state cfedistributortypes.State) {
	d.helperCfedistributorKeeper.SetState(ctx, state)
}

func (d *C4eDistributorUtils) VerifyStateAmount(ctx sdk.Context, stateName string, denom string, expectedRemains sdk.Dec) {
	state, _ := d.helperCfedistributorKeeper.GetState(ctx, stateName)

	coinRemains := state.CoinsStates
	require.EqualValues(d.t, expectedRemains, coinRemains.AmountOf(denom))
}

func (d *C4eDistributorUtils) VerifyDefaultDenomStateAmount(ctx sdk.Context, stateName string, expectedRemains sdk.Dec) {
	d.VerifyStateAmount(ctx, stateName, commontestutils.DefaultTestDenom, expectedRemains)
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

func (m *C4eDistributorUtils) SetParams(ctx sdk.Context, params cfedistributortypes.Params) {
	m.helperCfedistributorKeeper.SetParams(ctx, params)
}

type ContextC4eDistributorUtils struct {
	C4eDistributorUtils
	testContext commontestutils.TestContext
}

func NewContextC4eDistributorUtils(t *testing.T, testContext commontestutils.TestContext, helperCfedistributorKeeper *cfedistributormodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *commontestutils.BankUtils) *ContextC4eDistributorUtils {
	c4eDistributorUtils := NewC4eDistributorUtils(t, helperCfedistributorKeeper, helperAccountKeeper, bankUtils)
	return &ContextC4eDistributorUtils{C4eDistributorUtils: c4eDistributorUtils, testContext: testContext}
}

func (d *ContextC4eDistributorUtils) SetSubDistributorsParams(subdistributors []cfedistributortypes.SubDistributor) {
	d.C4eDistributorUtils.SetSubDistributorsParams(d.testContext.GetContext(), subdistributors)
}

func (d *ContextC4eDistributorUtils) VerifyStateAmount(stateName string, denom string, expectedRemains sdk.Dec) {
	d.C4eDistributorUtils.VerifyStateAmount(d.testContext.GetContext(), stateName, denom, expectedRemains)

}

func (d *ContextC4eDistributorUtils) VerifyDefaultDenomStateAmount(stateName string, expectedRemains sdk.Dec) {
	d.C4eDistributorUtils.VerifyDefaultDenomStateAmount(d.testContext.GetContext(), stateName, expectedRemains)
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

func (m *ContextC4eDistributorUtils) SetState( state cfedistributortypes.State) {
	m.C4eDistributorUtils.SetState(m.testContext.GetContext(), state)
}

func (m *ContextC4eDistributorUtils) SetParams(params cfedistributortypes.Params) {
	m.C4eDistributorUtils.SetParams(m.testContext.GetContext(), params)
}
