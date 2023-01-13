package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAirdropEntry(goCtx context.Context, msg *types.MsgCreateAirdropEntry) (*types.MsgCreateAirdropEntryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var airdropEntry = types.AirdropEntry{
		Address: msg.Address,
		Amount:  msg.Amount,
	}

	id := k.AppendAirdropEntry(
		ctx,
		airdropEntry,
	)

	return &types.MsgCreateAirdropEntryResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAirdropEntry(goCtx context.Context, msg *types.MsgUpdateAirdropEntry) (*types.MsgUpdateAirdropEntryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var airdropEntry = types.AirdropEntry{
		Id:      msg.Id,
		Address: msg.Address,
		Amount:  msg.Amount,
	}

	// Checks that the element exists
	//val, found := k.GetAirdropEntry(ctx, msg.Id)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	//}

	// Checks if the msg creator is the same as the current owner
	//if msg.Creator != val.Creator {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	//}

	k.SetAirdropEntry(ctx, airdropEntry)

	return &types.MsgUpdateAirdropEntryResponse{}, nil
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
