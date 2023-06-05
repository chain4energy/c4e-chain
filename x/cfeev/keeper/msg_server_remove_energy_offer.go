package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveEnergyOffer(goCtx context.Context, msg *types.MsgRemoveEnergyOffer) (*types.MsgRemoveEnergyOfferResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "remove energy transfer offer")

	ctx := sdk.UnwrapSDKContext(goCtx)

	offer, found := k.GetEnergyTransferOffer(ctx, msg.GetId())
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEnergyTransferOfferCannotBeRemoved, "energy transfer offer not found")
	}

	// offer can be removed only by offer owner
	if !(offer.GetOwner() == msg.GetCreator()) {
		return nil, sdkerrors.Wrap(types.ErrEnergyTransferOfferCannotBeRemoved, "Message sender is not an owner of the offer")
	}

	k.RemoveEnergyTransferOffer(ctx, msg.GetId())

	return &types.MsgRemoveEnergyOfferResponse{}, nil
}
