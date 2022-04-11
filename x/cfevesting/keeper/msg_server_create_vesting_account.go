package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateVestingAccount(goCtx context.Context, msg *types.MsgCreateVestingAccount) (*types.MsgCreateVestingAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	err := keeper.CreateVestingAccount(ctx, msg.FromAddress, msg.ToAddress, msg.Amount, msg.StartTime, msg.EndTime)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateVestingAccountResponse{}, nil
}
