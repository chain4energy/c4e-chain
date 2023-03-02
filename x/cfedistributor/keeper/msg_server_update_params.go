package keeper

import (
	"context"
	"cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAllSubDistributors(goCtx context.Context, msg *types.MsgUpdateAllSubDistributors) (*types.MsgUpdateAllSubDistributorsResponse, error) {
	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, types.Params{SubDistributors: msg.SubDistributors}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateAllSubDistributorsResponse{}, nil
}

func (k msgServer) UpdateSubDistributor(goCtx context.Context, distributor *types.MsgUpdateSubDistributor) (*types.MsgUpdateSubDistributorResponse, error) {
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
			return &types.MsgUpdateSubDistributorResponse{}, nil
		}
	}

	return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "distributor not found")
}

func (k msgServer) UpdateSubDistributorDestinationShare(goCtx context.Context, msg *types.MsgUpdateSubDistributorDestinationShare) (*types.MsgUpdateSubDistributorDestinationShareResponse, error) {
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
				return &types.MsgUpdateSubDistributorDestinationShareResponse{}, nil
			}
		}
	}
	return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "distributor not found")
}

func (k msgServer) UpdateMsgUpdateSubDistributorBurnShare(goCtx context.Context, msg *types.MsgUpdateSubDistributorBurnShare) (*types.MsgUpdateSubDistributorBurnShareResponse, error) {
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
			return &types.MsgUpdateSubDistributorBurnShareResponse{}, nil
		}
	}
	return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "distributor not found")

}
