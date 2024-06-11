package e2e

import (
	"cosmossdk.io/math"
	"fmt"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"github.com/chain4energy/c4e-chain/tests/e2e/util"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	"github.com/chain4energy/c4e-chain/testutil/utils"
	cfeclaimcli "github.com/chain4energy/c4e-chain/x/cfeclaim/client/cli"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfevestingmoduletypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
	"time"
)

type ClaimSetupSuite struct {
	BaseSetupSuite
}

func TestClaimSuite(t *testing.T) {
	suite.Run(t, new(ClaimSetupSuite))
}

func (s *ClaimSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(false, false, false)
}

func (s *ClaimSetupSuite) TestDefaultCampaign() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := utils.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)

	node.BankSendBaseBalanceFromNode(creatorAddress)

	s.NoError(err)

	free := sdk.ZeroDec()
	startTime := time.Now().Add(time.Minute)
	endTime := startTime.Add(time.Minute * 2)
	lockupPeriod := time.Hour
	vestingPeriod := time.Hour
	randName := utils.RandStringOfLength(10)
	randDescription := utils.RandStringOfLength(10)
	feegrantAmount := math.NewInt(2500000).String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.DefaultCampaign
	vestingPoolName := ""

	campaignId := node.CreateCampaign(randName, randDescription, campaignType.String(), "false", feegrantAmount, inititalClaimFreeAmount, free.String(),
		startTime.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupPeriod.String(), vestingPeriod.String(), vestingPoolName, creatorWalletName)
	campaignIdString := strconv.FormatUint(campaignId, 10)

	claimStartDate := startTime.Add(time.Minute)
	randMissionName := utils.RandStringOfLength(10)
	randMissionDescription := utils.RandStringOfLength(10)
	node.AddMission(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName)
	node.AddMissionError(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.8",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName, "all campaign missions weight sum is >= 1 (1.100000000000000000 > 1) error: wrong param valu")

	balances, err := node.QueryBalances(creatorAddress)
	require.NoError(s.T(), err)
	claimRecordEntries, claimRecordEntriesWalletNames := prepareNClaimRecords(node, 10, balances.AmountOf(appparams.MicroC4eUnit))
	destinationAddress := testcosmos.CreateRandomAccAddress()

	claimer := claimRecordEntriesWalletNames[2]
	node.ValidateClaimInitialClaimerNotFound(campaignIdString, destinationAddress, claimer)
	userEntriesJSONString, err := util.NewClaimRecordsListJson(claimRecordEntries)
	s.NoError(err)
	node.AddClaimRecordsError(campaignIdString, userEntriesJSONString, creatorWalletName, "owner balance is too small")
	node.BankSendBaseBalanceFromNode(creatorAddress)
	node.AddClaimRecords(campaignIdString, userEntriesJSONString, creatorWalletName)
	node.AddClaimRecordsError(campaignIdString, userEntriesJSONString, creatorWalletName, fmt.Sprintf("campaignId %s already exists for address", campaignIdString))

	node.DeleteClaimRecord(campaignIdString, claimRecordEntries[0].UserEntryAddress, creatorWalletName)
	node.ValidateDeleteClaimerNotFound(campaignIdString, claimRecordEntries[0].UserEntryAddress, creatorWalletName)

	node.ClaimInitialMissionError(campaignIdString, destinationAddress, claimer, "campaign is disabled")

	node.EnableCampaign(campaignIdString, "", "", creatorWalletName)
	node.RemoveCampaignError(campaignIdString, creatorWalletName, "campaign is enabled")
	node.DeleteClaimRecordError(campaignIdString, claimRecordEntries[2].UserEntryAddress, creatorWalletName,
		"campaign must have RemovableClaimRecords flag set to true to be able to delete its entries")

	node.ValidateClaimInitialCampaignNotStartedYet(campaignIdString, destinationAddress, claimer)
	node.WaitUntilSpecifiedTime(startTime)

	node.ClaimInitialMission(campaignIdString, destinationAddress, claimer)
	node.ClaimInitialMissionError(campaignIdString, destinationAddress, claimer, "missionId: 0: mission already completed")

	node.ClaimMissionError(campaignIdString, "1", claimer, "mission 1 not started yet")
	node.WaitUntilSpecifiedTime(claimStartDate)
	node.ClaimMission(campaignIdString, "1", claimer)
	node.ClaimMissionError(campaignIdString, "1", claimer, "missionId: 1: mission already completed")
	node.ValidateCampaignIsNotOverYet(campaignIdString, creatorWalletName)
	node.WaitUntilSpecifiedTime(endTime)
	node.CloseCampaign(campaignIdString, creatorWalletName)
	node.RemoveCampaign(campaignIdString, creatorWalletName)
}

