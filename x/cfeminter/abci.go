package cfeminter

import (
	"github.com/armon/go-metrics"
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
		k.Logger(ctx).Error("cfeminter beggin block mint error", "error", err.Error())
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

	k.Logger(ctx).Debug("cfeminter beggin block data", "amount", float32(amount.Int64()), "inflationStr", inflationStr)
	err = ctx.EventManager().EmitTypedEvent(&types.Mint{
		BondedRatio: k.BondedRatio(ctx).String(),
		Inflation:   inflationStr,
		Amount:      amount.String(),
	})
	if err != nil {
		k.Logger(ctx).Error("mint emit event error", "error", err.Error())
	}
}
