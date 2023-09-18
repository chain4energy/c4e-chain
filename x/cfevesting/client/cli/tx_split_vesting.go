package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSplitVesting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "split-vesting [to-address] [amount]",
		Short: "Broadcast message SplitVesting",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToAddress := args[0]
			argAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(argAmount)
			if err != nil {
				return err
			}

			msg := types.NewMsgSplitVesting(
				clientCtx.GetFromAddress().String(),
				argToAddress,
				coins,
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
