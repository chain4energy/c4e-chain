package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PublishEnergyTransferOffer(goCtx context.Context, msg *types.MsgPublishEnergyTransferOffer) (*types.MsgPublishEnergyTransferOfferResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "publish energy transfer offer")
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, err := k.Keeper.PublishEnergyTransferOffer(ctx,
		msg.Creator,
		msg.ChargerId,
		*msg.GetLocation(),
		msg.GetTariff(),
		msg.GetName(),
		msg.GetPlugType(),
	)
	if err != nil {
		k.Logger(ctx).Debug("publish energy transfer offer error", "error", err)
		return nil, err
	}

	return &types.MsgPublishEnergyTransferOfferResponse{Id: *id}, nil
}
