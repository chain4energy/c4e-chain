package keeper

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) StartEnergyTransfer(goCtx context.Context, msg *types.MsgStartEnergyTransfer) (*types.MsgStartEnergyTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	defer telemetry.IncrCounter(1, types.ModuleName, "start energy transfer")

	// handling tariffFromMsg string from message
	tariffFromMsg, err := strconv.ParseInt(msg.GetOfferedTariff(), 10, 32)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request, incorrectly defined tariff")
	}
	offeredTariff := int32(tariffFromMsg)
	dynamicCollateral := msg.GetCollateral()

	energyTransferId, err := keeper.StartEnergyTransfer(ctx,
		msg.OwnerAccountAddress,
		msg.Creator,
		msg.EnergyTransferOfferId,
		msg.ChargerId, offeredTariff,
		msg.GetEnergyToTransfer(),
		*dynamicCollateral)

	if err != nil {
		k.Logger(ctx).Error("start energy transfer failed", "error", err)
		return nil, err
	}

	return &types.MsgStartEnergyTransferResponse{Id: energyTransferId}, nil
}
