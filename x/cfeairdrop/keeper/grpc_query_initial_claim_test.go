package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestInitialClaimQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNInitialClaim(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetInitialClaimRequest
		response *types.QueryGetInitialClaimResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetInitialClaimRequest{
				CampaignId: msgs[0].CampaignId,
			},
			response: &types.QueryGetInitialClaimResponse{InitialClaim: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetInitialClaimRequest{
				CampaignId: msgs[1].CampaignId,
			},
			response: &types.QueryGetInitialClaimResponse{InitialClaim: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetInitialClaimRequest{
				CampaignId: uint64(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.InitialClaim(wctx, tc.request)
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

func TestInitialClaimQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNInitialClaim(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllInitialClaimRequest {
		return &types.QueryAllInitialClaimRequest{
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
			resp, err := keeper.InitialClaimAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.InitialClaim), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.InitialClaim),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.InitialClaimAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.InitialClaim), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.InitialClaim),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.InitialClaimAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.InitialClaim),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.InitialClaimAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
