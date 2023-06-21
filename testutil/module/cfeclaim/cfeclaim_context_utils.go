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
	h.C4eClaimUtils.CreateCampaign(h.testContext.GetContext(), owner, campaign.Name, campaign.Description, campaign.CampaignType, campaign.RemovableClaimRecords,
		campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.Free, campaign.StartTime, campaign.EndTime,
		campaign.LockupPeriod, campaign.VestingPeriod, campaign.VestingPoolName)
}

func (h *ContextC4eClaimUtils) CreateCampaignError(owner string, campaign cfeclaimtypes.Campaign, errorMessage string) {
	h.C4eClaimUtils.CreateCampaignError(h.testContext.GetContext(), owner, campaign.Name, campaign.Description, campaign.CampaignType, campaign.RemovableClaimRecords,
		campaign.FeegrantAmount, campaign.InitialClaimFreeAmount, campaign.Free, campaign.StartTime, campaign.EndTime, campaign.LockupPeriod,
		campaign.VestingPeriod, campaign.VestingPoolName, errorMessage)
}

func (h *ContextC4eClaimUtils) AddMission(owner string, campaignId uint64, mission cfeclaimtypes.Mission) {
	h.C4eClaimUtils.AddMission(h.testContext.GetContext(), owner, campaignId, mission.Name, mission.Description, mission.MissionType, mission.Weight, mission.ClaimStartDate)
}

func (h *ContextC4eClaimUtils) AddMissionError(owner string, campaignId uint64, mission cfeclaimtypes.Mission, errorString string) {
	h.C4eClaimUtils.AddMissionError(h.testContext.GetContext(), owner, campaignId, mission.Name, mission.Description, mission.MissionType, mission.Weight, mission.ClaimStartDate, errorString)
}

func (h *ContextC4eClaimUtils) EnableCampaign(owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) {
	h.C4eClaimUtils.EnableCampaign(h.testContext.GetContext(), owner, campaignId, startTime, endTime)
}

func (h *ContextC4eClaimUtils) EnableCampaignError(owner string, campaignId uint64, startTime *time.Time, endTime *time.Time, errorString string) {
	h.C4eClaimUtils.EnableCampaignError(h.testContext.GetContext(), owner, campaignId, startTime, endTime, errorString)
}

func (h *ContextC4eClaimUtils) CloseCampaign(owner string, campaignId uint64) {
	h.C4eClaimUtils.CloseCampaign(h.testContext.GetContext(), owner, campaignId)
}

func (h *ContextC4eClaimUtils) CloseCampaignError(owner string, campaignId uint64, errorString string) {
	h.C4eClaimUtils.CloseCampaignError(h.testContext.GetContext(), owner, campaignId, errorString)
}

func (h *ContextC4eClaimUtils) RemoveCampaign(owner string, campaignId uint64) {
	h.C4eClaimUtils.RemoveCampaign(h.testContext.GetContext(), owner, campaignId)
}

func (h *ContextC4eClaimUtils) RemoveCampaignError(owner string, campaignId uint64, errorString string) {
	h.C4eClaimUtils.RemoveCampaignError(h.testContext.GetContext(), owner, campaignId, errorString)
}

func (h *ContextC4eClaimUtils) InitGenesis(genState cfeclaimtypes.GenesisState) {
	h.C4eClaimKeeperUtils.InitGenesis(h.testContext.GetContext(), genState)
}

func (h *ContextC4eClaimUtils) InitGenesisError(genState cfeclaimtypes.GenesisState, errorString string) {
	h.C4eClaimKeeperUtils.InitGenesisError(h.testContext.GetContext(), genState, errorString)
}

func (h *ContextC4eClaimUtils) ExportGenesis(genState cfeclaimtypes.GenesisState) {
	h.C4eClaimKeeperUtils.ExportGenesis(h.testContext.GetContext(), genState)
}

func (h *ContextC4eClaimUtils) AddClaimRecords(srcAddress sdk.AccAddress, campaignId uint64, claimRecordEntries []*cfeclaimtypes.ClaimRecordEntry) {
	h.C4eClaimUtils.AddClaimRecords(h.testContext.GetContext(), srcAddress, campaignId, claimRecordEntries)
}

