package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/stretchr/testify/require"
)

const (
	validReferenceKey = "4e0f41db62899e37307c959e7cc2de62def852609428667d3682609c7dd10582"
	validPayloadHash  = "YWFhYWZmZjQ0cnJmZmZkc2RlZGVmZXRldAo="
)

func TestGetPayloadLinkNotFound(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	_, found := k.GetPayloadLink(ctx, validReferenceKey)

	require.False(t, found)
}

func TestAddPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	referenceId, err := k.CreatePayloadLink(ctx, validPayloadHash)
	require.NoError(t, err)

	_, err = k.MustGetPayloadLinkByReferenceId(ctx, *referenceId)
	require.NoError(t, err)
}

func TestAddEmptyPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	referenceId, err := k.CreatePayloadLink(ctx, "")
	require.EqualError(t, err, "payloadHash cannot be empty: wrong param value")
	require.Nil(t, referenceId)
}

func TestVerifyPayloadLink(t *testing.T) {
	k, ctx := testkeeper.CfefingerprintKeeper(t)

	referenceId, err := k.CreatePayloadLink(ctx, validPayloadHash)
	require.NoError(t, err)

	_, err = k.MustGetPayloadLinkByReferenceId(ctx, *referenceId)
	require.NoError(t, err)

	err = k.VerifyPayloadLink(ctx, *referenceId, validPayloadHash)
	require.NoError(t, err)

	err = k.VerifyPayloadLink(ctx, *referenceId, "invalidPayloadHash")
	require.EqualError(t, err, "expected payload link 8735a3b5ce24f9fde4c14ce8f7ccc7ce9365b7134f9cdb79f2d881aec949ca55 "+
		"for reference id b6bd44eecbe53dd2080399d709d495f59111ef03282f3dfff07453862077fdab doesn't match payload "+
		"link 2fb35e5473e249502ffbf7c1cfd425f5fd4bf5c13ba185119e38851ac37afb8c: wrong param value")

	err = k.VerifyPayloadLink(ctx, "invalidReferenceId", validPayloadHash)
	require.EqualError(t, err, "payload link not found for reference id invalidReferenceId: not found")
}
