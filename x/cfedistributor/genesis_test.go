package cfedistributor_test

import (
	"github.com/chain4energy/c4e-chain/v2/testutil/app"
	"testing"

	subdistributortestutils "github.com/chain4energy/c4e-chain/v2/testutil/module/cfedistributor/subdistributor"
	"github.com/chain4energy/c4e-chain/v2/x/cfedistributor/types"
)

func TestGenesis(t *testing.T) {
	account := types.Account{
		Id:   "usage_incentives_collector",
		Type: "InternalAccount",
	}

	state := types.State{
		Account: &account,
		Burn:    false,
		Remains: nil,
	}

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	burningSubSistributor := subdistributortestutils.PrepareBurningDistributor(subdistributortestutils.MainCollector)
	subdistributors = append(subdistributors, burningSubSistributor)
	subdistributors = append(subdistributors, subdistributortestutils.PreparareHelperDistributorForDestination(burningSubSistributor.Destinations.PrimaryShare))

	genesisState.Params.SubDistributors = subdistributors

	testHelper := app.SetupTestApp(t)
	testHelper.C4eDistributorUtils.InitGenesis(genesisState)
	testHelper.C4eDistributorUtils.SetState(state)
	genesisState.States = []*types.State{&state}
	testHelper.C4eDistributorUtils.ExportGenesis(genesisState)
	testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
}

func TestGenesisImport(t *testing.T) {
	account := types.Account{
		Id:   "usage_incentives_collector",
		Type: "InternalAccount",
	}

	state := types.State{
		Account: &account,
		Burn:    false,
		Remains: nil,
	}
	var states []*types.State
	states = append(states, &state)
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		States: states,
	}

	var subdistributors []types.SubDistributor
	burningSubSistributor := subdistributortestutils.PrepareBurningDistributor(subdistributortestutils.MainCollector)
	subdistributors = append(subdistributors, burningSubSistributor)
	subdistributors = append(subdistributors, subdistributortestutils.PreparareHelperDistributorForDestination(burningSubSistributor.Destinations.PrimaryShare))
	genesisState.Params.SubDistributors = subdistributors
	testHelper := app.SetupTestApp(t)
	testHelper.C4eDistributorUtils.InitGenesis(genesisState)
	testHelper.C4eDistributorUtils.ExportGenesis(genesisState)
	testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
}

func TestGenesisNoStates(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	burningSubSistributor := subdistributortestutils.PrepareBurningDistributor(subdistributortestutils.MainCollector)
	subdistributors = append(subdistributors, burningSubSistributor)
	subdistributors = append(subdistributors, subdistributortestutils.PreparareHelperDistributorForDestination(burningSubSistributor.Destinations.PrimaryShare))
	genesisState.Params.SubDistributors = subdistributors

	testHelper := app.SetupTestApp(t)
	testHelper.C4eDistributorUtils.InitGenesis(genesisState)
	testHelper.C4eDistributorUtils.ExportGenesis(genesisState)
	testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
}

func TestGenesisBurnStateAccNotNil(t *testing.T) {
	account := types.Account{
		Id:   "usage_incentives_collector",
		Type: "InternalAccount",
	}

	state := types.State{
		Account: &account,
		Burn:    true,
		Remains: nil,
	}

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	burningSubSistributor := subdistributortestutils.PrepareBurningDistributor(subdistributortestutils.MainCollector)
	subdistributors = append(subdistributors, burningSubSistributor)
	subdistributors = append(subdistributors, subdistributortestutils.PreparareHelperDistributorForDestination(burningSubSistributor.Destinations.PrimaryShare))

	genesisState.Params.SubDistributors = subdistributors

	testHelper := app.SetupTestApp(t)
	testHelper.C4eDistributorUtils.InitGenesis(genesisState)
	testHelper.C4eDistributorUtils.SetState(state)
	state.Account = nil
	genesisState.States = []*types.State{&state}
	testHelper.C4eDistributorUtils.ExportGenesis(genesisState)
	testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
}

func TestGenesisBurnState(t *testing.T) {

	state := types.State{
		Account: nil,
		Burn:    true,
		Remains: nil,
	}

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	var subdistributors []types.SubDistributor
	burningSubSistributor := subdistributortestutils.PrepareBurningDistributor(subdistributortestutils.MainCollector)
	subdistributors = append(subdistributors, burningSubSistributor)
	subdistributors = append(subdistributors, subdistributortestutils.PreparareHelperDistributorForDestination(burningSubSistributor.Destinations.PrimaryShare))

	genesisState.Params.SubDistributors = subdistributors

	testHelper := app.SetupTestApp(t)
	testHelper.C4eDistributorUtils.InitGenesis(genesisState)
	testHelper.C4eDistributorUtils.SetState(state)
	genesisState.States = []*types.State{&state}
	testHelper.C4eDistributorUtils.ExportGenesis(genesisState)
	testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
}
