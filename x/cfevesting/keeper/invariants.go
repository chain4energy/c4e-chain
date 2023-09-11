package keeper

import (
	"cosmossdk.io/math"
	"fmt"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const NONNEGATIVE_AMOUNTS_INVARIANT = "nonnegative vesting pool amounts"

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
		allVestingPools := k.GetAllAccountVestingPools(ctx)

		for _, accountVestingPools := range allVestingPools {
			for _, vestingPool := range accountVestingPools.VestingPools {
				if vestingPool.Withdrawn.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, NONNEGATIVE_AMOUNTS_INVARIANT,
						fmt.Sprintf("\tnegative Withdrawn %s in vesting pool: %s for address: %s", vestingPool.Withdrawn, vestingPool.Name, accountVestingPools.Owner)), true
				} else if vestingPool.Sent.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, NONNEGATIVE_AMOUNTS_INVARIANT,
						fmt.Sprintf("\tnegative Sent %s in vesting pool: %s for address: %s", vestingPool.Sent, vestingPool.Name, accountVestingPools.Owner)), true
				} else if vestingPool.InitiallyLocked.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, NONNEGATIVE_AMOUNTS_INVARIANT,
						fmt.Sprintf("\tnegative InitiallyLocked %s in vesting pool: %s for address: %s", vestingPool.InitiallyLocked, vestingPool.Name, accountVestingPools.Owner)), true
				}
			}
		}

		return sdk.FormatInvariant(types.ModuleName, NONNEGATIVE_AMOUNTS_INVARIANT, "\tno negative amounts in vesting pools"), false
	}
}

func VestingPoolConsistentDataInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allVestingPools := k.GetAllAccountVestingPools(ctx)

		for _, accountVestingPools := range allVestingPools {
			for _, vestingPool := range accountVestingPools.VestingPools {
				if vestingPool.Withdrawn.Add(vestingPool.Sent).GT(vestingPool.InitiallyLocked) {
					return sdk.FormatInvariant(types.ModuleName, "vesting pool consistent data",
						fmt.Sprintf("\tWithdrawn (%s) + Sent (%s) GT InitiallyLocked (%s) in vesting pool: %s for address: %s",
							vestingPool.Withdrawn, vestingPool.Sent, vestingPool.InitiallyLocked, vestingPool.Name, accountVestingPools.Owner)), true
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

func getLockedSum(k Keeper, ctx sdk.Context) math.Int {
	allVestingPools := k.GetAllAccountVestingPools(ctx)
	sum := math.ZeroInt()
	for _, accountVestingPools := range allVestingPools {
		for _, vestingPool := range accountVestingPools.VestingPools {
			sum = sum.Add(vestingPool.GetCurrentlyLocked())
		}
	}
	return sum
}
