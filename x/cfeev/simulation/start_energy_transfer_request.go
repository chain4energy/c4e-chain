package simulation

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
			fmt.Println("warning: start energy transfer failed" + err.Error())
			return simtypes.NewOperationMsg(&types.MsgStartEnergyTransfer{}, false, "", nil), nil, nil
		}
		return simtypes.NewOperationMsg(&types.MsgStartEnergyTransfer{}, true, "", nil), nil, nil
	}
}

func startEnergyTransfer(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper,
	app *baseapp.BaseApp, r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, chainID string) (simtypes.Account, uint64, error) {
	simAccount, _ := simtypes.RandomAcc(r, accs)
	allEnergyTransferOffers := k.GetAllEnergyTransferOffers(ctx)
	energyTransferOffersLen := len(allEnergyTransferOffers)
	if energyTransferOffersLen == 0 {
		return simtypes.Account{}, 0, sdkerrors.ErrNotFound
	}
	energyTransferOffer := allEnergyTransferOffers[r.Intn(energyTransferOffersLen)]
	msgStartEnergyTransfer := &types.MsgStartEnergyTransfer{
		Creator:               simAccount.Address.String(),
		EnergyTransferOfferId: energyTransferOffer.Id,
		OfferedTariff:         energyTransferOffer.Tariff,
		EnergyToTransfer:      utils.RandUint64(r, 100),
	}
	spendable := bk.SpendableCoins(ctx, simAccount.Address)
	if !spendable.IsAllPositive() {
		return simtypes.Account{}, 0, fmt.Errorf("balance is negative")
	}
	amount, err := simtypes.RandPositiveInt(r, spendable.AmountOf(sdk.DefaultBondDenom))
	if err != nil {
		return simtypes.Account{}, 0, fmt.Errorf("unable to generate positive amount")
	}
	vestingPoolAmount := sdk.NewCoin(sdk.DefaultBondDenom, amount)
	fmt.Println(vestingPoolAmount)
	result, err := simulation.SendMessageWithRandomFeesAndResult(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgStartEnergyTransfer, chainID)
	if err != nil {
		fmt.Println("warning: start energy transfer failed" + err.Error())
		return simtypes.Account{}, 0, err
	}

	response, _ := result.MsgResponses[0].GetCachedValue().(*types.MsgStartEnergyTransferResponse)
	return simAccount, response.Id, nil
}
