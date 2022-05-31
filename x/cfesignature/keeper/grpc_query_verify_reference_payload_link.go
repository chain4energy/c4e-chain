package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifyReferencePayloadLink(goCtx context.Context, req *types.QueryVerifyReferencePayloadLinkRequest) (*types.QueryVerifyReferencePayloadLinkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	// createStorageKeyRequest := &types.QueryCreateStorageKeyRequest{TargetAccAddress: targetAccAddress, ReferenceId: referenceID}
	// storageKey, err := k.CreateStorageKey(goCtx, createStorageKeyRequest)
	// TODO: invoke k.getReferencePayloadLink()

	return &types.QueryVerifyReferencePayloadLinkResponse{}, nil
}
