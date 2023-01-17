package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAirdropEntries(goCtx context.Context, msg *types.MsgAddAirdropEntries) (*types.MsgAddAirdropEntriesResponse, error) {
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
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	// TODO: add logic

	return &types.MsgDeleteAirdropEntryResponse{}, nil
}
