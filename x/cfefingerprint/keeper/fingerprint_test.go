package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/util"
	"github.com/stretchr/testify/require"
)

func TestGetPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	referenceKey := "4e0f41db62899e37307c959e7cc2de62def852609428667d3682609c7dd10582"

	result := k.CheckIfPayloadLinkExists(ctx, referenceKey)

	require.EqualValues(t, false, !result)
}

func TestAddPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	payloadHash := "YWFhYWZmZjQ0cnJmZmZkc2RlZGVmZXRldAo="

	// create reference payload link
	referenceID, err := k.CreatePayloadLink(ctx, payloadHash)
	if err != nil {
		panic(err)
	}

	// Check if a Payload Link was stored at the given key
	referenceKey := util.CalculateHash(referenceID)
	result := k.CheckIfPayloadLinkExists(ctx, referenceKey)

	require.EqualValues(t, false, result)
}

func TestVerifyPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	payloadHash := "YWFhYWZmZjQ0cnJmZmZkc2RlZGVmZXRldAo="

	// create reference payload link
	referenceID, err := k.CreatePayloadLink(ctx, payloadHash)
	if err != nil {
		panic(err)
	}

	// Check if a Payload Link was stored at the given key
	referenceKey := util.CalculateHash(referenceID)
	result := k.CheckIfPayloadLinkExists(ctx, referenceKey)
	require.EqualValues(t, false, result)

	// verify payloadHash
	result, err = k.VerifyPayloadLink(ctx, referenceID, payloadHash)
	if err != nil {
		require.Fail(t, err.Error())
	}
	require.Equal(t, true, result)

}
