package keeper

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants register cfedistribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "claim-claims-left-sum-check",
		CampaignAmountLeftSumCheckInvariant(k))
}

// CampaignAmountLeftSumCheckInvariant checks that sum of claim claims left is equal to cfeaidrop module account balance
func CampaignAmountLeftSumCheckInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		claimClaimsLeftList := k.GetAllCampaignAmountLeft(ctx)

		if len(claimClaimsLeftList) == 0 {
			return sdk.FormatInvariant(types.ModuleName, "claim claims left sum check", "claim claims left sum is empty"), false
		}

		var claimClaimsLeftSum = sdk.NewCoins()
		for _, claimClaimsLeft := range claimClaimsLeftList {
			claimClaimsLeftSum = claimClaimsLeftSum.Add(claimClaimsLeft.Amount...)
		}
		cfeaidropAccountCoins := k.GetAccountCoinsForModuleAccount(ctx, types.ModuleName)
		if !cfeaidropAccountCoins.IsEqual(claimClaimsLeftSum) {
			return sdk.FormatInvariant(types.ModuleName, "claim claims left sum check",
				fmt.Sprintf("claim claims left sum is equal to cfeclaim module account balance (%v != %v)", claimClaimsLeftSum, cfeaidropAccountCoins)), true
		}
		return sdk.FormatInvariant(types.ModuleName, "claim claims left sum check", "claim claims left sum is equal to cfeclaim module account balance"), false
	}
}
