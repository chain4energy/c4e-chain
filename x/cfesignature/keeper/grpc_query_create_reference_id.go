package keeper

import (
	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
	"crypto/rand"
	"encoding/hex"
)

func (k Keeper) CreateReferenceId(goCtx context.Context, req *types.QueryCreateReferenceIdRequest) (*types.QueryCreateReferenceIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: verify that the golang crypto lib returns random numbers that are good enough to be used here!
	rand32 := make([]byte, 32)
	_, err := rand.Read(rand32)
	if err != nil {
		// return "", errorcode.Internal.WithMessage("failed to generate referenceID, %v", err).LogReturn()
		return nil, err
	}

	// k.CreateStorageKey()

	// encode random numbers to hex string
	referenceID := hex.EncodeToString(rand32)

	// TODO: Process the query
	_ = ctx

	return &types.QueryCreateReferenceIdResponse{ReferenceId: referenceID}, nil
}
