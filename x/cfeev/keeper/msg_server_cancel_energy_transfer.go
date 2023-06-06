package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelEnergyTransferRequest(goCtx context.Context, msg *types.MsgCancelEnergyTransferRequest) (*types.MsgCancelEnergyTransferRequestResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "cancel energy transfer")
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Debug("CancelEnergyTransferRequest - reason=%s, code=%s", msg.GetErrorInfo(), msg.GetErrorCode())
	keeper := k.Keeper

	err := keeper.CancelEnergyTransferRequest(ctx, msg.GetEnergyTransferId())
	if err != nil {
		k.Logger(ctx).Error("cancel energy transfer failed", "error", err)
		return nil, err
	}

	return &types.MsgCancelEnergyTransferRequestResponse{}, nil
}
