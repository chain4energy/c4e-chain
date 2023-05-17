package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
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
		msg.FeegrantAmount,
		msg.InitialClaimFreeAmount,
		msg.Free,
		msg.StartTime,
		msg.EndTime,
		msg.LockupPeriod,
		msg.VestingPeriod,
		msg.VestingPoolName,
	)
	if err != nil {
		k.Logger(ctx).Debug("create campaign", "err", err.Error())
		return nil, err
	}

	event := &types.NewCampaign{
		Owner:                  campaign.Owner,
		Name:                   campaign.Name,
		Description:            campaign.Description,
		CampaignType:           campaign.CampaignType.String(),
		FeegrantAmount:         campaign.FeegrantAmount.String(),
		InitialClaimFreeAmount: campaign.InitialClaimFreeAmount.String(),
		Enabled:                "false",
		StartTime:              campaign.StartTime.String(),
		EndTime:                campaign.EndTime.String(),
		LockupPeriod:           campaign.LockupPeriod.String(),
		VestingPeriod:          campaign.VestingPeriod.String(),
		VestingPoolName:        campaign.VestingPoolName,
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("create campaign emit event error", "event", event, "error", err.Error())
	}

	return &types.MsgCreateCampaignResponse{}, nil
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

	event := &types.RemoveCampaign{
		Owner:      msg.Owner,
		CampaignId: strconv.FormatUint(msg.CampaignId, 10),
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("remove campaign emit event error", "event", event, "error", err.Error())
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
		k.Logger(ctx).Debug("start campaign", "err", err.Error())
		return nil, err
	}

	event := &types.EnableCampaign{
		Owner:      msg.Owner,
		CampaignId: strconv.FormatUint(msg.CampaignId, 10),
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("start campaign emit event error", "event", event, "error", err.Error())
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

	event := &types.CloseCampaign{
		Owner:      msg.Owner,
		CampaignId: strconv.FormatUint(msg.CampaignId, 10),
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("close campaign emit event error", "event", event, "error", err.Error())
	}

	return &types.MsgCloseCampaignResponse{}, nil
}
