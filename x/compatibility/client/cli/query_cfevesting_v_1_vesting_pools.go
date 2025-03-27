package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/compatibility/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCfevestingV1VestingPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cfevesting-v-1-vesting-pools [address]",
		Short: "Query cfevesting-v1-vesting-pools",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCfevestingV1VestingPoolsRequest{

				Address: reqAddress,
			}

			res, err := queryClient.CfevestingV1VestingPools(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
