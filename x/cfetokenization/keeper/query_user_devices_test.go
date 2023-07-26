package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
)

func TestUserDevicesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUserDevices(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetUserDevicesRequest
		response *types.QueryGetUserDevicesResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetUserDevicesRequest{Owner: msgs[0].Owner},
			response: &types.QueryGetUserDevicesResponse{UserDevices: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetUserDevicesRequest{Owner: msgs[1].Owner},
			response: &types.QueryGetUserDevicesResponse{UserDevices: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetUserDevicesRequest{Owner: "abdc"},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.UserDevices(wctx, tc.request)
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

func TestUserDevicesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CfetokenizationKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUserDevices(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUserDevicesRequest {
		return &types.QueryAllUserDevicesRequest{
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
			resp, err := keeper.UserDevicesAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserDevices), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserDevices),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UserDevicesAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserDevices), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserDevices),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.UserDevicesAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.UserDevices),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.UserDevicesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
