package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var _ = strconv.Itoa(0)

func CmdGetAccountInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-account-info [acc-address-string]",
		Short: "Query getAccountInfo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAccAddressString := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetAccountInfoRequest{

				AccAddressString: reqAccAddressString,
			}

			res, err := queryClient.GetAccountInfo(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
