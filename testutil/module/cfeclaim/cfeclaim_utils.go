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

func (h *C4eClaimUtils) AddClaimRecords(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, claimRecords []*cfeclaimtypes.ClaimRecord) {
	amountSum := sdk.NewCoins()
	for _, claimRecord := range claimRecords {
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
	srcBalance := h.BankUtils.GetAccountAllBalances(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	campaign, _ := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)

	userEntriesBefore := h.helpeCfeclaimkeeper.GetUsersEntries(ctx)
	claimRecordBeforeCount := 0
	for _, userEntry := range userEntriesBefore {
		for _, claimRecord := range userEntry.ClaimRecords {
			if claimRecord.CampaignId == campaignId {
				claimRecordBeforeCount++
			}
		}
	}
	require.NoError(h.t, h.helpeCfeclaimkeeper.AddClaimRecords(ctx, srcAddress.String(), campaignId, claimRecords))
	claimClaimsLeftAfter, _ := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	claimDistrubitionsAfter, _ := h.helpeCfeclaimkeeper.GetCampaignTotalAmount(ctx, campaignId)
	claimClaimsLeftBefore.Amount = claimClaimsLeftBefore.Amount.Add(amountSum...)
	claimDistrubitionsBefore.Amount = claimDistrubitionsBefore.Amount.Add(amountSum...)
	require.EqualValues(h.t, claimClaimsLeftBefore, claimClaimsLeftAfter)
	require.EqualValues(h.t, claimDistrubitionsBefore, claimDistrubitionsAfter)

	userEntriesAfter := h.helpeCfeclaimkeeper.GetUsersEntries(ctx)
	userEntryAfterCount := 0
	for _, userEntry := range userEntriesAfter {
		for _, claimRecord := range userEntry.ClaimRecords {
			if claimRecord.CampaignId == campaignId {
				userEntryAfterCount++
			}
		}
	}
	require.EqualValues(h.t, claimRecordBeforeCount+len(claimRecords), userEntryAfterCount)

	for _, claimRecord := range claimRecords {
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

	if campaign.FeegrantAmount.GT(math.ZeroInt()) {
		_, feegrandModuleAddress := cfeclaimmodulekeeper.CreateFeegrantAccountAddress(campaignId)
		feegrantSum := campaign.FeegrantAmount.MulRaw(int64(len(claimRecords)))
		feegrantCoins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, feegrantSum))
		h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBalance.Add(amountSum...))
		h.BankUtils.VerifyAccountAllBalances(ctx, feegrandModuleAddress, feegrantCoins)
		h.BankUtils.VerifyAccountAllBalances(ctx, srcAddress, srcBalance.Sub(feegrantCoins...).Sub(amountSum...))
	} else {
		if campaign.CampaignType != cfeclaimtypes.VestingPoolCampaign {
			h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBalance.Add(amountSum...))
			h.BankUtils.VerifyAccountAllBalances(ctx, srcAddress, srcBalance.Sub(amountSum...))
		} else {
			if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
				_, vestingPool, accFound := h.helperCfevestingKeeper.GetAccountVestingPool(ctx, srcAddress.String(), campaign.VestingPoolName)
				require.True(h.t, accFound)
				require.EqualValues(h.t, amountSum.AmountOf(testenv.DefaultTestDenom), vestingPool.GetReservation(campaignId).Amount)
			}
		}

	}
}

