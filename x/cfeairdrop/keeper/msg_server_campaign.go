package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateCampaign(goCtx context.Context, msg *types.MsgCreateCampaign) (*types.MsgCreateCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.CreateAidropCampaign(
		ctx,
		msg.Owner,
		msg.Name,
		msg.Description,
		msg.FeegrantAmount,
		msg.InitialClaimFreeAmount,
		msg.StartTime,
		msg.EndTime,
		msg.LockupPeriod,
		msg.VestingPeriod,
	); err != nil {
		return nil, err
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
	defer telemetry.IncrCounter(1, types.ModuleName, "start airdrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.RemoveCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
	); err != nil {
		return nil, err
	}

	return &types.MsgRemoveCampaignResponse{}, nil
}

func (k msgServer) StartCampaign(goCtx context.Context, msg *types.MsgStartCampaign) (*types.MsgStartCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "start airdrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.StartCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
	); err != nil {
		return nil, err
	}

	return &types.MsgStartCampaignResponse{}, nil
}

func (k msgServer) CloseCampaign(goCtx context.Context, msg *types.MsgCloseCampaign) (*types.MsgCloseCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "close airdrop campaign message")
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

	return &types.MsgCloseCampaignResponse{}, nil
}
