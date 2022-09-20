package cfedistributor

import (
	"github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func calculatePercentage(sharePercent sdk.Dec, coinsToDistributeDec sdk.DecCoins) sdk.DecCoins {
	if !coinsToDistributeDec.IsAllPositive() {
		return sdk.NewDecCoins()
	}
	percentInDecForm := sharePercent.QuoInt64(100)
	return coinsToDistributeDec.MulDecTruncate(percentInDecForm)
}

func findBurnState(states *[]types.State) int {
	for pos, state := range *states {
		if state.Burn {
			return pos
		}
	}
	return -1
}

func findAccountState(states *[]types.State, account *types.Account) int {
	for pos, state := range *states {
		if state.Account.Id == account.Id && state.Account.Id != "" && &state.Account.Id != nil {
			return pos
		} else if state.Account.Id == account.Id && state.Account.Id == "" {
			return pos
		}
	}
	return -1
}

func getRamainsSum(states *[]types.State) sdk.DecCoins {
	sum := sdk.NewDecCoins()
	for _, state := range *states {
		sum = sum.Add(state.CoinsStates...)
	}
	return sum
}

func prepareCoinToDistributeForMainAccount(ctx sdk.Context, k keeper.Keeper, coinsToDistribute sdk.DecCoins, states []types.State) sdk.DecCoins {
	coinsToDistribute = sdk.NewDecCoinsFromCoins(k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)...)
	k.Logger(ctx).Debug("IsMainCollector: " + coinsToDistribute.String())
	if len(coinsToDistribute) > 0 {
		sum := getRamainsSum(&states)
		coinsToDistribute = coinsToDistribute.Sub(sum)
	}

	return coinsToDistribute
}

func prepareCoinToDistributeForModuleAccount(ctx sdk.Context, k keeper.Keeper, coinsToDistribute sdk.DecCoins, source types.Account) sdk.DecCoins {
	k.Logger(ctx).Debug("Module account: " + source.Id)
	coinsToSend := k.GetAccountCoinsForModuleAccount(ctx, source.Id)
	coinsToDistribute = sdk.NewDecCoinsFromCoins(coinsToSend...)
	k.Logger(ctx).Debug("IsModuleAccount: " + source.Id + " - " + coinsToDistribute.String())

	if len(coinsToDistribute) > 0 {
		k.SendCoinsFromModuleToModule(ctx, coinsToSend, source.Id, types.DistributorMainAccount)
	}

	return coinsToDistribute
}

func prepareCoinToDistributeForInternalAccount(ctx sdk.Context, k keeper.Keeper, coinsToDistribute sdk.DecCoins, source types.Account) sdk.DecCoins {
	k.Logger(ctx).Debug("Internal account: " + source.Id)

	srcAccount, _ := sdk.AccAddressFromBech32(source.Id)
	coinsToSend := k.GetAccountCoins(ctx, srcAccount)
	coinsToDistribute = sdk.NewDecCoinsFromCoins(coinsToSend...)
	k.Logger(ctx).Debug("BaseAccount: " + source.Id + " - " + coinsToDistribute.String())

	if len(coinsToDistribute) > 0 {
		k.SendCoinsToModuleAccount(ctx, coinsToSend, srcAccount, types.DistributorMainAccount)
	}

	return coinsToDistribute
}

func prepareLeftedCoinToDistribute(coinsToDistribute sdk.DecCoins, source types.Account, states []types.State) sdk.DecCoins {
	pos := findAccountState(&states, &source)
	if pos >= 0 {
		coin := states[pos].CoinsStates
		if !coin.IsZero() {
			states[pos].CoinsStates = sdk.NewDecCoins()
			coinsToDistribute = coinsToDistribute.Add(coin...)
		}
	}

	return coinsToDistribute
}

