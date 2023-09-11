package keeper_test

import (
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

func TestMissionQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CfeclaimKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createAndSaveNTestMissions(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryMissionRequest
		response *types.QueryMissionResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryMissionRequest{
				CampaignId: msgs[0].CampaignId,
				MissionId:  msgs[0].Id,
			},
			response: &types.QueryMissionResponse{Mission: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryMissionRequest{
				CampaignId: msgs[1].CampaignId,
				MissionId:  msgs[1].Id,
			},
			response: &types.QueryMissionResponse{Mission: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryMissionRequest{
				CampaignId: 100000,
				MissionId:  100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Mission(wctx, tc.request)
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

func TestMissionQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CfeclaimKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createAndSaveNTestMissions(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryMissionsRequest {
		return &types.QueryMissionsRequest{
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
			resp, err := keeper.Missions(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Missions), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Missions),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.Missions(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Missions), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Missions),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.Missions(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Missions),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.Missions(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
