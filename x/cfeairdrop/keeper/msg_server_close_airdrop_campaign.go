package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CloseAirdropCampaign(goCtx context.Context, msg *types.MsgCloseAirdropCampaign) (*types.MsgCloseAirdropCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.CloseAirdropCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.Burn,
		msg.CommunityPoolSend,
	); err != nil {
		return nil, err
	}

	return &types.MsgCloseAirdropCampaignResponse{}, nil
}
