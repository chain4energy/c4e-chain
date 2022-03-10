package keeper

import (
	"context"
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k msgServer) BeginRedelegate(goCtx context.Context, msg *types.MsgBeginRedelegate) (*types.MsgBeginRedelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	keeper := k.Keeper

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

	// TODO !!!! getting all reward coins, not only bond denom
	// balanceBefore := k.bank.GetBalance(ctx, delegableAddress, keeper.staking.BondDenom(ctx))

	delagateMsg := stakingtypes.MsgBeginRedelegate{DelegatorAddress: accVestings.DelegableAddress,
		ValidatorSrcAddress: msg.ValidatorSrcAddress, ValidatorDstAddress: msg.ValidatorDstAddress, Amount: msg.Amount}
	resp, err := k.stakingMsgServer.BeginRedelegate(goCtx, &delagateMsg)
	if err != nil {
		return nil, err
	}

	// balanceAfter := k.bank.GetBalance(ctx, delegableAddress, keeper.staking.BondDenom(ctx))
	// coinToSend := balanceAfter.Sub(balanceBefore)

	// keeper.SetAccountVestings(ctx, accVestings)

	// k.bank.SendCoins(ctx, delegableAddress, accountAddress, sdk.NewCoins(coinToSend))
	// _ = coinToSend
	// _ = accountAddress
	return &types.MsgBeginRedelegateResponse{CompletionTime: resp.CompletionTime}, nil
}