func (h *C4eClaimUtils) DeleteClaimRecord(ctx sdk.Context, ownerAddress sdk.AccAddress, campaignId uint64, userAddress string, amoutDiff sdk.Coins) {
	claimClaimsLeftBefore, ok := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	if !ok {
		claimClaimsLeftBefore = cfeclaimtypes.CampaignAmountLeft{
			Amount:     sdk.NewCoins(),
			CampaignId: campaignId,
		}
	}

	srcBalance := h.BankUtils.GetAccountAllBalances(ctx, ownerAddress)
	moduleBalanceBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	campaign, _ := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)

	var vestingPoolBefore *cfevestingtypes.VestingPool
	if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
		var accFound bool
		_, vestingPoolBefore, accFound = h.helperCfevestingKeeper.GetAccountVestingPool(ctx, ownerAddress.String(), campaign.VestingPoolName)
		require.True(h.t, accFound)
	}
	userEntryBefore, found := h.helpeCfeclaimkeeper.GetUserEntry(ctx, userAddress)
	require.True(h.t, found)
	found = false
	for _, claimRecord := range userEntryBefore.ClaimRecords {
		require.NotEqual(h.t, claimRecord.CampaignId, campaignId)
	}
	require.NoError(h.t, h.helpeCfeclaimkeeper.DeleteClaimRecord(ctx, ownerAddress.String(), campaignId, userAddress))
	claimClaimsLeftAfter, _ := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	require.EqualValues(h.t, claimClaimsLeftBefore, claimClaimsLeftAfter.Amount.Add(amoutDiff...))

	userEntry, found := h.helpeCfeclaimkeeper.GetUserEntry(ctx, userAddress)
	require.True(h.t, found)
	for _, claimRecord := range userEntry.ClaimRecords {
		require.NotEqual(h.t, claimRecord.CampaignId, campaignId)
	}
	moduleBalanceAfter := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	require.Equal(h.t, moduleBalanceAfter, moduleBalanceBefore.Add(amoutDiff...))
	if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
		_, vestingPool, _ := h.helperCfevestingKeeper.GetAccountVestingPool(ctx, ownerAddress.String(), campaign.VestingPoolName)
		require.True(h.t, vestingPoolBefore.Sent.Sub(amoutDiff.AmountOf(testenv.DefaultTestDenom)).LT(vestingPool.Sent))
		// TODO: add feegrant verification
	} else {
		h.BankUtils.VerifyAccountAllBalances(ctx, ownerAddress, srcBalance.Add(amoutDiff...))
	}
}

func (h *C4eClaimUtils) DeleteClaimRecordError(ctx sdk.Context, ownerAddress sdk.AccAddress, campaignId uint64, userAddress string, errorMessage string) {
	srcBalance := h.BankUtils.GetAccountAllBalances(ctx, ownerAddress)
	moduleBalanceBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeclaimkeeper.DeleteClaimRecord(ctx, ownerAddress.String(), campaignId, userAddress), errorMessage)
	h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBalanceBefore)
	h.BankUtils.VerifyAccountAllBalances(ctx, ownerAddress, srcBalance)
}

func (h *C4eClaimUtils) AddClaimRecordsError(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, claimEntries []*cfeclaimtypes.ClaimRecord, errorMessage string) {
	ownerBalanceBefore := h.BankUtils.GetAccountAllBalances(ctx, srcAddress)
	moduleBalanceBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	err := h.helpeCfeclaimkeeper.AddClaimRecords(ctx, srcAddress.String(), campaignId, claimEntries)
	require.EqualError(h.t, err, errorMessage)

	h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBalanceBefore)
	h.BankUtils.VerifyAccountAllBalances(ctx, srcAddress, ownerBalanceBefore)
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
		h.BankUtils.AddCoinsToAccount(ctx, sum, srcAddress)
	}
	usersEntriesBefore := h.helpeCfeclaimkeeper.GetUsersEntries(ctx)

	srcBalance := h.BankUtils.GetAccountAllBalances(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeclaimkeeper.AddClaimRecords(ctx, srcAddress.String(), campaignId, claimEntries), errorMessage)

	usersEntriesAfter := h.helpeCfeclaimkeeper.GetUsersEntries(ctx)
	require.ElementsMatch(h.t, usersEntriesBefore, usersEntriesAfter)
	h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBalance)
	h.BankUtils.VerifyAccountAllBalances(ctx, srcAddress, srcBalance)

}

