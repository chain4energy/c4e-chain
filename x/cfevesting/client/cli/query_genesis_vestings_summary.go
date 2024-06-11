package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGenesisVestingsSummary() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genesis-summary",
		Short: "Query genesis vestings summary",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query genesis vestings summary.
			
Query response data:
  vesting_all_amount          sum of all currently locked tokens in genesis vesting pools, genesis vesting accounts and vesting accounts created with %s module directly from genesis vesting pools or genesis vesting accounts 
  vesting_in_pools_amount     sum of all currently locked tokens in genesis vesting pools
  vesting_in_accounts_amount  sum of all currently locked tokens in genesis vesting accounts and vesting accounts created with %s module directly from genesis vesting pools or genesis vesting accounts
  delegated_vesting_amount    sum of all currently locked and delegated tokens from genesis vesting accounts and vesting accounts created with %s module directly from genesis vesting pools or genesis vesting account

Example:
$ %s query %s genesis-summary
`,
				types.ModuleName, types.ModuleName, types.ModuleName, version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGenesisVestingsSummaryRequest{}

			res, err := queryClient.GenesisVestingsSummary(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
