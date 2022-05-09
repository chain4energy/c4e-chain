package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PublishReferencePayloadLink(goCtx context.Context, msg *types.MsgPublishReferencePayloadLink) (*types.MsgPublishReferencePayloadLinkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPublishReferencePayloadLinkResponse{}, nil
}
