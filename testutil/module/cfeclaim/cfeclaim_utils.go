package cfeclaim

import (
	"cosmossdk.io/math"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfeclaimmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingmodulekeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/require"
	"time"
)

type C4eClaimUtils struct {
	C4eClaimKeeperUtils
	helperAccountKeeper    *authkeeper.AccountKeeper
	BankUtils              *testcosmos.BankUtils
	DistrUtils             *testcosmos.DistributionUtils
	FeegrantUtils          *testcosmos.FeegrantUtils
	StakingUtils           *testcosmos.StakingUtils
	GovUtils               *testcosmos.GovUtils
	DistributionUtils      *testcosmos.DistributionUtils
	helperCfevestingKeeper *cfevestingmodulekeeper.Keeper
}

func NewC4eClaimUtils(t require.TestingT, helpeCfeclaimmodulekeeper *cfeclaimmodulekeeper.Keeper, helperCfevestingKeeper *cfevestingmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *testcosmos.BankUtils, stakingUtils *testcosmos.StakingUtils, govUtils *testcosmos.GovUtils, feegrantUtils *testcosmos.FeegrantUtils, distributionUtils *testcosmos.DistributionUtils) C4eClaimUtils {
	return C4eClaimUtils{C4eClaimKeeperUtils: NewC4eClaimKeeperUtils(t, helpeCfeclaimmodulekeeper),
		helperAccountKeeper: helperAccountKeeper, BankUtils: bankUtils, StakingUtils: stakingUtils, GovUtils: govUtils,
		FeegrantUtils: feegrantUtils, DistributionUtils: distributionUtils, helperCfevestingKeeper: helperCfevestingKeeper}
}

func (h *C4eClaimUtils) SendToRepeatedContinuousVestingAccount(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, missionType cfeclaimtypes.MissionType) {
	coins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amount))
	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	accBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)

	previousOriginalVesting := sdk.NewCoins()
	var previousPeriods []cfevestingtypes.ContinuousVestingPeriod
	if accountBefore != nil {
		if claimAccount, ok := accountBefore.(*cfevestingtypes.PeriodicContinuousVestingAccount); ok {
			previousOriginalVesting = previousOriginalVesting.Add(claimAccount.OriginalVesting...)
			previousPeriods = claimAccount.VestingPeriods
		}
	}
	userEntry := &cfeclaimtypes.UserEntry{
		Address:      toAddress.String(),
		ClaimAddress: toAddress.String(),
	}

	require.NoError(h.t, h.helpeCfeclaimkeeper.SendToNewPeriodicContinuousVestingAccount(ctx,
		userEntry,
		coins,
		startTime,
		endTime,
		missionType,
	))

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance.Add(amount))
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBalance.Sub(amount))

	claimAccount, ok := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	require.True(h.t, ok)
	newPeriods := append(previousPeriods, cfevestingtypes.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: coins})
	h.VerifyRepeatedContinuousVestingAccount(ctx, toAddress, previousOriginalVesting.Add(coins...), startTime, endTime, newPeriods, missionType)
	require.NoError(h.t, claimAccount.Validate())
}

func (h *C4eClaimUtils) SendToRepeatedContinuousVestingAccountError(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, missionType cfeclaimtypes.MissionType) {
	coins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amount))
	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	accBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	wasAccount := false
	if accountBefore != nil {
		_, wasAccount = accountBefore.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	}
	userEntry := &cfeclaimtypes.UserEntry{
		Address:      toAddress.String(),
		ClaimAddress: toAddress.String(),
	}
	require.EqualError(h.t, h.helpeCfeclaimkeeper.SendToNewPeriodicContinuousVestingAccount(ctx,
		userEntry,
		coins,
		startTime,
		endTime,
		missionType,
	), errorMessage)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBalance)

	accountAfter := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	_, isAccount := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	_, ok := accountBefore.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	if ok && missionType == cfeclaimtypes.MissionInitialClaim {
		require.EqualValues(h.t, true, isAccount)
		h.VerifyRepeatedContinuousVestingAccount(ctx, toAddress, sdk.NewCoins(), startTime, endTime, []cfevestingtypes.ContinuousVestingPeriod{}, missionType)

	} else {
		require.EqualValues(h.t, wasAccount, isAccount)
		require.EqualValues(h.t, accountBefore, accountAfter)
	}

}

