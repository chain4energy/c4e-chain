package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveTransfer(goCtx context.Context, msg *types.MsgRemoveTransfer) (*types.MsgRemoveTransferResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "remove energy transfer")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.RemoveTransfer(ctx, msg.GetId()); err != nil {
		k.Logger(ctx).Debug("remove energy transfer error", "error", err)
		return nil, err
	}

	return &types.MsgRemoveTransferResponse{}, nil
}
