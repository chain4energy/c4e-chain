package keeper_test

import (
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMint(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)
	params := types.DefaultParams()
	params.MintDenom = "uc4e"
	k.SetParams(ctx, params)

	minterState := types.MinterState{CurrentOrderingId: 1, AmountMinted: sdk.NewInt(0)}
	k.SetMinterState(ctx, minterState)

	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.Local)
	endTime1 := startTime.Add(time.Duration(345600000000 * 1000000))
	endTime2 := endTime1.Add(time.Duration(345600000000 * 1000000))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &endTime1, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{OrderingId: 2, PeriodEnd: &endTime2, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{OrderingId: 3, Type: types.MintingPeriod_NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	minter := types.Minter{Start: startTime, Periods: periods}

	k.SetMinter(ctx, minter)

	ctx = ctx.WithBlockTime(startTime)
	amount, err := k.Mint(ctx)
	require.NoError(t, err)
	require.EqualValues(t, sdk.NewInt(0), amount)



	// blockTime := startTime.Add(time.Duration(345600000000 * 1000000 / 2))

	
}
