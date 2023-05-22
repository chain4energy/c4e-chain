package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAccountInfo(goCtx context.Context, req *types.QueryGetAccountInfoRequest) (*types.QueryGetAccountInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	accAddress, _ := sdk.AccAddressFromBech32(req.AccAddressString)
	// if account is not found GetAccount method returns nil
	accountInfo := k.authKeeper.GetAccount(ctx, accAddress)

	// TODO: add errorCode and errorMessage to QueryGetAccountInfoResponse
	if accountInfo == nil {
		return &types.QueryGetAccountInfoResponse{AccAddress: "Account Not found", PubKey: ""}, nil
	}

	return &types.QueryGetAccountInfoResponse{AccAddress: accountInfo.GetAddress().String(), PubKey: accountInfo.GetPubKey().String()}, nil
}
