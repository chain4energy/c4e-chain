package cfevesting_test

import (

	"testing"

	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"

	"github.com/chain4energy/c4e-chain/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/chain4energy/c4e-chain/x/cfevesting/internal/testutils"

)
func TestGenesisWholeApp(t *testing.T) {

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		// this line is used by starport scaffolding # genesis/test/state
		VestingTypes: types.VestingTypes{},
	}

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	cfevesting.InitGenesis(ctx, app.CfevestingKeeper, genesisState, app.AccountKeeper)
	got := cfevesting.ExportGenesis(ctx, app.CfevestingKeeper)
	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

func TestGenesisVestingTypes(t *testing.T) {
	vestingTypesArray := testutils.GenerateVestingTypes(10, 1)
	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),
		VestingTypes: types.VestingTypes{VestingTypes: vestingTypesArray},
	}

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	cfevesting.InitGenesis(ctx, k, genesisState, ak)
	got := cfevesting.ExportGenesis(ctx, k)

	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}

func TestGenesisAccountVestingsList(t *testing.T) {
	accountVestingsListArray := testutils.GenerateAccountVestingsWithRandomVestings(10, 10, 1, 1)

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        types.VestingTypes{},
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	cfevesting.InitGenesis(ctx, k, genesisState, ak)
	got := cfevesting.ExportGenesis(ctx, k)
	require.NotNil(t, got)
	require.EqualValues(t, genesisState.Params, got.GetParams())
	require.EqualValues(t, genesisState.VestingTypes, (*got).VestingTypes)
	require.EqualValues(t, len(accountVestingsListArray), len((*got).AccountVestingsList.Vestings))

	testutils.EqualAccountVestings(t, accountVestingsListArray, (*got).AccountVestingsList.Vestings)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

}
