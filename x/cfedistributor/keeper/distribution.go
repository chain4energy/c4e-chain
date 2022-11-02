package keeper

import (
	"github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) PrepareCoinToDistributeForMainAccount(ctx sdk.Context, states []types.State, subDistributorName string) sdk.DecCoins {
	coinsToDistribute := sdk.NewDecCoinsFromCoins(k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)...)
	if len(coinsToDistribute) > 0 {
		sum := getRamainsSum(&states)
		coinsToDistribute = coinsToDistribute.Sub(sum)
	}
	k.Logger(ctx).Debug("prepare coins to distribute for main account", "subDistr", subDistributorName, "coins", coinsToDistribute.String())

	return coinsToDistribute
}

func prepareCoinToDistributeForModuleAccount(ctx sdk.Context, k Keeper, source types.Account, subDistributorName string) sdk.DecCoins {
	coinsToSend := k.GetAccountCoinsForModuleAccount(ctx, source.Id)
	coinsToDistribute := sdk.NewDecCoinsFromCoins(coinsToSend...)

	if len(coinsToDistribute) > 0 {
		err := k.SendCoinsFromModuleToModule(ctx, coinsToSend, source.Id, types.DistributorMainAccount)
		if err != nil {
			k.Logger(ctx).Error("prep coins module - send coins to main account", "subDistributorName", subDistributorName, "source", source, "error", err.Error())
			return nil
		}
	}
	k.Logger(ctx).Debug("prepare coins to distribute for module account", "subDistr", subDistributorName,
		"account", source.Id, "coinsToDistribute", coinsToDistribute.String())
	return coinsToDistribute
}

func prepareCoinToDistributeForBaseAccount(ctx sdk.Context, k Keeper, source types.Account, subDistributorName string) sdk.DecCoins {
	srcAccount, _ := sdk.AccAddressFromBech32(source.Id)
	coinsToSend := k.GetAccountCoins(ctx, srcAccount)
	coinsToDistribute := sdk.NewDecCoinsFromCoins(coinsToSend...)

	if len(coinsToDistribute) > 0 {
		err := k.SendCoinsToModuleAccount(ctx, coinsToSend, srcAccount, types.DistributorMainAccount)
		if err != nil {
			k.Logger(ctx).Error("prepare coin to distribute for internal account error", "error", err.Error())
			return nil
		}
	}
	k.Logger(ctx).Debug("prepare coins to distribute for base account", "subDistr", subDistributorName,
		"account", source.Id, "coinsToDistribute", coinsToDistribute.String())
	return coinsToDistribute
}

