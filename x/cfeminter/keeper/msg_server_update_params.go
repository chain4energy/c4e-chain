package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/telemetry"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateMinters(goCtx context.Context, msg *types.MsgUpdateMinters) (*types.MsgUpdateMintersResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update minters")

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	params.Minters = msg.Minters
	if err := k.SetParams(ctx, params); err != nil {
		return nil, errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}

	return &types.MsgUpdateMintersResponse{}, nil
}

func (k msgServer) UpdateMintDenom(goCtx context.Context, msg *types.MsgUpdateMintDenom) (*types.MsgUpdateMintDenomResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update mint denom")

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	params.MintDenom = msg.MintDenom
	if err := k.SetParams(ctx, params); err != nil {
		return nil, errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)

	}

	return &types.MsgUpdateMintDenomResponse{}, nil
}

func (k msgServer) UpdateStartTime(goCtx context.Context, msg *types.MsgUpdateStartTime) (*types.MsgUpdateStartTimeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	params.StartTime = msg.StartTime
	if err := k.SetParams(ctx, params); err != nil {
		return nil, errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)

	}

	return &types.MsgUpdateStartTimeResponse{}, nil
}
