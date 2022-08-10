package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTokenParams(goCtx context.Context, msg *types.MsgCreateTokenParams) (*types.MsgCreateTokenParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var token = types.TokenParams{
		Index:          msg.Name,
		Name:           msg.Name,
		TradingCompany: msg.TradingCompany,
		BurningTime:    msg.BurningTime,
		BurningType:    msg.BurningType,
		SendPrice:      msg.SendPrice,
		MintAccount:    msg.Creator,
	}

	_, isFound := k.GetTokenParams(ctx, token.Name)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Token with this name already exists")
	}

	k.SetTokenParams(ctx, token)

	return &types.MsgCreateTokenParamsResponse{}, nil
}
