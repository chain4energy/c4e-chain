package cfeairdrop

import (
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfeairdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"testing"
	"time"
)

type ContextC4eAirdropUtils struct {
	C4eAirdropUtils
	testContext testenv.TestContext
}

func NewContextC4eAirdropUtils(t *testing.T, testContext testenv.TestContext, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *testcosmos.BankUtils, stakingUtils *testcosmos.StakingUtils, govUtils *testcosmos.GovUtils) *ContextC4eAirdropUtils {
	c4eAirdropUtils := NewC4eAirdropUtils(t, helpeCfeairdropmodulekeeper, helperAccountKeeper, bankUtils, stakingUtils, govUtils)
	return &ContextC4eAirdropUtils{C4eAirdropUtils: c4eAirdropUtils, testContext: testContext}
}

func (h *ContextC4eAirdropUtils) CreateAirdropCampaign(owner string, name string, description string, feegrantAmount sdk.Int, initialClaimFreeAmount sdk.Int, startTime time.Time,
	endTime time.Time, lockupPeriod time.Duration, vestingPeriod time.Duration) {
	h.C4eAirdropUtils.CreateAirdropCampaign(h.testContext.GetContext(), owner, name, description, feegrantAmount, initialClaimFreeAmount, startTime, endTime, lockupPeriod, vestingPeriod)
}

func (h *ContextC4eAirdropUtils) AddMissionToAirdropCampaign(owner string, campaignId uint64, name string, description string, missionType cfeairdroptypes.MissionType,
	weight sdk.Dec) {
	h.C4eAirdropUtils.AddMissionToAirdropCampaign(h.testContext.GetContext(), owner, campaignId, name, description, missionType, weight)
}

func (h *ContextC4eAirdropUtils) StartAirdropCampaign(owner string, campaignId uint64) {
	h.C4eAirdropUtils.StartAirdropCampaign(h.testContext.GetContext(), owner, campaignId)
}

func (h *ContextC4eAirdropUtils) SendToAirdropAccount(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, missionType cfeairdroptypes.MissionType) {
	h.C4eAirdropUtils.SendToAirdropAccount(h.testContext.GetContext(), toAddress, amount, startTime, endTime, missionType)
}

func (h *ContextC4eAirdropUtils) SendToAirdropAccountError(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, missionType cfeairdroptypes.MissionType) {
	h.C4eAirdropUtils.SendToAirdropAccountError(h.testContext.GetContext(), toAddress, amount, startTime, endTime, createAccount, errorMessage, missionType)
}

func (h *ContextC4eAirdropUtils) VerifyAirdropAccount(address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfevestingtypes.ContinuousVestingPeriod, missionType cfeairdroptypes.MissionType) {
	h.C4eAirdropUtils.VerifyAirdropAccount(h.testContext.GetContext(), address, expectedOriginalVesting, expectedStartTime, expectedEndTime, expectedPeriods, missionType)
}

func (h *ContextC4eAirdropUtils) InitGenesis(genState cfeairdroptypes.GenesisState) {
	h.C4eAirdropKeeperUtils.InitGenesis(h.testContext.GetContext(), genState)
}

func (h *ContextC4eAirdropUtils) AddAirdropEntries(srcAddress sdk.AccAddress, campaignId uint64, airdropEntries []*cfeairdroptypes.AirdropEntry) {
	h.C4eAirdropUtils.AddAirdropEntries(h.testContext.GetContext(), srcAddress, campaignId, airdropEntries)
}

func (h *ContextC4eAirdropUtils) ClaimInitial(campaignId uint64, claimer sdk.AccAddress, expectedAmount int64) {
	h.C4eAirdropUtils.ClaimInitial(h.testContext.GetContext(), campaignId, claimer, expectedAmount)
}

func (h *ContextC4eAirdropUtils) ClaimInitialError(campaignId uint64, claimer sdk.AccAddress, errorMessage string) {
	h.C4eAirdropUtils.ClaimInitialError(h.testContext.GetContext(), campaignId, claimer, errorMessage)
}

func (h *ContextC4eAirdropUtils) GetUserAirdropEntries(address string) *cfeairdroptypes.UserAirdropEntries {
	return h.C4eAirdropUtils.GetUserAirdropEntries(h.testContext.GetContext(), address)
}

func (h *ContextC4eAirdropUtils) SetUserAirdropEntries(userAirdropEntries *cfeairdroptypes.UserAirdropEntries) {
	h.C4eAirdropUtils.SetUserAirdropEntries(h.testContext.GetContext(), userAirdropEntries)
}

func (h *ContextC4eAirdropUtils) CompleteMission(campaignId uint64, missionId uint64, claimer sdk.AccAddress) {
	h.C4eAirdropUtils.CompleteMission(h.testContext.GetContext(), campaignId, missionId, claimer)

}

func (h *ContextC4eAirdropUtils) CompleteMissionError(campaignId uint64, missionId uint64, claimer sdk.AccAddress, errorMessage string) {
	h.C4eAirdropUtils.CompleteMissionError(h.testContext.GetContext(), campaignId, missionId, claimer, errorMessage)

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

func (h *ContextC4eAirdropUtils) CreateAirdropAccout(address sdk.AccAddress, originalVesting sdk.Coins, startTime int64,
	endTime int64, periods ...cfevestingtypes.ContinuousVestingPeriod) *cfevestingtypes.RepeatedContinuousVestingAccount {
	return h.C4eAirdropUtils.CreateAirdropAccout(h.testContext.GetContext(), address, originalVesting, startTime, endTime, periods...)
}

func (h *ContextC4eAirdropUtils) CompleteDelegationMission(campaignId uint64,
	claimer sdk.AccAddress, deleagtionAmount sdk.Int) {
	h.C4eAirdropUtils.CompleteDelegationMission(h.testContext.GetContext(), campaignId, claimer, deleagtionAmount)
}

func (h *ContextC4eAirdropUtils) CompleteVoteMission(campaignId uint64, claimer sdk.AccAddress) {
	h.C4eAirdropUtils.CompleteVoteMission(h.testContext.GetContext(), campaignId, claimer)

}
