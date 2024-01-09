package keeper

import (
	"cosmossdk.io/errors"
	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"crypto/rand"
	"encoding/hex"
)

// CreateReferenceID creates a referenceID and verifies that is has not been used yet
func (k Keeper) CreateReferenceId(goCtx context.Context, req *types.QueryCreateReferenceIdRequest) (*types.QueryCreateReferenceIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: verify that the golang crypto lib returns random numbers that are good enough to be used here!
	rand32 := make([]byte, 32)
	_, err := rand.Read(rand32)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrLogic, "failed to generate referenceID")
	}

	// encode random numbers to hex string
	referenceID := hex.EncodeToString(rand32)

	// make sure that there is no such referenceID for this account address on the ledger yet:
	targetAccAddress := req.Creator
	createStorageKeyRequest := &types.QueryCreateStorageKeyRequest{TargetAccAddress: targetAccAddress, ReferenceId: referenceID}
	storageKey, err := k.CreateStorageKey(goCtx, createStorageKeyRequest)
	if err != nil {
		// it is safe to return local errors
		return nil, err
	}

	var err1 error

	data, err1 := k.GetSignature(ctx, storageKey.StorageKey)
	if err != nil {
		if !errors.IsOf(err1, sdkerrors.ErrKeyNotFound) {
			return nil, err1
		}
	}

	if data != nil {
		return nil, errors.Wrap(sdkerrors.ErrLogic, "KVStore data for this referenceID already exists.")
	}

	return &types.QueryCreateReferenceIdResponse{ReferenceId: referenceID}, nil
}
