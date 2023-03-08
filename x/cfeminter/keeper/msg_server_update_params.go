package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/telemetry"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateMintersParams(goCtx context.Context, msg *types.MsgUpdateMintersParams) (*types.MsgUpdateMintersParamsResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update minters")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	params := k.GetParams(ctx)
	params.StartTime = msg.StartTime
	params.Minters = msg.Minters

	if err := k.SetParams(ctx, params); err != nil {
		return nil, errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}

	return &types.MsgUpdateMintersParamsResponse{}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update mint denom")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	var params types.Params
	params.MintDenom = msg.MintDenom
	params.StartTime = msg.StartTime
	params.Minters = msg.Minters

	if err := k.SetParams(ctx, params); err != nil {
		return nil, errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
