package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/util"
	"github.com/stretchr/testify/require"
)

func TestGetPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	referenceID := "4e0f41db62899e37307c959e7cc2de62def852609428667d3682609c7dd10582"

	result := k.CheckIfPayloadLinkExists(ctx, referenceID)

	require.EqualValues(t, false, !result)
}

func TestAddPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	referenceID := "4e0f41db62899e37307c959e7cc2de62def852609428667d3682609c7dd10582"
	payloadHash := "YWFhYWZmZjQ0cnJmZmZkc2RlZGVmZXRldAo="
	// create reference payload link
	referenceKey := util.CalculateHash(referenceID)
	referenceValue := util.CalculateHash(util.HashConcat(referenceID, payloadHash))

	// store payload link
	k.AppendPayloadLink(ctx, referenceKey, referenceValue)

	// Check if a Payload Link was stored at the given key
	result := k.CheckIfPayloadLinkExists(ctx, referenceKey)

	require.EqualValues(t, false, result)
}

func TestVerifyPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	referenceID := "4e0f41db62899e37307c959e7cc2de62def852609428667d3682609c7dd10582"
	payloadHash := "YWFhYWZmZjQ0cnJmZmZkc2RlZGVmZXRldAo="
	// create reference payload link
	referenceKey := util.CalculateHash(referenceID)
	referenceValue := util.CalculateHash(util.HashConcat(referenceID, payloadHash))

	// store payload link
	k.AppendPayloadLink(ctx, referenceKey, referenceValue)

	// Check if a Payload Link was stored at the given key
	result := k.CheckIfPayloadLinkExists(ctx, referenceKey)
	if result {
		require.Fail(t, "PayloadLink not found")
	}

	// fetch data published on ledger
	ledgerPayloadLink, err := k.GetPayloadLink(ctx, referenceID)
	if err != nil {
		require.Fail(t, err.Error())
	}

	// calculate expeced data based on payload hash
	// so called reference value
	expectedPayloadLink := util.CalculateHash(util.HashConcat(referenceID, payloadHash))

	// verify ledger matches the payloadhash
	require.Equal(t, expectedPayloadLink, ledgerPayloadLink)

}
