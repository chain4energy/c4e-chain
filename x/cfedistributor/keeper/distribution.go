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
	return coinsToDistributeDec.MulDecTruncate(sharePercent)
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
		if state.Account.Id == account.Id && state.Account.Id != "" {
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
		sum = sum.Add(state.Remains...)
	}
	return sum
}

func (k Keeper) PrepareCoinsToDistribute(sources []*types.Account, ctx sdk.Context, states []types.State, subDistributorName string) sdk.DecCoins {
	allCoinsToDistribute := sdk.NewDecCoins()
	for _, source := range sources {
		var coinsToDistribute sdk.DecCoins
		if source.Type == types.Main {
			coinsToDistribute = k.prepareCoinToDistributeForMainAccount(ctx, states, subDistributorName)
		} else {
			coinsToDistribute = k.prepareCoinToDistributeForNotMainAccount(ctx, *source, states, subDistributorName)
		}

		if len(coinsToDistribute) == 0 {
			continue
		}
		allCoinsToDistribute = allCoinsToDistribute.Add(coinsToDistribute...)
	}
	return allCoinsToDistribute
}

func (k Keeper) prepareCoinToDistributeForMainAccount(ctx sdk.Context, states []types.State, subDistributorName string) sdk.DecCoins {
	coinsToDistribute := sdk.NewDecCoinsFromCoins(k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)...)
	if len(coinsToDistribute) > 0 {
		sum := getRamainsSum(&states)
		coinsToDistribute = coinsToDistribute.Sub(sum)
	}
	k.Logger(ctx).Debug("prepare coins to distribute for main account", "subDistr", subDistributorName, "coins", coinsToDistribute.String())

	return coinsToDistribute
}

func (k Keeper) prepareCoinToDistributeForNotMainAccount(ctx sdk.Context, source types.Account, states []types.State, subDistributorName string) sdk.DecCoins {
	var coinsToDistribute sdk.DecCoins
	if types.ModuleAccount == source.Type {
		coinsToDistribute = k.prepareCoinToDistributeForModuleAccount(ctx, source, subDistributorName)
	} else if types.InternalAccount != source.Type {
		coinsToDistribute = k.prepareCoinToDistributeForBaseAccount(ctx, source, subDistributorName)
	} else {
		coinsToDistribute = sdk.NewDecCoins()

	}
	return prepareLeftCoinToDistribute(coinsToDistribute, source, states)
}

func (k Keeper) prepareCoinToDistributeForModuleAccount(ctx sdk.Context, source types.Account, subDistributorName string) sdk.DecCoins {
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

func (k Keeper) prepareCoinToDistributeForBaseAccount(ctx sdk.Context, source types.Account, subDistributorName string) sdk.DecCoins {
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
		coin := states[pos].Remains
		if !coin.IsZero() {
			states[pos].Remains = sdk.NewDecCoins()
			coinsToDistribute = coinsToDistribute.Add(coin...)
		}
	}

	return coinsToDistribute
}

func (k Keeper) burnCoins(ctx sdk.Context, state *types.State) {
	toSend, change := state.Remains.TruncateDecimal()
	if err := k.BurnCoinsForSpecifiedModuleAccount(ctx, toSend, types.DistributorMainAccount); err != nil {
		ctx.Logger().Error("burn coins error", "state", state, "error", err.Error())
	} else {
		k.Logger(ctx).Debug("Coins burned", "coins", toSend)
		defer func() {
			sendCoinsSetGuage(types.BurnDestination, toSend)
		}()
		state.Remains = change
	}
}

func (k Keeper) sendCoinsToModuleAccount(ctx sdk.Context, state *types.State) {
	toSend, change := state.Remains.TruncateDecimal()
	if err := k.SendCoinsFromModuleToModule(ctx, toSend, types.DistributorMainAccount, state.Account.Id); err != nil {
		ctx.Logger().Error("send coins to module account dst error", "accountId", state.Account.Id, "error", err.Error())
	} else {
		k.Logger(ctx).Debug("coins sent to module account dst", "accountId", state.Account.Id, "toSend", toSend.String())
		defer func() {
			sendCoinsSetGuage(state.Account.Id, toSend)
		}()
		state.Remains = change
	}
}

func (k Keeper) sendCoinsToBaseAccount(ctx sdk.Context, state *types.State) {
	toSend, change := state.Remains.TruncateDecimal()
	if dstAccount, err := sdk.AccAddressFromBech32(state.Account.Id); err != nil {
		k.Logger(ctx).Error("destination base account address parsing error", "accountId", state.Account.Id, "error", err.Error())
	} else if err := k.SendCoinsFromModuleAccount(ctx, toSend, types.DistributorMainAccount, dstAccount); err != nil {
		k.Logger(ctx).Error("send coins to base account dst error", "accountId", state.Account.Id, "toSend", toSend, "error", err.Error())
	} else {
		k.Logger(ctx).Debug("coins sent to base account dst", "accountId", state.Account.Id, "toSend", toSend)
		defer func() {
			sendCoinsSetGuage(state.Account.Id, toSend)
		}()
		state.Remains = change
	}
}

