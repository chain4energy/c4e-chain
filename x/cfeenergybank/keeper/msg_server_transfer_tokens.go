package keeper

import (
	"context"
	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func (k msgServer) TransferTokens(goCtx context.Context, msg *types.MsgTransferTokens) (*types.MsgTransferTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	eneryToken, isFound := k.GetEnergyToken(ctx, msg.TokenId)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "This token doesn't exists.")
	}
	if eneryToken.UserAddress != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "This token doesn't belong to you.")
	}
	if eneryToken.Amount < msg.Amount || msg.Amount <= 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong amount of tokens.")
	}
	if msg.Amount != eneryToken.Amount {
		var newEnergyToken = types.EnergyToken{
			Name:        eneryToken.Name,
			Amount:      msg.Amount,
			UserAddress: msg.AddressTo,
			CreatedAt:   eneryToken.CreatedAt,
		}
		k.AppendEnergyToken(ctx, newEnergyToken)
		eneryToken.Amount -= msg.Amount
		k.SetEnergyToken(ctx, eneryToken)
	} else {
		eneryToken.UserAddress = msg.AddressTo
		k.SetEnergyToken(ctx, eneryToken)
	}
	var tokenHistory = types.TokensHistory{
		IssuerAddress: msg.Creator,
		UserAddress:   msg.Creator,
		CreatedAt:     uint64(time.Now().Unix()),
		Amount:        msg.Amount,
		TokenName:     eneryToken.Name,
		TargetAddress: msg.AddressTo,
	}
	k.AppendTokensHistory(ctx, tokenHistory)
	_ = ctx

	return &types.MsgTransferTokensResponse{}, nil
}
