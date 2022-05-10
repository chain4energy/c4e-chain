package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
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
	accountInfo := k.authKeeper.GetAccount(ctx, accAddress)

	return &types.QueryGetAccountInfoResponse{AccAddress: accountInfo.GetAddress().String(), PubKey: accountInfo.GetPubKey().String()}, nil
}
