package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAirdropCampaign(goCtx context.Context, msg *types.MsgCreateAirdropCampaign) (*types.MsgCreateAirdropCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.CreateAidropCampaign(
		ctx,
		msg.Owner,
		msg.Name,
		msg.Description,
		msg.StartTime,
		msg.EndTime,
		msg.LockupPeriod,
		msg.VestingPeriod,
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateAirdropCampaignResponse{}, nil
}

func (k msgServer) EditAirdropCampaign(goCtx context.Context, msg *types.MsgEditAirdropCampaign) (*types.MsgEditAirdropCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "edit aidrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.EditAirdropCampaign(
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

	return &types.MsgEditAirdropCampaignResponse{}, nil
}

func (k msgServer) CloseAirdropCampaign(goCtx context.Context, msg *types.MsgCloseAirdropCampaign) (*types.MsgCloseAirdropCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "close airdrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.CloseAirdropCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.AirdropCloseAction,
	); err != nil {
		return nil, err
	}

	return &types.MsgCloseAirdropCampaignResponse{}, nil
}

func (k msgServer) StartAirdropCampaign(goCtx context.Context, msg *types.MsgStartAirdropCampaign) (*types.MsgStartAirdropCampaignResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "start airdrop campaign message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	if err := keeper.StartAirdropCampaign(
		ctx,
		msg.Owner,
		msg.CampaignId,
	); err != nil {
		return nil, err
	}

	return &types.MsgStartAirdropCampaignResponse{}, nil
}
