package cferoutingdistributor_test

import (
	"testing"

	"github.com/chain4energy/c4e-chain/app"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestAbci(t *testing.T) {

	// Setup main app
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// app.CferoutingdistributorKeeper.SetRoutingDistributor(ctx) TODO

	//Setup minter params - TODO
	// minterNew := app.CfeminterKeeper.GetHalvingMinter(ctx)
	// minterNew.MintDenom = "uC4E"
	// minterNew.NewCoinsMint = 20596877
	// minterNew.BlocksPerYear = 4855105
	// app.CfeminterKeeper.SetHalvingMinter(ctx, minterNew)

	for i := 1; i < 100; i++ {
		ctx = ctx.WithBlockHeight(int64(i))
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
		app.EndBlocker(ctx, abci.RequestEndBlock{})
	}

	//app.BankKeeper. TODO
	// require.Equal(t, app.BankKeeper.GetSupply(ctx, "uC4E").Amount.String(), 20596877, "asd")

	// -------------------
	// types.RoutingDistributor{
	// 	SubDistributor: [
	// 	{Name:inflation_distributor Sources:[inflation_collector]
	// 	Destination:{DefaultShareAccount:address:"validators_rewards_collector" is_module_account:true
	// 	Share:[{Name:users_incentive_share Percent:30 Account:{Address:users_incentive_collector IsModuleAccount:true}}] BurnShare:0} Order:1}

	// 	{Name:fee_and_payment_distributor Sources:[fee_collector payment_collector] Destination:{DefaultShareAccount:address:"validators_rewards_collector" is_module_account:true  Share:[{Name:community_pool_rewards_share Percent:30 Account:{Address:community_pool_rewards_collector IsModuleAccount:true}}] BurnShare:0} Order:2} {Name:community_pool_rewards_distributor Sources:[community_pool_rewards_collector] Destination:{DefaultShareAccount:address:"c4e1wejevyydp409tz0necwfg4mzj8md4vfy9n95xu"  Share:[{Name:liquidity_and_gov_rewards_share Percent:30 Account:{Address:c4e132g4u3qzf890cqaz9yhaegc6v45ew7qzmzlywg IsModuleAccount:false11:03}} {Name:strategic_reserve_share Percent:30 Account:{Address:c4e1avc7vz3khvlf6fgd3a2exnaqnhhk0sxzzgxc4n IsModuleAccount:false}}] BurnShare:0} Order:3}] ModuleAccounts:[fee_collector inflation_collector validators_rewards_collector payment_collector liquididty_rewards_collector governance_locking_rewards_collector users_incentive_collector community_pool_rewards_collector]}
	// }
	// --------------------
	//
	//func prepareRoutingDistributor() (routingDistributor types.RoutingDistributor) {
	//
	//	routingDistributor types.RoutingDistributor {
	//		SubDistributor: types.SubDistributor{
	//			Name: "inflation_distributor",
	//			Sources: "inflation_collector",
	//
	//		}
	//	}
	//}

	//inflation_distributor := types.SubDistributor {
	//	Order: 1,
	//	Sources: [] string {"inflation_collector"},
	//	Destination: types.Desti
	//	}
	//}
	//
	//
	//
	//routingDistributor := types.RoutingDistributor {
	//
	//
	//
	//
	//	SubDistributor: [] types.SubDistributor {
	//		[			Name: "inflation_distributor",
	//		Sources: [] string {"inflation_collector"},
	//		],
	//
	//
	//	},
	//	ModuleAccounts: [] string {"asdasd","asdasdas"},
	//
	//}

}
