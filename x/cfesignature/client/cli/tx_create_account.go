package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdCreateAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-account [acc-address-string] [pub-key-string]",
		Short: "Broadcast message createAccount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAccAddressString := args[0]
			argPubKeyString := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAccount(
				clientCtx.GetFromAddress().String(),
				argAccAddressString,
				argPubKeyString,
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
