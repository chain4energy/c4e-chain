package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(k.Campaigns(ctx))
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) Campaigns(ctx sdk.Context) (res []*types.Campaign) {
	k.paramstore.Get(ctx, types.KeyCampaigns, &res)
	return
} //

func (k Keeper) Campaign(ctx sdk.Context, campaignId uint64) *types.Campaign {
	campaigns := k.Campaigns(ctx)
	for _, campaign := range campaigns {
		if campaign.CampaignId == campaignId {
			return campaign
		}
	}
	return nil
} //
