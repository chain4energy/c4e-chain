package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PublishEnergyTransferOffer(goCtx context.Context, msg *types.MsgPublishEnergyTransferOffer) (*types.MsgPublishEnergyTransferOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// there is a 1-1 relation between the offer and the charger
	// check if another offer for this chargerId has been added
	_, found := k.GetTransferOfferByChargerId(ctx, msg.ChargerId)
	if found {
		return nil, sdkerrors.Wrap(types.ErrOfferForChargerAlreadyExists, "energy transfer offer for this charger already exists")
	}

	var energyTransferOffer = types.EnergyTransferOffer{
		Owner:         msg.Creator,
		ChargerId:     msg.ChargerId,
		ChargerStatus: types.ChargerStatus_ACTIVE,
		Location:      msg.GetLocation(),
		Tariff:        msg.GetTariff(),
		Name:          msg.GetName(),
		PlugType:      msg.GetPlugType(),
	}

	id := k.AppendEnergyTransferOffer(ctx, energyTransferOffer)

	// place offer ID into response
	return &types.MsgPublishEnergyTransferOfferResponse{Id: id}, nil
}
