package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateReferencePayloadLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-reference-payload-link [reference-id] [payload-hash]",
		Short: "Query CreateReferencePayloadLink",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReferenceId := args[0]
			reqPayloadHash := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCreateReferencePayloadLinkRequest{

				ReferenceId: reqReferenceId,
				PayloadHash: reqPayloadHash,
			}

			res, err := queryClient.CreateReferencePayloadLink(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
