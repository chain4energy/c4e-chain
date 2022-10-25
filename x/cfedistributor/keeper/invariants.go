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
			for _, coinState := range state.CoinsStates {
				if coinState.IsNegative() {
					return sdk.FormatInvariant(types.ModuleName, "nonnegative coin state",
						fmt.Sprintf("\tnegative coin state %s in state %s", coinState, state.StateIdString())), true
				}
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "nonnegative coin state", "\tno negative coin states"), false
	}
}

// StateSumBalanceCheckInvariant checks that sum on state equal module account balance
func StateSumBalanceCheckInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		statesSum := getStatesSum(k, ctx)
		coinsStatesSum, change := statesSum.TruncateDecimal()
		if !change.IsZero() {
			changeString := change.String()
			fmt.Println(statesSum.String())
			fmt.Println(changeString)
			return sdk.FormatInvariant(types.ModuleName, "state sum balance check",
				fmt.Sprintf(
					"\tthe sum of the states should be integer: sum: %v",
					statesSum)), true
		}
		var broken bool

		distributorAccountCoins := k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		if coinsStatesSum.IsZero() && distributorAccountCoins.IsZero() {
			ctx.Logger().Debug("Coin state and distributor account is empty possible start of blockchain")
			broken = false
		} else {
			broken = !coinsStatesSum.IsEqual(distributorAccountCoins)
		}

		return sdk.FormatInvariant(types.ModuleName, "state sum balance check",
			fmt.Sprintf(
				"\tsum of states coins: %v\n"+
					"\tdistributor account balance: %v",
				coinsStatesSum, distributorAccountCoins)), broken
	}
}

func getStatesSum(k Keeper, ctx sdk.Context) sdk.DecCoins {
	states := k.GetAllStates(ctx)
	statesSum := sdk.NewDecCoins()

	for _, state := range states {
		statesSum = statesSum.Add(state.CoinsStates...)
	}

	return statesSum
}
