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
	keeper := k.Keeper

	id, err := keeper.PostEnergyTransferOffer(ctx,
		msg.Creator,
		msg.ChargerId,
		types.ChargerStatus_ACTIVE,
		*msg.GetLocation(),
		msg.GetTariff(),
		msg.GetName(),
		msg.GetPlugType(),
	)
	if err != nil {
		k.Logger(ctx).Error("publish energy transfer offer failed", "error", err)
		return nil, err
	}

	// place offer ID into response
	return &types.MsgPublishEnergyTransferOfferResponse{Id: id}, nil
}
