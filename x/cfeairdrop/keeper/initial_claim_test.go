package keeper_test

import (
	"strconv"
	"testing"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNInitialClaim(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.InitialClaim {
	items := make([]types.InitialClaim, n)
	for i := range items {
		items[i].CampaignId = uint64(i)

		keeper.SetInitialClaim(ctx, items[i])
	}
	return items
}

func TestInitialClaimGet(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNInitialClaim(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetInitialClaim(ctx,
			item.CampaignId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestInitialClaimRemove(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNInitialClaim(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInitialClaim(ctx,
			item.CampaignId,
		)
		_, found := keeper.GetInitialClaim(ctx,
			item.CampaignId,
		)
		require.False(t, found)
	}
}

func TestInitialClaimGetAll(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	items := createNInitialClaim(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllInitialClaim(ctx)),
	)
}

func TestClaimInitial(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	ctx := testHelper.Context

	end := testHelper.Context.BlockTime().Add(1000)
	params := types.Params{Campaigns: []*types.Campaign{
		{
			CampaignId:    1,
			Enabled:       true,
			StartTime:     testHelper.Context.BlockTime(),
			EndTime:       &end,
			LockupPeriod:  1000,
			VestingPeriod: 2000,
			Description:   "test-campaign",
		},
	}}
	initialClaims := []types.InitialClaim{{Enabled: true, CampaignId: 1, MissionId: 3}}
	missions := []types.Mission{{CampaignId: 1, MissionId: 3, Description: "test-mission", Weight: sdk.MustNewDecFromStr("0.2")}}
	genesisState := types.GenesisState{Params: params, InitialClaims: initialClaims, Missions: missions}
	cfeairdrop.InitGenesis(ctx, testHelper.App.CfeairdropKeeper, genesisState)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(sdk.NewInt(10000), types.ModuleName)

	records := []*keeper.CampaignRecordData{{Address: acountsAddresses[0].String(), Claimable: sdk.NewInt(10000)}}
	require.Nil(t, testHelper.App.AccountKeeper.GetAccount(ctx, acountsAddresses[0]))
	require.NoError(t, testHelper.App.CfeairdropKeeper.AddCampaignRecords(ctx, 1, records))
	require.Nil(t, testHelper.App.AccountKeeper.GetAccount(ctx, acountsAddresses[0]))

	require.NoError(t, testHelper.App.CfeairdropKeeper.ClaimInitial(ctx, 1, 3, acountsAddresses[0].String()))

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(acountsAddresses[0], sdk.NewInt(2000))
}
