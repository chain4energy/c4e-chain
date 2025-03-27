package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/compatibility/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCfevestingV2VestingPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cfevesting-v-2-vesting-pools [address]",
		Short: "Query cfevesting-v2-vesting-pools",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCfevestingV2VestingPoolsRequest{

				Address: reqAddress,
			}

			res, err := queryClient.CfevestingV2VestingPools(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
