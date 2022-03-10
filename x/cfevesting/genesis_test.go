package cfevesting_test

import (
	"fmt"
	"strconv"
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/stretchr/testify/require"
	tmdb "github.com/tendermint/tm-db"

	"github.com/chain4energy/c4e-chain/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// TODO verify how to create more than one keeper for testing
func TestGenesis(t *testing.T) {
	t.Skip("Skipping test - test to fix or remove later")
	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		// this line is used by starport scaffolding # genesis/test/state
		VestingTypes: types.VestingTypes{},
	}
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	k, ctx := keepertest.CfevestingKeeperWithBlockHeightAndStore(t, 0, db, stateStore)
	ak, _ := keepertest.AccountKeeperWithBlockHeight(t, ctx, stateStore, db)

	cfevesting.InitGenesis(ctx, *k, genesisState, ak)
	got := cfevesting.ExportGenesis(ctx, *k)
	require.NotNil(t, got)
	require.EqualValues(t, genesisState, *got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

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

	vestingType1 := types.VestingType{"test1", 2324, 42423, 4243, true}
	vestingType2 := types.VestingType{"test2", 1111, 112233, 445566, false}
	vestingTypesArray := []*types.VestingType{&vestingType1, &vestingType2}
	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		// this line is used by starport scaffolding # genesis/test/state
		VestingTypes: types.VestingTypes{vestingTypesArray},
	}

	// k, ctx := keepertest.CfevestingKeeper(t)

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

	// this line is used by starport scaffolding # genesis/test/assert
}

func TestGenesisAccountVestingsList(t *testing.T) {

	accountVestings1 := types.AccountVestings{}
	accountVestings1.Address = "someAddr1"
	vesting11 := types.Vesting{"test1", 2324, 42423, 4243, 14243, 24243, 34243, 44243, 54243, true, 0}
	vesting12 := types.Vesting{"test2", 92324, 942423, 94243, 914243, 924243, 934243, 944243, 954243, false, 0}
	vestingsArray1 := []*types.Vesting{&vesting11, &vesting12}
	accountVestings1.Vestings = vestingsArray1

	accountVestings2 := types.AccountVestings{}
	accountVestings2.Address = "someAddr2"
	vesting21 := types.Vesting{"test3", 2324, 42423, 4243, 14243, 24243, 34243, 44243, 54243, true, 0}
	vesting22 := types.Vesting{"test4", 92324, 942423, 94243, 914243, 924243, 934243, 944243, 954243, false, 0}
	vestingsArray2 := []*types.Vesting{&vesting21, &vesting22}
	accountVestings2.Vestings = vestingsArray2

	accountVestingsListArray := []*types.AccountVestings{&accountVestings1, &accountVestings2}

	genesisState := types.GenesisState{
		Params: types.NewParams("test_denom"),

		VestingTypes:        types.VestingTypes{},
		AccountVestingsList: types.AccountVestingsList{Vestings: accountVestingsListArray},
	}

	// k, ctx := keepertest.CfevestingKeeper(t)

	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	k := app.CfevestingKeeper
	ak := app.AccountKeeper

	cfevesting.InitGenesis(ctx, k, genesisState, ak)
	got := cfevesting.ExportGenesis(ctx, k)
	require.NotNil(t, got)
	require.EqualValues(t, genesisState.Params, got.GetParams())
	require.EqualValues(t, genesisState.VestingTypes, (*got).VestingTypes)
	require.EqualValues(t, len(genesisState.AccountVestingsList.Vestings), len((*got).AccountVestingsList.Vestings))

	for _, accVest := range genesisState.AccountVestingsList.Vestings {
		found := false
		for i, accVestExp := range (*got).AccountVestingsList.Vestings {
			fmt.Println("sdasa: " + strconv.Itoa(i) + " - " + accVestExp.Address)
			if accVest.Address == accVestExp.Address {
				require.EqualValues(t, accVest, accVestExp)
				found = true
			}
		}
		require.True(t, found, "not found: "+accVest.Address)

	}

	nullify.Fill(&genesisState)
	nullify.Fill(got)

}
