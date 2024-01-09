package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PublishReferencePayloadLink(goCtx context.Context, msg *types.MsgPublishReferencePayloadLink) (*types.MsgPublishReferencePayloadLinkResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "publish reference payload link message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	var err error

	// Check if a Payload Link was already stored at the given key
	if !(k.checkIfPayloadLinkExists(ctx, msg.Key)) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "data was found at the given key, cannot overwrite present payloadlinks")
	}

	// store payload link
	err = k.AppendPayloadLink(ctx, msg.Key, msg.Value)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "failed to store payload link")
	}

	timestampString := ctx.BlockTime().String()

	return &types.MsgPublishReferencePayloadLinkResponse{TxTimestamp: timestampString}, nil
}
