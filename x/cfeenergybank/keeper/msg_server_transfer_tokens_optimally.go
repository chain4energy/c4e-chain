package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferTokensOptimally(goCtx context.Context, msg *types.MsgTransferTokensOptimally) (*types.MsgTransferTokensOptimallyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	eneryTokens := k.GetAllUserEnergyTokens(ctx, msg.AddressFrom, msg.TokenName)

	if len(eneryTokens) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "There is no this energyTokens")
	}
	if eneryTokens[0].UserAddress != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "This token doesn't belong to you.")
	}
	var accumulatedEnergy uint64 = 0
	for i := 0; i < len(eneryTokens); i++ {
		accumulatedEnergy += eneryTokens[i].Amount
	}
	if accumulatedEnergy < msg.Amount || msg.Amount <= 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong amount of tokens.")
	}

	amount := msg.Amount
	i := 0
	for amount > 0 {
		if eneryTokens[i].Amount <= amount {
			eneryTokens[i].UserAddress = msg.AddressTo
			k.SetEnergyToken(ctx, eneryTokens[i])
			amount -= eneryTokens[i].Amount
		} else {
			var newEnergyToken = types.EnergyToken{
				Name:        eneryTokens[i].Name,
				Amount:      amount,
				UserAddress: msg.AddressTo,
				CreatedAt:   eneryTokens[i].CreatedAt,
			}
			eneryTokens[i].Amount -= amount
			amount = 0
			k.AppendEnergyToken(ctx, newEnergyToken)
			k.SetEnergyToken(ctx, eneryTokens[i])
		}
		i++
	}
	var tokenHistory = types.TokensHistory{
		IssuerAddress: msg.Creator,
		UserAddress:   msg.Creator,
		CreatedAt:     uint64(time.Now().Unix()),
		Amount:        msg.Amount,
		TokenName:     eneryTokens[i].Name,
		TargetAddress: msg.AddressTo,
	}
	k.AppendTokensHistory(ctx, tokenHistory)

	return &types.MsgTransferTokensOptimallyResponse{}, nil
}
