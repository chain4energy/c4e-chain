package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateCampaign(goCtx context.Context, msg *types.MsgCreateCampaign) (*types.MsgCreateCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, err := k.Keeper.CreateCampaign(
		ctx,
		msg.Owner,
		msg.Name,
		msg.Description,
		msg.CampaignType,
		msg.RemovableClaimRecords,
		*msg.FeegrantAmount,
		*msg.InitialClaimFreeAmount,
		*msg.Free,
		*msg.StartTime,
		*msg.EndTime,
		*msg.LockupPeriod,
		*msg.VestingPeriod,
		msg.VestingPoolName,
	)
	if err != nil {
		k.Logger(ctx).Debug("create campaign", "err", err.Error())
		return nil, err
	}

	return &types.MsgCreateCampaignResponse{CampaignId: campaign.Id}, nil
}

func (k msgServer) RemoveCampaign(goCtx context.Context, msg *types.MsgRemoveCampaign) (*types.MsgRemoveCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "start claim campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.RemoveCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
	); err != nil {
		return nil, err
	}

	return &types.MsgRemoveCampaignResponse{}, nil
}

func (k msgServer) EnableCampaign(goCtx context.Context, msg *types.MsgEnableCampaign) (*types.MsgEnableCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "start claim campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.EnableCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.StartTime,
		msg.EndTime,
	); err != nil {
		k.Logger(ctx).Debug("enable campaign", "err", err.Error())
		return nil, err
	}

	return &types.MsgEnableCampaignResponse{}, nil
}

func (k msgServer) CloseCampaign(goCtx context.Context, msg *types.MsgCloseCampaign) (*types.MsgCloseCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "close claim campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.CloseCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
	); err != nil {
		k.Logger(ctx).Debug("close campaign", "err", err.Error())
		return nil, err
	}

	return &types.MsgCloseCampaignResponse{}, nil
}
