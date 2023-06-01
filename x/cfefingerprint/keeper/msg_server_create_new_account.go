package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)
// Temporary message - to remove in the future
func (k msgServer) CreateNewAccount(goCtx context.Context, msg *types.MsgCreateNewAccount) (*types.MsgCreateNewAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: add logic to check that only a user with ADMIN role can add new blockchain users

	accAddress, _ := sdk.AccAddressFromBech32(msg.AccAddressString)
	newAccount := k.authKeeper.NewAccountWithAddress(ctx, accAddress)

	k.authKeeper.SetAccount(ctx, newAccount)

	return &types.MsgCreateNewAccountResponse{AccountNumber: newAccount.GetAccountNumber()}, nil
}
