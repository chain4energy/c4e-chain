package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnergyTransferStarted(goCtx context.Context, msg *types.MsgEnergyTransferStarted) (*types.MsgEnergyTransferStartedResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "energy transfer started")
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	err := keeper.EnergyTransferStarted(ctx, msg.GetEnergyTransferId())
	if err != nil {
		k.Logger(ctx).Error("energy transfer started failed", "error", err)
		return nil, err
	}

	return &types.MsgEnergyTransferStartedResponse{}, nil
}
