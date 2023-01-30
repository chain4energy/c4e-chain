package cfeairdrop

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"testing"
	"time"

	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfeairdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"
)

type C4eAirdropUtils struct {
	C4eAirdropKeeperUtils
	helperAccountKeeper *authkeeper.AccountKeeper
	BankUtils           *testcosmos.BankUtils
	FeegrantUtils       *testcosmos.FeegrantUtils
	StakingUtils        *testcosmos.StakingUtils
	GovUtils            *testcosmos.GovUtils
}

func NewC4eAirdropUtils(t *testing.T, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *testcosmos.BankUtils, stakingUtils *testcosmos.StakingUtils, govUtils *testcosmos.GovUtils, feegrantUtils *testcosmos.FeegrantUtils) C4eAirdropUtils {
	return C4eAirdropUtils{C4eAirdropKeeperUtils: NewC4eAirdropKeeperUtils(t, helpeCfeairdropmodulekeeper),
		helperAccountKeeper: helperAccountKeeper, BankUtils: bankUtils, StakingUtils: stakingUtils, GovUtils: govUtils, FeegrantUtils: feegrantUtils}
}

func (h *C4eAirdropUtils) SendToAirdropAccount(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, missionType cfeairdroptypes.MissionType) {
	coins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amount))
	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	accBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)

	previousOriginalVesting := sdk.NewCoins()
	var previousPeriods []cfevestingtypes.ContinuousVestingPeriod
	if accountBefore != nil {
		if airdropAccount, ok := accountBefore.(*cfevestingtypes.RepeatedContinuousVestingAccount); ok {
			previousOriginalVesting = previousOriginalVesting.Add(airdropAccount.OriginalVesting...)
			previousPeriods = airdropAccount.VestingPeriods
		}
	}
	userAirdropEntries := &cfeairdroptypes.UserAirdropEntries{
		Address:      toAddress.String(),
		ClaimAddress: toAddress.String(),
	}

	require.NoError(h.t, h.helpeCfeairdropkeeper.SendToNewRepeatedContinuousVestingAccount(ctx,
		userAirdropEntries,
		coins,
		startTime,
		endTime,
		missionType,
	))

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance.Add(amount))
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance.Sub(amount))

	airdropAccount, ok := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	require.True(h.t, ok)
	newPeriods := append(previousPeriods, cfevestingtypes.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: coins})
	h.VerifyAirdropAccount(ctx, toAddress, previousOriginalVesting.Add(coins...), startTime, endTime, newPeriods, missionType)
	require.NoError(h.t, airdropAccount.Validate())
}

func (h *C4eAirdropUtils) SendToAirdropAccountError(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, missionType cfeairdroptypes.MissionType) {
	coins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amount))
	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	accBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	wasAirdropAccount := false
	if accountBefore != nil {
		_, wasAirdropAccount = accountBefore.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	}
	userAirdropEntries := &cfeairdroptypes.UserAirdropEntries{
		Address:      toAddress.String(),
		ClaimAddress: toAddress.String(),
	}
	require.EqualError(h.t, h.helpeCfeairdropkeeper.SendToNewRepeatedContinuousVestingAccount(ctx,
		userAirdropEntries,
		coins,
		startTime,
		endTime,
		missionType,
	), errorMessage)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance)

	accountAfter := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	_, isAirdropAccount := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	_, ok := accountBefore.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if ok && missionType == cfeairdroptypes.MissionInitialClaim {
		require.EqualValues(h.t, true, isAirdropAccount)
		h.VerifyAirdropAccount(ctx, toAddress, sdk.NewCoins(), startTime, endTime, []cfevestingtypes.ContinuousVestingPeriod{}, missionType)

	} else {
		require.EqualValues(h.t, wasAirdropAccount, isAirdropAccount)
		require.EqualValues(h.t, accountBefore, accountAfter)
	}

}

