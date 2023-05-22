package keeper_test

import (
	"context"
	"encoding/hex"
	"strconv"
	"testing"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EnergychainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestCreateAccount(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	logger := log.TestingLogger()
	creatorAccAddressString := "cosmos1lcx66v2yqna3tk2urfpjmyq6rdj9c8uey3pzel"

	_, publicKeyFromTestUtil, accAddressFromUtil := testdata.KeyTestPubAddr()
	_, publicKeyFromTestUtil2, accAddressFromUtil2 := testdata.KeyTestPubAddr()

	// create two accounts with two different account addresses
	msg := types.NewMsgCreateAccount(creatorAccAddressString, accAddressFromUtil.String(), getPubKeyJSON(publicKeyFromTestUtil.String()))
	msgCreateAccountResponse, err := msgServer.CreateAccount(ctx, msg)
	if err != nil {
		require.Fail(t, "test failed")
	}
	logger.Debug(msgCreateAccountResponse.GetAccountNumber())

	msg2 := types.NewMsgCreateAccount(creatorAccAddressString, accAddressFromUtil2.String(), getPubKeyJSON(publicKeyFromTestUtil2.String()))
	msgCreateAccountResponse2, err := msgServer.CreateAccount(ctx, msg2)
	if err != nil {
		require.Fail(t, "test failed")
	}
	logger.Debug(msgCreateAccountResponse2.GetAccountNumber())

	// check if we get different and sequential account numbers
	accountNumber1, err := strconv.Atoi(msgCreateAccountResponse.GetAccountNumber())
	if err != nil {
		require.Fail(t, "test failed")
	}
	accountNumber2, err := strconv.Atoi(msgCreateAccountResponse2.GetAccountNumber())
	if err != nil {
		require.Fail(t, "test failed")
	}

	result := (accountNumber1 + 1) == accountNumber2
	require.EqualValues(t, result, true)
}

// getPubKeyJSON formats public key using Protobufs JSON
func getPubKeyJSON(pk string) string {

	pubKey := getPublicKeyObjectFromString(pk)
	apk, _ := codectypes.NewAnyWithValue(pubKey)
	bz, _ := codec.ProtoMarshalJSON(apk, nil)
	return string(bz)
}

func getPublicKeyObjectFromString(pk string) cryptotypes.PubKey {
	pubKeyBytes, _ := hex.DecodeString(pk)
	return &secp256k1.PubKey{Key: pubKeyBytes}
}
