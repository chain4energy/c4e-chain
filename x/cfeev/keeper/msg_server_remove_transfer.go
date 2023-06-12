package keeper

import (
	"context"
	"cosmossdk.io/errors"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveTransfer(goCtx context.Context, msg *types.MsgRemoveTransfer) (*types.MsgRemoveTransferResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "remove energy transfer")

	ctx := sdk.UnwrapSDKContext(goCtx)

	transfer, found := k.GetEnergyTransfer(ctx, msg.GetId())
	if !found {
		return nil, errors.Wrap(types.ErrEnergyTransferCannotBeRemoved, "energy transfer not found")
	}

	if transfer.Status == types.TransferStatus_PAID || transfer.Status == types.TransferStatus_CANCELLED {
		k.RemoveEnergyTransfer(ctx, msg.GetId())
	} else {
		return nil, errors.Wrap(types.ErrWrongEnergyTransferStatus, "energy transfer status is not PAID or CANCELLED")
	}

	return &types.MsgRemoveTransferResponse{}, nil
}
