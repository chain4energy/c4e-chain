package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintToken(goCtx context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenParams, isFound := k.GetTokenParams(ctx, msg.Name)
	// TODO: accountAddress verification
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Token with this name doesn't exist.")
	}

	if tokenParams.MintAccount != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You aren't allowed to mint this token.")
	}

	var energyToken = types.EnergyToken{
		Name:        msg.Name,
		Amount:      msg.Amount,
		UserAddress: msg.UserAddress,
		CreatedAt:   uint64(time.Now().Unix()),
	}
	k.AppendEnergyToken(ctx, energyToken)

	var tokenHistory = types.TokensHistory{
		IssuerAddress: msg.Creator,
		UserAddress:   energyToken.UserAddress,
		CreatedAt:     energyToken.CreatedAt,
		Amount:        energyToken.Amount,
		TokenName:     tokenParams.Name,
		TargetAddress: energyToken.UserAddress,
		OperationType: "mint",
	}
	k.AppendTokensHistory(ctx, tokenHistory)

	return &types.MsgMintTokenResponse{}, nil
}
