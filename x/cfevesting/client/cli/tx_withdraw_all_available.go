package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWithdrawAllAvailable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-all-available",
		Short: "Withdraw all available tokens from all vesting pools of the broadcaster's address.",
		Long: strings.TrimSpace(fmt.Sprintf(`Withdraw all available tokens from all vesting pools of the broadcaster's' address.
Token are available when vesting pool lock period expires.

Example:
$ %s tx %s withdraw-all-available --from mykey
`,
			version.AppName, types.ModuleName,
		),
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawAllAvailable(
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
