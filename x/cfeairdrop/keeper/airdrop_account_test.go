package keeper_test

//
//func TestCreateAirdropAccount(t *testing.T) {
//	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
//	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
//	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
//
//	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
//
//	moduleAmount := sdk.NewInt(1000000000000)
//	amount := sdk.NewInt(10000000000)
//
//	startTimeUnix := startTime.Unix()
//	endTimeUnix := endTime.Unix()
//	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
//
//	testHelper.C4eAirdropUtils.SendToNewRepeatedContinuousVestingAccount(acountsAddresses[0],
//		amount,
//		startTimeUnix,
//		endTimeUnix, cfeairdropmoduletypes.MissionVote,
//	)
//
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.Sub(types.OneForthC4e.Amount))
//	testHelper.SetContextBlockTime(testenv.TestEnvTime)
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.Sub(types.OneForthC4e.Amount).QuoRaw(2))
//	testHelper.SetContextBlockTime(endTime)
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], sdk.ZeroInt())
//
//	testHelper.SetContextBlockTime(startTime)
//	testHelper.C4eAirdropUtils.SendToNewRepeatedContinuousVestingAccount(acountsAddresses[0],
//		amount,
//		startTimeUnix,
//		endTimeUnix, cfeairdropmoduletypes.MissionVote,
//	)
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.MulRaw(2).Sub(types.OneForthC4e.Amount))
//	testHelper.SetContextBlockTime(testenv.TestEnvTime)
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.Sub(types.OneForthC4e.Amount).Add(types.OneForthC4e.Amount.QuoRaw(2)))
//	testHelper.SetContextBlockTime(endTime)
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], sdk.ZeroInt())
//
//	testHelper.SetContextBlockTime(startTime)
//	testHelper.C4eAirdropUtils.SendToNewRepeatedContinuousVestingAccount(acountsAddresses[0],
//		amount,
//		startTimeUnix,
//		endTimeUnix, cfeairdropmoduletypes.MissionVote,
//	)
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.MulRaw(3).Sub(types.OneForthC4e.Amount))
//	testHelper.SetContextBlockTime(testenv.TestEnvTime)
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], amount.QuoRaw(2).MulRaw(3).Sub(types.OneForthC4e.Amount).Add(types.OneForthC4e.Amount.QuoRaw(2)))
//	testHelper.SetContextBlockTime(endTime)
//	testHelper.BankUtils.VerifyAccountDefultDenomLocked(acountsAddresses[0], sdk.ZeroInt())
//}
//
//func TestCreateAirdropAccountSendDisabled(t *testing.T) {
//	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
//	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
//	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
//
//	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
//
//	moduleAmount := sdk.NewInt(10000)
//	amount := sdk.NewInt(1000)
//
//	startTimeUnix := startTime.Unix()
//	endTimeUnix := endTime.Unix()
//	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
//	testHelper.BankUtils.DisableDefaultSend()
//	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
//		amount,
//		startTimeUnix,
//		endTimeUnix, true, "send to airdrop account - send coins disabled: uc4e transfers are currently disabled: send transactions are disabled",
//		cfeairdropmoduletypes.MissionVote,
//	)
//}
//
//func TestCreateAirdropAccountBlockedAddress(t *testing.T) {
//	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
//	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
//	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
//	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
//	blockedAccounts := testHelper.App.ModuleAccountAddrs()
//	blockedAccounts[acountsAddresses[0].String()] = true
//	testHelper.App.BankKeeper = bankkeeper.NewBaseKeeper(
//		testHelper.App.AppCodec(), testHelper.App.GetKey(banktypes.StoreKey), testHelper.App.AccountKeeper, testHelper.App.GetSubspace(banktypes.ModuleName), blockedAccounts,
//	)
//	testHelper.App.CfeairdropKeeper = *cfeairdropmodulekeeper.NewKeeper(
//		testHelper.App.AppCodec(),
//		testHelper.App.GetKey(cfeairdropmoduletypes.StoreKey),
//		testHelper.App.GetKey(cfeairdropmoduletypes.MemStoreKey),
//		testHelper.App.GetSubspace(cfeairdropmoduletypes.ModuleName),
//
//		testHelper.App.AccountKeeper,
//		testHelper.App.BankKeeper,
//		testHelper.App.FeeGrantKeeper,
//	)
//
//	moduleAmount := sdk.NewInt(10000)
//	amount := sdk.NewInt(1000)
//
//	startTimeUnix := startTime.Unix()
//	endTimeUnix := endTime.Unix()
//	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
//	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
//		amount,
//		startTimeUnix,
//		endTimeUnix, true,
//		fmt.Sprintf("send to airdrop account - account address: %s is not allowed to receive funds error: unauthorized", acountsAddresses[0]),
//		cfeairdropmoduletypes.MissionVote,
//	)
//}
//
//func TestCreateAirdropAccountNotExist(t *testing.T) {
//	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
//	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
//	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
//
//	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
//
//	moduleAmount := sdk.NewInt(10000)
//	amount := sdk.NewInt(1000)
//
//	startTimeUnix := startTime.Unix()
//	endTimeUnix := endTime.Unix()
//	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
//	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
//		amount,
//		startTimeUnix,
//		endTimeUnix, false, fmt.Sprintf("create airdrop account - account does not exist: %s: entity does not exist", acountsAddresses[0]),
//		cfeairdropmoduletypes.MissionVote,
//	)
//}
//
//func TestCreateAirdropAccountWrongAccountType(t *testing.T) {
//	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
//	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
//	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
//
//	acountsAddresses, _ := testcosmos.CreateAccounts(1, 0)
//
//	baseAccount := testHelper.App.AccountKeeper.NewAccountWithAddress(testHelper.Context, acountsAddresses[0])
//	testHelper.App.AccountKeeper.SetAccount(testHelper.Context, baseAccount)
//	moduleAmount := sdk.NewInt(10000)
//	amount := sdk.NewInt(10000000000)
//
//	startTimeUnix := startTime.Unix()
//	endTimeUnix := endTime.Unix()
//	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)
//	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
//		amount,
//		startTimeUnix,
//		endTimeUnix, false, "send to airdrop account - expected RepeatedContinuousVestingAccount, got: *types.BaseAccount: invalid account type",
//		cfeairdropmoduletypes.MissionVote,
//	)
//}
//
//func TestCreateAirdropAccountSendError(t *testing.T) {
//	startTime := testenv.TestEnvTime.Add(-24 * 100 * time.Hour)
//	endTime := testenv.TestEnvTime.Add(24 * 100 * time.Hour)
//	testHelper := testapp.SetupTestAppWithHeightAndTime(t, 1000, startTime)
//
//	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
//	amount := sdk.NewInt(10000000000)
//
//	startTimeUnix := startTime.Unix()
//	endTimeUnix := endTime.Unix()
//	testHelper.BankUtils.AddDefaultDenomCoinsToModule(amount, types.ModuleName)
//
//	testHelper.C4eAirdropUtils.SendToAirdropAccountError(acountsAddresses[0],
//		amount.AddRaw(1),
//		startTimeUnix,
//		endTimeUnix, true, "send to airdrop account - send coins to airdrop account insufficient funds error (to: cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq, amount: 10000000001uc4e): insufficient funds",
//		cfeairdropmoduletypes.MissionInitialClaim,
//	)
//}
