package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnergyTransferCompletedRequest(goCtx context.Context, msg *types.MsgEnergyTransferCompletedRequest) (*types.MsgEnergyTransferCompletedRequestResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "energy transfer completed")

	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	err := keeper.EnergyTransferCompletedRequest(ctx, msg.EnergyTransferId, msg.GetUsedServiceUnits())
	if err != nil {
		k.Logger(ctx).Error("complete energy transfer failed", "error", err)
		return nil, err
	}

	// TODO: Handling the response
	return &types.MsgEnergyTransferCompletedRequestResponse{}, nil
}
