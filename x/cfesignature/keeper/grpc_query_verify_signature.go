package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/chain4energy/c4e-chain/x/cfesignature/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifySignature(goCtx context.Context, req *types.QueryVerifySignatureRequest) (*types.QueryVerifySignatureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	referenceId := req.ReferenceId
	targetAccAddress := req.TargetAccAddress

	var signature types.Signature

	param := types.QueryCreateStorageKeyRequest{TargetAccAddress: req.TargetAccAddress, ReferenceId: req.ReferenceId}

	// fetch storage keys for signature and document hash
	storageKeySignature, err := k.CreateStorageKey(goCtx, &param)
	if err != nil {
		//it is safe to forward local errors
		return nil, err
	}

	signature, err = k.GetSignature(ctx, storageKeySignature.StorageKey)
	if err != nil {
		// it is safe to forward local errors
		return nil, err
	}

	// fetch reference payload link
	referencePayloadLink, err := k.GetPayloadLink(ctx, req.ReferenceId)
	if err != nil {
		// it is safe to forward local errors
		return nil, err
	}

	// reconstruct signature payload
	signaturePayload := util.CalculateHash(util.HashConcat(targetAccAddress, referenceId, referencePayloadLink))

	k.isValidSignature(goCtx, targetAccAddress, signaturePayload, signature.Signature, signature.Algorithm, signature.Certificate)

	return &types.QueryVerifySignatureResponse{}, nil
}

func (k Keeper) isValidSignature(goCtx context.Context, targetAccAddress, signaturePayload, signature, signatureAlgorithm, certificate string) error {
	return nil
}
