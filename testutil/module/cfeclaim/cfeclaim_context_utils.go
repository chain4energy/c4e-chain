package cfeclaim

import (
	"cosmossdk.io/math"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfeclaimmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	cfeclaimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingmodulekeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/stretchr/testify/require"
	"time"
)

type ContextC4eClaimUtils struct {
	C4eClaimUtils
	testContext testenv.TestContext
}

func NewContextC4eClaimUtils(t require.TestingT, testContext testenv.TestContext, helpeCfeclaimmodulekeeper *cfeclaimmodulekeeper.Keeper,
	helperCfevestingKeeper *cfevestingmodulekeeper.Keeper, helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *testcosmos.BankUtils, stakingUtils *testcosmos.StakingUtils, govUtils *testcosmos.GovUtils, feegrantUtils *testcosmos.FeegrantUtils, distributionUtils *testcosmos.DistributionUtils) *ContextC4eClaimUtils {
	c4eClaimUtils := NewC4eClaimUtils(t, helpeCfeclaimmodulekeeper, helperCfevestingKeeper, helperAccountKeeper, bankUtils, stakingUtils, govUtils, feegrantUtils, distributionUtils)
	return &ContextC4eClaimUtils{C4eClaimUtils: c4eClaimUtils, testContext: testContext}
}

func (h *ContextC4eClaimUtils) CreateCampaign(owner string, campaign cfeclaimtypes.Campaign) {
	h.C4eClaimUtils.CreateCampaign(h.testContext.GetContext(), owner, campaign.Name, campaign.Description, campaign.CampaignType, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod, campaign.VestingPoolName)
}

func (h *ContextC4eClaimUtils) CreateCampaignError(owner string, campaign cfeclaimtypes.Campaign, errorMessage string) {
	h.C4eClaimUtils.CreateCampaignError(h.testContext.GetContext(), owner, campaign.Name, campaign.Description, campaign.CampaignType, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod, campaign.VestingPoolName, errorMessage)
}

func (h *ContextC4eClaimUtils) AddMissionToCampaign(owner string, campaignId uint64, mission cfeclaimtypes.Mission) {
	h.C4eClaimUtils.AddMissionToCampaign(h.testContext.GetContext(), owner, campaignId, mission.Name, mission.Description, mission.MissionType, mission.Weight, mission.ClaimStartDate)
}

func (h *ContextC4eClaimUtils) AddMissionToCampaignError(owner string, campaignId uint64, mission cfeclaimtypes.Mission, errorString string) {
	h.C4eClaimUtils.AddMissionToCampaignError(h.testContext.GetContext(), owner, campaignId, mission.Name, mission.Description, mission.MissionType, mission.Weight, mission.ClaimStartDate, errorString)
}

func (h *ContextC4eClaimUtils) StartCampaign(owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) {
	h.C4eClaimUtils.StartCampaign(h.testContext.GetContext(), owner, campaignId, startTime, endTime)
}

func (h *ContextC4eClaimUtils) StartCampaignError(owner string, campaignId uint64, startTime *time.Time, endTime *time.Time, errorString string) {
	h.C4eClaimUtils.StartCampaignError(h.testContext.GetContext(), owner, campaignId, startTime, endTime, errorString)
}

func (h *ContextC4eClaimUtils) CloseCampaign(owner string, campaignId uint64, CloseAction cfeclaimtypes.CloseAction) {
	h.C4eClaimUtils.CloseCampaign(h.testContext.GetContext(), owner, campaignId, CloseAction)
}

func (h *ContextC4eClaimUtils) CloseCampaignError(owner string, campaignId uint64, CloseAction cfeclaimtypes.CloseAction, errorString string) {
	h.C4eClaimUtils.CloseCampaignError(h.testContext.GetContext(), owner, campaignId, CloseAction, errorString)
}

func (h *ContextC4eClaimUtils) SendToRepeatedContinuousVestingAccount(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, missionType cfeclaimtypes.MissionType) {
	h.C4eClaimUtils.SendToRepeatedContinuousVestingAccount(h.testContext.GetContext(), toAddress, amount, startTime, endTime, missionType)
}

func (h *ContextC4eClaimUtils) SendToRepeatedContinuousVestingAccountError(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, missionType cfeclaimtypes.MissionType) {
	h.C4eClaimUtils.SendToRepeatedContinuousVestingAccountError(h.testContext.GetContext(), toAddress, amount, startTime, endTime, createAccount, errorMessage, missionType)
}

func (h *ContextC4eClaimUtils) VerifyRepeatedContinuousVestingAccount(address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfevestingtypes.ContinuousVestingPeriod, missionType cfeclaimtypes.MissionType) {
	h.C4eClaimUtils.VerifyRepeatedContinuousVestingAccount(h.testContext.GetContext(), address, expectedOriginalVesting, expectedStartTime, expectedEndTime, expectedPeriods, missionType)
}

func (h *ContextC4eClaimUtils) InitGenesis(genState cfeclaimtypes.GenesisState) {
	h.C4eClaimKeeperUtils.InitGenesis(h.testContext.GetContext(), genState)
}

func (h *ContextC4eClaimUtils) AddClaimRecords(srcAddress sdk.AccAddress, campaignId uint64, claimEntries []*cfeclaimtypes.ClaimRecord) {
	h.C4eClaimUtils.AddClaimRecords(h.testContext.GetContext(), srcAddress, campaignId, claimEntries)
}

