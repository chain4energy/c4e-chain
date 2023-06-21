package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelEnergyTransfer(goCtx context.Context, msg *types.MsgCancelEnergyTransfer) (*types.MsgCancelEnergyTransferResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "cancel energy transfer")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.CancelEnergyTransfer(ctx, msg.GetEnergyTransferId()); err != nil {
		k.Logger(ctx).Debug("cancel energy transfer error", "error", err)
		return nil, err
	}

	return &types.MsgCancelEnergyTransferResponse{}, nil
}
