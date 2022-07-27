package cferoutingdistributor

import (
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
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

func saveRemainsToMap(ctx sdk.Context, k keeper.Keeper, destinationAddress string, remainsCount sdk.Dec, routingDistributor *types.RoutingDistributor) {
	remains := routingDistributor.GetRemainsMap()[destinationAddress]
	remains.LeftoverCoin = remains.LeftoverCoin.Add(remainsCount)
	routingDistributor.GetRemainsMap()[destinationAddress] = remains
}

func createBurnRemainsIfNotExist(ctx sdk.Context, k keeper.Keeper, routingDistributor *types.RoutingDistributor) {
	account := types.Account{
		Address:         "burn",
		IsModuleAccount: false,
	}
	createRemainsIfNotExist(ctx, k, account, routingDistributor)
}

func createRemainsIfNotExist(ctx sdk.Context, k keeper.Keeper, account types.Account, routingDistributor *types.RoutingDistributor) {
	if routingDistributor.RemainsMap == nil {
		routingDistributor.RemainsMap = make(map[string]types.Remains, 10)
	}

	if _, ok := routingDistributor.RemainsMap[account.Address]; !ok {
		remains := types.Remains{
			Account:      account,
			LeftoverCoin: sdk.MustNewDecFromStr("0"),
		}
		routingDistributor.GetRemainsMap()[account.Address] = remains
	}
}

func calculateAndSendCoin(ctx sdk.Context, k keeper.Keeper, account types.Account, sharePercent sdk.Dec, coinsToDistributeDec sdk.Dec,
	distributorName string, sourceModuleAccount string, routingDistributor *types.RoutingDistributor) {
	if !coinsToDistributeDec.IsPositive() {
		return
	}

	dividedCoins := coinsToDistributeDec.Mul(sharePercent).QuoTruncate(sdk.MustNewDecFromStr("100"))
	coinsToTransfer := dividedCoins.TruncateInt()
	coinsLeftNoTransferred := dividedCoins.Sub(sdk.NewDecFromInt(coinsToTransfer))
	createRemainsIfNotExist(ctx, k, account, routingDistributor)
	saveRemainsToMap(ctx, k, account.Address, coinsLeftNoTransferred, routingDistributor)
	sendCoinToProperAccount(ctx, k, account.Address, account.IsModuleAccount, coinsToTransfer, sourceModuleAccount)
	k.Logger(ctx).Debug("Coin left no transferred: " + coinsLeftNoTransferred.String())
	k.Logger(ctx).Debug(distributorName + " amount of coins transferred : " + coinsToTransfer.String() + " to " + account.Address)
}

func calculateAndBurnCoin(ctx sdk.Context, k keeper.Keeper, coinsToDistributeDec sdk.Dec, share types.BurnShare, source string, routingDistributor *types.RoutingDistributor) {
	if !coinsToDistributeDec.IsPositive() {
		return
	}
	dividedCoins := coinsToDistributeDec.Mul(share.Percent).QuoTruncate(sdk.MustNewDecFromStr("100"))
	coinsToBurn := dividedCoins.TruncateInt()
	coinsLeftNoBurned := dividedCoins.Sub(sdk.NewDecFromInt(coinsToBurn))
	createBurnRemainsIfNotExist(ctx, k, routingDistributor)
	saveRemainsToMap(ctx, k, "burn", coinsLeftNoBurned, routingDistributor)
	burnCoinForModuleAccount(ctx, k, coinsToBurn, source)
}

func burnCoinForModuleAccount(ctx sdk.Context, k keeper.Keeper, coinsToBurn sdk.Int, sourceModule string) {
	k.BurnCoinsForSpecifiedModuleAccount(ctx, sdk.NewCoins(sdk.NewCoin("uc4e", coinsToBurn)), sourceModule)
	telemetry.IncrCounter(float32(coinsToBurn.Int64()), "burn-counter")
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	routingDistributor := k.GetRoutingDistributorr(ctx)

	//sort.SliceStable(routingDistributor.SubDistributor, func(i, j int) bool {
	//	return routingDistributor.SubDistributor[i].Order < routingDistributor.SubDistributor[j].Order
	//})

	for _, subDistributor := range routingDistributor.SubDistributor {

		for _, source := range subDistributor.Sources {
			percentShareSum := sdk.MustNewDecFromStr("0")

			coinsToDistribute := k.GetAccountCoinsForModuleAccount(ctx, source)
			coinsToDistributeDec := sdk.NewDecFromInt(coinsToDistribute.AmountOf("uc4e"))
			if !coinsToDistributeDec.IsPositive() {
				break
			}

			k.Logger(ctx).Info("Coin to distribute: " + coinsToDistribute.String() + " from source distributor name: " + subDistributor.Name)

			for _, share := range subDistributor.Destination.Share {
				percentShareSum = percentShareSum.Add(share.Percent)
				calculateAndSendCoin(ctx, k, share.Account, share.Percent, coinsToDistributeDec,
					subDistributor.Name, source, &routingDistributor)
			}

			if subDistributor.Destination.BurnShare.Percent != sdk.MustNewDecFromStr("0") {
				percentShareSum = percentShareSum.Add(subDistributor.Destination.BurnShare.Percent)
				calculateAndBurnCoin(ctx, k, coinsToDistributeDec, subDistributor.Destination.BurnShare, source, &routingDistributor)
			}

			defaultSharePercent := sdk.MustNewDecFromStr("100").Sub(percentShareSum)
			accountDefault := subDistributor.Destination.GetAccount()
			calculateAndSendCoin(ctx, k, accountDefault, defaultSharePercent, coinsToDistributeDec, subDistributor.Name, source, &routingDistributor)

			coinsLeftToTransferToRemainsAccount := k.GetAccountCoinsForModuleAccount(ctx, source)
			k.Logger(ctx).Debug("Send coin to remains account name: " + routingDistributor.RemainsCoinModuleAccount + " count " + coinsLeftToTransferToRemainsAccount.String())
			k.SendCoinsFromModuleToModule(ctx, coinsLeftToTransferToRemainsAccount, source, routingDistributor.RemainsCoinModuleAccount)
		}
	}

	sendRemains(ctx, k, &routingDistributor)
	k.SetRoutingDistributor(ctx, routingDistributor)

}

func sendRemains(ctx sdk.Context, k keeper.Keeper, routingDistributor *types.RoutingDistributor) {
	remainsMap := routingDistributor.RemainsMap
	source := k.GetRoutingDistributorr(ctx).RemainsCoinModuleAccount

	for key, remains := range remainsMap {
		//check if remains coin is greater then 1
		if remains.LeftoverCoin.Sub(sdk.MustNewDecFromStr("1")).IsPositive() {

			account := remains.Account
			coinToTransferInt := remains.LeftoverCoin.TruncateInt()

			if remains.Account.Address == "burn" {
				burnCoinForModuleAccount(ctx, k, coinToTransferInt, source)
			} else {
				sendCoinToProperAccount(ctx, k, account.Address, account.IsModuleAccount, coinToTransferInt, source)
			}
			remains.LeftoverCoin = remains.LeftoverCoin.Sub(coinToTransferInt.ToDec())
			remainsMap[key] = remains
		}
	}
	routingDistributor.RemainsMap = remainsMap
}
