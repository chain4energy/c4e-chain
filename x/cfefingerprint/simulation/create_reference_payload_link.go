package simulation

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/keeper"
	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"math/rand"
)

func SimulateMsgCreateReferencePayloadLink(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		randonPayloadHash := simtypes.RandStringOfLength(r, 32)
		msg := types.NewMsgCreateReferencePayloadLink(simAccount.Address.String(), randonPayloadHash)

		result, err := simulation.SendMessageWithRandomFeesWithResult(ctx, r, ak.(authkeeper.AccountKeeper), bk.(bankkeeper.Keeper), app, simAccount, msg, chainID)
		if err != nil {
			return simtypes.NewOperationMsg(msg, false, "", nil), nil, nil
		}

		response, _ := result.MsgResponses[0].GetCachedValue().(*types.MsgCreateReferencePayloadLinkResponse)
		if err = k.VerifyPayloadLink(ctx, response.ReferenceId, randonPayloadHash); err != nil {
			return simtypes.NewOperationMsg(msg, false, "", nil), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
