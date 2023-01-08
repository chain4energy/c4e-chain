package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVestingPools() *cobra.Command {
	bech32PrefixAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()

	cmd := &cobra.Command{
		Use:   "vesting-pools [address]",
		Short: "Query all vesting pools of a given address",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for all vesting pools of a given address.

Arguments:
  [address] vesting pools owner address

Example:
$ %s query %s vesting-pools %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj
`,
				version.AppName, types.ModuleName, bech32PrefixAddr,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryVestingPoolsRequest{

				Address: reqAddress,
			}

			res, err := queryClient.VestingPools(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