func sendCoinsSetGuage(accountName string, coins sdk.Coins) {
	for _, coin := range coins {
		telemetry.SetGaugeWithLabels(
			[]string{types.ModuleName, "coin_send", coin.Denom, accountName},
			float32(coin.Amount.Int64()),
			[]metrics.Label{telemetry.NewLabel("denom", coin.Denom)},
		)
	}
}

func (k Keeper) SendCoinsFromStates(ctx sdk.Context, states []types.State) {
	for _, state := range states {
		if types.InternalAccount != state.Account.Type && checkIfAnyCoinIsGTE1(state.Remains) {
			if state.Burn {
				k.burnCoins(ctx, &state)
			} else if types.ModuleAccount == state.Account.Type {
				k.sendCoinsToModuleAccount(ctx, &state)
			} else {
				k.sendCoinsToBaseAccount(ctx, &state)
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

func (k Keeper) addSharesToBurnState(ctx sdk.Context, localRemains *[]types.State, calculatedShare sdk.DecCoins, findState func() int) *[]types.State {
	return k.addSharesToState(ctx, localRemains, true, nil, calculatedShare, findState)
}

func (k Keeper) addSharesToAccountState(ctx sdk.Context, localRemains *[]types.State, account *types.Account, calculatedShare sdk.DecCoins, findState func() int) *[]types.State {
	return k.addSharesToState(ctx, localRemains, false, account, calculatedShare, findState)
}

func (k Keeper) addSharesToState(ctx sdk.Context, localRemains *[]types.State, burn bool, account *types.Account, calculatedShare sdk.DecCoins, findState func() int) *[]types.State {
	pos := findState()
	logger := k.Logger(ctx).With("localRemains", localRemains, "account", account, "burn", burn,
		"calculatedShare", calculatedShare.String(), "pos", pos)
	if pos < 0 {
		var state types.State
		if burn || account == nil {
			state = types.State{Account: &types.Account{}, Remains: sdk.NewDecCoins(), Burn: true}
		} else {
			state = types.State{Account: account, Remains: sdk.NewDecCoins(), Burn: false}
		}
		withAppended := append(*localRemains, state)

		localRemains = &withAppended
		pos = len(*localRemains) - 1
		logger.Debug("add shares to state", "state", state)
	} else {
		logger.Debug("add shares to state")
	}
	(*localRemains)[pos].Remains = (*localRemains)[pos].Remains.Add(calculatedShare...)
	return localRemains
}

func (k Keeper) StartDistributionProcess(ctx sdk.Context, states *[]types.State, coinsToDistributeDec sdk.DecCoins, subDistributor types.SubDistributor) (localRemains *[]types.State, distributions []*types.Distribution, burn *types.DistributionBurn) {
	k.Logger(ctx).Debug("start distribution process", "subDistributor", subDistributor.String(),
		"coinsToDistributeDec", coinsToDistributeDec.String())
	localRemains = states
	defaultShare := coinsToDistributeDec
	for _, share := range subDistributor.Destinations.Shares {
		if share.Destination.Type == types.Main {
			continue
		}
		calculatedShare := calculatePercentage(share.Share, coinsToDistributeDec)
		defaultShare = defaultShare.Sub(calculatedShare)
		if !calculatedShare.IsZero() {
			findFunc := func() int {
				return findAccountState(localRemains, &share.Destination)
			}

			localRemains = k.addSharesToAccountState(ctx, localRemains, &share.Destination, calculatedShare, findFunc)
			distributions = append(distributions, &types.Distribution{
				Subdistributor: subDistributor.Name,
				ShareName:      share.Name,
				Sources:        subDistributor.Sources,
				Destination:    &share.Destination,
				Amount:         calculatedShare,
			})
		}
	}

	if subDistributor.Destinations.BurnShare != sdk.ZeroDec() {
		calculatedShare := calculatePercentage(subDistributor.Destinations.BurnShare, coinsToDistributeDec)
		defaultShare = defaultShare.Sub(calculatedShare)
		if !calculatedShare.IsZero() {
			findFunc := func() int {
				return findBurnState(localRemains)
			}
			localRemains = k.addSharesToBurnState(ctx, localRemains, calculatedShare, findFunc)
			burn = &types.DistributionBurn{
				Subdistributor: subDistributor.Name,
				Sources:        subDistributor.Sources,
				Amount:         calculatedShare,
			}
		}
	}

	accountDefault := subDistributor.Destinations.GetPrimaryShare()

	if accountDefault.Type != types.Main {
		findFunc := func() int {
			return findAccountState(localRemains, &accountDefault)
		}
		localRemains = k.addSharesToAccountState(ctx, localRemains, &accountDefault, defaultShare, findFunc)
		distributions = append(distributions, &types.Distribution{
			Subdistributor: subDistributor.Name,
			ShareName:      subDistributor.GetPrimaryShareName(),
			Sources:        subDistributor.Sources,
			Destination:    &subDistributor.Destinations.PrimaryShare,
			Amount:         defaultShare,
		})
	}
	k.Logger(ctx).Debug("start distribution process ret", "subDistributor", subDistributor.String(), "localRemains", localRemains)
	return
}
