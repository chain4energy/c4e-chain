package cli_test

import (
	"strconv"
	"testing"

	"github.com/chain4energy/c4e-chain/testutil/network"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithInitialClaimObjects(t *testing.T, n int) (*network.Network, []types.InitialClaim) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		initialClaim := types.InitialClaim{
			CampaignId: uint64(i),
		}
		nullify.Fill(&initialClaim)
		state.InitialClaims = append(state.InitialClaims, initialClaim)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.InitialClaims
}

//
//func TestShowInitialClaim(t *testing.T) {
//	net, objs := networkWithInitialClaimObjects(t, 2)
//
//	ctx := net.Validators[0].ClientCtx
//	common := []string{
//		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
//	}
//	for _, tc := range []struct {
//		desc         string
//		idCampaignId string
//
//		args []string
//		err  error
//		obj  types.InitialClaim
//	}{
//		{
//			desc:         "found",
//			idCampaignId: strconv.FormatUint(objs[0].CampaignId, 10),
//
//			args: common,
//			obj:  objs[0],
//		},
//		{
//			desc:         "not found",
//			idCampaignId: strconv.Itoa(100000),
//
//			args: common,
//			err:  status.Error(codes.NotFound, "not found"),
//		},
//	} {
//		t.Run(tc.desc, func(t *testing.T) {
//			args := []string{
//				tc.idCampaignId,
//			}
//			args = append(args, tc.args...)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowInitialClaim(), args)
//			if tc.err != nil {
//				stat, ok := status.FromError(tc.err)
//				require.True(t, ok)
//				require.ErrorIs(t, stat.Err(), tc.err)
//			} else {
//				require.NoError(t, err)
//				var resp types.QueryGetInitialClaimResponse
//				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//				require.NotNil(t, resp.InitialClaim)
//				require.Equal(t,
//					nullify.Fill(&tc.obj),
//					nullify.Fill(&resp.InitialClaim),
//				)
//			}
//		})
//	}
//}
//
//func TestListInitialClaim(t *testing.T) {
//	net, objs := networkWithInitialClaimObjects(t, 5)
//
//	ctx := net.Validators[0].ClientCtx
//	request := func(next []byte, offset, limit uint64, total bool) []string {
//		args := []string{
//			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
//		}
//		if next == nil {
//			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
//		} else {
//			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
//		}
//		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
//		if total {
//			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
//		}
//		return args
//	}
//	t.Run("ByOffset", func(t *testing.T) {
//		step := 2
//		for i := 0; i < len(objs); i += step {
//			args := request(nil, uint64(i), uint64(step), false)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListInitialClaim(), args)
//			require.NoError(t, err)
//			var resp types.QueryAllInitialClaimResponse
//			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//			require.LessOrEqual(t, len(resp.InitialClaim), step)
//			require.Subset(t,
//				nullify.Fill(objs),
//				nullify.Fill(resp.InitialClaim),
//			)
//		}
//	})
//	t.Run("ByKey", func(t *testing.T) {
//		step := 2
//		var next []byte
//		for i := 0; i < len(objs); i += step {
//			args := request(next, 0, uint64(step), false)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListInitialClaim(), args)
//			require.NoError(t, err)
//			var resp types.QueryAllInitialClaimResponse
//			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//			require.LessOrEqual(t, len(resp.InitialClaim), step)
//			require.Subset(t,
//				nullify.Fill(objs),
//				nullify.Fill(resp.InitialClaim),
//			)
//			next = resp.Pagination.NextKey
//		}
//	})
//	t.Run("Total", func(t *testing.T) {
//		args := request(nil, 0, uint64(len(objs)), true)
//		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListInitialClaim(), args)
//		require.NoError(t, err)
//		var resp types.QueryAllInitialClaimResponse
//		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//		require.NoError(t, err)
//		require.Equal(t, len(objs), int(resp.Pagination.Total))
//		require.ElementsMatch(t,
//			nullify.Fill(objs),
//			nullify.Fill(resp.InitialClaim),
//		)
//	})
//}