func (h *C4eClaimUtils) ClaimInitial(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress, expectedAmount int64) {
	acc := h.helperAccountKeeper.GetAccount(ctx, claimer)
	claimerAccountBefore, ok := acc.(*cfevestingtypes.PeriodicContinuousVestingAccount)
	accExisted := acc != nil
	if !accExisted {
		claimerAccountBefore = nil
	}
	moduleBalanceBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	claimerBalanceBefore := h.BankUtils.GetAccountAllBalances(ctx, claimer)
	campaign, _ := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	claimClaimsLeftBefore, ok := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)

	userEntry, _ := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	_, granterAddr := cfeclaimmodulekeeper.CreateFeegrantAccountAddress(campaignId)
	if campaign.FeegrantAmount.GT(math.ZeroInt()) {
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
	claimClaimsLeftBefore.Amount = claimClaimsLeftBefore.Amount.Sub(sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(expectedAmount)))...)

	require.EqualValues(h.t, claimClaimsLeftBefore, claimClaimsLeftAfter)

	for _, coin := range claimerBalanceBefore {
		h.BankUtils.VerifyAccountBalanceByDenom(ctx, claimer, coin.Denom, coin.Amount.AddRaw(expectedAmount))
	}
	for _, coin := range moduleBalanceBefore {
		h.BankUtils.VerifyModuleAccountBalanceByDenom(ctx, cfeclaimtypes.ModuleName, coin.Denom, coin.Amount.SubRaw(expectedAmount))
	}

	if claimerAccountBefore == nil {
		baseAccount := h.helperAccountKeeper.NewAccountWithAddress(ctx, claimer)
		claimerAccountBefore = cfevestingtypes.NewRepeatedContinuousVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(), 100000000, 100000000, nil)
	}

	vestingAmount := math.NewInt(expectedAmount)
	if campaign.InitialClaimFreeAmount.GT(math.ZeroInt()) {
		vestingAmount = math.NewInt(expectedAmount).Sub(campaign.InitialClaimFreeAmount)
	}
	vestingCoins := sdk.NewCoins(sdk.NewCoin(h.helperCfevestingKeeper.GetParams(ctx).Denom, vestingAmount))
	claimerAccountBefore = h.addExpectedDataToAccount(ctx, campaignId, claimerAccountBefore, vestingCoins)

	claimerAccount, ok := h.helperAccountKeeper.GetAccount(ctx, claimer).(*cfevestingtypes.PeriodicContinuousVestingAccount)

	claimerAccountBefore.AccountNumber = claimerAccount.AccountNumber

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
	balanceBefore := h.BankUtils.GetAccountAllBalances(ctx, claimer)
	userEntryBefore, foundBefore := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())

	moduleBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeclaimkeeper.InitialClaim(ctx, claimer.String(), campaignId, claimer.String()), errorMessage)

	h.BankUtils.VerifyAccountAllBalances(ctx, claimer, balanceBefore)

	h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBefore)

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
		h.StakingUtils.SetupValidators(ctx, []sdk.ValAddress{valAddress}, math.NewInt(1))
		h.StakingUtils.MessageDelegate(ctx, 2, 0, valAddress, claimer, deleagtionAmount)
		return nil
	}
	beforeCheck := func(accBefore authtypes.AccountI, accAfter authtypes.AccountI, claimerAmountBefore sdk.Coins) (authtypes.AccountI, sdk.Coins) {
		veBefore, okBefore := accBefore.(*cfevestingtypes.PeriodicContinuousVestingAccount)
		veAfter, okAfter := accAfter.(*cfevestingtypes.PeriodicContinuousVestingAccount)
		if okBefore && okAfter {
			veBefore.DelegatedFree = veAfter.DelegatedFree
			veBefore.DelegatedVesting = veAfter.DelegatedVesting
		}
		return veBefore, claimerAmountBefore.Sub(sdk.NewCoin(h.StakingUtils.GetStakingDenom(ctx), deleagtionAmount))
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
	claimer sdk.AccAddress, action func() error, beforeCheck func(before authtypes.AccountI, after authtypes.AccountI, ampountBefore sdk.Coins) (authtypes.AccountI, sdk.Coins)) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountAllBalances(ctx, claimer)
	userEntryBefore, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.True(h.t, foundCr)

	require.NoError(h.t, action())
	claimerAccountAfter := h.helperAccountKeeper.GetAccount(ctx, claimer)
	if beforeCheck != nil {
		claimerAccountBefore, claimerBefore = beforeCheck(claimerAccountBefore, claimerAccountAfter, claimerBefore)
	}
	require.EqualValues(h.t, claimerAccountBefore, claimerAccountAfter)
	h.BankUtils.VerifyAccountAllBalances(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBefore)

	userEntry, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.True(h.t, foundCr)
	userEntryBefore.GetClaimRecord(campaignId).CompletedMissions = append(userEntryBefore.GetClaimRecord(campaignId).CompletedMissions, missionId)
	require.EqualValues(h.t, userEntryBefore, userEntry)
}