func (h *C4eAirdropUtils) VerifyAirdropAccount(ctx sdk.Context, address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfevestingtypes.ContinuousVestingPeriod, missionType cfeairdroptypes.MissionType) {

	airdropAccount, ok := h.helperAccountKeeper.GetAccount(ctx, address).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	require.True(h.t, ok)

	require.EqualValues(h.t, len(expectedPeriods), len(airdropAccount.VestingPeriods))
	require.EqualValues(h.t, expectedStartTime, airdropAccount.StartTime)
	require.EqualValues(h.t, expectedEndTime, airdropAccount.EndTime)
	require.True(h.t, expectedOriginalVesting.IsEqual(airdropAccount.OriginalVesting))
	for i := 0; i < len(expectedPeriods); i++ {
		require.EqualValues(h.t, expectedPeriods[i].StartTime, airdropAccount.VestingPeriods[i].StartTime)
		require.EqualValues(h.t, expectedPeriods[i].EndTime, airdropAccount.VestingPeriods[i].EndTime)
		require.EqualValues(h.t, expectedPeriods[i].Amount, airdropAccount.VestingPeriods[i].Amount)
	}
	require.NoError(h.t, airdropAccount.Validate())
}

func (h *C4eAirdropUtils) AddAirdropEntries(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, airdropEntries []*cfeairdroptypes.AirdropEntry) {
	airdropCoinsSum := sdk.NewCoins()
	for _, airdropEntry := range airdropEntries {
		airdropCoinsSum = airdropCoinsSum.Add(airdropEntry.AirdropCoins...)
	}
	usersAirdropEntriesBefore := h.helpeCfeairdropkeeper.GetUsersAirdropEntries(ctx)
	airdropClaimsLeftBefore, ok := h.helpeCfeairdropkeeper.GetAirdropClaimsLeft(ctx, campaignId)
	if !ok {
		airdropClaimsLeftBefore = cfeairdroptypes.AirdropClaimsLeft{
			AirdropCoins: sdk.NewCoins(),
			CampaignId:   campaignId,
		}
	}
	airdropDistrubitionsBefore, ok := h.helpeCfeairdropkeeper.GetAirdropDistrubitions(ctx, campaignId)
	if !ok {
		airdropDistrubitionsBefore = cfeairdroptypes.AirdropDistrubitions{
			AirdropCoins: sdk.NewCoins(),
			CampaignId:   campaignId,
		}
	}
	srcBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	require.NoError(h.t, h.helpeCfeairdropkeeper.AddUserAirdropEntries(ctx, srcAddress.String(), campaignId, airdropEntries))
	airdropClaimsLeftAfter, _ := h.helpeCfeairdropkeeper.GetAirdropClaimsLeft(ctx, campaignId)
	airdropDistrubitionsAfter, _ := h.helpeCfeairdropkeeper.GetAirdropDistrubitions(ctx, campaignId)
	airdropClaimsLeftBefore.AirdropCoins = airdropClaimsLeftBefore.AirdropCoins.Add(airdropCoinsSum...)
	airdropDistrubitionsBefore.AirdropCoins = airdropDistrubitionsBefore.AirdropCoins.Add(airdropCoinsSum...)
	require.EqualValues(h.t, airdropClaimsLeftBefore, airdropClaimsLeftAfter)
	require.EqualValues(h.t, airdropDistrubitionsBefore, airdropDistrubitionsAfter)

	// TODO: add check if len(beforeAllAirdropEntry) + len(airdropEntries) == len(afterAllAirdropEntry)
	for _, airdropEntry := range airdropEntries {
		userAirdropEntries, found := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, airdropEntry.Address)
		require.True(h.t, found)
		var userAirdropEntryBefore *cfeairdroptypes.UserAirdropEntries = nil
		for _, before := range usersAirdropEntriesBefore {
			if before.Address == airdropEntry.Address {
				userAirdropEntryBefore = &before
				break
			}
		}
		if userAirdropEntryBefore == nil {
			require.EqualValues(h.t, 1, len(userAirdropEntries.AirdropEntries))
		} else {
			require.EqualValues(h.t, len(userAirdropEntryBefore.AirdropEntries)+1, len(userAirdropEntries.AirdropEntries))
		}
		require.EqualValues(h.t, airdropEntry.Address, userAirdropEntries.Address)
		require.EqualValues(h.t, "", userAirdropEntries.ClaimAddress)
		if userAirdropEntryBefore == nil {
			require.EqualValues(h.t, campaignId, userAirdropEntries.AirdropEntries[0].CampaignId)
			require.True(h.t, airdropEntry.AirdropCoins.IsEqual(userAirdropEntries.AirdropEntries[0].AirdropCoins))
			require.EqualValues(h.t, 0, len(userAirdropEntries.AirdropEntries[0].CompletedMissions))
		} else {
			expectedCaipaignRecords := append(userAirdropEntryBefore.AirdropEntries, &cfeairdroptypes.AirdropEntry{CampaignId: campaignId, AirdropCoins: airdropEntry.AirdropCoins, CompletedMissions: nil})

			require.ElementsMatch(h.t, expectedCaipaignRecords, userAirdropEntries.AirdropEntries)
		}

	}
	campaign, _ := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignId)

	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		_, feegrandModuleAddress := cfeairdropmodulekeeper.FeegrantAccountAddress(campaignId)
		feegrantSum := campaign.FeegrantAmount.MulRaw(int64(len(airdropEntries)))
		h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance.Add(airdropCoinsSum.AmountOf(testenv.DefaultTestDenom)))
		h.BankUtils.VerifyAccountDefultDenomBalance(ctx, feegrandModuleAddress, feegrantSum)
		h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance.Sub(feegrantSum).Sub(airdropCoinsSum.AmountOf(testenv.DefaultTestDenom)))
	} else {
		h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance.Add(airdropCoinsSum.AmountOf(testenv.DefaultTestDenom)))
		h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance.Sub(airdropCoinsSum.AmountOf(testenv.DefaultTestDenom)))
	}
}

