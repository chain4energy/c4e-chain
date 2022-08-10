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
		Amount:      msg.Amount * 1000000,
		UserAddress: msg.UserAddress,
		CreatedAt:   uint64(time.Now().Unix()),
	}

	k.AppendEnergyToken(ctx, energyToken)

	return &types.MsgMintTokenResponse{}, nil
}
