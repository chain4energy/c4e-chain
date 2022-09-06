package cfevesting

import (
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, ak types.AccountKeeper, bk types.BankKeeper, sk types.StakingKeeper) {
	err := ValidateAccountsOnGenesis(ctx, k, genState, ak, bk, sk)
	if err != nil {
		panic(err)
	}
	k.Logger(ctx).Info("Init genesis")
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.Logger(ctx).Info("Init genesis params: ")
	vestingTypes := types.VestingTypes{}
	for _, gVestingType := range genState.VestingTypes {
		vt := types.VestingType{
			Name:          gVestingType.Name,
			LockupPeriod:  keeper.DurationFromUnits(keeper.PeriodUnit(gVestingType.LockupPeriodUnit), gVestingType.LockupPeriod),
			VestingPeriod: keeper.DurationFromUnits(keeper.PeriodUnit(gVestingType.VestingPeriodUnit), gVestingType.VestingPeriod),
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

func ValidateAccountsOnGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState,
	ak types.AccountKeeper, bk types.BankKeeper, sk types.StakingKeeper) error {
	accsVestings := genState.AccountVestingsList.Vestings
	undelegableAmount := sdk.ZeroInt()
	delegableAmounts := make(map[string]sdk.Int)

	for _, accVestings := range accsVestings {
		for _, v := range accVestings.VestingPools {
			undelegableAmount = undelegableAmount.Add(v.LastModificationVested)
		}
	}

	mAcc := ak.GetModuleAccount(ctx, types.ModuleName)
	modBalance := bk.GetBalance(ctx, mAcc.GetAddress(), genState.Params.Denom)
	if !undelegableAmount.Equal(modBalance.Amount) {
		return fmt.Errorf("module: %s account balance of denom %s not equal of sum of undelegable vestings: %s <> %s",
			types.ModuleName, genState.Params.Denom, modBalance.Amount.String(), undelegableAmount.String())
	}

	for delAddr, amount := range delegableAmounts {
		acc, err := sdk.AccAddressFromBech32(delAddr)
		if err != nil {
			return fmt.Errorf("account vestings delegable address: %s: %s", delAddr, err.Error())
		}
		accBalance := bk.GetBalance(ctx, acc, genState.Params.Denom)

		if !accBalance.Amount.LTE(amount) {
			return fmt.Errorf("module: %s - delegable account: %s balance of denom %s is bigger than sum of delegable vestings: %s > %s",
				types.ModuleName, delAddr, genState.Params.Denom, accBalance.Amount.String(), amount.String())
		}
	}
	return nil
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
