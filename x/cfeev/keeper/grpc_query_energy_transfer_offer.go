package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllEnergyTransferOffers(c context.Context, req *types.QueryAllEnergyTransferOffersRequest) (*types.QueryAllEnergyTransferOfferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var energyTransferOffers []types.EnergyTransferOffer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	energyTransferOfferStore := prefix.NewStore(store, types.EnergyTransferOfferKey)

	pageRes, err := query.Paginate(energyTransferOfferStore, req.Pagination, func(key []byte, value []byte) error {
		var energyTransferOffer types.EnergyTransferOffer
		if err := k.cdc.Unmarshal(value, &energyTransferOffer); err != nil {
			return err
		}

		energyTransferOffers = append(energyTransferOffers, energyTransferOffer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEnergyTransferOfferResponse{EnergyTransferOffer: energyTransferOffers, Pagination: pageRes}, nil
}

func (k Keeper) EnergyTransferOffer(c context.Context, req *types.QueryEnergyTransferOfferRequest) (*types.QueryEnergyTransferOfferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	energyTransferOffer, found := k.GetEnergyTransferOffer(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryEnergyTransferOfferResponse{EnergyTransferOffer: energyTransferOffer}, nil
}

func (k Keeper) EnergyTransferOffers(goCtx context.Context, req *types.QueryEnergyTransferOffersRequest) (*types.QueryEnergyTransferOffersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	_, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "wrong owner address (%s)", err.Error())
	}

	var energyTransferOffers []types.EnergyTransferOffer
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	energyTransferOfferStore := prefix.NewStore(store, types.EnergyTransferOfferKey)

	pageRes, err := query.Paginate(energyTransferOfferStore, req.Pagination, func(key []byte, value []byte) error {
		var energyTransferOffer types.EnergyTransferOffer
		if err := k.cdc.Unmarshal(value, &energyTransferOffer); err != nil {
			return err
		}

		if energyTransferOffer.GetOwner() == req.GetOwner() {
			energyTransferOffers = append(energyTransferOffers, energyTransferOffer)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryEnergyTransferOffersResponse{EnergyTransferOffers: energyTransferOffers, Pagination: pageRes}, nil
}