func (h *C4eAirdropUtils) AddAirdropEntriesError(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, airdropEntries []*cfeairdroptypes.AirdropEntry, errorMessage string) {
	ownerBalanceBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)
	moduleBalanceBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	err := h.helpeCfeairdropkeeper.AddUserAirdropEntries(ctx, srcAddress.String(), campaignId, airdropEntries)
	require.EqualError(h.t, err, errorMessage)

	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalanceBefore)
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, ownerBalanceBefore)
}

func (h *C4eAirdropUtils) AddCampaignRecordsError(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, airdropEntries []*cfeairdroptypes.AirdropEntry, errorMessage string, addRequiredCoinsToSrc bool) {
	if addRequiredCoinsToSrc {
		sum := sdk.NewCoins()
		for _, airdropEntry := range airdropEntries {
			sum = sum.Add(airdropEntry.AirdropCoins...)
		}
		h.BankUtils.AddDefaultDenomCoinsToAccount(ctx, sum.AmountOf(testenv.DefaultTestDenom), srcAddress)
	}
	usersAirdropEntriesBefore := h.helpeCfeairdropkeeper.GetUsersAirdropEntries(ctx)

	srcBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeairdropkeeper.AddUserAirdropEntries(ctx, srcAddress.String(), campaignId, airdropEntries), errorMessage)

	usersAirdropEntriesAfter := h.helpeCfeairdropkeeper.GetUsersAirdropEntries(ctx)
	require.ElementsMatch(h.t, usersAirdropEntriesBefore, usersAirdropEntriesAfter)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance)
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance)

}

