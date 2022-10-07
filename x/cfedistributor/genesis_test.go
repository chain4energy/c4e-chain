package cfedistributor_test

import (
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	account := types.Account{
		Id:   "usage_incentives_collector",
		Type: "INTERNAL_ACCOUNT",
	}

	state := types.State{
		Account:     &account,
		Burn:        false,
		CoinsStates: nil,
	}

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(MainCollector))
	genesisState.Params.SubDistributors = subdistributors
	k, ctx := keepertest.CfedistributorKeeper(t)
	app := testapp.Setup(false)
	cfedistributor.InitGenesis(ctx, *k, genesisState, app.AccountKeeper)

	k.SetState(ctx, state)
	got := cfedistributor.ExportGenesis(ctx, *k)
	require.EqualValues(t, sdk.MustNewDecFromStr("51"), got.Params.SubDistributors[0].Destination.BurnShare.Percent)
	require.EqualValues(t, "MAIN", got.Params.SubDistributors[0].Destination.Account.Type)
	require.EqualValues(t, "usage_incentives_collector", got.GetStates()[0].Account.Id)
	require.EqualValues(t, "INTERNAL_ACCOUNT", got.GetStates()[0].Account.Type)
}

func TestGenesisImport(t *testing.T) {
	account := types.Account{
		Id:   "usage_incentives_collector",
		Type: "INTERNAL_ACCOUNT",
	}

	state := types.State{
		Account:     &account,
		Burn:        false,
		CoinsStates: nil,
	}
	var states []*types.State
	states = append(states, &state)
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		States: states,
	}

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(MainCollector))
	genesisState.Params.SubDistributors = subdistributors
	k, ctx := keepertest.CfedistributorKeeper(t)
	app := testapp.Setup(false)
	cfedistributor.InitGenesis(ctx, *k, genesisState, app.AccountKeeper)
	require.EqualValues(t, "usage_incentives_collector", k.GetAllStates(ctx)[0].Account.Id)
	require.EqualValues(t, "INTERNAL_ACCOUNT", k.GetAllStates(ctx)[0].Account.Type)
}
