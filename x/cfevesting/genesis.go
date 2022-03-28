package cfevesting

import (
	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)



// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, ak types.AccountKeeper) {

	k.Logger(ctx).Info("Init genesis")
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.Logger(ctx).Info("Init genesis params: ")
	vestingTypes := types.VestingTypes{}
	for _, gVestingType := range genState.VestingTypes {
		vt := types.VestingType{
			Name:                 gVestingType.Name,
			LockupPeriod:         keeper.DurationFromUnits(keeper.PeriodUnit(gVestingType.LockupPeriodUnit), gVestingType.LockupPeriod),
			VestingPeriod:        keeper.DurationFromUnits(keeper.PeriodUnit(gVestingType.VestingPeriodUnit), gVestingType.VestingPeriod),
			TokenReleasingPeriod: keeper.DurationFromUnits(keeper.PeriodUnit(gVestingType.TokenReleasingPeriodUnit), gVestingType.TokenReleasingPeriod),
			DelegationsAllowed:   gVestingType.DelegationsAllowed,
		}
		vestingTypes.VestingTypes = append(vestingTypes.VestingTypes, &vt)
	}

	k.SetVestingTypes(ctx, vestingTypes)
	allVestings := genState.AccountVestingsList.Vestings

	for _, av := range allVestings {
		k.SetAccountVestings(ctx, *av)
	}
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	vestingTypes := k.GetVestingTypes(ctx)
	genesis.VestingTypes = keeper.ConvertVestingTypesToGenesisVestingTypes(&vestingTypes)
	allVestings := k.GetAllAccountVestings(ctx)

	for i := 0; i < len(allVestings); i++ {
		genesis.AccountVestingsList.Vestings = append(genesis.AccountVestingsList.Vestings, &allVestings[i])
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}


