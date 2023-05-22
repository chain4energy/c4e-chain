package cfeclaim_test

//func TestGenesis(t *testing.T) {
//	startTime := time.Now().Add(time.Hour)
//	testHelper := testapp.SetupTestAppWithHeight(t, 1000)
//
//	acountsAddresses, _ := testcosmos.CreateAccounts(2, 0)
//	testHelper.C4eVestingUtils.AddTestVestingPool(acountsAddresses[0], "Vesting1", math.NewInt(10000), 100, 100)
//	var genesisState = types.GenesisState{
//		Params: types.DefaultParams(),
//		Campaigns: []types.Campaign{
//			{
//				Id:                     0,
//				Owner:                  acountsAddresses[0].String(),
//				Name:                   "Campaign 1",
//				Description:            "Campaign 1 description",
//				CampaignType:           2,
//				FeegrantAmount:         math.NewInt(300),
//				InitialClaimFreeAmount: math.NewInt(500),
//				Enabled:                false,
//				StartTime:              time.Now(),
//				EndTime:                time.Now().Add(time.Hour),
//				LockupPeriod:           time.Hour * 10000,
//				VestingPeriod:          time.Hour * 10000,
//				VestingPoolName:        "Vesting1",
//			},
//			{
//				Id:                     1,
//				Owner:                  acountsAddresses[1].String(),
//				Name:                   "Campaign 2",
//				Description:            "Campaign 2 description",
//				CampaignType:           1,
//				FeegrantAmount:         math.NewInt(100),
//				InitialClaimFreeAmount: math.NewInt(300),
//				Enabled:                true,
//				StartTime:              time.Now(),
//				EndTime:                time.Now().Add(time.Hour),
//				LockupPeriod:           1234,
//				VestingPeriod:          1000,
//			},
//		},
//		Missions: []types.Mission{
//			{
//				Id:             0,
//				CampaignId:     0,
//				Name:           "Missin 1",
//				Description:    "Missin 1 description",
//				MissionType:    2,
//				Weight:         sdk.MustNewDecFromStr("0.5"),
//				ClaimStartDate: &startTime,
//			},
//			{
//				Id:             1,
//				CampaignId:     1,
//				Name:           "Missin 2",
//				Description:    "Missin 2 description",
//				MissionType:    1,
//				Weight:         sdk.MustNewDecFromStr("0.4"),
//				ClaimStartDate: nil,
//			},
//		},
//		UsersEntries: []types.UserEntry{
//			{
//				Address:      "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
//				ClaimAddress: "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
//				ClaimRecords: []*types.ClaimRecord{
//					{
//						CampaignId:        0,
//						Address:           "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
//						Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1234))),
//						CompletedMissions: []uint64{0, 1},
//						ClaimedMissions:   []uint64{0},
//					},
//					{
//						CampaignId:        1,
//						Address:           "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
//						Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
//						CompletedMissions: []uint64{0},
//						ClaimedMissions:   []uint64{0},
//					},
//				},
//			},
//			{
//				Address:      "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
//				ClaimAddress: "cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj",
//				ClaimRecords: []*types.ClaimRecord{
//					{
//						CampaignId:        0,
//						Address:           "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
//						Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(1234))),
//						CompletedMissions: []uint64{0, 1},
//						ClaimedMissions:   []uint64{0},
//					},
//					{
//						CampaignId:        1,
//						Address:           "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
//						Amount:            sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000))),
//						CompletedMissions: []uint64{0, 1},
//						ClaimedMissions:   []uint64{0, 1},
//					},
//				},
//			},
//		},
//
//		// this line is used by starport scaffolding # genesis/test/state
//	}
//
//	testHelper.C4eClaimUtils.C4eClaimUtils.InitGenesis(testHelper.Context, genesisState)
//	testHelper.C4eClaimUtils.C4eClaimKeeperUtils.InitGenesis(testHelper.Context, genesisState)
//	testHelper.C4eClaimUtils.C4eClaimKeeperUtils.ExportGenesis(testHelper.Context, genesisState)
//	// this line is used by starport scaffolding # genesis/test/assert
//}
