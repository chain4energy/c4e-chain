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
	keeper := k.Keeper

	if err := keeper.CreateCampaign(
		ctx,
		msg.Owner,
		msg.Name,
		msg.Description,
		msg.CampaignType,
		msg.FeegrantAmount,
		msg.InitialClaimFreeAmount,
		msg.StartTime,
		msg.EndTime,
		msg.LockupPeriod,
		msg.VestingPeriod,
	); err != nil {
		k.Logger(ctx).Debug("create campaign", "err", err.Error())
		return nil, err
	}

	event := &types.NewCampaign{
		Owner:                  msg.Owner,
		Name:                   msg.Name,
		Description:            msg.Description,
		CampaignType:           msg.CampaignType.String(),
		FeegrantAmount:         msg.FeegrantAmount.String(),
		InitialClaimFreeAmount: msg.InitialClaimFreeAmount.String(),
		Enabled:                "false",
		StartTime:              msg.StartTime.String(),
		EndTime:                msg.EndTime.String(),
		LockupPeriod:           msg.LockupPeriod.String(),
		VestingPeriod:          msg.VestingPeriod.String(),
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("create campaign emit event error", "event", event, "error", err.Error())
	}

	return &types.MsgCreateCampaignResponse{}, nil
}

func (k msgServer) EditCampaign(goCtx context.Context, msg *types.MsgEditCampaign) (*types.MsgEditCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "edit aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.EditCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.Name,
		msg.Description,
		msg.StartTime,
		msg.EndTime,
		msg.LockupPeriod,
		msg.VestingPeriod,
	); err != nil {
		return nil, err
	}

	return &types.MsgEditCampaignResponse{}, nil
}

func (k msgServer) RemoveCampaign(goCtx context.Context, msg *types.MsgRemoveCampaign) (*types.MsgRemoveCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "start claim campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.RemoveCampaign(
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

func (k msgServer) StartCampaign(goCtx context.Context, msg *types.MsgStartCampaign) (*types.MsgStartCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "start claim campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.StartCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
	); err != nil {
		return nil, err
	}

	event := &types.StartCampaign{
		Owner:      msg.Owner,
		CampaignId: strconv.FormatUint(msg.CampaignId, 10),
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("start campaign emit event error", "event", event, "error", err.Error())
	}

	return &types.MsgStartCampaignResponse{}, nil
}

func (k msgServer) CloseCampaign(goCtx context.Context, msg *types.MsgCloseCampaign) (*types.MsgCloseCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "close claim campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.CloseCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.CampaignCloseAction,
	); err != nil {
		return nil, err
	}

	event := &types.CloseCampaign{
		Owner:               msg.Owner,
		CampaignId:          strconv.FormatUint(msg.CampaignId, 10),
		CampaignCloseAction: msg.CampaignCloseAction.String(),
	}
	err := ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("close campaign emit event error", "event", event, "error", err.Error())
	}

	return &types.MsgCloseCampaignResponse{}, nil
}
