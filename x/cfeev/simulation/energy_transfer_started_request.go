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

func SimulateMsgEnergyTransferStarted(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		allEnergyTransfers := k.GetAllEnergyTransfers(ctx)
		energyTransfersLen := len(allEnergyTransfers)
		if energyTransfersLen == 0 {
			return simtypes.NewOperationMsg(&types.MsgEnergyTransferStarted{}, false, "", nil), nil, nil
		}
		energyTransferOffer := allEnergyTransfers[r.Intn(energyTransfersLen)]
		simAccount, found := simtypes.FindAccount(accs, sdk.MustAccAddressFromBech32(energyTransferOffer.Driver))
		if !found {
			return simtypes.NewOperationMsg(&types.MsgEnergyTransferStarted{}, false, "", nil), nil, nil
		}
		msgEnergyTransferStarted := &types.MsgEnergyTransferStarted{
			Creator:          simAccount.Address.String(),
			EnergyTransferId: energyTransferOffer.Id,
			Info:             "",
		}
		if err := simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgEnergyTransferStarted, chainID); err != nil {
			return simtypes.NewOperationMsg(msgEnergyTransferStarted, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(msgEnergyTransferStarted, true, "", nil), nil, nil
	}
}
