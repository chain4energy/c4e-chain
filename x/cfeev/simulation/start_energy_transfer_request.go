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

func SimulateMsgStartEnergyTransfer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		_, _, err := startEnergyTransfer(ak, bk, k, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsg(&types.MsgStartEnergyTransfer{}, false, "", nil), nil, nil
		}
		return simtypes.NewOperationMsg(&types.MsgStartEnergyTransfer{}, true, "", nil), nil, nil
	}
}

func startEnergyTransfer(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper,
	app *baseapp.BaseApp, r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, chainID string) (simtypes.Account, uint64, error) {
	simAccount, _ := simtypes.RandomAcc(r, accs)
	allEnergyTransferOffers := k.GetAllEnergyTransferOffers(ctx)
	energyTransferOffer := allEnergyTransferOffers[r.Intn(len(allEnergyTransferOffers))]
	msgStartEnergyTransfer := &types.MsgStartEnergyTransfer{
		Creator:               simAccount.Address.String(),
		EnergyTransferOfferId: energyTransferOffer.Id,
		ChargerId:             energyTransferOffer.ChargerId,
		OwnerAccountAddress:   energyTransferOffer.Owner,
		OfferedTariff:         energyTransferOffer.Tariff,
		Collateral:            nil,
		EnergyToTransfer:      utils.RandUint64(r, 10000),
	}

	result, err := simulation.SendMessageWithRandomFeesAndResult(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgStartEnergyTransfer, chainID)
	if err != nil {
		return simtypes.Account{}, 0, err
	}
	response, _ := result.MsgResponses[0].GetCachedValue().(*types.MsgStartEnergyTransferResponse)
	return simAccount, response.Id, nil
}
