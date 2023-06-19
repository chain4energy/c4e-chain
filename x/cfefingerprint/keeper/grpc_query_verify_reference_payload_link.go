package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifyReferencePayloadLink(goCtx context.Context, req *types.QueryVerifyReferencePayloadLinkRequest) (*types.QueryVerifyReferencePayloadLinkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.VerifyPayloadLink(ctx, req.ReferenceId, req.PayloadHash); err != nil {
		k.Logger(ctx).Error("verify reference payload link error", "error", err)
		return &types.QueryVerifyReferencePayloadLinkResponse{IsValid: false}, err
	}

	return &types.QueryVerifyReferencePayloadLinkResponse{IsValid: true}, nil
}