func (s *ClaimSetupSuite) TestDefaultCampaignNoFeegrantAndUpdatedEnableTimes() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := utils.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)

	node.BankSendBaseBalanceFromNode(creatorAddress)

	s.NoError(err)

	free := sdk.ZeroDec()
	campaignStartTimeBefore := time.Now().Add(time.Second * 35)
	endTime := campaignStartTimeBefore.Add(time.Second * 90)
	lockupPeriod := time.Hour
	vestingPeriod := time.Hour
	randName := utils.RandStringOfLength(10)
	randDescription := utils.RandStringOfLength(10)
	feegrantAmount := math.ZeroInt().String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.DefaultCampaign
	vestingPoolName := ""

	campaignId := node.CreateCampaign(randName, randDescription, campaignType.String(), "false", feegrantAmount, inititalClaimFreeAmount, free.String(),
		campaignStartTimeBefore.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupPeriod.String(), vestingPeriod.String(), vestingPoolName, creatorWalletName)
	campaignIdString := strconv.FormatUint(campaignId, 10)

	balances, err := node.QueryBalances(creatorAddress)
	require.NoError(s.T(), err)
	claimRecordEntries, claimRecordEntriesWalletNames := prepareNClaimRecords(node, 10, balances.AmountOf(appparams.MicroC4eUnit))
	destinationAddress := testcosmos.CreateRandomAccAddress()
	randomUserEntryIndex := utils.RandInt(10)
	claimer := claimRecordEntriesWalletNames[randomUserEntryIndex]
	node.ValidateClaimInitialClaimerNotFound(campaignIdString, destinationAddress, claimer)
	userEntriesJSONString, err := util.NewClaimRecordsListJson(claimRecordEntries)
	s.NoError(err)
	node.AddClaimRecords(campaignIdString, userEntriesJSONString, creatorWalletName)
	node.ClaimInitialMissionError(campaignIdString, destinationAddress, claimer, "campaign is disabled")
	node.BankSendBaseBalanceFromNode(claimRecordEntries[randomUserEntryIndex].UserEntryAddress)

	updatedStartTime := time.Now().Add(time.Minute)
	updatedEndTime := updatedStartTime.Add(time.Minute * 2)
	node.EnableCampaign(campaignIdString, updatedStartTime.Format(cfeclaimcli.TimeLayout), updatedEndTime.Format(cfeclaimcli.TimeLayout), creatorWalletName)
	node.WaitUntilSpecifiedTime(campaignStartTimeBefore)
	node.ValidateClaimInitialCampaignNotStartedYet(campaignIdString, destinationAddress, claimer)
	node.WaitUntilSpecifiedTime(updatedStartTime)
	node.ClaimInitialMission(campaignIdString, destinationAddress, claimer)
}

