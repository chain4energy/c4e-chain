package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CertificateTypeAll(goCtx context.Context, req *types.QueryAllCertificateTypeRequest) (*types.QueryAllCertificateTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var certificateTypes []types.CertificateType
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	certificateTypeStore := prefix.NewStore(store, types.KeyPrefix(types.CertificateTypeKey))

	pageRes, err := query.Paginate(certificateTypeStore, req.Pagination, func(key []byte, value []byte) error {
		var certificateType types.CertificateType
		if err := k.cdc.Unmarshal(value, &certificateType); err != nil {
			return err
		}

		certificateTypes = append(certificateTypes, certificateType)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCertificateTypeResponse{CertificateType: certificateTypes, Pagination: pageRes}, nil
}

func (k Keeper) CertificateType(goCtx context.Context, req *types.QueryGetCertificateTypeRequest) (*types.QueryGetCertificateTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	certificateType, found := k.GetCertificateType(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCertificateTypeResponse{CertificateType: certificateType}, nil
}
