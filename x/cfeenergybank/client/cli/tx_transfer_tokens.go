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

func CmdTransferTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-tokens [address-from] [address-to] [amount] [token-id]",
		Short: "Broadcast message transferTokens",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddressFrom := args[0]
			argAddressTo := args[1]
			argAmount, err := cast.ToUint64E(args[2])
			argTokenId, err := cast.ToUint64E(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferTokens(
				clientCtx.GetFromAddress().String(),
				argAddressFrom,
				argAddressTo,
				argAmount,
				argTokenId,
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
