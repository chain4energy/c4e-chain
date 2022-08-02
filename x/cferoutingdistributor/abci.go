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

func saveRemainsToMap(ctx sdk.Context, k keeper.Keeper, destinationAddress string, remainsCount sdk.Dec) {
	k.GetRemains(ctx, destinationAddress)
	remains, _ := k.GetRemains(ctx, destinationAddress)
	remains.LeftoverCoin = remains.LeftoverCoin.Add(remainsCount)
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
			LeftoverCoin: sdk.MustNewDecFromStr("0"),
		}
		k.SetRemains(ctx, remains)
	}
}

func calculateAndSendCoin(ctx sdk.Context, k keeper.Keeper, account types.Account, sharePercent sdk.Dec, coinsToDistributeDec sdk.Dec,
	distributorName string, sourceModuleAccount string) {
	if !coinsToDistributeDec.IsPositive() {
		return
	}

	dividedCoins := coinsToDistributeDec.Mul(sharePercent).QuoTruncate(sdk.MustNewDecFromStr("100"))
	coinsToTransfer := dividedCoins.TruncateInt()
	coinsLeftNoTransferred := dividedCoins.Sub(sdk.NewDecFromInt(coinsToTransfer))
	createRemainsIfNotExist(ctx, k, account)
	saveRemainsToMap(ctx, k, account.Address, coinsLeftNoTransferred)
	sendCoinToProperAccount(ctx, k, account.Address, account.IsModuleAccount, coinsToTransfer, sourceModuleAccount)
}

func calculateAndBurnCoin(ctx sdk.Context, k keeper.Keeper, coinsToDistributeDec sdk.Dec, share types.BurnShare, source string) {
	if !coinsToDistributeDec.IsPositive() {
		return
	}
	dividedCoins := coinsToDistributeDec.Mul(share.Percent).QuoTruncate(sdk.MustNewDecFromStr("100"))
	coinsToBurn := dividedCoins.TruncateInt()
	coinsLeftNoBurned := dividedCoins.Sub(sdk.NewDecFromInt(coinsToBurn))
	createBurnRemainsIfNotExist(ctx, k)
	saveRemainsToMap(ctx, k, "burn", coinsLeftNoBurned)
	burnCoinForModuleAccount(ctx, k, coinsToBurn, source)
}

func burnCoinForModuleAccount(ctx sdk.Context, k keeper.Keeper, coinsToBurn sdk.Int, sourceModule string) {
	k.BurnCoinsForSpecifiedModuleAccount(ctx, sdk.NewCoins(sdk.NewCoin("uc4e", coinsToBurn)), sourceModule)
	telemetry.IncrCounter(float32(coinsToBurn.Int64()), "burn-counter")
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	routingDistributor := k.GetParams(ctx).RoutingDistributor

	for _, subDistributor := range routingDistributor.SubDistributor {
		for _, source := range subDistributor.Sources {
			coinsToDistribute := k.GetAccountCoinsForModuleAccount(ctx, source)
			coinsToDistributeInt := coinsToDistribute.AmountOf("uc4e")
			if coinsToDistributeInt.IsZero() {
				break
			}
			coinsToDistributeDec := sdk.NewDecFromInt(coinsToDistribute.AmountOf("uc4e"))
			StartDistributionProcess(ctx, k, source, routingDistributor.RemainsCoinModuleAccount, coinsToDistributeDec, subDistributor)
		}
	}

	sendRemains(ctx, k, routingDistributor.RemainsCoinModuleAccount)
}

func StartDistributionProcess(ctx sdk.Context, k keeper.Keeper, source string, remainsAccountName string, coinsToDistributeDec sdk.Dec, subDistributor types.SubDistributor) {
	percentShareSum := sdk.MustNewDecFromStr("0")
	for _, share := range subDistributor.Destination.Share {
		percentShareSum = percentShareSum.Add(share.Percent)
		calculateAndSendCoin(ctx, k, share.Account, share.Percent, coinsToDistributeDec,
			subDistributor.Name, source)
	}

	if subDistributor.Destination.BurnShare.Percent != sdk.MustNewDecFromStr("0") {
		percentShareSum = percentShareSum.Add(subDistributor.Destination.BurnShare.Percent)
		calculateAndBurnCoin(ctx, k, coinsToDistributeDec, subDistributor.Destination.BurnShare, source)
	}

	defaultSharePercent := sdk.MustNewDecFromStr("100").Sub(percentShareSum)
	accountDefault := subDistributor.Destination.GetAccount()
	calculateAndSendCoin(ctx, k, accountDefault, defaultSharePercent, coinsToDistributeDec, subDistributor.Name, source)

	coinsLeftToTransferToRemainsAccount := k.GetAccountCoinsForModuleAccount(ctx, source)
	k.SendCoinsFromModuleToModule(ctx, coinsLeftToTransferToRemainsAccount, source, remainsAccountName)
}

func sendRemains(ctx sdk.Context, k keeper.Keeper, remainsAccountSource string) {

	for _, remains := range k.GetAllRemains(ctx) {
		//check if remains coin is greater then 10
		if remains.LeftoverCoin.Sub(sdk.MustNewDecFromStr("1")).IsPositive() {

			account := remains.Account
			coinToTransferInt := remains.LeftoverCoin.TruncateInt()

			if remains.Account.Address == "burn" {
				burnCoinForModuleAccount(ctx, k, coinToTransferInt, remainsAccountSource)
			} else {
				sendCoinToProperAccount(ctx, k, account.Address, account.IsModuleAccount, coinToTransferInt, remainsAccountSource)
			}
			remains.LeftoverCoin = remains.LeftoverCoin.Sub(coinToTransferInt.ToDec())
			k.SetRemains(ctx, remains)
		}
	}
}
