package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	testcosmos "github.com/chain4energy/c4e-chain/testutil/cosmossdk"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"
)

const (
	vPool1 = "v-pool-1"
	vPool2 = "v-pool-2"
)

func TestCampaigns(t *testing.T) {
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)

	correctName := "Name"
	correctDescription := "Description"
	startTime := testenv.TestEnvTime
	endTime := startTime.Add(time.Hour)
	lockupPeriod := time.Hour
	vestingPeriod := time.Hour

	for _, tc := range []struct {
		desc                   string
		Name                   string
		Description            string
		InitialClaimFreeAmount math.Int
		FeegrantAmount         math.Int
		StartTime              time.Time
		EndTime                time.Time
		LockupPeriod           time.Duration
		VestingPeriod          time.Duration
		CampaignType           types.CampaignType
		vestingPoolName        string
		blockTime              time.Time
		error                  bool
		expectedError          string
	}{
		{
			desc:            "create campaign",
			Name:            correctName,
			Description:     correctDescription,
			StartTime:       startTime,
			EndTime:         endTime,
			LockupPeriod:    lockupPeriod,
			VestingPeriod:   vestingPeriod,
			CampaignType:    types.DefaultCampaign,
			vestingPoolName: "",
			blockTime:       startTime,
			error:           false,
			expectedError:   "",
		},
		{
			desc:            "create campaign - empty name",
			Name:            "",
			Description:     correctDescription,
			StartTime:       startTime,
			EndTime:         endTime,
			LockupPeriod:    lockupPeriod,
			VestingPeriod:   vestingPeriod,
			CampaignType:    types.DefaultCampaign,
			vestingPoolName: "",
			blockTime:       startTime,
			error:           true,
			expectedError:   "campaign name is empty: wrong param value",
		},
		{
			desc:            "create campaign - start time equal end time",
			Name:            correctName,
			Description:     correctDescription,
			StartTime:       startTime,
			EndTime:         startTime,
			LockupPeriod:    lockupPeriod,
			VestingPeriod:   vestingPeriod,
			CampaignType:    types.DefaultCampaign,
			vestingPoolName: "",
			blockTime:       startTime,
			error:           true,
			expectedError:   fmt.Sprintf("start time is equal to end time (%s = %s): wrong param value", startTime, startTime),
		},
		{
			desc:            "create campaign - start time after end time",
			Name:            correctName,
			Description:     correctDescription,
			StartTime:       startTime.Add(time.Hour),
			EndTime:         startTime,
			LockupPeriod:    lockupPeriod,
			VestingPeriod:   vestingPeriod,
			CampaignType:    types.DefaultCampaign,
			vestingPoolName: "",
			blockTime:       startTime,
			error:           true,
			expectedError:   fmt.Sprintf("start time is after end time (%s > %s): wrong param value", startTime.Add(time.Hour), startTime),
		},
		{
			desc:            "create campaign - start time in the past",
			Name:            correctName,
			Description:     correctDescription,
			StartTime:       startTime.Add(-time.Hour),
			EndTime:         startTime,
			LockupPeriod:    lockupPeriod,
			VestingPeriod:   vestingPeriod,
			CampaignType:    types.DefaultCampaign,
			vestingPoolName: "",
			blockTime:       startTime,
			error:           true,
			expectedError:   fmt.Sprintf("start time in the past error (%s < %s): wrong param value", startTime.Add(-time.Hour), startTime.UTC()),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			testHelper := testapp.SetupTestAppWithHeight(t, 1000)
			testHelper.SetContextBlockTime(tc.blockTime)

			campaign := types.Campaign{
				Name:                   tc.Name,
				Description:            tc.Description,
				StartTime:              tc.StartTime,
				EndTime:                tc.EndTime,
				LockupPeriod:           tc.LockupPeriod,
				VestingPeriod:          tc.VestingPeriod,
				CampaignType:           tc.CampaignType,
				VestingPoolName:        tc.vestingPoolName,
				InitialClaimFreeAmount: tc.InitialClaimFreeAmount,
				FeegrantAmount:         tc.FeegrantAmount,
			}

			if tc.error {
				testHelper.C4eClaimUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, tc.expectedError)
			} else {
				testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
			}
		})
	}
}

func TestCreateManyClaimCampaigns(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaigns := prepareNTestCampaigns(testHelper.Context, 10)
	for _, campaign := range campaigns {
		testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	}
}

func TestCreateCampaignNegativeInitialClaimAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.InitialClaimFreeAmount = sdk.NewInt(-100)
	testHelper.C4eClaimUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("initial claim free amount (%s) cannot be negative: wrong param value", campaign.InitialClaimFreeAmount))
}

func TestCreateCampaignNegativeFeegrantAmount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(-100)
	testHelper.C4eClaimUtils.CreateCampaignError(acountsAddresses[0].String(), campaign, fmt.Sprintf("feegrant amount (%s) cannot be negative: wrong param value", campaign.FeegrantAmount))
}

func TestCreateCampaignAndStart(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
}

func TestCreateManyCampaignsAndStart(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaigns := prepareNTestCampaigns(testHelper.Context, 10)
	for i, campaign := range campaigns {
		testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
		testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), uint64(i), nil, nil)
	}
}

func TestCreateCampaignAndStartTimeAfterTimeNowError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	blockTime := campaign.StartTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.StartCampaignError(acountsAddresses[0].String(), 0, nil, nil, fmt.Sprintf("start time in the past error (%s < %s): wrong param value", campaign.StartTime, blockTime))
}

func TestCreateCampaignAndStartOwnerNotValidError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaignError(acountsAddresses[1].String(), 0, nil, nil, "you are not the campaign owner: tx intended signer does not match the given signer")
}

func TestCreateCampaignCampaignDoesntExistError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaignError(acountsAddresses[0].String(), 1, nil, nil, "campaign with id 1 not found: entity does not exist")
}

func TestCreateCampaignCampaignEnabledError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.StartCampaignError(acountsAddresses[0].String(), 0, nil, nil, "campaign is enabled")
}

func TestCreateVestingPoolCampaign(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(ownerAddress.String(), 0, nil, nil)
}

func TestCreateVestingPoolCampaignWrongPoolName(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool2

	testHelper.C4eClaimUtils.CreateCampaignError(ownerAddress.String(), campaign, fmt.Sprintf("vesting pool %s not found for address %s: entity does not exist", vPool2, acountsAddresses[0]))
}

func TestCreateVestingPoolCampaignWrongVestingPeriod(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	campaign.VestingPeriod = time.Minute
	testHelper.C4eClaimUtils.CreateCampaignError(ownerAddress.String(), campaign, fmt.Sprintf("the duration of campaign vesting period must be equal to or greater than the vesting type vesting period (%s > %s): wrong param value",
		(time.Hour*100).String(), campaign.VestingPeriod.String()))
}

func TestCreateVestingPoolCampaignWrongLockupPeriod(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	campaign.LockupPeriod = time.Minute
	testHelper.C4eClaimUtils.CreateCampaignError(ownerAddress.String(), campaign, fmt.Sprintf("the duration of campaign lockup period must be equal to or greater than the vesting type lockup period (%s > %s): wrong param value",
		(time.Hour*100).String(), campaign.LockupPeriod.String()))
}

func TestCreateVestingPoolCampaignWrongOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eClaimUtils.CreateCampaignError(acountsAddresses[1].String(), campaign, fmt.Sprintf("vesting pool %s not found for address %s: entity does not exist", vPool1, acountsAddresses[1]))
}

func TestCreateVestingPoolCampaignWrongType(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	ownerAddress := acountsAddresses[0]
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.DynamicCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eClaimUtils.CreateCampaignError(ownerAddress.String(), campaign, "vesting pool name can be set only for VESTING_POOL type campaigns: wrong param value")
}

func TestCreateCampaignCloseCloseActionBurn(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CloseBurn)
}

func TestCreateCampaignCloseCloseActionSendToOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CloseSendToOwner)
}

func TestCreateCampaignVestingPoolCloseCloseActionSendToOwner(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	ownerAddress := acountsAddresses[0]
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)

	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(ownerAddress.String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 300)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(ownerAddress.String(), 0, types.CloseSendToOwner)
}

func TestCreateCampaignCloseCloseActionSendToCommunityPool(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CloseSendToCommunityPool)
}

