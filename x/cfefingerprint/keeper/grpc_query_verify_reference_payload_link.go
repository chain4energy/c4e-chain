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

	result, err := k.VerifyPayloadLink(ctx, req.ReferenceId, req.PayloadHash)
	if err != nil {
		k.Logger(ctx).Error("verify reference payload link - validation error", "error", err)
		return &types.QueryVerifyReferencePayloadLinkResponse{IsValid: "false"}, err
	}

	if result {
		return &types.QueryVerifyReferencePayloadLinkResponse{IsValid: "true"}, nil
	}

	// verification failed
	ctx.Logger().Debug("PayloadLink verification failed: payloadHash:", req.PayloadHash)

	return &types.QueryVerifyReferencePayloadLinkResponse{IsValid: "false"}, nil
}
