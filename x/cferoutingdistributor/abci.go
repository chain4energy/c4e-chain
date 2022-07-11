package cferoutingdistributor

import (
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sort"
)

func sendCoinToProperAccount(k keeper.Keeper, address string, isModuleAccount bool,
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

func calculateAndSendCoin(ctx sdk.Context, k keeper.Keeper, account types.Account, sharePercent sdk.Dec, coinsToDistributeDec sdk.Dec,
	distributorName string, sourceModuleAccount string) sdk.Dec {

	dividedCoins := coinsToDistributeDec.Mul(sharePercent).QuoTruncate(sdk.MustNewDecFromStr("100"))
	coinsToTransfer := dividedCoins.TruncateInt()
	coinsLeftNoTransferred := dividedCoins.Sub(sdk.NewDecFromInt(coinsToTransfer))
	account.LeftoverCoin = account.LeftoverCoin.Add(coinsLeftNoTransferred)
	k.Logger(ctx).Debug("Coin left no transferred: " + coinsLeftNoTransferred.String() + " " + account.LeftoverCoin.String())

	k.Logger(ctx).Debug(distributorName + " amount of coins transferred : " + coinsToTransfer.String() + " to " + account.Address)
	sendCoinToProperAccount(k, account.Address, account.IsModuleAccount, ctx, coinsToTransfer, sourceModuleAccount)
	return account.LeftoverCoin
}

func calculateAndBurnCoin(ctx sdk.Context, k keeper.Keeper, coinsToDistributeDec sdk.Dec, share types.BurnShare, source string) sdk.Dec {
	dividedCoins := coinsToDistributeDec.Mul(share.Percent).QuoTruncate(sdk.MustNewDecFromStr("100"))
	coinsToBurn := dividedCoins.TruncateInt()
	coinsLeftNoBurned := dividedCoins.Sub(sdk.NewDecFromInt(coinsToBurn))
	k.BurnCoinsForSpecifiedModuleAccount(ctx, sdk.NewCoins(sdk.NewCoin("uc4e", coinsToBurn)), source)
	return share.LeftoverCoin.Add(coinsLeftNoBurned)
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {

	routingDistributor := k.GetRoutingDistributorr(ctx)

	sort.SliceStable(routingDistributor.SubDistributor, func(i, j int) bool {
		return routingDistributor.SubDistributor[i].Order < routingDistributor.SubDistributor[j].Order
	})

	for i, subDistributor := range routingDistributor.SubDistributor {
		percentShareSum := sdk.MustNewDecFromStr("0")

		for _, source := range subDistributor.Sources {
			coinsToDistribute := k.GetAccountCoinsForModuleAccount(ctx, source)
			coinsToDistributeDec := sdk.NewDecFromInt(coinsToDistribute.AmountOf("uc4e"))
			k.Logger(ctx).Info("Coin to distribute: " + coinsToDistribute.String() + " from source distributor name: " + subDistributor.Name)

			for shareIndex, _ := range subDistributor.Destination.Share {
				share := routingDistributor.SubDistributor[i].Destination.Share[shareIndex]
				percentShareSum = percentShareSum.Add(share.Percent)

				routingDistributor.SubDistributor[i].Destination.Share[shareIndex].Account.LeftoverCoin =
					calculateAndSendCoin(ctx, k, share.Account, share.Percent, coinsToDistributeDec,
						subDistributor.Name, source)
			}

			if subDistributor.Destination.BurnShare.Percent != sdk.MustNewDecFromStr("0") {
				percentShareSum = percentShareSum.Add(subDistributor.Destination.BurnShare.Percent)
				calculateAndBurnCoin(ctx, k, coinsToDistributeDec, subDistributor.Destination.BurnShare, source)
			}

			defaultSharePercent := sdk.MustNewDecFromStr("100").Sub(percentShareSum)
			accountDefault := routingDistributor.SubDistributor[i].Destination.GetAccount()
			routingDistributor.SubDistributor[i].Destination.Account.LeftoverCoin =
				calculateAndSendCoin(ctx, k, accountDefault, defaultSharePercent, coinsToDistributeDec, subDistributor.Name, source)

			coinsLeftToTransferToRemainsAccount := k.GetAccountCoinsForModuleAccount(ctx, source)
			k.Logger(ctx).Debug("Send coin to remains account name: " + routingDistributor.RemainsCoinModuleAccount + " count " + coinsLeftToTransferToRemainsAccount.String())
			k.SendCoinsFromModuleToModule(ctx, coinsLeftToTransferToRemainsAccount, source, routingDistributor.RemainsCoinModuleAccount)
		}
	}

	k.SetRoutingDistributor(ctx, routingDistributor)

	//send remains
	for i, subDistributors := range routingDistributor.SubDistributor {
		if subDistributors.Destination.BurnShare.LeftoverCoin.Sub(sdk.MustNewDecFromStr("1")).IsPositive() {

			//burn from remains
			coinsToBurnInt := subDistributors.Destination.BurnShare.LeftoverCoin.TruncateInt()
			coinsToBurn := sdk.NewCoins(sdk.NewCoin("uc4e", coinsToBurnInt))
			k.BurnCoinsForSpecifiedModuleAccount(ctx, coinsToBurn, routingDistributor.RemainsCoinModuleAccount)
			routingDistributor.SubDistributor[i].Destination.BurnShare.LeftoverCoin =
				routingDistributor.SubDistributor[i].Destination.BurnShare.LeftoverCoin.Sub(coinsToBurnInt.ToDec())
		}

		if subDistributors.Destination.Account.LeftoverCoin.Sub(sdk.MustNewDecFromStr("1")).IsPositive() {
			coinToTransferInt := subDistributors.Destination.Account.LeftoverCoin.TruncateInt()
			sendCoinToProperAccount(k, subDistributors.Destination.Account.Address, subDistributors.Destination.Account.IsModuleAccount,
				ctx, coinToTransferInt, routingDistributor.RemainsCoinModuleAccount)
			routingDistributor.SubDistributor[i].Destination.Account.LeftoverCoin =
				routingDistributor.SubDistributor[i].Destination.Account.LeftoverCoin.Sub(coinToTransferInt.ToDec())
		}

		for y, share := range subDistributors.Destination.Share {
			if share.Account.LeftoverCoin.Sub(sdk.MustNewDecFromStr("1")).IsPositive() {
				coinToTransferInt := share.Account.LeftoverCoin.TruncateInt()
				sendCoinToProperAccount(k, share.Account.Address, share.Account.IsModuleAccount, ctx, coinToTransferInt, routingDistributor.RemainsCoinModuleAccount)
				routingDistributor.SubDistributor[i].Destination.Share[y].Account.LeftoverCoin =
					routingDistributor.SubDistributor[i].Destination.Share[y].Account.LeftoverCoin.Sub(coinToTransferInt.ToDec())
			}
		}

	}
	k.SetRoutingDistributor(ctx, routingDistributor)

}