func (h *C4eClaimUtils) CompleteMissionFromHookError(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	moduleBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountAllBalances(ctx, claimer)
	claimRecordBefore, foundCrBefore := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())

	require.EqualError(h.t, h.helpeCfeclaimkeeper.CompleteMissionFromHook(ctx, campaignId, missionId, claimer.String()), errorMessage)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))
	h.BankUtils.VerifyAccountAllBalances(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBefore)
	userEntry, foundCr := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())
	require.Equal(h.t, foundCrBefore, foundCr)
	require.EqualValues(h.t, claimRecordBefore, userEntry)
}

func (h *C4eClaimUtils) ClaimMission(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.ClaimMissionToAddress(ctx, campaignId, missionId, claimer, claimer)
}

func (h *C4eClaimUtils) ClaimMissionToAddress(ctx sdk.Context, campaignId uint64, missionId uint64, claimer sdk.AccAddress, claimerDstAddress sdk.AccAddress) {
	claimerAccountBefore, ok := h.helperAccountKeeper.GetAccount(ctx, claimerDstAddress).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	campaign, _ := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	moduleName := cfeclaimtypes.ModuleName
	if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
		moduleName = cfevestingtypes.ModuleName
	}
	require.True(h.t, ok)
	moduleBalanceBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, moduleName)
	claimerBalanceBefore := h.BankUtils.GetAccountAllBalances(ctx, claimerDstAddress)
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

	var expectedCoins sdk.Coins
	for _, coin := range userEntry.GetClaimRecord(campaignId).Amount {
		expectedAmount := mission.Weight.MulInt(coin.Amount).TruncateInt()
		claimerCoinBefore := claimerBalanceBefore.AmountOf(coin.Denom)
		h.BankUtils.VerifyAccountBalanceByDenom(ctx, claimer, coin.Denom, claimerCoinBefore.Add(expectedAmount))
		moduleCoinBefore := moduleBalanceBefore.AmountOf(coin.Denom)
		h.BankUtils.VerifyModuleAccountBalanceByDenom(ctx, moduleName, coin.Denom, moduleCoinBefore.Sub(expectedAmount))
		expectedCoins = expectedCoins.Add(sdk.NewCoin(coin.Denom, expectedAmount))
	}

	h.addExpectedDataToAccount(ctx, campaignId, claimerAccountBefore, expectedCoins)

	claimerAccount, ok := h.helperAccountKeeper.GetAccount(ctx, claimerDstAddress).(*cfevestingtypes.PeriodicContinuousVestingAccount)
	require.True(h.t, ok)
	require.NoError(h.t, claimerAccount.Validate())
	require.EqualValues(h.t, claimerAccountBefore, claimerAccount)
}

func (h *C4eClaimUtils) addExpectedDataToAccount(ctx sdk.Context, campaignId uint64,
	claimerAccountBefore *cfevestingtypes.PeriodicContinuousVestingAccount, expectedAmount sdk.Coins) *cfevestingtypes.PeriodicContinuousVestingAccount {
	campaign, _ := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	expectedStartTime := ctx.BlockTime().Add(campaign.LockupPeriod)
	expectedEndTime := expectedStartTime.Add(campaign.VestingPeriod)
	expectedOriginalVesting := sdk.NewCoins(expectedAmount...)
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
	moduleBefore := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	claimerBefore := h.BankUtils.GetAccountAllBalances(ctx, claimer)
	claimRecordBefore, foundCrBefore := h.helpeCfeclaimkeeper.GetUserEntry(ctx, claimer.String())

	require.EqualError(h.t, h.helpeCfeclaimkeeper.Claim(ctx, campaignId, missionId, claimer.String()), errorMessage)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))
	h.BankUtils.VerifyAccountAllBalances(ctx, claimer, claimerBefore)
	h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, moduleBefore)
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

