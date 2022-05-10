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
	_ = txHash

	// try to extract all values from the given JSON
	// .signature
	signatureObject.Signature, err = util.ExtractFieldFromJSON(signatureJSON, "signature")
	if err != nil {
		// it is safe to forward local errors
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to parse signature"+signatureJSON)
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

	txId := hex.EncodeToString(txHash[:])

	return &types.MsgStoreSignatureResponse{TxId: txId, TxTimestamp: timestamp}, nil
}
