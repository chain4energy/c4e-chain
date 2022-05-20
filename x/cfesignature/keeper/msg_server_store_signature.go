package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/chain4energy/c4e-chain/x/cfesignature/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) StoreSignature(goCtx context.Context, msg *types.MsgStoreSignature) (*types.MsgStoreSignatureResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	var signatureObject types.Signature
	var signatureJSON = msg.SignatureJSON
	var err error

	txHash := sha256.Sum256(ctx.TxBytes())
	txId := hex.EncodeToString(txHash[:])

	// try to extract all values from the given JSON
	// .signature
	signatureObject.Signature, err = util.ExtractFieldFromJSON(signatureJSON, "signature")
	if err != nil {
		// it is safe to forward local errors
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to parse signature: "+signatureJSON)
	}
	// .algorithm
	signatureObject.Algorithm, err = util.ExtractFieldFromJSON(signatureJSON, "algorithm")
	if err != nil {
		// it is safe to forward local errors
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to parse algorithm")
	}
	// .certificate
	signatureObject.Certificate, err = util.ExtractFieldFromJSON(signatureJSON, "certificate")
	if err != nil {
		// it is safe to forward local errors
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to parse certificate")
	}

	// .timestamp
	signatureObject.Timestamp = ctx.BlockTime().String()
	timestamp := k.AppendSignature(ctx, msg.StorageKey, signatureObject)

	// TODO: extract and verify user cert
	// TODO: if support for multiple signatures is added then TODO: Check if the certificate was used for signing before

	return &types.MsgStoreSignatureResponse{TxId: txId, TxTimestamp: timestamp}, nil
}
