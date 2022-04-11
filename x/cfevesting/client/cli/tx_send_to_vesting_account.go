package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

var _ = strconv.Itoa(0)

func CmdSendToVestingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-to-vesting-account [to-address] [vesting-id] [amount] [restart-vesting]",
		Short: "Broadcast message sendToVestingAccount",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToAddress := args[0]
			argVestingId := args[1]
			argAmount := args[2]
			argRestartVesting := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			vestingIdInt32, err := strconv.ParseInt(argVestingId, 10, 32)
			if err != nil || vestingIdInt32 < 1 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "vesting-id must be a positive integer")
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
				int32(vestingIdInt32),
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