func (h *C4eClaimUtils) VerifyRepeatedContinuousVestingAccount(ctx sdk.Context, address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfevestingtypes.ContinuousVestingPeriod, missionType cfeclaimtypes.MissionType) {

	claimAccount, ok := h.helperAccountKeeper.GetAccount(ctx, address).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	require.True(h.t, ok)

	require.EqualValues(h.t, len(expectedPeriods), len(claimAccount.VestingPeriods))
	require.EqualValues(h.t, expectedStartTime, claimAccount.StartTime)
	require.EqualValues(h.t, expectedEndTime, claimAccount.EndTime)
	require.True(h.t, expectedOriginalVesting.IsEqual(claimAccount.OriginalVesting))
	for i := 0; i < len(expectedPeriods); i++ {
		require.EqualValues(h.t, expectedPeriods[i].StartTime, claimAccount.VestingPeriods[i].StartTime)
		require.EqualValues(h.t, expectedPeriods[i].EndTime, claimAccount.VestingPeriods[i].EndTime)
		require.EqualValues(h.t, expectedPeriods[i].Amount, claimAccount.VestingPeriods[i].Amount)
	}
	require.NoError(h.t, claimAccount.Validate())
}

func (h *C4eClaimUtils) AddClaimRecords(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, claimEntries []*cfeclaimtypes.ClaimRecord) {
	amountSum := sdk.NewCoins()
	for _, claimRecord := range claimEntries {
		amountSum = amountSum.Add(claimRecord.Amount...)
	}
	usersEntriesBefore := h.helpeCfeclaimkeeper.GetUsersEntries(ctx)
	claimClaimsLeftBefore, ok := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	if !ok {
		claimClaimsLeftBefore = cfeclaimtypes.CampaignAmountLeft{
			Amount:     sdk.NewCoins(),
			CampaignId: campaignId,
		}
	}
	claimDistrubitionsBefore, ok := h.helpeCfeclaimkeeper.GetCampaignTotalAmount(ctx, campaignId)
	if !ok {
		claimDistrubitionsBefore = cfeclaimtypes.CampaignTotalAmount{
			Amount:     sdk.NewCoins(),
			CampaignId: campaignId,
		}
	}
	srcBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	campaign, _ := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	var vestingPoolBefore cfevestingtypes.VestingPool
	if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
		var accFound bool
		vestingPoolBefore, accFound = h.helperCfevestingKeeper.GetAccountVestingPool(ctx, srcAddress.String(), campaign.VestingPoolName)
		require.True(h.t, accFound)
	}

	require.NoError(h.t, h.helpeCfeclaimkeeper.AddClaimRecords(ctx, srcAddress.String(), campaignId, claimEntries))
	claimClaimsLeftAfter, _ := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	claimDistrubitionsAfter, _ := h.helpeCfeclaimkeeper.GetCampaignTotalAmount(ctx, campaignId)
	claimClaimsLeftBefore.Amount = claimClaimsLeftBefore.Amount.Add(amountSum...)
	claimDistrubitionsBefore.Amount = claimDistrubitionsBefore.Amount.Add(amountSum...)
	require.EqualValues(h.t, claimClaimsLeftBefore, claimClaimsLeftAfter)
	require.EqualValues(h.t, claimDistrubitionsBefore, claimDistrubitionsAfter)

	// TODO: add check if len(beforeAllEntry) + len(claimEntries) == len(afterAllEntry)
	for _, claimRecord := range claimEntries {
		userEntry, found := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimRecord.Address)
		require.True(h.t, found)
		var userEntryBefore *cfeclaimtypes.UserEntry = nil
		for _, before := range usersEntriesBefore {
			if before.Address == claimRecord.Address {
				userEntryBefore = &before
				break
			}
		}
		if userEntryBefore == nil {
			require.EqualValues(h.t, 1, len(userEntry.ClaimRecords))
		} else {
			require.EqualValues(h.t, len(userEntryBefore.ClaimRecords)+1, len(userEntry.ClaimRecords))
		}
		require.EqualValues(h.t, claimRecord.Address, userEntry.Address)
		require.EqualValues(h.t, "", userEntry.ClaimAddress)
		if userEntryBefore == nil {
			require.EqualValues(h.t, campaignId, userEntry.ClaimRecords[0].CampaignId)
			require.True(h.t, claimRecord.Amount.IsEqual(userEntry.ClaimRecords[0].Amount))
			require.EqualValues(h.t, 0, len(userEntry.ClaimRecords[0].CompletedMissions))
		} else {
			expectedCaipaignRecords := append(userEntryBefore.ClaimRecords, &cfeclaimtypes.ClaimRecord{CampaignId: campaignId, Amount: claimRecord.Amount, CompletedMissions: nil})

			require.ElementsMatch(h.t, expectedCaipaignRecords, userEntry.ClaimRecords)
		}

	}

	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		_, feegrandModuleAddress := cfeclaimmodulekeeper.CreateFeegrantAccountAddress(campaignId)
		feegrantSum := campaign.FeegrantAmount.MulRaw(int64(len(claimEntries)))
		h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBalance.Add(amountSum.AmountOf(testenv.DefaultTestDenom)))
		h.BankUtils.VerifyAccountDefultDenomBalance(ctx, feegrandModuleAddress, feegrantSum)
		h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance.Sub(feegrantSum).Sub(amountSum.AmountOf(testenv.DefaultTestDenom)))
	} else {
		h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBalance.Add(amountSum.AmountOf(testenv.DefaultTestDenom)))
		if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
			vestingPool, accFound := h.helperCfevestingKeeper.GetAccountVestingPool(ctx, srcAddress.String(), campaign.VestingPoolName)
			require.True(h.t, accFound)
			require.EqualValues(h.t, vestingPoolBefore.Sent.Add(amountSum.AmountOf(testenv.DefaultTestDenom)), vestingPool.Sent)
		} else {
			h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance.Sub(amountSum.AmountOf(testenv.DefaultTestDenom)))
		}

	}
}

