package cli

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/energybank/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListEnergyToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-energy-token",
		Short: "list all EnergyToken",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEnergyTokenRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EnergyTokenAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowEnergyToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-energy-token [id]",
		Short: "shows a EnergyToken",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetEnergyTokenRequest{
				Id: id,
			}

			res, err := queryClient.EnergyToken(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
