package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/chain4energy/c4e-chain/x/cfesignature/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CreateReferencePayloadLink(goCtx context.Context, req *types.QueryCreateReferencePayloadLinkRequest) (*types.QueryCreateReferencePayloadLinkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	if len(req.ReferenceId) != 64 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid input size of referenceID")
	}

	referenceKey := util.CalculateHash(req.ReferenceId)
	referenceValue := util.CalculateHash(util.HashConcat(req.ReferenceId, req.PayloadHash))

	ctx.Logger().Debug("referenceKey   = %s", referenceKey)

	return &types.QueryCreateReferencePayloadLinkResponse{ReferenceKey: referenceKey, ReferenceValue: referenceValue}, nil
}
