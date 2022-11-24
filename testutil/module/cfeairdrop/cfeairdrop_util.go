package cfeairdrop

import (
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfeairdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/stretchr/testify/require"
)

type C4eAirdropUtils struct {
	C4eAirdropKeeperUtils
	helperAccountKeeper *authkeeper.AccountKeeper
	BankUtils           *commontestutils.BankUtils
}

func NewC4eAirdropUtils(t *testing.T, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *commontestutils.BankUtils) C4eAirdropUtils {
	return C4eAirdropUtils{C4eAirdropKeeperUtils: NewC4eAirdropKeeperUtils(t, helpeCfeairdropmodulekeeper), helperAccountKeeper: helperAccountKeeper, BankUtils: bankUtils}
}

func (h *C4eAirdropUtils) SendToAirdropAccount(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool) {
	coins := sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, amount))
	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	accBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)

	previousOriginalVesting := sdk.NewCoins()
	previousPeriods := []cfeairdroptypes.ContinuousVestingPeriod{}
	if accountBefore != nil {
		if airdropAccount, ok := accountBefore.(*cfeairdroptypes.AirdropVestingAccount); ok {
			previousOriginalVesting = previousOriginalVesting.Add(airdropAccount.OriginalVesting...)
			previousPeriods = airdropAccount.VestingPeriods
		}
	}

	require.NoError(h.t, h.helpeCfeairdropkeeper.SendToAirdropAccount(ctx,
		toAddress,
		coins,
		startTime,
		endTime, createAccount,
	))

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance.Add(amount))
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance.Sub(amount))

	airdropAccount, ok := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfeairdroptypes.AirdropVestingAccount)
	require.True(h.t, ok)
	// periodPosition := len(airdropAccount.VestingPeriods) - 1
	// require.EqualValues(h.t, periodsAmount+1, len(airdropAccount.VestingPeriods))
	// require.EqualValues(h.t, startTime, airdropAccount.StartTime)
	// require.EqualValues(h.t, endTime, airdropAccount.EndTime)
	// require.EqualValues(h.t, previousOriginalVesting.Add(coins...), airdropAccount.OriginalVesting)
	// require.EqualValues(h.t, startTime, airdropAccount.VestingPeriods[periodPosition].StartTime)
	// require.EqualValues(h.t, endTime, airdropAccount.VestingPeriods[periodPosition].EndTime)
	// require.EqualValues(h.t, coins, airdropAccount.VestingPeriods[periodPosition].Amount)

	newPeriods := append(previousPeriods, cfeairdroptypes.ContinuousVestingPeriod{StartTime: startTime, EndTime: endTime, Amount: coins})
	h.VerifyAirdropAccount(ctx, toAddress, previousOriginalVesting.Add(coins...), startTime, endTime, newPeriods)
	require.NoError(h.t, airdropAccount.Validate())

}

func (h *C4eAirdropUtils) SendToAirdropAccountError(ctx sdk.Context, toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool, errorMessage string, expectNewAccount bool) {
	coins := sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, amount))
	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)
	accBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, toAddress)

	accountBefore := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	wasAirdropAccount := false
	if accountBefore != nil {
		_, wasAirdropAccount = accountBefore.(*cfeairdroptypes.AirdropVestingAccount)
	}

	require.EqualError(h.t, h.helpeCfeairdropkeeper.SendToAirdropAccount(ctx,
		toAddress,
		coins,
		startTime,
		endTime, createAccount,
	), errorMessage)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, toAddress, accBalance)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance)

	accountAfter := h.helperAccountKeeper.GetAccount(ctx, toAddress)
	_, isAirdropAccount := h.helperAccountKeeper.GetAccount(ctx, toAddress).(*cfeairdroptypes.AirdropVestingAccount)

	if accountBefore == nil && expectNewAccount {
		require.EqualValues(h.t, true, isAirdropAccount)
		h.VerifyAirdropAccount(ctx, toAddress, sdk.NewCoins(), startTime, endTime, []cfeairdroptypes.ContinuousVestingPeriod{})

	} else {
		require.EqualValues(h.t, wasAirdropAccount, isAirdropAccount)
		require.EqualValues(h.t, accountBefore, accountAfter)
	}

}

