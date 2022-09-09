package cfeminter

import (
	"github.com/chain4energy/c4e-chain/x/cfeminter/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	amount, err := k.Mint(ctx)
	if err != nil {
		panic(err)
	}

	if amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(amount.Int64()), "minted_tokens")
	}

	inflation, err := k.GetCurrentInflation(ctx)
	var inflationStr string
	if err != nil {
		inflationStr = types.UndefinedInflation
	} else {
		inflationStr = inflation.String()
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyBondedRatio, k.BondedRatio(ctx).String()),
			sdk.NewAttribute(types.AttributeKeyInflation, inflationStr),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
		),
	)
}
