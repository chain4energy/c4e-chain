package keeper

import (
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants register cfedistribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "nonnegative-vesting-pool-amount",
		NonNegativeVestingPoolAmountsInvariant(k))
	ir.RegisterRoute(types.ModuleName, "vesting-pool-consistent-data",
		VestingPoolConsistentDataInvariant(k))
	ir.RegisterRoute(types.ModuleName, "module-account",
		ModuleAccountInvariant(k))
}

// NonNegativeCoinStateInvariant checks that any locked coins amount in vesting pools is non negative
func NonNegativeVestingPoolAmountsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allVestingPools := k.GetAllAccountVestings(ctx)

		for _, accountVestingPools := range allVestingPools {
			for _, vestingPool := range accountVestingPools.VestingPools {
				if vestingPool.LastModificationVested.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, "nonnegative vesting pool amounts",
						fmt.Sprintf("\tnegative LastModificationVested %s in vesting pool: %s for address: %s", vestingPool.LastModificationVested, vestingPool.Name, accountVestingPools.Address)), true
				} else if vestingPool.Withdrawn.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, "nonnegative vesting pool amounts",
						fmt.Sprintf("\tnegative Withdrawn %s in vesting pool: %s for address: %s", vestingPool.Withdrawn, vestingPool.Name, accountVestingPools.Address)), true
				} else if vestingPool.Sent.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, "nonnegative vesting pool amounts",
						fmt.Sprintf("\tnegative Sent %s in vesting pool: %s for address: %s", vestingPool.Sent, vestingPool.Name, accountVestingPools.Address)), true
				} else if vestingPool.LastModificationWithdrawn.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, "nonnegative vesting pool amounts",
						fmt.Sprintf("\tnegative LastModificationWithdrawn %s in vesting pool: %s for address: %s", vestingPool.LastModificationWithdrawn, vestingPool.Name, accountVestingPools.Address)), true
				} else if vestingPool.Vested.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, "nonnegative vesting pool amounts",
						fmt.Sprintf("\tnegative Vested %s in vesting pool: %s for address: %s", vestingPool.Vested, vestingPool.Name, accountVestingPools.Address)), true
				}
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "nonnegative vesting pool amounts", "\tno negative amounts in vesting pools"), false
	}
}

func VestingPoolConsistentDataInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allVestingPools := k.GetAllAccountVestings(ctx)

		for _, accountVestingPools := range allVestingPools {
			for _, vestingPool := range accountVestingPools.VestingPools {
				if vestingPool.LastModificationWithdrawn.GT(vestingPool.LastModificationVested) {
					return sdk.FormatInvariant(types.ModuleName, "vesting pool consistent data",
						fmt.Sprintf("\tLastModificationWithdrawn (%s) GT LastModificationVested (%s) in vesting pool: %s for address: %s",
							vestingPool.LastModificationWithdrawn, vestingPool.LastModificationVested, vestingPool.Name, accountVestingPools.Address)), true
				} else if vestingPool.Withdrawn.Add(vestingPool.Sent).GT(vestingPool.Vested) {
					return sdk.FormatInvariant(types.ModuleName, "vesting pool consistent data",
						fmt.Sprintf("\tWithdrawn (%s) + Sent (%s) GT Vested (%s) in vesting pool: %s for address: %s",
							vestingPool.Withdrawn, vestingPool.Sent, vestingPool.Vested, vestingPool.Name, accountVestingPools.Address)), true
				} else if !vestingPool.Vested.Sub(vestingPool.Withdrawn).Sub(vestingPool.Sent).Equal(vestingPool.LastModificationVested.Sub(vestingPool.LastModificationWithdrawn)) {
					return sdk.FormatInvariant(types.ModuleName, "vesting pool consistent data",
						fmt.Sprintf("\t Vested (%s) - Withdrawn (%s) - Sent (%s) <> LastModificationVested (%s) - LastModificationWithdrawn (%s) in vesting pool: %s for address: %s",
							vestingPool.Vested, vestingPool.Withdrawn, vestingPool.Sent, vestingPool.LastModificationVested, vestingPool.LastModificationWithdrawn, vestingPool.Name, accountVestingPools.Address)), true
				}
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "vesting pool consistent data", "\tno inconsistent vesting pools"), false
	}
}

// ModuleAccountInvariant checks that sum on locked in vesting pools equals to module account balance
func ModuleAccountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		sum := getLockedSum(k, ctx)

		account := k.account.GetModuleAccount(ctx, types.ModuleName)
		denom := k.GetParams(ctx).Denom
		balance := k.bank.GetBalance(ctx, account.GetAddress(), denom)

		if !balance.Amount.Equal(sum) {
			return sdk.FormatInvariant(types.ModuleName, "module account", fmt.Sprintf("\tamount (%s) inconsistent with vesting pools (%s)", balance.Amount, sum)), true
		}
		return sdk.FormatInvariant(types.ModuleName, "module account", "\tamount consistent with vesting pools"), false
	}
}

func getLockedSum(k Keeper, ctx sdk.Context) sdk.Int {
	allVestingPools := k.GetAllAccountVestings(ctx)
	sum := sdk.ZeroInt()
	for _, accountVestingPools := range allVestingPools {
		for _, vestingPool := range accountVestingPools.VestingPools {
			sum = sum.Add(vestingPool.LastModificationVested.Sub(vestingPool.LastModificationWithdrawn))
		}
	}
	return sum
}
