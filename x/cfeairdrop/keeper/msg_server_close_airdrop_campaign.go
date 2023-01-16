package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CloseAirdropCampaign(goCtx context.Context, msg *types.MsgCloseAirdropCampaign) (*types.MsgCloseAirdropCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCloseAirdropCampaignResponse{}, nil
}
