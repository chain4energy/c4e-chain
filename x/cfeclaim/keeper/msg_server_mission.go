package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddMission(goCtx context.Context, msg *types.MsgAddMission) (*types.MsgAddMissionResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "add mission to aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	mission, err := keeper.AddMission(ctx, msg.Owner, msg.CampaignId, msg.Name, msg.Description, msg.MissionType, *msg.Weight, msg.ClaimStartDate)
	if err != nil {
		k.Logger(ctx).Debug("add mission", "err", err.Error())
		return nil, err
	}

	return &types.MsgAddMissionResponse{MissionId: mission.Id}, nil
}