func prepareLeftCoinToDistribute(coinsToDistribute sdk.DecCoins, source types.Account, states []types.State) sdk.DecCoins {
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

func (k Keeper) PrepareCoinToDistributeForNotMainAccount(ctx sdk.Context, source types.Account, states []types.State, subDistributorName string) sdk.DecCoins {
	var coinsToDistribute sdk.DecCoins
	if types.MODULE_ACCOUNT == source.Type {
		coinsToDistribute = prepareCoinToDistributeForModuleAccount(ctx, k, source, subDistributorName)
	} else if types.INTERNAL_ACCOUNT != source.Type {
		coinsToDistribute = prepareCoinToDistributeForBaseAccount(ctx, k, source, subDistributorName)
	} else {
		coinsToDistribute = sdk.NewDecCoins()

	}
	return prepareLeftCoinToDistribute(coinsToDistribute, source, states)
}
func burnCoins(ctx sdk.Context, k Keeper, state *types.State) {
	toSend, change := state.CoinsStates.TruncateDecimal()

	if err := k.BurnCoinsForSpecifiedModuleAccount(ctx, toSend, types.DistributorMainAccount); err != nil {
		ctx.Logger().Error("burn coins error", "state", state, "error", err.Error())
	} else {
		k.Logger(ctx).Debug("Coins burned", "coins", toSend)
		defer telemetry.SetGaugeWithLabels(
			[]string{types.ModuleName, "coin_send", types.BurnDestination},
			float32(toSend.AmountOf(types.DenomToTrace).Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", types.DenomToTrace)},
		)
		state.CoinsStates = change
	}
}

func sendCoinsToModuleAccount(ctx sdk.Context, k Keeper, state *types.State) {
	toSend, change := state.CoinsStates.TruncateDecimal()

	if err := k.SendCoinsFromModuleToModule(ctx, toSend, types.DistributorMainAccount, state.Account.Id); err != nil {
		ctx.Logger().Error("send coins to module account dst error", "accountId", state.Account.Id, "error", err.Error())
	} else {
		k.Logger(ctx).Debug("coins sent to module account dst", "accountId", state.Account.Id, "toSend", toSend.String())
		defer telemetry.SetGaugeWithLabels(
			[]string{types.ModuleName, "coin_send", state.Account.Id},
			float32(toSend.AmountOf(types.DenomToTrace).Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", types.DenomToTrace)},
		)
		state.CoinsStates = change
	}
}

func sendCoinsToBaseAccount(ctx sdk.Context, k Keeper, state *types.State) {
	toSend, change := state.CoinsStates.TruncateDecimal()

	if dstAccount, err := sdk.AccAddressFromBech32(state.Account.Id); err != nil {
		k.Logger(ctx).Error("destination base account address parsing error", "accountId", state.Account.Id, "error", err.Error())
	} else if err := k.SendCoinsFromModuleAccount(ctx, toSend, types.DistributorMainAccount, dstAccount); err != nil {
		k.Logger(ctx).Error("send coins to base account dst error", "accountId", state.Account.Id, "toSend", toSend, "error", err.Error())
	} else {
		k.Logger(ctx).Debug("coins sent to base account dst", "accountId", state.Account.Id, "toSend", toSend)
		defer telemetry.SetGaugeWithLabels(
			[]string{types.ModuleName, "coin_send", state.Account.Id},
			float32(toSend.AmountOf(types.DenomToTrace).Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", types.DenomToTrace)},
		)
		state.CoinsStates = change
	}
}

func (k Keeper) SendCoinsFromStates(ctx sdk.Context, states []types.State) {
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
		if coin.Amount.GTE(sdk.NewDec(1)) {
			return true
		}
	}

	return false
}

func addSharesToBurnState(ctx sdk.Context, k Keeper, localRemains *[]types.State, calculatedShare sdk.DecCoins, findState func() int) *[]types.State {
	return addSharesToState(ctx, k, localRemains, true, nil, calculatedShare, findState)
}

func addSharesToAccountState(ctx sdk.Context, k Keeper, localRemains *[]types.State, account *types.Account, calculatedShare sdk.DecCoins, findState func() int) *[]types.State {
	return addSharesToState(ctx, k, localRemains, false, account, calculatedShare, findState)
}

func addSharesToState(ctx sdk.Context, k Keeper, localRemains *[]types.State, burn bool, account *types.Account, calculatedShare sdk.DecCoins, findState func() int) *[]types.State {
	pos := findState()
	logger := k.Logger(ctx).With("localRemains", localRemains, "account", account, "burn", burn,
		"calculatedShare", calculatedShare.String(), "pos", pos)
	if pos < 0 {
		var state types.State
		if burn || account == nil {
			state = types.State{Account: &types.Account{}, CoinsStates: sdk.NewDecCoins(), Burn: true}
		} else {
			state = types.State{Account: account, CoinsStates: sdk.NewDecCoins(), Burn: false}
		}
		withAppended := append(*localRemains, state)

		localRemains = &withAppended
		pos = len(*localRemains) - 1
		logger.Debug("add shares to state", "state", state)
	} else {
		logger.Debug("add shares to state")
	}
	(*localRemains)[pos].CoinsStates = (*localRemains)[pos].CoinsStates.Add(calculatedShare...)
	return localRemains
}

func (k Keeper) StartDistributionProcess(ctx sdk.Context, states *[]types.State, coinsToDistributeDec sdk.DecCoins, subDistributor types.SubDistributor, list *types.DistributionsResult) *[]types.State {
	k.Logger(ctx).Debug("start distribution process", "subDistributor", subDistributor.String(),
		"coinsToDistributeDec", coinsToDistributeDec.String())
	localRemains := states
	defaultShare := coinsToDistributeDec
	for _, share := range subDistributor.Destination.Share {
		if share.Account.Type == types.MAIN {
			continue
		}
		calculatedShare := calculatePercentage(share.Percent, coinsToDistributeDec)
		defaultShare = defaultShare.Sub(calculatedShare)
		if !calculatedShare.IsZero() {
			findFunc := func() int {
				return findAccountState(localRemains, &share.Account)
			}

			localRemains = addSharesToAccountState(ctx, k, localRemains, &share.Account, calculatedShare, findFunc)
			list.DistributionResult = append(list.DistributionResult, &types.DistributionResult{
				Source:      subDistributor.Sources,
				Destination: &share.Account,
				CoinSend:    calculatedShare,
			})
		}
	}

	if subDistributor.Destination.BurnShare.Percent != sdk.ZeroDec() {
		calculatedShare := calculatePercentage(subDistributor.Destination.BurnShare.Percent, coinsToDistributeDec)
		defaultShare = defaultShare.Sub(calculatedShare)
		if !calculatedShare.IsZero() {
			findFunc := func() int {
				return findBurnState(localRemains)
			}
			localRemains = addSharesToBurnState(ctx, k, localRemains, calculatedShare, findFunc)
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
		localRemains = addSharesToAccountState(ctx, k, localRemains, &accountDefault, defaultShare, findFunc)
		list.DistributionResult = append(list.DistributionResult, &types.DistributionResult{
			Source:      subDistributor.Sources,
			Destination: &subDistributor.Destination.Account,
			CoinSend:    defaultShare,
		})
	}
	k.Logger(ctx).Debug("start distribution process ret", "subDistributor", subDistributor.String(), "localRemains", localRemains)
	return localRemains
}
