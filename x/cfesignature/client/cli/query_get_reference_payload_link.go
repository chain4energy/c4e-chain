package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var _ = strconv.Itoa(0)

func CmdGetReferencePayloadLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-reference-payload-link [reference-id]",
		Short: "Query getReferencePayloadLink",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReferenceId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetReferencePayloadLinkRequest{

				ReferenceId: reqReferenceId,
			}

			res, err := queryClient.GetReferencePayloadLink(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
