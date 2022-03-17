package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vest [amount] [vesting-type]",
		Short: "Broadcast message vest",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argVestingType := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			amountInt, ok := sdk.NewIntFromString(argAmount)
			if !ok {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be a positive integer")
			}
			msg := types.NewMsgVest(
				clientCtx.GetFromAddress().String(),
				amountInt,
				argVestingType,
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
