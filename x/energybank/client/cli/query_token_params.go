package cli

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/energybank/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListTokenParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-token-params",
		Short: "list all tokenParams",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTokenParamsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TokenParamsAll(context.Background(), params)
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

func CmdShowTokenParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-token-params [index]",
		Short: "shows a tokenParams",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetTokenParamsRequest{
				Index: argIndex,
			}

			res, err := queryClient.TokenParams(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
