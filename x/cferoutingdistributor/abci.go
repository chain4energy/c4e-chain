package cferoutingdistributor

import (
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func sendCoinToProperAccount(ctx sdk.Context, k keeper.Keeper, destinationAddress string,
	isModuleAccount bool, coinsToTransfer sdk.Int, source string) {

	if isModuleAccount {
		k.SendCoinsFromModuleToModule(ctx,
			sdk.NewCoins(sdk.NewCoin("uc4e", coinsToTransfer)), source, destinationAddress)
	} else {
		destinationAccount, _ := sdk.AccAddressFromBech32(destinationAddress)
		k.SendCoinsFromModuleAccount(ctx,
			sdk.NewCoins(sdk.NewCoin("uc4e", coinsToTransfer)), source, destinationAccount)
	}
	telemetry.IncrCounter(float32(coinsToTransfer.Int64()), destinationAddress+"-counter")

}

func saveRemainsToMap(ctx sdk.Context, k keeper.Keeper, destinationAddress string, remainsCount sdk.Dec) {
	k.GetRemains(ctx, destinationAddress)
	remains, _ := k.GetRemains(ctx, destinationAddress)
	remains.LeftoverCoin.Amount = remains.LeftoverCoin.Amount.Add(remainsCount)
	k.SetRemains(ctx, remains)
}

func createBurnRemainsIfNotExist(ctx sdk.Context, k keeper.Keeper) {
	account := types.Account{
		Address:         "burn",
		IsModuleAccount: false,
	}
	createRemainsIfNotExist(ctx, k, account)
}

func createRemainsIfNotExist(ctx sdk.Context, k keeper.Keeper, account types.Account) {

	_, isFound := k.GetRemains(ctx, account.Address)
	if !isFound {
		remains := types.Remains{
			Account:      account,
			LeftoverCoin: sdk.NewDecCoin("uc4e", sdk.ZeroInt()),
		}
		k.SetRemains(ctx, remains)
	}
}

func calculatePercentage(sharePercent sdk.Dec, coinsToDistributeDec sdk.Dec) sdk.Dec {
	if !coinsToDistributeDec.IsPositive() {
		return sdk.ZeroDec()
	}

	return coinsToDistributeDec.Mul(sharePercent).Quo(sdk.MustNewDecFromStr("100"))
}

func burnCoinForModuleAccount(ctx sdk.Context, k keeper.Keeper, coinsToBurn sdk.Int, sourceModule string) {
	k.BurnCoinsForSpecifiedModuleAccount(ctx, sdk.NewCoins(sdk.NewCoin("uc4e", coinsToBurn)), sourceModule)
	telemetry.IncrCounter(float32(coinsToBurn.Int64()), "burn-counter")
}

func findBurnState(states *[]types.Remains) int {
	for pos, state := range *states {
		if state.Account.Address == "burn" {
			return pos
		}
	}
	return -1
}

func findAccountState(states *[]types.Remains, account *types.Account) int {
	for pos, state := range *states {
		if state.Account.Address == account.Address {
			if account.IsInternalAccount {
				if state.Account.IsInternalAccount {
					return pos
				}
			} else if account.IsModuleAccount {
				if !state.Account.IsInternalAccount && state.Account.IsModuleAccount {
					return pos
				}
			} else {
				if !state.Account.IsInternalAccount && !state.Account.IsModuleAccount {
					return pos
				}
			}
		}
	}
	return -1
}

func getRamainsSum(states *[]types.Remains) sdk.DecCoins {
	sum := sdk.NewDecCoins();
	for _, state := range *states {
		sum = sum.Add(state.LeftoverCoin)
	}
	return sum;
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	routingDistributor := k.GetParams(ctx).RoutingDistributor
	states := k.GetAllRemains(ctx)
	println("BeginBlocker")

	for _, subDistributor := range routingDistributor.SubDistributor {
		allCoinsToDistribute := sdk.NewDecCoins()
		for _, source := range subDistributor.Sources {
			var coinsToDistribute = sdk.NewDecCoins()
			if source.IsMainCollector {
				coinsToDistribute = sdk.NewDecCoinsFromCoins(k.GetAccountCoinsForModuleAccount(ctx, types.CollectorName)...)
				println("IsMainCollector: " + coinsToDistribute.String())
				if len(coinsToDistribute) > 0 {
					sum := getRamainsSum(&states);
					coinsToDistribute = coinsToDistribute.Sub(sum)
				}
			} else if source.IsInternalAccount {
				pos := findAccountState(&states, source)
				if (pos >= 0) {
					coin := states[pos].LeftoverCoin
					println("IsInternalAccount: " + coin.String())
					if !coin.Amount.IsZero() {
						states[pos].LeftoverCoin.Amount = sdk.ZeroDec()
						coinsToDistribute = coinsToDistribute.Add(coin)
					}
					
				}
			} else if source.IsModuleAccount {
				coinsToSend := k.GetAccountCoinsForModuleAccount(ctx, source.Address)
				coinsToDistribute = sdk.NewDecCoinsFromCoins(coinsToSend...)
				println("IsModuleAccount: " + coinsToDistribute.String())

				if len(coinsToDistribute) > 0 {
					k.SendCoinsFromModuleToModule(ctx, coinsToSend, source.Address, types.CollectorName)
				}

				pos := findAccountState(&states, source)
				if (pos >= 0) {
					// state := states[pos];
					coin := states[pos].LeftoverCoin
					println("IsModuleAccount as internal: " + coin.String())
					if !coin.Amount.IsZero() {
						states[pos].LeftoverCoin.Amount = sdk.ZeroDec()
						coinsToDistribute = coinsToDistribute.Add(coin)
					}
					
				}
			} else {
				srcAccount, _ := sdk.AccAddressFromBech32(source.Address)
				coinsToSend := k.GetAccountCoins(ctx, srcAccount)
				coinsToDistribute = sdk.NewDecCoinsFromCoins(coinsToSend...)
				println("BaseAccount: " + coinsToDistribute.String())

				if len(coinsToDistribute) > 0 {
					k.SendCoinsToModuleAccount(ctx, coinsToSend, srcAccount, types.CollectorName)
				}
			}
			if len(coinsToDistribute) == 0 {
				continue
			}
			println("coinsToDistribute: " + coinsToDistribute.String())
			allCoinsToDistribute = allCoinsToDistribute.Add(coinsToDistribute...)
		}
		coinsToDistributeDec := allCoinsToDistribute.AmountOf("uc4e")
		println("coinsToDistributeDec: " + allCoinsToDistribute.String())
		if coinsToDistributeDec.IsZero() {
			break
		}
		states = *StartDistributionProcess(&states, coinsToDistributeDec, subDistributor)

	}
	println("BeginBlocker states: " + strconv.Itoa(len(states)))
	for _, state := range states {
		if state.Account.Address == "burn" {
			println("burn: " + state.LeftoverCoin.String())
			toSend, change := state.LeftoverCoin.TruncateDecimal()
			println("burn truncated: " + toSend.String())
			state.LeftoverCoin = change
			k.BurnCoinsForSpecifiedModuleAccount(ctx, sdk.NewCoins(toSend), types.CollectorName)
		} else if !state.Account.IsInternalAccount && state.Account.IsModuleAccount {
			println(state.Account.Address + " module acc: " + state.LeftoverCoin.String())
			toSend, change := state.LeftoverCoin.TruncateDecimal()
			println(state.Account.Address + " module acc truncated: " + toSend.String())
			state.LeftoverCoin = change
			k.SendCoinsFromModuleToModule(ctx, sdk.NewCoins(toSend), types.CollectorName, state.Account.Address)
		} else if !state.Account.IsInternalAccount && !state.Account.IsModuleAccount {
			println(state.Account.Address + " acc: " + state.LeftoverCoin.String())
			toSend, change := state.LeftoverCoin.TruncateDecimal()
			println(state.Account.Address + " acc truncated: " + toSend.String())
			state.LeftoverCoin = change
			dstAccount, _ := sdk.AccAddressFromBech32(state.Account.Address)
			k.SendCoinsFromModuleAccount(ctx, sdk.NewCoins(toSend), types.CollectorName, dstAccount)
		}
		println("remains: " + state.Account.Address)
		println("remains amount: " + state.LeftoverCoin.Amount.String())

		k.SetRemains(ctx, state)
	}

}

func StartDistributionProcess(states *[]types.Remains, coinsToDistributeDec sdk.Dec, subDistributor types.SubDistributor) *[]types.Remains {
	println("StartDistributionProcess")

	left := coinsToDistributeDec
	localRemains := states
	for _, share := range subDistributor.Destination.Share {
		if share.Account.IsMainCollector {
			continue;
		}
		calculatedShare := calculatePercentage(share.Percent, coinsToDistributeDec)
		println("calculatedShare: " + calculatedShare.String())

		if calculatedShare.GT(coinsToDistributeDec) {
			calculatedShare = coinsToDistributeDec
		}
		if !calculatedShare.IsZero() {
			left = left.Sub(calculatedShare)
			pos := findAccountState(localRemains, &share.Account)
			if pos < 0 {
				remains := types.Remains{Account: share.Account, LeftoverCoin: sdk.NewDecCoin("uc4e", sdk.ZeroInt())}
				withAppended := append(*localRemains, remains)
				println("remains append: " + share.Account.Address)

				localRemains = &withAppended
				pos = len(*localRemains) -  1;
			}
			(*localRemains)[pos].LeftoverCoin.Amount = (*localRemains)[pos].LeftoverCoin.Amount.Add(calculatedShare)
			println("state: " + (*localRemains)[pos].Account.Address)
			println("state: " + (*localRemains)[pos].LeftoverCoin.String())
		}
	}

	if subDistributor.Destination.BurnShare.Percent != sdk.MustNewDecFromStr("0") {
		calculatedShare := calculatePercentage(subDistributor.Destination.BurnShare.Percent, coinsToDistributeDec)
		if calculatedShare.GT(coinsToDistributeDec) {
			calculatedShare = coinsToDistributeDec
		}
		if !calculatedShare.IsZero() {
			left = left.Sub(calculatedShare)
			pos := findBurnState(localRemains)
			if pos < 0 {
				remains := types.Remains{Account: types.Account{Address: "burn"}, LeftoverCoin: sdk.NewDecCoin("uc4e", sdk.ZeroInt())}
				withAppended := append(*localRemains, remains)
				println("remains append: burn")
				localRemains = &withAppended
				pos = len(*localRemains) -  1;
			}
			(*localRemains)[pos].LeftoverCoin.Amount = (*localRemains)[pos].LeftoverCoin.Amount.Add(calculatedShare)
			println("burn state: " + (*localRemains)[pos].LeftoverCoin.String())
		}
	}

	accountDefault := subDistributor.Destination.GetAccount()
	if !accountDefault.IsMainCollector {
		pos := findAccountState(localRemains, &accountDefault)
		if pos < 0 {
			remains := types.Remains{Account: accountDefault, LeftoverCoin: sdk.NewDecCoin("uc4e", sdk.ZeroInt())}
			withAppended := append(*localRemains, remains)
			println("remains append: " + accountDefault.Address)
			localRemains = &withAppended
			pos = len(*localRemains) -  1;
		}
		(*localRemains)[pos].LeftoverCoin.Amount = (*localRemains)[pos].LeftoverCoin.Amount.Add(left)
		println("state: " + (*localRemains)[pos].Account.Address)
		println("state: " + (*localRemains)[pos].LeftoverCoin.String())
	}

	return localRemains
}
