package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateStorageKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-storage-key [target-acc-address] [reference-id]",
		Short: "Query createStorageKey",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqTargetAccAddress := args[0]
			reqReferenceId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCreateStorageKeyRequest{

				TargetAccAddress: reqTargetAccAddress,
				ReferenceId:      reqReferenceId,
			}

			res, err := queryClient.CreateStorageKey(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
