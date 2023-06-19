package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVerifyReferencePayloadLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-reference-payload-link [reference-id] [payload-hash]",
		Short: "Verify reference payload link by passing reference id and payload hash",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReferenceId := args[0]
			reqPayloadHash := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryVerifyReferencePayloadLinkRequest{

				ReferenceId: reqReferenceId,
				PayloadHash: reqPayloadHash,
			}

			res, err := queryClient.VerifyReferencePayloadLink(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
