package keeper

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/chain4energy/c4e-chain/x/cfesignature/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CreateStorageKey(goCtx context.Context, req *types.QueryCreateStorageKeyRequest) (*types.QueryCreateStorageKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// module specific logger
	logger := ctx.Logger()
	logger.Info("CreateStorageKey invoked")

	if len(req.ReferenceId) != 64 {
		errorMessage := "invalid input size of referenceID: " + strconv.Itoa(len(req.ReferenceId))
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, errorMessage)
	}

	if len(req.TargetAccAddress) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid input, targetAccAddress is empty")

	}

	hashInput := util.HashConcat(req.TargetAccAddress, req.ReferenceId)
	storageKey := util.CalculateHash(hashInput)

	return &types.QueryCreateStorageKeyResponse{StorageKey: storageKey}, nil
}