func (h *C4eAirdropUtils) ClaimInitial(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress, expectedAmount int64) {
	acc := h.helperAccountKeeper.GetAccount(ctx, claimer)
	claimerAccountBefore, ok := acc.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	accExisted := acc != nil
	if !accExisted {
		claimerAccountBefore = nil
	}
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	campaign, _ := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignId)
	airdropClaimsLeftBefore, ok := h.helpeCfeairdropkeeper.GetAirdropClaimsLeft(ctx, campaignId)

	userAirdropEntries, _ := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	_, granterAddr := cfeairdropmodulekeeper.FeegrantAccountAddress(campaignId)
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		allowance, err := h.FeegrantUtils.FeegrantKeeper.GetAllowance(ctx, granterAddr, claimer)
		require.NoError(h.t, err)
		require.NotNil(h.t, allowance)
	}
	err := h.helpeCfeairdropkeeper.InitialClaim(ctx, claimer.String(), campaignId, claimer.String())
	require.NoError(h.t, err)
	allowance, err := h.FeegrantUtils.FeegrantKeeper.GetAllowance(ctx, granterAddr, claimer)
	require.Error(h.t, err)
	require.Nil(h.t, allowance)
	airdropClaimsLeftAfter, ok := h.helpeCfeairdropkeeper.GetAirdropClaimsLeft(ctx, campaignId)
	airdropClaimsLeftBefore.AirdropCoins = airdropClaimsLeftBefore.AirdropCoins.Sub(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, sdk.NewInt(expectedAmount))))

	require.EqualValues(h.t, airdropClaimsLeftBefore, airdropClaimsLeftAfter)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore.AddRaw(expectedAmount))
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore.SubRaw(expectedAmount))

	if claimerAccountBefore == nil {
		baseAccount := h.helperAccountKeeper.NewAccountWithAddress(ctx, claimer)
		claimerAccountBefore = cfevestingtypes.NewRepeatedContinuousVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), 100000000, 100000000, nil)
	}

	vestingAmount := sdk.NewInt(expectedAmount)
	if campaign.InitialClaimFreeAmount.GT(sdk.ZeroInt()) {
		vestingAmount = sdk.NewInt(expectedAmount).Sub(campaign.InitialClaimFreeAmount)
	}
	claimerAccountBefore = h.addExpectedDataToAccount(ctx, campaignId, claimerAccountBefore, vestingAmount)

	claimerAccount, ok := h.helperAccountKeeper.GetAccount(ctx, claimer).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if !accExisted {
		claimerAccountBefore.AccountNumber = claimerAccount.AccountNumber
	}
	if campaign.FeegrantAmount.GT(sdk.ZeroInt()) {
		claimerAccount.AccountNumber++
	}
	require.True(h.t, ok)
	require.NoError(h.t, claimerAccount.Validate())
	require.EqualValues(h.t, claimerAccountBefore, claimerAccount)

	userAirdropEntries, found := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.True(h.t, found)
	airdropEntry := userAirdropEntries.GetAidropEntry(campaignId)
	require.NotNil(h.t, airdropEntry)
}

func (h *C4eAirdropUtils) ClaimInitialError(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	balanceBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	userAirdropEntriesBefore, foundBefore := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())

	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeairdropkeeper.InitialClaim(ctx, claimer.String(), campaignId, claimer.String()), errorMessage)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, balanceBefore)

	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))

	userAirdropEntries, found := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.Equal(h.t, foundBefore, found)
	if found {
		require.EqualValues(h.t, userAirdropEntriesBefore, userAirdropEntries)

	}
}

func (h *C4eAirdropUtils) GetUserAirdropEntries(
	ctx sdk.Context,
	address string,
) *cfeairdroptypes.UserAirdropEntries {
	val, found := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, address)
	if found {
		return &val
	}
	return nil
}

func (h *C4eAirdropUtils) SetUserAirdropEntries(
	ctx sdk.Context,
	userAirdropEntries *cfeairdroptypes.UserAirdropEntries,
) {
	h.helpeCfeairdropkeeper.SetUserAirdropEntries(ctx, *userAirdropEntries)
}

