package cli

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [amount]",
		Short: "Burns a specified amount of tokens",
		Long: strings.TrimSpace(fmt.Sprintf(`Burns a specified amount of tokens from the given address. This process permanently reduces the total supply of the tokens.

Arguments:
  [amount]            amount of tokens to burn

Example:
$ %s tx %s burn 100uc4e --from mykey
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurn(
				clientCtx.GetFromAddress().String(),
				argAmount,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
