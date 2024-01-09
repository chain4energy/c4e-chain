package keeper

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants register cfedistribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "nonnegative-coin-state",
		NonNegativeCoinStateInvariant(k))
	ir.RegisterRoute(types.ModuleName, "state-sum-balance-check",
		StateSumBalanceCheckInvariant(k))
}

// NonNegativeCoinStateInvariant checks that any coin state is negative
func NonNegativeCoinStateInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		states := k.GetAllStates(ctx)
		for _, state := range states {
			if err := state.IsNegative(); err != nil {
				return sdk.FormatInvariant(types.ModuleName, "nonnegative coin state", err.Error()), true
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "nonnegative coin state", "\tno negative coin states"), false
	}
}

// StateSumBalanceCheckInvariant checks that sum on state equal module account balance
func StateSumBalanceCheckInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		states := k.GetAllStates(ctx)
		remainsSum, err := types.StateSumIsInteger(states)
		if err != nil {
			return sdk.FormatInvariant(types.ModuleName, "state sum balance check", err.Error()), true
		}

		var broken bool

		distributorAccountCoins := k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		if remainsSum.IsZero() && distributorAccountCoins.IsZero() {
			ctx.Logger().Debug("Coin state and distributor account is empty possible start of blockchain")
			broken = false
		} else {
			broken = !remainsSum.IsEqual(distributorAccountCoins)
		}

		return sdk.FormatInvariant(types.ModuleName, "state sum balance check",
			fmt.Sprintf(
				"\tsum of states coins: %v\n"+
					"\tdistributor account balance: %v",
				remainsSum, distributorAccountCoins)), broken
	}
}
