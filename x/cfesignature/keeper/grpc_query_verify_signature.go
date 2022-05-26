package keeper

import (
	"context"
	"encoding/base64"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/chain4energy/c4e-chain/x/cfesignature/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	var signature *types.Signature

	queryCreateStorageKeyRequest := types.QueryCreateStorageKeyRequest{TargetAccAddress: req.TargetAccAddress, ReferenceId: req.ReferenceId}

	// fetch storage keys for signature and document hash
	storageKeySignature, err := k.CreateStorageKey(goCtx, &queryCreateStorageKeyRequest)
	if err != nil {
		//it is safe to forward local errors
		return nil, err
	}

	// get signature object from the ledger
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

	validationError := k.isValidSignature(goCtx, targetAccAddress, signaturePayload, signature.Signature, signature.Algorithm, signature.Certificate)
	if validationError != nil {
		// it is safe to forward local errors
		return nil, validationError
	}

	return &types.QueryVerifySignatureResponse{Signature: signature.Signature, Algorithm: signature.Algorithm, Certificate: signature.Signature,
		Timestamp: signature.Timestamp, Valid: "valid"}, nil
}

func (k Keeper) isValidSignature(goCtx context.Context, targetAccAddress, signaturePayload, signature, signatureAlgorithm, certificate string) error {

	// decode signature from base64
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "failed to decode signature string")

	}

	x509signatureAlgorithm, err := util.GetSignatureAlgorithmFromString(signatureAlgorithm)
	if err != nil {
		// it is safe to forward local errors
		return err
	}

	userCert, err := util.GetUserCertificateFromString([]byte(certificate))
	if err != nil {
		// it is safe to forward local errors
		return err
	}

	// verifies that signature is a valid signature
	if err = userCert.CheckSignature(x509signatureAlgorithm, []byte(signaturePayload), signatureBytes); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "signature validation failed")
	}

	return nil

}
