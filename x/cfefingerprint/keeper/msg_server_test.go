package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CfefingerprintKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestCreateAccount(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	logger := log.TestingLogger()
	creatorAccAddressString := "cosmos1lcx66v2yqna3tk2urfpjmyq6rdj9c8uey3pzel"

	_, _, accAddressFromUtil := testdata.KeyTestPubAddr()
	_, _, accAddressFromUtil2 := testdata.KeyTestPubAddr()

	// create two accounts with two different account addresses
	msg := types.NewMsgCreateNewAccount(creatorAccAddressString, accAddressFromUtil.String())
	msgCreateNewAccountResponse, err := msgServer.CreateNewAccount(ctx, msg)
	if err != nil {
		require.Fail(t, "test failed")
	}
	logger.Debug(strconv.FormatUint(msgCreateNewAccountResponse.GetAccountNumber(), 10))

	msg2 := types.NewMsgCreateNewAccount(creatorAccAddressString, accAddressFromUtil2.String())
	msgCreateAccountResponse2, err := msgServer.CreateNewAccount(ctx, msg2)
	if err != nil {
		require.Fail(t, "test failed")
	}
	logger.Debug(strconv.FormatUint(msgCreateAccountResponse2.GetAccountNumber(), 10))

	// check if we get different and sequential account numbers
	accountNumber1 := msgCreateNewAccountResponse.GetAccountNumber()
	accountNumber2 := msgCreateAccountResponse2.GetAccountNumber()

	result := (accountNumber1 + 1) == accountNumber2
	require.EqualValues(t, result, true)
}
