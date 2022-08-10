package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCurrentBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "current-balance [user-address] [token-name]",
		Short: "Query current-balance",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqUserAddress := args[0]
			reqTokenName := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCurrentBalanceRequest{

				UserAddress: reqUserAddress,
				TokenName:   reqTokenName,
			}

			res, err := queryClient.CurrentBalance(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