func (s *ClaimSetupSuite) TestCampaignRemovableClaimRecords() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := utils.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)

	node.BankSendBaseBalanceFromNode(creatorAddress)

	s.NoError(err)

	free := sdk.ZeroDec()
	startTime := time.Now().Add(time.Second * 35)
	endTime := startTime.Add(time.Second * 90)
	lockupPeriod := time.Hour
	vestingPeriod := time.Hour
	randName := utils.RandStringOfLength(10)
	randDescription := utils.RandStringOfLength(10)
	feegrantAmount := math.NewInt(2500000).String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.DefaultCampaign
	vestingPoolName := ""

	campaignId := node.CreateCampaign(randName, randDescription, campaignType.String(), "true", feegrantAmount, inititalClaimFreeAmount, free.String(),
		startTime.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupPeriod.String(), vestingPeriod.String(), vestingPoolName, creatorWalletName)
	campaignIdString := strconv.FormatUint(campaignId, 10)

	claimStartDate := startTime.Add(time.Second * 65)
	randMissionName := utils.RandStringOfLength(10)
	randMissionDescription := utils.RandStringOfLength(10)
	node.AddMission(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName)
	node.AddMission(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName)

	balances, err := node.QueryBalances(creatorAddress)
	require.NoError(s.T(), err)
	claimRecordEntries, claimRecordEntriesWalletNames := prepareNClaimRecords(node, 10, balances.AmountOf(appparams.MicroC4eUnit))
	destinationAddress := testcosmos.CreateRandomAccAddress()

	claimer := claimRecordEntriesWalletNames[2]
	node.ValidateClaimInitialClaimerNotFound(campaignIdString, destinationAddress, claimer)
	userEntriesJSONString, err := util.NewClaimRecordsListJson(claimRecordEntries)
	s.NoError(err)
	node.BankSendBaseBalanceFromNode(creatorAddress)
	node.AddClaimRecords(campaignIdString, userEntriesJSONString, creatorWalletName)
	node.DeleteClaimRecord(campaignIdString, claimRecordEntries[0].UserEntryAddress, creatorWalletName)
	node.EnableCampaign(campaignIdString, "", "", creatorWalletName)
	node.DeleteClaimRecord(campaignIdString, claimRecordEntries[1].UserEntryAddress, creatorWalletName)
	node.ValidateClaimInitialCampaignNotStartedYet(campaignIdString, destinationAddress, claimer)
	node.WaitUntilSpecifiedTime(startTime)

	node.ClaimInitialMission(campaignIdString, destinationAddress, claimer)
	node.WaitUntilSpecifiedTime(claimStartDate)
	node.ClaimMission(campaignIdString, "1", claimer)
	node.DeleteClaimRecord(campaignIdString, claimRecordEntries[2].UserEntryAddress, creatorWalletName)
	node.ClaimMissionError(campaignIdString, "2", claimer, fmt.Sprintf("claim record with campaign id %s not found for address %s", campaignIdString, claimRecordEntries[2].UserEntryAddress))
	node.WaitUntilSpecifiedTime(endTime)
	node.CloseCampaign(campaignIdString, creatorWalletName)
}

func (s *ClaimSetupSuite) TestRemoveCampaign() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := utils.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)

	node.BankSendBaseBalanceFromNode(creatorAddress)

	s.NoError(err)

	free := sdk.ZeroDec()
	startTime := time.Now().Add(time.Second * 35)
	endTime := startTime.Add(time.Second * 90)
	lockupPeriod := time.Hour
	vestingPeriod := time.Hour
	randName := utils.RandStringOfLength(10)
	randDescription := utils.RandStringOfLength(10)
	feegrantAmount := math.NewInt(2500000).String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.DefaultCampaign
	vestingPoolName := ""

	campaignId := node.CreateCampaign(randName, randDescription, campaignType.String(), "true", feegrantAmount, inititalClaimFreeAmount, free.String(),
		startTime.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupPeriod.String(), vestingPeriod.String(), vestingPoolName, creatorWalletName)
	campaignIdString := strconv.FormatUint(campaignId, 10)

	claimStartDate := startTime.Add(time.Second * 65)
	randMissionName := utils.RandStringOfLength(10)
	randMissionDescription := utils.RandStringOfLength(10)
	node.AddMission(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName)
	node.AddMission(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName)

	balances, err := node.QueryBalances(creatorAddress)
	require.NoError(s.T(), err)
	claimRecordEntries, _ := prepareNClaimRecords(node, 10, balances.AmountOf(appparams.MicroC4eUnit))
	userEntriesJSONString, err := util.NewClaimRecordsListJson(claimRecordEntries)
	s.NoError(err)
	node.BankSendBaseBalanceFromNode(creatorAddress)
	node.AddClaimRecords(campaignIdString, userEntriesJSONString, creatorWalletName)

	node.DeleteClaimRecord(campaignIdString, claimRecordEntries[0].UserEntryAddress, creatorWalletName)
	node.RemoveCampaign(campaignIdString, creatorWalletName)
	node.RemoveCampaignError(campaignIdString, creatorWalletName, "")
}

