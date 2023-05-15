package cfeclaim_test

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/chain4energy/c4e-chain/x/cfeclaim"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	startTime := time.Now().Add(time.Hour)
	var genesisState = types.GenesisState{
		Params: types.DefaultParams(),
		Campaigns: []types.Campaign{
			{
				Id:                     0,
				Owner:                  "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
				Name:                   "Campaign 1",
				Description:            "Campaign 1 description",
				CampaignType:           2,
				FeegrantAmount:         math.NewInt(300),
				InitialClaimFreeAmount: math.NewInt(500),
				Enabled:                false,
				StartTime:              time.Now(),
				EndTime:                time.Now().Add(time.Hour),
				LockupPeriod:           1000,
				VestingPeriod:          1234,
			},
			{
				Id:                     1,
				Owner:                  "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
				Name:                   "Campaign 2",
				Description:            "Campaign 2 description",
				CampaignType:           1,
				FeegrantAmount:         math.NewInt(100),
				InitialClaimFreeAmount: math.NewInt(300),
				Enabled:                true,
				StartTime:              time.Now(),
				EndTime:                time.Now().Add(time.Hour),
				LockupPeriod:           1234,
				VestingPeriod:          1000,
			},
		},
		Missions: []types.Mission{
			{
				Id:             0,
				CampaignId:     0,
				Name:           "Missin 1",
				Description:    "Missin 1 description",
				MissionType:    2,
				Weight:         sdk.MustNewDecFromStr("0.5"),
				ClaimStartDate: &startTime,
			},
			{
				Id:             1,
				CampaignId:     1,
				Name:           "Missin 2",
				Description:    "Missin 2 description",
				MissionType:    1,
				Weight:         sdk.MustNewDecFromStr("0.4"),
				ClaimStartDate: nil,
			},
		},
		UsersEntries: []types.UserEntry{
			{
				Address:      "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
				ClaimAddress: "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
				ClaimRecords: []*types.ClaimRecord{
					{
						CampaignId:        0,
						Address:           "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
						Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1234))),
						CompletedMissions: []uint64{0, 1},
						ClaimedMissions:   []uint64{0},
					},
					{
						CampaignId:        1,
						Address:           "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
						Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
						CompletedMissions: []uint64{0},
						ClaimedMissions:   []uint64{0},
					},
				},
			},
			{
				Address:      "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
				ClaimAddress: "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
				ClaimRecords: []*types.ClaimRecord{
					{
						CampaignId:        0,
						Address:           "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
						Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1234))),
						CompletedMissions: []uint64{0, 1},
						ClaimedMissions:   []uint64{0},
					},
					{
						CampaignId:        1,
						Address:           "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
						Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
						CompletedMissions: []uint64{0, 1},
						ClaimedMissions:   []uint64{0, 1},
					},
				},
			},
		},

		// this line is used by starport scaffolding # genesis/test/state
	}
	k, ctx := keepertest.CfeclaimKeeper(t)
	cfeclaim.InitGenesis(ctx, *k, genesisState)
	got := cfeclaim.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UsersEntries, got.UsersEntries)
	require.ElementsMatch(t, genesisState.Missions, got.Missions)
	require.ElementsMatch(t, genesisState.Campaigns, got.Campaigns)
	require.ElementsMatch(t, genesisState.CampaignsTotalAmount, got.CampaignsTotalAmount)
	require.ElementsMatch(t, genesisState.CampaignsAmountLeft, got.CampaignsAmountLeft)
	// this line is used by starport scaffolding # genesis/test/assert
}