func (h *C4eAirdropUtils) CompleteDelegationMission(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, deleagtionAmount sdk.Int) {
	action := func() error {
		validators := h.StakingUtils.StakingKeeper.GetValidators(ctx, 1)
		valAddr, err := sdk.ValAddressFromBech32(validators[0].OperatorAddress)
		if err != nil {
			return err
		}
		h.StakingUtils.MessageDelegate(ctx, 1, 0, valAddr, claimer, deleagtionAmount)
		return nil
	}
	beforeCheck := func(accBefore authtypes.AccountI, accAfter authtypes.AccountI, claimerAmountBefore sdk.Int) (authtypes.AccountI, sdk.Int) {
		veBefore, okBefore := accBefore.(*cfevestingtypes.RepeatedContinuousVestingAccount)
		veAfter, okAfter := accAfter.(*cfevestingtypes.RepeatedContinuousVestingAccount)
		if okBefore && okAfter {
			veBefore.DelegatedFree = veAfter.DelegatedFree
			veBefore.DelegatedVesting = veAfter.DelegatedVesting
		}
		return veBefore, claimerAmountBefore.Sub(deleagtionAmount)
	}
	h.completeAnyMission(ctx, campaignId, missionId, claimer, action, beforeCheck)
}

func (h *C4eAirdropUtils) CompleteVoteMission(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	action := func() error {
		depParams := h.GovUtils.GovKeeper.GetDepositParams(ctx)
		depositAmount := depParams.MinDeposit
		h.BankUtils.AddCoinsToAccount(ctx, depositAmount, claimer)

		testProposal := &govtypes.TextProposal{Title: "Title", Description: "Description"}
		proposal, err := h.GovUtils.GovKeeper.SubmitProposal(ctx, testProposal)
		if err != nil {
			return err
		}
		h.GovUtils.GovKeeper.AddDeposit(ctx, proposal.ProposalId, claimer, depositAmount)

		return h.GovUtils.GovKeeper.AddVote(ctx, proposal.ProposalId,
			claimer, []govtypes.WeightedVoteOption{{Option: govtypes.OptionAbstain, Weight: sdk.NewDec(1)}})
	}
	h.completeAnyMission(ctx, campaignId, missionId, claimer, action, nil)
}

func (h *C4eAirdropUtils) CompleteMissionFromHook(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	action := func() error {
		return h.helpeCfeairdropkeeper.CompleteMissionFromHook(ctx, campaignId, missionId, claimer.String())
	}
	h.completeAnyMission(ctx, campaignId, missionId, claimer, action, nil)
}

func (h *C4eAirdropUtils) completeAnyMission(ctx sdk.Context, campaignId uint64, missionId uint64,
	claimer sdk.AccAddress, action func() error, beforeCheck func(before authtypes.AccountI, after authtypes.AccountI, ampountBefore sdk.Int) (authtypes.AccountI, sdk.Int)) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	userAirdropEntriesBefore, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.True(h.t, foundCr)

	require.NoError(h.t, action())
	claimerAccountAfter := h.helperAccountKeeper.GetAccount(ctx, claimer)
	if beforeCheck != nil {
		claimerAccountBefore, claimerBefore = beforeCheck(claimerAccountBefore, claimerAccountAfter, claimerBefore)
	}
	require.EqualValues(h.t, claimerAccountBefore, claimerAccountAfter)
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore)

	userAirdropEntries, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.True(h.t, foundCr)
	userAirdropEntriesBefore.GetAidropEntry(campaignId).CompletedMissions = append(userAirdropEntriesBefore.GetAidropEntry(campaignId).CompletedMissions, missionId)
	require.EqualValues(h.t, userAirdropEntriesBefore, userAirdropEntries)
}

func (h *C4eAirdropUtils) CompleteMissionFromHookError(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	claimRecordBefore, foundCrBefore := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())

	require.EqualError(h.t, h.helpeCfeairdropkeeper.CompleteMissionFromHook(ctx, campaignId, missionId, claimer.String()), errorMessage)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore)
	userAirdropEntries, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.Equal(h.t, foundCrBefore, foundCr)
	require.EqualValues(h.t, claimRecordBefore, userAirdropEntries)
}

func (h *C4eAirdropUtils) ClaimMission(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.ClaimMissionToAddress(ctx, campaignId, missionId, claimer, claimer)
}

