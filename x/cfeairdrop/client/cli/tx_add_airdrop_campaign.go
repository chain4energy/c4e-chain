package cli

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateAirdropCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-airdrop-campaign [name] [description] [start-time] [end-time] [lockup-period] [vesting-period]",
		Short: "Broadcast message CreateAirdropCampaign",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argStartTime := args[2]
			argEndTime := args[3]
			argLockupPeriod := args[4]
			argVestingPeriod := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
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
			lockupPeriod, err := time.ParseDuration(argLockupPeriod)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}

			vestingPeriod, err := time.ParseDuration(argVestingPeriod)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}

			msg := types.NewMsgCreateAirdropCampaign(
				clientCtx.GetFromAddress().String(),
				argName,
				argDescription,
				startTime,
				endTime,
				lockupPeriod,
				vestingPeriod,
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
