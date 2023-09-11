package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "claim message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, err := k.Keeper.Claim(
		ctx,
		msg.CampaignId,
		msg.MissionId,
		msg.Claimer,
	)
	if err != nil {
		k.Logger(ctx).Debug("claim", "err", err.Error())
		return nil, err
	}

	return &types.MsgClaimResponse{Amount: amount}, nil
}

func (k msgServer) InitialClaim(goCtx context.Context, msg *types.MsgInitialClaim) (*types.MsgInitialClaimResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "initial claim message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, err := k.Keeper.InitialClaim(
		ctx,
		msg.Claimer,
		msg.CampaignId,
		msg.DestinationAddress,
	)
	if err != nil {
		k.Logger(ctx).Debug("initial claim", "err", err.Error())
		return nil, err
	}

	return &types.MsgInitialClaimResponse{Amount: amount}, nil
}
