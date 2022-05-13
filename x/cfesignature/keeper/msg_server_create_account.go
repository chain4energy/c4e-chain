package keeper

import (
	"context"
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: add logic to check that only a user with ADMIN role can add new blockchain users

	accAddress, _ := sdk.AccAddressFromBech32(msg.AccAddressString)
	newAccount := k.authKeeper.NewAccountWithAddress(ctx, accAddress)

	var pk cryptotypes.PubKey
	bytes := []byte(msg.PubKeyString)
	k.proto.UnmarshalInterfaceJSON(bytes, &pk)

	newAccount.SetPubKey(pk)
	k.authKeeper.SetAccount(ctx, newAccount)

	return &types.MsgCreateAccountResponse{AccountNumber: fmt.Sprint(newAccount.GetAccountNumber())}, nil

}
