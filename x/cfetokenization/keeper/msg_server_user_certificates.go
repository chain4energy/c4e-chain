package keeper

import (
	"context"
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUserCertificates(goCtx context.Context, msg *types.MsgCreateUserCertificates) (*types.MsgCreateUserCertificatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var userCertificates = types.UserCertificates{
		Owner:        msg.Owner,
		Certificates: msg.Certificates,
	}

	id := k.AppendUserCertificates(
		ctx,
		userCertificates,
	)

	return &types.MsgCreateUserCertificatesResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateUserCertificates(goCtx context.Context, msg *types.MsgUpdateUserCertificates) (*types.MsgUpdateUserCertificatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var userCertificates = types.UserCertificates{
		Owner:        msg.Owner,
		Id:           msg.Id,
		Certificates: msg.Certificates,
	}

	// Checks that the element exists
	val, found := k.GetUserCertificates(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Owner != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetUserCertificates(ctx, userCertificates)

	return &types.MsgUpdateUserCertificatesResponse{}, nil
}

func (k msgServer) DeleteUserCertificates(goCtx context.Context, msg *types.MsgDeleteUserCertificates) (*types.MsgDeleteUserCertificatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetUserCertificates(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Owner != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUserCertificates(ctx, msg.Id)

	return &types.MsgDeleteUserCertificatesResponse{}, nil
}