func TestCreateVestingPoolCampaignCloseCloseActionSendToCommunityPool(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	ownerAddress := acountsAddresses[0]

	campaign := prepareTestCampaign(testHelper.Context)
	campaign.CampaignType = types.VestingPoolCampaign
	campaign.VestingPoolName = vPool1
	testHelper.C4eVestingUtils.AddTestVestingPool(ownerAddress, vPool1, math.NewInt(10000), 100, 100)
	testHelper.C4eClaimUtils.CreateCampaign(ownerAddress.String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(ownerAddress.String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 300)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(ownerAddress, amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(ownerAddress, 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaignError(ownerAddress.String(), 0, types.CloseSendToCommunityPool, "in the case of sale campaigns and campaigns created from whitelist vesting accounts, it is not possible to use sendToCommunityPool close action: invalid type")
}

func TestCreateCampaignCloseCloseActionBurnAndFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(1000)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Add(campaign.FeegrantAmount.MulRaw(int64(len(claimEntries)))))
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CloseBurn)
}

func TestCreateCampaignCloseCloseActionSendToOwnerAndFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(1000)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Add(campaign.FeegrantAmount.MulRaw(int64(len(claimEntries)))))
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CloseSendToOwner)
}

func TestCreateCampaignCloseCloseActionSendToCommunityPoolAndFeegrant(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	campaign.FeegrantAmount = sdk.NewInt(1000)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum.Add(campaign.FeegrantAmount.MulRaw(int64(len(claimEntries)))))
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), 0, types.CloseSendToCommunityPool)
}

func TestCreateCampaignCloseCampaignWrongCloseAction(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	claimEntries, amountSum := createTestClaimRecords(acountsAddresses, 100000000)
	testHelper.C4eClaimUtils.AddCoinsToCampaignOwnerAcc(acountsAddresses[0], amountSum)
	testHelper.C4eClaimUtils.AddClaimRecords(acountsAddresses[0], 0, claimEntries)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaignError(acountsAddresses[0].String(), 0, types.CloseActionUnspecified, "wrong campaign close action type: invalid type")
}

func TestCreateManyCampaignsAndClose(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaigns := prepareNTestCampaigns(testHelper.Context, 10)
	contextNow := testHelper.Context.BlockTime()
	for i, campaign := range campaigns {
		testHelper.SetContextBlockTime(contextNow)
		testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
		testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), uint64(i), nil, nil)
		blockTime := campaign.EndTime.Add(time.Minute)
		testHelper.SetContextBlockTime(blockTime)
		testHelper.C4eClaimUtils.CloseCampaign(acountsAddresses[0].String(), uint64(i), types.CloseBurn)
	}
}

func TestCreateCampaignCloseCampaignCampaignNotOverYetError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	testHelper.C4eClaimUtils.CloseCampaignError(acountsAddresses[0].String(), 0, types.CloseBurn, fmt.Sprintf("campaign with id %d campaign is not over yet (endtime - %s < %s): wrong param value", 0, campaign.EndTime, testHelper.Context.BlockTime()))
}

func TestCreateCampaignCloseCampaignCampaignDoesntExistError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaignError(acountsAddresses[0].String(), 1, types.CloseAction_CLOSE_ACTION_UNSPECIFIED, "campaign with id 1 not found: entity does not exist")
}

func TestCreateCampaignCloseCampaignYouAreNotTheOwnerError(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
	campaign := prepareTestCampaign(testHelper.Context)
	testHelper.C4eClaimUtils.CreateCampaign(acountsAddresses[0].String(), campaign)
	testHelper.C4eClaimUtils.StartCampaign(acountsAddresses[0].String(), 0, nil, nil)
	blockTime := campaign.EndTime.Add(time.Minute)
	testHelper.SetContextBlockTime(blockTime)
	testHelper.C4eClaimUtils.CloseCampaignError(acountsAddresses[1].String(), 0, types.CloseAction_CLOSE_ACTION_UNSPECIFIED, "you are not the campaign owner: tx intended signer does not match the given signer")
}

func prepareNTestCampaigns(ctx sdk.Context, n int) []types.Campaign {
	campaigns := make([]types.Campaign, n)
	for i := range campaigns {
		campaigns[i] = prepareTestCampaign(ctx)
	}
	return campaigns
}

func prepareTestCampaign(ctx sdk.Context) types.Campaign {
	start := ctx.BlockTime()
	end := ctx.BlockTime().Add(time.Second * 10)
	lockupPeriod := time.Hour * 10000
	vestingPeriod := 3 * time.Hour * 10000
	return types.Campaign{
		Id:            0,
		Name:          "Name",
		Description:   "test-campaign",
		StartTime:     start,
		EndTime:       end,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
		CampaignType:  types.DefaultCampaign,
	}
}
