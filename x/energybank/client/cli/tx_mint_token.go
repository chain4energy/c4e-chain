package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/energybank/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMintToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-token [name] [amount] [user-address]",
		Short: "Broadcast message mint-token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argAmount, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argUserAddress := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintToken(
				clientCtx.GetFromAddress().String(),
				argName,
				argAmount,
				argUserAddress,
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
