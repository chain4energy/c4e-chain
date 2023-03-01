package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateDenomParam(goCtx context.Context, msg *types.MsgUpdateDenomParam) (*types.MsgUpdateDenomParamResponse, error) {
	if k.authority != msg.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	poolList := k.GetAllAccountVestingPools(ctx)

	if len(poolList) > 0 {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidProposalMsg, "Pool exist cannot change denom")

	}

	if err := k.SetParams(ctx, types.Params{Denom: msg.Denom}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateDenomParamResponse{}, nil
}
