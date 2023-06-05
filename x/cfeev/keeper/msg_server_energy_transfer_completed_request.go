package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnergyTransferCompletedRequest(goCtx context.Context, msg *types.MsgEnergyTransferCompletedRequest) (*types.MsgEnergyTransferCompletedRequestResponse, error) {
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
