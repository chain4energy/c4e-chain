package simulation

import (
	"math/rand"

	"github.com/chain4energy/c4e-chain/x/cfesignature/keeper"
	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgPublishReferencePayloadLink(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgPublishReferencePayloadLink{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the PublishReferencePayloadLink simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "PublishReferencePayloadLink simulation not implemented"), nil, nil
	}
}
