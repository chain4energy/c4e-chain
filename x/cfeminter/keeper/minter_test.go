package keeper_test

import (
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	testminter.CompareMinters(t, minter, k.GetMinter(ctx))
}