func (h *C4eAirdropUtils) ClaimMissionToAddress(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, claimerDstAddress sdk.AccAddress) {
	claimerAccountBefore, ok := h.helperAccountKeeper.GetAccount(ctx, claimerDstAddress).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	require.True(h.t, ok)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimerDstAddress)
	userAirdropEntriesBefore, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.True(h.t, foundCr)
	mission, _ := h.helpeCfeairdropkeeper.GetMission(ctx, campaignId, missionId)
	require.NoError(h.t, h.helpeCfeairdropkeeper.Claim(ctx, campaignId, missionId, claimer.String()))

	userAirdropEntriesBefore.GetAidropEntry(campaignId).ClaimedMissions = append(userAirdropEntriesBefore.GetAidropEntry(campaignId).ClaimedMissions, missionId)
	if mission.MissionType == cfeairdroptypes.MissionClaim {
		userAirdropEntriesBefore.GetAidropEntry(campaignId).CompletedMissions = append(userAirdropEntriesBefore.GetAidropEntry(campaignId).CompletedMissions, missionId)
	}
	userAirdropEntries, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.True(h.t, foundCr)

	require.EqualValues(h.t, userAirdropEntriesBefore, userAirdropEntries)

	expectedAmount := mission.Weight.MulInt(userAirdropEntries.GetAidropEntry(campaignId).AirdropCoins.AmountOf(testenv.DefaultTestDenom)).TruncateInt()

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimerDstAddress, claimerBefore.Add(expectedAmount))
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore.Sub(expectedAmount))

	h.addExpectedDataToAccount(ctx, campaignId, claimerAccountBefore, expectedAmount)

	claimerAccount, ok := h.helperAccountKeeper.GetAccount(ctx, claimerDstAddress).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	require.True(h.t, ok)
	require.NoError(h.t, claimerAccount.Validate())
	require.EqualValues(h.t, claimerAccountBefore, claimerAccount)
}

func (h *C4eAirdropUtils) addExpectedDataToAccount(ctx sdk.Context, campaignId uint64,
	claimerAccountBefore *cfevestingtypes.RepeatedContinuousVestingAccount, expectedAmount sdk.Int) *cfevestingtypes.RepeatedContinuousVestingAccount {
	campaign, _ := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignId)
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

func (h *C4eAirdropUtils) ClaimMissionError(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	claimRecordBefore, foundCrBefore := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())

	require.EqualError(h.t, h.helpeCfeairdropkeeper.Claim(ctx, campaignId, missionId, claimer.String()), errorMessage)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore)
	userAirdropEntries, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.Equal(h.t, foundCrBefore, foundCr)
	require.EqualValues(h.t, claimRecordBefore, userAirdropEntries)
}

func (h *C4eAirdropUtils) CreateAirdropAccout(ctx sdk.Context, address sdk.AccAddress, originalVesting sdk.Coins, startTime int64, endTime int64, periods ...cfevestingtypes.ContinuousVestingPeriod) *cfevestingtypes.RepeatedContinuousVestingAccount {
	baseAccount := h.helperAccountKeeper.NewAccountWithAddress(ctx, address)
	airdropAcc := cfevestingtypes.NewRepeatedContinuousVestingAccount(baseAccount.(*authtypes.BaseAccount), originalVesting, startTime, endTime, periods)
	h.helperAccountKeeper.SetAccount(ctx, airdropAcc)
	require.NoError(h.t, airdropAcc.Validate())
	return airdropAcc
}