func (h *C4eClaimUtils) AddClaimRecordsError(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, claimEntries []*cfeclaimtypes.ClaimRecord, errorMessage string) {
	ownerBalanceBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)
	moduleBalanceBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	err := h.helpeCfeclaimkeeper.AddClaimRecords(ctx, srcAddress.String(), campaignId, claimEntries)
	require.EqualError(h.t, err, errorMessage)

	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBalanceBefore)
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, ownerBalanceBefore)
}

func (h *C4eClaimUtils) AddClaimRecordsFromWhitelistedVestingAccount(ctx sdk.Context, from sdk.AccAddress, amountToSend sdk.Coins, unlockedAmount sdk.Coins) {
	err := h.helpeCfeclaimkeeper.AddClaimRecordsFromWhitelistedVestingAccount(ctx, from, amountToSend)
	require.NoError(h.t, err)
	h.BankUtils.VerifyAccountDefultDenomSpendableCoins(ctx, from, unlockedAmount.AmountOf(testenv.DefaultTestDenom))
}

func (h *C4eClaimUtils) AddCampaignRecordsError(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, claimEntries []*cfeclaimtypes.ClaimRecord, errorMessage string, addRequiredCoinsToSrc bool) {
	if addRequiredCoinsToSrc {
		sum := sdk.NewCoins()
		for _, claimRecord := range claimEntries {
			sum = sum.Add(claimRecord.Amount...)
		}
		h.BankUtils.AddDefaultDenomCoinsToAccount(ctx, sum.AmountOf(testenv.DefaultTestDenom), srcAddress)
	}
	usersEntriesBefore := h.helpeCfeclaimkeeper.GetUsersEntries(ctx)

	srcBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeclaimkeeper.AddClaimRecords(ctx, srcAddress.String(), campaignId, claimEntries), errorMessage)

	usersEntriesAfter := h.helpeCfeclaimkeeper.GetUsersEntries(ctx)
	require.ElementsMatch(h.t, usersEntriesBefore, usersEntriesAfter)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBalance)
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance)

}

func (h *C4eClaimUtils) ClaimInitial(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress, expectedAmount int64) {
	acc := h.helperAccountKeeper.GetAccount(ctx, claimer)
	claimerAccountBefore, ok := acc.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	accExisted := acc != nil
	if !accExisted {
		claimerAccountBefore = nil
	}
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	campaign, _ := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	claimClaimsLeftBefore, ok := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)

	userEntry, _ := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	_, granterAddr := cfeclaimmodulekeeper.CreateFeegrantAccountAddress(campaignId)
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		allowance, err := h.FeegrantUtils.FeegrantKeeper.GetAllowance(ctx, granterAddr, claimer)
		require.NoError(h.t, err)
		require.NotNil(h.t, allowance)
	}
	err := h.helpeCfeclaimkeeper.InitialClaim(ctx, claimer.String(), campaignId, claimer.String())
	require.NoError(h.t, err)
	allowance, err := h.FeegrantUtils.FeegrantKeeper.GetAllowance(ctx, granterAddr, claimer)
	require.Error(h.t, err)
	require.Nil(h.t, allowance)
	claimClaimsLeftAfter, ok := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	claimClaimsLeftBefore.Amount = claimClaimsLeftBefore.Amount.Sub(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(expectedAmount)))...)

	require.EqualValues(h.t, claimClaimsLeftBefore, claimClaimsLeftAfter)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore.AddRaw(expectedAmount))
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBefore.SubRaw(expectedAmount))

	if claimerAccountBefore == nil {
		baseAccount := h.helperAccountKeeper.NewAccountWithAddress(ctx, claimer)
		claimerAccountBefore = cfevestingtypes.NewRepeatedContinuousVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), 100000000, 100000000, nil)
	}

	vestingAmount := sdk.NewInt(expectedAmount)
	if campaign.InitialClaimFreeAmount.GT(sdk.ZeroInt()) {
		vestingAmount = sdk.NewInt(expectedAmount).Sub(campaign.InitialClaimFreeAmount)
	}
	claimerAccountBefore = h.addExpectedDataToAccount(ctx, campaignId, claimerAccountBefore, vestingAmount)

	claimerAccount, ok := h.helperAccountKeeper.GetAccount(ctx, claimer).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	if !accExisted {
		claimerAccountBefore.AccountNumber = claimerAccount.AccountNumber
	}
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		claimerAccount.AccountNumber++
	}
	require.True(h.t, ok)
	require.NoError(h.t, claimerAccount.Validate())
	require.EqualValues(h.t, claimerAccountBefore, claimerAccount)

	userEntry, found := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.True(h.t, found)
	claimRecord := userEntry.GetClaimRecord(campaignId)
	require.NotNil(h.t, claimRecord)
}

