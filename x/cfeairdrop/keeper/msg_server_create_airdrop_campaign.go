package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAirdropCampaign(goCtx context.Context, msg *types.MsgCreateAirdropCampaign) (*types.MsgCreateAirdropCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.CreateAidropCampaign(ctx, msg.Creator, msg.Owner, msg.); err != nil {
		return nil, err
	}


	return &types.MsgCreateAirdropCampaignResponse{}, nil
}
