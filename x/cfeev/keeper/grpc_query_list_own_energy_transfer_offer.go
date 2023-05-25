package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListOwnEnergyTransferOffer(goCtx context.Context, req *types.QueryListOwnEnergyTransferOfferRequest) (*types.QueryListOwnEnergyTransferOfferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.OwnerAccAddress == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request - CP owner not specified")
	}

	var energyTransferOffers []types.EnergyTransferOffer
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	energyTransferOfferStore := prefix.NewStore(store, types.KeyPrefix(types.EnergyTransferOfferKey))

	pageRes, err := query.Paginate(energyTransferOfferStore, req.Pagination, func(key []byte, value []byte) error {
		var energyTransferOffer types.EnergyTransferOffer
		if err := k.cdc.Unmarshal(value, &energyTransferOffer); err != nil {
			return err
		}

		if energyTransferOffer.GetOwner() == req.GetOwnerAccAddress() {
			energyTransferOffers = append(energyTransferOffers, energyTransferOffer)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListOwnEnergyTransferOfferResponse{EnergyTransferOffer: energyTransferOffers, Pagination: pageRes}, nil
}
