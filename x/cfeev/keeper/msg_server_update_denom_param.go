package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/telemetry"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateDenomParam(goCtx context.Context, msg *types.MsgUpdateDenomParam) (*types.MsgUpdateDenomParamResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "Update vesting denom")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	if err := k.SetParams(ctx, types.Params{Denom: msg.Denom}); err != nil {
		return nil, errors.Wrap(govtypes.ErrInvalidProposalMsg, err.Error())
	}

	return &types.MsgUpdateDenomParamResponse{}, nil
}
