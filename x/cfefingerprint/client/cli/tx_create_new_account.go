package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

// Temporary message - to remove in the future
func CmdCreateNewAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-new-account [acc-address-string]",
		Short: "Broadcast message create-new-account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAccAddressString := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNewAccount(
				clientCtx.GetFromAddress().String(),
				argAccAddressString,
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
