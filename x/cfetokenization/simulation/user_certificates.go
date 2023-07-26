package simulation

//
//func SimulateMsgCreateUserCertificates(
//	ak types.AccountKeeper,
//	bk types.BankKeeper,
//	k keeper.Keeper,
//) simtypes.Operation {
//	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//		simAccount, _ := simtypes.RandomAcc(r, accs)
//
//		msg := &types.MsgCreateUserCertificates{
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
//func SimulateMsgUpdateUserCertificates(
//	ak types.AccountKeeper,
//	bk types.BankKeeper,
//	k keeper.Keeper,
//) simtypes.Operation {
//	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//		var (
//			simAccount          = simtypes.Account{}
//			userCertificates    = types.UserCertificates{}
//			msg                 = &types.MsgUpdateUserCertificates{}
//			allUserCertificates = k.GetAllUserCertificates(ctx)
//			found               = false
//		)
//		for _, obj := range allUserCertificates {
//			simAccount, found = FindAccount(accs, obj.Owner)
//			if found {
//				userCertificates = obj
//				break
//			}
//		}
//		if !found {
//			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "userCertificates owner not found"), nil, nil
//		}
//		msg.Owner = simAccount.Address.String()
//		msg.Id = userCertificates.Id
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
//func SimulateMsgDeleteUserCertificates(
//	ak types.AccountKeeper,
//	bk types.BankKeeper,
//	k keeper.Keeper,
//) simtypes.Operation {
//	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//		var (
//			simAccount          = simtypes.Account{}
//			msg                 = &types.MsgUpdateUserCertificates{}
//			allUserCertificates = k.GetAllUserCertificates(ctx)
//			found               = false
//		)
//		for _, obj := range allUserCertificates {
//			simAccount, found = FindAccount(accs, obj.Owner)
//			if found {
//				userCertificates = obj
//				break
//			}
//		}
//		if !found {
//			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "userCertificates owner not found"), nil, nil
//		}
//		msg.Owner = simAccount.Address.String()
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