func (h *C4eClaimUtils) CreateCampaign(ctx sdk.Context, owner string, name string, description string, campaignType cfeclaimtypes.CampaignType, feegrantAmount math.Int, initialClaimFreeAmount math.Int, free sdk.Dec, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration, vestingPoolName string) {

	campaignCountBefore := h.helpeCfeclaimkeeper.GetCampaignCount(ctx)
	_, err := h.helpeCfeclaimkeeper.CreateCampaign(ctx, owner, name, description, campaignType, &feegrantAmount, &initialClaimFreeAmount, &free, &startTime, &endTime, &lockupPeriod, &vestingPeriod, vestingPoolName)
	missionCountAfter := h.helpeCfeclaimkeeper.GetMissionCount(ctx, campaignCountBefore)
	require.NoError(h.t, err)
	campaignCountAfter := h.helpeCfeclaimkeeper.GetCampaignCount(ctx)
	require.Equal(h.t, campaignCountBefore+1, campaignCountAfter)
	require.Equal(h.t, uint64(1), missionCountAfter)

	h.VerifyCampaign(ctx, campaignCountBefore, true, owner, name, description, false, &feegrantAmount, &initialClaimFreeAmount, startTime, endTime, lockupPeriod, vestingPeriod, vestingPoolName)
	h.VerifyMission(ctx, true, campaignCountBefore, 0, "Initial mission", "Initial mission - basic mission that must be claimed first", cfeclaimtypes.MissionInitialClaim, sdk.ZeroDec(), nil)
}

func (h *C4eClaimUtils) CreateCampaignError(ctx sdk.Context, owner string, name string, description string, campaignType cfeclaimtypes.CampaignType, feegrantAmount math.Int, initialClaimFreeAmount math.Int, free sdk.Dec, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration, vestingPoolName string, errorMessage string) {

	campaignCountBefore := h.helpeCfeclaimkeeper.GetCampaignCount(ctx)
	_, err := h.helpeCfeclaimkeeper.CreateCampaign(ctx, owner, name, description, campaignType, &feegrantAmount, &initialClaimFreeAmount, &free, &startTime, &endTime, &lockupPeriod, &vestingPeriod, vestingPoolName)
	require.EqualError(h.t, err, errorMessage)
	campaignCountAfter := h.helpeCfeclaimkeeper.GetCampaignCount(ctx)
	missionCountAfter := h.helpeCfeclaimkeeper.GetMissionCount(ctx, campaignCountBefore)
	require.Equal(h.t, campaignCountBefore, campaignCountAfter)
	require.Equal(h.t, uint64(0), missionCountAfter)
	_, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignCountBefore)
	require.False(h.t, ok)
}

func (h *C4eClaimUtils) EnableCampaign(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) {
	err := h.helpeCfeclaimkeeper.EnableCampaign(ctx, owner, campaignId, startTime, endTime)
	require.NoError(h.t, err)
	campaign, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	require.True(h.t, ok)
	h.VerifyCampaign(ctx, campaign.Id, true, owner, campaign.Name, campaign.Description, true, &campaign.FeegrantAmount, &campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod, campaign.VestingPoolName)
}

func (h *C4eClaimUtils) EnableCampaignError(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time, errorString string) {
	campaignBefore, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)

	err := h.helpeCfeclaimkeeper.EnableCampaign(ctx, owner, campaignId, startTime, endTime)
	require.EqualError(h.t, err, errorString)
	if !ok {
		return
	}
	enabled := campaignBefore.Enabled
	h.VerifyCampaign(ctx, campaignBefore.Id, true, campaignBefore.Owner, campaignBefore.Name, campaignBefore.Description, enabled, &campaignBefore.FeegrantAmount, &campaignBefore.InitialClaimFreeAmount, campaignBefore.StartTime, campaignBefore.EndTime, campaignBefore.LockupPeriod, campaignBefore.VestingPeriod, campaignBefore.VestingPoolName)
}