func (h *C4eAirdropUtils) VerifyAirdropAccount(ctx sdk.Context, address sdk.AccAddress,
	expectedOriginalVesting sdk.Coins, expectedStartTime int64, expectedEndTime int64, expectedPeriods []cfeairdroptypes.ContinuousVestingPeriod) {

	airdropAccount, ok := h.helperAccountKeeper.GetAccount(ctx, address).(*cfeairdroptypes.AirdropVestingAccount)
	require.True(h.t, ok)

	require.EqualValues(h.t, len(expectedPeriods), len(airdropAccount.VestingPeriods))
	require.EqualValues(h.t, expectedStartTime, airdropAccount.StartTime)
	require.EqualValues(h.t, expectedEndTime, airdropAccount.EndTime)
	require.True(h.t, expectedOriginalVesting.IsEqual(airdropAccount.OriginalVesting))
	for i := 0; i < len(expectedPeriods); i++ {
		require.EqualValues(h.t, expectedPeriods[i], airdropAccount.VestingPeriods[i])
	}
	require.NoError(h.t, airdropAccount.Validate())

}

func (h *C4eAirdropUtils) AddCampaignRecords(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, campaignRecord map[string]sdk.Int) {
	sum := sdk.ZeroInt()
	for _, amount := range campaignRecord {
		sum = sum.Add(amount)
	}
	allRecordsBefore := h.helpeCfeairdropkeeper.GetAllClaimRecord(ctx)
	h.BankUtils.AddDefaultDenomCoinsToAccount(ctx, sum, srcAddress)
	srcBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	require.NoError(h.t, h.helpeCfeairdropkeeper.AddCampaignRecords(ctx, srcAddress, campaignId, campaignRecord))

	allRecords := h.helpeCfeairdropkeeper.GetAllClaimRecord(ctx)
	require.EqualValues(h.t, len(campaignRecord), len(allRecords))
	for address, claimable := range campaignRecord {
		claimRecord, found := h.helpeCfeairdropkeeper.GetClaimRecord(ctx, address)
		require.True(h.t, found)
		var recordBefore *cfeairdroptypes.ClaimRecord = nil
		for _, before := range allRecordsBefore {
			if before.Address == address {
				recordBefore = &before
				break
			}
		}
		if recordBefore == nil {
			require.EqualValues(h.t, 1, len(claimRecord.CampaignRecords))
		} else {
			require.EqualValues(h.t, len(recordBefore.CampaignRecords)+1, len(claimRecord.CampaignRecords))
		}
		require.EqualValues(h.t, address, claimRecord.Address)
		require.EqualValues(h.t, "", claimRecord.ClaimAddress)
		if recordBefore == nil {
			require.EqualValues(h.t, campaignId, claimRecord.CampaignRecords[0].CampaignId)
			require.True(h.t, claimable.Equal(claimRecord.CampaignRecords[0].Claimable))
			require.EqualValues(h.t, 0, len(claimRecord.CampaignRecords[0].CompletedMissions))
		} else {
			expectedCaipaignRecords := append(recordBefore.CampaignRecords, &cfeairdroptypes.CampaignRecord{CampaignId: campaignId, Claimable: claimable, CompletedMissions: nil})

			require.ElementsMatch(h.t, expectedCaipaignRecords, claimRecord.CampaignRecords)
		}

	}
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance.Add(sum))
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance.Sub(sum))

}

func (h *C4eAirdropUtils) AddCampaignRecordsError(ctx sdk.Context, srcAddress sdk.AccAddress, campaignId uint64, campaignRecord map[string]sdk.Int, errorMessage string, addRequiredCoinsToSrc bool) {
	if addRequiredCoinsToSrc {
		sum := sdk.ZeroInt()
		for _, amount := range campaignRecord {
			sum = sum.Add(amount)
		}
		h.BankUtils.AddDefaultDenomCoinsToAccount(ctx, sum, srcAddress)
	}
	allRecordsBefore := h.helpeCfeairdropkeeper.GetAllClaimRecord(ctx)

	srcBalance := h.BankUtils.GetAccountDefultDenomBalance(ctx, srcAddress)

	moduleBalance := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeairdropkeeper.AddCampaignRecords(ctx, srcAddress, campaignId, campaignRecord), errorMessage)

	allRecords := h.helpeCfeairdropkeeper.GetAllClaimRecord(ctx)
	require.ElementsMatch(h.t, allRecordsBefore, allRecords)
	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBalance)
	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, srcAddress, srcBalance)

}

