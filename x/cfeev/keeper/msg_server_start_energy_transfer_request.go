package keeper

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) StartEnergyTransferRequest(goCtx context.Context, msg *types.MsgStartEnergyTransferRequest) (*types.MsgStartEnergyTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	// handling tariffFromMsg string from message
	tariffFromMsg, err := strconv.ParseInt(msg.GetOfferedTariff(), 10, 32)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request, incorrectly defined tariff")
	}
	offeredTariff := int32(tariffFromMsg)
	dynamicCollateral := msg.GetCollateral()

	energyTransferId, err := keeper.StartEnergyTransferRequest(ctx,
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

	return &types.MsgStartEnergyTransferRequestResponse{Id: energyTransferId}, nil
}
