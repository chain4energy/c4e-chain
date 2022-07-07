package cferoutingdistributor

import (
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sort"
)

func sendCoinToProperAdd(k keeper.Keeper, address string, isModuleAccount bool,
	ctx sdk.Context, coinsToTransfer sdk.Int, source string) {
	if isModuleAccount {
		k.SendCoinsFromModuleToModule(ctx,
			sdk.NewCoins(sdk.NewCoin("uc4e", coinsToTransfer)), source, address)
	} else {

		destinationAccount, _ := sdk.AccAddressFromBech32(address)
		k.SendCoinsFromModuleAccount(ctx,
			sdk.NewCoins(sdk.NewCoin("uc4e", coinsToTransfer)), source, destinationAccount)
	}
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {

	routingDistributor := k.GetRoutingDistributorr(ctx)

	sort.SliceStable(routingDistributor.SubDistributor, func(i, j int) bool {
		return routingDistributor.SubDistributor[i].Order < routingDistributor.SubDistributor[j].Order
	})

	for i, subDistributors := range routingDistributor.SubDistributor {

		for _, source := range subDistributors.Sources {
			k.Logger(ctx).Info(source)

			coinsToDistribute := k.GetAccountCoinsForModuleAccount(ctx, source)
			coinsToDistributeDec := sdk.NewDecFromInt(coinsToDistribute.AmountOf("uc4e"))

			percentShareSum := sdk.MustNewDecFromStr("0")

			//new sdk.NewDec(coinsToDistribute.AmountOf("uc4e"))
			k.Logger(ctx).Info(subDistributors.Name + " The amount of coins to be distributed: " + coinsToDistributeDec.String())

			for i, share := range subDistributors.Destination.Share {
				percentShareSum = percentShareSum.Add(share.Percent)
				dividedCoins := coinsToDistributeDec.Mul(share.Percent).QuoTruncate(sdk.MustNewDecFromStr("100"))
				coinsToTransfer := dividedCoins.TruncateInt()
				coinsLeftNoTransfered := dividedCoins.Sub(sdk.NewDecFromInt(coinsToTransfer))
				//share.Account.LeftoverCoin = share.Account.LeftoverCoin.Add(coinsLeftNoTransfered)
				subDistributors.Destination.Share[i].Account.LeftoverCoin = share.Account.LeftoverCoin.Add(coinsLeftNoTransfered)

				share.Account.LeftoverCoin = share.Account.LeftoverCoin.Add(coinsLeftNoTransfered)
				k.Logger(ctx).Debug("Coin left no transfered: " + coinsLeftNoTransfered.String())

				k.Logger(ctx).Debug(subDistributors.Name + " amount of coins transferred : " + coinsToTransfer.String() + " to " + share.Account.Address)
				sendCoinToProperAdd(k, share.Account.Address, share.Account.IsModuleAccount, ctx, coinsToTransfer, source)
			}

			if subDistributors.Destination.BurnShare != nil && subDistributors.Destination.BurnShare.Percent != sdk.MustNewDecFromStr("0") {
				percentShareSum = percentShareSum.Add(subDistributors.Destination.BurnShare.Percent)
				dividedCoins := coinsToDistributeDec.Mul(subDistributors.Destination.BurnShare.Percent).QuoTruncate(sdk.MustNewDecFromStr("100"))
				coinsToBurn := dividedCoins.TruncateInt()
				coinsLeftNoBurned := dividedCoins.Sub(sdk.NewDecFromInt(coinsToBurn))
				subDistributors.Destination.BurnShare.LeftoverCoin = subDistributors.Destination.BurnShare.LeftoverCoin.Add(coinsLeftNoBurned)
				k.BurnCoinsForSpecifiedModuleAccount(ctx, sdk.NewCoins(sdk.NewCoin("uc4e", coinsToBurn)), source)
			}

			defaultSharePercent := sdk.MustNewDecFromStr("100").Sub(percentShareSum)
			dividedCoinsDefault := coinsToDistributeDec.Mul(defaultSharePercent.QuoTruncate(sdk.MustNewDecFromStr("100")))
			coinsToTransfer := dividedCoinsDefault.TruncateInt()
			coinsLeftNoTransfered := dividedCoinsDefault.Sub(sdk.NewDecFromInt(coinsToTransfer))

			var account = subDistributors.Destination.Account
			sendCoinToProperAdd(k, account.Address, account.IsModuleAccount, ctx, coinsToTransfer, source)
			subDistributors.Destination.Account.LeftoverCoin = account.LeftoverCoin.Add(coinsLeftNoTransfered)

			k.Logger(ctx).Info(subDistributors.Name + " Send money to default share account name: " + account.Address + " amount: " + coinsToTransfer.String())

			coinsLeftToTransferToRemainsAccount := k.GetAccountCoinsForModuleAccount(ctx, source)
			k.Logger(ctx).Debug("Send coin to remains account name: " + routingDistributor.RemainsCoinModuleAccount + " count " + coinsLeftToTransferToRemainsAccount.String())
			k.SendCoinsFromModuleToModule(ctx, coinsLeftToTransferToRemainsAccount, source, routingDistributor.RemainsCoinModuleAccount)
		}

		routingDistributor.SubDistributor[i] = subDistributors
	}

	k.Logger(ctx).Info(routingDistributor.String())

	k.SetRoutingDistributor(ctx, routingDistributor)

	//send remains
	for i, subDistributors := range routingDistributor.SubDistributor {
		k.Logger(ctx).Info("ASDAS2")
		if subDistributors.Destination.BurnShare != nil && subDistributors.Destination.BurnShare.LeftoverCoin.Sub(sdk.MustNewDecFromStr("1")).IsPositive() {
			k.Logger(ctx).Info("ASDAS")

			//burn from remains
			coinsToBurnInt := subDistributors.Destination.BurnShare.LeftoverCoin.TruncateInt()
			coinsToBurn := sdk.NewCoins(sdk.NewCoin("uc4e", coinsToBurnInt))
			k.BurnCoinsForSpecifiedModuleAccount(ctx, coinsToBurn, routingDistributor.RemainsCoinModuleAccount)
			routingDistributor.SubDistributor[i].Destination.BurnShare.LeftoverCoin =
				routingDistributor.SubDistributor[i].Destination.BurnShare.LeftoverCoin.Sub(coinsToBurnInt.ToDec())
		}

		if subDistributors.Destination.Account.LeftoverCoin.Sub(sdk.MustNewDecFromStr("1")).IsPositive() {
			coinToTransferInt := subDistributors.Destination.Account.LeftoverCoin.TruncateInt()
			sendCoinToProperAdd(k, subDistributors.Destination.Account.Address, subDistributors.Destination.Account.IsModuleAccount,
				ctx, coinToTransferInt, routingDistributor.RemainsCoinModuleAccount)
			routingDistributor.SubDistributor[i].Destination.Account.LeftoverCoin =
				routingDistributor.SubDistributor[i].Destination.Account.LeftoverCoin.Sub(coinToTransferInt.ToDec())
		}

		for y, share := range subDistributors.Destination.Share {
			if share.Account.LeftoverCoin.Sub(sdk.MustNewDecFromStr("1")).IsPositive() {
				coinToTransferInt := share.Account.LeftoverCoin.TruncateInt()
				sendCoinToProperAdd(k, share.Account.Address, share.Account.IsModuleAccount, ctx, coinToTransferInt, routingDistributor.RemainsCoinModuleAccount)
				routingDistributor.SubDistributor[i].Destination.Share[y].Account.LeftoverCoin =
					routingDistributor.SubDistributor[i].Destination.Share[y].Account.LeftoverCoin.Sub(coinToTransferInt.ToDec())
			}
		}

	}
	k.SetRoutingDistributor(ctx, routingDistributor)

}
