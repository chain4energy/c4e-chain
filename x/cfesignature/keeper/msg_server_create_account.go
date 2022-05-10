package keeper

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	accAddress, _ := sdk.AccAddressFromBech32(msg.AccAddressString)
	newAccount := k.authKeeper.NewAccountWithAddress(ctx, accAddress)

	var pk cryptotypes.PubKey
	bytes := []byte(msg.PubKeyString)
	k.proto.UnmarshalInterfaceJSON(bytes, &pk)

	publicKeyString := pk.String()

	newAccount.SetPubKey(pk)
	k.authKeeper.SetAccount(ctx, newAccount)

	return &types.MsgCreateAccountResponse{AccountId: newAccount.GetAddress().String() + publicKeyString}, nil

}
