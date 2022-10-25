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

	allAccountVestingPools := genState.AccountVestingPools

	for _, av := range allAccountVestingPools {
		k.Logger(ctx).Debug("set account vesting pools", "accountVestingPool", av)
		k.SetAccountVestingPools(ctx, *av)
	}
	ak.GetModuleAccount(ctx, types.ModuleName)
}

func ValidateAccountsOnGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState,
	ak types.AccountKeeper, bk types.BankKeeper, sk types.StakingKeeper) error {
	accsVestingPools := genState.AccountVestingPools
	vestingPoolsAmount := sdk.ZeroInt()

	for _, accVestingPools := range accsVestingPools {
		for _, v := range accVestingPools.VestingPools {
			vestingPoolsAmount = vestingPoolsAmount.Add(v.LastModificationVested).Sub(v.LastModificationWithdrawn)
		}
	}

	mAcc := ak.GetModuleAccount(ctx, types.ModuleName)
	modBalance := bk.GetBalance(ctx, mAcc.GetAddress(), genState.Params.Denom)
	k.Logger(ctx).Debug("cfevesting validate accounts on genesis data", "vestingPoolsAmount", vestingPoolsAmount,
		"moduleBalance", modBalance.Amount, "moduleAccount", mAcc)

	if !vestingPoolsAmount.Equal(modBalance.Amount) {
		k.Logger(ctx).Error("cfevesting module account balance not equal of sum of vesting pools error", "denom",
			genState.Params.Denom, "moduleBalance", modBalance.Amount, "moduleAccount", mAcc, "vestingPoolsAmount",
			vestingPoolsAmount)
		return fmt.Errorf("module: %s account balance of denom %s not equal of sum of vesting pools: %s <> %s",
			types.ModuleName, genState.Params.Denom, modBalance.Amount, vestingPoolsAmount)
	}

	return nil
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	vestingTypes := k.GetVestingTypes(ctx)
	genesis.VestingTypes = types.ConvertVestingTypesToGenesisVestingTypes(&vestingTypes)
	allAccountVestingPools := k.GetAllAccountVestingPools(ctx)

	for i := 0; i < len(allAccountVestingPools); i++ {
		genesis.AccountVestingPools = append(genesis.AccountVestingPools, &allAccountVestingPools[i])
	}

	genesis.VestingAccountList = k.GetAllVestingAccount(ctx)
	genesis.VestingAccountCount = k.GetVestingAccountCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
