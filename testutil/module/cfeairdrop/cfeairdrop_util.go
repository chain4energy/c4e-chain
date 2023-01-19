package cfeairdrop

import (
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"testing"
	"time"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
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
	BankUtils           *commontestutils.BankUtils
	StakingUtils        *commontestutils.StakingUtils
	GovUtils            *commontestutils.GovUtils
}

func NewC4eAirdropUtils(t *testing.T, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *commontestutils.BankUtils, stakingUtils *commontestutils.StakingUtils, govUtils *commontestutils.GovUtils) C4eAirdropUtils {
	return C4eAirdropUtils{C4eAirdropKeeperUtils: NewC4eAirdropKeeperUtils(t, helpeCfeairdropmodulekeeper),
		helperAccountKeeper: helperAccountKeeper, BankUtils: bankUtils, StakingUtils: stakingUtils, GovUtils: govUtils}
}

func (h *C4eAirdropUtils) SendToAirdropAccount(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, initialClaim bool) {
	coins := sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, amount))
	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	accBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)

	previousOriginalVesting := sdk.NewCoins()
	previousPeriods := []cfevestingtypes.ContinuousVestingPeriod{}
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
	require.NoError(h.t, h.helpeCfeairdropkeeper.SendToAirdropAccount(ctx,
		userAirdropEntries,
		coins,
		startTime,
		endTime, initialClaim,
	))

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance.Add(amount))
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance.Sub(amount))

	airdropAccount, ok := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	require.True(h.t, ok)
	newPeriods := append(previousPeriods, cfevestingtypes.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: coins})
	h.VerifyAirdropAccount(ctx, toAddress, previousOriginalVesting.Add(coins...), startTime, endTime, newPeriods, initialClaim)
	require.NoError(h.t, airdropAccount.Validate())
}

func (h *C4eAirdropUtils) SendToAirdropAccountError(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, initialClaim bool) {
	coins := sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, amount))
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
	require.EqualError(h.t, h.helpeCfeairdropkeeper.SendToAirdropAccount(ctx,
		userAirdropEntries,
		coins,
		startTime,
		endTime, createAccount,
	), errorMessage)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance)

	accountAfter := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	_, isAirdropAccount := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	_, ok := accountBefore.(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if ok && initialClaim {
		require.EqualValues(h.t, true, isAirdropAccount)
		h.VerifyAirdropAccount(ctx, toAddress, sdk.NewCoins(), startTime, endTime, []cfevestingtypes.ContinuousVestingPeriod{}, initialClaim)

	} else {
		require.EqualValues(h.t, wasAirdropAccount, isAirdropAccount)
		require.EqualValues(h.t, accountBefore, accountAfter)
	}

}

func (h *C4eAirdropUtils) VerifyAirdropAccount(ctx sdk.Context, address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfevestingtypes.ContinuousVestingPeriod, intitialClaim bool) {
	if intitialClaim && len(expectedOriginalVesting) > 0 {
		expectedOriginalVesting = expectedOriginalVesting.Sub(sdk.NewCoins(sdk.NewCoin(expectedOriginalVesting[0].Denom, cfeairdroptypes.OneToken)))
	}
	airdropAccount, ok := h.helperAccountKeeper.GetAccount(ctx, address).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	require.True(h.t, ok)

	require.EqualValues(h.t, len(expectedPeriods), len(airdropAccount.VestingPeriods))
	require.EqualValues(h.t, expectedStartTime, airdropAccount.StartTime)
	require.EqualValues(h.t, expectedEndTime, airdropAccount.EndTime)
	require.True(h.t, expectedOriginalVesting.IsEqual(airdropAccount.OriginalVesting))
	for i := 0; i < len(expectedPeriods); i++ {
		require.EqualValues(h.t, expectedPeriods[i].StartTime, airdropAccount.VestingPeriods[i].StartTime)
		require.EqualValues(h.t, expectedPeriods[i].EndTime, airdropAccount.VestingPeriods[i].EndTime)
		if intitialClaim {
			expectedPeriods[i].Amount = expectedPeriods[i].Amount.Sub(sdk.NewCoins(sdk.NewCoin(expectedOriginalVesting[0].Denom, cfeairdroptypes.OneToken)))
		}
		require.EqualValues(h.t, expectedPeriods[i].Amount, airdropAccount.VestingPeriods[i].Amount)
	}
	require.NoError(h.t, airdropAccount.Validate())

}

