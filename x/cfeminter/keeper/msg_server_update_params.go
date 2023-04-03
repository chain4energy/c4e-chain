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
	defer telemetry.IncrCounter(1, types.ModuleName, "Update minters params")
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	params.StartTime = msg.StartTime
	params.Minters = msg.Minters

	if err := k.Keeper.UpdateParams(ctx, msg.Authority, params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateMintersParamsResponse{}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update params")
	ctx := sdk.UnwrapSDKContext(goCtx)

	var params types.Params
	params.MintDenom = msg.MintDenom
	params.StartTime = msg.StartTime
	params.Minters = msg.Minters

	if err := k.Keeper.UpdateParams(ctx, msg.Authority, params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (k Keeper) UpdateParams(ctx sdk.Context, authority string, params types.Params) error {
	if k.authority != authority {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, govtypes.ErrInvalidSigner.Error())
	}

	minterState := k.GetMinterState(ctx)
	if !params.ContainsMinter(minterState.SequenceId) {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "minter state sequence id %d not found in minters", minterState.SequenceId)
	}

	if err := k.SetParams(ctx, params); err != nil {
		return errors.Wrap(govtypes.ErrInvalidProposalContent, err.Error())
	}
	return nil
}
