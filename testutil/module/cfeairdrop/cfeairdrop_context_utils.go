package cfeairdrop

import (
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfeairdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/stretchr/testify/require"
)

type ContextC4eAirdropUtils struct {
	C4eAirdropUtils
	testContext testenv.TestContext
}

func NewContextC4eAirdropUtils(t require.TestingT, testContext testenv.TestContext, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *testcosmos.BankUtils, stakingUtils *testcosmos.StakingUtils, govUtils *testcosmos.GovUtils, feegrantUtils *testcosmos.FeegrantUtils, distributionUtils *testcosmos.DistributionUtils) *ContextC4eAirdropUtils {
	c4eAirdropUtils := NewC4eAirdropUtils(t, helpeCfeairdropmodulekeeper, helperAccountKeeper, bankUtils, stakingUtils, govUtils, feegrantUtils, distributionUtils)
	return &ContextC4eAirdropUtils{C4eAirdropUtils: c4eAirdropUtils, testContext: testContext}
}

func (h *ContextC4eAirdropUtils) CreateCampaign(owner string, campaign cfeairdroptypes.Campaign) {
	h.C4eAirdropUtils.CreateCampaign(h.testContext.GetContext(), owner, campaign.Name, campaign.Description, campaign.CampaignType, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod)
}

func (h *ContextC4eAirdropUtils) CreateCampaignError(owner string, campaign cfeairdroptypes.Campaign, errorMessage string) {
	h.C4eAirdropUtils.CreateCampaignError(h.testContext.GetContext(), owner, campaign.Name, campaign.Description, campaign.CampaignType, campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod, campaign.VestingPeriod, errorMessage)
}

func (h *ContextC4eAirdropUtils) AddMissionToCampaign(owner string, campaignId uint64, mission cfeairdroptypes.Mission) {
	h.C4eAirdropUtils.AddMissionToCampaign(h.testContext.GetContext(), owner, campaignId, mission.Name, mission.Description, mission.MissionType, mission.Weight, mission.ClaimStartDate)
}

func (h *ContextC4eAirdropUtils) AddMissionToCampaignError(owner string, campaignId uint64, mission cfeairdroptypes.Mission, errorString string) {
	h.C4eAirdropUtils.AddMissionToCampaignError(h.testContext.GetContext(), owner, campaignId, mission.Name, mission.Description, mission.MissionType, mission.Weight, mission.ClaimStartDate, errorString)
}

func (h *ContextC4eAirdropUtils) StartCampaign(owner string, campaignId uint64) {
	h.C4eAirdropUtils.StartCampaign(h.testContext.GetContext(), owner, campaignId)
}

func (h *ContextC4eAirdropUtils) StartCampaignError(owner string, campaignId uint64, errorString string) {
	h.C4eAirdropUtils.StartCampaignError(h.testContext.GetContext(), owner, campaignId, errorString)
}

func (h *ContextC4eAirdropUtils) CloseCampaign(owner string, campaignId uint64, campaignCloseAction cfeairdroptypes.CampaignCloseAction) {
	h.C4eAirdropUtils.CloseCampaign(h.testContext.GetContext(), owner, campaignId, campaignCloseAction)
}

func (h *ContextC4eAirdropUtils) CloseCampaignError(owner string, campaignId uint64, campaignCloseAction cfeairdroptypes.CampaignCloseAction, errorString string) {
	h.C4eAirdropUtils.CloseCampaignError(h.testContext.GetContext(), owner, campaignId, campaignCloseAction, errorString)
}

func (h *ContextC4eAirdropUtils) SendToRepeatedContinuousVestingAccount(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, missionType cfeairdroptypes.MissionType) {
	h.C4eAirdropUtils.SendToRepeatedContinuousVestingAccount(h.testContext.GetContext(), toAddress, amount, startTime, endTime, missionType)
}

func (h *ContextC4eAirdropUtils) SendToRepeatedContinuousVestingAccountError(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, missionType cfeairdroptypes.MissionType) {
	h.C4eAirdropUtils.SendToRepeatedContinuousVestingAccountError(h.testContext.GetContext(), toAddress, amount, startTime, endTime, createAccount, errorMessage, missionType)
}

func (h *ContextC4eAirdropUtils) VerifyRepeatedContinuousVestingAccount(address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfevestingtypes.ContinuousVestingPeriod, missionType cfeairdroptypes.MissionType) {
	h.C4eAirdropUtils.VerifyRepeatedContinuousVestingAccount(h.testContext.GetContext(), address, expectedOriginalVesting, expectedStartTime, expectedEndTime, expectedPeriods, missionType)
}

func (h *ContextC4eAirdropUtils) InitGenesis(genState cfeairdroptypes.GenesisState) {
	h.C4eAirdropKeeperUtils.InitGenesis(h.testContext.GetContext(), genState)
}

