package cfeminter

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mint/types"
	"time"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// halvingMinter := k.GetHalvingMinter(ctx)

	// halvingMinter.NewCoinsMint = halvingMinter.NextCointCount(ctx.BlockHeight())
	// mintedCoin := sdk.NewCoin(halvingMinter.MintDenom, sdk.NewInt(halvingMinter.NewCoinsMint))
	// mintedCoins := sdk.NewCoins(mintedCoin)

	// k.SetHalvingMinter(ctx, halvingMinter)

	// // mint coin from bank
	// err := k.MintCoins(ctx, mintedCoins)
	// if err != nil {
	// 	panic(err)
	// }

	// // send coint to
	// err = k.SendCoinsToCommonAccount(ctx, mintedCoins)
	// if err != nil {
	// 	panic(err)
	// }

	amount, err := k.Mint(ctx)
	if err != nil {
		panic(err)
	}

	if amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(amount.Int64()), "minted_tokens")
	}
}
