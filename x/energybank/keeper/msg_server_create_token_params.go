package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/energybank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTokenParams(goCtx context.Context, msg *types.MsgCreateTokenParams) (*types.MsgCreateTokenParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateTokenParamsResponse{}, nil
}