func (h *C4eClaimUtils) ClaimInitialError(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	balanceBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	userEntryBefore, foundBefore := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())

	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeclaimkeeper.InitialClaim(ctx, claimer.String(), campaignId, claimer.String()), errorMessage)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, balanceBefore)

	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBefore)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))

	userEntry, found := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.Equal(h.t, foundBefore, found)
	if found {
		require.EqualValues(h.t, userEntryBefore, userEntry)

	}
}

func (h *C4eClaimUtils) GetUsersEntries(
	ctx sdk.Context,
	address string,
) *cfeclaimtypes.UserEntry {
	val, found := h.helpeCfeclaimkeeper.GetUserEntry(ctx, address)
	if found {
		return &val
	}
	return nil
}

func (h *C4eClaimUtils) SetUsersEntries(
	ctx sdk.Context,
	userEntry *cfeclaimtypes.UserEntry,
) {
	h.helpeCfeclaimkeeper.SetUserEntry(ctx, *userEntry)
}

func (h *C4eClaimUtils) CompleteDelegationMission(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, deleagtionAmount math.Int, valAddress sdk.ValAddress) {
	action := func() error {
		h.StakingUtils.SetupValidators(ctx, []sdk.ValAddress{valAddress}, sdk.NewInt(1))
		h.StakingUtils.MessageDelegate(ctx, 2, 0, valAddress, claimer, deleagtionAmount)
		return nil
	}
	beforeCheck := func(accBefore authtypes.AccountI, accAfter authtypes.AccountI, claimerAmountBefore sdk.Int) (authtypes.AccountI, sdk.Int) {
		veBefore, okBefore := accBefore.(*cfevestingtypes.PeriodicContinuousVestingAccount)
		veAfter, okAfter := accAfter.(*cfevestingtypes.PeriodicContinuousVestingAccount)
		if okBefore && okAfter {
			veBefore.DelegatedFree = veAfter.DelegatedFree
			veBefore.DelegatedVesting = veAfter.DelegatedVesting
		}
		return veBefore, claimerAmountBefore.Sub(deleagtionAmount)
	}
	h.completeAnyMission(ctx, campaignId, missionId, claimer, action, beforeCheck)
}

func (h *C4eClaimUtils) CompleteVoteMission(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	action := func() error {
		depParams := h.GovUtils.GovKeeper.GetDepositParams(ctx)
		depositAmount := depParams.MinDeposit
		h.BankUtils.AddCoinsToAccount(ctx, depositAmount, claimer)

		testProposal := &cfevestingtypes.MsgUpdateDenomParam{Authority: appparams.GetAuthority(), Denom: testenv.DefaultTestDenom}
		proposal, err := h.GovUtils.GovKeeper.SubmitProposal(ctx, []sdk.Msg{testProposal}, "=abc")
		if err != nil {
			return err
		}

		_, err = h.GovUtils.GovKeeper.AddDeposit(ctx, proposal.Id, claimer, depositAmount)
		if err != nil {
			return err
		}

		return h.GovUtils.GovKeeper.AddVote(ctx, proposal.Id, claimer,
			govv1types.WeightedVoteOptions{{Option: govv1types.VoteOption_VOTE_OPTION_YES, Weight: "1"}}, "=abc")
	}
	h.completeAnyMission(ctx, campaignId, missionId, claimer, action, nil)
}

func (h *C4eClaimUtils) CompleteMissionFromHook(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	action := func() error {
		return h.helpeCfeclaimkeeper.CompleteMissionFromHook(ctx, campaignId, missionId, claimer.String())
	}
	h.completeAnyMission(ctx, campaignId, missionId, claimer, action, nil)
}

