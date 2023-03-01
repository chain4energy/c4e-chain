package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateMinters(goCtx context.Context, msg *types.MsgUpdateMinters) (*types.MsgUpdateMintersResponse, error) {
	if k.authority != msg.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	params.MinterConfig.Minters = msg.Minters
	if err := k.SetParams(ctx, params); err != nil {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}

	return &types.MsgUpdateMintersResponse{}, nil
}

func (k msgServer) UpdateMintDenom(goCtx context.Context, msg *types.MsgUpdateMintDenom) (*types.MsgUpdateMintDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	params.MintDenom = msg.MintDenom
	if err := k.SetParams(ctx, params); err != nil {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)

	}

	return &types.MsgUpdateMintDenomResponse{}, nil
}
