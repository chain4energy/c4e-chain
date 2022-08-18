package cli

import (
	"github.com/spf13/cast"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdTransferTokensOptimally() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-tokens-optimally [address-from] [address-to] [amount] [token-name]",
		Short: "Broadcast message transferTokensOptimally",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddressFrom := args[0]
			argAddressTo := args[1]
			argAmount, err := cast.ToUint64E(args[2])
			argTokenName := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferTokensOptimally(
				clientCtx.GetFromAddress().String(),
				argAddressFrom,
				argAddressTo,
				argAmount,
				argTokenName,
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
