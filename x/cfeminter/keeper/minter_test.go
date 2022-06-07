package keeper_test

import (
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetMinter(t *testing.T) {
	k, ctx := testkeeper.CfeminterKeeper(t)
	now := time.Now()
	linearMinter := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	end := now.AddDate(0, 0, 4)
	period1 := types.MintingPeriod{OrderingId: 1, PeriodEnd: &end, Type: types.MintingPeriod_TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter}
	period2 := types.MintingPeriod{OrderingId: 1, Type: types.MintingPeriod_NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2}
	minter := types.Minter{Start: now, Periods: periods}

	k.SetMinter(ctx, minter)

	compareMinters(t, minter, k.GetMinter(ctx))
}

func compareMinters(t *testing.T, m1 types.Minter, m2 types.Minter) {
	require.True(t, m1.Start.Equal(m2.Start))
	for i, p1 := range m1.Periods {
		p2 := m2.Periods[i]
		if (p1.PeriodEnd == nil) {
			require.Nil(t, p2.PeriodEnd)
		} else {
			require.True(t, p1.PeriodEnd.Equal(*p2.PeriodEnd))
		}
		require.EqualValues(t, p1.OrderingId, p2.OrderingId)
		require.EqualValues(t, p1.TimeLinearMinter, p2.TimeLinearMinter)
		require.EqualValues(t, p1.Type, p2.Type)
	}
}
