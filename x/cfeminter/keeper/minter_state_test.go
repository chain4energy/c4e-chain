package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetMinterState(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)

	minterState := types.MinterState{CurrentPosition: 7, AmountMinted: sdk.NewInt(123412)}

	k.SetMinterState(ctx, minterState)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
}