func (h *C4eClaimUtils) completeAnyMission(ctx sdk.Context, campaignId uint64, missionId uint64,
	claimer sdk.AccAddress, action func() error, beforeCheck func(before authtypes.AccountI, after authtypes.AccountI, ampountBefore sdk.Int) (authtypes.AccountI, sdk.Int)) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	userEntryBefore, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.True(h.t, foundCr)

	require.NoError(h.t, action())
	claimerAccountAfter := h.helperAccountKeeper.GetAccount(ctx, claimer)
	if beforeCheck != nil {
		claimerAccountBefore, claimerBefore = beforeCheck(claimerAccountBefore, claimerAccountAfter, claimerBefore)
	}
	require.EqualValues(h.t, claimerAccountBefore, claimerAccountAfter)
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBefore)

	userEntry, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.True(h.t, foundCr)
	userEntryBefore.GetClaimRecord(campaignId).CompletedMissions = append(userEntryBefore.GetClaimRecord(campaignId).CompletedMissions, missionId)
	require.EqualValues(h.t, userEntryBefore, userEntry)
}

func (h *C4eClaimUtils) CompleteMissionFromHookError(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	claimRecordBefore, foundCrBefore := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())

	require.EqualError(h.t, h.helpeCfeclaimkeeper.CompleteMissionFromHook(ctx, campaignId, missionId, claimer.String()), errorMessage)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBefore)
	userEntry, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.Equal(h.t, foundCrBefore, foundCr)
	require.EqualValues(h.t, claimRecordBefore, userEntry)
}

func (h *C4eClaimUtils) ClaimMission(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.ClaimMissionToAddress(ctx, campaignId, missionId, claimer, claimer)
}

func (h *C4eClaimUtils) ClaimMissionToAddress(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, claimerDstAddress sdk.AccAddress) {
	claimerAccountBefore, ok := h.helperAccountKeeper.GetAccount(ctx, claimerDstAddress).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	require.True(h.t, ok)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimerDstAddress)
	userEntryBefore, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.True(h.t, foundCr)
	mission, _ := h.helpeCfeclaimkeeper.GetMission(ctx, campaignId, missionId)
	require.NoError(h.t, h.helpeCfeclaimkeeper.Claim(ctx, campaignId, missionId, claimer.String()))

	userEntryBefore.GetClaimRecord(campaignId).ClaimedMissions = append(userEntryBefore.GetClaimRecord(campaignId).ClaimedMissions, missionId)
	if mission.MissionType == cfeclaimtypes.MissionClaim {
		userEntryBefore.GetClaimRecord(campaignId).CompletedMissions = append(userEntryBefore.GetClaimRecord(campaignId).CompletedMissions, missionId)
	}
	userEntry, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.True(h.t, foundCr)

	require.EqualValues(h.t, userEntryBefore, userEntry)

	expectedAmount := mission.Weight.MulInt(userEntry.GetClaimRecord(campaignId).Amount.AmountOf(testenv.DefaultTestDenom)).TruncateInt()

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimerDstAddress, claimerBefore.Add(expectedAmount))
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBefore.Sub(expectedAmount))

	h.addExpectedDataToAccount(ctx, campaignId, claimerAccountBefore, expectedAmount)

	claimerAccount, ok := h.helperAccountKeeper.GetAccount(ctx, claimerDstAddress).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	require.True(h.t, ok)
	require.NoError(h.t, claimerAccount.Validate())
	require.EqualValues(h.t, claimerAccountBefore, claimerAccount)
}

func (h *C4eClaimUtils) addExpectedDataToAccount(ctx sdk.Context, campaignId uint64,
	claimerAccountBefore *cfevestingtypes.PeriodicContinuousVestingAccount, expectedAmount sdk.Int) *cfevestingtypes.PeriodicContinuousVestingAccount {
	campaign, _ := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	expectedStartTime := ctx.BlockTime().Add(campaign.LockupPeriod)
	expectedEndTime := expectedStartTime.Add(campaign.VestingPeriod)
	expectedOriginalVesting := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, expectedAmount))
	if len(claimerAccountBefore.VestingPeriods) == 0 {
		claimerAccountBefore.StartTime = expectedStartTime.Unix()
		claimerAccountBefore.EndTime = expectedEndTime.Unix()

	} else {
		if claimerAccountBefore.StartTime > expectedStartTime.Unix() {
			claimerAccountBefore.StartTime = expectedStartTime.Unix()
		}
		if claimerAccountBefore.EndTime < expectedEndTime.Unix() {
			claimerAccountBefore.EndTime = expectedEndTime.Unix()
		}
	}
	claimerAccountBefore.OriginalVesting = claimerAccountBefore.OriginalVesting.Add(expectedOriginalVesting...)
	claimerAccountBefore.VestingPeriods = append(claimerAccountBefore.VestingPeriods, cfevestingtypes.ContinuousVestingPeriod{StartTime: expectedStartTime.Unix(), EndTime: expectedEndTime.Unix(), Amount: expectedOriginalVesting})
	return claimerAccountBefore
}

