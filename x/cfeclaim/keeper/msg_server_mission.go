package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddMissionToCampaign(goCtx context.Context, msg *types.MsgAddMissionToCampaign) (*types.MsgAddMissionToCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "add mission to aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.AddMissionToCampaign(ctx, msg.Owner, msg.CampaignId, msg.Name, msg.Description, msg.MissionType, *msg.Weight, msg.ClaimStartDate); err != nil {
		return nil, err
	}

	return &types.MsgAddMissionToCampaignResponse{}, nil
}
