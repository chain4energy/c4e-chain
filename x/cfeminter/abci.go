package cfeminter

import (
	"time"

	"github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/keeper"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	amount, err := k.Mint(ctx)
	if err != nil {
		k.Logger(ctx).Error("mint error", "error", err.Error())
		panic(err)
	}

	if amount.IsInt64() {
		defer telemetry.SetGaugeWithLabels(
			[]string{types.ModuleName, "minted_tokens"},
			float32(amount.Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", k.MintDenom(ctx))},
		)
	}

	inflation, err := k.GetCurrentInflation(ctx)
	var inflationStr string
	if err != nil {
		inflationStr = types.UndefinedInflation
	} else {
		inflationStr = inflation.String()
	}
	bondedRatio := k.BondedRatio(ctx).String()
	k.Logger(ctx).Debug("minted", "amount", amount, "bondedRatio", bondedRatio, "inflation", inflationStr)
	event := &types.EventMint{
		BondedRatio: bondedRatio,
		Inflation:   inflationStr,
		Amount:      amount.String(),
	}
	err = ctx.EventManager().EmitTypedEvent(event)
	if err != nil {
		k.Logger(ctx).Error("mint emit event error", "event", event, "error", err.Error())
	}
}
