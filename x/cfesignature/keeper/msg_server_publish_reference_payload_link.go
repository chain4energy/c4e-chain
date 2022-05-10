package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PublishReferencePayloadLink(goCtx context.Context, msg *types.MsgPublishReferencePayloadLink) (*types.MsgPublishReferencePayloadLinkResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	var err error
	
	// Check if a Payload Link was already stored at the given key
	if !(k.checkIfPayloadLinkExists(ctx, msg.Key)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data was found at the given key, cannot overwrite present payloadlinks")
	}

	// store payload link
	err = k.AppendPayloadLink(ctx, msg.Key, msg.Value)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to store payload link")
	}

	timestampString := ctx.BlockTime().String()

	return &types.MsgPublishReferencePayloadLinkResponse{TxTimestamp: timestampString}, nil
	// return &types.MsgPublishReferencePayloadLinkResponse{TxTimestamp: ""}, nil
}
