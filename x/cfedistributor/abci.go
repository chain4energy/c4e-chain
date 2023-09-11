package cfedistributor

import (
	"time"

	"github.com/chain4energy/c4e-chain/v2/x/cfedistributor/keeper"
	"github.com/chain4energy/c4e-chain/v2/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	subDistributors := k.GetParams(ctx).SubDistributors
	states := k.GetAllStates(ctx)

	for _, subDistributor := range subDistributors {
		allCoinsToDistribute := k.PrepareCoinsToDistribute(subDistributor.Sources, ctx, states, subDistributor.Name)
		if allCoinsToDistribute.IsZero() {
			continue
		}

		newStates, distributionEvents, burnEvent := k.StartDistributionProcess(ctx, &states, allCoinsToDistribute, subDistributor)
		for _, event := range distributionEvents {
			err := ctx.EventManager().EmitTypedEvent(event)
			if err != nil {
				k.Logger(ctx).Error("distributions emit event error", "event", event, "error", err.Error())
			}
		}
		if burnEvent != nil {
			err := ctx.EventManager().EmitTypedEvent(burnEvent)
			if err != nil {
				k.Logger(ctx).Error("distributions emit event error", "distribution_burn", burnEvent, "error", err.Error())
			}
		}
		states = *newStates
	}

	k.SendCoinsFromStates(ctx, states)
}
