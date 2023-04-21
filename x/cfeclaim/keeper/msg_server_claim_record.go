package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddClaimRecords(goCtx context.Context, msg *types.MsgAddClaimRecords) (*types.MsgAddClaimRecordsResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "add claim entries message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.AddUsersEntries(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.ClaimRecords,
	); err != nil {
		k.Logger(ctx).Debug("add user entries", "err", err.Error())
		return nil, err
	}

	return &types.MsgAddClaimRecordsResponse{}, nil
}

func (k msgServer) DeleteClaimRecord(goCtx context.Context, msg *types.MsgDeleteClaimRecord) (*types.MsgDeleteClaimRecordResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "delete claim entry message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.DeleteClaimRecord(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.UserAddress,
		msg.DeleteClaimRecordAction,
	); err != nil {
		k.Logger(ctx).Debug("delete claim record", "err", err.Error())
		return nil, err
	}

	return &types.MsgDeleteClaimRecordResponse{}, nil
}
