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

func SimulateMsgRemoveEnergyOffer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, id, err := publishEnergyTransferOffer(ak, bk, k, app, r, ctx, accs, chainID)
		if err != nil {
			return simtypes.NewOperationMsg(&types.MsgRemoveEnergyOffer{}, false, "", nil), nil, nil
		}

		msgRemoveEnergyOffer := &types.MsgRemoveEnergyOffer{
			Owner: simAccount.Address.String(),
			Id:    id,
		}

		if err = simulation.SendMessageWithRandomFees(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msgRemoveEnergyOffer, chainID); err != nil {
			return simtypes.NewOperationMsg(msgRemoveEnergyOffer, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(msgRemoveEnergyOffer, true, "", nil), nil, nil
	}
}
