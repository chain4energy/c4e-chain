package simulation

//
//func SimulateMsgCreateUserDevices(
//	ak types.AccountKeeper,
//	bk types.BankKeeper,
//	k keeper.Keeper,
//) simtypes.Operation {
//	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//		simAccount, _ := simtypes.RandomAcc(r, accs)
//
//		msg := &types.MsgCreateUserDevices{
//			Owner: simAccount.Address.String(),
//		}
//
//		txCtx := simulation.OperationInput{
//			R:               r,
//			App:             app,
//			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
//			Cdc:             nil,
//			Msg:             msg,
//			MsgType:         msg.Type(),
//			Context:         ctx,
//			SimAccount:      simAccount,
//			ModuleName:      types.ModuleName,
//			CoinsSpentInMsg: sdk.NewCoins(),
//			AccountKeeper:   ak,
//			Bankkeeper:      bk,
//		}
//		return simulation.GenAndDeliverTxWithRandFees(txCtx)
//	}
//}
//
//func SimulateMsgUpdateUserDevices(
//	ak types.AccountKeeper,
//	bk types.BankKeeper,
//	k keeper.Keeper,
//) simtypes.Operation {
//	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//		var (
//			simAccount = simtypes.Account{}
//			userDevices = types.UserDevices{}
//			msg = &types.MsgUpdateUserDevices{}
//			allUserDevices = k.GetAllUserDevices(ctx)
//			found = false
//		)
//		for _, obj := range allUserDevices {
//			simAccount, found = FindAccount(accs, obj.Owner)
//			if found {
//				userDevices = obj
//				break
//			}
//		}
//		if !found {
//			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "userDevices owner not found"), nil, nil
//		}
//		msg.Owner = simAccount.Address.String()
//		msg.Id = userDevices.Id
//
//		txCtx := simulation.OperationInput{
//			R:               r,
//			App:             app,
//			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
//			Cdc:             nil,
//			Msg:             msg,
//			MsgType:         msg.Type(),
//			Context:         ctx,
//			SimAccount:      simAccount,
//			ModuleName:      types.ModuleName,
//			CoinsSpentInMsg: sdk.NewCoins(),
//			AccountKeeper:   ak,
//			Bankkeeper:      bk,
//		}
//		return simulation.GenAndDeliverTxWithRandFees(txCtx)
//	}
//}
//
//func SimulateMsgDeleteUserDevices(
//	ak types.AccountKeeper,
//	bk types.BankKeeper,
//	k keeper.Keeper,
//) simtypes.Operation {
//	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//		var (
//			simAccount = simtypes.Account{}
//			userDevices = types.UserDevices{}
//			msg = &types.MsgUpdateUserDevices{}
//			allUserDevices = k.GetAllUserDevices(ctx)
//			found = false
//		)
//		for _, obj := range allUserDevices {
//			simAccount, found = FindAccount(accs, obj.Owner)
//			if found {
//				userDevices = obj
//				break
//			}
//		}
//		if !found {
//			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "userDevices owner not found"), nil, nil
//		}
//		msg.Owner = simAccount.Address.String()
//		msg.Id = userDevices.Id
//
//		txCtx := simulation.OperationInput{
//			R:               r,
//			App:             app,
//			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
//			Cdc:             nil,
//			Msg:             msg,
//			MsgType:         msg.Type(),
//			Context:         ctx,
//			SimAccount:      simAccount,
//			ModuleName:      types.ModuleName,
//			CoinsSpentInMsg: sdk.NewCoins(),
//			AccountKeeper:   ak,
//			Bankkeeper:      bk,
//		}
//		return simulation.GenAndDeliverTxWithRandFees(txCtx)
//	}
//}
