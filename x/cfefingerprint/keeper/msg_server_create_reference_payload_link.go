package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateReferencePayloadLink(goCtx context.Context, msg *types.MsgCreateReferencePayloadLink) (*types.MsgCreateReferencePayloadLinkResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create reference payloadLink")
	ctx := sdk.UnwrapSDKContext(goCtx)

	referenceId, err := k.CreatePayloadLink(ctx, msg.PayloadHash)
	if err != nil {
		k.Logger(ctx).Error("create reference payload link error", "error", err)
		return nil, err
	}

	return &types.MsgCreateReferencePayloadLinkResponse{ReferenceId: *referenceId}, nil
}
