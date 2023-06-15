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

	if err := k.Keeper.RemoveTransfer(ctx, msg.GetId()); err != nil {
		k.Logger(ctx).Error("remove energy transfer error", "error", err)
		return nil, err
	}

	return &types.MsgRemoveTransferResponse{}, nil
}

func (k Keeper) RemoveTransfer(ctx sdk.Context, id uint64) error {
	transfer, err := k.MustGetEnergyTransfer(ctx, id)
	if err != nil {
		return err
	}

	if transfer.Status != types.TransferStatus_PAID && transfer.Status != types.TransferStatus_CANCELLED {
		return errors.Wrap(types.ErrWrongEnergyTransferStatus, "energy transfer status is not PAID or CANCELLED")
	}

	k.RemoveEnergyTransfer(ctx, id)
	return nil
}
