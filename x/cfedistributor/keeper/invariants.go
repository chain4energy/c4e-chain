package keeper

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants register cfedistribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "state-sum-balance-check",
		StateSumBalanceCheckInvariant(k))
}

// StateSumBalanceCheckInvariant checks that sum on state equal module account balance
func StateSumBalanceCheckInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		statesSum := GetStatesSum(k, ctx)
		coinsStatesSum, change := statesSum.TruncateDecimal()
		if !change.IsZero() {
			return sdk.FormatInvariant(types.ModuleName, "state sum balance check",
				fmt.Sprintf(
					"\tthe sum of the states should be integer change: %v\n",
					change)), true
		}
		var broken bool

		distributorAccountCoins := k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		if coinsStatesSum.IsZero() && distributorAccountCoins.IsZero() {
			ctx.Logger().Debug("Coin state and distributor account is empty possible start of blockchain")
			broken = false
		} else {
			broken = coinsStatesSum.IsEqual(distributorAccountCoins)
		}

		return sdk.FormatInvariant(types.ModuleName, "state sum balance check",
			fmt.Sprintf(
				"\tsum of states coins: %v\n"+
					"\tdistributor account balance:          %v\n",
				coinsStatesSum, distributorAccountCoins)), broken
	}
}

func GetStatesSum(k Keeper, ctx sdk.Context) sdk.DecCoins {
	states := k.GetAllStates(ctx)
	statesSum := sdk.NewDecCoins()

	for _, state := range states {
		statesSum.Add(state.CoinsStates...)
	}

	return statesSum
}
