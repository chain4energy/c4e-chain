package keeper_test

import (
	"strconv"
)

// Prevent strconv unused error
var _ = strconv.IntSize

//
//func TestClaimRecordQuerySingle(t *testing.T) {
//	keeper, ctx := keepertest.CfeairdropKeeper(t)
//	wctx := sdk.WrapSDKContext(ctx)
//	msgs := createNUsersEntries(keeper, ctx, 2, 10, true, true)
//	for _, tc := range []struct {
//		desc     string
//		request  *types.QueryUsersEntriesRequest
//		response *types.QueryUsersEntriesResponse
//		err      error
//	}{
//		{
//			desc: "First",
//			request: &types.QueryUsersEntriesRequest{
//				Address: msgs[0].Address,
//			},
//			response: &types.QueryUsersEntriesResponse{UserEntry: msgs[0]},
//		},
//		{
//			desc: "Second",
//			request: &types.QueryUsersEntriesRequest{
//				Address: msgs[1].Address,
//			},
//			response: &types.QueryUsersEntriesResponse{UserEntry: msgs[1]},
//		},
//		{
//			desc: "KeyNotFound",
//			request: &types.QueryUsersEntriesRequest{
//				Address: strconv.Itoa(100000),
//			},
//			err: status.Error(codes.NotFound, "not found"),
//		},
//		{
//			desc: "InvalidRequest",
//			err:  status.Error(codes.InvalidArgument, "invalid request"),
//		},
//	} {
//		t.Run(tc.desc, func(t *testing.T) {
//			response, err := keeper.UserEntry(wctx, tc.request)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//				require.Equal(t,
//					nullify.Fill(tc.response),
//					nullify.Fill(response),
//				)
//			}
//		})
//	}
//}
//
//func TestClaimRecordQueryPaginated(t *testing.T) {
//	keeper, ctx := keepertest.CfeairdropKeeper(t)
//	wctx := sdk.WrapSDKContext(ctx)
//	msgs := createNUsersEntries(keeper, ctx, 5, 0, false, false)
//
//	request := func(next []byte, offset, limit uint64, total bool) *types.QueryUsersEntriesRequest {
//		return &types.QueryUsersEntriesRequest{
//			Pagination: &query.PageRequest{
//				Key:        next,
//				Offset:     offset,
//				Limit:      limit,
//				CountTotal: total,
//			},
//		}
//	}
//	t.Run("ByOffset", func(t *testing.T) {
//		step := 2
//		for i := 0; i < len(msgs); i += step {
//			resp, err := keeper.UsersEntries(wctx, request(nil, uint64(i), uint64(step), false))
//			require.NoError(t, err)
//			require.LessOrEqual(t, len(resp.UsersEntries), step)
//			require.Subset(t,
//				nullify.Fill(msgs),
//				nullify.Fill(resp.UsersEntries),
//			)
//		}
//	})
//	t.Run("ByKey", func(t *testing.T) {
//		step := 2
//		var next []byte
//		for i := 0; i < len(msgs); i += step {
//			resp, err := keeper.UsersEntries(wctx, request(next, 0, uint64(step), false))
//			require.NoError(t, err)
//			require.LessOrEqual(t, len(resp.UsersEntries), step)
//			require.Subset(t,
//				nullify.Fill(msgs),
//				nullify.Fill(resp.UsersEntries),
//			)
//			next = resp.Pagination.NextKey
//		}
//	})
//	t.Run("Total", func(t *testing.T) {
//		resp, err := keeper.UsersEntries(wctx, request(nil, 0, 0, true))
//		require.NoError(t, err)
//		require.Equal(t, len(msgs), int(resp.Pagination.Total))
//		require.ElementsMatch(t,
//			nullify.Fill(msgs),
//			nullify.Fill(resp.UsersEntries),
//		)
//	})
//	t.Run("InvalidRequest", func(t *testing.T) {
//		_, err := keeper.UsersEntries(wctx, nil)
//		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
//	})
//}