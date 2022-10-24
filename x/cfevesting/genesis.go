package cfevesting

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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
	// Set all the vestingAccount
	for _, elem := range genState.VestingAccountList {
		k.Logger(ctx).Debug("set vesting account", "vestingAccount", elem)
		k.SetVestingAccount(ctx, elem)
	}

	// Set vestingAccount count
	k.SetVestingAccountCount(ctx, genState.VestingAccountCount)
	k.Logger(ctx).Debug("set vesting account count", "vestingAccountCount", genState.VestingAccountCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	vestingTypes := types.VestingTypes{}
	for _, gVestingType := range genState.VestingTypes {
		lockupPeriod, err := types.DurationFromUnits(types.PeriodUnit(gVestingType.LockupPeriodUnit), gVestingType.LockupPeriod)
		if err != nil {
			k.Logger(ctx).Error("init genesis lockup period duration error", "unit", gVestingType.LockupPeriodUnit, "period", gVestingType.LockupPeriod)
			panic(sdkerrors.Wrapf(err, "init genesis lockup period duration error: unit: %s", gVestingType.LockupPeriodUnit))
		}
		vestingPeriod, err := types.DurationFromUnits(types.PeriodUnit(gVestingType.VestingPeriodUnit), gVestingType.VestingPeriod)
		if err != nil {
			k.Logger(ctx).Error("init genesis vesting period duration error", "unit", gVestingType.VestingPeriodUnit, "period", gVestingType.VestingPeriod)
			panic(sdkerrors.Wrapf(err, "init genesis lockup period duration error: unit: %s", gVestingType.VestingPeriodUnit))
		}
		vt := types.VestingType{
			Name:          gVestingType.Name,
			LockupPeriod:  lockupPeriod,
			VestingPeriod: vestingPeriod,
		}
		k.Logger(ctx).Debug("append vesting type", "vestingType", &vt)

		vestingTypes.VestingTypes = append(vestingTypes.VestingTypes, &vt)
	}

	k.SetVestingTypes(ctx, vestingTypes)

	allVestings := genState.Vestings

	for _, av := range allVestings {
		k.Logger(ctx).Debug("set account vesting pools", "accountVestingPool", av)
		k.SetAccountVestingPools(ctx, *av)
	}
	ak.GetModuleAccount(ctx, types.ModuleName)
}

func ValidateAccountsOnGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState,
	ak types.AccountKeeper, bk types.BankKeeper, sk types.StakingKeeper) error {
	accsVestings := genState.Vestings
	undelegableAmount := sdk.ZeroInt()
	delegableAmounts := make(map[string]sdk.Int)

	for _, accVestingPools := range accsVestings {
		for _, v := range accVestingPools.VestingPools {
			undelegableAmount = undelegableAmount.Add(v.LastModificationVested).Sub(v.LastModificationWithdrawn)
		}
	}

	mAcc := ak.GetModuleAccount(ctx, types.ModuleName)
	modBalance := bk.GetBalance(ctx, mAcc.GetAddress(), genState.Params.Denom)
	k.Logger(ctx).Debug("cfevesting validate accounts on genesis data", "undelegableAmount", undelegableAmount,
		"moduleBalance", modBalance.Amount, "moduleAccount", mAcc)

	if !undelegableAmount.Equal(modBalance.Amount) {
		k.Logger(ctx).Error("cfevesting module account balance not equal of sum of undelegable vestings error", "denom",
			genState.Params.Denom, "moduleBalance", modBalance.Amount, "moduleAccount", mAcc, "undelegableAmount",
			undelegableAmount)
		return fmt.Errorf("module: %s account balance of denom %s not equal of sum of undelegable vestings: %s <> %s",
			types.ModuleName, genState.Params.Denom, modBalance.Amount, undelegableAmount)
	}

	for delAddr, amount := range delegableAmounts {
		acc, err := sdk.AccAddressFromBech32(delAddr)
		if err != nil {
			k.Logger(ctx).Error("cfevesting account vestings delegable address error", "error", err.Error(),
				"delegableAddress", delAddr)
			return fmt.Errorf("account vestings delegable address: %s: %s", delAddr, err.Error())
		}
		accBalance := bk.GetBalance(ctx, acc, genState.Params.Denom)

		if !accBalance.Amount.LTE(amount) {
			k.Logger(ctx).Error("cfevesting delegable account balance is bigger than sum of delegable vestings error",
				"accountBalance", accBalance.Amount, "delegableVestings", amount)
			return fmt.Errorf("module: %s - delegable account: %s balance of denom %s is bigger than sum of delegable vestings: %s > %s",
				types.ModuleName, delAddr, genState.Params.Denom, accBalance.Amount, amount)
		}
	}
	return nil
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	vestingTypes := k.GetVestingTypes(ctx)
	genesis.VestingTypes = types.ConvertVestingTypesToGenesisVestingTypes(&vestingTypes)
	allVestings := k.GetAllAccountVestingPools(ctx)

	for i := 0; i < len(allVestings); i++ {
		genesis.Vestings = append(genesis.Vestings, &allVestings[i])
	}

	genesis.VestingAccountList = k.GetAllVestingAccount(ctx)
	genesis.VestingAccountCount = k.GetVestingAccountCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