func prepareCoinToDistributeForNotMainAccount(ctx sdk.Context, k keeper.Keeper, coinsToDistribute sdk.DecCoins, source types.Account, states []types.State) sdk.DecCoins {
	if types.MODULE_ACCOUNT == source.Type {
		coinsToDistribute = prepareCoinToDistributeForModuleAccount(ctx, k, coinsToDistribute, source)

	} else if types.INTERNAL_ACCOUNT != source.Type {
		coinsToDistribute = prepareCoinToDistributeForInternalAccount(ctx, k, coinsToDistribute, source)
	}

	return prepareLeftedCoinToDistribute(coinsToDistribute, source, states)
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	subDistributors := k.GetParams(ctx).SubDistributors
	states := k.GetAllStates(ctx)
	distributionsResult := types.DistributionsResult{}

	for _, subDistributor := range subDistributors {
		k.Logger(ctx).Debug("BeginBlock - cfedistr: " + subDistributor.Name)
		allCoinsToDistribute := sdk.NewDecCoins()
		for _, source := range subDistributor.Sources {
			k.Logger(ctx).Debug("Sources: " + source.String())

			var coinsToDistribute = sdk.NewDecCoins()
			if source.Type == types.MAIN {
				coinsToDistribute = prepareCoinToDistributeForMainAccount(ctx, k, coinsToDistribute, states)
			} else {
				coinsToDistribute = prepareCoinToDistributeForNotMainAccount(ctx, k, coinsToDistribute, *source, states)
			}

			if len(coinsToDistribute) == 0 {
				continue
			}
			allCoinsToDistribute = allCoinsToDistribute.Add(coinsToDistribute...)
		}

		if allCoinsToDistribute.IsZero() {
			continue
		}
		states = *StartDistributionProcess(&states, allCoinsToDistribute, subDistributor, &distributionsResult)
	}

	ctx.EventManager().EmitTypedEvent(&distributionsResult)
	sendCoinsFromStates(ctx, k, states)
}

