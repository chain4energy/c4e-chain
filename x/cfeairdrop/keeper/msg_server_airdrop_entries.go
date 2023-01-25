package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAirdropEntries(goCtx context.Context, msg *types.MsgAddAirdropEntries) (*types.MsgAddAirdropEntriesResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "add airdrop entries message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper
	if err := keeper.AddUserAirdropEntries(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.AirdropEntries,
	); err != nil {
		return nil, err
	}

	return &types.MsgAddAirdropEntriesResponse{}, nil
}

func (k msgServer) DeleteAirdropEntry(goCtx context.Context, msg *types.MsgDeleteAirdropEntry) (*types.MsgDeleteAirdropEntryResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "delete airdrop entry message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper
	if err := keeper.DeleteUserAirdropEntry(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.UserAddress,
	); err != nil {
		return nil, err
	}

	return &types.MsgDeleteAirdropEntryResponse{}, nil
}