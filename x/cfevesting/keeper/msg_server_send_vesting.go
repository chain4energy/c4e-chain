package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendVesting(goCtx context.Context, msg *types.MsgSendVesting) (*types.MsgSendVestingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendVestingResponse{}, nil
}
