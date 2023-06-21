package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"math/rand"

	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgPublishEnergyTransferOffer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, _, err := publishEnergyTransferOffer(ak, bk, k, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsg(&types.MsgPublishEnergyTransferOffer{}, false, "", nil), nil, nil
		}
		return simtypes.NewOperationMsg(&types.MsgPublishEnergyTransferOffer{}, true, "", nil), nil, nil
	}
}

func publishEnergyTransferOffer(ak types.AccountKeeper, bk types.BankKeeper, cfeev keeper.Keeper,
	app *baseapp.BaseApp, r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, chainID string) (simtypes.Account, uint64, error) {

	simAccount, _ := simtypes.RandomAcc(r, accs)
	msgPublishEnergyTransferOffer := &types.MsgPublishEnergyTransferOffer{
		Creator:   simAccount.Address.String(),
		ChargerId: simtypes.RandStringOfLength(r, 10),
		Tariff:    utils.RandUint64(r, 1000),
		Location:  randomLocation(r),
		Name:      simtypes.RandStringOfLength(r, 10),
		PlugType:  types.PlugType(utils.RandInt64(r, 4)),
	}

	result, err := simulation.SendMessageWithRandomFeesAndResult(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgPublishEnergyTransferOffer, chainID)
	if err != nil {
		return simtypes.Account{}, 0, err
	}
	response, _ := result.MsgResponses[0].GetCachedValue().(*types.MsgPublishEnergyTransferOfferResponse)
	return simAccount, response.Id, nil
}

func randomLocation(r *rand.Rand) *types.Location {
	latitude := utils.RandomDecBetween(r, -90, 90)
	longitude := utils.RandomDecBetween(r, -180, 190)
	return &types.Location{
		Latitude:  &latitude,
		Longitude: &longitude,
	}
}