func (s *ClaimSetupSuite) TestVestingPoolCampaign() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := utils.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)

	node.BankSendBaseBalanceFromNode(creatorAddress)

	s.NoError(err)
	balanceBefore, err := node.QueryBalances(creatorAddress)
	randVestingPoolName := utils.RandStringOfLength(5)
	vestingPoolDuration := time.Hour
	vestingTypes := node.QueryVestingTypes()
	s.Greater(len(vestingTypes), 0)
	vestingType := vestingTypes[0]
	node.CreateVestingPool(randVestingPoolName, balanceBefore.AmountOf(appparams.MicroC4eUnit).String(), vestingPoolDuration.String(), vestingType.Name, creatorWalletName)
	free := sdk.ZeroDec()
	startTime := time.Now().Add(time.Minute)
	endTime := startTime.Add(time.Minute * 2)
	lockupPeriod, err := cfevestingmoduletypes.DurationFromUnits(cfevestingmoduletypes.PeriodUnit(vestingType.LockupPeriodUnit), vestingType.LockupPeriod)
	s.NoError(err)
	vestingPeriod, err := cfevestingmoduletypes.DurationFromUnits(cfevestingmoduletypes.PeriodUnit(vestingType.VestingPeriodUnit), vestingType.VestingPeriod)
	s.NoError(err)
	randName := utils.RandStringOfLength(10)
	randDescription := utils.RandStringOfLength(10)
	feegrantAmount := math.NewInt(2500000).String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.VestingPoolCampaign

	campaignId := node.CreateCampaign(randName, randDescription, campaignType.String(), "false", feegrantAmount, inititalClaimFreeAmount, free.String(),
		startTime.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupPeriod.String(), vestingPeriod.String(), randVestingPoolName, creatorWalletName)
	campaignIdString := strconv.FormatUint(campaignId, 10)

	claimStartDate := startTime.Add(time.Minute)
	randMissionName := utils.RandStringOfLength(10)
	randMissionDescription := utils.RandStringOfLength(10)
	node.AddMission(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName)
	node.AddMissionError(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.8",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName, "all campaign missions weight sum is >= 1 (1.100000000000000000 > 1) error: wrong param valu")

	claimRecordEntries, claimRecordEntriesWalletNames := prepareNClaimRecords(node, 10, balanceBefore.AmountOf(appparams.MicroC4eUnit))
	destinationAddress := testcosmos.CreateRandomAccAddress()

	claimer := claimRecordEntriesWalletNames[2]
	node.ValidateClaimInitialClaimerNotFound(campaignIdString, destinationAddress, claimer)
	userEntriesJSONString, err := util.NewClaimRecordsListJson(claimRecordEntries)
	s.NoError(err)
	node.AddClaimRecordsError(campaignIdString, userEntriesJSONString, creatorWalletName, "owner balance is too small")
	node.BankSendBaseBalanceFromNode(creatorAddress)
	node.AddClaimRecords(campaignIdString, userEntriesJSONString, creatorWalletName)
	node.AddClaimRecordsError(campaignIdString, userEntriesJSONString, creatorWalletName, fmt.Sprintf("campaignId %s already exists for address", campaignIdString))

	node.DeleteClaimRecord(campaignIdString, claimRecordEntries[0].UserEntryAddress, creatorWalletName)
	node.ValidateDeleteClaimerNotFound(campaignIdString, claimRecordEntries[0].UserEntryAddress, creatorWalletName)

	node.ClaimInitialMissionError(campaignIdString, destinationAddress, claimer, "campaign is disabled")

	node.EnableCampaign(campaignIdString, "", "", creatorWalletName)
	node.RemoveCampaignError(campaignIdString, creatorWalletName, "campaign is enabled")
	node.DeleteClaimRecordError(campaignIdString, claimRecordEntries[2].UserEntryAddress, creatorWalletName,
		"campaign must have RemovableClaimRecords flag set to true to be able to delete its entries")

	node.ValidateClaimInitialCampaignNotStartedYet(campaignIdString, destinationAddress, claimer)
	node.WaitUntilSpecifiedTime(startTime)

	node.ClaimInitialMission(campaignIdString, destinationAddress, claimer)
	node.ClaimInitialMissionError(campaignIdString, destinationAddress, claimer, "missionId: 0: mission already completed")

	node.ClaimMissionError(campaignIdString, "1", claimer, "mission 1 not started yet")
	node.WaitUntilSpecifiedTime(claimStartDate)
	node.ClaimMission(campaignIdString, "1", claimer)
	node.ClaimMissionError(campaignIdString, "1", claimer, "missionId: 1: mission already completed")
	node.ValidateCampaignIsNotOverYet(campaignIdString, creatorWalletName)
	node.WaitUntilSpecifiedTime(endTime)
	node.CloseCampaign(campaignIdString, creatorWalletName)
	node.RemoveCampaign(campaignIdString, creatorWalletName)
}

