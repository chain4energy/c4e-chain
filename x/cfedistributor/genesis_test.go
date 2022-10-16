package cfedistributor_test

import (
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"testing"
)

func TestGenesis(t *testing.T) {
	account := types.Account{
		Id:   "usage_incentives_collector",
		Type: "INTERNAL_ACCOUNT",
	}

	state := types.State{
		Account:     &account,
		Burn:        false,
		CoinsStates: nil,
	}

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(MainCollector))
	genesisState.Params.SubDistributors = subdistributors

	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eDistributorUtils.InitGenesis(genesisState)
	testHelper.C4eDistributorUtils.SetState(state)
	genesisState.States = []*types.State{&state}
	testHelper.C4eDistributorUtils.ExportGenesis(genesisState)

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
	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eDistributorUtils.InitGenesis(genesisState)
	testHelper.C4eDistributorUtils.ExportGenesis(genesisState)
}

func TestGenesisNoStates(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(MainCollector))
	genesisState.Params.SubDistributors = subdistributors

	testHelper := testapp.SetupTestApp(t)
	testHelper.C4eDistributorUtils.InitGenesis(genesisState)
	testHelper.C4eDistributorUtils.ExportGenesis(genesisState)

}
