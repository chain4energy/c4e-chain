package keeper

import (
	"context"
	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveEnergyOffer(goCtx context.Context, msg *types.MsgRemoveEnergyOffer) (*types.MsgRemoveEnergyOfferResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "remove energy transfer offer")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.RemoveEnergyOffer(ctx, msg.Creator, msg.Id); err != nil {
		k.Logger(ctx).Error("remove energy offer error", "error", err)
		return nil, err
	}

	return &types.MsgRemoveEnergyOfferResponse{}, nil
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