func burnCoins(ctx sdk.Context, k keeper.Keeper, state *types.State) {
	toSend, change := state.CoinsStates.TruncateDecimal()

	if error := k.BurnCoinsForSpecifiedModuleAccount(ctx, toSend, types.DistributorMainAccount); error != nil {
		ctx.Logger().Error("Can not burn coin: " + error.Error())

	} else {
		k.Logger(ctx).Debug("Successful burn coin: " + toSend.String())
		defer telemetry.SetGaugeWithLabels(
			[]string{"coin_send", types.BurnDestination},
			float32(toSend.AmountOf(types.DenomToTrace).Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", types.DenomToTrace)},
		)
		state.CoinsStates = change
	}
}

func sendCoinsToModuleAccount(ctx sdk.Context, k keeper.Keeper, state *types.State) {
	toSend, change := state.CoinsStates.TruncateDecimal()

	if error := k.SendCoinsFromModuleToModule(ctx, toSend, types.DistributorMainAccount, state.Account.Id); error != nil {
		ctx.Logger().Error("Can not send coin: " + error.Error())

	} else {
		k.Logger(ctx).Debug("Successful send to: " + state.Account.Id + " - " + toSend.String())
		defer telemetry.SetGaugeWithLabels(
			[]string{"coin_send", state.Account.Id},
			float32(toSend.AmountOf(types.DenomToTrace).Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", types.DenomToTrace)},
		)
		state.CoinsStates = change
	}
}

func sendCoinsToBaseAccount(ctx sdk.Context, k keeper.Keeper, state *types.State) {
	toSend, change := state.CoinsStates.TruncateDecimal()

	if dstAccount, error := sdk.AccAddressFromBech32(state.Account.Id); error != nil {
		ctx.Logger().Error("Can not get addr from bech32: " + error.Error())

	} else if error := k.SendCoinsFromModuleAccount(ctx, toSend, types.DistributorMainAccount, dstAccount); error != nil {
		ctx.Logger().Error("Can not send coin: " + error.Error())

	} else {
		k.Logger(ctx).Debug("Successful send to : " + state.Account.Id + " - " + toSend.String())
		defer telemetry.SetGaugeWithLabels(
			[]string{"coin_send", state.Account.Id},
			float32(toSend.AmountOf(types.DenomToTrace).Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", types.DenomToTrace)},
		)
		state.CoinsStates = change
	}
}

func sendCoinsFromStates(ctx sdk.Context, k keeper.Keeper, states []types.State) {
	for _, state := range states {
		if types.INTERNAL_ACCOUNT != state.Account.Type && checkIfAnyCoinIsGTE1(state.CoinsStates) {

			if state.Burn {
				burnCoins(ctx, k, &state)

			} else if types.MODULE_ACCOUNT == state.Account.Type {
				sendCoinsToModuleAccount(ctx, k, &state)
			} else {
				sendCoinsToBaseAccount(ctx, k, &state)
			}
		}
		k.SetState(ctx, state)
	}
}

func checkIfAnyCoinIsGTE1(coins sdk.DecCoins) bool {
	if len(coins) == 0 {
		return false
	}
	for _, coin := range coins {
		if coin.Amount.GTE(sdk.MustNewDecFromStr("1")) {
			return true
		}
	}

	return false
}

func addSharesToState(localRemains *[]types.State, account types.Account, calculatedShare sdk.DecCoins, findState func() int) *[]types.State {
	pos := findState()
	if pos < 0 {
		state := types.State{}
		if &account.Type == nil || account.Type == "" {

			state = types.State{Account: &account, CoinsStates: sdk.NewDecCoins(), Burn: true}
		} else {
			state = types.State{Account: &account, CoinsStates: sdk.NewDecCoins(), Burn: false}
		}
		withAppended := append(*localRemains, state)

		localRemains = &withAppended
		pos = len(*localRemains) - 1
	}
	(*localRemains)[pos].CoinsStates = (*localRemains)[pos].CoinsStates.Add(calculatedShare...)
	return localRemains
}

func StartDistributionProcess(states *[]types.State, coinsToDistributeDec sdk.DecCoins, subDistributor types.SubDistributor, list *types.DistributionsResult) *[]types.State {
	percentShareSum := sdk.MustNewDecFromStr("0")
	localRemains := states
	for _, share := range subDistributor.Destination.Share {
		percentShareSum = percentShareSum.Add(share.Percent)
		if share.Account.Type == types.MAIN {
			continue
		}
		calculatedShare := calculatePercentage(share.Percent, coinsToDistributeDec)

		if !calculatedShare.IsZero() {
			findFunc := func() int {
				return findAccountState(localRemains, &share.Account)
			}
			localRemains = addSharesToState(localRemains, share.Account, calculatedShare, findFunc)
			list.DistributionResult = append(list.DistributionResult, &types.DistributionResult{
				Source:      subDistributor.Sources,
				Destination: &share.Account,
				CoinSend:    calculatedShare,
			})
		}
	}

	if subDistributor.Destination.BurnShare.Percent != sdk.MustNewDecFromStr("0") {
		percentShareSum = percentShareSum.Add(subDistributor.Destination.BurnShare.Percent)
		calculatedShare := calculatePercentage(subDistributor.Destination.BurnShare.Percent, coinsToDistributeDec)

		if !calculatedShare.IsZero() {
			findFunc := func() int {
				return findBurnState(localRemains)
			}
			localRemains = addSharesToState(localRemains, types.Account{}, calculatedShare, findFunc)
			list.DistributionResult = append(list.DistributionResult, &types.DistributionResult{
				Source: subDistributor.Sources,
				Destination: &types.Account{
					Id:   types.BurnDestination,
					Type: "",
				},
				CoinSend: calculatedShare,
			})
		}
	}

	accountDefault := subDistributor.Destination.GetAccount()

	if accountDefault.Type != types.MAIN {
		findFunc := func() int {

			return findAccountState(localRemains, &accountDefault)
		}

		defaultSharePercent := sdk.MustNewDecFromStr("100").Sub(percentShareSum)
		calculatedShare := calculatePercentage(defaultSharePercent, coinsToDistributeDec)
		localRemains = addSharesToState(localRemains, accountDefault, calculatedShare, findFunc)
		list.DistributionResult = append(list.DistributionResult, &types.DistributionResult{
			Source:      subDistributor.Sources,
			Destination: &subDistributor.Destination.Account,
			CoinSend:    calculatedShare,
		})
	}

	return localRemains
}
