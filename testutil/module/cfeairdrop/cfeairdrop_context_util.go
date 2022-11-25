package cfeairdrop

import (
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfeairdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

type ContextC4eAirdropUtils struct {
	C4eAirdropUtils
	testContext commontestutils.TestContext
}

func NewContextC4eAirdropUtils(t *testing.T, testContext commontestutils.TestContext, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *commontestutils.BankUtils, stakingUtils *commontestutils.StakingUtils, govUtils *commontestutils.GovUtils) *ContextC4eAirdropUtils {
	c4eAirdropUtils := NewC4eAirdropUtils(t, helpeCfeairdropmodulekeeper, helperAccountKeeper, bankUtils, stakingUtils, govUtils)
	return &ContextC4eAirdropUtils{C4eAirdropUtils: c4eAirdropUtils, testContext: testContext}
}

func (h *ContextC4eAirdropUtils) SendToAirdropAccount(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool) {
	h.C4eAirdropUtils.SendToAirdropAccount(h.testContext.GetContext(), toAddress, amount, startTime, endTime, createAccount)
}

func (h *ContextC4eAirdropUtils) SendToAirdropAccountError(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, expectNewAccount bool) {
	h.C4eAirdropUtils.SendToAirdropAccountError(h.testContext.GetContext(), toAddress, amount, startTime, endTime, createAccount, errorMessage, expectNewAccount)

}

func (h *ContextC4eAirdropUtils) VerifyAirdropAccount(address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfeairdroptypes.ContinuousVestingPeriod) {
	h.C4eAirdropUtils.VerifyAirdropAccount(h.testContext.GetContext(), address, expectedOriginalVesting, expectedStartTime, expectedEndTime, expectedPeriods)
}

func (h *ContextC4eAirdropUtils) InitGenesis(genState cfeairdroptypes.GenesisState) {
	h.C4eAirdropKeeperUtils.InitGenesis(h.testContext.GetContext(), genState)
}

func (h *ContextC4eAirdropUtils) AddCampaignRecords(srcAddress sdk.AccAddress, campaignId uint64, campaignRecords map[string]sdk.Int) {
	h.C4eAirdropUtils.AddCampaignRecords(h.testContext.GetContext(), srcAddress, campaignId, campaignRecords)

}

func (h *ContextC4eAirdropUtils) ClaimInitial(campaignId uint64, claimer sdk.AccAddress) {
	h.C4eAirdropUtils.ClaimInitial(h.testContext.GetContext(), campaignId, claimer)
}

func (h *ContextC4eAirdropUtils) ClaimInitialError(campaignId uint64, claimer sdk.AccAddress, errorMessage string) {
	h.C4eAirdropUtils.ClaimInitialError(h.testContext.GetContext(), campaignId, claimer, errorMessage)
}

func (h *ContextC4eAirdropUtils) GetClaimRecord(address string) *cfeairdroptypes.ClaimRecord {
	return h.C4eAirdropUtils.GetClaimRecord(h.testContext.GetContext(), address)
}

func (h *ContextC4eAirdropUtils) SetClaimRecord(claimRecord *cfeairdroptypes.ClaimRecord) {
	h.C4eAirdropUtils.SetClaimRecord(h.testContext.GetContext(), claimRecord)
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
	endTime int64, periods ...cfeairdroptypes.ContinuousVestingPeriod) *cfeairdroptypes.AirdropVestingAccount {
	return h.C4eAirdropUtils.CreateAirdropAccout(h.testContext.GetContext(), address, originalVesting, startTime, endTime, periods...)
}

func (h *ContextC4eAirdropUtils) CompleteDelegationMission(campaignId uint64,
	claimer sdk.AccAddress, deleagtionAmount sdk.Int) {
	h.C4eAirdropUtils.CompleteDelegationMission(h.testContext.GetContext(), campaignId, claimer, deleagtionAmount)
}

func (h *ContextC4eAirdropUtils) CompleteVoteMission(campaignId uint64, claimer sdk.AccAddress) {
	h.C4eAirdropUtils.CompleteVoteMission(h.testContext.GetContext(), campaignId, claimer)

}
