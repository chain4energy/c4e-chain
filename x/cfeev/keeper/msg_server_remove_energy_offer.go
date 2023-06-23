package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveEnergyOffer(goCtx context.Context, msg *types.MsgRemoveEnergyOffer) (*types.MsgRemoveEnergyOfferResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "remove energy transfer offer")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.RemoveEnergyOffer(ctx, msg.GetOwner(), msg.Id); err != nil {
		k.Logger(ctx).Debug("remove energy offer error", "error", err)
		return nil, err
	}

	return &types.MsgRemoveEnergyOfferResponse{}, nil
}
