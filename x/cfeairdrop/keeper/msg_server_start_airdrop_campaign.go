package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StartAirdropCampaign(goCtx context.Context, msg *types.MsgStartAirdropCampaign) (*types.MsgStartAirdropCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.StartAirdropCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
	); err != nil {
		return nil, err
	}

	return &types.MsgStartAirdropCampaignResponse{}, nil
}