func (s *ClaimSetupSuite) TestMissionToDefine() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	creatorWalletName := utils.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)

	node.BankSendBaseBalanceFromNode(creatorAddress)

	s.NoError(err)

	free := sdk.ZeroDec()
	startTime := time.Now().Add(time.Second * 35)
	endTime := startTime.Add(time.Second * 90)
	lockupPeriod := time.Hour
	vestingPeriod := time.Hour
	randName := utils.RandStringOfLength(10)
	randDescription := utils.RandStringOfLength(10)
	feegrantAmount := math.NewInt(2500000).String()
	inititalClaimFreeAmount := math.ZeroInt().String()
	campaignType := types.DefaultCampaign
	vestingPoolName := ""

	campaignId := node.CreateCampaign(randName, randDescription, campaignType.String(), "true", feegrantAmount, inititalClaimFreeAmount, free.String(),
		startTime.Format(cfeclaimcli.TimeLayout), endTime.Format(cfeclaimcli.TimeLayout), lockupPeriod.String(), vestingPeriod.String(), vestingPoolName, creatorWalletName)
	campaignIdString := strconv.FormatUint(campaignId, 10)

	claimStartDate := startTime.Add(time.Second * 25)
	randMissionName := utils.RandStringOfLength(10)
	randMissionDescription := utils.RandStringOfLength(10)
	node.AddMissionError(campaignIdString, randMissionName, randMissionDescription, types.MissionInitialClaim.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName, "there can be only one mission with InitialClaim type and must be first in the campaign")
	node.AddMission(campaignIdString, randMissionName, randMissionDescription, types.MissionClaim.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName)
	node.AddMission(campaignIdString, randMissionName, randMissionDescription, types.MissionToDefine.String(), "0.3",
		claimStartDate.Format(cfeclaimcli.TimeLayout), creatorWalletName)

	balances, err := node.QueryBalances(creatorAddress)
	require.NoError(s.T(), err)
	claimRecordEntries, claimRecordEntriesWalletNames := prepareNClaimRecords(node, 10, balances.AmountOf(appparams.MicroC4eUnit))
	userEntriesJSONString, err := util.NewClaimRecordsListJson(claimRecordEntries)
	s.NoError(err)
	node.BankSendBaseBalanceFromNode(creatorAddress)
	node.AddClaimRecords(campaignIdString, userEntriesJSONString, creatorWalletName)
	node.EnableCampaign(campaignIdString, "", "", creatorWalletName)
	node.WaitUntilSpecifiedTime(claimStartDate)
	destinationAddress := testcosmos.CreateRandomAccAddress()
	claimer := claimRecordEntriesWalletNames[2]
	node.ClaimInitialMission(campaignIdString, destinationAddress, claimer)
	node.ClaimMission(campaignIdString, "1", claimer)
	node.ClaimMissionError(campaignIdString, "2", claimer, "cannot claim mission with type TO_DEFINE: mission claiming error")
}

func prepareNClaimRecords(node *chain.NodeConfig, n int, allEntriesAmountSum math.Int) (claimRecordEntries []types.ClaimRecordEntry, claimRecordEntriesWalletNames []string) {
	amountPerClaimRecord := allEntriesAmountSum.Quo(math.NewInt(int64(n)))
	for i := 0; i < n; i++ {
		walletName := utils.RandStringOfLength(10)
		address := node.CreateWallet(walletName)
		claimRecordEntry := types.ClaimRecordEntry{
			UserEntryAddress: address,
			Amount:           sdk.NewCoins(sdk.NewCoin(appparams.MicroC4eUnit, amountPerClaimRecord)),
		}
		claimRecordEntries = append(claimRecordEntries, claimRecordEntry)
		claimRecordEntriesWalletNames = append(claimRecordEntriesWalletNames, walletName)
	}
	return
}
