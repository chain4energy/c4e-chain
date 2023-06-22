package cli_test

import (
	"cosmossdk.io/math"
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chain4energy/c4e-chain/testutil/network"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeev/client/cli"
	"github.com/chain4energy/c4e-chain/x/cfeev/types"
)

func networkWithEnergyTransferObjects(t *testing.T, n int) (*network.Network, []types.EnergyTransfer) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		energyTransfer := types.EnergyTransfer{
			Id:                    uint64(i),
			EnergyTransferOfferId: 0,
			ChargerId:             "",
			OwnerAccountAddress:   "",
			DriverAccountAddress:  "",
			OfferedTariff:         0,
			Status:                0,
			Collateral:            math.NewInt(1000),
			EnergyToTransfer:      0,
			EnergyTransferred:     0,
			PaidDate:              time.Time{},
		}
		nullify.Fill(&energyTransfer)
		state.EnergyTransfers = append(state.EnergyTransfers, energyTransfer)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.EnergyTransfers
}

func TestShowEnergyTransfer(t *testing.T) {
	net, objs := networkWithEnergyTransferObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		id   string
		args []string
		err  error
		obj  types.EnergyTransfer
	}{
		{
			desc: "found",
			id:   fmt.Sprintf("%d", objs[0].Id),
			args: common,
			obj:  objs[0],
		},
		{
			desc: "not found",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{tc.id}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdEnergyTransfer(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryEnergyTransferResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.EnergyTransfer)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.EnergyTransfer),
				)
			}
		})
	}
}

func TestListEnergyTransfer(t *testing.T) {
	net, objs := networkWithEnergyTransferObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdAllEnergyTransfers(), args)
			require.NoError(t, err)
			var resp types.QueryEnergyTransfersResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.EnergyTransfers), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.EnergyTransfers),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdAllEnergyTransfers(), args)
			require.NoError(t, err)
			var resp types.QueryEnergyTransfersResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.EnergyTransfers), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.EnergyTransfers),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdAllEnergyTransfers(), args)
		require.NoError(t, err)
		var resp types.QueryEnergyTransfersResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.EnergyTransfers),
		)
	})
}
