package cferoutingdistributor

import (
	"time"

	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
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

		//if Internal == findAccountType(*account) && state.Account.InternalName == account.InternalName {
		//	return pos
		//} else if Module == findAccountType(*account) && state.Account.ModuleName == account.ModuleName {
		//	return pos
		//}
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

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	subDistributors := k.GetParams(ctx).SubDistributors
	states := k.GetALlStates(ctx)
	k.Logger(ctx).Info("BeginBlock - cfedistr")
	for _, subDistributor := range subDistributors {
		k.Logger(ctx).Info("BeginBlock - cfedistr: " + subDistributor.Name)
		allCoinsToDistribute := sdk.NewDecCoins()
		for _, source := range subDistributor.Sources {
			k.Logger(ctx).Debug("Sources: " + source.String())

			var coinsToDistribute = sdk.NewDecCoins()
			if source.Type == types.MAIN {
				coinsToDistribute = sdk.NewDecCoinsFromCoins(k.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)...)
				k.Logger(ctx).Debug("IsMainCollector: " + coinsToDistribute.String())
				if len(coinsToDistribute) > 0 {
					sum := getRamainsSum(&states)
					coinsToDistribute = coinsToDistribute.Sub(sum)
				}
			} else {

				if types.MODULE_ACCOUNT == source.Type {
					k.Logger(ctx).Debug("Module account: " + source.Id)
					coinsToSend := k.GetAccountCoinsForModuleAccount(ctx, source.Id)
					coinsToDistribute = sdk.NewDecCoinsFromCoins(coinsToSend...)
					k.Logger(ctx).Debug("IsModuleAccount: " + source.Id + " - " + coinsToDistribute.String())

					if len(coinsToDistribute) > 0 {
						k.SendCoinsFromModuleToModule(ctx, coinsToSend, source.Id, types.DistributorMainAccount)
					}
				} else if types.INTERNAL_ACCOUNT != source.Type {
					k.Logger(ctx).Debug("Internal account: " + source.Id)

					srcAccount, _ := sdk.AccAddressFromBech32(source.Id)
					coinsToSend := k.GetAccountCoins(ctx, srcAccount)
					coinsToDistribute = sdk.NewDecCoinsFromCoins(coinsToSend...)
					k.Logger(ctx).Debug("BaseAccount: " + source.Id + " - " + coinsToDistribute.String())

					if len(coinsToDistribute) > 0 {
						k.SendCoinsToModuleAccount(ctx, coinsToSend, srcAccount, types.DistributorMainAccount)
					}
				}

				pos := findAccountState(&states, source)
				if pos >= 0 {
					coin := states[pos].CoinsStates
					if !coin.IsZero() {
						states[pos].CoinsStates = sdk.NewDecCoins()
						coinsToDistribute = coinsToDistribute.Add(coin...)
					}

				}
			}
			if len(coinsToDistribute) == 0 {
				continue
			}
			allCoinsToDistribute = allCoinsToDistribute.Add(coinsToDistribute...)
		}
		//coinsToDistributeDec := allCoinsToDistribute.AmountOf("uc4e")
		if allCoinsToDistribute.IsZero() {
			continue
		}
		states = *StartDistributionProcess(&states, allCoinsToDistribute, subDistributor)

	}
	for _, state := range states {
		if types.INTERNAL_ACCOUNT != state.Account.Type {
			toSend, change := state.CoinsStates.TruncateDecimal()

			if state.Burn {
				if error := k.BurnCoinsForSpecifiedModuleAccount(ctx, toSend, types.DistributorMainAccount); error != nil {
					ctx.Logger().Error("Can not burn coin: " + error.Error())

				} else {
					k.Logger(ctx).Debug("Successful burn coin: " + toSend.String())
					state.CoinsStates = change
				}

			} else if types.MODULE_ACCOUNT == state.Account.Type {
				if error := k.SendCoinsFromModuleToModule(ctx, toSend, types.DistributorMainAccount, state.Account.Id); error != nil {
					ctx.Logger().Error("Can not send coin: " + error.Error())

				} else {
					k.Logger(ctx).Debug("Successful send to: " + state.Account.Id + " - " + toSend.String())
					state.CoinsStates = change
				}

			} else {
				if dstAccount, error := sdk.AccAddressFromBech32(state.Account.Id); error != nil {
					ctx.Logger().Error("Can not get addr from bech32: " + error.Error())

				} else if error := k.SendCoinsFromModuleAccount(ctx, toSend, types.DistributorMainAccount, dstAccount); error != nil {
					ctx.Logger().Error("Can not send coin: " + error.Error())

				} else {
					k.Logger(ctx).Debug("Successful send to : " + state.Account.Id + " - " + toSend.String())
					state.CoinsStates = change
				}
			}
		}

		k.SetState(ctx, state)
	}

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

func StartDistributionProcess(states *[]types.State, coinsToDistributeDec sdk.DecCoins, subDistributor types.SubDistributor) *[]types.State {
	percentShareSum := sdk.MustNewDecFromStr("0")
	//left := coinsToDistributeDec
	localRemains := states
	for _, share := range subDistributor.Destination.Share {
		percentShareSum = percentShareSum.Add(share.Percent)
		if share.Account.Type == types.MAIN {
			continue
		}
		calculatedShare := calculatePercentage(share.Percent, coinsToDistributeDec)
		//if calculatedShare.GT(coinsToDistributeDec) {
		//	calculatedShare = coinsToDistributeDec
		//}
		if !calculatedShare.IsZero() {
			//left = left.Sub(calculatedShare)
			findFunc := func() int {
				return findAccountState(localRemains, &share.Account)
			}
			localRemains = addSharesToState(localRemains, share.Account, calculatedShare, findFunc)
		}
	}

	if subDistributor.Destination.BurnShare.Percent != sdk.MustNewDecFromStr("0") {
		percentShareSum = percentShareSum.Add(subDistributor.Destination.BurnShare.Percent)
		calculatedShare := calculatePercentage(subDistributor.Destination.BurnShare.Percent, coinsToDistributeDec)

		//if calculatedShare.GT(coinsToDistributeDec) {
		//	calculatedShare = coinsToDistributeDec
		//}
		if !calculatedShare.IsZero() {
			//left = left.Sub(calculatedShare)

			findFunc := func() int {
				return findBurnState(localRemains)
			}
			localRemains = addSharesToState(localRemains, types.Account{}, calculatedShare, findFunc)
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
	}

	return localRemains
}