func (h *ContextC4eClaimUtils) DeleteClaimRecord(ownerAddress sdk.AccAddress, campaignId uint64, userAddress string, amoutDiff sdk.Coins) {
	h.C4eClaimUtils.DeleteClaimRecord(h.testContext.GetContext(), ownerAddress, campaignId, userAddress, amoutDiff)
}

func (h *ContextC4eClaimUtils) DeleteClaimRecordError(ownerAddress sdk.AccAddress, campaignId uint64, userAddress string, errorMessage string) {
	h.C4eClaimUtils.DeleteClaimRecordError(h.testContext.GetContext(), ownerAddress, campaignId, userAddress, errorMessage)
}

func (h *ContextC4eClaimUtils) AddCoinsToCampaignOwnerAcc(srcAddress sdk.AccAddress, coins sdk.Coins) {
	h.BankUtils.AddCoinsToAccount(h.testContext.GetContext(), coins, srcAddress)
}

func (h *ContextC4eClaimUtils) AddClaimRecordsError(srcAddress sdk.AccAddress, campaignId uint64, claimRecordEntries []*cfeclaimtypes.ClaimRecordEntry, errorMessage string) {
	h.C4eClaimUtils.AddClaimRecordsError(h.testContext.GetContext(), srcAddress, campaignId, claimRecordEntries, errorMessage)
}

func (h *ContextC4eClaimUtils) ClaimInitial(claimer sdk.AccAddress, campaignId uint64) {
	h.C4eClaimUtils.ClaimInitial(h.testContext.GetContext(), campaignId, claimer)
}

func (m *ContextC4eClaimUtils) ValidateGenesisAndInvariants() {
	m.C4eClaimUtils.ExportGenesisAndValidate(m.testContext.GetContext())
	m.C4eClaimUtils.ValidateInvariants(m.testContext.GetContext())
}

func (h *ContextC4eClaimUtils) ClaimInitialError(claimer sdk.AccAddress, campaignId uint64, errorMessage string) {
	h.C4eClaimUtils.ClaimInitialError(h.testContext.GetContext(), campaignId, claimer, errorMessage)
}

func (h *ContextC4eClaimUtils) GetUsersEntry(address string) *cfeclaimtypes.UserEntry {
	return h.C4eClaimUtils.GetUsersEntry(h.testContext.GetContext(), address)
}

func (h *ContextC4eClaimUtils) GetAllUsersEntries() []cfeclaimtypes.UserEntry {
	return h.C4eClaimUtils.GetAllUsersEntries(h.testContext.GetContext())
}

func (h *ContextC4eClaimUtils) GetCampaigns() []cfeclaimtypes.Campaign {
	return h.C4eClaimUtils.GetCampaigns(h.testContext.GetContext())
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

func (h *ContextC4eClaimUtils) CreatePeriodicContinuousVestingAccount(address sdk.AccAddress, originalVesting sdk.Coins, startTime int64,
	endTime int64, periods ...cfevestingtypes.ContinuousVestingPeriod) *cfevestingtypes.PeriodicContinuousVestingAccount {
	return h.C4eClaimUtils.CreatePeriodicContinuousVestingAccount(h.testContext.GetContext(), address, originalVesting, startTime, endTime, periods...)
}

func (h *ContextC4eClaimUtils) CompleteDelegationMission(campaignId uint64, missionId uint64,
	claimer sdk.AccAddress, deleagtionAmount math.Int, valAddress sdk.ValAddress) {
	h.C4eClaimUtils.CompleteDelegationMission(h.testContext.GetContext(), campaignId, missionId, claimer, deleagtionAmount, valAddress)
}

func (h *ContextC4eClaimUtils) CompleteVoteMission(campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.C4eClaimUtils.CompleteVoteMission(h.testContext.GetContext(), campaignId, missionId, claimer)

}

func (h *ContextC4eClaimUtils) CheckNonNegativeCoinStateInvariant(ctx sdk.Context, failed bool, message string) {
	invariant := cfeclaimmodulekeeper.CampaignCurrentAmountSumCheckInvariant(*h.helpeCfeclaimkeeper)
	testcosmos.CheckInvariant(h.t, ctx, invariant, failed, message)
}
