package cferoutingdistributor

import (
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func calculatePercentage(sharePercent sdk.Dec, coinsToDistributeDec sdk.Dec) sdk.Dec {
	if !coinsToDistributeDec.IsPositive() {
		return sdk.ZeroDec()
	}

	return coinsToDistributeDec.Mul(sharePercent).Quo(sdk.MustNewDecFromStr("100"))
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
	sum := sdk.NewDecCoins()
	for _, state := range *states {
		sum = sum.Add(state.LeftoverCoin)
	}
	return sum
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	routingDistributor := k.GetParams(ctx).RoutingDistributor
	states := k.GetAllRemains(ctx)
	println("BeginBlocker")
	k.Logger(ctx).Info("BeginBlock - cfedistr")
	for _, subDistributor := range routingDistributor.SubDistributor {
		k.Logger(ctx).Info("BeginBlock - cfedistr: " + subDistributor.Name)
		allCoinsToDistribute := sdk.NewDecCoins()
		for _, source := range subDistributor.Sources {
			k.Logger(ctx).Debug("Sources: " +  source.Address)

			var coinsToDistribute = sdk.NewDecCoins()
			if source.IsMainCollector {
				coinsToDistribute = sdk.NewDecCoinsFromCoins(k.GetAccountCoinsForModuleAccount(ctx, types.CollectorName)...)
				println("IsMainCollector: " + coinsToDistribute.String())
				k.Logger(ctx).Debug("IsMainCollector: " + coinsToDistribute.String())
				if len(coinsToDistribute) > 0 {
					sum := getRamainsSum(&states)
					coinsToDistribute = coinsToDistribute.Sub(sum)
				}
			} else {
				if source.IsModuleAccount {
					coinsToSend := k.GetAccountCoinsForModuleAccount(ctx, source.Address)
					coinsToDistribute = sdk.NewDecCoinsFromCoins(coinsToSend...)
					println("IsModuleAccount: " + coinsToDistribute.String())
					k.Logger(ctx).Debug("IsModuleAccount: " + source.Address + " - " + coinsToDistribute.String())

					if len(coinsToDistribute) > 0 {
						k.SendCoinsFromModuleToModule(ctx, coinsToSend, source.Address, types.CollectorName)
					}
				} else if !source.IsInternalAccount {
					srcAccount, _ := sdk.AccAddressFromBech32(source.Address)
					coinsToSend := k.GetAccountCoins(ctx, srcAccount)
					coinsToDistribute = sdk.NewDecCoinsFromCoins(coinsToSend...)
					println("BaseAccount: " + coinsToDistribute.String())
					k.Logger(ctx).Debug("BaseAccount: " + source.Address + " - " + coinsToDistribute.String())

					if len(coinsToDistribute) > 0 {
						k.SendCoinsToModuleAccount(ctx, coinsToSend, srcAccount, types.CollectorName)
					}
				}

				pos := findAccountState(&states, source)
				if pos >= 0 {
					coin := states[pos].LeftoverCoin
					if !coin.Amount.IsZero() {
						states[pos].LeftoverCoin.Amount = sdk.ZeroDec()
						coinsToDistribute = coinsToDistribute.Add(coin)
					}

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
			continue
		}
		states = *StartDistributionProcess(&states, coinsToDistributeDec, subDistributor)

	}
	println("BeginBlocker states: " + strconv.Itoa(len(states)))
	for _, state := range states {

		if !state.Account.IsInternalAccount {
			println("burn: " + state.LeftoverCoin.String())
			toSend, change := state.LeftoverCoin.TruncateDecimal()
			println("burn truncated: " + toSend.String())
			state.LeftoverCoin = change

			if state.Account.Address == "burn" {
				k.Logger(ctx).Debug("Burn: " + toSend.String())

				k.BurnCoinsForSpecifiedModuleAccount(ctx, sdk.NewCoins(toSend), types.CollectorName)
			} else if state.Account.IsModuleAccount {
				k.Logger(ctx).Debug("Send to : " + state.Account.Address + " - " + toSend.String())

				k.SendCoinsFromModuleToModule(ctx, sdk.NewCoins(toSend), types.CollectorName, state.Account.Address)
			} else {
				k.Logger(ctx).Debug("Send to : " + state.Account.Address + " - " + toSend.String())
				dstAccount, _ := sdk.AccAddressFromBech32(state.Account.Address)
				k.SendCoinsFromModuleAccount(ctx, sdk.NewCoins(toSend), types.CollectorName, dstAccount)
			}
		}
		println("remains: " + state.Account.Address)
		println("remains amount: " + state.LeftoverCoin.Amount.String())

		k.SetRemains(ctx, state)
	}

}

func addSharesToState(localRemains *[]types.Remains, account types.Account, calculatedShare sdk.Dec, findState func() int) *[]types.Remains {
	pos := findState()
	if pos < 0 {
		remains := types.Remains{Account: account, LeftoverCoin: sdk.NewDecCoin("uc4e", sdk.ZeroInt())}
		withAppended := append(*localRemains, remains)
		println("remains append: " + account.Address)

		localRemains = &withAppended
		pos = len(*localRemains) - 1
	}
	(*localRemains)[pos].LeftoverCoin.Amount = (*localRemains)[pos].LeftoverCoin.Amount.Add(calculatedShare)
	return localRemains
}

func StartDistributionProcess(states *[]types.Remains, coinsToDistributeDec sdk.Dec, subDistributor types.SubDistributor) *[]types.Remains {
	println("StartDistributionProcess")

	left := coinsToDistributeDec
	localRemains := states
	for _, share := range subDistributor.Destination.Share {
		if share.Account.IsMainCollector {
			continue
		}
		calculatedShare := calculatePercentage(share.Percent, coinsToDistributeDec)
		println("calculatedShare: " + calculatedShare.String())

		if calculatedShare.GT(coinsToDistributeDec) {
			calculatedShare = coinsToDistributeDec
		}
		if !calculatedShare.IsZero() {
			left = left.Sub(calculatedShare)
			findFunc := func() int {
				return findAccountState(localRemains, &share.Account)
			}
			localRemains = addSharesToState(localRemains, share.Account, calculatedShare, findFunc)
		}
	}

	if subDistributor.Destination.BurnShare.Percent != sdk.MustNewDecFromStr("0") {
		calculatedShare := calculatePercentage(subDistributor.Destination.BurnShare.Percent, coinsToDistributeDec)
		if calculatedShare.GT(coinsToDistributeDec) {
			calculatedShare = coinsToDistributeDec
		}
		if !calculatedShare.IsZero() {
			left = left.Sub(calculatedShare)

			findFunc := func() int {
				return findBurnState(localRemains)
			}
			localRemains = addSharesToState(localRemains, types.Account{Address: "burn"}, calculatedShare, findFunc)
		}
	}

	accountDefault := subDistributor.Destination.GetAccount()
	if !accountDefault.IsMainCollector {
		findFunc := func() int {
			return findAccountState(localRemains, &accountDefault)
		}
		localRemains = addSharesToState(localRemains, accountDefault, left, findFunc)
	}

	return localRemains
}
