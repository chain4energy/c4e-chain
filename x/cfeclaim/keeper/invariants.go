package keeper

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const campaignCurrentAmountSumInvariant = "campaigns current amount sum"

// RegisterInvariants register cfedistribution invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "campaigns-current-amount",
		CampaignCurrentAmountSumCheckInvariant(k))
}

// CampaignCurrentAmountSumCheckInvariant checks that sum of claim claims left is equal to cfeaidrop module account balance
func CampaignCurrentAmountSumCheckInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		campaigns := k.GetAllCampaigns(ctx)
		if len(campaigns) == 0 {
			return sdk.FormatInvariant(types.ModuleName, campaignCurrentAmountSumInvariant, "campaigns list is empty"), false
		}

		var vestingCampaignsCurrentAmount sdk.Coins
		var lockedInReservations = math.ZeroInt()
		var defaultCampaignsCurrentAmount sdk.Coins

		for _, campaign := range campaigns {
			if campaign.CampaignType == types.VestingPoolCampaign {
				vestingCampaignsCurrentAmount = vestingCampaignsCurrentAmount.Add(campaign.CampaignCurrentAmount...)
				reservation, err := k.vestingKeeper.MustGetVestingPoolReservation(ctx, campaign.Owner, campaign.VestingPoolName, campaign.Id)
				if err == nil {
					lockedInReservations = lockedInReservations.Add(reservation.Amount)
				}
			} else if campaign.CampaignType == types.DefaultCampaign {
				defaultCampaignsCurrentAmount = defaultCampaignsCurrentAmount.Add(campaign.CampaignCurrentAmount...)
			}
		}
		cfeaidropAccountCoins := k.GetAccountCoinsForModuleAccount(ctx, types.ModuleName)
		if !cfeaidropAccountCoins.IsEqual(defaultCampaignsCurrentAmount) {
			return sdk.FormatInvariant(types.ModuleName, campaignCurrentAmountSumInvariant,
				fmt.Sprintf("campaigns current amount sum is not equal to cfeclaim module account balance (%v != %v)", defaultCampaignsCurrentAmount, cfeaidropAccountCoins)), true
		}

		if !vestingCampaignsCurrentAmount.IsEqual(sdk.NewCoins(sdk.NewCoin(k.vestingKeeper.Denom(ctx), lockedInReservations))) {
			return sdk.FormatInvariant(types.ModuleName, campaignCurrentAmountSumInvariant,
				fmt.Sprintf("campaigns current amount sum is not equal to lock tokens in vesting pools reservations (%v != %v)", vestingCampaignsCurrentAmount, lockedInReservations)), true
		}

		return sdk.FormatInvariant(types.ModuleName, campaignCurrentAmountSumInvariant, "claim claims left sum is equal to cfeclaim module account balance"), false
	}
}
