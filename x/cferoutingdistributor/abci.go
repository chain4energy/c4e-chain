package cferoutingdistributor

import (
	"github.com/chain4energy/c4e-chain/x/cferoutingdistributor/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sort"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {

	routingDistributor := k.GetRoutingDistributorr(ctx)

	sort.SliceStable(routingDistributor.SubDistributor, func(i, j int) bool {
		return routingDistributor.SubDistributor[i].Order < routingDistributor.SubDistributor[j].Order
	})

	for _, subDistributors := range routingDistributor.SubDistributor {

		for _, source := range subDistributors.Sources {

			coinsToDistribute := k.GetAccountCoinsForModuleAccount(ctx, source)
			k.Logger(ctx).Info(subDistributors.Name + " The amount of coins to be distributed: " + coinsToDistribute.String())

			for _, share := range subDistributors.Destination.Share {
				dividedCoins := coinsToDistribute.AmountOf("uc4e").Mul(sdk.NewInt(share.Percent)).Quo(sdk.NewInt(100))
				k.Logger(ctx).Info(subDistributors.Name + " amount of coins transferred : " + dividedCoins.String() + " to " + share.Account.Address)

				if share.Account.IsModuleAccount {
					k.SendCoinsFromModuleToModule(ctx, sdk.NewCoins(sdk.NewCoin("uc4e", dividedCoins)), source, share.Account.Address)
				} else {

					destinationAccount, _ := sdk.AccAddressFromBech32(share.Account.Address)
					k.SendCoinsFromModuleAccount(ctx, sdk.NewCoins(sdk.NewCoin("uc4e", dividedCoins)), source, destinationAccount)
				}
			}

			coinsLeftToDistributeToDeafultShare := k.GetAccountCoinsForModuleAccount(ctx, source)
			k.Logger(ctx).Info(subDistributors.Name + " Send money to default share account name: " + subDistributors.Destination.DefaultShareAccount.Address + " amount: " + coinsLeftToDistributeToDeafultShare.String())
			if subDistributors.Destination.DefaultShareAccount.IsModuleAccount {
				k.SendCoinsFromModuleToModule(ctx, coinsLeftToDistributeToDeafultShare, source, subDistributors.Destination.DefaultShareAccount.Address)

			} else {
				// todo Handle this error later if address is empty only, but this should not happend
				addressToSend, _ := sdk.AccAddressFromBech32(subDistributors.Destination.DefaultShareAccount.Address)
				k.SendCoinsFromModuleAccount(ctx, coinsLeftToDistributeToDeafultShare, source, addressToSend)
			}
		}
	}
}