func (h *C4eAirdropUtils) CreateAirdropCampaign(ctx sdk.Context, owner string, name string, description string, feegrantAmount sdk.Int, initialClaimFreeAmount sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) {

	campaignCountBefore := h.helpeCfeairdropkeeper.GetCampaignCount(ctx)
	err := h.helpeCfeairdropkeeper.CreateAidropCampaign(ctx, owner, name, description, &feegrantAmount, &initialClaimFreeAmount, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	missionCountAfter := h.helpeCfeairdropkeeper.GetMissionCount(ctx, campaignCountBefore)
	require.NoError(h.t, err)
	campaignCountAfter := h.helpeCfeairdropkeeper.GetCampaignCount(ctx)
	require.Equal(h.t, campaignCountBefore+1, campaignCountAfter)
	require.Equal(h.t, uint64(1), missionCountAfter)

	h.VerifyAirdropCampaign(ctx, campaignCountBefore, true, owner, name, description, false, &feegrantAmount, &initialClaimFreeAmount, startTime, endTime, lockupPeriod, vestingPeriod)
	h.VerifyAirdropMission(ctx, true, campaignCountBefore, 0, "Initial mission", "Initial mission - basic mission that must be claimed first", cfeairdroptypes.MissionInitialClaim, sdk.ZeroDec(), nil)
}

func (h *C4eAirdropUtils) CreateAirdropCampaignError(ctx sdk.Context, owner string, name string, description string, feegrantAmount sdk.Int, initialClaimFreeAmount sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration, errorMessage string) {

	campaignCountBefore := h.helpeCfeairdropkeeper.GetCampaignCount(ctx)
	err := h.helpeCfeairdropkeeper.CreateAidropCampaign(ctx, owner, name, description, &feegrantAmount, &initialClaimFreeAmount, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	require.EqualError(h.t, err, errorMessage)
	campaignCountAfter := h.helpeCfeairdropkeeper.GetCampaignCount(ctx)
	missionCountAfter := h.helpeCfeairdropkeeper.GetMissionCount(ctx, campaignCountBefore)
	require.Equal(h.t, campaignCountBefore, campaignCountAfter)
	require.Equal(h.t, uint64(0), missionCountAfter)
	_, ok := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignCountBefore)
	require.False(h.t, ok)
}

func (h *C4eAirdropUtils) StartAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64) {
	err := h.helpeCfeairdropkeeper.StartAirdropCampaign(ctx, owner, campaignId)
	require.NoError(h.t, err)
	campaign, ok := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignId)
	require.True(h.t, ok)
	h.VerifyAirdropCampaign(ctx, campaign.Id, true, owner, campaign.Name, campaign.Description, true, &campaign.FeegrantAmount, &campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
}

func (h *C4eAirdropUtils) StartAirdropCampaignError(ctx sdk.Context, owner string, campaignId uint64, errorString string) {
	campaignBefore, ok := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignId)

	err := h.helpeCfeairdropkeeper.StartAirdropCampaign(ctx, owner, campaignId)
	require.EqualError(h.t, err, errorString)
	if !ok {
		return
	}
	enabled := campaignBefore.Enabled
	h.VerifyAirdropCampaign(ctx, campaignBefore.Id, true, campaignBefore.Owner, campaignBefore.Name, campaignBefore.Description, enabled, &campaignBefore.FeegrantAmount, &campaignBefore.InitialClaimFreeAmount, campaignBefore.StartTime, campaignBefore.EndTime, campaignBefore.LockupPeriod, campaignBefore.VestingPeriod)
}

func (h *C4eAirdropUtils) CloseAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64, airdropCloseAction cfeairdroptypes.AirdropCloseAction) {
	err := h.helpeCfeairdropkeeper.CloseAirdropCampaign(ctx, owner, campaignId, airdropCloseAction)
	require.NoError(h.t, err)
	campaign, ok := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignId)
	require.True(h.t, ok)
	h.VerifyAirdropCampaign(ctx, campaign.Id, true, owner, campaign.Name, campaign.Description, false, &campaign.FeegrantAmount, &campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
}

func (h *C4eAirdropUtils) CloseAirdropCampaignError(ctx sdk.Context, owner string, campaignId uint64, airdropCloseAction cfeairdroptypes.AirdropCloseAction, errorString string) {
	campaignBefore, ok := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignId)

	err := h.helpeCfeairdropkeeper.CloseAirdropCampaign(ctx, owner, campaignId, airdropCloseAction)
	require.EqualError(h.t, err, errorString)
	if !ok {
		return
	}
	enabled := campaignBefore.Enabled
	h.VerifyAirdropCampaign(ctx, campaignBefore.Id, true, campaignBefore.Owner, campaignBefore.Name, campaignBefore.Description, enabled, &campaignBefore.FeegrantAmount, &campaignBefore.InitialClaimFreeAmount, campaignBefore.StartTime, campaignBefore.EndTime, campaignBefore.LockupPeriod, campaignBefore.VestingPeriod)
}