func (h *C4eAirdropUtils) ClaimInitial(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress) {

	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	require.NoError(h.t, h.helpeCfeairdropkeeper.ClaimInitial(ctx, campaignId, claimer.String()))
	initialClaim, foundIc := h.helpeCfeairdropkeeper.GetInitialClaim(ctx, campaignId)
	require.True(h.t, foundIc)

	mission, _ := h.helpeCfeairdropkeeper.GetMission(ctx, campaignId, initialClaim.MissionId)
	claimRecord, _ := h.helpeCfeairdropkeeper.GetClaimRecord(ctx, claimer.String())

	expectedAmount := mission.Weight.MulInt(claimRecord.GetCampaignRecord(campaignId).Claimable).TruncateInt()

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, expectedAmount)

	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore.Sub(expectedAmount))

	campaign := h.helpeCfeairdropkeeper.Campaign(ctx, campaignId)

	expectedOriginalVesting := sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, expectedAmount))
	expectedStartTime := ctx.BlockTime().Add(campaign.LockupPeriod)
	expectedEndTime := expectedStartTime.Add(campaign.VestingPeriod)
	expectedPeriod := cfeairdroptypes.ContinuousVestingPeriod{StartTime: expectedStartTime.Unix(), EndTime: expectedEndTime.Unix(), Amount: expectedOriginalVesting}
	h.VerifyAirdropAccount(ctx, claimer, expectedOriginalVesting,
		expectedStartTime.Unix(), expectedEndTime.Unix(), []cfeairdroptypes.ContinuousVestingPeriod{expectedPeriod})

	claimRecord, found := h.helpeCfeairdropkeeper.GetClaimRecord(ctx, claimer.String())
	require.True(h.t, found)
	camapaignRecord := claimRecord.GetCampaignRecord(campaignId)
	require.NotNil(h.t, camapaignRecord)
	require.ElementsMatch(h.t, []uint64{initialClaim.MissionId}, camapaignRecord.CompletedMissions)
	require.ElementsMatch(h.t, []uint64{initialClaim.MissionId}, camapaignRecord.ClaimedMissions)
}

func (h *C4eAirdropUtils) ClaimInitialError(ctx sdk.Context, campaignId uint64, claimer sdk.AccAddress, errorMessage string) {
	claimerAccountBefore := h.helperAccountKeeper.GetAccount(ctx, claimer)
	claimRecordBefore, foundBefore := h.helpeCfeairdropkeeper.GetClaimRecord(ctx, claimer.String())

	moduleBefore := h.BankUtils.GetModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName)

	require.EqualError(h.t, h.helpeCfeairdropkeeper.ClaimInitial(ctx, campaignId, claimer.String()), errorMessage)

	h.BankUtils.VerifyAccountDefultDenomBalance(ctx, claimer, sdk.ZeroInt())

	h.BankUtils.VerifyModuleAccountDefultDenomBalance(ctx, cfeairdroptypes.ModuleName, moduleBefore)

	require.EqualValues(h.t, claimerAccountBefore, h.helperAccountKeeper.GetAccount(ctx, claimer))

	claimRecord, found := h.helpeCfeairdropkeeper.GetClaimRecord(ctx, claimer.String())
	require.Equal(h.t, foundBefore, found)
	if found {
		require.EqualValues(h.t, claimRecordBefore, claimRecord)

	}
}

func (h *C4eAirdropUtils) GetClaimRecord(
	ctx sdk.Context,
	address string,
) *cfeairdroptypes.ClaimRecord {
	val, found := h.helpeCfeairdropkeeper.GetClaimRecord(ctx, address)
	if found {
		return &val
	}
	return nil
}

func (h *C4eAirdropUtils) SetClaimRecord(
	ctx sdk.Context,
	claimRecord *cfeairdroptypes.ClaimRecord,
) {
	h.helpeCfeairdropkeeper.SetClaimRecord(ctx, *claimRecord)
}
