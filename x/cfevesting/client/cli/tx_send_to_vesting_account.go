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

func CmdSendToVestingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-to-vesting-account [to-address] [vesting-pool-name] [amount] [restart-vesting]",
		Short: "Broadcast message sendToVestingAccount",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToAddress := args[0]
			argVestingPoolName := args[1]
			argAmount := args[2]
			argRestartVesting := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if argVestingPoolName == "" {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "vesting-pool name cannot be empty")
			}
			amountInt, ok := sdk.NewIntFromString(argAmount)
			if !ok || amountInt.LTE(sdk.ZeroInt()) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be a positive integer")
			}
			restartVestingBool, err := strconv.ParseBool(argRestartVesting)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "restart-vesting must be a boolean [1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False]")
			}

			msg := types.NewMsgSendToVestingAccount(
				clientCtx.GetFromAddress().String(),
				argToAddress,
				argVestingPoolName,
				amountInt,
				restartVestingBool,
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