func (h *ContextC4eClaimUtils) DeleteClaimRecord(ownerAddress sdk.AccAddress, campaignId uint64, userAddress string, deleteClaimRecordAction cfeclaimtypes.CloseAction, amoutDiff sdk.Coins) {
	h.C4eClaimUtils.DeleteClaimRecord(h.testContext.GetContext(), ownerAddress, campaignId, userAddress, deleteClaimRecordAction, amoutDiff)
}

func (h *ContextC4eClaimUtils) DeleteClaimRecordError(ownerAddress sdk.AccAddress, campaignId uint64, userAddress string, deleteClaimRecordAction cfeclaimtypes.CloseAction, errorMessage string) {
	h.C4eClaimUtils.DeleteClaimRecordError(h.testContext.GetContext(), ownerAddress, campaignId, userAddress, deleteClaimRecordAction, errorMessage)
}

func (h *ContextC4eClaimUtils) AddClaimRecordsFromWhitelistedVestingAccount(srcAddress sdk.AccAddress, amountToSend sdk.Int, unlockedAmount sdk.Int) {
	coinsToSend := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amountToSend))
	unlockedCoins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, unlockedAmount))
	h.C4eClaimUtils.AddClaimRecordsFromWhitelistedVestingAccount(h.testContext.GetContext(), srcAddress, coinsToSend, unlockedCoins)
}

func (h *ContextC4eClaimUtils) AddCoinsToCampaignOwnerAcc(srcAddress sdk.AccAddress, amountOfCoins sdk.Int) {
	h.BankUtils.AddDefaultDenomCoinsToAccount(h.testContext.GetContext(), amountOfCoins, srcAddress)
}

func (h *ContextC4eClaimUtils) AddClaimRecordsError(srcAddress sdk.AccAddress, campaignId uint64, claimEntries []*cfeclaimtypes.ClaimRecord, errorMessage string) {
	h.C4eClaimUtils.AddClaimRecordsError(h.testContext.GetContext(), srcAddress, campaignId, claimEntries, errorMessage)
}

func (h *ContextC4eClaimUtils) ClaimInitial(claimer sdk.AccAddress, campaignId uint64, expectedAmount int64) {
	h.C4eClaimUtils.ClaimInitial(h.testContext.GetContext(), campaignId, claimer, expectedAmount)
}

func (h *ContextC4eClaimUtils) ClaimInitialError(claimer sdk.AccAddress, campaignId uint64, errorMessage string) {
	h.C4eClaimUtils.ClaimInitialError(h.testContext.GetContext(), campaignId, claimer, errorMessage)
}

func (h *ContextC4eClaimUtils) GetUsersEntries(address string) *cfeclaimtypes.UserEntry {
	return h.C4eClaimUtils.GetUsersEntries(h.testContext.GetContext(), address)
}

func (h *ContextC4eClaimUtils) SetUsersEntries(userEntry *cfeclaimtypes.UserEntry) {
	h.C4eClaimUtils.SetUsersEntries(h.testContext.GetContext(), userEntry)
}

func (h *ContextC4eClaimUtils) CompleteMissionFromHook(campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.C4eClaimUtils.CompleteMissionFromHook(h.testContext.GetContext(), campaignId, missionId, claimer)

}

func (h *ContextC4eClaimUtils) CompleteMissionFromHookError(campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	h.C4eClaimUtils.CompleteMissionFromHookError(h.testContext.GetContext(), campaignId, missionId, claimer, errorMessage)

}

func (h *ContextC4eClaimUtils) ClaimMission(campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.C4eClaimUtils.ClaimMission(h.testContext.GetContext(), campaignId, missionId, claimer)
}

func (h *ContextC4eClaimUtils) ClaimMissionToAddress(campaignId uint64, missionId uint64, claimer sdk.AccAddress, claimerDstAddress sdk.AccAddress) {
	h.C4eClaimUtils.ClaimMissionToAddress(h.testContext.GetContext(), campaignId, missionId, claimer, claimerDstAddress)
}

func (h *ContextC4eClaimUtils) ClaimMissionError(campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	h.C4eClaimUtils.ClaimMissionError(h.testContext.GetContext(), campaignId, missionId, claimer, errorMessage)
}

func (h *ContextC4eClaimUtils) CreateRepeatedContinuousVestingAccount(address sdk.AccAddress, originalVesting sdk.Coins, startTime int64,
	endTime int64, periods ...cfevestingtypes.ContinuousVestingPeriod) *cfevestingtypes.PeriodicContinuousVestingAccount {
	return h.C4eClaimUtils.CreateRepeatedContinuousVestingAccount(h.testContext.GetContext(), address, originalVesting, startTime, endTime, periods...)
}

func (h *ContextC4eClaimUtils) CompleteDelegationMission(campaignId uint64, missionId uint64,
	claimer sdk.AccAddress, deleagtionAmount math.Int, valAddress sdk.ValAddress) {
	h.C4eClaimUtils.CompleteDelegationMission(h.testContext.GetContext(), campaignId, missionId, claimer, deleagtionAmount, valAddress)
}

func (h *ContextC4eClaimUtils) CompleteVoteMission(campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.C4eClaimUtils.CompleteVoteMission(h.testContext.GetContext(), campaignId, missionId, claimer)

}

func (h *ContextC4eClaimUtils) CheckNonNegativeCoinStateInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfeclaimmodulekeeper.CampaignAmountLeftSumCheckInvariant(*h.helpeCfeclaimkeeper)
	testcosmos.CheckInvariant(h.t, ctx, invariant, failed, message)
}
