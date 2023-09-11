package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVestingsSummary() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Query vestings summary",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query vestings summary.
			
Query response data:
  vesting_all_amount          sum of all currently locked tokens in vesting pools and vesting accounts created with %s module
  vesting_in_pools_amount     sum of all currently locked tokens in vesting pools
  vesting_in_accounts_amount  sum of all currently locked tokens in vesting accounts created with %s module
  delegated_vesting_amount    sum of all currently locked and delegated tokens from vesting accounts created with %s module

Example:
$ %s query %s summary
`,
				types.ModuleName, types.ModuleName, types.ModuleName, version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryVestingsSummaryRequest{}

			res, err := queryClient.VestingsSummary(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