func (h *C4eClaimUtils) ClaimMissionError(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	claimRecordBefore, foundCrBefore := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())

	require.EqualError(h.t, h.helpeCfeclaimkeeper.Claim(ctx, campaignId, missionId, claimer.String()), errorMessage)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, moduleBefore)
	userEntry, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.Equal(h.t, foundCrBefore, foundCr)
	require.EqualValues(h.t, claimRecordBefore, userEntry)
}

func (h *C4eClaimUtils) CreateRepeatedContinuousVestingAccount(ctx sdk.Context, address sdk.AccAddress, originalVesting sdk.Coins, startTime int64, endTime int64, periods ...cfevestingtypes.ContinuousVestingPeriod) *cfevestingtypes.PeriodicContinuousVestingAccount {
	baseAccount := h.helperAccountKeeper.NewAccountWithAddress(ctx, address)
	claimAcc := cfevestingtypes.NewRepeatedContinuousVestingAccount(baseAccount.(*authtypes.BaseAccount), originalVesting, startTime, endTime, periods)
	h.helperAccountKeeper.SetAccount(ctx, claimAcc)
	require.NoError(h.t, claimAcc.Validate())
	return claimAcc
}

func (h *C4eClaimUtils) CreateCampaign(ctx sdk.Context, owner string, name string, description string, campaignType cfeclaimtypes.CampaignType, feegrantAmount sdk.Int, initialClaimFreeAmount sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration, vestingPoolName string) {

	campaignCountBefore := h.helpeCfeclaimkeeper.GetCampaignCount(ctx)
	_, err := h.helpeCfeclaimkeeper.CreateCampaign(ctx, owner, name, description, campaignType, &feegrantAmount, &initialClaimFreeAmount, &startTime, &endTime, &lockupPeriod, &vestingPeriod, vestingPoolName)
	missionCountAfter := h.helpeCfeclaimkeeper.GetMissionCount(ctx, campaignCountBefore)
	require.NoError(h.t, err)
	campaignCountAfter := h.helpeCfeclaimkeeper.GetCampaignCount(ctx)
	require.Equal(h.t, campaignCountBefore+1, campaignCountAfter)
	require.Equal(h.t, uint64(1), missionCountAfter)

	h.VerifyCampaign(ctx, campaignCountBefore, true, owner, name, description, false, &feegrantAmount, &initialClaimFreeAmount, startTime, endTime, lockupPeriod, vestingPeriod, vestingPoolName)
	h.VerifyMission(ctx, true, campaignCountBefore, 0, "Initial mission", "Initial mission - basic mission that must be claimed first", cfeclaimtypes.MissionInitialClaim, sdk.ZeroDec(), nil)
}

func (h *C4eClaimUtils) CreateCampaignError(ctx sdk.Context, owner string, name string, description string, campaignType cfeclaimtypes.CampaignType, feegrantAmount sdk.Int, initialClaimFreeAmount sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration, vestingPoolName string, errorMessage string) {

	campaignCountBefore := h.helpeCfeclaimkeeper.GetCampaignCount(ctx)
	_, err := h.helpeCfeclaimkeeper.CreateCampaign(ctx, owner, name, description, campaignType, &feegrantAmount, &initialClaimFreeAmount, &startTime, &endTime, &lockupPeriod, &vestingPeriod, vestingPoolName)
	require.EqualError(h.t, err, errorMessage)
	campaignCountAfter := h.helpeCfeclaimkeeper.GetCampaignCount(ctx)
	missionCountAfter := h.helpeCfeclaimkeeper.GetMissionCount(ctx, campaignCountBefore)
	require.Equal(h.t, campaignCountBefore, campaignCountAfter)
	require.Equal(h.t, uint64(0), missionCountAfter)
	_, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignCountBefore)
	require.False(h.t, ok)
}

func (h *C4eClaimUtils) StartCampaign(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) {
	err := h.helpeCfeclaimkeeper.StartCampaign(ctx, owner, campaignId, startTime, endTime)
	require.NoError(h.t, err)
	campaign, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	require.True(h.t, ok)
	h.VerifyCampaign(ctx, campaign.Id, true, owner, campaign.Name, campaign.Description, true, &campaign.FeegrantAmount, &campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod, campaign.VestingPoolName)
}

