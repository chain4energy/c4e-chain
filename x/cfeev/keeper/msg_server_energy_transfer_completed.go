package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnergyTransferCompleted(goCtx context.Context, msg *types.MsgEnergyTransferCompleted) (*types.MsgEnergyTransferCompletedResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "energy transfer completed")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.EnergyTransferCompleted(ctx, msg.EnergyTransferId, msg.GetUsedServiceUnits()); err != nil {
		k.Logger(ctx).Debug("complete energy transfer error", "error", err)
		return nil, err
	}

	return &types.MsgEnergyTransferCompletedResponse{}, nil
}
