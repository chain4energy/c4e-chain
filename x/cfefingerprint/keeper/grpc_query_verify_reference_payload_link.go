package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifyReferencePayloadLink(goCtx context.Context, req *types.QueryVerifyReferencePayloadLinkRequest) (*types.QueryVerifyReferencePayloadLinkResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// fetch data published on ledger
	ledgerPayloadLinkValue, err := k.GetPayloadLink(ctx, req.ReferenceId)
	if err != nil {
		return &types.QueryVerifyReferencePayloadLinkResponse{IsValid: "false"}, err
	}

	// calculate expeced data based on payload hash
	// so called reference value
	expectedPayloadLinkValue := util.CalculateHash(util.HashConcat(req.ReferenceId, req.PayloadHash))

	// verify ledger matches the payloadhash
	if expectedPayloadLinkValue == ledgerPayloadLinkValue {
		return &types.QueryVerifyReferencePayloadLinkResponse{IsValid: "true"}, nil
	}

	// something failed
	return &types.QueryVerifyReferencePayloadLinkResponse{IsValid: "false"}, nil
}
