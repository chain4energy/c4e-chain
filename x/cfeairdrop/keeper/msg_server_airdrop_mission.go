package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddMissionToAidropCampaign(goCtx context.Context, msg *types.MsgAddMissionToAidropCampaign) (*types.MsgAddMissionToAidropCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.AddMissionToAirdropCampaign(ctx, msg.Owner, msg.CampaignId, msg.Name, msg.Description, msg.MissionType, msg.Weight, msg.ClaimStartDate); err != nil {
		return nil, err
	}
	return &types.MsgAddMissionToAidropCampaignResponse{}, nil
}
