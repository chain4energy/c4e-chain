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
	bankUtils *commontestutils.BankUtils) *ContextC4eAirdropUtils {
	c4eAirdropUtils := NewC4eAirdropUtils(t, helpeCfeairdropmodulekeeper, helperAccountKeeper, bankUtils)
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
