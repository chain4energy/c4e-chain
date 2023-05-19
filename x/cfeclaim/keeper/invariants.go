package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants register cfedistribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "claim-claims-left-sum-check",
		CampaignCurrentAmountSumCheckInvariant(k))
}

// TODO: add if reservation amount = campaignCurrentAmount
// CampaignCurrentAmountSumCheckInvariant checks that sum of claim claims left is equal to cfeaidrop module account balance
func CampaignCurrentAmountSumCheckInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		claimClaimsLeftList := k.GetAllCampaignCurrentAmount(ctx)
		campaigns := k.GetCampaigns(ctx)
		if len(claimClaimsLeftList) == 0 {
			return sdk.FormatInvariant(types.ModuleName, "claim claims left sum check", "claim claims left sum is empty"), false
		}

		var claimClaimsLeftSum = sdk.NewCoins()
		for _, claimClaimsLeft := range claimClaimsLeftList {
			campaignType, err := findCampaignType(campaigns, claimClaimsLeft.CampaignId)
			if err != nil {
				return sdk.FormatInvariant(types.ModuleName, "claim claims left sum check",
					err.Error()), true
			}
			if *campaignType != types.VestingPoolCampaign {
				claimClaimsLeftSum = claimClaimsLeftSum.Add(claimClaimsLeft.Amount...)
			}

		}
		cfeaidropAccountCoins := k.GetAccountCoinsForModuleAccount(ctx, types.ModuleName)
		if !cfeaidropAccountCoins.IsEqual(claimClaimsLeftSum) {
			return sdk.FormatInvariant(types.ModuleName, "claim claims left sum check",
				fmt.Sprintf("claim claims left sum is equal to cfeclaim module account balance (%v != %v)", claimClaimsLeftSum, cfeaidropAccountCoins)), true
		}
		return sdk.FormatInvariant(types.ModuleName, "claim claims left sum check", "claim claims left sum is equal to cfeclaim module account balance"), false
	}
}

func findCampaignType(campaigns []types.Campaign, campaignId uint64) (*types.CampaignType, error) {
	for _, campaign := range campaigns {
		if campaign.Id == campaignId {
			return &campaign.CampaignType, nil
		}
	}
	return nil, errors.Wrapf(c4eerrors.ErrNotExists, "campaign with id: %d doesn't exist", campaignId)
}
