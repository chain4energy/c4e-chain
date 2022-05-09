package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateReferenceId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-reference-id [creator]",
		Short: "Query CreateReferenceId",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqCreator := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCreateReferenceIdRequest{

				Creator: reqCreator,
			}

			res, err := queryClient.CreateReferenceId(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
