package cli

import (
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05 -0700 MST"
)

func CmdAddMeasurement() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-measurement [timestamp] [active-power] [reverse-power]",
		Short: "Create a new UserDevices",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTimestamp, err := time.Parse(TimeLayout, args[0])
			if err != nil {
				return err
			}

			argActivePower, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argReversePower, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddMeasurement(clientCtx.GetFromAddress().String(), &argTimestamp,
				argActivePower, argReversePower)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
