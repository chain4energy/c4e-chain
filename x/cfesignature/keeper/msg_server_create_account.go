package keeper

import (
	"context"
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	defer telemetry.IncrCounter(1, types.ModuleName, "create account message")
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: add logic to check that only a user with ADMIN role can add new blockchain users

	accAddress, err := sdk.AccAddressFromBech32(msg.AccAddressString)
	if err != nil {
		k.Logger(ctx).Error("create account parsing error", "error", err.Error())
		return nil, err
	}
	newAccount := k.authKeeper.NewAccountWithAddress(ctx, accAddress)

	var pk cryptotypes.PubKey
	bytes := []byte(msg.PubKeyString)
	err = k.proto.UnmarshalInterfaceJSON(bytes, &pk)
	if err != nil {
		k.Logger(ctx).Error("unmarshal interface JSON error", "error", err.Error())
		return nil, err
	}

	err = newAccount.SetPubKey(pk)
	if err != nil {
		k.Logger(ctx).Error("new account set pub key error", "error", err.Error())
		return nil, err
	}
	k.Logger(ctx).Debug("auth keeper set account", "newAccount", newAccount.String())
	k.authKeeper.SetAccount(ctx, newAccount)

	return &types.MsgCreateAccountResponse{AccountNumber: fmt.Sprint(newAccount.GetAccountNumber())}, nil

}
