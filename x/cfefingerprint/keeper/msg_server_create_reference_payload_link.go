package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) CreateReferencePayloadLink(goCtx context.Context, msg *types.MsgCreateReferencePayloadLink) (*types.MsgCreateReferencePayloadLinkResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create reference payloadLink")
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.Logger().Debug("create payload link for a given payload: ", msg.PayloadHash)

	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	referenceId, err := k.CreatePayloadLink(ctx, msg.PayloadHash)
	if err != nil {
		k.Logger(ctx).Error("move available vesting by denoms - validation error", "error", err)
		return nil, err
	}

	return &types.MsgCreateReferencePayloadLinkResponse{ReferenceId: referenceId}, nil
}
