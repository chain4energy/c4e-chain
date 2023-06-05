package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnergyTransferStartedRequest(goCtx context.Context, msg *types.MsgEnergyTransferStartedRequest) (*types.MsgEnergyTransferStartedRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	err := keeper.EnergyTransferStartedRequest(ctx, msg.GetEnergyTransferId())
	if err != nil {
		k.Logger(ctx).Error("energy transfer started failed", "error", err)
		return nil, err
	}

	return &types.MsgEnergyTransferStartedRequestResponse{}, nil
}
