package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/telemetry"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAllSubDistributorsParams(goCtx context.Context, msg *types.MsgUpdateAllSubDistributorsParams) (*types.MsgUpdateAllSubDistributorsParamsResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update all sub distributors")

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, types.Params{SubDistributors: msg.SubDistributors}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateAllSubDistributorsParamsResponse{}, nil
}

func (k msgServer) UpdateSubDistributorParam(goCtx context.Context, distributor *types.MsgUpdateSubDistributorParam) (*types.MsgUpdateSubDistributorParamResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update sub distributor")

	if k.authority != distributor.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, distributor.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	subDistributors := k.Keeper.GetParams(ctx).SubDistributors
	for i, subDistributor := range subDistributors {
		if subDistributor.Name == distributor.SubDistributor.Name {
			subDistributors[i] = *distributor.SubDistributor
			if err := k.SetParams(ctx, types.Params{SubDistributors: subDistributors}); err != nil {
				return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "validation error: %s", err)
			}
			return &types.MsgUpdateSubDistributorParamResponse{}, nil
		}
	}

	return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "distributor not found")
}

func (k msgServer) UpdateSubDistributorDestinationShareParam(goCtx context.Context, msg *types.MsgUpdateSubDistributorDestinationShareParam) (*types.MsgUpdateSubDistributorDestinationShareParamResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update sub distributor destination share")

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	subDistributors := k.Keeper.GetParams(ctx).SubDistributors
	for i, subDistributor := range subDistributors {
		for y, share := range subDistributor.Destinations.Shares {
			if share.Name == msg.DestinationName {
				subDistributors[i].Destinations.Shares[y].Share = msg.Share
				if err := k.SetParams(ctx, types.Params{SubDistributors: subDistributors}); err != nil {
					return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "validation error: %s", err)
				}
				return &types.MsgUpdateSubDistributorDestinationShareParamResponse{}, nil
			}
		}
	}
	return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "distributor not found")
}

func (k msgServer) UpdateMsgUpdateSubDistributorBurnShareParam(goCtx context.Context, msg *types.MsgUpdateSubDistributorBurnShareParam) (*types.MsgUpdateSubDistributorBurnShareParamResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update sub distributor burn share")

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	subDistributors := k.Keeper.GetParams(ctx).SubDistributors
	for i, subDistributor := range subDistributors {
		if subDistributor.Name == msg.SubDistributorName {
			subDistributors[i].Destinations.BurnShare = msg.BurnShare
			if err := k.SetParams(ctx, types.Params{SubDistributors: subDistributors}); err != nil {
				return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "validation error: %s", err)
			}
			return &types.MsgUpdateSubDistributorBurnShareParamResponse{}, nil
		}
	}
	return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "distributor not found")

}