func (h *C4eAirdropUtils) AddMissionToAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType cfeairdroptypes.MissionType,
	weight sdk.Dec, missionClaimDate *time.Time) {
	missionCountBefore := h.helpeCfeairdropkeeper.GetMissionCount(ctx, campaignId)
	err := h.helpeCfeairdropkeeper.AddMissionToAirdropCampaign(ctx, owner, campaignId, name, description, missionType, weight, nil)
	missionCountAfter := h.helpeCfeairdropkeeper.GetMissionCount(ctx, campaignId)
	require.NoError(h.t, err)
	require.Equal(h.t, missionCountBefore+1, missionCountAfter)
	h.VerifyAirdropMission(ctx, true, campaignId, missionCountBefore, name, description, missionType, weight, missionClaimDate)
}

func (h *C4eAirdropUtils) AddMissionToAirdropCampaignError(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType cfeairdroptypes.MissionType,
	weight sdk.Dec, missionClaimDate *time.Time, errorString string) {
	missionCountBefore := h.helpeCfeairdropkeeper.GetMissionCount(ctx, campaignId)
	err := h.helpeCfeairdropkeeper.AddMissionToAirdropCampaign(ctx, owner, campaignId, name, description, missionType, weight, missionClaimDate)
	missionCountAfter := h.helpeCfeairdropkeeper.GetMissionCount(ctx, campaignId)
	require.EqualError(h.t, err, errorString)
	require.Equal(h.t, missionCountBefore, missionCountAfter)
}

func (h *C4eAirdropUtils) VerifyAirdropCampaign(ctx sdk.Context, campaignId uint64, mustExist bool, owner string, name string, description string, enabled bool, feegrantAmount *sdk.Int, initialClaimFreeAmount *sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) {

	airdropCampaign, ok := h.helpeCfeairdropkeeper.GetCampaign(ctx, campaignId)

	if mustExist {
		require.True(h.t, ok)
	} else {
		require.False(h.t, ok)
	}
	require.EqualValues(h.t, airdropCampaign.Id, campaignId)
	require.EqualValues(h.t, airdropCampaign.Owner, owner)
	require.EqualValues(h.t, airdropCampaign.Name, name)
	require.EqualValues(h.t, airdropCampaign.Description, description)

	if feegrantAmount.IsNil() {
		require.EqualValues(h.t, airdropCampaign.FeegrantAmount, sdk.ZeroInt())
	} else {
		require.True(h.t, airdropCampaign.FeegrantAmount.Equal(*feegrantAmount))
	}

	if initialClaimFreeAmount.IsNil() {
		require.EqualValues(h.t, airdropCampaign.InitialClaimFreeAmount, sdk.ZeroInt())
	} else {
		require.True(h.t, airdropCampaign.InitialClaimFreeAmount.Equal(*initialClaimFreeAmount))
	}

	require.EqualValues(h.t, airdropCampaign.Enabled, enabled)
	require.EqualValues(h.t, airdropCampaign.StartTime, startTime)
	require.EqualValues(h.t, airdropCampaign.EndTime, endTime)
	require.EqualValues(h.t, airdropCampaign.VestingPeriod, vestingPeriod)
	require.EqualValues(h.t, airdropCampaign.LockupPeriod, lockupPeriod)
}

func (h *C4eAirdropUtils) VerifyAirdropMission(ctx sdk.Context, mustExist bool, campaignId uint64, missionId uint64, name string, description string, missionType cfeairdroptypes.MissionType,
	weight sdk.Dec, claimStartDate *time.Time) {

	mission, ok := h.helpeCfeairdropkeeper.GetMission(ctx, campaignId, missionId)
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
