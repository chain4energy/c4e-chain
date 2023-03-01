package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAllSubDistributors(goCtx context.Context, msg *types.MsgUpdateAllSubDistributors) (*types.MsgUpdateAllSubDistributorsResponse, error) {
	if k.authority != msg.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, types.Params{SubDistributors: msg.SubDistributors}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateAllSubDistributorsResponse{}, nil
}

func (k msgServer) UpdateSingleSubdistributor(goCtx context.Context, distributor *types.MsgUpdateSingleSubdistributor) (*types.MsgUpdateSingleSubdistributorResponse, error) {
	if k.authority != distributor.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, distributor.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	subDistributors := k.Keeper.GetParams(ctx).SubDistributors
	for i, subDistributor := range subDistributors {
		if subDistributor.Name == distributor.SubDistributor.Name {
			subDistributors[i] = *distributor.SubDistributor
			if err := k.SetParams(ctx, types.Params{SubDistributors: subDistributors}); err != nil {
				return nil, sdkerrors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
			}
			return &types.MsgUpdateSingleSubdistributorResponse{}, nil
		}
	}

	return nil, sdkerrors.Wrapf(govtypes.ErrInvalidProposalContent, "distributor not found")
}
