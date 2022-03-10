package keeper

import (
	"context"
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

)

func (k msgServer) WithdrawDelegatorReward(goCtx context.Context, msg *types.MsgWithdrawDelegatorReward) (*types.MsgWithdrawDelegatorRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper
	keeper.Logger(ctx).Info("WithdrawDelegatorReward: " + msg.DelegatorAddress)

	// stakingKeeper := keeper.staking

	accVestings, found := keeper.GetAccountVestings(ctx, msg.DelegatorAddress)
	if !found {
		return nil, fmt.Errorf("No vestings for account: %q", msg.DelegatorAddress)
	}
	if len(accVestings.DelegableAddress) == 0 {
		return nil, fmt.Errorf("No delegable vestings for account: %q", msg.DelegatorAddress)
	}

	// delegableAddress, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
	// if err != nil {
	// 	return nil, err
	// }

	// accountAddress, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	// if err != nil {
	// 	return nil, err
	// }

	withdrawMsg := distrtypes.MsgWithdrawDelegatorReward{DelegatorAddress: accVestings.DelegableAddress,
		ValidatorAddress: msg.ValidatorAddress}
	_, err := k.distrMsgServer.WithdrawDelegatorReward(goCtx, &withdrawMsg)
	if err != nil {
		return nil, err
	}

	// TODO: Handling the message
	_ = ctx

	return &types.MsgWithdrawDelegatorRewardResponse{}, nil
}
