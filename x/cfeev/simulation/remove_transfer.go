package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"math/rand"

	"github.com/chain4energy/c4e-chain/x/cfeev/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRemoveTransfer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allEnergyTransferOffers := k.GetAllEnergyTransferOffers(ctx)
		energyTransferOffer := allEnergyTransferOffers[r.Intn(len(allEnergyTransferOffers))]
		simAccount, found := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(energyTransferOffer.Owner))
		if !found {

		}
		msgRemoveTransfer := &types.MsgRemoveTransfer{
			Creator: simAccount.Address.String(),
			Id:      energyTransferOffer.Id,
		}
		if err := simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgRemoveTransfer, chainID); err != nil {
			return simtypes.NewOperationMsg(msgRemoveTransfer, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(msgRemoveTransfer, true, "", nil), nil, nil
	}
}
