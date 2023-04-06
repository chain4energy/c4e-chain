package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "claim message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.Claim(
		ctx,
		msg.CampaignId,
		msg.MissionId,
		msg.Claimer,
	); err != nil {
		return nil, err
	}

	return &types.MsgClaimResponse{}, nil
}

func (k msgServer) InitialClaim(goCtx context.Context, msg *types.MsgInitialClaim) (*types.MsgInitialClaimResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "initial claim message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.InitialClaim(
		ctx,
		msg.Claimer,
		msg.CampaignId,
		msg.AddressToClaim,
	); err != nil {
		return nil, err
	}

	return &types.MsgInitialClaimResponse{}, nil
}
