package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/telemetry"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update vesting denom")

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	poolList := k.GetAllAccountVestingPools(ctx)

	if len(poolList) > 0 {
		return nil, errors.Wrapf(govtypes.ErrInvalidProposalMsg, "Pool exist cannot change denom")

	}

	if err := k.SetParams(ctx, types.Params{Denom: msg.Denom}); err != nil {
		return nil, errors.Wrap(govtypes.ErrInvalidProposalMsg, err.Error())
	}

	return &types.MsgUpdateDenomResponse{}, nil
}