func (h *C4eAirdropUtils) AddAirdropEntries(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, airdropEntries []*cfeairdroptypes.AirdropEntry) {
	sum := sdk.ZeroInt()
	for _, airdropEntry := range airdropEntries {
		sum = sum.Add(airdropEntry.Amount)
	}
	usersAirdropEntriesBefore := h.helpeCfeairdropkeeper.GetUsersAirdropEntries(ctx)
	h.BankUtils.AddDefaultDenomCoinToAccount(ctx, sum, srcAddress)
	srcBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	require.NoError(h.t, h.helpeCfeairdropkeeper.AddUserAirdropEntries(ctx, srcAddress.String(), campaignId, airdropEntries))
	// TODO: add check if len(beforeAllAirdropEntry) + len(airdropEntries) == len(afterAllAirdropEntry)
	for _, airdropEntry := range airdropEntries {
		userAirdropEntries, found := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, airdropEntry.Address)
		require.True(h.t, found)
		var recordBefore *cfeairdroptypes.UserAirdropEntries = nil
		for _, before := range usersAirdropEntriesBefore {
			if before.Address == airdropEntry.Address {
				recordBefore = &before
				break
			}
		}
		if recordBefore == nil {
			require.EqualValues(h.t, 1, len(userAirdropEntries.AirdropEntries))
		} else {
			require.EqualValues(h.t, len(recordBefore.AirdropEntries)+1, len(userAirdropEntries.AirdropEntries))
		}
		require.EqualValues(h.t, airdropEntry.Address, userAirdropEntries.Address)
		require.EqualValues(h.t, "", userAirdropEntries.ClaimAddress)
		if recordBefore == nil {
			require.EqualValues(h.t, campaignId, userAirdropEntries.AirdropEntries[0].CampaignId)
			require.True(h.t, airdropEntry.Amount.Equal(userAirdropEntries.AirdropEntries[0].Amount))
			require.EqualValues(h.t, 0, len(userAirdropEntries.AirdropEntries[0].CompletedMissions))
		} else {
			expectedCaipaignRecords := append(recordBefore.AirdropEntries, &cfeairdroptypes.AirdropEntry{CampaignId: campaignId, Amount: airdropEntry.Amount, CompletedMissions: nil})

			require.ElementsMatch(h.t, expectedCaipaignRecords, userAirdropEntries.AirdropEntries)
		}

	}
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance.Add(sum))
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance.Sub(sum))

}

