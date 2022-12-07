package keeper_test

import (
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetMinterState(t *testing.T) {
	k, ctx, _ := testkeeper.CfeminterKeeper(t)

	minterState := types.MinterState{
		SequenceId:                  7,
		AmountMinted:                sdk.NewInt(123412),
		RemainderToMint:             sdk.ZeroDec(),
		RemainderFromPreviousPeriod: sdk.ZeroDec(),
		LastMintBlockTime:           time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC),
	}

	k.SetMinterState(ctx, minterState)
	require.EqualValues(t, minterState, k.GetMinterState(ctx))
}
