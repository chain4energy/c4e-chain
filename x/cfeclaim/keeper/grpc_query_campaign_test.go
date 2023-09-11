package keeper_test

import (
	"fmt"
	keepertest "github.com/chain4energy/c4e-chain/v2/testutil/keeper"
	"github.com/chain4energy/c4e-chain/v2/testutil/nullify"
	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"testing"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCampaignQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CfeclaimKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createAndSaveNTestCampaigns(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryCampaignRequest
		response *types.QueryCampaignResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryCampaignRequest{
				CampaignId: msgs[0].Id,
			},
			response: &types.QueryCampaignResponse{Campaign: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryCampaignRequest{
				CampaignId: msgs[1].Id,
			},
			response: &types.QueryCampaignResponse{Campaign: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryCampaignRequest{
				CampaignId: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Campaign(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestCampaignQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CfeclaimKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createAndSaveNTestCampaigns(keeper, ctx, 5)
	campaigns := keeper.GetAllCampaigns(ctx)
	fmt.Println("campaigns", campaigns)
	fmt.Println("campaignsLen", len(campaigns))
	request := func(next []byte, offset, limit uint64, total bool) *types.QueryCampaignsRequest {
		return &types.QueryCampaignsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.Campaigns(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Campaigns), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Campaigns),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.Campaigns(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Campaigns), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Campaigns),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.Campaigns(wctx, request(nil, 0, 0, true))
		fmt.Println(resp)
		fmt.Println(campaigns)
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Campaigns),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.Missions(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