func (h *C4eClaimUtils) StartCampaignError(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time, errorString string) {
	campaignBefore, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)

	err := h.helpeCfeclaimkeeper.StartCampaign(ctx, owner, campaignId, startTime, endTime)
	require.EqualError(h.t, err, errorString)
	if !ok {
		return
	}
	enabled := campaignBefore.Enabled
	h.VerifyCampaign(ctx, campaignBefore.Id, true, campaignBefore.Owner, campaignBefore.Name, campaignBefore.Description, enabled, &campaignBefore.FeegrantAmount, &campaignBefore.InitialClaimFreeAmount, campaignBefore.StartTime, campaignBefore.EndTime, campaignBefore.LockupPeriod, campaignBefore.VestingPeriod, campaignBefore.VestingPoolName)
}

func (h *C4eClaimUtils) CloseCampaign(ctx sdk.Context, owner string, campaignId uint64, CloseAction cfeclaimtypes.CloseAction) {
	campaignAmoutLeftBefore, _ := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	cfeclaimModuleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName)
	campaign, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	_, feegrantAccountAddress := cfeclaimmodulekeeper.CreateFeegrantAccountAddress(campaign.Id)
	feegrantAmountLefBefore := math.ZeroInt()
	if campaign.FeegrantAmount.IsPositive() {
		feegrantAmountLefBefore = h.BankUtils.GetAccountDefultDenomBalance(ctx, feegrantAccountAddress)
	}
	ownerAccAddress, _ := sdk.AccAddressFromBech32(campaign.Owner)
	ownerBalanceBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, ownerAccAddress)
	var vestingPoolBefore cfevestingtypes.VestingPool

	if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
		vestingPoolBefore, _ = h.helperCfevestingKeeper.GetAccountVestingPool(ctx, owner, campaign.VestingPoolName)
	}

	err := h.helpeCfeclaimkeeper.CloseCampaign(ctx, owner, campaignId, CloseAction)
	require.NoError(h.t, err)

	campaign, _ = h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	campaignAmoutLeft, _ := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	require.True(h.t, campaignAmoutLeft.Amount.IsEqual(sdk.NewCoins()))

	if campaign.FeegrantAmount.IsPositive() {
		feegrantAmountLef := h.BankUtils.GetAccountDefultDenomBalance(ctx, feegrantAccountAddress)
		require.True(h.t, feegrantAmountLef.IsZero())
	}
	switch CloseAction {
	case cfeclaimtypes.CloseSendToCommunityPool:
		feePool := h.DistributionUtils.DistrKeeper.GetFeePool(ctx)
		feePoolAmount := feePool.CommunityPool.AmountOf(testenv.DefaultTestDenom)
		expectedAmount := sdk.NewDecFromInt(campaignAmoutLeftBefore.Amount.AmountOf(testenv.DefaultTestDenom).Add(feegrantAmountLefBefore))
		require.True(h.t, feePoolAmount.Equal(expectedAmount))
	case cfeclaimtypes.CloseSendToOwner:
		if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
			vestingPool, _ := h.helperCfevestingKeeper.GetAccountVestingPool(ctx, owner, campaign.VestingPoolName)
			if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
				require.True(h.t, vestingPoolBefore.Sent.
					Sub(campaignAmoutLeftBefore.Amount.AmountOf(testenv.DefaultTestDenom)).Sub(feegrantAmountLefBefore).Equal(vestingPool.Sent))
			} else {
				require.True(h.t, vestingPoolBefore.Sent.Sub(campaignAmoutLeftBefore.Amount.AmountOf(testenv.DefaultTestDenom)).Equal(vestingPool.Sent))
			}

		} else {
			h.BankUtils.VerifyAccountDefultDenomBalance(ctx, ownerAccAddress, ownerBalanceBefore.Add(campaignAmoutLeftBefore.Amount.AmountOf(testenv.DefaultTestDenom).Add(feegrantAmountLefBefore)))
		}
	}

	require.True(h.t, ok)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeclaimtypes.ModuleName, cfeclaimModuleBalance.Sub(campaignAmoutLeftBefore.Amount.AmountOf(testenv.DefaultTestDenom)))
	h.VerifyCloseAction(ctx, campaignId, CloseAction, campaignAmoutLeftBefore.Amount)
	h.VerifyCampaign(ctx, campaign.Id, true, owner, campaign.Name, campaign.Description, false, &campaign.FeegrantAmount, &campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod, campaign.VestingPoolName)
}

func (h *C4eClaimUtils) CloseCampaignError(ctx sdk.Context, owner string, campaignId uint64, CloseAction cfeclaimtypes.CloseAction, errorString string) {
	campaignBefore, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)

	err := h.helpeCfeclaimkeeper.CloseCampaign(ctx, owner, campaignId, CloseAction)
	require.EqualError(h.t, err, errorString)
	if !ok {
		return
	}
	enabled := campaignBefore.Enabled
	h.VerifyCampaign(ctx, campaignBefore.Id, true, campaignBefore.Owner, campaignBefore.Name, campaignBefore.Description, enabled, &campaignBefore.FeegrantAmount, &campaignBefore.InitialClaimFreeAmount, campaignBefore.StartTime, campaignBefore.EndTime, campaignBefore.LockupPeriod, campaignBefore.VestingPeriod, campaignBefore.VestingPoolName)
}

