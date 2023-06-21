package keeper

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) PublishEnergyTransferOffer(
	ctx sdk.Context,
	creator string,
	chargerId string,
	location types.Location,
	tariff uint64,
	name string,
	plugType types.PlugType,
) (*uint64, error) {
	k.Logger(ctx).Debug("publish energy transfer offer", "creator", creator, "chargerId", chargerId, "location",
		location, "tariff", tariff, "name", name, "plugType", plugType)

	if err := types.ValidatePublishEnergyTransferOffer(chargerId, name, location); err != nil {
		return nil, err
	}

	// there is a 1-1 relation between the offer and the charger
	// check if another offer for this chargerId has been added
	_, found := k.GetTransferOfferByChargerId(ctx, chargerId)
	if found {
		// Rule: either log the error or throw it but never do both
		// Rule: pass all relevant information to errors to make them informative as much as possible
		return nil, errors.Wrapf(c4eerrors.ErrAlreadyExists, "energy transfer offer for this charger %s/%s already exists", name, chargerId)
	}

	var energyTransferOffer = types.EnergyTransferOffer{
		Owner:         creator,
		ChargerId:     chargerId,
		ChargerStatus: types.ChargerStatus_ACTIVE,
		Location:      &location,
		Tariff:        tariff,
		Name:          name,
		PlugType:      plugType,
	}

	id := k.AppendEnergyTransferOffer(ctx, energyTransferOffer)

	event := &types.EventPublishOffer{
		EnergyTransferOfferId: id,
		Owner:                 creator,
		ChargerId:             chargerId,
		Tariff:                tariff,
		Name:                  name,
		PlugType:              plugType,
	}
	if err := ctx.EventManager().EmitTypedEvent(event); err != nil {
		k.Logger(ctx).Error("new publish energy transfer offer emit event error", "event", event, "error", err.Error())
	}
	k.Logger(ctx).Debug("new publish energy transfer ret", "id", id)
	return &id, nil
}

func (k Keeper) RemoveEnergyOffer(ctx sdk.Context, creator string, energyOfferId uint64) error {
	offer, err := k.MustGetEnergyTransferOffer(ctx, energyOfferId)
	if err != nil {
		return err
	}

	if offer.Owner != creator {
		return errors.Wrapf(sdkerrors.ErrorInvalidSigner, "address %s is not a creator of energy offer with id %d", creator, offer.Id)
	}

	k.RemoveEnergyTransferOffer(ctx, energyOfferId)
	return nil
}
