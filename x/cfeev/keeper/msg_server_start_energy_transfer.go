package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StartEnergyTransfer(goCtx context.Context, msg *types.MsgStartEnergyTransfer) (*types.MsgStartEnergyTransferResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "start energy transfer")
	ctx := sdk.UnwrapSDKContext(goCtx)

	energyTransferId, err := k.Keeper.StartEnergyTransfer(ctx,
		msg.Creator,
		msg.EnergyTransferOfferId,
		msg.GetOfferedTariff(),
		msg.GetEnergyToTransfer(),
	)
	if err != nil {
		k.Logger(ctx).Debug("start energy transfer failed", "error", err)
		return nil, err
	}

	return &types.MsgStartEnergyTransferResponse{Id: *energyTransferId}, nil
}
