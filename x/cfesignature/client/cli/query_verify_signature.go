package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVerifySignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-signature [reference-id] [target-acc-address]",
		Short: "Query VerifySignature",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReferenceId := args[0]
			reqTargetAccAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryVerifySignatureRequest{

				ReferenceId:      reqReferenceId,
				TargetAccAddress: reqTargetAccAddress,
			}

			res, err := queryClient.VerifySignature(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
