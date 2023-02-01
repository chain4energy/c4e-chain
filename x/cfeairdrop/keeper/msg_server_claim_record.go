package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddClaimRecords(goCtx context.Context, msg *types.MsgAddClaimRecords) (*types.MsgAddClaimRecordsResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "add airdrop entries message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper
	if err := keeper.AddUsersEntries(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.ClaimRecords,
	); err != nil {
		return nil, err
	}

	return &types.MsgAddClaimRecordsResponse{}, nil
}

func (k msgServer) DeleteClaimRecord(goCtx context.Context, msg *types.MsgDeleteClaimRecord) (*types.MsgDeleteClaimRecordResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "delete airdrop entry message")
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

	if err := keeper.DeleteClaimRecord(
		ctx,
		msg.Owner,
		msg.CampaignId,
		msg.UserAddress,
	); err != nil {
		return nil, err
	}

	return &types.MsgDeleteClaimRecordResponse{}, nil
}