func (h *C4eClaimUtils) AddMissionToCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType cfeclaimtypes.MissionType,
	weight sdk.Dec, missionClaimDate *time.Time) {
	missionCountBefore := h.helpeCfeclaimkeeper.GetMissionCount(ctx, campaignId)
	err := h.helpeCfeclaimkeeper.AddMissionToCampaign(ctx, owner, campaignId, name, description, missionType, weight, nil)
	missionCountAfter := h.helpeCfeclaimkeeper.GetMissionCount(ctx, campaignId)
	require.NoError(h.t, err)
	require.Equal(h.t, missionCountBefore+1, missionCountAfter)
	h.VerifyMission(ctx, true, campaignId, missionCountBefore, name, description, missionType, weight, missionClaimDate)
}

func (h *C4eClaimUtils) AddMissionToCampaignError(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType cfeclaimtypes.MissionType,
	weight sdk.Dec, missionClaimDate *time.Time, errorString string) {
	missionCountBefore := h.helpeCfeclaimkeeper.GetMissionCount(ctx, campaignId)
	err := h.helpeCfeclaimkeeper.AddMissionToCampaign(ctx, owner, campaignId, name, description, missionType, weight, missionClaimDate)
	missionCountAfter := h.helpeCfeclaimkeeper.GetMissionCount(ctx, campaignId)
	require.EqualError(h.t, err, errorString)
	require.Equal(h.t, missionCountBefore, missionCountAfter)
}

func (h *C4eClaimUtils) VerifyCampaign(ctx sdk.Context, campaignId uint64, mustExist bool, owner string, name string, description string, enabled bool, feegrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration, vestingPoolName string) {

	claimCampaign, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)

	if mustExist {
		require.True(h.t, ok)
	} else {
		require.False(h.t, ok)
	}
	require.EqualValues(h.t, claimCampaign.Id, campaignId)
	require.EqualValues(h.t, claimCampaign.Owner, owner)
	require.EqualValues(h.t, claimCampaign.Name, name)
	require.EqualValues(h.t, claimCampaign.Description, description)
	require.EqualValues(h.t, claimCampaign.VestingPoolName, vestingPoolName)

	if feegrantAmount.IsNil() {
		require.EqualValues(h.t, claimCampaign.FeegrantAmount, sdk.ZeroInt())
	} else {
		require.True(h.t, claimCampaign.FeegrantAmount.Equal(*feegrantAmount))
	}

	if initialClaimFreeAmount.IsNil() {
		require.EqualValues(h.t, claimCampaign.InitialClaimFreeAmount, sdk.ZeroInt())
	} else {
		require.True(h.t, claimCampaign.InitialClaimFreeAmount.Equal(*initialClaimFreeAmount))
	}

	require.EqualValues(h.t, claimCampaign.Enabled, enabled)
	require.EqualValues(h.t, claimCampaign.StartTime, startTime)
	require.EqualValues(h.t, claimCampaign.EndTime, endTime)
	require.EqualValues(h.t, claimCampaign.VestingPeriod, vestingPeriod)
	require.EqualValues(h.t, claimCampaign.LockupPeriod, lockupPeriod)
}

func (h *C4eClaimUtils) VerifyCloseAction(ctx sdk.Context, campaignId uint64, CloseAction cfeclaimtypes.CloseAction, campaignAmountLeftBefore sdk.Coins) {

}

func (h *C4eClaimUtils) VerifyMission(ctx sdk.Context, mustExist bool, campaignId uint64, missionId uint64, name string, description string, missionType cfeclaimtypes.MissionType,
	weight sdk.Dec, claimStartDate *time.Time) {

	mission, ok := h.helpeCfeclaimkeeper.GetMission(ctx, campaignId, missionId)
	if mustExist {
		require.True(h.t, ok)
	} else {
		require.False(h.t, ok)
	}
	require.EqualValues(h.t, mission.MissionType, missionType)
	require.EqualValues(h.t, mission.CampaignId, campaignId)
	require.EqualValues(h.t, mission.Weight, weight)
	require.EqualValues(h.t, mission.Name, name)
	require.EqualValues(h.t, mission.Description, description)
	require.EqualValues(h.t, mission.Id, missionId)
	require.EqualValues(h.t, mission.ClaimStartDate, claimStartDate)
}
