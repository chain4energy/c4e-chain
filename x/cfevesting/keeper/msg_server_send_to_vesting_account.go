package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendToVestingAccount(goCtx context.Context, msg *types.MsgSendToVestingAccount) (*types.MsgSendToVestingAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	_, err := keeper.SendToNewVestingAccount(ctx, msg.FromAddress, msg.ToAddress, msg.VestingId, msg.Amount, msg.RestartVesting)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendToVestingAccountResponse{}, nil
}
