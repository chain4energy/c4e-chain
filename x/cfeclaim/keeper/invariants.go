package keeper

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants register cfedistribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "campaigns-current-amount",
		CampaignCurrentAmountSumCheckInvariant(k))
}

// TODO: add if reservation amount = campaignCurrentAmount
// CampaignCurrentAmountSumCheckInvariant checks that sum of claim claims left is equal to cfeaidrop module account balance
func CampaignCurrentAmountSumCheckInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		campaigns := k.GetCampaigns(ctx)
		if len(campaigns) == 0 {
			return sdk.FormatInvariant(types.ModuleName, "campaigns current amount sum", "campaigns list is empty"), false
		}

		var vestingCampaignsCurrentAmount = sdk.NewCoins()
		var lockedInReservations = math.ZeroInt()
		var defaultCampaignsCurrentAmount = sdk.NewCoins()

		for _, campaign := range campaigns {
			if campaign.CampaignType == types.VestingPoolCampaign {
				vestingCampaignsCurrentAmount = vestingCampaignsCurrentAmount.Add(campaign.CampaignCurrentAmount...)
				reservation, err := k.vestingKeeper.GetVestingPoolReservation(ctx, campaign.Owner, campaign.VestingPoolName, campaign.Id)
				if err == nil {
					lockedInReservations = lockedInReservations.Add(reservation.Amount)
				}
			} else if campaign.CampaignType == types.DefaultCampaign {
				defaultCampaignsCurrentAmount = defaultCampaignsCurrentAmount.Add(campaign.CampaignCurrentAmount...)
			}
		}
		cfeaidropAccountCoins := k.GetAccountCoinsForModuleAccount(ctx, types.ModuleName)
		if !cfeaidropAccountCoins.IsEqual(defaultCampaignsCurrentAmount) {
			return sdk.FormatInvariant(types.ModuleName, "campaigns current amount sum",
				fmt.Sprintf("campaigns current amount sum is not equal to cfeclaim module account balance (%v != %v)", defaultCampaignsCurrentAmount, cfeaidropAccountCoins)), true
		}

		if !vestingCampaignsCurrentAmount.AmountOf(k.vestingKeeper.Denom(ctx)).Equal(lockedInReservations) {
			return sdk.FormatInvariant(types.ModuleName, "campaigns current amount sum",
				fmt.Sprintf("campaigns current amount sum is not equal to lock tokens in vesting pools reservations (%v != %v)", vestingCampaignsCurrentAmount, lockedInReservations)), true
		}

		return sdk.FormatInvariant(types.ModuleName, "campaigns current amount sum", "claim claims left sum is equal to cfeclaim module account balance"), false
	}
}
