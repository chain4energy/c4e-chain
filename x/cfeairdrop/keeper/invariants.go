package keeper

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants register cfedistribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "airdrop-claims-left-sum-check",
		AirdropClaimsLeftSumCheckInvariant(k))
}

// AirdropClaimsLeftSumCheckInvariant checks that sum of airdrop claims left is equal to cfeaidrop module account balance
func AirdropClaimsLeftSumCheckInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		airdropClaimsLeftList := k.GetAllAirdropClaimsLeft(ctx)
		airdropClaimsLeftSum := sdk.NewCoins()
		for _, airdropClaimsLeft := range airdropClaimsLeftList {
			airdropClaimsLeftSum.Add(airdropClaimsLeft.Amount)
		}
		cfeaidropAccountCoins := k.GetAccountCoinsForModuleAccount(ctx, types.ModuleName)
		if !cfeaidropAccountCoins.IsEqual(airdropClaimsLeftSum) {
			return sdk.FormatInvariant(types.ModuleName, "airdrop claims left sum check",
				fmt.Sprintf("airdrop claims left sum is equal to cfeairdrop module account balance ( %s != %s", airdropClaimsLeftSum.String(), cfeaidropAccountCoins.String())), true
		}
		return sdk.FormatInvariant(types.ModuleName, "airdrop claims left sum check", "airdrop claims left sum is equal to cfeairdrop module account balance"), false
	}
}