func (h *C4eAirdropUtils) AddCampaignRecordsError(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, airdropEntries []*cfeairdroptypes.AirdropEntry, errorMessage string, addRequiredCoinsToSrc bool) {
	if addRequiredCoinsToSrc {
		sum := sdk.ZeroInt()
		for _, airdropEntry := range airdropEntries {
			sum = sum.Add(airdropEntry.Amount)
		}
		h.BankUtils.AddDefaultDenomCoinToAccount(ctx, sum, srcAddress)
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
	if accExisted {
		require.True(h.t, ok)
	} else {
		claimerAccountBefore = nil
	}
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)

	userAirdropEntries, _ := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	err := h.helpeCfeairdropkeeper.InitialClaim(ctx, claimer.String(), campaignId, claimer.String())
	require.NoError(h.t, err)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, claimerBefore.AddRaw(expectedAmount))

	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore.SubRaw(expectedAmount))

	if claimerAccountBefore == nil {
		baseAccount := h.helperAccountKeeper.NewAccountWithAddress(ctx, claimer)
		claimerAccountBefore = cfevestingtypes.NewRepeatedContinuousVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), 100000000, 100000000, nil)
	}

	claimerAccountBefore = h.addExpectedDataToAccount(ctx, campaignId, claimerAccountBefore, sdk.NewInt(expectedAmount).Sub(cfeairdroptypes.OneToken))

	claimerAccount, ok := h.helperAccountKeeper.GetAccount(ctx, claimer).(*cfevestingtypes.RepeatedContinuousVestingAccount)
	if !accExisted {
		claimerAccountBefore.AccountNumber = claimerAccount.AccountNumber
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
	claimRecordBefore, foundBefore := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())

	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	//require.EqualError(h.t, h.helpeCfeairdropkeeper.ClaimInitial(ctx, campaignId, claimer.String()), errorMessage)
	// TODO: fix
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, sdk.ZeroInt())

	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))

	userAirdropEntries, found := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.Equal(h.t, foundBefore, found)
	if found {
		require.EqualValues(h.t, claimRecordBefore, userAirdropEntries)

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

func (h *C4eAirdropUtils) CompleteDelegationMission(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress, deleagtionAmount sdk.Int) {
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
	h.completeAnyMission(ctx, campaignId, 1, claimer, action, beforeCheck)
}

func (h *C4eAirdropUtils) CompleteVoteMission(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress) {
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
	h.completeAnyMission(ctx, campaignId, 2, claimer, action, nil)
}

func (h *C4eAirdropUtils) CompleteMission(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	action := func() error {
		return h.helpeCfeairdropkeeper.CompleteMission(ctx, campaignId, missionId, claimer.String(), false)
	}
	h.completeAnyMission(ctx, campaignId, missionId, claimer, action, nil)
}

func (h *C4eAirdropUtils) completeAnyMission(ctx sdk.Context, campaignId uint64, missionId uint64,
	claimer sdk.AccAddress, action func() error, beforeCheck func(before authtypes.AccountI, after authtypes.AccountI, ampountBefore sdk.Int) (authtypes.AccountI, sdk.Int)) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	claimRecordBefore, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
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
	claimRecordBefore.GetAidropEntry(campaignId).CompletedMissions = append(claimRecordBefore.GetAidropEntry(campaignId).CompletedMissions, missionId)
	require.EqualValues(h.t, claimRecordBefore, userAirdropEntries)
}

func (h *C4eAirdropUtils) CompleteMissionError(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountDefultDenomBalance(ctx, claimer)
	claimRecordBefore, foundCrBefore := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())

	require.EqualError(h.t, h.helpeCfeairdropkeeper.CompleteMission(ctx, campaignId, missionId, claimer.String(), false), errorMessage)

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
	claimRecordBefore, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.True(h.t, foundCr)

	require.NoError(h.t, h.helpeCfeairdropkeeper.Claim(ctx, campaignId, missionId, claimer.String()))

	claimRecordBefore.GetAidropEntry(campaignId).ClaimedMissions = append(claimRecordBefore.GetAidropEntry(campaignId).ClaimedMissions, missionId)
	userAirdropEntries, foundCr := h.helpeCfeairdropkeeper.GetUserAirdropEntries(ctx, claimer.String())
	require.True(h.t, foundCr)
	require.EqualValues(h.t, claimRecordBefore, userAirdropEntries)

	mission, _ := h.helpeCfeairdropkeeper.GetMission(ctx, campaignId, missionId)
	expectedAmount := mission.Weight.MulInt(userAirdropEntries.GetAidropEntry(campaignId).Amount).TruncateInt()

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
	expectedOriginalVesting := sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, expectedAmount))
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

func (h *C4eAirdropUtils) CreateAirdropCampaign(ctx sdk.Context, owner string, name string, description string, denom string, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) {

	err := h.helpeCfeairdropkeeper.CreateAidropCampaign(ctx, owner, name, description, denom, startTime, endTime, lockupPeriod, vestingPeriod)
	require.NoError(h.t, err)
}
func (h *C4eAirdropUtils) StartAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64) {
	err := h.helpeCfeairdropkeeper.StartAirdropCampaign(ctx, owner, campaignId)
	require.NoError(h.t, err)
}

func (h *C4eAirdropUtils) AddMissionToAirdropCampaign(ctx sdk.Context, owner string, campaignId uint64, name string, description string, missionType cfeairdroptypes.MissionType,
	weight sdk.Dec) {

	err := h.helpeCfeairdropkeeper.AddMissionToAirdropCampaign(ctx, owner, campaignId, name, description, missionType, weight, nil)
	require.NoError(h.t, err)
}