func (h *ContextC4eAirdropUtils) AddClaimRecords(srcAddress sdk.AccAddress, campaignId uint64, airdropEntries []*cfeairdroptypes.ClaimRecord) {
	h.C4eAirdropUtils.AddClaimRecords(h.testContext.GetContext(), srcAddress, campaignId, airdropEntries)
}

func (h *ContextC4eAirdropUtils) AddClaimRecordsFromWhitelistedVestingAccount(srcAddress sdk.AccAddress, amountToSend sdk.Int, unlockedAmount sdk.Int) {
	coinsToSend := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amountToSend))
	unlockedCoins := sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, unlockedAmount))
	h.C4eAirdropUtils.AddClaimRecordsFromWhitelistedVestingAccount(h.testContext.GetContext(), srcAddress, coinsToSend, unlockedCoins)
}

func (h *ContextC4eAirdropUtils) AddCoinsToCampaignOwnerAcc(srcAddress sdk.AccAddress, amountOfCoins sdk.Int) {
	h.BankUtils.AddDefaultDenomCoinsToAccount(h.testContext.GetContext(), amountOfCoins, srcAddress)
}

func (h *ContextC4eAirdropUtils) AddClaimRecordsError(srcAddress sdk.AccAddress, campaignId uint64, airdropEntries []*cfeairdroptypes.ClaimRecord, errorMessage string) {
	h.C4eAirdropUtils.AddClaimRecordsError(h.testContext.GetContext(), srcAddress, campaignId, airdropEntries, errorMessage)
}

func (h *ContextC4eAirdropUtils) ClaimInitial(claimer sdk.AccAddress, campaignId uint64, expectedAmount int64) {
	h.C4eAirdropUtils.ClaimInitial(h.testContext.GetContext(), campaignId, claimer, expectedAmount)
}

func (h *ContextC4eAirdropUtils) ClaimInitialError(claimer sdk.AccAddress, campaignId uint64, errorMessage string) {
	h.C4eAirdropUtils.ClaimInitialError(h.testContext.GetContext(), campaignId, claimer, errorMessage)
}

func (h *ContextC4eAirdropUtils) GetUsersEntries(address string) *cfeairdroptypes.UserEntry {
	return h.C4eAirdropUtils.GetUsersEntries(h.testContext.GetContext(), address)
}

func (h *ContextC4eAirdropUtils) SetUsersEntries(userEntry *cfeairdroptypes.UserEntry) {
	h.C4eAirdropUtils.SetUsersEntries(h.testContext.GetContext(), userEntry)
}

func (h *ContextC4eAirdropUtils) CompleteMissionFromHook(campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.C4eAirdropUtils.CompleteMissionFromHook(h.testContext.GetContext(), campaignId, missionId, claimer)

}

func (h *ContextC4eAirdropUtils) CompleteMissionFromHookError(campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	h.C4eAirdropUtils.CompleteMissionFromHookError(h.testContext.GetContext(), campaignId, missionId, claimer, errorMessage)

}

func (h *ContextC4eAirdropUtils) ClaimMission(campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.C4eAirdropUtils.ClaimMission(h.testContext.GetContext(), campaignId, missionId, claimer)
}

func (h *ContextC4eAirdropUtils) ClaimMissionToAddress(campaignId uint64, missionId uint64, claimer sdk.AccAddress, claimerDstAddress sdk.AccAddress) {
	h.C4eAirdropUtils.ClaimMissionToAddress(h.testContext.GetContext(), campaignId, missionId, claimer, claimerDstAddress)
}

func (h *ContextC4eAirdropUtils) ClaimMissionError(campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	h.C4eAirdropUtils.ClaimMissionError(h.testContext.GetContext(), campaignId, missionId, claimer, errorMessage)
}

func (h *ContextC4eAirdropUtils) CreateRepeatedContinuousVestingAccount(address sdk.AccAddress, originalVesting sdk.Coins, startTime int64,
	endTime int64, periods ...cfevestingtypes.ContinuousVestingPeriod) *cfevestingtypes.RepeatedContinuousVestingAccount {
	return h.C4eAirdropUtils.CreateRepeatedContinuousVestingAccount(h.testContext.GetContext(), address, originalVesting, startTime, endTime, periods...)
}

func (h *ContextC4eAirdropUtils) CompleteDelegationMission(campaignId uint64, missionId uint64,
	claimer sdk.AccAddress, deleagtionAmount sdk.Int) {
	h.C4eAirdropUtils.CompleteDelegationMission(h.testContext.GetContext(), campaignId, missionId, claimer, deleagtionAmount)
}

func (h *ContextC4eAirdropUtils) CompleteVoteMission(campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.C4eAirdropUtils.CompleteVoteMission(h.testContext.GetContext(), campaignId, missionId, claimer)

}