func (h *C4eClaimUtils) CloseCampaign(ctx sdk.Context, owner string, campaignId uint64) {
	campaignAmoutLeftBefore, _ := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	cfeclaimModuleBalance := h.BankUtils.GetModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName)
	campaign, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	require.True(h.t, ok)
	_, feegrantAccountAddress := cfeclaimmodulekeeper.CreateFeegrantAccountAddress(campaign.Id)
	feegrantAmountLefBefore := sdk.NewCoins()
	if campaign.FeegrantAmount.IsPositive() {
		feegrantAmountLefBefore = h.BankUtils.GetAccountAllBalances(ctx, feegrantAccountAddress)
	}
	ownerAccAddress, _ := sdk.AccAddressFromBech32(campaign.Owner)
	ownerBalanceBefore := h.BankUtils.GetAccountAllBalances(ctx, ownerAccAddress)
	var vestingPoolBefore *cfevestingtypes.VestingPool

	if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
		_, vestingPoolBefore, _ = h.helperCfevestingKeeper.GetAccountVestingPool(ctx, owner, campaign.VestingPoolName)
	}

	err := h.helpeCfeclaimkeeper.CloseCampaign(ctx, owner, campaignId)
	require.NoError(h.t, err)

	campaign, _ = h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)
	campaignAmoutLeft, _ := h.helpeCfeclaimkeeper.GetCampaignAmountLeft(ctx, campaignId)
	require.True(h.t, campaignAmoutLeft.Amount.IsEqual(sdk.NewCoins()))

	if campaign.FeegrantAmount.IsPositive() {
		feegrantAmountLef := h.BankUtils.GetAccountAllBalances(ctx, feegrantAccountAddress)
		require.True(h.t, feegrantAmountLef.IsZero())
	}
	amountDiff := campaignAmoutLeftBefore.Amount.Sub(feegrantAmountLefBefore...)

	if campaign.CampaignType == cfeclaimtypes.VestingPoolCampaign {
		_, vestingPool, _ := h.helperCfevestingKeeper.GetAccountVestingPool(ctx, owner, campaign.VestingPoolName)
		if campaign.FeegrantAmount.GT(math.ZeroInt()) {
			require.True(h.t, vestingPoolBefore.Sent.Sub(amountDiff.AmountOf(h.helperCfevestingKeeper.GetParams(ctx).Denom)).Equal(vestingPool.Sent))
		} else {
			require.Nil(h.t, vestingPool.GetReservation(campaignId))
		}

	} else {
		h.BankUtils.VerifyModuleAccountAllBalances(ctx, cfeclaimtypes.ModuleName, cfeclaimModuleBalance.Sub(campaignAmoutLeftBefore.Amount...))
		h.BankUtils.VerifyAccountAllBalances(ctx, ownerAccAddress, ownerBalanceBefore.Add(campaignAmoutLeftBefore.Amount.Add(feegrantAmountLefBefore...)...))
	}

	h.VerifyCloseAction(ctx, campaignId, campaignAmoutLeftBefore.Amount)
	h.VerifyCampaign(ctx, campaign.Id, true, owner, campaign.Name, campaign.Description, false, &campaign.FeegrantAmount, &campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod, campaign.VestingPoolName)
}

func (h *C4eClaimUtils) CloseCampaignError(ctx sdk.Context, owner string, campaignId uint64, errorString string) {
	campaignBefore, ok := h.helpeCfeclaimkeeper.GetCampaign(ctx, campaignId)

	err := h.helpeCfeclaimkeeper.CloseCampaign(ctx, owner, campaignId)
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

func (h *C4eClaimUtils) VerifyCampaign(ctx sdk.Context, campaignId uint64, mustExist bool, owner string, name string, description string, enabled bool, feegrantAmount *math.Int, initialClaimFreeAmount *math.Int, startTime time.Time,
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
		require.EqualValues(h.t, claimCampaign.FeegrantAmount, math.ZeroInt())
	} else {
		require.True(h.t, claimCampaign.FeegrantAmount.Equal(*feegrantAmount))
	}

	if initialClaimFreeAmount.IsNil() {
		require.EqualValues(h.t, claimCampaign.InitialClaimFreeAmount, math.ZeroInt())
	} else {
		require.True(h.t, claimCampaign.InitialClaimFreeAmount.Equal(*initialClaimFreeAmount))
	}

	require.EqualValues(h.t, claimCampaign.Enabled, enabled)
	require.True(h.t, claimCampaign.StartTime.Equal(startTime))
	require.True(h.t, claimCampaign.EndTime.Equal(endTime))
	require.EqualValues(h.t, claimCampaign.VestingPeriod, vestingPeriod)
	require.EqualValues(h.t, claimCampaign.LockupPeriod, lockupPeriod)
}

func (h *C4eClaimUtils) VerifyCloseAction(ctx sdk.Context, campaignId uint64, campaignAmountLeftBefore sdk.Coins) {

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
