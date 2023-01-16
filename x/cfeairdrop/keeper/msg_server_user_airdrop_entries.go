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

	// Checks that the element exists
	//val, found := k.GetAirdropEntry(ctx, msg.Id)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	//}

	// Checks if the msg creator is the same as the current owner
	//if msg.Creator != val.Creator {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	//}

	k.RemoveAirdropEntry(ctx, msg.Id)

	return &types.MsgDeleteAirdropEntryResponse{}, nil
}
