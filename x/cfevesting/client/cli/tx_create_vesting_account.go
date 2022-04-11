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

func CmdCreateVestingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vesting-account [to-address] [amount] [start-time] [end-time]",
		Short: "Broadcast message createVestingAccount",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToAddress := args[0]
			argAmount := args[1]
			argStartTime := args[2]
			argEndTime := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinsNormalized(argAmount)
			if err != nil {
				return err
			}

			startTime, err := strconv.ParseInt(argStartTime, 10, 64)
			if err != nil {
				return err
			}

			endTime, err := strconv.ParseInt(argEndTime, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVestingAccount(
				clientCtx.GetFromAddress().String(),
				argToAddress,
				amount,
				startTime,
				endTime,
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
