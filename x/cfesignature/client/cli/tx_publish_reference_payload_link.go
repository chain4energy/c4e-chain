package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfesignature/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPublishReferencePayloadLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish-reference-payload-link [key] [value]",
		Short: "Broadcast message publishReferencePayloadLink",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argKey := args[0]
			argValue := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPublishReferencePayloadLink(
				clientCtx.GetFromAddress().String(),
				argKey,
				argValue,
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
