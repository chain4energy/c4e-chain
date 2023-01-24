package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state",
		Short: fmt.Sprintf("Query current %s state", types.ModuleName),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the current %s state.

Example:
$ %s query %s state
`,
				types.ModuleName, version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStateRequest{}

			res, err := queryClient.State(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
