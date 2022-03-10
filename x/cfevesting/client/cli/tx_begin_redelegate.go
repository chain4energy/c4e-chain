package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBeginRedelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "begin-redelegate [validator-src-address] [validator-dst-address] [amount]",
		Short: "Broadcast message beginRedelegate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argValidatorSrcAddress := args[0]
			argValidatorDstAddress := args[1]
			argAmount := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(argAmount)
			if err != nil {
				return err
			}

			_, err = sdk.ValAddressFromBech32(argValidatorSrcAddress)
			if err != nil {
				return err
			}

			_, err = sdk.ValAddressFromBech32(argValidatorDstAddress)
			if err != nil {
				return err
			}

			msg := types.NewMsgBeginRedelegate(
				clientCtx.GetFromAddress().String(),
				argValidatorSrcAddress,
				argValidatorDstAddress,
				amount,
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
